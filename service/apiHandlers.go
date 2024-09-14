package service

import (
	"ToDoApp/db"

	"github.com/gin-gonic/gin"
)

type ApiHandler struct {
	db *db.PostgresDB
}

func GetHandlers(db *db.PostgresDB) *ApiHandler {
	return &ApiHandler{db: db}
}

func (a *ApiHandler) CreateToDoHandler(c *gin.Context) {
	status, response, err := createToDo(a.db, c)

	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(status, response)
}

func (a *ApiHandler) GetByIdHandler(c *gin.Context) {
	status, todo, err := getToDoById(a.db, c)

	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(status, todo)
}

func (a *ApiHandler) GetAllToDosHandler(c *gin.Context) {
	status, todo, err := GetAllToDos(a.db, c)

	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(status, todo)
}
