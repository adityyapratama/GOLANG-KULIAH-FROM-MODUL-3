package repository

import (
	"context"
	"golang-kuliah-from-modul-3/app/model"
	"golang-kuliah-from-modul-3/database"
	"time"
)

func GetAllAlumni(ctx context.Context) ([]model.Alumni, error) {
    rows, err := database.DB.QueryContext(ctx, `
        SELECT id, nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, created_at, updated_at
        FROM alumni ORDER BY created_at DESC`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []model.Alumni
    for rows.Next() {
        var a model.Alumni
        if err := rows.Scan(&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan,
            &a.TahunLulus, &a.Email, &a.NoTelepon, &a.Alamat, &a.CreatedAt, &a.UpdatedAt); err != nil {
            return nil, err
        }
        list = append(list, a)
    }
    return list, nil
}

func GetAlumniByID(ctx context.Context, id int) (*model.Alumni, error) {
    var a model.Alumni
    row := database.DB.QueryRowContext(ctx, `
        SELECT id, nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, created_at, updated_at
        FROM alumni WHERE id=$1`, id)

    if err := row.Scan(&a.ID, &a.NIM, &a.Nama, &a.Jurusan, &a.Angkatan,
        &a.TahunLulus, &a.Email, &a.NoTelepon, &a.Alamat, &a.CreatedAt, &a.UpdatedAt); err != nil {
        return nil, err
    }
    return &a, nil
}

func CreateAlumni(ctx context.Context, a *model.Alumni) error {
    return database.DB.QueryRowContext(ctx, `
        INSERT INTO alumni (nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat, created_at, updated_at)
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`,
        a.NIM, a.Nama, a.Jurusan, a.Angkatan, a.TahunLulus, a.Email, a.NoTelepon, a.Alamat, time.Now(), time.Now(),
    ).Scan(&a.ID)
}

func UpdateAlumni(ctx context.Context, a *model.Alumni) (int64, error) {
    result, err := database.DB.ExecContext(ctx, `
        UPDATE alumni SET nama=$1, jurusan=$2, angkatan=$3, tahun_lulus=$4, email=$5, no_telepon=$6, alamat=$7, updated_at=$8 WHERE id=$9`,
        a.Nama, a.Jurusan, a.Angkatan, a.TahunLulus, a.Email, a.NoTelepon, a.Alamat, time.Now(), a.ID)
    if err != nil {
        return 0, err
    }
    return result.RowsAffected()
}

func DeleteAlumni(ctx context.Context, id int) (int64, error) {
    result, err := database.DB.ExecContext(ctx, "DELETE FROM alumni WHERE id=$1", id)
    if err != nil {
        return 0, err
    }
    return result.RowsAffected()
}
