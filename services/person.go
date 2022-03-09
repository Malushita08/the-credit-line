package services

import (
	"github.com/Malushita08/the-credit-line/database"
	"github.com/Malushita08/the-credit-line/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type Data struct {
	DbSession *gorm.DB
}

func New() *Data {
	db, _ := database.ConnectDB()
	db.AutoMigrate(&models.Person{})
	db.AutoMigrate(&models.CreditLine{})
	return &Data{DbSession: db}
}

//get persons
// @Summary Get all people
// @Description get all people
// @Tags people
// @Router /people [get]
func (repository *Data) GetPeople(c *gin.Context) {
	var person []models.Person
	err := models.GetPeople(repository.DbSession, &person)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, person)
}
