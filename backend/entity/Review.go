package entity

import (
	"time"

	"gorm.io/gorm"
)


// สร้างตารางชื่อ Department
type Systemwork struct {
	gorm.Model
	Name string
	// ส่ง Department_id ไปตาราง Review เพื่อเป็น foreignKey
	Review []Review `gorm:"foreignKey:DepartmentID"`
}


// สร้างตารางชื่อ Review เป็นตารางหลัก
type Review struct {
	gorm.Model

	Comment string
	Start int
	Reviewdate time.Time


	// CustomerID ทำหน้าที่เป็น FK
	CustomerID *uint
	Customer   Customer `gorm:"references:ID"`

	// CustomerID ทำหน้าที่เป็น FK
	SystemworkID *uint
	Systemwork   Systemwork `gorm:"references:ID"`

	// DepartmentID ทำหน้าที่เป็น FK
	DepartmentID *uint
	Department   Department `gorm:"references:ID"`
}