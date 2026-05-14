package main

import (
	"fmt"
	"log"
	"os"      // Bridge between Go program and operating system
	"strconv" // package to convert int to string becuase args for os are string

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

	if len(os.Args) < 2 {
		fmt.Println("usage:")
		fmt.Println("add <task title>")
		fmt.Println("list")
		return
	}

	command := os.Args[1]
	fmt.Println(os.Args)

	switch command {

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("please provide task title")
			return
		}

		title := os.Args[2]

		err := taskService.AddTask(title)
		if err != nil {
			log.Fatalf("failed to add task: %v", err)
		}

		fmt.Println("task added successfully")

	case "list":
		tasks, err := taskService.GetTasks()
		if err != nil {
			log.Fatalf("failed to get tasks: %v", err)
		}

		for _, task := range tasks {
			fmt.Printf(
				"ID: %d | Title: %s | Completed: %v\n",
				task.ID,
				task.Title,
				task.Completed,
			)
		}

	case "complete": 
	if len(os.Args) < 3 {
		fmt.Println("Please provide the task id")
		return
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("invalid task id")
		return
	}

	err = taskService.CompleteTask(id)

	if err != nil {
		log.Fatalf("failed to complete task: %v", err)
	}

	fmt.Println("task completed successfully")

	default:
		fmt.Println("unknown command")
	}
}