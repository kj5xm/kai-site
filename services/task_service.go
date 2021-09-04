package services

import (
	"time"

	"github.com/kj5xm/task-manager/lib/logger"
)

type TaskService struct {
}

type CreateTaskRequest struct {
	Size int `json:"size"`
}

func (s *TaskService) CreateTask(request CreateTaskRequest) {
	logger.Info("MakePizza", "Creating pizza")

	return &CreateTaskResponse{
		ID: 
	}
}
