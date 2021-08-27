package handlers

import (
	"net/http"

	"github.com/kj5xm/task-manager/services"
)

type TaskHandler struct {
	TaskService services.TaskService
}

func (s *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
}

func (s *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
}

func (s *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
}

func (s *TaskHandler) UpdateTaskByID(w http.ResponseWriter, r *http.Request) {
}

func (s *TaskHandler) DeleteTaskByID(w http.ResponseWriter, r *http.Request) {
}
