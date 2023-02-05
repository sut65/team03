package entity

import (
	"time"

	"gorm.io/gorm"
)

//	type Staff struct {
//		gorm.Model
//		Name     string
//		Email    string `gorm:"uniqueIndex"`
//		Password string
//		Rooms    []Room `gorm:"foreignKey:StaffID"`
//	}
type RoomType struct {
	gorm.Model
	Size    string
	Bedsize string
	Bedtype string
	Rooms   []Room `gorm:"foreignKey:RoomTypeID"`
}
type RoomZone struct {
	gorm.Model
	Name  string
	Rooms []Room `gorm:"foreignKey:RoomZoneID"`
}
type State struct {
	gorm.Model
	Name  string
	Rooms []Room `gorm:"foreignKey:StateID"`
}
type Room struct {
	gorm.Model
	Room_No string
	Time    time.Time `valid:"required"`

	//StaffID ทำหน้าที่เป็น FK
	// StaffID *uint
	// Staff   Staff
	EmployeeID *uint `valid:"required"`
	Employee   Employee

	//RoomTypeID ทำหน้าที่เป็น FK
	RoomTypeID *uint `valid:"required"`
	RoomType   RoomType

	//ZoneID ทำหน้าที่เป็น FK
	RoomZoneID *uint `valid:"required"`
	RoomZone   RoomZone

	//StateID ทำหน้าที่เป็น FK
	StateID *uint `valid:"required"`
	State   State

	Bookings  []Booking   `gorm:"foreignKey:RoomID"`
	Checkroom []Checkroom `gorm:"foreignKey:RoomID"`
}
