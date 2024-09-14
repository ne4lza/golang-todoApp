package routes

import (
	"ToDoApp/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, a *service.ApiHandler) {
	r.GET("/tasks/:id", a.GetByIdHandler)
	r.GET("/tasks", a.GetAllToDosHandler)
	r.POST("/tasks/create", a.CreateToDoHandler)
}
