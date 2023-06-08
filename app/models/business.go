package models

import "github.com/google/uuid"

type Business struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string     `json:"name"`
	Employees []Employee `json:"employees" gorm:"foreignKey:BusinessID"`
	//Products  []Product  `json:"products" gorm:"foreignkey:BusinessID;references:ID"`
	OwnerID uuid.UUID `gorm:"type:uuid" json:"-"`
	Owner   User      `gorm:"foreignKey:OwnerID" json:"owner"`
}
