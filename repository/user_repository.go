package repository

import (
	"reservation-system/models/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// crear usuario
func (r *UserRepository) Create(user entity.User) error {
	return r.DB.Create(&user).Error
}

// consultar por correo
func (r *UserRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

// Consultar todos los usuarios
func (r *UserRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	err := r.DB.Find(&users).Error
	return users, err
}
