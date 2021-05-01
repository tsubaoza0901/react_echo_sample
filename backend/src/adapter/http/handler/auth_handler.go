package handler

import (
	"net/http"
	"os"
	"react-echo-sample/adapter/http/interfaces/request"
	"react-echo-sample/adapter/http/interfaces/response"
	"react-echo-sample/conf"
	"react-echo-sample/domain/model"
	"react-echo-sample/usecase"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// AuthHandler AuthHandler interface
type AuthHandler interface {
	Signup(echo.Context) error
	Login(echo.Context) error
}

// authHandler authHandler構造体
// 役割：埋め込んだinterfaceに定義されたメソッドを自身の構造体のメソッドとして取得
type authHandler struct {
	usecase.AuthUseCase
}

// NewAuthHandler NewUserHandler関数
// 役割：authHandlerのコンストラクタ関数
func NewAuthHandler(u usecase.AuthUseCase) AuthHandler {
	return &authHandler{u}
}

// Signup ...
// 役割：userの登録
// @Resource /v1/signup
// @Router api/v1/signup [post]
func (h *authHandler) Signup(c echo.Context) error {
	ctx := c.Request().Context()

	// HTTP request bodyのbind処理
	req := &request.SignupInfo{}
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

	input := &model.User{
		LastName:  req.LastName,
		FirstName: req.FirstName,
		UserName:  req.UserName,
		Password:  req.Password,
		Email:     req.Email,
	}

	userID, err := h.AuthUseCase.Signup(ctx, input)
	if err != nil {
		// zap.S().Errorw("create error", zap.Error(err))
		code := conf.ErrFailedToServer
		if apperr, ok := errors.Cause(err).(*conf.AppError); ok {
			code = apperr.Code
			err = apperr.Wrap()
		} else {
			err = conf.NewAppError(conf.ErrFailedToServer, err).Wrap()
		}
		return c.JSON(http.StatusOK, response.NewAPIResponse(code, err.Error(), nil))
	}

	// トークン生成
	token := makeToken(userID)

	return c.JSON(http.StatusOK, response.NewAPIResponse(0, response.StatusText(response.StatusSuccess), response.Jwt{Token: token}))
}

// Login ...
// 役割：ログイン
// @Resource /v1/login
// @Router api/v1/login [post]
func (h *authHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()

	// HTTP request bodyのbind処理
	req := &request.LoginInfo{}
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

	input := &model.User{
		Password: req.Password,
		Email:    req.Email,
	}

	userID, err := h.AuthUseCase.Login(ctx, input)
	if err != nil {
		// zap.S().Errorw("create error", zap.Error(err))
		code := conf.ErrFailedToServer
		if apperr, ok := errors.Cause(err).(*conf.AppError); ok {
			code = apperr.Code
			err = apperr.Wrap()
		} else {
			err = conf.NewAppError(conf.ErrFailedToServer, err).Wrap()
		}
		return c.JSON(http.StatusOK, response.NewAPIResponse(code, err.Error(), nil))
	}

	// トークン生成
	token := makeToken(userID)

	return c.JSON(http.StatusOK, response.NewAPIResponse(0, response.StatusText(response.StatusSuccess), response.Jwt{Token: token}))
}

// makeToken トークン生成（TODO:実装場所検討）
func makeToken(userID uint) string {
	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)

	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = userID
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// 電子署名
	tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))

	return tokenString
}
