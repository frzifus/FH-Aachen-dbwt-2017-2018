package model

import (
	"time"
)

// User - FH-Nutzer
type User struct {
	ID        uint      `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Active    bool      `gorm:"column:active;default:true"`
	Firstname string    `gorm:"column:firstname;not null"`
	Lastname  string    `gorm:"column:lastname;not null"`
	Mail      string    `gorm:"column:mail;unique"`
	Loginname string    `gorm:"column:loginname;unique"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_DATE"`
	LastLogin time.Time `gorm:"column:last_login"`
	Stretch   uint      `gorm:"column:stretch;not null"`
	Algo      string    `gorm:"column:algo;not null"`
	Salt      string    `gorm:"column:salt;type:varchar(32);not null"`
	Hash      string    `gorm:"column:hash;type:varchar(64);not null"`
}
