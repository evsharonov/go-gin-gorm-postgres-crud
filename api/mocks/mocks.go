package mocks

import (
	"time"

	"github.com/evsharonov/go-gin-gorm-crud/api/models"
)

var Users = []models.User{
	{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Date(1, 1, 1, 0, 0, 0, 0, time.Local), Name: "John Paul"},
	{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Date(1, 1, 1, 0, 0, 0, 0, time.Local), Name: "Peter Pen"},
	{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Date(1, 1, 1, 0, 0, 0, 0, time.Local), Name: "Captain Hook"},
}

var Messages = []models.Message{
	{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Date(1, 1, 1, 0, 0, 0, 0, time.Local), UserFrom: 1, UserTo: 2, Message: "Hello Pavlo"},
	{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Date(1, 1, 1, 0, 0, 0, 0, time.Local), UserFrom: 2, UserTo: 3, Message: "Type me later"},
	{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Date(1, 1, 1, 0, 0, 0, 0, time.Local), UserFrom: 3, UserTo: 2, Message: "I`m here"},
}
