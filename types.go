package main

import "time"

type CreateProjectPayload struct {
	Name string `json:"name"`
}

type Project struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterPayload struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

type User struct {
	Id        int64     `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateTaskPayload struct {
	Name         string `json:"name"`
	ProjectId    int64  `json:"projectId"`
	AssignedToId int64  `json:"assignedTo"`
}

type Task struct {
	Id           int64     `json:"id"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	ProjectId    int64     `json:"projectId"`
	AssignedToId int64     `json:"assignedTo"`
	CreatedAt    time.Time `json:"createdAt"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}