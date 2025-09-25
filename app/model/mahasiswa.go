package model
import "time"

type Mahasiswa struct {
		ID int `json:"id"`
		NIM string `json:"nim"`
		Nama string `json:"nama"`
		Jurusan string `json:"jurusan"`
		Angkatan int `json:"angkatan"`
		Email string `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt  *time.Time `json:"deleted_at,omitempty"`

}

type CreateMahasiswaRequest struct{
		NIM string `json:"nim"`
		Nama string `json:"nama"`
		Jurusan string `json:"jurusan"`
		Angkatan int `json:"angkatan"`
		Email string `json:"email"`
}

type UpdateMahasiswaRequest struct{
		Nama string `json:"nama"`
		Jurusan string `json:"jurusan"`
		Angkatan int `json:"angkatan"`
		Email string `json:"email"`
}