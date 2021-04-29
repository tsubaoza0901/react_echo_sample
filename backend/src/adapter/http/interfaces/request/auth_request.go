package request

// SignupInfo Signup用構造体
type SignupInfo struct {
	LastName  string `json:"last_name" example:"山田"`        // 苗字
	FirstName string `json:"first_name" example:"太郎"`       // 氏名
	UserName  string `json:"user_name" example:"たろう"`       // ユーザー名
	Email     string `json:"email" example:"xxx@gmail.com"` // メールアドレス
	Password  string `json:"password" example:"password"`   // パスワード
}

// LoginInfo Login用構造体
type LoginInfo struct {
	Email    string `json:"email" example:"xxx@gmail.com"` // メールアドレス
	Password string `json:"password" example:"password"`   // パスワード
}
