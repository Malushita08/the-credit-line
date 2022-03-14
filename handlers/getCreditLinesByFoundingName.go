package handlers

import (
	"github.com/Malushita08/the-credit-line/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCreditLinesByFoundingName godoc
// @Summary Get all the creditLines requests a foundingName did
// @Description Get all the creditLines requests a foundingName did
// @Tags creditLine
// @Param foundingName path string true "creditLine foundingName"
// @Success 200 {array} models.CreditLine
// @Router /creditLines/foundingName/{foundingName} [get]
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
