package httpserver

import (
	"gin_auth/internal/app"
	"gin_auth/internal/middleware"
	"gin_auth/internal/user"
	"net/http"

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
	r.POST("/login", userHandler.Login)


	api := r.Group("/api")
	api.Use(middleware.AuthRequired(a.Config.JWT_SECRET)) // protecting routes by adding middleware


	api.GET("/files", func (c *gin.Context) {
		userId, _ := middleware.GetUserId(c)
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
			"userId": userId,
			"files": []any{},
		})
	})

	api.GET("products", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
			"products": []any{},
		})
	})

	return r
}
