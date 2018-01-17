package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// LoadFile -
func LoadFile(filename string) (*Speiseplan, error) {
	root := &Speiseplan{}
	dump, err := ioutil.ReadFile(filename)
	if err != nil {
		return root, err
	}
	return root, xml.Unmarshal(dump, &root)
}

// Menu -
type Menu struct {
	AttrHighlight     string   `xml:"Highlight,attr"`
	AttrKalenderwoche string   `xml:"Kalenderwoche,attr"`
	AttrTag           string   `xml:"Tag,attr"`
	Motto             Motto    `xml:"Motto,omitempty"`
	Produkte          Produkte `xml:"Produkte,omitempty"`
}

// Motto -
type Motto struct {
	Text string `xml:",chardata"`
}

// Produkt -
type Produkt struct {
	AttrProduktID string `xml:"ProduktID,attr"`
	AttrTyp       string `xml:"Typ,attr"`
}

// Produkte -
type Produkte struct {
	Produkt []Produkt `xml:"Produkt,omitempty"`
}

// Speiseplan -
type Speiseplan struct {
	Menu []Menu `xml:"Menu,omitempty"`
}

func main() {
	data, err := LoadFile("Speiseplan.xml")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Speiseplan ist fehlerhaft!")
	} else {
		fmt.Println(data)
	}
}
