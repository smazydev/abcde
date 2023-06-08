package repositories

import (
	"log"

	"github.com/google/uuid"
	"github.com/smazydev/abcde/app/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	GetByID(id uuid.UUID) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User, id uuid.UUID) (*models.User, error)
	Delete(userId string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *models.User) (*models.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Businesses").First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Update(user *models.User, id uuid.UUID) (*models.User, error) {
	log.Print("ID IN REPO", id)
	if err := r.db.Model(&models.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":       user.Name,
		"username":   user.Username,
		"email":      user.Email,
		"Businesses": user.Businesses,
		"roles":      user.Roles,
		// Add other fields that can be updated
	}).Error; err != nil {
		return nil, err
	}
	// Retrieve the updated user from the database
	updatedUser := &models.User{}
	if err := r.db.First(updatedUser, id).Error; err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (r *userRepository) Delete(userID string) error {
	return r.db.Delete(&models.User{}, userID).Error
}
