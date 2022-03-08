package main

// @title Gin Swagger Example API
// @version 1.0
// @description This is a sample server.

// @host localhost:8080
// @BasePath /
// @schemes http
import (
	"log"
	"net/http"

	"github.com/Malushita08/the-credit-line/database"

	_ "github.com/Malushita08/the-credit-line/docs"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	//DATABASE
	database.ConnectDB()

	// Gin instance
	r := gin.Default()

	// Routes
	//r.GET("/people/", GetPeople)
	//r.GET("/people/:id", GetPerson)
	//r.POST("/people", CreatePerson)
	//r.PUT("/people/:id", UpdatePerson)
	//r.DELETE("/people/:id", DeletePerson)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

//CONTROLLERS

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}
	c.JSON(http.StatusOK, res)
}

// CreateUser ... Create User
// @Summary Create new user based on paramters
// @Description Create new user
// @Tags Users
// @Accept json
// @Param user body Person true "User Data"
// @Success 200 {object} object
// @Failure 400,500 {object} object
// @Router / [post]
//func DeletePerson(c *gin.Context) {
//	id := c.Params.ByName("id")
//	var person Person
//	d := db.Where("id = ?", id).Delete(&person)
//	fmt.Println(d)
//	c.JSON(200, gin.H{"id #" + id: "deleted"})
//}
//
//func UpdatePerson(c *gin.Context) {
//
//	var person Person
//	id := c.Params.ByName("id")
//
//	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
//		c.AbortWithStatus(404)
//		fmt.Println(err)
//	}
//	err := c.BindJSON(&person)
//	if err != nil {
//		return
//	}
//
//	db.Save(&person)
//	c.JSON(200, person)
//
//}
//
//func CreatePerson(c *gin.Context) {
//
//	var person Person
//	err := c.BindJSON(&person)
//	if err != nil {
//		return
//	}
//
//	db.Create(&person)
//	c.JSON(200, person)
//}
//
//func GetPerson(c *gin.Context) {
//	id := c.Params.ByName("id")
//	var person Person
//	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
//		c.AbortWithStatus(404)
//		fmt.Println(err)
//	} else {
//		c.JSON(200, person)
//	}
//}
//
//// @Summary Get all people
//// @Description get all people
//// @Tags people
//// @Success 200 {array} Person
//// @Failure 404 {object} object
//// @Router /people [get]
//func GetPeople(c *gin.Context) {
//	var people []Person
//	if err := db.Find(&people).Error; err != nil {
//		c.AbortWithStatus(404)
//		fmt.Println(err)
//	} else {
//		c.JSON(200, people)
//	}

//}
