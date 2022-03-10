package utils

import (
	"database/sql"
	"fmt"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase(dbDriver string, dbAddress string) *gorp.DbMap {
	db, err := sql.Open(dbDriver, dbAddress)
	if(err != nil){
		fmt.Println(err.Error())
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	return dbmap
}
