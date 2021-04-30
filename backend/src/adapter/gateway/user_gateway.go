package gateway

import (
	"context"
	"react-echo-sample/domain/model"
)

// UserRepositoryAccess UserRepositoryAccessインターフェース
type UserRepositoryAccess interface {
	FetchByID(context.Context, uint, bool) (*model.User, error)
	FetchByLoginInfo(context.Context, *model.User) (*model.User, error)
	Search(context.Context) ([]*model.User, error)
	TxCreate(context.Context, *model.User) (uint, error)
	TxUpdate(context.Context, *model.User) error
	TxDelete(context.Context, *model.User) error
}
