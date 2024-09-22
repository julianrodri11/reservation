package routes

import (
	"reservation-system/controllers"

	"github.com/kataras/iris/v12"
)

func ConfigureRoutes(app *iris.Application, userController controllers.UserController) {
	users := app.Party("/users")
	{
		users.Get("/", userController.GetAllUsers)
		users.Post("/register", userController.Register)
		users.Get("/email/{email:string}", userController.GetUserByEmail)
		users.Delete("/id/{id:int}", userController.DeleteUser)
	}
	entities := app.Party("/entities")
	{
		entities.Get("/all", userController.GetAllUsers)
		entities.Post("/register", userController.Register)
		entities.Get("/email/{email:string}", userController.GetUserByEmail)
	}
}
