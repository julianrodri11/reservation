package docs

import (
	"reservation-system/cmd/app/docs"

	"github.com/iris-contrib/swagger/v12" // Importa el paquete correcto
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
)

func ConfigureSwagger(app *iris.Application) {
	// Registrar la documentación de Swag
	docs.SwaggerInfo.Title = "Reservation System API"
	docs.SwaggerInfo.Description = "API para gestionar reservas de canchas"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http"}
	// Configurar la ruta para Swagger
	app.Get("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
}
