package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name      string
	Price     int
	Storages  []Storage   `gorm:"foreignKey:ProductID"`
	Checkroom []Checkroom `gorm:"foreignKey:ProductID"`
}
type ProductType struct {
	gorm.Model
	Name     string
	Storages []Storage `gorm:"foreignKey:ProductTypeID"`
}
type Storage struct {
	gorm.Model
	ID       uint
	Quantity int       `valid:"required~กรุณากรอกจำนวนที่มากกว่าศูนย์, range(0|9223372036854775807)~กรุณากรอกจำนวนเป็นจำนวนเต็มบวก"`
	Time     time.Time `valid:"IsnotPast~วันที่และเวลาไม่ถูกต้อง"` // เป็นปัจจุบัน +- 3 นาที

	//EmployeeID ทำหน้าที่เป็น FK
	EmployeeID *uint
	Employee   Employee `valid:"-" gorm:"references:id"`

	//ProductID ทำหน้าที่เป็น FK
	ProductID *uint
	Product   Product `valid:"-" gorm:"references:id"`

	//ProductTypeID ทำหน้าที่เป็น FK
	ProductTypeID *uint
	ProductType   ProductType `valid:"-" gorm:"references:id"`
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
}
