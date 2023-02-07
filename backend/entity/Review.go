package entity

import (
	"time"

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

	Comment string `valid:"stringlength(0|200)~Comment length must be between 0 - 200,required~Comment not blank,"`
	Star int
	Reviewdate time.Time
	Reviewimage string

	// CustomerID ทำหน้าที่เป็น FK
	CustomerID *uint
	Customer   Customer `gorm:"references:ID"`

	// SystemworkID ทำหน้าที่เป็น FK
	SystemworkID *uint
	Systemwork   Systemwork `gorm:"references:ID"`

	// DepartmentID ทำหน้าที่เป็น FK
	DepartmentID *uint
	Department   Department `gorm:"references:ID"`

}