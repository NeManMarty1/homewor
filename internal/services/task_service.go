package services

import (
	"errors"
	"homework/internal/models"
	"homework/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetTasks() []models.Task {
	return s.repo.GetAllTasks()
}

func (s *TaskService) AddTask(title string) (models.Task, error) {
	if title == "" {
		return models.Task{}, errors.New("Pizdec drug , ti ne smog sozdat` tasku")
	}
	newTask := models.Task{
		Title: title,
	}
	return s.repo.AddTask(newTask), nil
}
