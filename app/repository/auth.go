package repository

import (
	"context"
	"database/sql"
	"golang-kuliah-from-modul-3/app/model"
	"golang-kuliah-from-modul-3/database"
)

// GetUserByIdentifier returns a user and its password hash by username or email
func GetUserByIdentifier(ctx context.Context, identifier string) (*model.User, string, error) {
	var user model.User
	var passwordHash string

	err := database.DB.QueryRowContext(ctx, `
		SELECT id, username, email, password_hash, role, created_at
		FROM users
		WHERE username = $1 OR email = $1
		LIMIT 1
	`, identifier).Scan(
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
