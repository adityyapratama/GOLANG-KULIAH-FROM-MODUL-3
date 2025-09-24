package service

import (
	"context"
	"database/sql"
	"golang-kuliah-from-modul-3/app/model"
	"golang-kuliah-from-modul-3/app/repository"
	"golang-kuliah-from-modul-3/utils"

	"github.com/gofiber/fiber/v2"
)

// Login performs login logic and returns a token and user
func Login(ctx context.Context, identifier string, password string) (*model.User, string, error) {
	user, passwordHash, err := repository.GetUserByIdentifier(ctx, identifier)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, "", fiber.ErrUnauthorized
		}
		return nil, "", err
	}

	if !utils.CheckPassword(password, passwordHash) {
		return nil, "", fiber.ErrUnauthorized
	}

	token, err := utils.GenerateToken(*user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

// LoginHandler is a Fiber handler example that uses the service Login function
func LoginHandler(c *fiber.Ctx) error {
	var req model.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Request body tidak valid"})
	}
	if req.Username == "" || req.Password == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Username dan password harus diisi"})
	}

	user, token, err := Login(c.Context(), req.Username, req.Password)
	if err != nil {
		if err == fiber.ErrUnauthorized {
			return c.Status(401).JSON(fiber.Map{"error": "Username atau password salah"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Error server"})
	}

	resp := model.LoginResponse{User: *user, Token: token}
	return c.JSON(fiber.Map{"success": true, "message": "Login berhasil", "data": resp})
}
