package usecase

import (
	"baberlab/domain"
)

// userUsecase ..
type userUsecase struct {
	repo domain.UserRepository
}

// NewUserUsecase ..
func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

// CreateUser ..
func (usecase *userUsecase) CreateUser(user domain.User) error {
	return usecase.repo.CreateUser(user)
}

// GetUserByID ..
func (usecase *userUsecase) GetUserByID(userID string) (domain.User, error) {
	return usecase.repo.GetUserByID(userID)

}

// GetUser ..
func (usecase *userUsecase) GetUser() ([]domain.User, error) {
	return usecase.repo.GetUser()
}
