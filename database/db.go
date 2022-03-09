package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DbSession *gorm.DB
	err       error
)

func ConnectDB() (*gorm.DB, error) {
	DbSession, err = gorm.Open("sqlite3", "./database/credit-line.sqlite3")
	//DbSession, _ = gorm.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/database?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	//defer DbSession.Close()
	return DbSession, nil
}
