package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/qianlongzt/hdu_electricity_fees_watcher/database"
	"github.com/qianlongzt/hdu_electricity_fees_watcher/elec"
	"github.com/robfig/cron/v3"
)

//CronRun 定时运行
func CronRun(db *database.DB, bot *tgbotapi.BotAPI) {
	fn := func() {
		us := db.ReadAllSubed()
		for _, u := range us {
			info, err := elec.GetInfo(elec.GetIDs(u.Room))
			text := ""
			if err != nil {
				log.Println("定时获取电费信息错误", err)
				continue
			}
			rom, err := strconv.ParseFloat(strings.Trim(info.Roaming, "元"), 32)
			if err != nil {
				log.Println("电费余额解析出错", err)
				continue
			}
			if rom > float64(u.MinLevel) {
				continue
			}

			text = fmt.Sprintf("寝室：%v,\n更新时间: %v,\n余额：%v", info.Community+info.Building+info.Floor+info.Room, info.FreshTime, info.Roaming)
			text += "\n\n 该交电费了！！！"
			uid, _ := strconv.Atoi(u.UserID)
			msg := tgbotapi.NewMessage(int64(uid), text)
			bot.Send(msg)
		}

	}
	c := cron.New()
	c.AddFunc("0 21 * * *", fn)
	c.Start()
}
