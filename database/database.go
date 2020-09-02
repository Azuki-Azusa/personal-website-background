package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // run driver
)

var Db *gorm.DB

func GetInstance() *gorm.DB {
	if Db == nil {
		fmt.Println("111")
		Db = &gorm.DB{}
	}
	return Db
}

// Database run.
func init() {
	var err error
	// Db, err = gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/personal_website?charset=utf8&parseTime=True&loc=Local")
	Db, err = gorm.Open("mysql", "azuki:orangestar@tcp(database-1.cpxxkljx2bm4.ap-northeast-1.rds.amazonaws.com:3306)/personal_website?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	Db.LogMode(true)
}
