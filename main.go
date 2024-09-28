package main

import (
	"reservation-system/config"
	"reservation-system/controllers" // Importar el paquete de Swagger
	"reservation-system/models/entity"
	"reservation-system/repository"
	"reservation-system/routes"
	"reservation-system/services"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	// Iniciar la base de datos
	db := config.InitDB()

	// Migrar el modelo de Usuario
	//db.AutoMigrate(&models.Usuario{})

	// Migrar las entidades
	db.AutoMigrate(&entity.Users{}, &entity.Reservations{})

	// Inicializar repositorios, servicios y controladores
	userRepo := repository.UserRepository{DB: db}
	userService := services.UserService{Repo: &userRepo}
	userController := controllers.UserController{Service: &userService}

	// Instanciar LoginController
	loginRepo := repository.LoginRepository{DB: db}
	loginService := services.LoginService{Repo: loginRepo} // Asumiendo que tienes un LoginService
	loginController := controllers.LoginController{Service: loginService}

	// Configurar rutas
	routes.ConfigureRoutes(app, userController, loginController)

	// Iniciar servidor
	app.Listen(":8080")
}
