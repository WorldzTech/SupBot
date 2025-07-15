package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"govno/internal/postrgesql"
	"log"
)


func main() {
	r := gin.Default()

	ctx := context.Background()
	_, err := postrgesql.NewRepository(ctx, "postgresql://79.174.88.248:18477/oprac?user=ozonprac&password=OzonPrac2025!")

	if err != nil {
		log.Fatal(err)
	}

	authGroup := r.Group("/auth")
	authGroup.POST("", )
}