package handlers

import (
    "github.com/gin-gonic/gin"
    "govno/internal/services"
)

type AdminHandlers struct {
	Service services.AdminService
}

func (handler *AdminHandlers) GetAdminList(c *gin.Context) {
	admins, err := handler.Service.GetAdminsList()
	if err != nil {
		c.JSON(400, gin.H{"error": err})
	}

	c.JSON(200, admins)
}