package model

// Member - Fh-Angehoerige
type Member struct {
	UserID uint `gorm:"column:user_id;primary_key"`
	User   User `gorm:"ForeignKey:UserID"`
}
