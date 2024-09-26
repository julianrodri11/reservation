package controllers

import (
	"net/http"
	"reservation-system/config"
	"reservation-system/models/dto"
	"reservation-system/services"

	"github.com/kataras/iris/v12"
)

type LoginController struct {
	Service services.LoginService
}

func (c *LoginController) Login(ctx iris.Context) {
	var loginDTO dto.LoginDTO
	err := ctx.ReadJSON(&loginDTO)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid input"})
		return
	}

	// Obtener el usuario desde el servicio de login
	user, err := c.Service.Login(loginDTO)
	if err != nil {
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.JSON(iris.Map{"error": err.Error()})
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
