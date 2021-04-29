package model

import "time"

// User ...
type User struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	LastName  string // 苗字
	FirstName string // 氏名
	UserName  string // ユーザー名
	Password  string // パスワード
	Email     string // メールアドレス
}
