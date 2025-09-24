package repository

import (
	"context"
	"golang-kuliah-from-modul-3/app/model"
	"golang-kuliah-from-modul-3/database"
	"time"
)

func GetAllPekerjaan(ctx context.Context) ([]model.Pekerjaan, error) {
    rows, err := database.DB.QueryContext(ctx, `
        SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at
        FROM pekerjaan_alumni ORDER BY created_at DESC`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []model.Pekerjaan
    for rows.Next() {
        var p model.Pekerjaan
        if err := rows.Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan,
            &p.BidangIndustri, &p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja,
            &p.TanggalSelesaiKerja, &p.StatusPekerjaan, &p.DeskripsiPekerjaan,
            &p.CreatedAt, &p.UpdatedAt); err != nil {
            return nil, err
        }
        list = append(list, p)
    }
    return list, nil
}

func GetPekerjaanByID(ctx context.Context, id int) (*model.Pekerjaan, error) {
    var p model.Pekerjaan
    row := database.DB.QueryRowContext(ctx, `
        SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at
        FROM pekerjaan_alumni WHERE id=$1`, id)

    if err := row.Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan,
        &p.BidangIndustri, &p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja,
        &p.TanggalSelesaiKerja, &p.StatusPekerjaan, &p.DeskripsiPekerjaan,
        &p.CreatedAt, &p.UpdatedAt); err != nil {
        return nil, err
    }
    return &p, nil
}


func GetPekerjaanByAlumniID(ctx context.Context, alumniID int) ([]model.Pekerjaan, error) {
    rows, err := database.DB.QueryContext(ctx, `
        SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, 
               lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, 
               status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at
        FROM pekerjaan_alumni WHERE alumni_id = $1 ORDER BY created_at DESC`, alumniID)
    
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []model.Pekerjaan
    for rows.Next() {
        var p model.Pekerjaan
        if err := rows.Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan,
            &p.BidangIndustri, &p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja,
            &p.TanggalSelesaiKerja, &p.StatusPekerjaan, &p.DeskripsiPekerjaan,
            &p.CreatedAt, &p.UpdatedAt); err != nil {
            return nil, err
        }
        list = append(list, p)
    }
    return list, nil
}

func CreatePekerjaan(ctx context.Context, p *model.Pekerjaan) error {
    return database.DB.QueryRowContext(ctx, `
        INSERT INTO pekerjaan_alumni (alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at)
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING id`,
        p.AlumniID, p.NamaPerusahaan, p.PosisiJabatan, p.BidangIndustri, p.LokasiKerja,
        p.GajiRange, p.TanggalMulaiKerja, p.TanggalSelesaiKerja, p.StatusPekerjaan, p.DeskripsiPekerjaan,
        time.Now(), time.Now(),
    ).Scan(&p.ID)
}

func UpdatePekerjaan(ctx context.Context, p *model.Pekerjaan) (int64, error) {
    result, err := database.DB.ExecContext(ctx, `
        UPDATE pekerjaan_alumni SET nama_perusahaan=$1, posisi_jabatan=$2, bidang_industri=$3, lokasi_kerja=$4, gaji_range=$5, tanggal_mulai_kerja=$6, tanggal_selesai_kerja=$7, status_pekerjaan=$8, deskripsi_pekerjaan=$9, updated_at=$10 WHERE id=$11`,
        p.NamaPerusahaan, p.PosisiJabatan, p.BidangIndustri, p.LokasiKerja,
        p.GajiRange, p.TanggalMulaiKerja, p.TanggalSelesaiKerja, p.StatusPekerjaan, p.DeskripsiPekerjaan,
        time.Now(), p.ID)
    if err != nil {
        return 0, err
    }
    return result.RowsAffected()
}

func DeletePekerjaan(ctx context.Context, id int) (int64, error) {
    result, err := database.DB.ExecContext(ctx, "DELETE FROM pekerjaan_alumni WHERE id=$1", id)
    if err != nil {
        return 0, err
    }
    return result.RowsAffected()
}
