package config

import (
	"net/http"
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

	// Parsear el token y validarlo
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, iris.NewProblem().Status(http.StatusUnauthorized).Detail("Método de firma no válido")
		}
		return jwtSecret, nil // jwtSecret es tu clave secreta para firmar los tokens
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Guardar la información del token en el contexto para usarla más adelante
		ctx.Values().Set("userID", claims["id"])
		ctx.Next()
	} else {
		// Responder con un mensaje claro si el token no es válido
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.JSON(iris.Map{
			"error":   "Token inválido",
			"details": err.Error(),
		})
		return
	}
}

// GenerateJWT genera un token JWT para un usuario dado
func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expiración de 24 horas
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar el token con la clave secreta
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
