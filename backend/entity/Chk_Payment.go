package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
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
	PaymentID *uint   `gorm:"uniqueIndex" valid:"required~กรุณาเลือกรายการชำระเงิน"`
	Payment   Payment `valid:"-" gorm:"references:id"`
	//รับเข้ามา
	CHK_PaymentStatusID *uint             `valid:"required~กรุณาเลือกสถานะการชำระเงิน"`
	CHK_PaymentStatus   CHK_PaymentStatus `valid:"-" gorm:"references:id"`
	Date_time           time.Time         `valid:"required~กรุณาเลือกวันที่ที่ทำการชำระเงิน, IsBeforeAndPresent~เวลาในการชำระเงินไม่ถูกต้อง(ห้ามเป็นอนาคต)"`
	Amount              float64           `valid:"required~กรุณาระบุจำนวนเงิน, range(0|10000000000)~จำนวนเงินไม่สามารถติดลบได้"`
	Description         string
	//รับเข้ามา
	EmployeeID *uint    `valid:"required~กรุณาเข้าสู่ระบบ"`
	Employee   Employee `valid:"-" gorm:"references:id"`
}

func init() {
	govalidator.CustomTypeTagMap.Set("IsBeforeAndPresent", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.Before(time.Now().AddDate(0, 0, 1)) // <- today
	})
}
