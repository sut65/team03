package entity

import (
	"time"

	"gorm.io/gorm"
)

// สร้างตารางชื่อ CheckInOutStatus
type CheckInOutStatus struct {
	gorm.Model
	Name string
	// ส่ง CheckInOutStatusID ไปตาราง CheckInOut เพื่อเป็น foreignKey
	CheckInOut []CheckInOut `gorm:"foreignKey:CheckInOutStatusID"`
}

// สร้างตารางชื่อ CheckInOut
type CheckInOut struct {
	gorm.Model

	// Booking uint
	BookingID *uint   `valid:"required~Please select booking"`
	Booking   Booking `valid:"-" gorm:"references:ID"`

	CheckInTime  time.Time `valid:"required~Please select check-in time"`
	CheckOutTime time.Time

	CheckInOutStatusID *uint            `valid:"required~Please select status"`
	CheckInOutStatus   CheckInOutStatus `valid:"-" gorm:"references:ID"`

	EmployeeID *uint    `valid:"required~Please Signin"`
	Employee   Employee `valid:"-" gorm:"references:ID"`
}
