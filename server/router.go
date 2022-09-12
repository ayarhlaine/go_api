package server

import (
	"net/http"

	"github.com/ayarhlaine/go_api/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// Render static files
	router.Static("/assets", "./static/assets")
	router.LoadHTMLGlob("static/*.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "This is an index page...",
		})
	})

	// API Routes
	router.GET("/ping", controllers.Pong)

	return router
}
