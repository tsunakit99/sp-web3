package main

import (
	"log"

	"github.com/tsunakit99/sp-web3/handler"
	"github.com/tsunakit99/sp-web3/infra"
	"github.com/tsunakit99/sp-web3/infra/postgres"
	"github.com/tsunakit99/sp-web3/middleware"
	"github.com/tsunakit99/sp-web3/usecase/task"

	"github.com/labstack/echo/v4"
)

func main() {
	db := infra.ConnectDB()
	taskRepo := postgres.NewTaskRepository(db)
	taskUsecase := task.New(taskRepo)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	e := echo.New()
	e.Use(middleware.SupabaseAuthMiddleware)

	e.GET("/tasks", taskHandler.GetTasks)

	log.Fatal(e.Start(":4000"))
}