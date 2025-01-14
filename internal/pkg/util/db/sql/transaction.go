package sql

import (
	"gorm.io/gorm"
)

func RunTransaction(fn func(tx *gorm.DB) error) error {
	return Db.Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}
