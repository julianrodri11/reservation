package main

import (
	"reservation-system/config"
	"reservation-system/controllers" // Importar el paquete de Swagger
	"reservation-system/repository"
	"reservation-system/routes"
	"reservation-system/services"

	"github.com/kataras/iris/v12"
)

func main() {

	app := iris.New()

	//Variables de entorno
	config.LoadEnv()

	// Iniciar la base de datos
	db := config.InitDB()

	// Inicializar repositorios, servicios y controladores
	userRepo := repository.UserRepository{DB: db}
	userService := services.UserService{Repo: &userRepo}
	userController := controllers.UserController{Service: &userService}

	// Instanciar LoginController
	loginRepo := repository.LoginRepository{DB: db}
	loginService := services.LoginService{Repo: loginRepo} // Asumiendo que tienes un LoginService
	loginController := controllers.LoginController{Service: loginService}

	// Instanciar CompanyController
	companyRepo := repository.CompanyRepository{DB: db}
	companyService := services.CompanyService{Repo: &companyRepo}
	companyController := controllers.CompanyController{Service: &companyService}

	// Configurar rutas
	routes.ConfigureRoutes(app, userController, loginController, companyController)

	// Iniciar servidor
	app.Listen(":8080")
}
