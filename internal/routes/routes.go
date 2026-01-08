package routes

import (
	"github.com/gin-gonic/gin"
	"nigeria-tax-api/internal/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/ping", handlers.PingHandler)
	r.POST("/api/tax/calculate", handlers.CalculateTax)
}