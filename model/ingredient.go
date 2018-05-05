package model

// Ingredient - Zutat
type Ingredient struct {
	ID          uint      `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Name        string    `gorm:"column:name;type:varchar(64);not null"`
	Description string    `gorm:"column:description;not null"`
	Glutenfree  bool      `gorm:"column:gluten_free;not null"`
	Bio         bool      `gorm:"column:bio;not null"`
	Vegetarian  bool      `gorm:"column:vegetarian;not null"`
	Vegan       bool      `gorm:"column:vegan;not null"`
	Products    []Product `gorm:"many2many:products_ingredients;"`
}
