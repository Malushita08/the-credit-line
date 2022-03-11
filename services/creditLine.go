package services

import (
	"errors"
	"github.com/Malushita08/the-credit-line/database"
	"github.com/Malushita08/the-credit-line/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
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
// @Param creditLine body models.CreditLineRequestBody true "creditLine Data"
// @Success 200 {object} models.CreditLine
// @Router /creditLines [post]
//
func (repository *CreditLineData) CreateCreditLine(c *gin.Context) {
	var creditLineRequestBody models.CreditLineRequestBody
	var creditLine models.CreditLine
	var lastCreditLine models.CreditLine
	var responseBody models.ResponseBody

	_ = c.BindJSON(&creditLineRequestBody)

	//Defining our creditLine in base of the requestBody
	timeStp, _ := time.Parse(time.RFC3339, creditLineRequestBody.RequestedDate)
	creditLine = models.CreditLine{
		FoundingType:        creditLineRequestBody.FoundingType,
		CashBalance:         creditLineRequestBody.CashBalance,
		MonthlyRevenue:      creditLineRequestBody.MonthlyRevenue,
		RequestedCreditLine: creditLineRequestBody.RequestedCreditLine,
		RequestedDate:       timeStp,
	}

	err := models.CreateCreditLine(repository.DbSession, &creditLine, &lastCreditLine)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Controlling the Date param in our requests
	if creditLine.State == "ACCEPTED" {
		if creditLine.AllowedRequest == true {
			//Defining our responseBody
			responseBody = models.ResponseBody{
				Data:    &creditLine,
				Message: "ACCEPTED",
				Error:   nil,
			}
			c.JSON(http.StatusOK, responseBody)
		} else {
			responseBody = models.ResponseBody{
				Data:    &creditLine,
				Message: "ACCEPTED",
				Error:   nil,
			}
			c.JSON(426, responseBody)
		}
	} else {
		if creditLine.AllowedRequest == true {
			responseBody = models.ResponseBody{
				Data:    nil,
				Message: "REJECTED CREDIT LINE REQUEST",
				Error:   nil,
			}
			c.JSON(http.StatusOK, responseBody)
		} else {
			responseBody = models.ResponseBody{
				Data:    &creditLine,
				Message: "Wait 30 seconds please",
				Error:   nil,
			}
			//c.AbortWithStatusJSON(426, gin.H{"error": "time?"})
			c.JSON(426, responseBody)
		}
	}
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
