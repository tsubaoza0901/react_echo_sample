package mysql

import (
	"react-echo-sample/domain/model"
	"react-echo-sample/infrastructure/rdb"

	"gorm.io/gorm"
)

// User
func convertCreateUserInputToRdb(input *model.User) *rdb.User {
	return &rdb.User{
		LastName:  input.LastName,
		FirstName: input.FirstName,
		UserName:  input.UserName,
		Password:  input.Password,
		Email:     input.Email,
	}
}

// User
func convertRdbUserModelToDomain(input *rdb.User) *model.User {
	return &model.User{
		ID:        input.ID,
		UpdatedAt: input.UpdatedAt,
		LastName:  input.LastName,
		FirstName: input.FirstName,
		UserName:  input.UserName,
		Password:  input.Password,
		Email:     input.Email,
	}
}

func convertUpdateUserInputToRdb(input *model.User) *rdb.User {
	return &rdb.User{
		Model: gorm.Model{
			ID:        input.ID,
			UpdatedAt: input.UpdatedAt,
		},
		LastName:  input.LastName,
		FirstName: input.FirstName,
		UserName:  input.UserName,
		Email:     input.Email,
		Password:  input.Password,
	}
}
