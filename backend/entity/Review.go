package entity

import (
	"time"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)


// สร้างตารางชื่อ Systemwork
type Systemwork struct {
	gorm.Model
	Name string
	// ส่ง Systemwork_id ไปตาราง Review เพื่อเป็น foreignKey
	Review []Review `gorm:"foreignKey:SystemworkID"`
}


// สร้างตารางชื่อ Review เป็นตารางหลัก
type Review struct {
	gorm.Model

	Comment string `valid:"matches([a-zA-Z0-9ก-๙]$)~Comment no special characters,stringlength(0|200)~Comment length must be between 0 - 200,required~Comment not blank"`
	Star int
	Reviewdate time.Time
	Reviewimage string `valid:"image_valid~Invalid Image"`

	// CustomerID ทำหน้าที่เป็น FK
	CustomerID *uint `valid:"-"`
	Customer   Customer `gorm:"references:ID" valid:"-"`

	// SystemworkID ทำหน้าที่เป็น FK
	SystemworkID *uint `valid:"-"`
	Systemwork   Systemwork `gorm:"references:ID" valid:"-"`

	// DepartmentID ทำหน้าที่เป็น FK
	DepartmentID *uint `valid:"-"`
	Department   Department `gorm:"references:ID" valid:"-"`

}

func init() {
	govalidator.TagMap["image_valid"] = govalidator.Validator(func(str string) bool {
		return govalidator.Matches(str, "^(data:image(.+);base64,.+)$")
	})
}