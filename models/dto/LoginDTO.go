package dto

// LoginDTO contiene las credenciales de inicio de sesión
type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
