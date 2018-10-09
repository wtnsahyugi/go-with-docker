package db

import (
	"../../../config"
	"../model"
	"github.com/jinzhu/gorm"
)

func MigrateProduct() error {

	db, err := config.GetConn()

	if err == nil {
		db.AutoMigrate(&model.Product{})

		//seeding jika init masukin datanya langsung agar langsung terisi !
		seedProduct(db)
		return nil
	}

	return err
}

func seedProduct(db *gorm.DB) {
	db.Create(&model.Product{Name: "Lucky Stretch", Price: 1000, TaxID: 2})
	db.Create(&model.Product{Name: "Big Mac", Price: 1000, TaxID: 1})
	db.Create(&model.Product{Name: "Movie", Price: 150, TaxID: 3})
}
