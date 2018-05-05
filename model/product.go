package model

// Product - Produkt
type Product struct {
	ID          int          `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Name        string       `gorm:"column:name;not null"`
	Description string       `gorm:"column:description;not null"`
	ImageID     uint         `gorm:"column:image_id"`
	Image       Image        `gorm:"ForeignKey:ImageID"`
	PriceID     uint         `gorm:"column:price_id"`
	Price       Price        `gorm:"ForeignKey:PriceID"`
	CategoryID  uint         `gorm:"column:category_id"`
	Category    Category     `gorm:"ForeignKey:CategoryID"`
	Orders      []Product    `gorm:"many2many:orders_products;"`
	Ingredients []Ingredient `gorm:"many2many:products_ingredients;"`
}

// IsVegan - Returns true if all product ingredients are vegan
func (p *Product) IsVegan() bool {
	for _, v := range p.Ingredients {
		if v.Vegan == false {
			return v.Vegan
		}
	}
	return true
}

// IsVegetarian - Returns true if all product ingredients are vegetarian
func (p *Product) IsVegetarian() bool {
	for _, v := range p.Ingredients {
		if v.Vegetarian == false {
			return v.Vegan
		}
	}
	return true
}
