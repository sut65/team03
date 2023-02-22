package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
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

	CheckInTime  time.Time `valid:"required~Please select check-in time, NowP~Invalid check-in time (do not be in the past), NowF~Invalid check-in time (do not be in the future)"`
	CheckOutTime time.Time `valid:"IsAfterCheckIn~check out time must be after in check in time"`

	CheckInOutStatusID *uint            `valid:"required~Please select status"`
	CheckInOutStatus   CheckInOutStatus `valid:"-" gorm:"references:ID"`

	EmployeeID *uint    `valid:"required~Please Signin"`
	Employee   Employee `valid:"-" gorm:"references:ID"`
}

func init() {
	// govalidator.CustomTypeTagMap.Set("IsAfterAndPresent", func(i interface{}, context interface{}) bool {
	// 	t := i.(time.Time)
	// 	return t.After(time.Now().AddDate(0, 0, -1)) // today ->
	// })

	govalidator.CustomTypeTagMap.Set("NowP", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.After(time.Now().Add(time.Second * -599))
	})

	govalidator.CustomTypeTagMap.Set("NowF", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.Before(time.Now().Add(time.Second * 599))
	})

	// govalidator.CustomTypeTagMap.Set("Future", func(i interface{}, context interface{}) bool {
	// 	t := i.(time.Time)
	// 	return t.Before(time.Now().Add(time.Second * 599))
	// })

	govalidator.CustomTypeTagMap.Set("IsAfterCheckIn", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		cio := context.(CheckInOut)
		return t.After(cio.CheckInTime.Add(time.Second)) //Start ->
	})
}
