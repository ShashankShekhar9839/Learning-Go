package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ShashankShekhar9839/go-task-manager/internal/cli"
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

	handler := cli.NewTaskHandler(taskService)

	if len(os.Args) < 2 {
		fmt.Println("usage:")
		fmt.Println("add <task title>")
		fmt.Println("list")
		fmt.Println("complete <task id>")
		fmt.Println("delete <task id>")
		return
	}

	command := os.Args[1]

	args := os.Args[2:]

	handler.Handle(command, args)
}