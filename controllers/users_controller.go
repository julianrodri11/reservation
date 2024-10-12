package controllers

import (
	"reservation-system/models/dto"
	"reservation-system/services"
	"reservation-system/utils"
	"strconv"

	"github.com/kataras/iris/v12"
)

type UserController struct {
	Service *services.UserService
}

func (c *UserController) RegisterUser(ctx iris.Context) {
	var user dto.UserDTO
	err := ctx.ReadJSON(&user)
	utils.HandleBadRequest(ctx, err)
	// Intentar registrar al usuario usando el servicio
	createdUser, err := c.Service.Register(user)

	if err != nil {
		// Si el servicio retorna un error, manejarlo aquí y enviar una respuesta adecuada
		utils.HandleFound(ctx, err)
		return
	}

	// Si no hay errores, retornar el usuario creado con un código de éxito
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"message": "User registered successfully", "user": createdUser})
}

func (c *UserController) UpdateUser(ctx iris.Context) {
	var user dto.UserDTO
	err := ctx.ReadJSON(&user)
	utils.HandleBadRequest(ctx, err)
	c.Service.Update(user)

	ctx.JSON(iris.Map{"message": "User updated successfully"})
}

func (c *UserController) GetAllUsers(ctx iris.Context) {
	users, err := c.Service.GetAllUsers()
	utils.HandleInternalServerError(ctx, err)
	ctx.JSON(users)
}

func (c *UserController) GetUserByEmail(ctx iris.Context) {
	email := ctx.Params().Get("email")
	user, err := c.Service.GetUserByEmail(email)
	utils.HandleNotFound(ctx, err)
	ctx.JSON(user)
}

func (c *UserController) DeleteUser(ctx iris.Context) {
	id := ctx.Params().Get("id")
	userID, err := strconv.Atoi(id)
	user, err := c.Service.DeleteUserById(userID)
	utils.HandleNotFound(ctx, err)
	ctx.JSON(user)
}
