package entity

import (
	"time"

	"gorm.io/gorm"
)

type CHK_PaymentStatus struct {
	gorm.Model
	Type string
	// ส่งไป
	CHK_Payments []CHK_Payment `gorm:"foreignKey:CHK_PaymentStatusID"`
}

type CHK_Payment struct {
	gorm.Model
	//รับเข้ามา
	// PaymentID *uint
	// Payment   Payment `gorm:"references:id"`
	//รับเข้ามา
	CHK_PaymentStatusID *uint
	CHK_PaymentStatus   CHK_PaymentStatus `gorm:"references:id"`
	Date_time           time.Time
	Amount              int
	Description         string
	//รับเข้ามา
	EmployeeID *uint
	Employee   Employee `gorm:"references:id"`
}
