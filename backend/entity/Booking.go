package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
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
	BranchID *uint  `valid:"required~กรุณาเลือกสาขาของโรงแรม"`
	Branch   Branch `valid:"-" gorm:"references:id"`
	//รับเข้ามา
	RoomID *uint `valid:"required~กรุณาเลือกห้องพัก"`
	Room   Room  `valid:"-" gorm:"references:id"`

	Start time.Time `valid:"required~กรุณาเลือกเวลาเริ่มเข้าพัก, IsAfterAndPresent~เวลาในการเข้าพักไม่ถูกต้อง(ห้ามเป็นอดีต)"`
	Stop  time.Time `valid:"required~กรุณาเลือกวันที่สิ้นสุดการพัก, IsAfterStartOneDay~เวลาสิ้นสุดการพักต้องอยู่หลังวันเข้าพักอย่างน้อย 1 วัน"`
	//รับเข้ามา
	CustomerID *uint    `valid:"required~กรุณาเข้าสู่ระบบ"`
	Customer   Customer `valid:"-" gorm:"references:id"`

	CheckInOut []CheckInOut `gorm:"foreignKey:BookingID"`
}

func init() {
	govalidator.CustomTypeTagMap.Set("IsAfterAndPresent", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.After(time.Now().AddDate(0, 0, -1)) // today ->
	})

	govalidator.CustomTypeTagMap.Set("IsAfterStartOneDay", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		b := context.(Booking)
		return t.After(b.Start) //Start ->
	})
}
