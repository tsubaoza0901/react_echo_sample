package usecase

import (
	"context"
	"errors"
	"react-echo-sample/adapter/gateway"
	"react-echo-sample/domain/model"
	"react-echo-sample/infrastructure/transaction"
)

// AuthUseCase AuthUseCase interface
type AuthUseCase interface {
	Signup(context.Context, *model.User) (uint, error)
	Login(context.Context, *model.User) (uint, error)
}

// authUseCase authUseCase構造体
// 役割：埋め込んだinterfaceに定義されたメソッドを自身の構造体のメソッドとして取得
type authUseCase struct {
	gateway.UserRepositoryAccess
	gateway.ManageTransaction
}

// NewAuthUseCase NewAuthUseCase関数
// 役割：authUseCaseのコンストラクタ関数
func NewAuthUseCase(userra gateway.UserRepositoryAccess, mtx gateway.ManageTransaction) AuthUseCase {
	return &authUseCase{userra, mtx}
}

func (u *authUseCase) Signup(ctx context.Context, input *model.User) (uint, error) {
	tx := u.ManageTransaction.Begin()
	ctx = transaction.NewContext(ctx, tx)

	// 重複チェック（TODO:ここでやるべきか検討。service?）
	user, err := u.FetchByLoginInfo(ctx, input)
	if err != nil {
		return 0, err
	}
	if user != nil {
		// TODO:エラーコードを返却するよう修正
		return 0, errors.New("すでに登録済みのユーザー")
	}

	// ユーザー登録
	userID, err := u.TxCreate(ctx, input)
	if err != nil {
		return 0, err
	}

	tx.Commit()

	return userID, nil
}

func (u *authUseCase) Login(ctx context.Context, input *model.User) (uint, error) {
	tx := u.ManageTransaction.Begin()
	ctx = transaction.NewContext(ctx, tx)

	// 重複チェック（TODO:ここでやるべきか検討。service?）
	user, err := u.FetchByLoginInfo(ctx, input)
	if err != nil {
		return 0, err
	}
	if user == nil {
		// TODO:エラーコードを返却するよう修正
		return 0, errors.New("ユーザー登録なし")
	}

	tx.Commit()

	return user.ID, nil
}
