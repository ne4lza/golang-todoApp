package service

import (
	"ToDoApp/db"
	"ToDoApp/model"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func createToDo(db *db.PostgresDB, c *gin.Context) (int, gin.H, error) {
	var todo model.ToDoModel

	if err := c.ShouldBindJSON(&todo); err != nil {
		return http.StatusBadRequest, nil, err
	}

	if len(todo.Title) == 0 || len(todo.Description) == 0 {
		return http.StatusBadRequest, nil, errors.New("Lütfen tüm alanları sağlayın")
	}

	query := "INSERT INTO TBL_Tasks (title, description) VALUES ($1, $2) RETURNING id"
	err := db.DB.QueryRow(query, todo.Title, todo.Description).Scan(&todo.ID)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	response := gin.H{
		"message": "Kayıt İşlemi Başarılı",
		"item": todo,
	}
	return http.StatusOK, response, nil
}

func getToDoById(db *db.PostgresDB, c *gin.Context) (int, *model.ToDoModel, error) {
	todoId := c.Param("id")
	id, err := strconv.Atoi(todoId)

	if err != nil {
		return http.StatusNotFound, nil, err
	}

	var todo model.ToDoModel

	query := "SELECT * FROM TBL_Tasks where id = $1"

	err = db.DB.QueryRow(query, id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.IsCompleted)

	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	}

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, &todo, nil
}

func GetAllToDos(db *db.PostgresDB, c *gin.Context) (int, *model.ToDoModel, error) {
	var todo model.ToDoModel

	query := "SELECT * FROM TBL_Tasks"

	err := db.DB.QueryRow(query).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.IsCompleted)

	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	}

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, &todo, nil
}
