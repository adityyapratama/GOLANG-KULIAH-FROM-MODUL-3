package repository

import (
	"context"
	"database/sql"
	"fmt"
	"golang-kuliah-from-modul-3/app/model"
	"golang-kuliah-from-modul-3/database"
	"log"
)

func UserLogin(ctx context.Context, userLogin string) (*model.User, string, error) {
	var user model.User
	var passwordHash string

	err := database.DB.QueryRowContext(ctx, `
		SELECT id, username, email, password_hash, role, created_at
		FROM users
		WHERE username = $1 OR email = $1
		LIMIT 1
	`, userLogin).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&passwordHash,
		&user.Role,
		&user.CreatedAt,
		
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, "", sql.ErrNoRows
		}
		return nil, "", err
	}

	return &user, passwordHash, nil
}


func CreateUser(ctx context.Context, user *model.User, passwordHash string) error{

	err := database.DB.QueryRowContext(ctx,
	`INSERT INTO users (username, email,password_hash,role)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at`,
		user.Username, user.Email, passwordHash,user.Role).Scan(&user.ID, &user.CreatedAt)

		return  err

}


func GetUserRepo(search, sortBy, order string, limit, offset int) ([]model.User, error) { 
    query := fmt.Sprintf(`
        SELECT id, name, email, created_at ,pekerjaan_alumni_id
        FROM users
        WHERE name ILIKE $1 OR email ILIKE $1
        ORDER BY %s %s
        LIMIT $2 OFFSET $3`, sortBy, order)

    rows, err := database.DB.Query(query, "%"+search+"%", limit, offset)
    
    if err != nil {
        log.Println("Query error :", err)
        return nil, err
    }
    defer rows.Close()
    
    var users []model.User
    
    for rows.Next() {
        var u model.User
    
        if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.CreatedAt); err != nil {
            return nil, err
        }
        users = append(users, u)
    }

    return users, nil
}


func CountUsersRepo(search string) (int, error) {
			var total int
			countQuery := `SELECT COUNT(*) FROM users WHERE name ILIKE $1 OR
			email ILIKE $1`
		err := database.DB.QueryRow(countQuery, "%"+search+"%").Scan(&total)
		if err != nil && err != sql.ErrNoRows {
			return 0, err
	}
			return total, nil
}


func DeleteUsers(ctx context.Context, id int) (int64, error) {
    result, err := database.DB.ExecContext(ctx, "UPDATE users set deleted_at= NOW() WHERE id = $1 AND deleted_at IS NULL", id)
    if err != nil {
        return 0, err
    }
    return result.RowsAffected()
}




func GetUserPekerjaan(ctx context.Context, userID int) ([]model.Pekerjaan, error) {
    query := `
        SELECT p.id, p.alumni_id, p.nama_perusahaan, p.posisi_jabatan, p.bidang_industri,
               p.lokasi_kerja, p.gaji_range, p.tanggal_mulai_kerja, p.tanggal_selesai_kerja,
               p.status_pekerjaan, p.deskripsi_pekerjaan, p.created_at, p.updated_at
        FROM pekerjaan_alumni p
        LEFT JOIN alumni a ON p.alumni_id = a.id
        WHERE a.user_id = $1 AND p.deleted_at IS NULL
    `
    rows, err := database.DB.QueryContext(ctx, query, userID)
    if err != nil {
        return nil, err
    }
    
    defer rows.Close()

    var listPekerjaan []model.Pekerjaan


    for rows.Next() {
        
        var p model.Pekerjaan

        
        if err := rows.Scan(
            &p.ID,
            &p.AlumniID,
            &p.NamaPerusahaan,
            &p.PosisiJabatan,
            &p.BidangIndustri,
            &p.LokasiKerja,
            &p.GajiRange,
            &p.TanggalMulaiKerja,
            &p.TanggalSelesaiKerja,
            &p.StatusPekerjaan,
            &p.DeskripsiPekerjaan,
            &p.CreatedAt,
            &p.UpdatedAt,
        ); err != nil {
            return nil, err
        }

        
        listPekerjaan = append(listPekerjaan, p)
    }

    return listPekerjaan, nil
}



func DeletePekerjaanByUser(ctx context.Context, pekerjaanID int, userID int) (int64, error) {
	query := `
		UPDATE pekerjaan_alumni pa
		SET deleted_at = NOW(), deleted_by = $1
		FROM alumni a
		WHERE pa.alumni_id = a.id
		  AND pa.id = $2
		  AND a.user_id = $1
		  AND pa.deleted_at IS NULL`
	result, err := database.DB.ExecContext(ctx, query, userID, pekerjaanID)

	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
