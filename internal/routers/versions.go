package routers

import (
	"github.com/gin-gonic/gin"
	"govno/internal/handlers"
)

func NewRouter() *gin.Engine {
	return gin.Default()
}

func GetAPIv1() *gin.Engine {
	router := NewRouter()

	adminHandlers := handlers.AdminHandlers{}
	adminGroup := router.Group("/admins")
	adminGroup.GET("/list", adminHandlers.GetAdminList)

	return router
}