package handler

import (
	"net/http"
	"react-echo-sample/adapter/http/interfaces/request"
	"react-echo-sample/adapter/http/interfaces/response"
	"react-echo-sample/conf"
	"react-echo-sample/domain/model"
	"react-echo-sample/usecase"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// UserHandler UserHandler interface
// 役割：user_handler.goに定義されたメソッドの呼び出しリスト
type UserHandler interface {
	GetUser(echo.Context) error
	GetUsers(echo.Context) error
	UpdateUser(echo.Context) error
	DeleteUser(echo.Context) error
}

// userHandler userHandler構造体
// 役割：埋め込んだinterfaceに定義されたメソッドを自身の構造体のメソッドとして取得
type userHandler struct {
	usecase.UserUseCase
}

// NewUserHandler NewUserHandler関数
// 役割：userHandlerのコンストラクタ関数
func NewUserHandler(u usecase.UserUseCase) UserHandler {
	return &userHandler{u}
}

// GetUser GetUserメソッド
// 役割：userの取得
// @Resource /v1/user
// @Router /api/v1/user/{id} [get]
func (h *userHandler) GetUser(c echo.Context) error {
	ctx := c.Request().Context()

	// HTTP request bodyのbind処理
	req := &request.SearchUser{}
	if err := c.Bind(req); err != nil {
		// zap.S().Errorw("failed to bind", zap.Error(err))
		err = conf.NewAppError(conf.ErrBadRequest, err).Wrap()
		return c.JSON(http.StatusOK, response.NewAPIResponse(conf.ErrBadRequest, err.Error(), nil))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		err = conf.NewAppError(conf.ErrBadRequest, err).Wrap()
		return c.JSON(http.StatusOK, response.NewAPIResponse(conf.ErrBadRequest, err.Error(), nil))
	}

	// // メモ:今のままだと、ぬるぽエラーになるためコメントアウト。おそらく、validator.goの実装内容のため、バリデーションが必要な際に再編集。
	// // validate req
	// if err := c.Validate(req); err != nil {
	// 	conf.WithContext(ctx).Debug("validate error", zap.Error(err))
	// 	err = model.NewAppError(model.ErrBadRequest, err).Wrap()
	// 	return c.JSON(http.StatusOK, model.NewAPIResponse(model.ErrBadRequest, err.Error(), nil))
	// }

	// // TODO：状況に応じて変更予定
	// input := &model.User{
	// 	ID: req.ID,
	// }

	user, err := h.UserUseCase.GetUser(ctx, uint(id), req.DemandPW)
	if err != nil {
		// zap.S().Errorw("get error", zap.Error(err))
		code := conf.ErrFailedToServer
		if apperr, ok := errors.Cause(err).(*conf.AppError); ok {
			code = apperr.Code
			err = apperr.Wrap()
		} else {
			err = conf.NewAppError(conf.ErrFailedToServer, err).Wrap()
		}
		return c.JSON(http.StatusOK, response.NewAPIResponse(code, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, response.NewAPIResponse(0, response.StatusText(response.StatusSuccess), response.ToUserResponse(user)))
}

// // GetUsers GetUsersメソッド
// 役割：userの全取得
// @Resource /v1/users
// @Router /api/v1/users [get]
func (h *userHandler) GetUsers(c echo.Context) error {
	ctx := c.Request().Context()

	// // HTTP request bodyのbind処理
	// req := &request.SearchUser{}
	// if err := c.Bind(req); err != nil {
	// 	// zap.S().Errorw("failed to bind", zap.Error(err))
	// 	err = conf.NewAppError(conf.ErrBadRequest, err).Wrap()
	// 	return c.JSON(http.StatusOK, response.NewAPIResponse(conf.ErrBadRequest, err.Error(), nil))
	// }

	// // メモ:今のままだと、ぬるぽエラーになるためコメントアウト。おそらく、validator.goの実装内容のため、バリデーションが必要な際に再編集。
	// // validate req
	// if err := c.Validate(req); err != nil {
	// 	conf.WithContext(ctx).Debug("validate error", zap.Error(err))
	// 	err = model.NewAppError(model.ErrBadRequest, err).Wrap()
	// 	return c.JSON(http.StatusOK, model.NewAPIResponse(model.ErrBadRequest, err.Error(), nil))
	// }

	// // TODO：状況に応じて変更予定
	// input := &model.User{
	// 	ID: req.ID,
	// }

	// 複数案件の取得
	users, err := h.UserUseCase.GetUsers(ctx)
	if err != nil {
		// zap.S().Errorw("gets error", zap.Error(err))
		code := conf.ErrFailedToServer
		if apperr, ok := errors.Cause(err).(*conf.AppError); ok {
			code = apperr.Code
			err = apperr.Wrap()
		} else {
			err = conf.NewAppError(conf.ErrFailedToServer, err).Wrap()
		}
		return c.JSON(http.StatusOK, response.NewAPIResponse(code, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, response.NewAPIResponse(0, response.StatusText(response.StatusSuccess), response.ToUserResponseList(users).Users))
}

// UpdateUser ...
// 役割：userの登録
// @Resource /v1/user
// @Router /auth/api/v1/user [put]
func (h *userHandler) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()

	// HTTP request bodyのbind処理
	req := &request.UpdateUser{}
	if err := c.Bind(req); err != nil {
		// zap.S().Errorw("failed to bind", zap.Error(err))
		err = conf.NewAppError(conf.ErrBadRequest, err).Wrap()
		return c.JSON(http.StatusOK, response.NewAPIResponse(conf.ErrBadRequest, err.Error(), nil))
	}

	// // validate req
	// if err := c.Validate(req); err != nil {
	// 	conf.WithContext(ctx).Debug("validate error", zap.Error(err))
	// 	err = model.NewAppError(model.ErrBadRequest, err).Wrap()
	// 	return c.JSON(http.StatusOK, model.NewAPIResponse(model.ErrBadRequest, err.Error(), nil))
	// }

	// bind結果などをusecase層で利用できるよう整形
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		err = conf.NewAppError(conf.ErrBadRequest, err).Wrap()
		return c.JSON(http.StatusOK, response.NewAPIResponse(conf.ErrBadRequest, err.Error(), nil))
	}

	// TODO:バリデーションを設定したら削除
	if req.UpdatedAt == "" {
		// zap.S().Errorw("UpdatedAt is empty", zap.Error(err))
		err = conf.NewAppError(conf.ErrBadRequest, err).Wrap()
		return c.JSON(http.StatusOK, response.NewAPIResponse(conf.ErrBadRequest, err.Error(), nil))
	}

	updatedAt, err := time.Parse(time.RFC3339, req.UpdatedAt)
	if err != nil {
		return c.JSON(http.StatusOK, response.NewAPIResponse(conf.ErrBadRequest, err.Error(), nil))
	}

	inputdata := &model.User{
		ID:        uint(id),
		UpdatedAt: updatedAt,
		LastName:  req.LastName,
		FirstName: req.FirstName,
		UserName:  req.UserName,
		Email:     req.Email,
		Password:  req.Password,
	}

	if err := h.UserUseCase.UpdateUser(ctx, inputdata); err != nil {
		// zap.S().Errorw("update error", zap.Error(err))
		code := conf.ErrFailedToServer
		if apperr, ok := errors.Cause(err).(*conf.AppError); ok {
			code = apperr.Code
			err = apperr.Wrap()
		} else {
			err = conf.NewAppError(conf.ErrFailedToServer, err).Wrap()
		}
		return c.JSON(http.StatusOK, response.NewAPIResponse(code, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, response.NewAPIResponse(0, response.StatusText(response.StatusSuccess), nil))
}

// DeleteUser ...
// 役割：userの削除
// @Resource /v1/user
// @Router /auth/api/v1/user [delete]
func (h *userHandler) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()

	// HTTP request bodyのbind処理
	req := &request.DeleteUser{}
	if err := c.Bind(req); err != nil {
		// zap.S().Errorw("failed to bind", zap.Error(err))
		err = conf.NewAppError(conf.ErrBadRequest, err).Wrap()
		return c.JSON(http.StatusOK, response.NewAPIResponse(conf.ErrBadRequest, err.Error(), nil))
	}

	// // validate req
	// if err := c.Validate(req); err != nil {
	// 	conf.WithContext(ctx).Debug("validate error", zap.Error(err))
	// 	err = model.NewAppError(model.ErrBadRequest, err).Wrap()
	// 	return c.JSON(http.StatusOK, model.NewAPIResponse(model.ErrBadRequest, err.Error(), nil))
	// }

	// bind結果などをusecase層で利用できるよう整形
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		err = conf.NewAppError(conf.ErrBadRequest, err).Wrap()
		return c.JSON(http.StatusOK, response.NewAPIResponse(conf.ErrBadRequest, err.Error(), nil))
	}

	// TODO:バリデーションを設定したら削除
	if req.UpdatedAt == "" {
		// zap.S().Errorw("UpdatedAt is empty", zap.Error(err))
		err = conf.NewAppError(conf.ErrBadRequest, err).Wrap()
		return c.JSON(http.StatusOK, response.NewAPIResponse(conf.ErrBadRequest, err.Error(), nil))
	}

	updatedAt, err := time.Parse(time.RFC3339, req.UpdatedAt)
	if err != nil {
		return err
	}
	input := &model.User{
		ID:        uint(id),
		UpdatedAt: updatedAt,
	}

	if err := h.UserUseCase.DeleteUser(ctx, input); err != nil {
		// zap.S().Errorw("delete error", zap.Error(err))
		code := conf.ErrFailedToServer
		if apperr, ok := errors.Cause(err).(*conf.AppError); ok {
			code = apperr.Code
			err = apperr.Wrap()
		} else {
			err = conf.NewAppError(conf.ErrFailedToServer, err).Wrap()
		}
		return c.JSON(http.StatusOK, response.NewAPIResponse(code, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, response.NewAPIResponse(0, response.StatusText(response.StatusSuccess), nil))
}
