package utils

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

// Gestion de errores
func HandleBadRequest(ctx iris.Context, err error) {
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid data"})
		return
	}
}

func HandleUnauthorized(ctx iris.Context, err error) {
	if err != nil {
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
}

func HandleInternalServerError(ctx iris.Context, err error) {
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
}

func HandleNotFound(ctx iris.Context, err error) {
	if err != nil {
		ctx.StatusCode(http.StatusNotFound)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
}
func HandleFound(ctx iris.Context, err error) {
	if err != nil {
		ctx.StatusCode(http.StatusFound)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
}

// HandleValidationError procesa los errores de validación y los organiza en un mapa estructurado por campo
func HandleValidationError(ctx iris.Context, err error) bool {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		// Crear un mapa para organizar los errores de validación por campo
		errorMap := make(map[string]string)
		for _, e := range validationErrors {
			// Obtener el nombre del campo y el mensaje de error
			fieldName := e.Field()
			errorMsg := e.Error() // Obtenemos el mensaje de error sin traducir

			// Agregar el campo y su error al mapa
			errorMap[fieldName] = errorMsg
		}

		// Responder con un código de estado 400 y los detalles de los errores
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Validation failed", "details": errorMap})
		return true
	}

	return false
}
