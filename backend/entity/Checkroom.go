package entity

import (
	"time"
	"gorm.io/gorm"
)

type Damage struct {
	gorm.Model
	Description string
	Checkroom    []Checkroom `gorm:"foreignKey:DamageID"`
}

type StatusCR struct {
	gorm.Model
	S_Name string
	Checkroom    []Checkroom `gorm:"foreignKey:StatusID"`
}	

type Checkroom struct {
	gorm.Model
	//roomzone
	RoomID *uint
	Room   Room `gorm:"references:id"`
	//room
	ProductID *uint
	Product  Product `gorm:"references:id"`

	DamageID *uint
	Damage  Damage `gorm:"references:id"`

	StatusID *uint
	Status  StatusCR `gorm:"references:id"`

	Date time.Time

	EmployeeID *uint
	Employee  Employee `gorm:"references:id"`
	
	

}
