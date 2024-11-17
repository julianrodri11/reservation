package routes

import (
	"reservation-system/controllers"

	"reservation-system/config"

	"github.com/kataras/iris/v12"
)

func ConfigureRoutes(app *iris.Application,
	userController controllers.UserController,
	loginController controllers.LoginController,
	companyController controllers.CompanyController) {
	users := app.Party("/users")
	{
		users.Get("/all", config.JWTMiddleware, userController.GetAllUsers)
		users.Post("/register", userController.RegisterUser)
		users.Post("/update", config.JWTMiddleware, userController.UpdateUser)
		users.Get("/email/{email:string}", config.JWTMiddleware, userController.GetUserByEmail)
		users.Delete("/id/{id:int}", config.JWTMiddleware, userController.DeleteUser)
	}
	app.Post("/login", loginController.Login) // Ruta para iniciar sesión (sin middleware JWT)

	entities := app.Party("/entities")
	{
		entities.Get("/all", config.JWTMiddleware, userController.GetAllUsers)
		entities.Post("/register", config.JWTMiddleware, userController.RegisterUser)
		entities.Get("/email/{email:string}", config.JWTMiddleware, userController.GetUserByEmail)
	}

	companies := app.Party("/companies")
	{
		companies.Get("/all", config.JWTMiddleware, companyController.GetAllCompanies)
		companies.Post("/register", companyController.RegisterCompany)
		companies.Post("/update", config.JWTMiddleware, companyController.UpdateCompany)
		companies.Get("/email/{email:string}", config.JWTMiddleware, companyController.GetCompanyByEmail)
		companies.Delete("/id/{id:int}", config.JWTMiddleware, companyController.DeleteCompany)
	}
}
