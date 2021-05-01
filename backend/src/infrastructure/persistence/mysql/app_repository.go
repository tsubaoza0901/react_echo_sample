package mysql

import (
	"react-echo-sample/adapter/gateway"

	"gorm.io/gorm"
)

type appRepository struct {
	Conn *gorm.DB
}

// NewAppRepository NewAppRepository関数
func NewAppRepository(Conn *gorm.DB) gateway.ManageTransaction {
	return &appRepository{Conn}
}

func (r *appRepository) Begin() *gorm.DB {
	return r.Conn.Begin()
}

func (r *appRepository) Rollback(tx *gorm.DB) *gorm.DB {
	return tx.Rollback()
}

func (r *appRepository) Commit(tx *gorm.DB) *gorm.DB {
	return tx.Commit()
}
