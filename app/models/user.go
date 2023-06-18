package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Businesses []Business `json:"businesses" gorm:"foreignKey:OwnerID"`
	Name       string     `json:"name"`
	Username   string     `json:"username"`
	Email      string     `gorm:"unique" json:"email"`
	Password   string     `json:"-"`
	Roles      string     `json:"roles"`
}
