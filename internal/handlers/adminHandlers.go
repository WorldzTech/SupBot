package handlers

import (
	"govno/internal/services"

	"github.com/gin-gonic/gin"
)

type AdminHandlers struct {
	Service services.AdminService
}

func (handler *AdminHandlers) GetAdminList(c *gin.Context) {
	admins, err := handler.Service.GetAdminsList() // Ошибка здесь
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()}) // Добавим .Error() для преобразования в строку
		return
	}

	c.JSON(200, admins)
}
