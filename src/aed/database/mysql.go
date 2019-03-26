package database

import (
	"database/sql"
	"github.com/gpmgo/gopm/modules/log"
	_"github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func init()  {
	var err error
	SqlDB, err = sql.Open("mysql","root:102030@tcp(127.0.0.1:3306)/aed")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
