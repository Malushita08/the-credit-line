package handlers

import (
	"github.com/Malushita08/the-credit-line/database"
	"github.com/Malushita08/the-credit-line/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InsertCreditLine(db database.CreditLineInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		creditLineRequestBody := models.CreditLineRequestBody{}
		err := c.BindJSON(&creditLineRequestBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		res, err := db.CreateCreditLine(creditLineRequestBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}
