package model

import (
	"database/sql"
	"time"
)

// User - FH-Nutzer
type User struct {
	Nr           uint           `gorm:"column:nr;AUTO_INCREMENT;primary_key"`
	Active       bool           `gorm:"column:active;default:true"`
	Firstname    string         `gorm:"column:firstname;not null"`
	Lastname     string         `gorm:"column:lastname;not null"`
	Mail         string         `gorm:"column:mail;unique"`
	Loginname    string         `gorm:"column:loginname"`
	CreatedAt    time.Time      `gorm:"column:created_at;default:CURRENT_DATE"`
	LastLogin    time.Time      `gorm:"column:last_login"`
	Stretch      uint           `gorm:"column:stretch;not null"`
	Algo         string         `gorm:"column:algo;not null"`
	Salt         string         `gorm:"column:salt;type:varchar(32);not null"`
	Hash         string         `gorm:"column:hash;type:varchar(64);not null"`
}

// Member - Fh-Angehoerige
type Member struct {
	Nr     uint `gorm:"column:nr;primary_key"`
    User   User `gorm:"ForeignKey:UserNr;AssociationForeignKey:Nr"`
    UserNr uint
}

// Guest - Gast
type Guest struct {
    Reason     string    `gorm:"column:reason;not null"`
	ExpiryDate time.Time `gorm:"column:expiry_date;default:CURRENT_DATE"`
	Nr         uint      `gorm:"column:nr;not null"`
}

// Student - Student
type Student struct {
	StudentId      int    `gorm:"column:student_id;not null"`
	Course         string `gorm:"column:course;not null"`
    Member         Member `gorm:"column:nr;ForeignKey:UserNr;AssociationForeignKey:Refer"`
    MemberNr       uint

	Nr             uint  `schema:"nr"`
}
//////////////////////////////////////////////
// Mitarbeiter -
type Mitarbeiter struct {
	Telefonnummer sql.NullInt64  `schema:"telefonnummer"`
	Buero         sql.NullString `schema:"buero"`
	Nr            int64          `schema:"nr"`
}

// Bestellung -
type Bestellung struct {
	Id        int64     `schema:"id"`
	Zeitpunkt time.Time `schema:"zeitpunkt"`
	Nutzernr  int64     `schema:"nutzernr"`
}

// Bild -
type Bild struct {
	Id               int64          `schema:"id"`
	Binaerdaten      []byte         `schema:"binaerdaten"`
	Alttext          sql.NullString `schema:"alttext"`
	Titel            sql.NullString `schema:"titel"`
	Bildunterschrift sql.NullString `schema:"bildunterschrift"`
}

// Kategorie -
type Kategorie struct {
	Id            int64          `schema:"id"`
	Bezeichnung   sql.NullString `schema:"bezeichnung"`
	Oberkategorie sql.NullInt64  `schema:"oberkategorie"`
	Kategoriebild sql.NullInt64  `schema:"kategoriebild"`
}

// Produkt -
type Product struct {
	Id            int           `schema:"id"`
	Beschreibung  string        `schema:"beschreibung"`
	Vegetarisch   sql.NullInt64 `schema:"vegetarisch"`
	Vegan         sql.NullInt64 `schema:"vegan"`
	ProduktbildId int           `schema:"produktbildId"`
	KategorieId   int           `schema:"kategorieId"`
}


// Preis -
type Preis struct {
	Gastbetrag        int64 `schema:"gastbetrag"`
	Studentenbetrag   int64 `schema:"studentenbetrag"`
	Mitarbeiterbetrag int64 `schema:"mitarbeiterbetrag"`
	Produkt           int64 `schema:"produkt"`
}

// Zutat
type Zutat struct {
	Id           int64          `schema:"id"`
	Glutenfrei   sql.NullInt64  `schema:"glutenfrei"`
	Bio          sql.NullInt64  `schema:"bio"`
	Vegetarisch  sql.NullInt64  `schema:"vegetarisch"`
	Vegan        sql.NullInt64  `schema:"vegan"`
	Beschreibung sql.NullString `schema:"beschreibung"`
	Name         string         `schema:"Name"`
}

// // TransProduktBestellung -
// type TransProduktBestellung struct {
//	Produktid int64 `schema:"produktid"`
//	Bestellid int64 `schema:"bestellid"`
// }

// // TransProduktZutat -
// type TransProduktZutat struct {
//	ProduktId int64 `schema:"produktId"`
//	ZutatId   int64 `schema:"zutatId"`
// }
