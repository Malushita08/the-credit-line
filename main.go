// @title Gin Swagger Example API
// @version 1.0
// @description This is a sample server.

// @host localhost:8080
// @BasePath /
// @schemes http
package main

import (
	"github.com/Malushita08/the-credit-line/database"
	_ "github.com/Malushita08/the-credit-line/docs"
	"github.com/Malushita08/the-credit-line/services"
	"github.com/gin-gonic/gin"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	//DATABASE
	_, err := database.ConnectDB()
	if err != nil {
		return
	}

	// Gin instance
	r := gin.Default()

	creditLine := services.NewCreditLine()

	// Routes
	r.GET("/creditLines/", creditLine.GetCreditLines)
	r.GET("/creditLines/:id", creditLine.GetCreditLine)
	r.POST("/creditLines/", creditLine.CreateCreditLine)
	r.PUT("/creditLines/:id", creditLine.UpdateCreditLine)
	r.DELETE("/creditLines/:id", creditLine.DeleteCreditLine)

	// Swagger documentation
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
