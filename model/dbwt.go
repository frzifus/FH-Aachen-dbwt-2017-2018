package model

import (
	"database/sql"
	"time"
)

// User - FH-Nutzer
type User struct {
	ID        uint      `gorm:"column:ID;AUTO_INCREMENT;primary_key"`
	Active    bool      `gorm:"column:active;default:true"`
	Firstname string    `gorm:"column:firstname;not null"`
	Lastname  string    `gorm:"column:lastname;not null"`
	Mail      string    `gorm:"column:mail;unique"`
	Loginname string    `gorm:"column:loginname"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_DATE"`
	LastLogin time.Time `gorm:"column:last_login"`
	Stretch   uint      `gorm:"column:stretch;not null"`
	Algo      string    `gorm:"column:algo;not null"`
	Salt      string    `gorm:"column:salt;type:varchar(32);not null"`
	Hash      string    `gorm:"column:hash;type:varchar(64);not null"`
}

// Member - Fh-Angehoerige
type Member struct {
	ID     uint `gorm:"column:nr;primary_key"`
	User   User `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	UserID uint
}

// Guest - Gast
type Guest struct {
	Reason     string    `gorm:"column:reason;not null"`
	ExpiryDate time.Time `gorm:"column:expiry_date;default:CURRENT_DATE"`
	ID         uint      `gorm:"column:nr;not null"`
}

// Student - Student
type Student struct {
	StudentID uint   `gorm:"column:student_id;not null"`
	Course    string `gorm:"column:course;not null"`
	Member    Member `gorm:"column:nr;ForeignKey:UserNr;AssociationForeignKey:Refer"`
	MemberNr  uint

	Nr uint `schema:"nr"`
}

// Employee - Mitarbeiter
type Employee struct {
	ID          uint           `schema:"nr"`
	PhoneNummer sql.NullInt64  `schema:"telefonnummer"`
	Office      sql.NullString `schema:"buero"`
}

// Order - Bestellung
type Order struct {
	ID     int64     `schema:"id"`
	Time   time.Time `schema:"zeitpunkt"`
	UserNr int64     `schema:"nutzernr"`
}

// Image - Bild
type Image struct {
	ID      int64          `schema:"id"`
	Bindata []byte         `schema:"binaerdaten"`
	Alttext sql.NullString `schema:"alttext"`
	Title   sql.NullString `schema:"titel"`
	Caption sql.NullString `schema:"bildunterschrift"`
}

// Category - Kategorie
type Category struct {
	ID            int64          `schema:"id"`
	Designation   sql.NullString `schema:"bezeichnung"`
	UpperCategory sql.NullInt64  `schema:"oberkategorie"`
	ImageID       sql.NullInt64  `schema:"kategoriebild"`
}

// Produkt -
type Product struct {
	ID          int `schema:"id"`
	Name        string
	Description string `schema:"beschreibung"`
	ImageID     int    `schema:"produktbildId"`
	CategoryID  int    `schema:"kategorieId"`
	PriceID     uint
	Price       Price        `gorm:"ForeignKey:PriceID"`
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
	Glutenfree  bool      `gorm:"column:gluten_free;not null"`
	Bio         bool      `gorm:"column:bio;not null"`
	Vegetarian  bool      `gorm:"column:vegetarian;not null"`
	Vegan       bool      `gorm:"column:vegan;not null"`
	Description string    `gorm:"column:description;not null"`
	Name        string    `gorm:"column:name;type:varchar(64)"`
	Products    []Product `gorm:"many2many:products_ingredients;"`
}
