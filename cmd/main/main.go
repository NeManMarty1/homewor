package main

import (
	"homework/internal/handlers"
	"homework/internal/repository"
	"homework/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	bookRepo := repository.NewRepository()
	bookService := services.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)
	taskRepo := repository.NewTaskRepository()
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	r := gin.Default()

	r.GET("/books", bookHandler.GetAllBooksHandler)
	r.POST("/books", bookHandler.AddBookHandler)
	r.GET("/tasks", taskHandler.GetAllTasksHandler)
	r.POST("/tasks", taskHandler.AddTasksHandler)
	r.Run(":8080")

}
