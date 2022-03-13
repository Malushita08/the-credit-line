package handlers

import (
	"github.com/Malushita08/the-credit-line/database"
	"github.com/Malushita08/the-credit-line/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InsertCreditLine godoc
// @Summary Create a creditLine
// @Description Create a creditLine
// @Tags creditLine
// @Accept json
// @Param creditLine body models.CreditLineRequestBody true "creditLine Data"
// @Success 200 {object} models.ResponseBody
// @Router /creditLines [post]
func InsertCreditLine(db database.CreditLineInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		creditLineRequestBody := models.CreditLineRequestBody{}
		err := c.BindJSON(&creditLineRequestBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		responseBody, err := db.CreateCreditLine(creditLineRequestBody)
		if err != nil {
			if responseBody.Message == "Please, wait 30 seconds" || responseBody.Message == "You've done more than 3 request within 2 minutes" {
				c.AbortWithStatusJSON(426, responseBody)
				return
			}
			c.AbortWithStatusJSON(200, responseBody)
			return
		}
		c.JSON(http.StatusOK, responseBody)
	}
}
