package models

import "github.com/google/uuid"

type Business struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name string    `json:"name"`
}
