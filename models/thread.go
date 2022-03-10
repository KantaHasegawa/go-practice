package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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

func (thread Thread) Get() *sql.DB {
	db, err := sql.Open("mysql","root:@(localhost:3306)/go_practice")
	if(err != nil){
		fmt.Println(err.Error())
	}
	return db
}
