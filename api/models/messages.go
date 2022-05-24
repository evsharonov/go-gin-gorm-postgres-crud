package models

import (
	"log"
	"time"

	"gorm.io/gorm"
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

func (m *Message) CreateMessage(db *gorm.DB) error {

	err := db.Create(m).Error

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (m *Message) GetMessages(db *gorm.DB, userId *string) (*[]Message, error) {

	var messages []Message

	err := db.Where("user_to = ?", userId).Find(&messages).Error

	if err != nil {
		log.Println(err)
		return &[]Message{}, err
	}

	return &messages, nil
}
