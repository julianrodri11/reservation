package routes

import (
	"reservation-system/controllers"

	"github.com/kataras/iris/v12"
)

func UserRoutes(app *iris.Application, userController controllers.UserController) {
	// @Router /users [get]
	users := app.Party("/users")
	{
		// @Router / [get]
		users.Get("/", userController.GetAllUsers) // Ruta para obtener todos los usuarios
		// @Router /users/register/ [post]
		users.Post("/register", userController.Register)
		// @Router /users/email/{email} [get]
		users.Get("/email/{email:string}", userController.GetUserByEmail) // Ruta para obtener usuario por correo
	}
}
