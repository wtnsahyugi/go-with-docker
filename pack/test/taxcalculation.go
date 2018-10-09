package test

import (
	"strconv"

	"./repo"
	"github.com/gin-gonic/gin"
)

func TaxCalculationHandler(c *gin.Context) {
	prodId, _ := strconv.ParseInt(c.Param("id"), 10, 32)

	items, err := repo.TaxCalculate(int(prodId))

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, items)
}
