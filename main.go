package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/qianlongzt/hdu_electricity_fees_watcher/database"
	"github.com/qianlongzt/hdu_electricity_fees_watcher/elec"
)

var (
	flagToken  = flag.String("token", "", "telegram bot api token")
	flagDbpath = flag.String("dir", "./dir", "directory location to databse")
	flagDebug  = flag.Bool("debug", false, "bot debug")
)

func main() {
	flag.Parse()
	if *flagToken == "" {
		log.Panic("telegram api token can't empty")
	}
	db, err := database.NewDb(*flagDbpath)
	if err != nil {
		log.Panic(err)
	}
	bot, err := tgbotapi.NewBotAPI(*flagToken)
	if err != nil {
		log.Panic(err)
	}
	if *flagDebug {
		bot.Debug = true
	}

	CronRun(db, bot)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		if !update.Message.IsCommand() {
			continue
		}

		text := ""
		args := update.Message.CommandArguments()
		userid := strconv.Itoa(update.Message.From.ID)
		switch update.Message.Command() {
		case "help", "start":
			text = "\n/elec - 电费查询\n/room - 设置寝室 如 1号楼南,1层,101\n\n/sub - 订阅通知\n\n/unsub - 取消订阅\n/clear - 清空寝室记录"
		case "elec":
			uif, err := db.Read(userid)
			log.Println(uif, err)
			if err != nil {
				text = "你没有保存寝室信息"
				break
			}
			info, err := elec.GetInfo(elec.GetIDs(uif.Room))
			if err != nil {
				text = "获取寝室电费时发生错误"
				break
			}
			text = fmt.Sprintf("寝室：%v,\n更新时间: %v,\n余额：%v", info.Community+info.Building+info.Floor+info.Room, info.FreshTime, info.Roaming)
			if uif.IsSubscribed {
				text += fmt.Sprintf("\n\n你订阅了提醒，将在每日 21 点，且余额低于 %v 元 提醒你", uif.MinLevel)
			}
		case "room":
			bid, fid, rid := elec.GetIDs(args)
			if bid == "" || fid == "" || rid == "" {
				text = "错误的寝室号"
				break
			}
			err := db.Write(userid, database.UserInfo{
				UserID:       userid,
				Room:         args,
				IsSubscribed: false,
			})
			if err != nil {
				text = "保存寝室时有错误发生，" + err.Error()
				break
			}
			text = "保存成功"

		case "sub":
			min := 15
			args = strings.Trim(args, " ")
			if args != "" {
				i, err := strconv.Atoi(args)
				if err != nil {
					text = "必须是整数"
					break
				}
				if i <= 0 {
					text = "订阅余额必须大于0"
					break
				}
				min = i
			}

			err = db.Sub(userid, true, min)
			if err != nil {
				text = "订阅出错" + err.Error()
				break
			}
			text = fmt.Sprintf("订阅成功，将在每天21点，且低于 %v 元提醒你", min)
		case "unsub":
			err := db.Sub(userid, false, 0)
			if err != nil {
				text = "取消订阅出错" + err.Error()
				break
			}
			text = "取消订阅成功"
		case "clear":
			err := db.Delete(userid)
			if err != nil {
				text = "清除寝室记录时有错误发生，" + err.Error()
			} else {
				text = "清除记录成功"
			}
		default:
			text = "未知命令"
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}
