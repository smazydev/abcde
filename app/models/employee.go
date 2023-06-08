package models

import (
	"github.com/google/uuid"
)

type Employee struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name       string    `json:"name"`
	Position   string    `json:"position"`
	BusinessID uuid.UUID `gorm:"type:uuid" json:"businessId"`
	Business   Business  `gorm:"foreignKey:BusinessID" json:"business"`
}
