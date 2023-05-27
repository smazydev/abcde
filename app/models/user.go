package models

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	BusinessID uuid.UUID `json:"businessId"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Roles      string    `json:"roles"`
}
