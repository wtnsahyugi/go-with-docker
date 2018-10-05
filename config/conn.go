
package config

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/postgres"

func GetConn() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 sslmode=disable user=postgres dbname=test password=mypass")

	return db, err
}