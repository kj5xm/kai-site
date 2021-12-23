package services

import (
	"fmt"
	"time"

	"github.com/kj5xm/task-manager/lib/logger"
)

type TaskService struct {
}

type CreateTaskRequest struct {
	Size int `json:"size"`
}

type CreateTaskResponse struct {
	Size       int
	Price      string
	CreateDate time.Time
}

func (s *TaskService) CreateTask(request CreateTaskRequest) (*CreateTaskResponse, error) {
	logger.Info("MakePizza", "Creating pizza")

	if request.Size > 10 {
		return nil, fmt.Errorf("pizza too big")
	}

	return &CreateTaskResponse{
		Size:       request.Size,
		Price:      "$10",
		CreateDate: time.Now(),
	}, nil
}
