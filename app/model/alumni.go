package model

import "time"

type Alumni struct {
	ID         int       `json:"id"`
	UserID     int        `json:"user_id"`
	NIM        string    `json:"nim"`
	Nama       string    `json:"nama"`
	Jurusan    string    `json:"jurusan"`
	Angkatan   int       `json:"angkatan"`
	TahunLulus int       `json:"tahun_lulus"`
	Email      string    `json:"email"`
	NoTelepon  *string   `json:"no_telepon"`
	Alamat     *string   `json:"alamat"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
	
}

// Request for filtering alumni employment status
// type AlumniEmploymentStatusRequest struct {
// 	ID                *int    json:"id" query:"id"
// 	Nama              *string json:"nama" query:"nama"
// 	Jurusan           *string json:"jurusan" query:"jurusan"
// 	Angkatan          *int    json:"angkatan" query:"angkatan"
// 	BidangIndustri    *string json:"bidang_industri" query:"bidang_industri"
// 	NamaPerusahaan    *string json:"nama_perusahaan" query:"nama_perusahaan"
// 	PosisiJabatan     *string json:"posisi_jabatan" query:"posisi_jabatan"
// 	LebihDari1Tahun   *int    json:"lebih_dari_1_tahun" query:"lebih_dari_1_tahun"
// 	Page              int     json:"page" query:"page"
// 	Limit             int     json:"limit" query:"limit"
// }