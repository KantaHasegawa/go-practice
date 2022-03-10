package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase(dbDriver string, dbAddress string) *sql.DB {
	db, err := sql.Open(dbDriver, dbAddress)
	if(err != nil){
		fmt.Println(err.Error())
	}
	return db
}
