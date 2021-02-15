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
func (r *userRepository) CreateUser(user domain.User) (domain.User, error) {
	if err := r.db.
		Create(&user).
		Error; err != nil {

		log.Println("CreateUser error: ", err)
		return user, err
	}

	return user, nil
}
