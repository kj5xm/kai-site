package services

type TaskService struct {
}

type CreateTaskRequest struct {
	Size int `json:"size"`
}

func (s *TaskService) CreateService(request CreateTaskRequest)
