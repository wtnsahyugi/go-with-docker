package migration

import (
	db "../pack/test/db"
)

func TaxMigration() {
	db.MigrateTax()
}
