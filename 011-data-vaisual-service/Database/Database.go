package Database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

//
//func GetDB() *gorm.DB  {
//	return DB
//}

func Migrate(dsn string, models ...interface{}) *gorm.DB {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect Database")
	}
	//defer db.Close()
	db.AutoMigrate(models...)
	return db
}

func init() {

	fmt.Println("Init Database")
}
