package model

import (
	"database/sql"
	"time"
)

// Angehoerige -
type Angehoerige struct {
	Nr int64 `schema:"nr"`
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

// Gast -
type Gast struct {
	Grund  string    `schema:"grund"`
	Ablauf time.Time `schema:"ablauf"`
	Nr     int64     `schema:"nr"`
}

// Kategorie -
type Kategorie struct {
	Id            int64          `schema:"id"`
	Bezeichnung   sql.NullString `schema:"bezeichnung"`
	Oberkategorie sql.NullInt64  `schema:"oberkategorie"`
	Kategoriebild sql.NullInt64  `schema:"kategoriebild"`
}

// Mitarbeiter -
type Mitarbeiter struct {
	Telefonnummer sql.NullInt64  `schema:"telefonnummer"`
	Buero         sql.NullString `schema:"buero"`
	Nr            int64          `schema:"nr"`
}

// Nutzer -
type Nutzer struct {
	Nr           int64          `schema:"nr"`
	Aktiv        int64          `schema:"aktiv"`
	Vorname      string         `schema:"vorname"`
	Nachname     string         `schema:"nachname"`
	Mail         string         `schema:"mail"`
	Loginname    string         `schema:"loginname"`
	Anlegedatum  time.Time      `schema:"anlegedatum"`
	LetzterLogin time.Time      `schema:"letzterLogin"`
	Stretch      int64          `schema:"stretch"`
	Algorithmus  sql.NullString `schema:"algorithmus"`
	Salt         string         `schema:"salt"`
	Hash         string         `schema:"hash"`
}

// Preis -
type Preis struct {
	Gastbetrag        int64 `schema:"gastbetrag"`
	Studentenbetrag   int64 `schema:"studentenbetrag"`
	Mitarbeiterbetrag int64 `schema:"mitarbeiterbetrag"`
	Produkt           int64 `schema:"produkt"`
}

// Produkt -
type Produkt struct {
	Id            int64         `schema:"id"`
	Beschreibung  string        `schema:"beschreibung"`
	Vegetarisch   sql.NullInt64 `schema:"vegetarisch"`
	Vegan         sql.NullInt64 `schema:"vegan"`
	ProduktbildId int64         `schema:"produktbildId"`
	KategorieId   int64         `schema:"kategorieId"`
}

// Student -
type Student struct {
	Matrikelnummer int64  `schema:"matrikelnummer"`
	Studiengang    string `schema:"studiengang"`
	Nr             int64  `schema:"nr"`
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
