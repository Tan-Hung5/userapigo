package repository

import (
	"errors"
	"net/http"
	"strconv"
	"sync"
	"usersservice/domain/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
	mu sync.Mutex
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUsers() ([]models.User, error) {

	var users []models.User

	err := r.db.Table("user").Unscoped().Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserByID(userID int) (models.User, error) {
	var user models.User
	err := r.db.Table("user").First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByEmail(userEmail string) (models.User, error) {
	var user models.User
	err := r.db.Table("user").Where("email= ?", userEmail).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) AddUser(user models.CreateUser) (int, error) {
	err := r.db.Table("user").Create(&user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *UserRepository) UpdateUser(userID int, updatedUser models.User) error {
	var existingUser models.User

	// Kiểm tra xem người dùng có tồn tại không
	if err := r.db.First(&existingUser, userID).Error; err != nil {
		r.db.Table("user").Where("id = ? ", userID).Model(&existingUser).Updates(updatedUser)
	}

	return nil
}

func (r *UserRepository) DeleteUser(w http.ResponseWriter, res *http.Request) {
	vars := mux.Vars(res)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	result := r.db.Table("user").Delete(&models.User{}, userID)
	if result.RowsAffected == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
