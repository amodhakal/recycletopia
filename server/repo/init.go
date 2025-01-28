package repo

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

const SQLITE_DB = "recycletopia.db"

func setupAdmin() {
	var admins []User
	db.Where("name = ?", "Admin").Find(&admins)

	if len(admins) != 0 {
		return
	}

	admin := User{Name: "Admin", Email: "admin@test.com", Password: "admin123"}
	result := db.Create(&admin)

	if result.Error != nil {
		log.Fatal("Couldn't create admin", result.Error)
	}

	println("Created admin")
}

func InitDB() {
	gormDB, err := gorm.Open(sqlite.Open(SQLITE_DB), &gorm.Config{})

	if err != nil {
		log.Fatal("Couldn't intialize DB: ", err)
	}

	db = gormDB

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Couldn't migrate user table")
	}

	setupAdmin()
}
