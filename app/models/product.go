package models

import (
	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name        string    `json:"productName"`
	Description string    `json:"productDescription"`
	Images      []string  `gorm:"type:text" json:"images"`
	BusinessID  uuid.UUID `gorm:"type:uuid" json:"businessId"`
	Business    *Business `gorm:"foreignKey:BusinessID" json:"business"`
}
