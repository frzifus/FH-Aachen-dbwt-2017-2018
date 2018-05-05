package model

// Employee - Mitarbeiter
type Employee struct {
	ID          uint   `gorm:"column:id;primary_key;AUTO_INCREMENT"` // MA
	MemberID    uint   `gorm:"column:member_id"`
	Member      User   `gorm:"ForeignKey:ID"`
	PhoneNumber int    `gorm:"column:phone_number;not null"`
	Office      string `gorm:"office;not null"`
}
