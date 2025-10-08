package service

import (
	"database/sql"
	"golang-kuliah-from-modul-3/app/model"
	"golang-kuliah-from-modul-3/app/repository"
	"golang-kuliah-from-modul-3/utils"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

func Login(c *fiber.Ctx) error {
	// Parsing Request Body (Tugas Handler)
	var req model.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Request body tidak valid"})
	}
	if req.Username == "" || req.Password == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Username dan password harus diisi"})
	}

	
	user, passwordHash, err := repository.UserLogin(c.Context(), req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			
			return c.Status(401).JSON(fiber.Map{"error": "Username atau password salah"})
		}

		
		return c.Status(500).JSON(fiber.Map{"error": "Error server"})
	}

	
	if !utils.CheckPassword(req.Password, passwordHash) {
		return c.Status(401).JSON(fiber.Map{"error": "Username atau password salah"})
	}

	
	token, err := utils.GenerateToken(*user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat token"})
	}

	

	
	response := model.LoginResponse{User: *user, Token: token}
	return c.JSON(fiber.Map{"success": true, "message": "Login berhasil", "data": response})
}



func Register(c *fiber.Ctx) error {
	// Parsing Request Body
	var req model.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Request body tidak valid"})
	}

	// Validasi Input
	if req.Username == "" || req.Email == "" || req.Password == ""{
		return c.Status(400).JSON(fiber.Map{"error": "Username, email, dan password harus diisi"})
	}
	
	// Default role adalah 'user' jika tidak diisi
	req.Role = strings.ToLower(req.Role)
	if req.Role != "admin" && req.Role != "user" {
		req.Role = "user"
	}


	//  Hash Password
	passwordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal memproses password"})
	}

	// Siapkan Model User
	newUser := model.User{
		Username: req.Username,
		Email:    req.Email,
		Role:     req.Role,
	}

	//  Panggil Repository untuk Menyimpan User
	if err := repository.CreateUser(c.Context(), &newUser, passwordHash); err != nil {
		// Cek apakah error disebabkan oleh unique constraint (username/email sudah ada)
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code.Name() == "unique_violation" {
			return c.Status(409).JSON(fiber.Map{"error": "Username atau email sudah terdaftar"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat user"})
	}

	// Kirim Response
	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "User berhasil dibuat",
		"data":    newUser,
	})
}





// func GetProfile(c *fiber.Ctx) error {
// 		userID := c.Locals("user_id").(int)
// 		username := c.Locals("username").(string)
// 		role := c.Locals("role").(string)
		
// 		return c.JSON(fiber.Map{
// 				"success": true,
// 				"message": "Profile berhasil diambil",
// 				"data": fiber.Map{
// 					"user_id": userID,
// 					"username": username,
// 					"role": role,
					
// 		},
// 	})
// }


func GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	username := c.Locals("username").(string)
	role := c.Locals("role").(string)

	listPekerjaan, err := repository.GetUserPekerjaan(c.Context(), userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Gagal mengambil data pekerjaan",
		})
	}


	response := model.ProfileResponse{
		Profile: model.ProfileData{
			UserID:   userID,
			Username: username,
			Role:     role,
		},
		Pekerjaan: listPekerjaan,
	}

	
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Profile berhasil diambil",
		"data":    response,
	})
}



func GetUsersService(c *fiber.Ctx) error {
	
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sortBy := c.Query("sortBy", "id")
	order := c.Query("order", "asc")
	search := c.Query("search", "")

	
	offset := (page - 1) * limit

	
	sortByWhitelist := map[string]bool{
		"id":         true,
		"name":       true,
		"email":      true,
		"created_at": true,
	}
	if !sortByWhitelist[sortBy] {
		sortBy = "id" 
	}
	if strings.ToLower(order) != "desc" {
		order = "asc" 
	}

	
	users, err := repository.GetUserRepo(search, sortBy, order, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data pengguna"}) 
	}

	
	total, err := repository.CountUsersRepo(search)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghitung total pengguna"}) 
	}

	
	response := model.UserResponse{
		Data: users,
		Meta: model.MetaInfo{
			Page:   page,
			Limit:  limit,
			Total:  total,
			Pages:  (total + limit - 1) / limit, 
			SortBy: sortBy,
			Order:  order,
			Search: search,
		},
	}
	
	return c.JSON(response) 
}


func DeletedUsers(c *fiber.Ctx) error {
    ctx := c.Context()
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
    }

    rows, err := repository.DeleteUsers(ctx, id)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Gagal hapus DATA USERS"})
    }
    if rows == 0 {
        return c.Status(404).JSON(fiber.Map{"error": "USERS tidak ditemukan"})
    }

    return c.JSON(fiber.Map{"success": true, "message": "USERS berhasil dihapus"})
}



func Getuserbyalumni(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(int)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Tidak dapat mengidentifikasi pengguna",
		})
	}
	list, err := repository.GetUserPekerjaan(c.Context(), userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data pekerjaan"})
	}

	return c.JSON(fiber.Map{"success": true, "data": list})
}

func DeletePekerjaanByUser(c *fiber.Ctx) error {
	pekerjaanID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "id pekerjaan tidak valid"})
	}

	role := c.Locals("role").(string)
	userID := c.Locals("user_id").(int)
	ctx := c.Context()
	var rowsAffected int64

	if role == "admin" {
		rowsAffected, err = repository.DeletePekerjaan(ctx, pekerjaanID, userID)
	} else {
		rowsAffected, err = repository.DeletePekerjaanByUser(ctx, pekerjaanID, userID)
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus pekerjaan"})
	}

	
	if rowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Pekerjaan tidak ditemukan atau Anda tidak memiliki akses"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Pekerjaan berhasil dihapus",
	})
}

