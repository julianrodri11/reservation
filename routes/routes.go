package routes

import (
	"reservation-system/controllers"

	"reservation-system/config"

	"github.com/kataras/iris/v12"
)

func ConfigureRoutes(app *iris.Application,
	userController controllers.UserController,
	loginController controllers.LoginController) {
	users := app.Party("/users")
	{
		users.Get("/", config.JWTMiddleware, userController.GetAllUsers)
		users.Post("/register", config.JWTMiddleware, userController.Register)
		users.Post("/update", config.JWTMiddleware, userController.Update)
		users.Get("/email/{email:string}", config.JWTMiddleware, userController.GetUserByEmail)
		users.Delete("/id/{id:int}", config.JWTMiddleware, userController.DeleteUser)
	}
	app.Post("/login", loginController.Login) // Ruta para iniciar sesión (sin middleware JWT)

	entities := app.Party("/entities")
	{
		entities.Get("/all", config.JWTMiddleware, userController.GetAllUsers)
		entities.Post("/register", config.JWTMiddleware, userController.Register)
		entities.Get("/email/{email:string}", config.JWTMiddleware, userController.GetUserByEmail)
	}
}
