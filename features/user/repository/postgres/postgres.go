package repository

import (
	"baberlab/domain"
	"log"

	"gorm.io/gorm"
)

// userRepository ..
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository ..
func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

// CreateUser ..
func (r *userRepository) CreateUser(user domain.User) error {
	if err := r.db.Debug().Create(&user).
		Error; err != nil {

		log.Println("CreateUser error: ", err)
		return err
	}

	return nil
}

// GetUserByID ..
func (r *userRepository) GetUserByID(userID string) (domain.User, error) {
	user := domain.User{}

	if err := r.db.Debug().Table(`users`).Where("user_id = ?", userID).Scan(&user).Error; err != nil {
		log.Println("GetUserById error: ", err)
		return user, err
	}

	return user, nil
}

// GetUser ..
func (r *userRepository) GetUser() ([]domain.User, error) {

	user := []domain.User{}

	if err := r.db.Table(`users`).
		Scan(&user).
		Error; err != nil {

		log.Println("GetUser error: ", err)
		return user, err
	}

	return user, nil
}
