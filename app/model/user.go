package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string   `bson:"username" json:"username"`
	Email     string   `bson:"email" json:"email"`
	PasswordHash string `bson:"password_hash" json:"-"`
	Role      string   `bson:"role" json:"role"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	DeletedAt  *time.Time `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}


type LoginRequest struct {
		Username string `bson:"username" json:"username"`
		Password string `bson:"password" json:"password"`
}

type RegisterRequest struct {
	Username string `bson:"username" json:"username"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
	Role     string `bson:"role" json:"role"` // 'admin' atau 'user'
}



type LoginResponse struct {
		User User `bson:"user" json:"user"`
		Token string `bson:"token" json:"token"`
}


type JWTClaims struct {
		UserID   string `json:"user_id"`
		Username string `json:"username"`
		Role     string `json:"role"`
		jwt.RegisteredClaims
		
}


type ProfileData struct {
	UserID   string    `json:"user_id"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
}

// ProfileResponse adalah struktur lengkap untuk endpoint /profile
type ProfileResponse struct {
	Profile   ProfileData `bson:"profile" json:"profile"`
	Pekerjaan []Pekerjaan `bson:"pekerjaan" json:"pekerjaan"`
}