package services

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/abhinavthapa1998/task-manager/internal/config"
	"github.com/abhinavthapa1998/task-manager/internal/store"
	"github.com/abhinavthapa1998/task-manager/internal/types"
	"github.com/abhinavthapa1998/task-manager/internal/util"
	"github.com/gorilla/mux"
)

var ErrEmailRequired = errors.New("email is required")
var ErrFirstNameRequired = errors.New("first name is required")
var ErrLastNameRequired = errors.New("last name is required")
var ErrPasswordRequired = errors.New("password is required")

type UserService struct {
	store store.Store
}

func NewUserService(s store.Store) *UserService {
	return &UserService{store: s}
}

func (s *UserService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users/register", s.HandleUserRegister).Methods("POST")
	r.HandleFunc("/users/login", s.handleUserLogin).Methods("POST")
}

func (s *UserService) HandleUserRegister(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var payload *types.User
	err = json.Unmarshal(body, &payload)
	if err != nil {
		util.WriteJSON(w, http.StatusBadRequest, util.ErrorResponse{Error: "Invalid request payload"})
		return
	}

	if err := ValidateUserPayload(payload); err != nil {
		util.WriteJSON(w, http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
		return
	}

	hashedPassword, err := HashPassword(payload.Password)
	if err != nil {
		util.WriteJSON(w, http.StatusInternalServerError, util.ErrorResponse{Error: "Error creating user"})
		return
	}
	payload.Password = hashedPassword

	u, err := s.store.CreateUser(payload)
	if err != nil {
		util.WriteJSON(w, http.StatusInternalServerError, util.ErrorResponse{Error: "Error creating user"})
		return
	}

	token, err := createAndSetAuthCookie(u.Id, w)
	if err != nil {
		util.WriteJSON(w, http.StatusInternalServerError, util.ErrorResponse{Error: "Error creating user"})
		return
	}

	util.WriteJSON(w, http.StatusCreated, token)
}

func (s *UserService) handleUserLogin(w http.ResponseWriter, r *http.Request) {
	// 1. Find user in db by email
	// 2. Compare password with hashed password
	// 3. Create JWT and set it in a cookie
	// 4. Return JWT in response
}

func ValidateUserPayload(user *types.User) error {
	if user.Email == "" {
		return ErrEmailRequired
	}

	if user.FirstName == "" {
		return ErrFirstNameRequired
	}

	if user.LastName == "" {
		return ErrLastNameRequired
	}

	if user.Password == "" {
		return ErrPasswordRequired
	}

	return nil
}

func createAndSetAuthCookie(userId int64, w http.ResponseWriter) (string, error) {
	secret := []byte(config.Envs.JWTSecret)
	token, err := CreateJWT(secret, userId)
	if err != nil {
		return "", err
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "Authorization",
		Value: token,
	})

	return token, nil
}
