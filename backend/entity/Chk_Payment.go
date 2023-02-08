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
	PaymentID *uint   `valid:"required~Please select Payment"`
	Payment   Payment `valid:"-" gorm:"references:id"`
	//รับเข้ามา
	CHK_PaymentStatusID *uint             `valid:"required~Please select Status"`
	CHK_PaymentStatus   CHK_PaymentStatus `valid:"-" gorm:"references:id"`
	Date_time           time.Time         `valid:"required~Please select date time"  govalidator:"func=ValidateDateTimeNotFuture"`
	Amount              int               `valid:"required~Please input Amount"`
	Description         string
	//รับเข้ามา
	EmployeeID *uint    `valid:"required~Please Signin"`
	Employee   Employee `valid:"-" gorm:"references:id"`
}
