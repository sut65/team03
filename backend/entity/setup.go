package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db
}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("Team03.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		//จัดการข้อมูลพนักงาน
		&UserRole{},
		&Signin{},
		&Officer{},
		&Department{},
		&Position{},
		&Location{},
		&Employee{},
		&RoomType{},
		&RoomZone{},
		&State{},
		&Room{},
	)
	db = database

	SetupIntoDatabase(db)
}
