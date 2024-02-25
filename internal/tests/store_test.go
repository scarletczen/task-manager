package tests

import "github.com/abhinavthapa1998/task-manager/internal/types"

// Mocks

type MockStore struct{}

func (s *MockStore) CreateProject(p *types.Project) error {
	return nil
}

func (s *MockStore) GetProject(id string) (*types.Project, error) {
	return &types.Project{Name: "Super cool project"}, nil
}

func (s *MockStore) DeleteProject(id string) error {
	return nil
}

func (s *MockStore) CreateUser(u *types.User) (*types.User, error) {
	return &types.User{}, nil
}

func (s *MockStore) GetUserById(id string) (*types.User, error) {
	return &types.User{}, nil
}

func (s *MockStore) CreateTask(t *types.Task) (*types.Task, error) {
	return &types.Task{}, nil
}

func (s *MockStore) GetTask(id string) (*types.Task, error) {
	return &types.Task{}, nil
}
