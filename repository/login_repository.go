package repository

import (
	"errors"
	"reservation-system/models/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRepository struct {
	DB *gorm.DB
}

// Buscar por correo y contraseña
func (r *LoginRepository) FindByEmailAndPassword(email string, password string) (entity.Users, error) {
	var user entity.Users
	// Buscar usuario por correo
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	// Verificar si la contraseña es correcta
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("invalid password")
	}

	// Retornar el usuario si la contraseña es válida
	return user, nil
}
