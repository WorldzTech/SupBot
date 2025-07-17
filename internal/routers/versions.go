package routers

import (
	"govno/internal/handlers"
	"govno/internal/repo"
	"govno/internal/services"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	return gin.Default()
}

func GetAPIv1() *gin.Engine {
	router := NewRouter()

	adminRepo := repo.AdminRepo{}
	adminService := services.AdminService{Repo: adminRepo}
	authHandler := handlers.NewAuthHandler(adminService, "your_strong_secret_key")

	router.POST("/auth/login", authHandler.GenJWTToken)

	adminGroup := router.Group("/admins")
	adminGroup.Use(authHandler.JWTAuthMiddleware())
	{
		adminHandlers := handlers.AdminHandlers{Service: adminService}
		adminGroup.GET("/list", adminHandlers.GetAdminList)
	}

	return router
}
