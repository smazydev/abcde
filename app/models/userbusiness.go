package models

import (
	"github.com/google/uuid"
)

type UserBusiness struct {
	UserID     uuid.UUID `gorm:"type:uuid"`
	BusinessID uuid.UUID `gorm:"type:uuid"`
}
