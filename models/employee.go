package models

import "github.com/google/uuid"

type Employee struct {
	Id      uuid.UUID `gorm:"primaryKey"`
	Name    string
	Address Address `gorm:"foreignKey:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
