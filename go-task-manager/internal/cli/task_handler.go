package cli

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ShashankShekhar9839/go-task-manager/internal/service"
)

type TaskHandler struct {
	taskService *service.TaskService
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (h *TaskHandler) Handle(command string, args []string) {
	switch command {

	case "add":
		h.handleAdd(args)

	case "list":
		h.handleList()

	case "complete":
		h.handleComplete(args)

	case "delete":
		h.handleDelete(args)

	default:
		fmt.Println("unknown command")
	}
}

func (h *TaskHandler) handleAdd(args []string) {
	if len(args) < 1 {
		fmt.Println("please provide task title")
		return
	}

	title := args[0]

	err := h.taskService.AddTask(title)
	if err != nil {
		log.Fatalf("failed to add task: %v", err)
	}

	fmt.Println("task added successfully")
}

func (h *TaskHandler) handleList() {
	tasks, err := h.taskService.GetTasks()
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
}

func (h *TaskHandler) handleComplete(args []string) {
	if len(args) < 1 {
		fmt.Println("please provide task id")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("invalid task id")
		return
	}

	err = h.taskService.CompleteTask(id)
	if err != nil {
		log.Fatalf("failed to complete task: %v", err)
	}

	fmt.Println("task completed successfully")
}

func (h *TaskHandler) handleDelete(args []string) {
	if len(args) < 1 {
		fmt.Println("please provide task id")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("invalid task id")
		return
	}

	err = h.taskService.DeleteTask(id)
	if err != nil {
		log.Fatalf("failed to delete task: %v", err)
	}

	fmt.Println("task deleted successfully")
}