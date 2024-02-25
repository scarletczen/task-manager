package routes

import (
	"log"
	"net/http"

	"github.com/abhinavthapa1998/task-manager/internal/services"
	"github.com/abhinavthapa1998/task-manager/internal/store"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store store.Store
}

func NewAPIServer(addr string, store store.Store) *APIServer {
	return &APIServer{
		addr:  addr,
		store: store,
	}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	projectService := services.NewProjectService(s.store)
	projectService.RegisterRoutes(subrouter)

	userService := services.NewUserService(s.store)
	userService.RegisterRoutes(subrouter)

	tasksService := services.NewTasksService(s.store)
	tasksService.RegisterRoutes(subrouter)

	log.Println("Starting the API server at", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
