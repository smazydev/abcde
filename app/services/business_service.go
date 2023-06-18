package services

import (
	"github.com/google/uuid"
	"github.com/smazydev/abcde/app/models"
	"github.com/smazydev/abcde/app/repositories"
)

type BusinessService struct {
	container          *Container
	businessRepository repositories.BusinessRepository
}

func NewBusinessService(businessRepository repositories.BusinessRepository) *BusinessService {
	return &BusinessService{businessRepository: businessRepository}
}

func (s *BusinessService) Create(business *models.Business) (*models.Business, error) {
	// Perform any necessary business logic or validation here
	createdBusiness, err := s.businessRepository.Create(business)
	createdBusiness.Owner.Businesses = nil
	return createdBusiness, err

}

func (s *BusinessService) GetByID(id uuid.UUID) (*models.Business, error) {
	return s.businessRepository.GetByID(id)
}

func (s *BusinessService) UpdateBusiness(business *models.Business, businessId string) error {
	// Perform any necessary business logic or validation here
	return s.businessRepository.Update(business, businessId)
}

func (s *BusinessService) GetBusinessesByOwnerID(userId string) ([]*models.Business, error) {
	return s.businessRepository.GetBusinessesByOwnerID(userId)
}

// func (s *BusinessService) DeleteBusiness(businessId string) error {
// 	return s.businessRepository.DeleteBusiness(businessId)
// }
