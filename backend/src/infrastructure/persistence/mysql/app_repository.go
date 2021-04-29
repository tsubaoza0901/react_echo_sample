package mysql

import (
	"react-echo-sample/infrastructure/transaction"

	"gorm.io/gorm"
)

type appRepository struct {
	Conn *gorm.DB
}

// NewAppRepository NewAppRepository関数
func NewAppRepository(Conn *gorm.DB) transaction.ManageTransaction {
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
