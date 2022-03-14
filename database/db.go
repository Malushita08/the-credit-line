package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "gorm.io/gorm"
)

var (
	DbSession *gorm.DB
	err       error
)

func ConnectDB() (*gorm.DB, error) {
	//SQLITE3
	//DbSession, err = gorm.Open("sqlite3", "./database/credit-line.sqlite3")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//return DbSession, nil

	//MYSQL
	USER := "root"
	PASS := "123456"
	HOST := "localhost"
	PORT := "3308"
	DBNAME := "creditLine"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", USER, PASS, HOST, PORT)
	DbSession, err = gorm.Open("mysql", dsn)
	_ = DbSession.Exec("CREATE DATABASE IF NOT EXISTS " + DBNAME + ";")
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)

	DbSession, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	return DbSession, nil

}
