package test

import (
	"../../core"
	"./repo"
	"github.com/gin-gonic/gin"
)

func ProductHandler(c *gin.Context) {
	items, err := repo.GetAllProduct()

	if err != nil {
		c.JSON(200, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, items)
}

func CreateProductHandler(c *gin.Context) {

	items, err := repo.CreateProduct(c)

	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}
	if items != nil {

	}
	core.SendResponse(c, 200, "succesfully insert data", items)
}
