package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abhinavthapa1998/task-manager/internal/services"
	"github.com/abhinavthapa1998/task-manager/internal/types"
	"github.com/abhinavthapa1998/task-manager/internal/util"
	"github.com/gorilla/mux"
)

func TestCreateTask(t *testing.T) {
	ms := &MockStore{}
	service := services.NewTasksService(ms)

	t.Run("should return error if name is empty", func(t *testing.T) {
		payload := &types.CreateTaskPayload{
			Name: "",
		}

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/tasks", service.HandleCreateTask)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}

		var response util.ErrorResponse
		err = json.NewDecoder(rr.Body).Decode(&response)
		if err != nil {
			t.Fatal(err)
		}

		if response.Error != services.ErrNameRequired.Error() {
			t.Errorf("expected error message %s, got %s", response.Error, services.ErrNameRequired.Error())
		}
	})

	t.Run("should create a task", func(t *testing.T) {
		payload := &types.CreateTaskPayload{
			Name:         "Creating a REST API in go",
			ProjectId:    1,
			AssignedToId: 42,
		}

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/tasks", service.HandleCreateTask)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}
