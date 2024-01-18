package schemas

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Description string
	IsDone      bool
}
