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
	Location []Location `gorm:"foreignKey:DepartmentID"`
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

	DepartmentID *uint `valid:"-"`
	Department   Department `gorm:"references:ID"`
	// ส่ง Location_id ไปตาราง Employee เพื่อเป็น foreignKey
}

// สร้างตารางชื่อ Employee เป็นตารางหลัก
type Employee struct {
	gorm.Model

	// รับข้อมูล PersonalID ที่ไม่ซ้ำกัน
	PersonalID   string `gorm:"uniqueIndex" valid:"matches(^([0-9]{13})$)~PersonalID is not valid,required~PersonalID not blank"`
	Employeename string `valid:"required~Name not blank"`
	Email        string `gorm:"uniqueIndex" valid:"email~Email is not valid,required~Email not blank"`

	Eusername string `valid:"matches(^[E][A-Z][a-zA-Z]+$)~Username must be is Begin with E and The second letter must start with A-Z and must not number,required~Username not blank"`
	Password  string `valid:"minstringlength(6)~Password must be more than or equal to 6 characters,matches([A-Z])~Password must contain at least 1 character A-Z.,required~Password not blank"`

	SigninID *uint `valid:"-"`
	Signin   Signin `gorm:"references:ID" valid:"-"` 

	Salary      uint64
	Phonenumber string `valid:"matches(^(0)([0-9]{9})$)~Phonenumber is not valid,required~Tel not blank"`
	Gender      string
	DateOfBirth time.Time
	YearOfStart time.Time
	Address     string `valid:"required~Address not blank"`

	// OfficerID ทำหน้าที่เป็น FK
	OfficerID *uint
	Officer   Officer `gorm:"references:ID"`

	// DepartmentID ทำหน้าที่เป็น FK
	DepartmentID *uint `valid:"-"`
	Department   Department `gorm:"references:ID" valid:"-"` 

	// PositionID ทำหน้าที่เป็น FK
	PositionID *uint `valid:"-"`
	Position   Position `gorm:"references:ID" valid:"-"`

	Rooms []Room `gorm:"foreignKey:EmployeeID"`
	// ส่ง EmployeeID ไปตาราง CheckInOut เพื่อเป็น foreignKey
	CheckInOut []CheckInOut `grom:"foreignKey:EmployeeID"`
	// ส่ง EmployeeID ไปตาราง CHK_Payment เพื่อเป็น foreignKey
	CHK_Payments []CHK_Payment `gorm:"foreignKey:EmployeeID"`
	// ส่ง EmployeeID ไปตาราง Storage เพื่อเป็น foreignKey
	Storages  []Storage   `gorm:"foreignKey:EmployeeID"`
	Checkroom []Checkroom `gorm:"foreignKey:EmployeeID"`
}
