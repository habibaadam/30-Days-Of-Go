package httpserver

import (
	"gin_auth/internal/app"
	"gin_auth/internal/user"

	"github.com/gin-gonic/gin"
)


func NewRouter(a *app.App) *gin.Engine {
	r := gin.New() // empty/ blank engine

	//gin.default() is exactly this under the hood
	r.Use(gin.Logger())

	r.Use(gin.Recovery()) // catches any panic happening & returns 500 instead of crashing.

	r.GET("/health",  health)

	userRepo := user.UserRepo(a.Db)
	userService := user.NewService(userRepo, a.Config.JWT_SECRET)
	userHandler := user.NewHandler(userService)

	r.POST("/register", userHandler.Register)

	return r
}