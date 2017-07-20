package Models

import (
	"github.com/jinzhu/gorm"
)

type Email struct {
	gorm.Model
	UserID int    `gorm:"index"`                          // Foreign key (belongs to), tag `index` will create index for this column
	Email  string `gorm:"type:varchar(100);unique_index"` // `type` set sql type, `unique_index` will create unique index for this column
	Active bool
}
