package migration

import (
	db "../pack/test/db"
)

func ProductMigration() {
	db.MigrateProduct()
}