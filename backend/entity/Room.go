package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

//	type Staff struct {
//		gorm.Model
//		Name     string
//		Email    string `gorm:"uniqueIndex"`
//		Password string
//		Rooms    []Room `gorm:"foreignKey:StaffID"`
//	}
type RoomType struct {
	gorm.Model
	Size    string
	Bedsize string
	Bedtype string
	Amount  int
	Rooms   []Room `gorm:"foreignKey:RoomTypeID"`
}
type RoomZone struct {
	gorm.Model
	Name  string
	Rooms []Room `gorm:"foreignKey:RoomZoneID"`
}
type State struct {
	gorm.Model
	Name  string
	Rooms []Room `gorm:"foreignKey:StateID"`
}
type Room struct {
	gorm.Model
	Amount  int       `valid:"required~กรุณากรอกราคา, range(0|9223372036854775807)~กรุณากรอกราคาเป็นจำนวนเต็มบวก"`
	Room_No string    `valid:"matches(^[A-D]\\d{2}$)~หมายเลขห้องต้องขึ้นต้นด้วย A-D ตามด้วยตัวเลข 2 หลัก, required~กรุณากรอกหมายเลขห้อง"`
	Time    time.Time `valid:"DelayNow3Min~วันที่และเวลาไม่ถูกต้อง"`

	//StaffID ทำหน้าที่เป็น FK
	// StaffID *uint
	// Staff   Staff
	EmployeeID *uint
	Employee   Employee `valid:"-" gorm:"references:id"`

	//RoomTypeID ทำหน้าที่เป็น FK
	RoomTypeID *uint
	RoomType   RoomType `valid:"-" gorm:"references:id"`

	//ZoneID ทำหน้าที่เป็น FK
	RoomZoneID *uint
	RoomZone   RoomZone `valid:"-" gorm:"references:id"`

	//StateID ทำหน้าที่เป็น FK
	StateID *uint
	State   State `valid:"-" gorm:"references:id"`

	Bookings  []Booking   `gorm:"foreignKey:RoomID"`
	Checkroom []Checkroom `gorm:"foreignKey:RoomID"`
}

func init() {
	govalidator.CustomTypeTagMap.Set("IsFuture", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.After(time.Now())
	})

	govalidator.CustomTypeTagMap.Set("IsPresent", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.After(time.Now().AddDate(0, 0, -1)) && t.Before(time.Now().AddDate(0, 0, 1))
	})

	govalidator.CustomTypeTagMap.Set("IsPast", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.Before(time.Now())
	})
	govalidator.CustomTypeTagMap.Set("IsnotPast", func(i interface{}, o interface{}) bool {
		t := i.(time.Time)
		// ย้อนหลังไม่เกิน 1 วัน
		return t.After(time.Now().AddDate(0, 0, -1))
	})
	govalidator.CustomTypeTagMap.Set("DelayNow3Min", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		a := t.After(time.Now().Add(-3 * time.Minute))
		b := t.Before(time.Now().Add(+3 * time.Minute))
		sts := a && b
		println(a)
		println(b)
		return sts
	})
}
