package entity

import (
	"time"

	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	Type string
	// ส่งไป
	CHK_Payments []CHK_Payment `gorm:"foreignKey:StatusID"`
}

type CHK_Payment struct {
	gorm.Model
	//รับเข้ามา
	// PaymentID *uint
	// Payment   Payment `gorm:"references:id"`
	//รับเข้ามา
	StatusID    *uint
	Status      Status `gorm:"references:id"`
	Date_time   time.Time
	Amount      *uint
	Description string
	//รับเข้ามา
	CustomerID *uint
	Customer   Customer `gorm:"references:id"`
}
