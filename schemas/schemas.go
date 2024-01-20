package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Description string
	IsDone      bool
}

type TaskResponse struct {
	Id          uint      `json:id`
	CreatedAt   time.Time `json:createAt`
	UpdatedAt   time.Time `json:updateAt`
	DeletedAt   time.Time `json:deletedAt,omitermpty`
	Description string    `json:description`
	IsDone      bool      `json:is-done`
}
