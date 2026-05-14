package main

import (
	"fmt"
	"log"

	"github.com/ShashankShekhar9839/go-task-manager/internal/config"
	"github.com/ShashankShekhar9839/go-task-manager/internal/service"
	"github.com/ShashankShekhar9839/go-task-manager/internal/storage"
)

func main() {
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	store := storage.NewJSONStorage(cfg.Storage.FilePath)

	taskService := service.NewTaskService(store)

	err = taskService.Addtask("Learn Go");
	err = taskService.Addtask("Learn Go Again");
	if err != nil {
		log.Fatalf("Failed to add task: %v", err)
	}

	tasks, err := taskService.GetTasks()

	if err != nil {
		log.Fatalf("Failed to get tasks: %v", err)
	}

	fmt.Println(tasks)
}