package handlers

import (
	"homework/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service *services.TaskService
}

func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) GetAllTasksHandler(c *gin.Context) {
	tasks := h.service.GetTasks()
	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) AddTasksHandler(c *gin.Context) {
	var input struct {
		Title string `json="title"`
	}
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	task, err := h.service.AddTask(input.Title)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, task	)
}
