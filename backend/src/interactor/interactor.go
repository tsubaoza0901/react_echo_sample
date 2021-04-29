package interactor

import (
	"react-echo-sample/adapter/gateway"
	"react-echo-sample/adapter/http/handler"
	"react-echo-sample/infrastructure/persistence/mysql"
	"react-echo-sample/infrastructure/transaction"
	"react-echo-sample/usecase"

	"gorm.io/gorm"
)

// Interactor Interactor
// 役割：すべての依存関係を把握する設定ファイル（安易コンテナ）
type Interactor interface {
	NewAppHandler() handler.AppHandler
}

type interactor struct {
	Conn *gorm.DB
}

// NewInteractor NewInteractor関数
// 役割：interactorのコンストラクタ関数
func NewInteractor(Conn *gorm.DB) Interactor {
	return &interactor{Conn}
}

type appHandler struct {
	handler.UserHandler
	handler.AuthHandler
}

// NewAppHandler ...
func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{
		i.NewUserHandler(),
		i.NewAuthHandler(),
	}
	return appHandler
}

// NewAppRepository NewAppRepositoryメソッド
// 役割：インスタンス生成
func (i *interactor) NewAppRepository() transaction.ManageTransaction {
	return mysql.NewAppRepository(i.Conn)
}

// < Auth > ---------------------------------------------

// NewAuthHandler NewUserHandlerメソッド
// 役割：インスタンス生成
func (i *interactor) NewAuthHandler() handler.AuthHandler {
	return handler.NewAuthHandler(i.NewAuthUseCase())
}

// NewAuthUseCase NewAuthUseCaseメソッド
// 役割：インスタンス生成
func (i *interactor) NewAuthUseCase() usecase.AuthUseCase {
	return usecase.NewAuthUseCase(
		i.NewUserRepository(),
		i.NewAppRepository(),
	)
}

// < User > ---------------------------------------------

// NewUserHandler NewUserHandlerメソッド
// 役割：インスタンス生成
func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserUseCase())
}

// NewUserUseCase NewUserUseCaseメソッド
// 役割：インスタンス生成
func (i *interactor) NewUserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(
		i.NewUserRepository(),
		i.NewAppRepository(),
		// i.NewUserService(),
	)
}

// NewUserRepository NewUserRepositoryメソッド
// 役割：インスタンス生成
func (i *interactor) NewUserRepository() gateway.UserRepositoryAccess {
	return mysql.NewUserRepository(
		i.Conn,
		// i.NewAppRepository(),
	)
}
