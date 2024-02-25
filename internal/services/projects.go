package services

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/abhinavthapa1998/task-manager/internal/store"
	"github.com/abhinavthapa1998/task-manager/internal/types"
	"github.com/abhinavthapa1998/task-manager/internal/util"
	"github.com/gorilla/mux"
)

type ProjectService struct {
	store store.Store
}

func NewProjectService(s store.Store) *ProjectService {
	return &ProjectService{store: s}
}

func (s *ProjectService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/projects", WithJWTAuth(s.handleCreateProject, s.store)).Methods("POST")
	r.HandleFunc("/projects/{id}", WithJWTAuth(s.handleGetProject, s.store)).Methods("GET")
	r.HandleFunc("/projects/{id}", WithJWTAuth(s.handleDeleteProject, s.store)).Methods("DELETE")
}

func (s *ProjectService) handleCreateProject(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var project *types.Project
	err = json.Unmarshal(body, &project)
	if err != nil {
		util.WriteJSON(w, http.StatusBadRequest, util.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	if project.Name == "" {
		util.WriteJSON(w, http.StatusBadRequest, util.ErrorResponse{Error: "Name is required"})
		return
	}

	err = s.store.CreateProject(project)
	if err != nil {
		util.WriteJSON(w, http.StatusInternalServerError, util.ErrorResponse{Error: "Error creating project"})
		return
	}

	util.WriteJSON(w, http.StatusCreated, project)
}

func (s *ProjectService) handleGetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	project, err := s.store.GetProject(id)
	if err != nil {
		util.WriteJSON(w, http.StatusInternalServerError, util.ErrorResponse{Error: "Error getting project"})
		return
	}

	util.WriteJSON(w, http.StatusOK, project)
}

func (s *ProjectService) handleDeleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := s.store.DeleteProject(id)
	if err != nil {
		util.WriteJSON(w, http.StatusInternalServerError, util.ErrorResponse{Error: "Error deleting project"})
		return
	}

	util.WriteJSON(w, http.StatusNoContent, nil)
}
