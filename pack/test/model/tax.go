package model

import "github.com/jinzhu/gorm"

type Tax struct {
	gorm.Model
	TaxCode  uint `form:"TaxCode" json:"TaxCode" binding:"required"`
	TaxName uint `form:"TaxName" json:"TaxName" binding:"required"`
}