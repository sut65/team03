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

type Status struct {
	gorm.Model
	S_Name string
	Checkroom    []Checkroom `gorm:"foreignKey:StatusID"`
}	

type Checkroom struct {
	gorm.Model
	//roomzone
	RoomID *uint  `valid:"required~Please select Room"`
	Room   Room `valid:"-" gorm:"references:id"`
	//room
	ProductID *uint  `valid:"required~กรุณาระบุอุปกรณ์ที่ตรวจสอบในห้องพัก"`
	Product  Product `valid:"-" gorm:"references:id"`

	DamageID *uint  `valid:"required~กรุณาระบุความเสียหาของอุปกรณ์"`
	Damage  Damage `valid:"-" gorm:"references:id"`

	StatusID *uint 	`valid:"required~Please select Status"`
	Status  Status `valid:"-" gorm:"references:id"` 

	Date time.Time	`valid:"required~Please select Start date" govalidator:"func=ValidateStartBeforeStop, func=ValidateStopAfterStartOneDay()"`

	EmployeeID *uint  `valid:"required~Please select login"`
	Employee   Employee  `valid:"-" gorm:"references:id"`
	
}

/////////////////////////////////////////////////เพิ่ม Fk routes main //////////////////////////////////////
