package repo

import (
	"../../../config"
	"../model"
	"github.com/gin-gonic/gin"
)

func GetAllProduct() ([]model.GetProduct, error) {
	var mod []model.GetProduct

	db, err := config.GetConn()

	if err != nil {
		return nil, err
	}
	res := db.Table("products").Select("products.name, products.price, taxes.tax_name, taxes.refundable").Joins("left join taxes on products.tax_id = taxes.id").Scan(&mod)

	return mod, res.Error
}

func CreateProduct(c *gin.Context) (*model.Product, error) {

	var mod model.Product

	if err := c.ShouldBindJSON(&mod); err != nil { //validator
		return nil, err
	}

	db, err := config.GetConn()

	if err != nil {
		return nil, err
	}

	res := db.Create(&mod)

	return &mod, res.Error
}
