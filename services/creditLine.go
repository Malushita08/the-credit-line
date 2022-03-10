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

// CreateCreditLine godoc
// @Summary Create a creditLine
// @Description Create a creditLine
// @Tags creditLine
// @Accept json
// @Param creditLine body models.CreditLine true "creditLine Data"
// @Success 200 {object} models.CreditLine
// @Router /creditLines [post]
func (repository *CreditLineData) CreateCreditLine(c *gin.Context) {
	var creditLine models.CreditLine
	err := c.BindJSON(&creditLine)
	if err != nil {
		return
	}
	err = models.CreateCreditLine(repository.DbSession, &creditLine)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, creditLine)
}

// UpdateCreditLine godoc
func (repository *CreditLineData) UpdateCreditLine(c *gin.Context) {
	var creditLine models.CreditLine
	id, _ := c.Params.Get("id")
	err := models.GetCreditLine(repository.DbSession, &creditLine, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&creditLine)
	err = models.UpdateCreditLine(repository.DbSession, &creditLine)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, creditLine)
}

// delete book
func (repository *CreditLineData) DeleteCreditLine(c *gin.Context) {
	var book models.CreditLine
	id, _ := c.Params.Get("id")
	err := models.DeleteCreditLine(repository.DbSession, &book, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "CreditLine deleted successfully"})
}
