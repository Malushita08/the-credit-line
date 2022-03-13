package services

import (
	"errors"
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

// GetCreditLines godoc
// @Summary Get all creditLines
// @Description get all creditLines
// @Tags creditLine
// @Success 200 {array} models.CreditLine
// @Router /creditLines [get]
func (repository *CreditLineData) GetCreditLines(c *gin.Context) {
	var creditLine []models.CreditLine
	err := models.GetCreditLines(repository.DbSession, &creditLine)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, creditLine)
}

// GetCreditLine godoc
// @Summary Get a creditLine by id
// @Description Get a creditLine by id
// @Tags creditLine
// @Param id path string true "creditLine ID"
// @Success 200 {object} models.CreditLine
// @Router /creditLines/{id} [get]
func (repository *CreditLineData) GetCreditLine(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var creditLine models.CreditLine

	err := models.GetCreditLine(repository.DbSession, &creditLine, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, creditLine)
}

func (repository *CreditLineData) GetCreditLineByFoundingName(c *gin.Context) {
	foundingName, _ := c.Params.Get("foundingName")
	var creditLine []models.CreditLine

	err := models.GetCreditLinesByFoundingName(repository.DbSession, &creditLine, foundingName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, creditLine)
}
