package response

import (
	"react-echo-sample/domain/model"
	"time"
)

// User User構造体
type User struct {
	ID        uint      `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	LastName  string    `json:"last_name"`
	FirstName string    `json:"first_name"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
}

// UsersList Userのレスポンスリスト
type UsersList struct {
	Users []*User `json:"users"`
}

// ToUserResponse ToUserResponse関数
func ToUserResponse(output *model.User) *User {
	return &User{
		ID:        output.ID,
		UpdatedAt: output.UpdatedAt,
		LastName:  output.LastName,
		FirstName: output.FirstName,
		UserName:  output.UserName,
		Email:     output.Email,
		Password:  output.Password,
	}
}

// ToUserResponseList ToUserResponseList関数
func ToUserResponseList(output []*model.User) *UsersList {
	userl := &UsersList{
		Users: []*User{},
	}
	for _, user := range output {
		userl.Users = append(userl.Users, ToUserResponse(user))
	}
	return userl
}
