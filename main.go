package main

import (
	"ToDoApp/db"
	"ToDoApp/routes"
	"ToDoApp/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.GetDb()
	defer db.DB.Close()
	r := gin.Default()
	apiHandlers := service.GetHandlers(db)

	routes.RegisterRoutes(r, apiHandlers)

	_ = r.Run("localhost:8080")
}
