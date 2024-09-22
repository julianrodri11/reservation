package controllers

import (
	"net/http"
	"reservation-system/models/dto"
	"reservation-system/services"

	"github.com/kataras/iris/v12"
	//"reservation-system/utils"
)

type UserController struct {
	Service *services.UserService
}

// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags users
// @Produce json
// @Success 200 {array} entity.User
// @Router /users [get]
func (c *UserController) Register(ctx iris.Context) {
	var user dto.UserDTO
	err := ctx.ReadJSON(&user)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid data"})
		return
	}
	c.Service.Register(user)
	ctx.JSON(iris.Map{"message": "User registered successfully"})
}

// Ruta para obtener todos los usuarios
func (c *UserController) GetAllUsers(ctx iris.Context) {
	users, err := c.Service.GetAllUsers()
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	ctx.JSON(users)
}

// @Summary Get user by email
// @Description Get user details by email
// @Tags users
// @Produce json
// @Param email path string true "User Email"
// @Success 200 {object} entity.User
// @Router /users/email/{email} [get]
func (c *UserController) GetUserByEmail(ctx iris.Context) {
	email := ctx.Params().Get("email")
	user, err := c.Service.GetUserByEmail(email)
	if err != nil {
		ctx.StatusCode(http.StatusNotFound)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	ctx.JSON(user)
}
