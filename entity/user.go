package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName     string
	LastName      string
	Email         string
	BirthDay      time.Time
	CitizenID     uint
	Walletaccount string
	RoleID        *uint
	Role          Role
}

type Role struct {
	gorm.Model
	Name  string
	Users []User `gorm:"foreignKey:RoleID"`
}
