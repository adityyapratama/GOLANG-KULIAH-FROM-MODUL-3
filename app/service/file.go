package service

import (
	"fmt"
	"golang-kuliah-from-modul-3/app/model"
	"golang-kuliah-from-modul-3/app/repository"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileService interface {
	UploadFile(c *fiber.Ctx) error
	GetAllFiles(c *fiber.Ctx) error
	GetFileByID(c *fiber.Ctx) error
	DeleteFile(c *fiber.Ctx) error

	// ✅ Endpoint khusus untuk tugas
	UploadFoto(c *fiber.Ctx) error       // Max 1MB, jpg/jpeg/png
	UploadSertifikat(c *fiber.Ctx) error // Max 2MB, pdf
}

type fileService struct {
	repo       repository.FileRepository
	uploadPath string
}

// Constructor
func NewFileService(repo repository.FileRepository, uploadPath string) FileService {
	return &fileService{
		repo:       repo,
		uploadPath: uploadPath,
	}
}

// upload File
func (s *fileService) UploadFile(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "No file uploaded",
			"error":   err.Error(),
		})
	}

	// Validasi ukuran file (maks 10MB)
	if fileHeader.Size > 10*1024*1024 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "File size exceeds 10MB",
		})
	}

	// Validasi tipe file
	allowedTypes := map[string]bool{
		"image/jpeg":      true,
		"image/png":       true,
		"image/jpg":       true,
		"application/pdf": true,
	}

	contentType := fileHeader.Header.Get("Content-Type")
	if !allowedTypes[contentType] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "File type not allowed",
		})
	}

	// untuk generate nama file unik
	ext := filepath.Ext(fileHeader.Filename)
	newFileName := uuid.New().String() + ext
	filePath := filepath.Join(s.uploadPath, newFileName)

	// Buat folder jika belum ada
	if err := os.MkdirAll(s.uploadPath, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to create upload directory",
			"error":   err.Error(),
		})
	}

	// Simpan file ke direktori
	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to open file",
			"error":   err.Error(),
		})
	}
	defer file.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to save file",
			"error":   err.Error(),
		})
	}
	defer out.Close()

	if _, err := out.ReadFrom(file); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to write file",
			"error":   err.Error(),
		})
	}

	// Simpan metadata ke database
	fileModel := &model.File{
		FileName:     newFileName,
		OriginalName: fileHeader.Filename,
		FilePath:     filePath,
		FileSize:     fileHeader.Size,
		FileType:     contentType,
		UploadedAt:   time.Now(),
	}

	if err := s.repo.Create(fileModel); err != nil {
		os.Remove(filePath)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to save file metadata",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "File uploaded successfully",
		"data":    s.toFileResponse(fileModel),
	})
}


func (s *fileService) GetAllFiles(c *fiber.Ctx) error {
	files, err := s.repo.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to get files",
			"error":   err.Error(),
		})
	}

	var responses []model.FileResponse
	for _, f := range files {
		responses = append(responses, *s.toFileResponse(&f))
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Files retrieved successfully",
		"data":    responses,
	})
}


func (s *fileService) GetFileByID(c *fiber.Ctx) error {
	id := c.Params("id")

	file, err := s.repo.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "File not found",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "File retrieved successfully",
		"data":    s.toFileResponse(file),
	})
}


func (s *fileService) DeleteFile(c *fiber.Ctx) error {
	fileID := c.Params("id")

	
	role, ok := c.Locals("role").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "Unauthorized: Role tidak ditemukan",
		})
	}

	userID, ok := c.Locals("user_id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "Unauthorized: User ID tidak ditemukan",
		})
	}

	
	file, err := s.repo.FindByID(fileID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "File not found",
			"error":   err.Error(),
		})
	}

	var deletedCount int64

	
	if role == "admin" {
		err = s.repo.Delete(fileID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "Failed to delete file",
				"error":   err.Error(),
			})
		}
		deletedCount = 1 
	} else {
		deletedCount, err = s.repo.DeleteByUser(fileID, userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": "Failed to delete file",
				"error":   err.Error(),
			})
		}
	}

	
	if deletedCount == 0 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"success": false,
			"error":   "File tidak ditemukan atau Anda tidak memiliki akses",
		})
	}

	
	if err := os.Remove(file.FilePath); err != nil {
		fmt.Println("Warning: Failed to delete file from storage:", err)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "File deleted successfully",
	})
}


func (s *fileService) toFileResponse(file *model.File) *model.FileResponse {
	return &model.FileResponse{
		ID:           file.ID.Hex(),
		UserID:       file.UserID.Hex(),
		Category:     file.Category,
		FileName:     file.FileName,
		OriginalName: file.OriginalName,
		FilePath:     file.FilePath,
		FileSize:     file.FileSize,
		FileType:     file.FileType,
		UploadedAt:   file.UploadedAt,
		UploadedBy:   file.UploadedBy.Hex(),
	}
}

// upload foto max 1mb
func (s *fileService) UploadFoto(c *fiber.Ctx) error {
	return s.uploadWithValidation(c, "foto", 1*1024*1024, map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
	})
}

//upload Sertifikat - Max 2MB, pdf 
func (s *fileService) UploadSertifikat(c *fiber.Ctx) error {
	return s.uploadWithValidation(c, "sertifikat", 2*1024*1024, map[string]bool{
		"application/pdf": true,
	})
}

// Helper untuk Upload dengan validasi custom
func (s *fileService) uploadWithValidation(c *fiber.Ctx, category string, maxSize int64, allowedTypes map[string]bool) error {
	// ambil user yg sedang login
	uploadedByStr, ok := c.Locals("user_id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "Unauthorized: User ID tidak ditemukan",
		})
	}

	uploadedBy, err := primitive.ObjectIDFromHex(uploadedByStr)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "Unauthorized: User ID tidak valid",
		})
	}

	
	var targetUserID primitive.ObjectID
	targetUserIDStr := c.Params("user_id")

	if targetUserIDStr != "" {
		targetUserID, err = primitive.ObjectIDFromHex(targetUserIDStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "User ID tidak valid",
			})
		}
	} else {
		targetUserID = uploadedBy
	}

	// Get file from form
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "No file uploaded",
			"error":   err.Error(),
		})
	}

	
	if fileHeader.Size > maxSize {
		sizeInMB := maxSize / (1024 * 1024)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintf("File size exceeds %dMB", sizeInMB),
		})
	}

	
	contentType := fileHeader.Header.Get("Content-Type")
	if !allowedTypes[contentType] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintf("File type not allowed for %s", category),
		})
	}

	
	ext := filepath.Ext(fileHeader.Filename)
	newFileName := uuid.New().String() + ext

	// Buat subfolder untuk kategori (foto/sertifikat)
	categoryPath := filepath.Join(s.uploadPath, category)
	filePath := filepath.Join(categoryPath, newFileName)

	// Buat folder jika belum ada
	if err := os.MkdirAll(categoryPath, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to create upload directory",
			"error":   err.Error(),
		})
	}

	// Simpan file ke direktori
	if err := c.SaveFile(fileHeader, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to save file",
			"error":   err.Error(),
		})
	}

	// Simpan metadata ke database
	fileModel := &model.File{
		UserID:       targetUserID,
		Category:     category,
		FileName:     newFileName,
		OriginalName: fileHeader.Filename,
		FilePath:     filePath,
		FileSize:     fileHeader.Size,
		FileType:     contentType,
		UploadedAt:   time.Now(),
		UploadedBy:   uploadedBy,
	}

	if err := s.repo.Create(fileModel); err != nil {
		// Hapus file jika gagal simpan ke database
		os.Remove(filePath)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to save file metadata",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("%s uploaded successfully", category),
		"data":    s.toFileResponse(fileModel),
	})
}
