package model

import "github.com/jinzhu/gorm"

type Tax struct {
	gorm.Model
	TaxName    string `form:"TaxName" json:"TaxName" binding:"required,max=20,min=1"`
	Refundable bool   `form:"Refundable" json:"Refundable" binding:"required"`
}
