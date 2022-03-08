package database

import (
	"fmt"
	"github.com/Malushita08/the-credit-line/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() {
	// NOTE: See we're using = to assign the global var
	//         	instead of := which would assign it only in this function
	db, err = gorm.Open("sqlite3", "./database/credit-line.sqlite3")
	//db, _ = gorm.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/database?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//Migraciones
	db.AutoMigrate(&models.Person{})
}

func MigrateDB() {
	//db.AutoMigrate(&models.Person{})
}
