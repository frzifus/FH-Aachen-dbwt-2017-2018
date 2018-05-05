package model

import (
	"database/sql"
)

// Image - Bild
type Image struct {
	ID      uint           `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Bindata []byte         `gorm:"column:blob_data"`
	Alttext sql.NullString `gorm:"column:alttext"`
	Title   sql.NullString `gorm:"column:title"`
	Caption string         `gorm:"column:caption;not null"`
}
