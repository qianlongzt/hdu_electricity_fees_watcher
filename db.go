package main

import (
	scribble "github.com/nanobox-io/golang-scribble"
)

//UserInfo 用户订阅信息
type UserInfo struct {
	UserID       string `json:"user_id"`
	Room         string `json:"room"`
	IsSubscribed bool   `json:"is_subscribed"`
}

//DB json 数据库
type DB struct {
	db *scribble.Driver
}

//NewDb 新数据库
func NewDb(dir string) (*DB, error) {
	d, err := scribble.New(dir, nil)
	if err != nil {
		return nil, err
	}
	return &DB{
		db: d,
	}, nil
}

//Write 写记录
func (d *DB) Write(id string, data UserInfo) error {
	return d.db.Write("db", id, data)
}

//Read 读记录
func (d *DB) Read(id string) (*UserInfo, error) {
	u := UserInfo{}
	err := d.db.Read("db", id, &u)
	return &u, err
}

//Delete 删除记录
func (d *DB) Delete(id string) error {
	return d.db.Delete("db", id)
}
