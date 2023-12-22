package services

import (
	"net/http"
	"strconv"
	"usersservice/domain/models"
	"usersservice/repository"

	"github.com/gorilla/mux"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return s.userRepository.GetUsers()
}

func (s *UserService) GetUserByID(userID int) (models.User, error) {
	return s.userRepository.GetUserByID(userID)
}

func (s *UserService) GetUserByEmail(userEmail string) (models.User, error) {
	return s.userRepository.GetUserByEmail(userEmail)
}

func (s *UserService) AddUser(user models.CreateUser) (int, error) {
	return s.userRepository.AddUser(user)
}

func (s *UserService) UpdateUser(userID int, updatedUser models.User) error {
	return s.userRepository.UpdateUser(userID, updatedUser)
}

func (s *UserService) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	_, err = s.userRepository.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	s.userRepository.DeleteUser(w, r)
}

// Additional methods for more complex business logic if needed
