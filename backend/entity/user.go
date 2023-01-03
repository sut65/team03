package entity

import (
	"gorm.io/gorm"
)


type UserRole struct { //สร้างเพื่อเอาไว้แยก Officer, Employee with Customer
	gorm.Model //ไลเบอร์รัี่สำเร็จรูป เอาไว้ใช้และใน model จะมี ไอดี ลบ อัพเดพ สร้าง
	RoleName string

	//[] อาเรย์

	Signin []Signin `gorm:"foreignKey:UserRoleID"`
	// Officerlogin []Officer `gorm:"foreignKey:UserRoleID"`
	// Employeelogin []Employee `gorm:"foreignKey:UserRoleID"`
	// Customerlogin []Customer `gorm:"foreignKey:UserRoleID"`
}

type Signin struct {
	gorm.Model
	Username string
	Password string

	UserRoleID *uint
	UserRole   UserRole `gorm:"references:ID"`
}