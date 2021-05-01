package gateway

import "gorm.io/gorm"

// ManageTransaction ...
type ManageTransaction interface {
	Begin() *gorm.DB
	Rollback(tx *gorm.DB) *gorm.DB
	Commit(tx *gorm.DB) *gorm.DB
}
