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
	var loginDTO dto.LoginDTO
	err := ctx.ReadJSON(&loginDTO)
	utils.HandleBadRequest(ctx, err)

	// Obtener el usuario desde el servicio de login
	user, err := c.Service.Login(loginDTO)
	utils.HandleUnauthorized(ctx, err)

	// Generar el JWT para el usuario usando el ID
	token, err := config.GenerateJWT(user.ID)
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(iris.Map{"token": token})
}
