package model

// Admin
type Admin struct {
	User   User `gorm:"ForeignKey:UserID"`
	UserID uint `gorm:"column:user_id"`
}
