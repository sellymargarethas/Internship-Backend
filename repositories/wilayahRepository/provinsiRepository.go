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

type provinsiRepository struct {
	RepoDB repositories.Repository
}

func NewProvinsiRepository(repoDB repositories.Repository) provinsiRepository {
	return provinsiRepository{
		RepoDB: repoDB,
	}
}

const defineColumn = `uraian`

func (ctx provinsiRepository) IsProvinsiExistsByIndex(provinsi models.Provinsi) (models.Provinsi, bool, error) {
	var args []interface{}
	var result models.Provinsi

	query := `
	SELECT id, ` + defineColumn + `
	FROM provinsi
	WHERE deleted_at IS NULL`

	if provinsi.Id != constants.EMPTY_VALUE_INT {
		query += ` AND provinsi.id = ? `
		args = append(args, provinsi.Id)
	}

	if provinsi.Uraian != constants.EMPTY_VALUE {
		query += ` AND provinsi.uraian = ? `
		args = append(args, provinsi.Uraian)
	}

	query += ` ORDER BY id ASC`

	query = ReplaceSQL(query, "?")

	rows, err := ctx.RepoDB.DB.Query(query, args...)

	if err != nil {
		return result, constants.FALSE_VALUE, err
	}
	defer rows.Close()

	data, err := ProvinsiDto(rows)
	if err != nil {
		return result, constants.FALSE_VALUE, err
	}

	if len(data) == constants.EMPTY_VALUE_INT {
		return result, constants.FALSE_VALUE, err
	}
	return data[0], constants.TRUE_VALUE, err
}

// Get All Provinsi
func (ctx provinsiRepository) GetListProvinsi() ([]models.Provinsi, error) {
	var result []models.Provinsi

	query := `SELECT id, ` + defineColumn + `
	FROM provinsi
	WHERE deleted_at IS NULL
	ORDER BY id ASC`

	rows, err := ctx.RepoDB.DB.Query(query)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	data, err := ProvinsiDto(rows)

	if err != nil {
		log.Println("Error querying GetProvinsi : ", err)
	}

	return data, err
}

// Insert Provinsi
func (ctx provinsiRepository) InsertProvinsi(provinsi models.RequestAddProvinsi) (id int, err error) {
	query := `INSERT INTO provinsi (` + defineColumn + `, created_at) 
	VALUES (?, ?) RETURNING id`
	query = ReplaceSQL(query, "?")
	err = ctx.RepoDB.DB.QueryRow(query, &provinsi.Uraian, time.Now()).Scan(&id)

	if err != nil {
		log.Println("Error querying InsertProvinsi : ", err)
	}

	return
}

// Update Provinsi
func (ctx provinsiRepository) UpdateProvinsi(provinsi models.RequestUpdateProvinsi) (id int, err error) {
	query := `
	UPDATE provinsi SET 
		uraian=?, updated_at=?
	WHERE id=? RETURNING id
	`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, provinsi.Uraian, time.Now(), provinsi.Id).Scan(&id)

	if err != nil {
		log.Println("Error querying UpdateProvinsi : ", err)
	}

	return
}

// Delete Provinsi
func (ctx provinsiRepository) DeleteProvinsi(pembayaran models.RequestDeleteProvinsi) (id int, err error) {
	query := `
		UPDATE provinsi SET deleted_at=? WHERE id=? RETURNING id
	`
	query = ReplaceSQL(query, "?")
	err = ctx.RepoDB.DB.QueryRow(query, time.Now(), pembayaran.Id).Scan(&id)

	if err != nil {
		log.Println("Error querying DeleteProvinsi : ", err)
	}

	return
}

func ProvinsiDto(rows *sql.Rows) (result []models.Provinsi, err error) {
	for rows.Next() {
		var provinsi models.Provinsi
		err = rows.Scan(&provinsi.Id, &provinsi.Uraian)
		if err != nil {
			return
		}
		result = append(result, provinsi)
	}

	return
}
