package repositories

import (
	"github.com/smazydev/abcde/app/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	GetByID(id string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User, id string) error
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

func (r *userRepository) GetByID(id string) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
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

func (r *userRepository) Update(user *models.User, id string) error {
	return r.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"name":       user.Name,
		"username":   user.Username,
		"email":      user.Email,
		"businesses": user.Businesses,
		"roles":      user.Roles,
		// Add other fields that can be updated
	}).Error
}

func (r *userRepository) Delete(userID string) error {
	return r.db.Delete(&models.User{}, userID).Error
}
