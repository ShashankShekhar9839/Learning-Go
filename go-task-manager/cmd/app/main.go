package main

import (
	"fmt"
	"log"

	"github.com/ShashankShekhar9839/go-task-manager/internal/config"
)


func main() {
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Application Name:", cfg.App.Name)
	fmt.Println("Version:", cfg.App.Version)
	fmt.Println("Storage File:", cfg.Storage.FilePath)
	fmt.Println("Nothing File:", cfg.Nothing.Nothing)
}