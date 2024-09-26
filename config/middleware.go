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

	// Parse the JWT string and validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, iris.NewProblem().Status(http.StatusUnauthorized).Detail("Invalid signing method")
		}
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Save the token data into the request context
		ctx.Values().Set("userID", claims["id"])
		ctx.Next()
	} else {
		ctx.StopWithStatus(http.StatusUnauthorized)
		ctx.JSON(iris.Map{"error": "Invalid token", "details": err.Error()})
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
