package entity

import (
	"time"

	"gorm.io/gorm"
)

type Branch struct {
	gorm.Model
	B_name string
	// ส่งไป
	Bookings []Booking `gorm:"foreignKey:BranchID"`
}

type Booking struct {
	gorm.Model
	//รับเข้ามา
	BranchID *uint
	Branch   Branch `gorm:"references:id"`
	//รับเข้ามา
	RoomID *uint
	Room   Room `gorm:"references:id"`

	Start time.Time
	Stop  time.Time
	//รับเข้ามา
	CustomerID *uint
	Customer   Customer `gorm:"references:id"`

	CheckInOut []CheckInOut `gorm:"foreignKey:BookingID"`
}
