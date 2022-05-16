package models

import (
	"time"
)

type Message struct {
	ID        int       `json:"id"  gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"index"`
	UserFrom  int       `json:"userfrom"`
	UserTo    int       `json:"userto"`
	Message   string    `json:"message"`
}
