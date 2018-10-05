package repo

import (
	"../../../config"
	"../model"
	"github.com/gin-gonic/gin"
)

func GetAllTax() ([]model.Tax, error) {
	var mod []model.Tax

	db, err := config.GetConn()

	if err != nil {
		return nil, err
	}
	res := db.Find(&mod)

	return mod, res.Error
}

func CreateTax(c *gin.Context) (*model.Tax, error) {
	var mod model.Tax

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