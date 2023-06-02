package repositories

import (
	"github.com/smazydev/abcde/app/models"
	"gorm.io/gorm"
)

type BusinessRepository interface {
	Create(business *models.Business) error
	GetByID(id uint) (*models.Business, error)
	GetBusinessesByOwnerID(id string) ([]*models.Business, error)
}

type businessRepository struct {
	db *gorm.DB
}

func NewBusinessRepository(db *gorm.DB) BusinessRepository {
	return &businessRepository{
		db: db,
	}
}

func (r *businessRepository) Create(business *models.Business) error {

	return r.db.Create(business).Error
}

func (r *businessRepository) GetByID(id uint) (*models.Business, error) {
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
