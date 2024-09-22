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

// Consultar todos los usuarios
func (r *UserRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	err := r.DB.Find(&users).Error
	return users, err
}

// consultar por correo
func (r *UserRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

// consultar por id
func (r *UserRepository) FindById(id int) (entity.User, error) {
	var user entity.User
	err := r.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

// eliminar por id
func (r *UserRepository) DeleteUser(id int) (entity.User, error) {
	var user entity.User

	// Buscar el usuario por ID
	err := r.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	// Si lo encuentra, lo elimina
	err = r.DB.Delete(&user).Error
	return user, err
}
