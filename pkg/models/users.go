package models

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int       `json:"id"  gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"index"`
	Name      string    `json:"name"`
}

func (u *User) GetAllUsers(db *gorm.DB) (*[]User, error) {

	var users []User

	err := db.Find(&users).Error

	if err != nil {
		log.Println(err)
		return &[]User{}, err
	}

	return &users, nil
}

func (u *User) GetUserById(db *gorm.DB, id *string) (*User, error) {

	err := db.First(u, *id).Error

	if err != nil {
		log.Println(err)
		return &User{}, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &User{}, errors.New("User not found")
	}

	return u, nil
}

func (u *User) CreateUser(db *gorm.DB) error {

	err := db.Create(u).Error

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
