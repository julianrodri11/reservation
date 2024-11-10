package controllers

import (
	"net/http"
	"reservation-system/config"
	"reservation-system/models/dto"
	"reservation-system/services"
	"reservation-system/utils"

	"github.com/kataras/iris/v12"
)

type LoginController struct {
	Service services.LoginService
}

func (c *LoginController) Login(ctx iris.Context) {
	// Validar objeto json
	var loginDTO dto.LoginDTO
	err := ctx.ReadJSON(&loginDTO)
	if err != nil {
		utils.HandleBadRequest(ctx, err)
		return
	}
	// Validar el DTO usando la función de utilidades
	err = validate.Struct(loginDTO)
	if utils.HandleValidationError(ctx, err) {
		return
	}
	// Obtener el usuario desde el servicio de login
	user, err := c.Service.Login(loginDTO)
	if err != nil {
		utils.HandleUnauthorized(ctx, err)
		return
	}

	// Generar el JWT para el usuario usando el ID
	token, err := config.GenerateJWT(user.ID)
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(iris.Map{"token": token})
}
