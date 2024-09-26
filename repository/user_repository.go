package repository

import (
	"reservation-system/models/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// crear usuario
func (r *UserRepository) CreateUser(user entity.Users) error {
	return r.DB.Create(&user).Error
}

// Actualizar usuario existente
func (r *UserRepository) UpdateUser(user entity.Users) error {
	return r.DB.Save(&user).Error
}

// Consultar todos los usuarios
func (r *UserRepository) FindAll() ([]entity.Users, error) {
	var users []entity.Users
	err := r.DB.Find(&users).Error
	return users, err
}

// consultar por correo
func (r *UserRepository) FindByEmail(email string) (entity.Users, error) {
	var user entity.Users
	err := r.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

// eliminar por id
func (r *UserRepository) DeleteUser(id int) (entity.Users, error) {
	var user entity.Users

	// Buscar el usuario por ID
	err := r.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	// Si lo encuentra, lo elimina
	err = r.DB.Delete(&user).Error
	return user, err
}

// consultar por id
func (r *UserRepository) FindById(id int) (entity.Users, error) {
	var user entity.Users
	err := r.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

// Verificar si existe un usuario por ID
func (r *UserRepository) UserExistsById(id int) (bool, error) {
	var user entity.Users
	err := r.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		// Si ocurre un error de tipo RecordNotFound, devolver false
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		// Si es otro tipo de error, devolver false y el error
		return false, err
	}
	// Si el usuario es encontrado, devolver true
	return true, nil
}
