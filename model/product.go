package model

import (
	"time"
)

// just a Product
type Product struct {
	ID        int       `schema:"-"`
	Test      string    `schema:"test"`
	CreatedAt time.Time `schema:"-"`
	UpdatedAt time.Time `schema:"-"`
}
