package models

import (
	"github.com/google/uuid"
)

type Product struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
}
