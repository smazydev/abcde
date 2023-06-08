package services

import (
	"github.com/google/uuid"
	"github.com/smazydev/abcde/app/models"
	"github.com/smazydev/abcde/app/repositories"
)

type ProductService struct {
	container         *Container
	productRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (s *ProductService) CreateProduct(product *models.Product) (*models.Product, error) {
	// Perform any necessary business logic or validation here
	return s.productRepository.Create(product)

}

func (s *ProductService) GetProductByID(id uuid.UUID) (*models.Product, error) {
	return s.productRepository.GetByID(id)
}

func (s *ProductService) UpdateProduct(product *models.Product, productId uuid.UUID) (*models.Product, error) {
	// Perform any necessary business logic or validation here
	updatedProduct, err := s.productRepository.Update(product, productId)
	return updatedProduct, err
}

func (s *ProductService) DeleteProduct(productId uuid.UUID) (*models.Product, error) {
	deletedProduct, err := s.productRepository.Delete(productId)
	return deletedProduct, err
}
