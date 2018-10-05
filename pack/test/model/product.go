package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name  string `form:"Name" json:"Name" binding:"required,max=20,min=1"`
	TaxCode uint `form:"TaxCode" json:"TaxCode" binding:"required"`
	Price uint  `form:"Price" json:"Price" binding:"required"`
}