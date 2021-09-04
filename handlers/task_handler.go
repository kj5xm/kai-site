package handlers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/kj5xm/task-manager/lib/logger"
	rr "github.com/kj5xm/task-manager/lib/restutils"
	"github.com/kj5xm/task-manager/services"
)

type TaskHandler struct {
	TaskService services.TaskService
}

func (s *TaskHandler) Init(db *sql.DB) {
	
}

func (s *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("Create Task", err, "error reading body")
		rr.JsonResp(w, 500, rr.Message{Message: "error reading request body", Error: err.Error()})
		return
	}

	var req services.CreateTaskRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		rr.JsonResp(w, 500, rr.Message{Message: "unmarshal request fail", Error: err.Error()})
	}

	resp, err := s.TaskService.CreateTask(req)
	if err != nil {
		rr.JsonResp(w, 400, rr.Message{Message: "create pizza failed", Error: err.Error()})
		return
	}

	rr.JsonResp(w, 200, resp)
	return
}

func (s *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
}

func (s *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
}

func (s *TaskHandler) UpdateTaskByID(w http.ResponseWriter, r *http.Request) {
}

func (s *TaskHandler) DeleteTaskByID(w http.ResponseWriter, r *http.Request) {
}
