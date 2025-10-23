package repository

// import (
// 	"context"
// 	"golang-kuliah-from-modul-3/app/model"
// 	"golang-kuliah-from-modul-3/database"
// 	"time"
// )

// func GetAllMahasiswa(ctx context.Context) ([]model.Mahasiswa, error) {
//     rows, err := database.DB.QueryContext(ctx, `
//         SELECT id, nim, nama, jurusan, angkatan, email, created_at, updated_at
//         FROM mahasiswa ORDER BY created_at DESC`)
//     if err != nil {
//         return nil, err
//     }
//     defer rows.Close()

//     var list []model.Mahasiswa
//     for rows.Next() {
//         var m model.Mahasiswa
//         if err := rows.Scan(&m.ID, &m.NIM, &m.Nama, &m.Jurusan,
//             &m.Angkatan, &m.Email, &m.CreatedAt, &m.UpdatedAt); err != nil {
//             return nil, err
//         }
//         list = append(list, m)
//     }
//     return list, nil
// }

// func GetMahasiswaByID(ctx context.Context, id int) (*model.Mahasiswa, error) {
//     var m model.Mahasiswa
//     row := database.DB.QueryRowContext(ctx, `
//         SELECT id, nim, nama, jurusan, angkatan, email, created_at, updated_at
//         FROM mahasiswa WHERE id=$1`, id)

//     if err := row.Scan(&m.ID, &m.NIM, &m.Nama, &m.Jurusan,
//         &m.Angkatan, &m.Email, &m.CreatedAt, &m.UpdatedAt); err != nil {
//         return nil, err
//     }
//     return &m, nil
// }

// func CreateMahasiswa(ctx context.Context, m *model.Mahasiswa) error {
//     return database.DB.QueryRowContext(ctx, `
//         INSERT INTO mahasiswa (nim, nama, jurusan, angkatan, email, created_at, updated_at)
//         VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`,
//         m.NIM, m.Nama, m.Jurusan, m.Angkatan, m.Email, time.Now(), time.Now(),
//     ).Scan(&m.ID)
// }

// func UpdateMahasiswa(ctx context.Context, m *model.Mahasiswa) (int64, error) {
//     result, err := database.DB.ExecContext(ctx, `
//         UPDATE mahasiswa SET nama=$1, jurusan=$2, angkatan=$3, email=$4, updated_at=$5 WHERE id=$6`,
//         m.Nama, m.Jurusan, m.Angkatan, m.Email, time.Now(), m.ID)
//     if err != nil {
//         return 0, err
//     }
//     return result.RowsAffected()
// }

// func DeleteMahasiswa(ctx context.Context, id int) (int64, error) {
//     result, err := database.DB.ExecContext(ctx, "DELETE FROM mahasiswa WHERE id=$1", id)
//     if err != nil {
//         return 0, err
//     }
//     return result.RowsAffected()
// }
