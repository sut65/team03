package entity

import (
	"time"

	"gorm.io/gorm"
)

// สร้างตารางชื่อ Officer
type Officer struct {
	gorm.Model

	SigninID *uint
	Signin   Signin `gorm:"references:ID"`

	Officername string
	Tel         string
	// ส่ง admin_id ไปตาราง Employee เพื่อเป็น foreignKey
	Employee []Employee `gorm:"foreignKey:OfficerID"`
}

// สร้างตารางชื่อ Department
type Department struct {
	gorm.Model
	Name string
	// ส่ง Department_id ไปตาราง Employee เพื่อเป็น foreignKey
	Employee []Employee `gorm:"foreignKey:DepartmentID"`
	Review   []Review   `gorm:"foreignKey:DepartmentID"`
}

// สร้างตารางชื่อ Position
type Position struct {
	gorm.Model
	Name string
	// ส่ง Position _id ไปตาราง Employee เพื่อเป็น foreignKey
	Employee []Employee `gorm:"foreignKey:PositionID"`
}

// สร้างตารางชื่อ Location
type Location struct {
	gorm.Model
	Name string
	// ส่ง Location_id ไปตาราง Employee เพื่อเป็น foreignKey
	Employee []Employee `gorm:"foreignKey:LocationID"`
}

// สร้างตารางชื่อ Employee เป็นตารางหลัก
type Employee struct {
	gorm.Model

	// รับข้อมูล PersonalID ที่ไม่ซ้ำกัน
	PersonalID   uint64 `gorm:"uniqueIndex"`
	Employeename string `valid:"required~Name not blank"`
	Email        string `gorm:"uniqueIndex"`

	Eusername string
	Password  string

	SigninID *uint
	Signin   Signin `gorm:"references:ID"`

	Salary      uint64
	Phonenumber string
	Gender      string
	DateOfBirth time.Time
	YearOfStart time.Time
	Address     string

	// OfficerID ทำหน้าที่เป็น FK
	OfficerID *uint
	Officer   Officer `gorm:"references:ID"`

	// DepartmentID ทำหน้าที่เป็น FK
	DepartmentID *uint
	Department   Department `gorm:"references:ID"`

	// PositionID ทำหน้าที่เป็น FK
	PositionID *uint
	Position   Position `gorm:"references:ID"`

	// LocationID ทำหน้าที่เป็น FK
	LocationID *uint
	Location   Location `gorm:"references:ID"`

	Rooms []Room `gorm:"foreignKey:EmployeeID"`
	// ส่ง EmployeeID ไปตาราง CheckInOut เพื่อเป็น foreignKey
	CheckInOut []CheckInOut `grom:"foreignKey:EmployeeID"`
	// ส่ง EmployeeID ไปตาราง CHK_Payment เพื่อเป็น foreignKey
	CHK_Payments []CHK_Payment `gorm:"foreignKey:EmployeeID"`
	// ส่ง EmployeeID ไปตาราง Storage เพื่อเป็น foreignKey
	Storages  []Storage   `gorm:"foreignKey:EmployeeID"`
	Checkroom []Checkroom `gorm:"foreignKey:EmployeeID"`
}
