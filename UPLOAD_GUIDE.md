# 📤 Panduan Upload File - Go Fiber MongoDB

## 🎯 Penjelasan `BodyLimit` di Fiber

### Apa itu `BodyLimit`?
`BodyLimit` adalah konfigurasi **global** di Fiber untuk membatasi ukuran maksimal request body yang bisa diterima server.

```go
app := fiber.New(fiber.Config{
    BodyLimit: 10 * 1024 * 1024, // 10MB maksimal
})
```

### Kenapa Perlu `BodyLimit`?
✅ **Keamanan**: Melindungi server dari serangan DoS dengan upload file besar  
✅ **Performa**: Mencegah memory overflow  
✅ **Resource Management**: Mengontrol penggunaan bandwidth  

---

## 📁 Struktur File Upload di Project Ini

```
GOLANG-KULIAH-FROM-MODUL-3/
├── app/
│   ├── model/
│   │   └── file.go              # Model data file
│   ├── repository/
│   │   └── file.go              # Database operations
│   └── service/
│       └── file.go              # Business logic + validasi
├── config/
│   └── app.go                   # ✅ BodyLimit dikonfigurasi di sini
├── route/
│   └── file_route.go            # Endpoint upload
├── middleware/
│   └── auth.go                  # Middleware autentikasi
├── uploads/                     # Folder penyimpanan file
│   ├── foto/                    # Subfolder untuk foto
│   └── sertifikat/              # Subfolder untuk sertifikat
└── .env                         # Konfigurasi environment
```

---

## ⚙️ Konfigurasi di Project Ini

### 1. **Global Config** (`config/app.go`)
```go
app := fiber.New(fiber.Config{
    BodyLimit: 10 * 1024 * 1024, // 10MB - limit global
})

app.Static("/uploads", "./uploads") // Serve uploaded files
```

### 2. **Environment Variables** (`.env`)
```env
UPLOAD_PATH=./uploads
APP_PORT=3000
```

### 3. **Validasi di Service Layer** (`app/service/file.go`)

#### ✅ Upload Foto (Max 1MB)
```go
// Validasi ukuran: 1MB
if fileHeader.Size > 1*1024*1024 {
    return error("File size exceeds 1MB")
}

// Validasi tipe: jpg/jpeg/png
allowedTypes := map[string]bool{
    "image/jpeg": true,
    "image/jpg":  true,
    "image/png":  true,
}
```

#### ✅ Upload Sertifikat (Max 2MB)
```go
// Validasi ukuran: 2MB
if fileHeader.Size > 2*1024*1024 {
    return error("File size exceeds 2MB")
}

// Validasi tipe: pdf
allowedTypes := map[string]bool{
    "application/pdf": true,
}
```

---

## 🔐 Implementasi Middleware untuk Tugas

### Requirement dari Modul:
1. **Admin** → Bisa upload untuk semua user
2. **User** → Hanya bisa upload untuk dirinya sendiri

### Solusi: Role-based Middleware

```go
// middleware/upload_auth.go
func CheckUploadPermission(c *fiber.Ctx) error {
    role := c.Locals("role").(string)
    userID := c.Locals("user_id").(string)
    targetUserID := c.Params("user_id") // dari URL

    // Admin bisa upload untuk siapa saja
    if role == "admin" {
        return c.Next()
    }

    // User hanya bisa upload untuk dirinya sendiri
    if role == "user" && userID == targetUserID {
        return c.Next()
    }

    return c.Status(403).JSON(fiber.Map{
        "error": "Forbidden: Anda tidak punya akses",
    })
}
```

---

## 🚀 Endpoint API untuk Tugas

### 1. Upload Foto Profil
```http
POST /api/users/:user_id/foto
Authorization: Bearer {token}
Content-Type: multipart/form-data

Body:
- file: [foto.jpg] (max 1MB, jpg/jpeg/png)
```

**Response Success:**
```json
{
  "success": true,
  "message": "Foto uploaded successfully",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "file_name": "uuid-generated.jpg",
    "original_name": "foto-profil.jpg",
    "file_path": "./uploads/foto/uuid-generated.jpg",
    "file_size": 512000,
    "file_type": "image/jpeg",
    "uploaded_at": "2025-10-30T10:30:00Z"
  }
}
```

### 2. Upload Sertifikat/Ijazah
```http
POST /api/users/:user_id/sertifikat
Authorization: Bearer {token}
Content-Type: multipart/form-data

Body:
- file: [ijazah.pdf] (max 2MB, pdf only)
```

**Response Success:**
```json
{
  "success": true,
  "message": "Sertifikat uploaded successfully",
  "data": {
    "id": "507f1f77bcf86cd799439012",
    "file_name": "uuid-generated.pdf",
    "original_name": "ijazah.pdf",
    "file_path": "./uploads/sertifikat/uuid-generated.pdf",
    "file_size": 1800000,
    "file_type": "application/pdf",
    "uploaded_at": "2025-10-30T10:35:00Z"
  }
}
```

---

## 📝 Testing dengan Postman

### Test Upload Foto (sebagai User)
1. Login dulu untuk dapat token
2. Request:
   - Method: `POST`
   - URL: `http://localhost:3000/api/users/{your_user_id}/foto`
   - Headers: `Authorization: Bearer {token}`
   - Body: form-data
     - Key: `file` (type: File)
     - Value: pilih file foto (jpg/png, max 1MB)

### Test Upload Sertifikat (sebagai Admin)
1. Login sebagai admin
2. Request:
   - Method: `POST`
   - URL: `http://localhost:3000/api/users/{any_user_id}/sertifikat`
   - Headers: `Authorization: Bearer {admin_token}`
   - Body: form-data
     - Key: `file` (type: File)
     - Value: pilih file PDF (max 2MB)

---

## 🔍 Perbedaan: BodyLimit vs Validasi di Service

| Aspek | BodyLimit (Global) | Validasi Service (Per-endpoint) |
|-------|-------------------|--------------------------------|
| **Lokasi** | `config/app.go` | `app/service/file.go` |
| **Scope** | Semua request | Spesifik per endpoint |
| **Ukuran** | 10MB (contoh) | 1MB foto, 2MB sertifikat |
| **Error** | `413 Request Entity Too Large` | Custom error message |
| **Kapan dipakai** | Before parsing | After parsing, before save |

### Contoh Flow:
```
Request 15MB → ❌ Ditolak oleh BodyLimit (413)
Request 3MB → ✅ Lolos BodyLimit
            → Masuk service UploadFoto
            → ❌ Ditolak validasi (max 1MB untuk foto)
Request 800KB foto → ✅ Lolos semua validasi → Saved
```

---

## 🛡️ Best Practices

1. **Gunakan UUID untuk nama file** - Hindari collision
2. **Validasi MIME type** - Jangan hanya cek extension
3. **Bersihkan file jika gagal save ke DB** - Hindari orphan files
4. **Gunakan subfolder** - `/uploads/foto/`, `/uploads/sertifikat/`
5. **Log semua aktivitas upload** - Audit trail
6. **Batasi request rate** - Rate limiting per user

---

## 🐛 Troubleshooting

### Error: "Request Entity Too Large"
- **Penyebab**: File lebih besar dari `BodyLimit`
- **Solusi**: Naikkan `BodyLimit` di `config/app.go` atau kompres file

### Error: "File type not allowed"
- **Penyebab**: MIME type tidak sesuai allowlist
- **Solusi**: Cek Content-Type header, pastikan file asli (bukan rename extension)

### Error: "Failed to save file"
- **Penyebab**: Permission folder uploads atau disk penuh
- **Solusi**: `chmod 755 uploads/` atau bersihkan disk

---

## 📚 Referensi
- [Fiber Documentation - Config](https://docs.gofiber.io/api/fiber#config)
- [MongoDB Go Driver](https://www.mongodb.com/docs/drivers/go/current/)
- [Go UUID Package](https://github.com/google/uuid)

---

**✅ Status**: Konfigurasi sudah diterapkan di project ini  
**🔧 Perlu dikerjakan**: Implementasi 2 endpoint sesuai tugas modul
