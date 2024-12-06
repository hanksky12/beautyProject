package model

import (
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt int64
	UpdatedAt int64
}

func (model *Base) BeforeCreate(db *gorm.DB) (err error) {
	uTime := time.Now().Unix()
	model.CreatedAt = uTime
	model.UpdatedAt = uTime
	return
}

func (model *Base) BeforeUpdate(db *gorm.DB) (err error) {
	uTime := time.Now().Unix()
	model.UpdatedAt = uTime
	return
}
