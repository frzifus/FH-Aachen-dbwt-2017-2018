package model

import (
	"time"
)

//Test is just a test
type Test struct {
	ID        int       `schema:"-"`
	Test      string    `schema:"test"`
	CreatedAt time.Time `schema:"-"`
	UpdatedAt time.Time `schema:"-"`
}
