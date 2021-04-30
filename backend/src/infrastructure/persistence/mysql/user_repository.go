package mysql

import (
	"context"
	"react-echo-sample/adapter/gateway"
	"react-echo-sample/conf"
	"react-echo-sample/domain/model"
	"react-echo-sample/infrastructure/rdb"
	"react-echo-sample/infrastructure/transaction"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// userRepository userRepository構造体
type userRepository struct {
	DBConn *gorm.DB // トランザクションを使用しない場合
}

// NewUserRepository NewUserRepository関数
// 役割：userRepositoryのコンストラクタ関数
func NewUserRepository(DBConn *gorm.DB) gateway.UserRepositoryAccess {
	return &userRepository{DBConn}
}

// FetchByID FetchByIDメソッド
// 役割：指定されたIDに対応する単一レコードの取得
func (r *userRepository) FetchByID(ctx context.Context, id uint, demandPW bool) (*model.User, error) {
	return r.fetchByID(id, demandPW)
}

// fetchByID FetchByIDの実体
// 役割：
func (r *userRepository) fetchByID(id uint, demandPW bool) (*model.User, error) {
	db := r.DBConn

	user := &rdb.User{}
	if !demandPW {
		// PWが必要な場合以外、PWの抽出は省略
		db = db.Omit("password")
	}
	if err := db.Debug().
		Where("id = ?", id).First(user).
		Error; err != nil {
		db.Rollback()
		err = conf.NewAppError(conf.ErrFailedToServer, errors.Wrap(err, "failed to fetch user"))
		return nil, err
	}

	return convertRdbUserModelToDomain(user), nil
}

// FetchByLoginInfo FetchByLoginInfoメソッド
// 役割：指定されたIDに対応する単一レコードの取得
func (r *userRepository) FetchByLoginInfo(ctx context.Context, loginInfo *model.User) (*model.User, error) {
	db := r.DBConn

	user := &rdb.User{}
	if err := db.Debug().Where("email = ? AND password = ?", loginInfo.Email, loginInfo.Password).First(user).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		db.Rollback()
		err = conf.NewAppError(conf.ErrFailedToServer, errors.Wrap(err, "failed to fetch user"))
		return nil, err
	}

	return convertRdbUserModelToDomain(user), nil
}

func (r *userRepository) Search(ctx context.Context) ([]*model.User, error) {
	db := r.DBConn

	users := []*rdb.User{}

	// PWの抽出は省略
	db = db.Omit("password")

	if err := db.Debug().Find(&users).
		Error; err != nil {
		db.Rollback()
		err = conf.NewAppError(conf.ErrFailedToServer, errors.Wrap(err, "failed to search users"))
		return nil, errors.WithStack(err)
	}

	return convertRdbUserModelsToDomains(users), nil
}

func (r *userRepository) TxCreate(ctx context.Context, createUserInput *model.User) (uint, error) {
	tx, _ := transaction.WithContext(ctx)

	user := convertCreateUserInputToRdb(createUserInput)
	if err := tx.Debug().Create(user).Error; err != nil {
		tx.Rollback()
		err = conf.NewAppError(conf.ErrFailedToServer, errors.Wrap(err, "failed to create user"))
		return 0, errors.WithStack(err)
	}

	return user.ID, nil
}

func (r *userRepository) TxUpdate(ctx context.Context, updateUserInput *model.User) error {
	tx, _ := transaction.WithContext(ctx)

	result := tx.Debug().Where("updated_at = ?", updateUserInput.UpdatedAt).Updates(convertUpdateUserInputToRdb(updateUserInput))

	err := result.Error
	if err != nil {
		tx.Rollback()
		err = conf.NewAppError(conf.ErrFailedToServer, errors.Wrap(err, "failed to update user"))
		return errors.WithStack(err)
	}
	// 対象レコードがなくUpdate処理が実行されなかった場合のエラーハンドリング
	if result.RowsAffected == 0 {
		tx.Rollback()
		err = conf.NewAppError(conf.ErrFailedToServer, errors.New("failed to update user"))
		return errors.WithStack(err)
	}
	return nil
}

func (r *userRepository) TxDelete(ctx context.Context, deleteUserInput *model.User) error {
	tx, _ := transaction.WithContext(ctx)

	result := tx.Debug().Where("id = ? AND updated_at = ?", deleteUserInput.ID, deleteUserInput.UpdatedAt).Delete(&rdb.User{})

	err := result.Error
	if err != nil {
		tx.Rollback()
		err := conf.NewAppError(conf.ErrFailedToServer, errors.Wrap(err, "failed to delete user"))
		return errors.WithStack(err)
	}
	// 対象のレコードが見つからなかった場合のエラーハンドリング
	if result.RowsAffected == 0 {
		tx.Rollback()
		err := conf.NewAppError(conf.ErrFailedToServer, errors.New("failed to delete user"))
		return errors.WithStack(err)
	}

	return nil
}
