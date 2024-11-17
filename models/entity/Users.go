package entity

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID           uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string         `gorm:"size:100;not null" json:"name"`
	Email        string         `gorm:"size:100;unique;not null" json:"email"`
	Password     string         `gorm:"size:100;not null" json:"password"`
	CompanyID    uint           `json:"company_id"` // Relación con Company
	Company      Company        `gorm:"foreignKey:CompanyID" json:"company"`
	IsActive     bool           `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Reservations []Reservations `gorm:"foreignKey:UserID" json:"reservations"`
}
