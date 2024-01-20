package handler

import (
	"silvanotfound/tasks/config"
	"silvanotfound/tasks/schemas"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("Handler")
	db = config.GetSqlite()
}

func CreateTaskHandler(ctx *gin.Context) {
	request := CreateTaskRequest{}
	ctx.BindJSON(&request)
	if err := request.validade(); err != nil {
		sendResponseError(ctx, err.Error())
		return
	}
	task := schemas.Task{
		Description: request.Description,
		IsDone:      *request.IsDone,
	}
	db.Create(&task)
	sendResponseSuccess(ctx, "Atividade criada com sucesso", task)
}

type UpdateTask struct {
	Description string
	IsDone      *bool
}

func UpdateTaskHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendResponseError(ctx, "Aprameter Id is querired")
	}
	var taskModel schemas.Task
	if err := db.First(&taskModel, id).Error; err != nil {
		sendResponseError(ctx, "task id not found")
		return
	}
	request := CreateTaskRequest{}
	ctx.BindJSON(&request)
	if err := request.validade(); err != nil {
		sendResponseError(ctx, err.Error())
		return
	}
	task := UpdateTask{
		Description: request.Description,
		IsDone:      request.IsDone,
	}
	db.Model(taskModel).Updates(task)
	sendResponseSuccess(ctx, "Atividade atualizada com sucesso", task)
}

func DeleteTaskHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendResponseError(ctx, "Aprameter Id is querired")
	}
	task := schemas.Task{}
	if err := db.First(&task, id).Error; err != nil {
		sendResponseError(ctx, "task id not found")
		return
	}
	db.Delete(&task)
	sendResponseSuccess(ctx, "Atividade deletada com sucesso", task)
}

func ShowTaskHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendResponseError(ctx, "Aprameter Id is querired")
	}
	task := schemas.Task{}
	if err := db.First(&task, id).Error; err != nil {
		sendResponseError(ctx, "task id not found")
		return
	}
	sendResponseSuccess(ctx, "", task)
}

func ShowAllTasksHandler(ctx *gin.Context) {
	tasks := []schemas.Task{}
	db.Find(&tasks)
	sendResponseSuccess(ctx, "", tasks)
}
