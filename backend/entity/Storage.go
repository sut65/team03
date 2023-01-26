package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string
	Price    int
	Storages []Storage `gorm:"foreignKey:ProductID"`
}
type ProductType struct {
	gorm.Model
	Name     string
	Storages []Storage `gorm:"foreignKey:ProductTypeID"`
}
type Storage struct {
	gorm.Model
	Quantity int
	Time     time.Time

	//EmployeeID ทำหน้าที่เป็น FK
	EmployeeID *uint
	Employee   Employee

	//ProductID ทำหน้าที่เป็น FK
	ProductID *uint
	Product   Product

	//ProductTypeID ทำหน้าที่เป็น FK
	ProductTypeID *uint
	ProductType   ProductType
}
