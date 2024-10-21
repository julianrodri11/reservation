package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Custom validation function for email
func ValidateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	// Ejemplo: Validar que el correo termine en un dominio específico (por ejemplo, @example.com)
	// Cambia el patrón de acuerdo a tus necesidades
	//regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@example\.com$`)
	// Expresión regular para validar el formato de un correo electrónico
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}
