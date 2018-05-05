package model

import (
	"time"
)

// Order - Bestellung
type Order struct {
	ID      uint      `gorm:"column:ID;AUTO_INCREMENT;primary_key"`
	Time    time.Time `gorm:"column:time;default:CURRENT_DATE"`
	UserID  uint      `gorm:"column:user_id"`
	User    User      `gorm:"ForeignKey:ID"`
	Product []Product `gorm:"many2many:orders_products;"`
}
