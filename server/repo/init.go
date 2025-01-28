package repo

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

const SQLITE_DB = "recycletopia.db"

func InitDB() {
	gormDB, err := gorm.Open(sqlite.Open(SQLITE_DB), &gorm.Config{})

	if err != nil {
		log.Fatal("Couldn't intialize DB: ", err)
	}

	db = gormDB
}
