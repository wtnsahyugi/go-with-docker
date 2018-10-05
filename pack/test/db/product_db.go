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

		//seeding jika inin masukin datanya langsung biar crot !
		seedProduct(db)
		return nil
	}
	return err
}

func seedProduct(db *gorm.DB) {
	// db.Create(&model.Product{Name: "Lucky Stretch", Price: 1000, TaxCode: 2})
	// db.Create(&model.Product{Name: "Big Mac", Price: 1000, TaxCode: 1})
	// db.Create(&model.Product{Name: "Movie", Price: 150, TaxCode: 3})
}