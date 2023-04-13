package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hy-reza/mygram-api/helper"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome to MyGram API!!")
	})

	err := router.Run(":8080")
	helper.ErrorPanic(err)
}
