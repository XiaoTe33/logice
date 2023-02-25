package model

import (
	"time"
)

type Student struct {
	ID       int       `grom:"id"`
	Name     string    `gorm:"name"`
	Age      int       `gorm:"age"`
	Birthday time.Time `gorm:"birthday"`
}
