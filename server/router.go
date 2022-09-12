package server

import (
	"github.com/ayarhlaine/go_api/controllers/ping"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", ping.Pong)

	return router
}
