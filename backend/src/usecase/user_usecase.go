package usecase

import (
	"context"
	"react-echo-sample/adapter/gateway"
	"react-echo-sample/domain/model"
	"react-echo-sample/infrastructure/transaction"

	"github.com/pkg/errors"
)

// UserUseCase UserUseCase interface
type UserUseCase interface {
	GetUser(context.Context, uint, bool) (*model.User, error)
	GetUsers(context.Context) ([]*model.User, error)
	UpdateUser(context.Context, *model.User) error
	DeleteUser(context.Context, *model.User) error
}

// userUseCase userUseCase構造体
// 役割：埋め込んだinterfaceに定義されたメソッドを自身の構造体のメソッドとして取得
type userUseCase struct {
	gateway.UserRepositoryAccess
	gateway.ManageTransaction
}

// NewUserUseCase NewUserUseCase関数
// 役割：userUseCaseのコンストラクタ関数
func NewUserUseCase(userra gateway.UserRepositoryAccess, mtx gateway.ManageTransaction) UserUseCase {
	return &userUseCase{userra, mtx}
}

// GetUser GetUserメソッド
// 役割：指定したIDのユーザー情報取得
func (u *userUseCase) GetUser(ctx context.Context, id uint, demandPW bool) (*model.User, error) {
	user, err := u.UserRepositoryAccess.FetchByID(ctx, id, demandPW)
	if err != nil {
		return nil, errors.Wrap(err, "faild to FetchByID()")
	}
	return user, nil
}

// GetUsers GetUsersメソッド
// 役割：ユーザー情報取得
func (u *userUseCase) GetUsers(ctx context.Context) ([]*model.User, error) {
	users, err := u.UserRepositoryAccess.Search(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "faild to Search()")
	}
	return users, nil
}

// UpdateUser ...
func (u *userUseCase) UpdateUser(ctx context.Context, input *model.User) error {
	tx := u.ManageTransaction.Begin()
	ctx = transaction.NewContext(ctx, tx)

	if err := u.TxUpdate(ctx, input); err != nil {
		return errors.Wrap(err, "faild to UpdateUser()")
	}

	tx.Commit()
	return nil
}

// Delete ...
func (u *userUseCase) DeleteUser(ctx context.Context, input *model.User) error {
	tx := u.ManageTransaction.Begin()
	ctx = transaction.NewContext(ctx, tx)

	if err := u.UserRepositoryAccess.TxDelete(ctx, input); err != nil {
		return errors.Wrap(err, "faild to TxDelete()")
	}

	tx.Commit()
	return nil
}
