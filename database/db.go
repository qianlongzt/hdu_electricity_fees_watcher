package database

import (
	"encoding/json"
	"log"

	scribble "github.com/nanobox-io/golang-scribble"
)

//UserInfo 用户订阅信息
type UserInfo struct {
	UserID       string `json:"user_id"`
	Room         string `json:"room"`
	IsSubscribed bool   `json:"is_subscribed"`
	MinLevel     int    `json:"min_level"`
}

//DB json 数据库
type DB struct {
	db     *scribble.Driver
	dbname string
}

//NewDb 新数据库
func NewDb(dir string) (*DB, error) {
	d, err := scribble.New(dir, nil)
	if err != nil {
		return nil, err
	}
	return &DB{
		db:     d,
		dbname: "db",
	}, nil
}

//Write 写记录
func (d *DB) Write(id string, data UserInfo) error {
	return d.db.Write(d.dbname, id, data)
}

//Read 读记录
func (d *DB) Read(id string) (*UserInfo, error) {
	u := UserInfo{}
	err := d.db.Read(d.dbname, id, &u)
	return &u, err
}

//Sub 更新用户sub信息
func (d *DB) Sub(id string, sub bool, min int) error {
	u, err := d.Read(id)
	if err != nil {
		return err
	}
	d.Delete(id)
	u.IsSubscribed = sub
	u.MinLevel = min
	return d.Write(id, *u)
}

//Delete 删除记录
func (d *DB) Delete(id string) error {
	return d.db.Delete(d.dbname, id)
}

//ReadAllSubed 获取所有订阅用户信息
func (d *DB) ReadAllSubed() []UserInfo {
	us := d.ReadAll()
	re := []UserInfo{}
	for _, u := range us {
		if u.IsSubscribed {
			re = append(re, u)
		}
	}
	return []UserInfo{}
}

//ReadAll 获取所有用户信息
func (d *DB) ReadAll() []UserInfo {
	rs, err := d.db.ReadAll(d.dbname)
	if err != nil {
		return []UserInfo{}
	}
	re := []UserInfo{}
	for _, f := range rs {
		tmp := UserInfo{}
		if err := json.Unmarshal([]byte(f), &tmp); err != nil {
			log.Println("Error", err)
			continue
		}
		re = append(re, tmp)
	}
	return re
}
