package repositories

import (
	"github.com/google/uuid"
	"github.com/smazydev/abcde/app/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) (*models.Product, error)
	GetByID(id uuid.UUID) (*models.Product, error)
	Update(product *models.Product, id uuid.UUID) (*models.Product, error)
	Delete(productId uuid.UUID) (*models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) Create(product *models.Product) (*models.Product, error) {
	err := r.db.Create(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) GetByID(id uuid.UUID) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Businesses").First(&product, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Update(product *models.Product, businessId uuid.UUID) (*models.Product, error) {
	if err := r.db.Model(&models.Product{}).Where("id = ?", businessId).Updates(map[string]interface{}{
		"name":        product.Name,
		"description": product.Description,
		"images":      product.Images,
		"businessId":  product.BusinessID,
	}).Error; err != nil {
		return nil, err
	}
	// Retrieve the updated user from the database
	updatedProduct := &models.Product{}
	if err := r.db.First(updatedProduct, businessId).Error; err != nil {
		return nil, err
	}
	return updatedProduct, nil
}

func (r *productRepository) Delete(productId uuid.UUID) (*models.Product, error) {
	product := &models.Product{}
	if err := r.db.First(product, productId).Error; err != nil {
		return nil, err
	}

	if err := r.db.Delete(&models.Product{}, productId).Error; err != nil {
		return nil, err
	}

	return product, nil
}
