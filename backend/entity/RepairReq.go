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

type RepairStatus struct {
	gorm.Model
	Name      string
	RepairReq []RepairReq `gorm:"foreignKey:RepairStatusID"`
}

type RepairReq struct {
	gorm.Model
	Room uint
	// RoomID *uint
	// Room Room `gorm:"references:ID"`

	RepairTypeID *uint
	RepairType   RepairType `gorm:"references:ID"`

	Note string
	Time time.Time

	RepairStatusID *uint
	RepairStatus   RepairStatus `gorm:"references:ID"`

	User uint
	//UserID *uint
	//User User `gorm:"references:ID"`
}
