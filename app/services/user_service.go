package services

import (
	"github.com/google/uuid"
	"github.com/smazydev/abcde/app/models"
	"github.com/smazydev/abcde/app/repositories"
)

type UserService struct {
	container      *Container
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	// Perform any necessary business logic or validation here
	return s.userRepository.Create(user)

}

func (s *UserService) GetUserByID(id uuid.UUID) (*models.User, error) {
	return s.userRepository.GetByID(id)
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepository.GetByEmail(email)
}

func (s *UserService) UpdateUser(user *models.User, userId uuid.UUID) (*models.User, error) {
	// Perform any necessary business logic or validation here
	updatedUser, err := s.userRepository.Update(user, userId)
	return updatedUser, err
}

func (s *UserService) DeleteUser(userId string) error {
	return s.userRepository.Delete(userId)
}
