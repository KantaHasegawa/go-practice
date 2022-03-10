package models

import (
	"encoding/json"
	"errors"
)

type Thread struct {
	Id int `json:"id"`
	Title string `json:"title"`
	CreatedAt string `json:"createdAt"`
}

func (thread Thread) GetAll() ([]byte, error) {
	thread1 := Thread{1,"first title", "100000"}
	thread2 := Thread{2,"second title", "200000"}
	threads := []Thread{thread1, thread2}
	result, err := json.Marshal(threads)
	if err != nil {
		err := errors.New("json format error:" + err.Error())
		return nil, err
	}
	return result, err
}
