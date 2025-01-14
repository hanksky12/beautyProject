package repository

import (
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/util/db/sql"
)

type Hardware struct {
}

func (h *Hardware) FindAll() ([]model.Hardware, error) {
	var hardware []model.Hardware
	err := sql.Db.Find(&hardware).Error
	return hardware, err
}
