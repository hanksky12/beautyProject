package repository

import (
	"beautyProject/internal/pkg/model"
	"beautyProject/internal/pkg/util/db/sql"
)

type MouseAction struct {
}

func (h *MouseAction) FindAll() ([]model.MouseAction, error) {
	var action []model.MouseAction
	err := sql.Db.Find(&action).Error
	return action, err
}
