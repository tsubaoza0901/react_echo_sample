package response

import "time"

// UserResponse UserResponse構造体
// 役割：
type UserResponse struct {
	ID        uint      `json:"id" example:"1"`
	UpdatedAt time.Time `json:"updated_at" example:"2020-02-02T10:03:48.1292"`
	LastName  string    `json:"last_name" example:"山田"`
	FirstName string    `json:"first_name" example:"太郎"`
}
