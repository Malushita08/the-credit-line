package services

import (
	"github.com/Malushita08/the-credit-line/database"
	"github.com/Malushita08/the-credit-line/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type CreditLineData struct {
	DbSession *gorm.DB
}

func NewCreditLine() *CreditLineData {
	db, _ := database.ConnectDB()
	db.AutoMigrate(&models.CreditLine{})
	return &CreditLineData{DbSession: db}
}

//create creditline
func (repository *CreditLineData) CreateCreditLine(c *gin.Context) {
	var book models.CreditLine
	err := c.BindJSON(&book)
	if err != nil {
		return
	}
	err = models.CreateCreditLine(repository.DbSession, &book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, book)
}

//get creditLines
// @Summary Get all people
// @Description get all people
// @Tags creditLine
// @Router /creditLine [get]
func (repository *CreditLineData) GetCreditLine(c *gin.Context) {
	var creditLine []models.CreditLine
	err := models.GetCreditLine(repository.DbSession, &creditLine)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, creditLine)
}
