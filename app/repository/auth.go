package repository

import (
	"context"
	"database/sql"
	"golang-kuliah-from-modul-3/app/model"
	"golang-kuliah-from-modul-3/database"
)

func UserLogin(ctx context.Context, userLogin string) (*model.User, string, error) {
	var user model.User
	var passwordHash string

	err := database.DB.QueryRowContext(ctx, `
		SELECT id, username, email, password_hash, role, created_at
		FROM users
		WHERE username = $1 OR email = $1
		LIMIT 1
	`, userLogin).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&passwordHash,
		&user.Role,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, "", sql.ErrNoRows
		}
		return nil, "", err
	}

	return &user, passwordHash, nil
}


func CreateUser(ctx context.Context, user *model.User, passwordHash string) error{

	err := database.DB.QueryRowContext(ctx,
	`INSERT INTO users (username, email, password_hash, role)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at`,
		user.Username, user.Email, passwordHash, user.Role).Scan(&user.ID, &user.CreatedAt)

		return  err

}