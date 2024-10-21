package controllers

import (
	"reservation-system/models/dto"
	"reservation-system/services"
	"reservation-system/utils"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

// Crear una instancia del validador
var validate = validator.New()

type UserController struct {
	Service *services.UserService
}

func init() {
	// Registrar la validación personalizada para el correo
	validate.RegisterValidation("customEmail", utils.ValidateEmail)
}

func (c *UserController) RegisterUser(ctx iris.Context) {
	var user dto.UserDTO
	err := ctx.ReadJSON(&user)
	if err != nil {
		utils.HandleBadRequest(ctx, err)
		return
	}
	// Validar el DTO
	err = validate.Struct(user)
	if err != nil {
		// Si las validaciones fallan, devolver un error 400 con los detalles
		utils.HandleValidationsBadRequest(ctx, err)
		return
	}
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
