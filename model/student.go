package model

// Student - Student
type Student struct {
	ID       uint   `gorm:"column:id;primary_key;AUTO_INCREMENT"` // MA
	MemberID uint   `gorm:"column:member_id"`
	Member   Member `gorm:"ForeignKey:ID"`
	Course   string `gorm:"column:course;not null"`
}
