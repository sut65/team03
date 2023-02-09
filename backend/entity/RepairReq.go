package entity

import (
	"time"

	"gorm.io/gorm"
)

type RepairType struct {
	gorm.Model
	Name      string
	RepairReq []RepairReq `gorm:"foreignKey:RepairTypeID"`
}

type RepairReq struct {
	gorm.Model

	RoomID *uint `valid:"required~Please select room"`
	Room   Room  `valid:"-" gorm:"references:ID"`

	RepairTypeID *uint      `valid:"required~Please select type of problem"`
	RepairType   RepairType `valid:"-" gorm:"references:ID"`

	Note string    `valid:"stringlength(0|200)~Comment length must be between 0 - 200,required~Please enter comment,"`
	Time time.Time `valid:"required~Please select time"`

	CustomerID *uint    `valid:"required~Please Signin"`
	Customer   Customer `valid:"-" gorm:"references:ID"`
}
