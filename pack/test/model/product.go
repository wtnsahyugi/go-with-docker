package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name  string `form:"Name" json:"Name" binding:"required,max=20,min=1"`
	Price int    `form:"Price" json:"Price" binding:"required"`
	TaxID int    `form:"TaxID" json:"TaxID" binding:"required"`
	Tax   Tax
}

type GetProduct struct {
	gorm.Model
	Name       string
	Price      float32
	TaxID      string
	TaxName    string
	Refundable bool
	Tax        float32
	Amount     float32
}
