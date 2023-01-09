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
	Time time.Time

	//StaffID ทำหน้าที่เป็น FK
	// StaffID *uint
	// Staff   Staff
	EmployeeID *uint
	Employee   Employee

	//RoomTypeID ทำหน้าที่เป็น FK
	RoomTypeID *uint
	RoomType   RoomType

	//ZoneID ทำหน้าที่เป็น FK
	RoomZoneID *uint
	RoomZone   RoomZone

	//StateID ทำหน้าที่เป็น FK
	StateID *uint
	State   State

	Bookings []Booking `gorm:"foreignKey:RoomID"`
}
