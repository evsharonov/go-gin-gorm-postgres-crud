package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id"  gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"index"`
	Name      string    `json:"name"`
}