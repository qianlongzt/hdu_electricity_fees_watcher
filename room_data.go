package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	getRoomListURL     = "https://miniprogram.greencampus.cc/base/common/getRoomList"
	getBuildingListURL = "https://miniprogram.greencampus.cc/base/common/getBuildingList"
	getFloorListURL    = "https://miniprogram.greencampus.cc/base/common/getFloorList"
)

type builidng struct {
	BuildingID   string `json:"building_id"`
	BuildingName string `json:"building_name"`
}

type floor struct {
	FloorID   string `json:"floor_id"`
	FloorName string `json:"floor_name"`
}

type room struct {
	RoomID   string `json:"room_id"`
	RoomName string `json:"room_name"`
}

//GetIDs 获取学校的信息
//raw string `12号楼南,1层,101`
func GetIDs(raw string) (buildingID, FloorID, RoomID string) {
	client := newHTTPTimeoutClient()
	id := "57" //57 下沙校区，硬编码

	d := strings.Split(raw, ",")
	if len(d) != 3 {
		return
	}
	b, f, r := d[0], d[1], d[2]

	bs := getBuildings(client, id)
	for _, t := range bs {
		if t.BuildingName == b {
			id, buildingID = t.BuildingID, t.BuildingID
			break
		}
	}
	fs := getFloors(client, id)
	for _, t := range fs {
		if t.FloorName == f {
			id, FloorID = t.FloorID, t.FloorID
			break
		}
	}
	rs := getRooms(client, id)
	for _, t := range rs {
		if t.RoomName == r {
			id, RoomID = t.RoomID, t.RoomID
			break
		}
	}
	return
}

func newHTTPTimeoutClient() *http.Client {
	timeout := time.Duration(3 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	return &client
}

func getBuildings(client *http.Client, id string) []builidng {
	str := httpPost(client, getBuildingListURL, id)

	ret := struct {
		Data []builidng `json:"data"`
	}{}
	json.Unmarshal([]byte(str), &ret)
	return ret.Data
}

func getFloors(client *http.Client, id string) []floor {
	str := httpPost(client, getFloorListURL, id)
	ret := struct {
		Data []floor `json:"data"`
	}{}
	json.Unmarshal([]byte(str), &ret)
	return ret.Data
}

func getRooms(client *http.Client, id string) []room {
	str := httpPost(client, getRoomListURL, id)
	ret := struct {
		Data []room `json:"data"`
	}{}
	json.Unmarshal([]byte(str), &ret)
	return ret.Data

}
func httpPost(client *http.Client, link, id string) string {
	data := url.Values{}
	data.Set("id", id)
	r, err := (*client).PostForm(link, data)
	if err != nil {
		return ""
	}
	defer r.Body.Close()
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return ""
	}
	return string(d)
}
