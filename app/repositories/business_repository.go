package repositories

import (
	"github.com/google/uuid"
	"github.com/smazydev/abcde/app/models"
	"gorm.io/gorm"
)

type BusinessRepository interface {
	Create(business *models.Business) (*models.Business, error)
	GetByID(id uuid.UUID) (*models.Business, error)
	GetBusinessesByOwnerID(id string) ([]*models.Business, error)
	Update(business *models.Business, id string) error
}

type businessRepository struct {
	db *gorm.DB
}

func NewBusinessRepository(db *gorm.DB) BusinessRepository {
	return &businessRepository{
		db: db,
	}
}

func (r *businessRepository) Create(business *models.Business) (*models.Business, error) {
	err := r.db.Create(business).Error
	if err != nil {
		return nil, err
	}
	return business, nil
}

func (r *businessRepository) GetByID(id uuid.UUID) (*models.Business, error) {
	var business models.Business
	err := r.db.First(&business, id).Error
	if err != nil {
		return nil, err
	}
	return &business, nil
}

func (r *businessRepository) GetBusinessesByOwnerID(id string) ([]*models.Business, error) {
	var businesses []*models.Business
	err := r.db.Where("owner_id = ?", id).Find(&businesses).Error
	if err != nil {
		return nil, err
	}
	return businesses, nil
}

func (r *businessRepository) Update(business *models.Business, id string) error {
	return r.db.Model(&models.User{}).Where("id = ?", business.ID).Updates(map[string]interface{}{
		"name":      business.Name,
		"employees": business.Employees,
		"ownerID":   business.OwnerID,
		// Add other fields that can be updated
	}).Error
}
