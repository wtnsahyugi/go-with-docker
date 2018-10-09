package config

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/postgres"

func GetConn() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=172.15.0.1 port=5432 sslmode=disable user=go dbname=go password=go")

	return db, err
}
