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
	// Validar el DTO usando la función de utilidades
	err = validate.Struct(user)
	if utils.HandleValidationError(ctx, err) {
		return
	}
	// Intentar registrar al usuario usando el servicio
	createdUser, err := c.Service.Register(user)
	if err != nil {
		utils.HandleFound(ctx, err)
		return
	}
	// Si no hay errores, retornar el usuario creado con un código de éxito
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"message": "Register saved successfully", "user": createdUser})
}

func (c *UserController) UpdateUser(ctx iris.Context) {

	//leer Json
	var user dto.UserDTO
	err := ctx.ReadJSON(&user)
	if err != nil {
		utils.HandleBadRequest(ctx, err)
		return
	}
	// Validar el DTO usando la función de utilidades
	err = validate.Struct(user)
	if utils.HandleValidationError(ctx, err) {
		// Si hubo errores de validación, ya se manejaron, simplemente retornar
		return
	}
	// Intentar actualizar al usuario usando el servicio
	updateUser, err := c.Service.Update(user)
	if err != nil {
		utils.HandleNotFound(ctx, err)
		return
	}
	// Si no hay errores, retornar el usuario creado con un código de éxito
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Register updated successfully", "user": updateUser})
}

func (c *UserController) GetAllUsers(ctx iris.Context) {
	users, err := c.Service.GetAllUsers()
	utils.HandleInternalServerError(ctx, err)
	ctx.JSON(users)
}

func (c *UserController) GetUserByEmail(ctx iris.Context) {
	email := ctx.Params().Get("email")
	user, err := c.Service.GetUserByEmail(email)
	if err != nil {
		utils.HandleNotFound(ctx, err)
		return
	}
	ctx.StatusCode(iris.StatusFound)
	ctx.JSON(iris.Map{"message": "Register found successfully", "user": user})

}

func (c *UserController) DeleteUser(ctx iris.Context) {
	id := ctx.Params().Get("id")
	userID, err := strconv.Atoi(id)
	user, err := c.Service.DeleteUserById(userID)
	if err != nil {
		utils.HandleNotFound(ctx, err)
		return
	}
	ctx.StatusCode(iris.StatusFound)
	ctx.JSON(iris.Map{"message": "Register deleted successfully", "user": user})
}
