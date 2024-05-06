package main

import (
	"imgtool/handler"
	"imgtool/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Welcome to Image Tool API, use the /resize endpoint to resize an image."})
	})

	router.POST("/resize", handler.ResizeImage)

	router.Run()
}
