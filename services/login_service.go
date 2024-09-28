package services

import (
	"errors"
	"reservation-system/models/dto"
	"reservation-system/models/entity"
	"reservation-system/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type LoginService struct {
	Repo      repository.LoginRepository
	JWTSecret string // Este es el secreto utilizado para firmar el JWT
}

func (s *LoginService) Login(loginDTO dto.LoginDTO) (entity.Users, error) {
	// Buscar usuario por email
	user, err := s.Repo.FindByEmailAndPassword(loginDTO.Email, loginDTO.Password)
	if err != nil {
		return entity.Users{}, errors.New("usuario o contraseña incorrectos")
	}

	// Si todo es correcto, devolver el usuario
	return user, nil
}

// generateJWT genera un token JWT con los datos del usuario
func (s *LoginService) generateJWT(user entity.Users) (string, error) {
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Expiración de 24 horas
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar el token con la clave secreta
	tokenString, err := token.SignedString([]byte(s.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
