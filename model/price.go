package model

// Price - Preis
type Price struct {
	ID       uint `gorm:"column:ID;AUTO_INCREMENT;primary_key"`
	Guest    uint `gorm:"column:guest"`
	Student  uint `gorm:"column:student"`
	Employee uint `gorm:"column:employee"`
}
