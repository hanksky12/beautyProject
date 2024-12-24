package repository

import (
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/util/db/sql"
	"beautyProject/internal/pkg/util/web/requestConversion"
	"gorm.io/gorm"
)

type StatusRecordAverage struct {
}

func (s *StatusRecordAverage) Add(tx *gorm.DB, record *model.StatusRecordAverage) error {
	return tx.Create(record).Error
}

func (s *StatusRecordAverage) FindByUserId(userId uint64, cond map[string]any, paging *requestConversion.PagingSchema) ([]map[string]any, int, error) {
	query := sql.Db.Model(&model.StatusRecordAverage{})
	query = query.Where("user_id = ?", userId)
	return DynamicQuery(cond, query, paging)
}
