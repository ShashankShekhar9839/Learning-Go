package storage

import "github.com/ShashankShekhar9839/go-task-manager/internal/task"

type Storage interface {
	LoadTasks() ([]task.Task, error)
	SaveTasks([]task.Task) error
}