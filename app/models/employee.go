package models

import (
	"github.com/google/uuid"
)

type Employee struct {
	ID         uuid.UUID `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	Position   string    `json:"position"`
	BusinessID int       `json:"businessId"`
}
