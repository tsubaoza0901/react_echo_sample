package rdb

import "gorm.io/gorm"

// User ...
type User struct {
	gorm.Model
	LastName  string `gorm:"last_name; not null"`  // 苗字
	FirstName string `gorm:"first_name; not null"` // 氏名
	UserName  string `gorm:"user_name; not null"`  // ユーザー名
	Password  string `gorm:"password; not null"`   // パスワード
	Email     string `gorm:"email; not null"`      // メールアドレス
}
