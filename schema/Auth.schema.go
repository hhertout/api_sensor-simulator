package schema

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	Key string `gorm:"unique"`
}
