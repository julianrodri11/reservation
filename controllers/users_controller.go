package controllers

import (
	"net/http"
	"reservation-system/models/dto"
	"reservation-system/services"
	"strconv"

	"github.com/kataras/iris/v12"
)

type UserController struct {
	Service *services.UserService
}

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

func (c *UserController) Update(ctx iris.Context) {
	var user dto.UserDTO
	err := ctx.ReadJSON(&user)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid data"})
		return
	}
	c.Service.Update(user)
	ctx.JSON(iris.Map{"message": "User updated successfully"})
}

func (c *UserController) UpdateUser(ctx iris.Context) {
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

func (c *UserController) GetAllUsers(ctx iris.Context) {
	users, err := c.Service.GetAllUsers()
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	ctx.JSON(users)
}

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

func (c *UserController) DeleteUser(ctx iris.Context) {
	id := ctx.Params().Get("id")
	userID, err := strconv.Atoi(id)
	user, err := c.Service.DeleteUserById(userID)
	if err != nil {
		ctx.StatusCode(http.StatusNotFound)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	ctx.JSON(user)
}
