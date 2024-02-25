package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var errNameRequired = errors.New("name is required")
var errProjectIdRequired = errors.New("project id is required")
var errUserIdRequired = errors.New("user id is required")

type TasksService struct {
	store Store
}

func NewTasksService(s Store) *TasksService {
	return &TasksService{store: s}
}

func (s *TasksService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", s.handleCreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", s.handleGetTask).Methods("GET")
}

func (s *TasksService) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()

	var task *Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		return
	}
	if err := validateTaskPayload(task); err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	t, err := s.store.CreateTask(task)
	
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating task"})
		return
	}
	WriteJSON(w, http.StatusCreated, t)

}

func (s *TasksService) handleGetTask(w http.ResponseWriter, r *http.Request) {

}

func validateTaskPayload(task *Task) error{
	if task.Name == "" {
		return errNameRequired
	}
	if task.ProjectId == 0 {
		return errProjectIdRequired
	}
	if task.AssignedToId == 0 {
		return errUserIdRequired
	}
	return nil
}
