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
	BranchID *uint  `valid:"required~กรุณาระบุสาขาของโรงแรม"`
	Branch   Branch `gorm:"references:id"`
	//รับเข้ามา
	RoomID *uint `valid:"required~กรุณาระบุห้องของโรงแรม"`
	Room   Room  `gorm:"references:id"`

	Start time.Time `valid:"required~กรุณาระบุเวลาเริ่มเข้าพัก"`
	Stop  time.Time `valid:"required~กรุณาระบุเวลาที่สิ้นสุดการพัก"`
	//รับเข้ามา
	CustomerID *uint    `valid:"required~กรุณาระบุชื่อผู้จอง"`
	Customer   Customer `gorm:"references:id"`

	CheckInOut []CheckInOut `gorm:"foreignKey:BookingID"`
}
