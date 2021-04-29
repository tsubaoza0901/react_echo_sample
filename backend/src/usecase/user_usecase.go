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
	GetUser(context.Context, uint) (*model.User, error)
	// GetUsers(context.Context, *iodata.UserSearchInputData) ([]*iodata.UserSearchOutputData, error)
	UpdateUser(context.Context, *model.User) error
	// DeleteUser(context.Context, *iodata.UserDeleteInputData) error
}

// userUseCase userUseCase構造体
// 役割：埋め込んだinterfaceに定義されたメソッドを自身の構造体のメソッドとして取得
type userUseCase struct {
	gateway.UserRepositoryAccess
	transaction.ManageTransaction
}

// NewUserUseCase NewUserUseCase関数
// 役割：userUseCaseのコンストラクタ関数
func NewUserUseCase(userra gateway.UserRepositoryAccess, mtx transaction.ManageTransaction) UserUseCase {
	return &userUseCase{userra, mtx}
}

// GetUser GetUserメソッド
// 役割：指定したIDのユーザー情報取得
func (u *userUseCase) GetUser(ctx context.Context, id uint) (*model.User, error) {
	user, err := u.UserRepositoryAccess.FetchByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "faild to FetchByID()")
	}
	return user, nil
}

// // GetUsers GetUsersメソッド
// // 役割：ユーザー情報取得
// func (u *userUseCase) GetUsers(ctx context.Context, inputdata *iodata.UserSearchInputData) ([]*iodata.UserSearchOutputData, error) {
// 	users, err := u.UserRepositoryAccess.Search(ctx, inputdata)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "faild to Search()")
// 	}
// 	return convertUserModelsToOutputData(users), nil
// }

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

// // Delete ...
// func (u *userUseCase) DeleteUser(ctx context.Context, inputdata *iodata.UserDeleteInputData) error {
// 	if err := u.UserRepositoryAccess.TxDelete(ctx, inputdata); err != nil {
// 		return errors.Wrap(err, "faild to TxDelete()")
// 	}
// 	return nil
// }
