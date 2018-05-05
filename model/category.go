package model

// Category - Kategorie
type Category struct {
	ID              uint      `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Designation     string    `gorm:"column:designation;not null"`
	UpperCategoryID uint      `gorm:"column:upper_category_id"`
	UpperCategory   *Category `gorm:"ForeignKey:UpperCaterogyID"`
	ImageID         uint      `gorm:"column:image_id"`
	Image           Image     `gorm:"ForeignKey:ImageID"`
}
