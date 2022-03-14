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
	"github.com/Malushita08/the-credit-line/handlers"
	"github.com/gin-gonic/gin"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	//DATABASE
	db, err := database.ConnectDB()
	if err != nil {
		return
	}

	//CreditLine instance
	client := &database.CreditLineClient{
		DbSession: db,
	}

	//Gin instance
	r := gin.Default()

	// Routes
	creditLine := r.Group("/creditLines")
	creditLine.POST("/", handlers.CreateCreditLine(client))
	creditLine.GET("/foundingName/:foundingName", handlers.GetCreditLinesByFoundingName(client))

	// Swagger documentation
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	//Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
