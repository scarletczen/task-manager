package store

import (
	"database/sql"

	"github.com/abhinavthapa1998/task-manager/internal/types"
)

type Storage struct {
	db *sql.DB
}

type Store interface {
	// Users
	CreateUser(u *types.User) (*types.User, error)
	GetUserById(id string) (*types.User, error)
	// Projects
	CreateProject(p *types.Project) error
	GetProject(id string) (*types.Project, error)
	DeleteProject(id string) error
	// Tasks
	CreateTask(t *types.Task) (*types.Task, error)
	GetTask(id string) (*types.Task, error)
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateProject(p *types.Project) error {
	_, err := s.db.Exec("INSERT INTO projects (name) VALUES (?)", p.Name)
	return err
}

func (s *Storage) GetProject(id string) (*types.Project, error) {
	var p types.Project
	err := s.db.QueryRow("SELECT id, name, createdAt FROM projects WHERE id = ?", id).Scan(&p.Id, &p.Name, &p.CreatedAt)
	return &p, err
}

func (s *Storage) DeleteProject(id string) error {
	_, err := s.db.Exec("DELETE FROM projects WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) CreateUser(u *types.User) (*types.User, error) {
	rows, err := s.db.Exec("INSERT INTO users (email, firstName, lastName, password) VALUES (?, ?, ?, ?)", u.Email, u.FirstName, u.LastName, u.Password)
	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	u.Id = id
	return u, nil
}

func (s *Storage) GetUserById(id string) (*types.User, error) {
	var u types.User
	err := s.db.QueryRow("SELECT id, email, firstName, lastName, createdAt FROM users WHERE id = ?", id).Scan(&u.Id, &u.Email, &u.FirstName, &u.LastName, &u.CreatedAt)
	return &u, err
}

func (s *Storage) CreateTask(t *types.Task) (*types.Task, error) {
	rows, err := s.db.Exec("INSERT INTO tasks (name, status, project_id, assigned_to) VALUES (?, ?, ?, ?)", t.Name, t.Status, t.ProjectId, t.AssignedToId)

	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	t.Id = id
	return t, nil
}

func (s *Storage) GetTask(id string) (*types.Task, error) {
	var t types.Task
	err := s.db.QueryRow("SELECT id, name, status, project_id, assigned_to, createdAt FROM tasks WHERE id = ?", id).Scan(&t.Id, &t.Name, &t.Status, &t.ProjectId, &t.AssignedToId, &t.CreatedAt)
	return &t, err
}
