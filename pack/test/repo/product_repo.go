package repo

import (
	"../../../config"
	"../model"
	"github.com/gin-gonic/gin"
)

func GetAllProduct() ([]model.Product, error) {
	var mod []model.Product

	db, err := config.GetConn()

	if err != nil {
		return nil, err
	}
	res := db.Find(&mod)

	return mod, res.Error
}

func CreateProduct(c *gin.Context) (*model.Product, error) {
	var mod model.Product

	if err_ := c.ShouldBindJSON(&mod); err_ != nil { //validator
		return nil, err_
	}

	db, err := config.GetConn()

	if err != nil {
		return nil, err
	}

	res := db.Create(&mod)

	return &mod, res.Error
}