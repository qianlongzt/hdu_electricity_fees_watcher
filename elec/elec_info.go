package elec

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

//Info 电费信息
type Info struct {
	FreshTime string
	Roaming   string
	Community string
	Building  string
	Floor     string
	Room      string
}

var (
	//http://wap.xt.beescrm.com 这个域名也可以，不过看起来HTTPS 更安全
	elecTemp = "https://miniprogram.greencampus.cc/base/electricityHd/queryResult/ele_id/7/community_id/57/building_id/%s/floor_id/%s/room_id/%s/flag/1"

	re = regexp.MustCompile(`(?m).*<div class="field">截至.*<b>(.*)</b></font></div>\s*<div.*>\s*<p>实际所剩金额：<span class="price" style="color:#2d9fd3">(.*)</span></p>\s*<p>园区：(.*)</p>\s*<p>楼幢：(.*)</p>\s*<p>楼层：(.*)</p>\s*<p>寝室号：(.*)</p>\s*</div>.*`)
)

func buildURL(buildingID string, floorID string, roomID string) string {
	return fmt.Sprintf(elecTemp, buildingID, floorID, roomID)
}

func GetInfo(buildingID string, floorID string, roomID string) (ret Info, err error) {
	url := buildURL(buildingID, floorID, roomID)

	timeout := time.Duration(3 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	r, err := client.Get(url)
	if err != nil {
		return ret, err
	}
	defer r.Body.Close()
	t, err2 := ioutil.ReadAll(r.Body)
	if err2 != nil {
		return ret, err2
	}
	return getElectInfoFromString(string(t))
}

func getElectInfoFromString(raw string) (ret Info, err error) {
	data2 := re.FindAllStringSubmatch(raw, -1)
	if len(data2) != 1 {
		return ret, errors.New("no match")
	}
	if len(data2[0]) != 7 {
		return ret, errors.New("match length error")
	}
	data := data2[0]
	ret.FreshTime = data[1]
	ret.Roaming = data[2]
	ret.Community = data[3]
	ret.Building = data[4]
	ret.Floor = data[5]
	ret.Room = data[6]
	return
}
