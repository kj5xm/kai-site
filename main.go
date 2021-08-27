package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kj5xm/task-manager/handlers"
	"github.com/kj5xm/task-manager/services"
)

const (
	Timeout  = 60
	Throttle = 10
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	//r.Use(logger.NewMiddlewareLogger())
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.Timeout(Timeout * time.Second))
	r.Use(middleware.Throttle(Throttle))

	taskHandler := handlers.TaskHandler{
		TaskService: services.TaskService{},
	}

	r.Post("/tasks", taskHandler.CreateTask)
	r.Get("/tasks", taskHandler.GetTasks)
	r.Get("/tasks/:id", taskHandler.GetTaskByID)
	r.Put("/tasks/:id", taskHandler.UpdateTaskByID)
	r.Delete("/tasks/:id", taskHandler.DeleteTaskByID)

	http.ListenAndServe(":8080", r)
	//logger.Info("Init", "Server started")

}
