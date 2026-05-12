package storage

import (
	"encoding/json"
	"os"

	"github.com/ShashankShekhar9839/go-task-manager/internal/task"
)

type JSONStorage struct {
	FilePath string
}

func NewJSONStorage(filePath string) *JSONStorage {
	return  &JSONStorage{
		FilePath: filePath,
	}
}

func (j *JSONStorage) LoadTasks() ([]task.Task, error) {
	file, err := os.ReadFile(j.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []task.Task{}, nil
		}

		return nil, err
	}

	var tasks []task.Task
    
	if len(file) == 0 {
		return []task.Task{}, nil
	}

	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (j *JSONStorage) SaveTasks(tasks []task.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(j.FilePath, data, 0644)
}

