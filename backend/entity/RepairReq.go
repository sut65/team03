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

	RoomID *uint
	Room   Room `gorm:"references:ID"`

	RepairTypeID *uint
	RepairType   RepairType `gorm:"references:ID"`

	Note string
	Time time.Time

	CustomerID *uint
	Customer   Customer `gorm:"references:ID"`
}
