package models

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-gorp/gorp"
)

type Thread struct {
	Id int `json:"id"`
	Title string `json:"title"`
	CreatedAt string `json:"createdAt"`
}

func (thread Thread) GetAll(db *gorp.DbMap) ([]byte, error) {
	var threads []Thread
	_, err := db.Select(&threads, "select * from thread")
	if err !=nil {
		fmt.Println(err.Error())
	}
	result, err := json.Marshal(threads)
	if err != nil {
		err := errors.New("json format error:" + err.Error())
		return nil, err
	}
	return result, err
}

func (thread Thread) Get(id string, db *gorp.DbMap) ([]byte, error) {
	thread1 := Thread{}
	err := db.SelectOne(&thread1, "select * from thread where id=?", id)
	if err !=nil {
		fmt.Println(err.Error())
	}
	result, err := json.Marshal(thread1)
	if err != nil {
		err := errors.New("json format error:" + err.Error())
		return nil, err
	}
	return result, err
}
