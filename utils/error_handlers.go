package utils

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

// Gestion de errores
func HandleBadRequest(ctx iris.Context, err error) {
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid data"})
		return
	}
}

func HandleUnauthorized(ctx iris.Context, err error) {
	ctx.StatusCode(http.StatusUnauthorized)
	ctx.JSON(iris.Map{"error": err.Error()})
}

func HandleInternalServerError(ctx iris.Context, err error) {
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
}

func HandleNotFound(ctx iris.Context, err error) {
	if err != nil {
		ctx.StatusCode(http.StatusNotFound)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
}
