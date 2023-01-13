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
	BookingID *uint
	Booking   Booking `gorm:"references:ID"`

	CheckInTime  time.Time
	CheckOutTime time.Time

	CheckInOutStatusID *uint
	CheckInOutStatus   CheckInOutStatus `gorm:"references:ID"`

	EmployeeID *uint
	Employee   Employee `gorm:"references:ID"`
}
