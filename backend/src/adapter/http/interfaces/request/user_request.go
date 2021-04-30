package request

type UpdateUser struct {
	ID        uint   `json:"id" example:"1"`                                 // ユーザーID
	UpdatedAt string `json:"updated_at" example:"2020-11-30T11:25:30+09:00"` // 更新日時
	LastName  string `json:"last_name" example:"山田"`                         // 苗字
	FirstName string `json:"first_name" example:"太郎"`                        // 氏名
	UserName  string `json:"user_name" example:"Tom"`                        // ユーザー名
	Email     string `json:"email" example:"xxx@gmail.com"`                  // メールアドレス
	Password  string `json:"password" example:"password"`                    // パスワード
}

type SearchUser struct {
	ID       uint `query:"user_id"`   // ユーザーID
	DemandPW bool `query:"demand_pw"` // パスワード取得の要否
}

// UserDeleteRequest ...
type DeleteUser struct {
	ID        uint   `json:"id" example:"1"`
	UpdatedAt string `json:"updated_at" example:"2020-11-30T11:25:30+09:00"`
}
