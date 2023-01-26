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
		// ระบบห้องพัก
		&RoomType{},
		&RoomZone{},
		&State{},
		&Room{},
		//checkInOut
		&CheckInOutStatus{},
		&CheckInOut{},
		// แจ้งซ่อม
		&RepairType{},
		&RepairStatus{},
		&RepairReq{},
		//ระบบสมัครสมาชิก(ข้อมูลลูกค้า)
		&Customer{},
		&Gender{},
		&Province{},
		&Memberlevel{},
		// ระบบบริการเสริม
		&Food{},
		&Drink{},
		&Accessories{},
		// ระบบจองห้องพัก
		&Booking{},
		&Branch{},
		// ระบบตรวจสอบการชำระเงิน
		&CHK_Payment{},
		&CHK_PaymentStatus{},
		//Review---------
		&Review{},
		&Systemwork{},
		//ระบบคลังสินค้าห้องพัก
		&Product{},
		&ProductType{},
		&Storage{},
	)
	db = database

	SetupIntoDatabase(db)
}
