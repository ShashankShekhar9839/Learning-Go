package main

import (
	"fmt"
	"log"

	"github.com/ShashankShekhar9839/go-task-manager/internal/config"
	"github.com/ShashankShekhar9839/go-task-manager/internal/storage"
)

func main() {
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	store := storage.NewJSONStorage(cfg.Storage.FilePath)

	tasks, err := store.LoadTasks()
	if err != nil {
		log.Fatalf("failed to load tasks: %v", err)
	}

	fmt.Println("Loaded Tasks:", tasks)
}