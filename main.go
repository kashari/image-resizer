package main

import (
	"imgtool/handler"
	"imgtool/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	router.POST("/resize", handler.ResizeImage)

	router.Run(":8080")
}
