package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/vyantik/todo-backend"
	"github.com/vyantik/todo-backend/pkg/handler"
	"github.com/vyantik/todo-backend/pkg/repository"
	"github.com/vyantik/todo-backend/pkg/service"
)

func main() {
	clearScreen()

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
