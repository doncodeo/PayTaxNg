package main

import (
	"os"
	
	"github.com/gin-gonic/gin"
	"nigeria-tax-api/internal/routes"
) 

func main() {
	r := gin.Default()

	routes.RegisterRoutes(r)
	port := os.Getenv("PORT")
	if port  == ""{
		port = "8080" // Fallback to default port if not specified
	}

	r.Run(":" + port) // listen and serve on
}