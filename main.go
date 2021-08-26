package main

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kj5xm/task-manager/handlers"
)

const (
	Timeout  = 60
	Throttle = 10
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(logger.NewMiddlewareLogger())
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.Timeout(Timeout * time.Second))
	r.Use(middleware.Throttle(Throttle))

	taskHandler := handlers.TaskHandler{}

	r.Get("/task/showtasks")

}
