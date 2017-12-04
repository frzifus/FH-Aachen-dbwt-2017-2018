package model

import (
	"database/sql"
	"time"
)

// User - FH-Nutzer
type User struct {
	ID        uint      `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Active    bool      `gorm:"column:active;default:true"`
	Firstname string    `gorm:"column:firstname;not null"`
	Lastname  string    `gorm:"column:lastname;not null"`
	Mail      string    `gorm:"column:mail;unique"`
	Loginname string    `gorm:"column:loginname;unique"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_DATE"`
	LastLogin time.Time `gorm:"column:last_login"`
	Stretch   uint      `gorm:"column:stretch;not null"`
	Algo      string    `gorm:"column:algo;not null"`
	Salt      string    `gorm:"column:salt;type:varchar(32);not null"`
	Hash      string    `gorm:"column:hash;type:varchar(64);not null"`
    MemberID  uint
    GuestID   uint
    Orders    []Order   `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
}

// Member - Fh-Angehoerige
type Member struct {
	ID     uint `gorm:"column:id"`
    User   User `gorm:"ForeignKey:ID;AssociationForeignKey:ID"`

    StudentID  uint
    EmployeeID uint
}

// Guest - Gast
type Guest struct {
    ID         uint      `gorm:"column:id;not null"`
	Reason     string    `gorm:"column:reason;not null"`
	ExpiryDate time.Time `gorm:"column:expiry_date;default:CURRENT_DATE"`
}

// Student - Student
type Student struct {
    ID        uint 
	StudentID uint   `gorm:"column:student_id;not null"`
	Course    string `gorm:"column:course;not null"`
    Member    Member `gorm:"ForeignKey:StudentID;AssociationForeignKe1y:ID"`
}

// Employee - Mitarbeiter
type Employee struct {
	ID          uint   `gorm:"id;not null"`
	PhoneNumber int    `gorm:"column:phone_number;not null"`
	Office      string `gorm:"office;not null"`
    Member      Member `gorm:"ForeignKey:EmployeeID;AssociationForeignKe1y:ID"`
}

// Order - Bestellung
type Order struct {
	ID      uint      `gorm:"column:ID;AUTO_INCREMENT;primary_key"`
	Time    time.Time `gorm:"column:time;default:CURRENT_DATE"`
	UserID  uint  
    Product []Product `gorm:"many2many:orders_products;"`
}

// Image - Bild
type Image struct {
	ID      uint            `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Bindata []byte          `gorm:"column:blob_data"`
	Alttext sql.NullString  `gorm:"column:alttext"`
	Title   sql.NullString  `gorm:"column:title"`
	Caption string          `gorm:"column:caption;not null"`
}

// Category - Kategorie
type Category struct {
	ID              uint   `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Designation     string `gorm:"column:designation;not null"`
	UpperCategoryID uint
    UpperCategory   *Category `gorm:"ForeignKey:UpperCaterogyID"`
	ImageID         uint  
}

// Product - Produkt
type Product struct {
	ID          int    `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Name        string `gorm:"column:name;not null"`
	Description string `gorm:"column:description;not null"`
	ImageID     int    
	CategoryID  int    
	PriceID     uint
	Price       Price        `gorm:"ForeignKey:PriceID"`
    Orders      []Product    `gorm:"many2many:orders_products;"`
	Ingredients []Ingredient `gorm:"many2many:products_ingredients;"`
}

// Price - Preis
type Price struct {
	ID       uint `gorm:"column:ID;AUTO_INCREMENT;primary_key"`
	Guest    uint `gorm:"column:guest"`
	Student  uint `gorm:"column:student"`
	Employee uint `gorm:"column:employee"`
}

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
