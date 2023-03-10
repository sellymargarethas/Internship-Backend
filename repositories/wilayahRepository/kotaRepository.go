package wilayahRepository

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/repositories"
	. "Internship-Backend/utils"
	"database/sql"
	"log"
	"time"
)

type kotaRepository struct {
	RepoDB repositories.Repository
}

func NewKotaRepository(repoDB repositories.Repository) kotaRepository {
	return kotaRepository{
		RepoDB: repoDB,
	}
}

const defineColumn2 = `uraian`

func (ctx kotaRepository) IsKotaExistsByIndex(kota models.Kota) (models.Kota, bool, error) {
	var args []interface{}
	var result models.Kota

	query := `
	SELECT kota.id, kota.uraian, idprovinsi, provinsi.uraian
	FROM kota
	INNER JOIN provinsi on kota.idprovinsi=provinsi.id
	WHERE kota.deleted_at IS NULL`

	if kota.Id != constants.EMPTY_VALUE_INT {
		query += ` AND kota.id = ? `
		args = append(args, kota.Id)
	}

	if kota.Uraian != constants.EMPTY_VALUE {
		query += ` AND kota.uraian = ? `
		args = append(args, kota.Uraian)
	}

	query += ` ORDER BY id ASC`

	query = ReplaceSQL(query, "?")

	rows, err := ctx.RepoDB.DB.Query(query, args...)

	if err != nil {
		return result, constants.FALSE_VALUE, err
	}
	defer rows.Close()

	data, err := KotaDto(rows)
	if err != nil {
		return result, constants.FALSE_VALUE, err
	}

	if len(data) == constants.EMPTY_VALUE_INT {
		return result, constants.FALSE_VALUE, err
	}
	return data[0], constants.TRUE_VALUE, err
}

// Get All Kota
func (ctx kotaRepository) GetListKota() ([]models.Kota, error) {
	var result []models.Kota

	query := `SELECT kota.id, kota.uraian, idprovinsi, provinsi.uraian
	FROM kota
	INNER JOIN provinsi on kota.idprovinsi=provinsi.id
	WHERE kota.deleted_at IS NULL
	ORDER BY kota.id ASC`

	rows, err := ctx.RepoDB.DB.Query(query)

	if err != nil {
		return result, err
	}
	defer rows.Close()

	data, err := KotaDto(rows)

	if err != nil {
		log.Println("Error querying GetKota : ", err)
	}

	return data, err
}

// Insert Kota
func (ctx kotaRepository) InsertKota(kota models.RequestAddKota) (id int, err error) {
	query := `
	INSERT INTO kota (` + defineColumn2 + `, idprovinsi, created_at) 
	VALUES (?, ?, ?) returning id`
	query = ReplaceSQL(query, "?")
	err = ctx.RepoDB.DB.QueryRow(query, &kota.Uraian, kota.IdProvinsi, time.Now()).Scan(&id)

	if err != nil {
		log.Println("Error querying InsertKota : ", err)
	}

	return
}

// Update Kota
func (ctx kotaRepository) UpdateKota(kota models.RequestUpdateKota) (id int, err error) {
	query := `
	UPDATE kota SET 
		uraian=?, idprovinsi=?, updated_at=?
	WHERE id=? RETURNING id
	`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, kota.Uraian, kota.IdProvinsi, time.Now(), kota.Id).Scan(&id)

	if err != nil {
		log.Println("Error querying UpdateKota : ", err)
	}

	return
}

// Delete Kota
func (ctx kotaRepository) DeleteKota(kota models.RequestDeleteKota) (id int, err error) {
	query := `
		UPDATE kota SET deleted_at=? WHERE id=? RETURNING id
	`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, time.Now(), kota.Id).Scan(&id)

	if err != nil {
		log.Println("Error querying DeleteKota : ", err)
	}

	return
}

func KotaDto(rows *sql.Rows) (result []models.Kota, err error) {
	for rows.Next() {
		var kota models.Kota
		err = rows.Scan(&kota.Id, &kota.Uraian, &kota.IdProvinsi, &kota.NamaProvinsi)
		if err != nil {
			return
		}
		result = append(result, kota)
	}

	return
}
