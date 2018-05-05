package model

import (
	"time"
)

// Guest - Gast
type Guest struct {
	UserID     uint      `gorm:"column:user_id;primary_key"`
	User       User      `gorm:"ForeignKey:UserID"`
	Reason     string    `gorm:"column:reason;not null"`
	ExpiryDate time.Time `gorm:"column:expiry_date;default:CURRENT_DATE"`
}
