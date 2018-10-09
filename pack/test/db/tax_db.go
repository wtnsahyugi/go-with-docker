package db

import (
	"../../../config"
	"../model"
	"github.com/jinzhu/gorm"
)

func MigrateTax() error {

	db, err := config.GetConn()

	if err == nil {
		db.AutoMigrate(&model.Tax{})

		//seeding jika inin masukin datanya langsung biar crot !
		seedTax(db)
		return nil
	}
	return err
}

func seedTax(db *gorm.DB) {
	db.Create(&model.Tax{TaxName: "food", Refundable: true})
	db.Create(&model.Tax{TaxName: "tobacco", Refundable: false})
	db.Create(&model.Tax{TaxName: "entertainment", Refundable: false})
}
