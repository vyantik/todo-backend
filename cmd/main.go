package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/vyantik/todo-backend"
	"github.com/vyantik/todo-backend/pkg/handler"
	"github.com/vyantik/todo-backend/pkg/repository"
	"github.com/vyantik/todo-backend/pkg/service"
)

func main() {
	clearScreen()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
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
