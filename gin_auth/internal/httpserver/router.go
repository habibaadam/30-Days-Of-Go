package httpserver

import "github.com/gin-gonic/gin"


func NewRouter() *gin.Engine {
	r := gin.New() // empty/ blank engine

	//gin.default() is exactly this under the hood
	r.Use(gin.Logger())

	r.Use(gin.Recovery()) // catches any panic happening & returns 500 instead of crashing.

	r.GET("/health",  health)

	return r
}