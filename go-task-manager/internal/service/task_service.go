package service

import (
	"fmt"
	"time"

	"github.com/ShashankShekhar9839/go-task-manager/internal/storage"
	"github.com/ShashankShekhar9839/go-task-manager/internal/task"
)

type TaskService struct {
	storage storage.Storage
}

func NewTaskService(storage storage.Storage) *TaskService {
	return &TaskService{
		storage: storage,
	}
}

func (s *TaskService) AddTask(title string) error {
	tasks, err := s.storage.LoadTasks()
	if err != nil {
		return  err
	}

	newTask := task.Task{
		ID: len(tasks) + 1, 
		Title: title,
		Completed: false,
		CreatedAt: time.Now(), 

	}

	tasks = append(tasks, newTask)

	err = s.storage.SaveTasks(tasks)

	if err != nil {
		return  err
	}
   
	return nil

}

func (s *TaskService) GetTasks() ([]task.Task, error) {
       return  s.storage.LoadTasks()
}

func (s *TaskService) CompleteTask(id int) error  {
    
	tasks, err := s.storage.LoadTasks();
	if err != nil {
         return  err
	}

	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			tasks[i].CompletedAt = time.Now()

			found = true;
			break
		}
	}
     
	if !found {
		return  fmt.Errorf("task with is %d not found", id)
	}
     
	err = s.storage.SaveTasks(tasks)

	if err != nil {
		 return  err
	}

	return  nil 

}
