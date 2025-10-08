package repository

import (
	"context"
	"database/sql"
	"fmt"
	"golang-kuliah-from-modul-3/app/model"
	"golang-kuliah-from-modul-3/database"
	"time"

	
)

func GetAllPekerjaan(ctx context.Context) ([]model.Pekerjaan, error) {
    rows, err := database.DB.QueryContext(ctx, `
       SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at, deleted_at, deleted_by
       FROM pekerjaan_alumni WHERE deleted_by IS NULL ORDER BY created_at DESC`)
        
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
            &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt,&p.DeletedBy); err != nil {
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

// func DeletePekerjaan(ctx context.Context, id int) (int64, error) {
//     result, err := database.DB.ExecContext(ctx, "DELETE FROM pekerjaan_alumni WHERE id=$1", id)
//     if err != nil {
//         return 0, err
//     }
//     return result.RowsAffected()
// }


func DeletePekerjaan(ctx context.Context, pekerjaanID int, deletedByID int) (int64, error) {
	result, err := database.DB.ExecContext(ctx,
		"UPDATE pekerjaan_alumni SET deleted_at = NOW(), deleted_by = $1 WHERE id = $2 AND deleted_at IS NULL",
		deletedByID, pekerjaanID) 
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}



func CountWorkAlumni(ctx context.Context, alumniID int) (*model.TotalMasaKerja, error) {
    query := `
        SELECT
            FLOOR(total_bulan_all / 12) AS total_tahun,
            (total_bulan_all::integer % 12) AS total_bulan
        FROM (
            SELECT
                SUM(EXTRACT(YEAR FROM age(COALESCE(tanggal_selesai_kerja, NOW()), tanggal_mulai_kerja)) * 12 +
                    EXTRACT(MONTH FROM age(COALESCE(tanggal_selesai_kerja, NOW()), tanggal_mulai_kerja))) AS total_bulan_all
            FROM
                pekerjaan_alumni
            WHERE
                alumni_id = $1 AND deleted_at IS NULL
        ) AS subquery`

    var totalMasaKerja model.TotalMasaKerja
    row := database.DB.QueryRowContext(ctx, query, alumniID)
    
    err := row.Scan(&totalMasaKerja.Tahun, &totalMasaKerja.Bulan)
    if err != nil {
        if err == sql.ErrNoRows {
            return &model.TotalMasaKerja{Tahun: 0, Bulan: 0}, nil
        }
        return nil, fmt.Errorf("gagal memindai hasil query: %w", err)
    }

    return &totalMasaKerja, nil
}







// func GetAllPekerjaanTrash(ctx context.Context) (*model.Pekerjaan, error) {
//     var p model.Pekerjaan
//     row := database.DB.QueryRowContext(ctx, `
//         SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at
//         deleted_at,deleted_by FROM pekerjaan_alumni WHERE deleted_by IS NOT NULL ORDER BY created_at DESC`)

//     if err := row.Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan,
//         &p.BidangIndustri, &p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja,
//         &p.TanggalSelesaiKerja, &p.StatusPekerjaan, &p.DeskripsiPekerjaan,
//         &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt, &p.DeletedBy); err != nil {
//         return nil, err
//     }
//     return &p, nil
// }




func GetAllPekerjaanTrash(ctx context.Context) ([]model.Pekerjaan, error) {
    rows, err := database.DB.QueryContext(ctx, `
       SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at, deleted_at, deleted_by
       FROM pekerjaan_alumni WHERE deleted_by IS NOT NULL ORDER BY created_at DESC`)
        
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
            &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt,&p.DeletedBy); err != nil {
            return nil, err
        }
        list = append(list, p)
    }
    return list, nil
}



func DeletePekerjaanTrash(ctx context.Context, id int) (int64, error) {
    result, err := database.DB.ExecContext(ctx, "DELETE FROM pekerjaan_alumni WHERE deleted_by IS NOT NULL and id=$1", id)
    if err != nil {
        return 0, err
    }
    return result.RowsAffected()
}

func DeletePekerjaanTrashByUser(ctx context.Context, pekerjaanID int, userID int) (int64, error) {
    query := `
        DELETE FROM pekerjaan_alumni
        WHERE id = $1 
          AND deleted_at IS NOT NULL
          AND alumni_id IN (
              SELECT id FROM alumni WHERE user_id = $2
          )`
    result, err := database.DB.ExecContext(ctx, query, pekerjaanID, userID)
    if err != nil {
        return 0, err
    }
    return result.RowsAffected()
}







// restore buat admin
func DeletePekerjaanTrashRestore(ctx context.Context, pekerjaanID int, deletedByID int) (int64, error) {
    result, err := database.DB.ExecContext(ctx, 
        "UPDATE pekerjaan_alumni SET deleted_at = NULL, deleted_by = NULL WHERE id = $1 AND deleted_by IS NOT NULL", 
        pekerjaanID) 
    if err != nil {
        return 0, err
    }
    return result.RowsAffected()
}

// func DeleteAlumni(ctx context.Context, id int) (int64, error) {
//     result, err := database.DB.ExecContext(ctx, "UPDATE alumni set deleted_at= NOW() WHERE id = $1 AND deleted_at IS NULL", id)
//     if err != nil {
//         return 0, err
//     }
//     return result.RowsAffected()
// }

// buat user
func RestoreTrashPekerjaanByUser(ctx context.Context, pekerjaanID int, userID int) (int64, error) {
    query := `
        UPDATE pekerjaan_alumni pa
        SET deleted_at = NULL, deleted_by = NULL
        FROM alumni a
        WHERE pa.alumni_id = a.id
          AND pa.id = $1
          AND a.user_id = $2
          AND pa.deleted_at IS NOT NULL`
    result, err := database.DB.ExecContext(ctx, query, pekerjaanID, userID)

    if err != nil {
        return 0, err
    }
    return result.RowsAffected()
}



// hard delete buat user


func DeletePekerjaanTrashRestoreByUser(ctx context.Context, pekerjaanID int, userID int) (int64, error) {
    query := `
        Delete pekerjaan_alumni pa
        SET deleted_at = NULL, deleted_by = NULL
        FROM alumni a
        WHERE pa.alumni_id = a.id
          AND pa.id = $1
          AND a.user_id = $2
          AND pa.deleted_at IS NOT NULL`
    result, err := database.DB.ExecContext(ctx, query, pekerjaanID, userID)

    if err != nil {
        return 0, err
    }
    return result.RowsAffected()
}
