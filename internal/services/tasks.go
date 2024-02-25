package services

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/abhinavthapa1998/task-manager/internal/store"
	"github.com/abhinavthapa1998/task-manager/internal/types"
	"github.com/abhinavthapa1998/task-manager/internal/util"
	"github.com/gorilla/mux"
)

var ErrNameRequired = errors.New("name is required")
var ErrProjectIDRequired = errors.New("project id is required")
var ErrUserIDRequired = errors.New("user id is required")

type TasksService struct {
	store store.Store
}

func NewTasksService(s store.Store) *TasksService {
	return &TasksService{store: s}
}

func (s *TasksService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", WithJWTAuth(s.HandleCreateTask, s.store)).Methods("POST")
	r.HandleFunc("/tasks/{id}", WithJWTAuth(s.handleGetTask, s.store)).Methods("GET")
}

func (s *TasksService) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var task *types.Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		util.WriteJSON(w, http.StatusBadRequest, util.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	if err := validateTaskPayload(task); err != nil {
		util.WriteJSON(w, http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
		return
	}

	t, err := s.store.CreateTask(task)
	if err != nil {
		util.WriteJSON(w, http.StatusInternalServerError, util.ErrorResponse{Error: "Error creating task"})
		return
	}

	util.WriteJSON(w, http.StatusCreated, t)
}

func (s *TasksService) handleGetTask(w http.ResponseWriter, r *http.Request) {

}

func validateTaskPayload(task *types.Task) error {
	if task.Name == "" {
		return ErrNameRequired
	}

	if task.ProjectId == 0 {
		return ErrProjectIDRequired
	}

	if task.AssignedToId == 0 {
		return ErrUserIDRequired
	}

	return nil
}
