package entity

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Address   string         `gorm:"type:varchar(255)" json:"address"`
	Phone     string         `gorm:"type:varchar(15)" json:"phone"`
	Email     string         `gorm:"type:varchar(100);unique;not null" json:"email"`
	IsActive  bool           `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Users     []Users        `gorm:"foreignKey:CompanyID" json:"users,omitempty"` // Relación con los usuarios
}
