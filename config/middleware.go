package config

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kataras/iris/v12"
)

// Secret key used for signing the JWT
var jwtSecret = []byte("your-secret-key")

// Middleware para validar JWT
func JWTMiddleware(ctx iris.Context) {
	tokenString := ctx.GetHeader("Authorization")

	if tokenString == "" {
		// Si no se proporciona el token
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.JSON(iris.Map{"error": "Token no proporcionado", "details": "Debe incluir un token de autorización en la cabecera"})
		return
	}

	// Quitar prefijo "Bearer " si está presente
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Parsear el token y validarlo
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, iris.NewProblem().Status(http.StatusUnauthorized).Detail("Método de firma no válido")
		}
		return jwtSecret, nil // jwtSecret es tu clave secreta para firmar los tokens
	})

	if err != nil {
		// Responder con el error específico del token
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.JSON(iris.Map{
			"error":   "Token inválido",
			"details": err.Error(),
		})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Guardar la información del token en el contexto para usarla más adelante
		ctx.Values().Set("userID", claims["id"])
		ctx.Next()
	} else {
		// Responder con un mensaje claro si el token no es válido
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.JSON(iris.Map{
			"error":   "Token inválido",
			"details": "Token inválido o no autorizado",
		})
		return
	}
}

// GenerateJWT genera un token JWT para un usuario dado
func GenerateJWT(userID uint) (string, error) {

	// Convierte la variable de entorno EXPIRE_TOKEN_HOURS a int
	expireTokenHours, err := strconv.Atoi(os.Getenv("EXPIRE_TOKEN_HOURS"))
	if err != nil {
		log.Printf("Error al convertir EXPIRE_TOKEN_HOURS: %v. Usando valor predeterminado de 24 horas", err)
		//expireTokenHours = 0 // Valor predeterminado de 24 horas si no se puede convertir
	}

	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * time.Duration(expireTokenHours)).Unix(), // Expiración de 24 horas
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar el token con la clave secreta
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
