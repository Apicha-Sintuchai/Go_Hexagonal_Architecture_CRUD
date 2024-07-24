package main

import (
	"go/clean/adapters"
	"go/clean/core"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	dsn := "host=localhost user=postgres password=0126 dbname=CleanArchitecture port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&core.Todolist{})

	TodoRepo := adapters.NewGormOrderRepository(db)
	TodoService := core.NewTodolistService(TodoRepo)
	TodoHandler := adapters.NewHttpTodoHandler(TodoService)
	r.GET("/order", TodoHandler.GetTodo)
	r.GET("/order/:id", TodoHandler.GetID)
	r.POST("/order", TodoHandler.CreateTodo)
	r.PUT("/order/:id", TodoHandler.UpdateTodo)
	r.DELETE("/order/:id", TodoHandler.DeleteID)
	r.Run("localhost:8080")
}
