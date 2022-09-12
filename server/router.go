package server

import (
	"github.com/ayarhlaine/go_api/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", controllers.Pong)

	return router
}
