package models

import "github.com/google/uuid"

type Business struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string     `json:"name"`
	Employees []Employee `json:"employees" gorm:"many2many:business_employees;"`
	OwnerID   uuid.UUID  `gorm:"type:uuid" json:"-"`
	User      User       `gorm:"foreignKey:OwnerID" json:"Owner"`
}
