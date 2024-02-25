package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store Store
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	// register services
	usersService := NewUserService(s.store)
	usersService.RegisterRoutes(subrouter)
	tasksService := NewTasksService(s.store)
	tasksService.RegisterRoutes(subrouter)
	log.Println("Starting the API server at", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
