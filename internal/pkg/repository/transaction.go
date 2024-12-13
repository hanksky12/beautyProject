package repository

import (
	"beautyProject/internal/pkg/util/db/sql"
	"gorm.io/gorm"
)

func RunTransaction(fn func(tx *gorm.DB) error) error {
	return sql.Db.Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}
