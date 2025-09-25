package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID        int      `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Role      string   `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}


type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"` // 'admin' atau 'user'
}



type LoginResponse struct {
		User User `json:"user"`
		Token string `json:"token"`
}


type JWTClaims struct {
		UserID int `json:"user_id"`
		Username string `json:"username"`
		Role string `json:"role"`
		jwt.RegisteredClaims
}