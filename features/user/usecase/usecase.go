package usecase

import (
	"baberlab/domain"
	"fmt"
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
func (usecase *userUsecase) CreateUser(u domain.User) (domain.UserResponse, error) {
	user, err := usecase.repo.CreateUser(u)
	if err != nil {
		return domain.UserResponse{}, err
	}

	userWithoutPassword := domain.UserResponse{}
	fmt.Println(user)
	// utils.MapStruct(user, &userWithoutPassword)
	return userWithoutPassword, err
}

// GetUser ..
func (usecase *userUsecase) GetUserById(userID string) (domain.User, error) {
	return usecase.repo.GetUserById(userID)
	// user, err := usecase.repo.GetUserById(userID)
	// if err != nil {
	// 	return domain.UserResponse{}, err
	// }

	// userWithoutPassword := domain.UserResponse{}
	// utils.MapStruct(user, &userWithoutPassword)
	// return userWithoutPassword, err
}
