package repository

import "homework/internal/models"

type TaskRepository struct {
	tasks []models.Task
	nextID int
}

func NewTaskRepository() *TaskRepository{
	return &TaskRepository{
		tasks: []models.Task{},
		nextID: 1,
	}
}

func (r *TaskRepository) GetAllTasks() []models.Task{
	return r.tasks
}

func (r *TaskRepository) AddTask(task models.Task) models.Task{
	task.ID = r.nextID
	r.nextID++
	r.tasks = append(r.tasks, task)
	return task
}