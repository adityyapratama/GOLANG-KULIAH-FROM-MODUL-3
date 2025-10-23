package service

import (
	"golang-kuliah-from-modul-3/app/model"
	"golang-kuliah-from-modul-3/app/repository"
	"golang-kuliah-from-modul-3/utils"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)



type IAuthService interface{
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type AuthService struct{
	repo repository.IAuthRepository
}

func NewAuthService(repo repository.IAuthRepository) IAuthService{
	return &AuthService{repo: repo}
}


func (s *AuthService) Login(c *fiber.Ctx) error {
    var req model.LoginRequest

    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "error": "request body tidak valid"})
    }
    if req.Username == "" || req.Password == "" {
        return c.Status(400).JSON(fiber.Map{
            "error": "username dan password kudu di isi"})
    }
    user, err := s.repo.GetUserByLogin(c.Context(), req.Username)

    if err != nil {
        if err == mongo.ErrNoDocuments {
            return c.Status(401).JSON(fiber.Map{
                "error": "username atau password salah"})
        }
        return c.Status(500).JSON(fiber.Map{
            "error": "error server"})
    }

    if user == nil {
        return c.Status(401).JSON(fiber.Map{
            "error": "username atau password salah"})
    }

    if !utils.CheckPassword(req.Password, user.PasswordHash) {
        return c.Status(401).JSON(fiber.Map{"error": "Username atau password salah"})
    }

    
    token, err := utils.GenerateToken(*user) 
    
    
    if err != nil { 
        return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat token"})
    }

    
    user.PasswordHash = "" 

    
    response := model.LoginResponse{User: *user, Token: token} 
    
    return c.JSON(fiber.Map{
        "success": true,
        "message": "login berhasil",
        "data":    response})

    
}


func (s *AuthService)Register (c *fiber.Ctx) error {

	var req model.RegisterRequest
	if err :=c.BodyParser(&req); err !=nil{
		return c.Status(400).JSON(fiber.Map{"error":"Request body tidak valid"})
	}
	if req.Username == "" || req.Email == "" || req.Password == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Username, email, dan password harus diisi"})
	}

	req.Role = strings.ToLower(req.Role)
	if req.Role != "admin" && req.Role != "user" {
		req.Role = "user"
	}

	passwordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal memproses password"})
	}

	user := &model.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: passwordHash, // Masukkan hash ke model
		Role:         req.Role,
	}

	if err := s.repo.CreateUser(c.Context(), user); err != nil {
		if mongo.IsDuplicateKeyError(err){
			return c.Status(409).JSON(fiber.Map{"error": "Username atau email sudah terdaftar"})
		}
		log.Println("!!! ERROR SAAT CREATE USER:", err)
		return  c.Status(500).JSON(fiber.Map{"error":"gagal membuat user"})
		}

		
		user.PasswordHash = ""
		return c.Status(201).JSON(fiber.Map{"success": true, "message": "User berhasil didaftarkan", "data": user})	

}
