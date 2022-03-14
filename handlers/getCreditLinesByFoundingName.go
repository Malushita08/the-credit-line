package handlers

import (
	"github.com/Malushita08/the-credit-line/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCreditLinesByFoundingName godoc
// @Summary Create a creditLine
// @Description Create a creditLine
// @Tags creditLine
// @Accept json
// @Param creditLine body models.CreditLineRequestBody true "creditLine Data"
// @Success 200 {object} models.ResponseBody
// @Router /creditLines [post]
func GetCreditLinesByFoundingName(db database.CreditLineInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		foundingName := c.Param("foundingName")
		res, err := db.GetCreditLinesByFoundingName(foundingName)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}
