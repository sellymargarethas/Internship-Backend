package paketrepository

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/repositories"
	. "Internship-Backend/utils"
	"database/sql"
	"log"
	"time"
)

type paketRepository struct {
	RepoDB repositories.Repository
}

func NewPaketRepository(repoDB repositories.Repository) paketRepository {
	return paketRepository{
		RepoDB: repoDB,
	}
}

const defineColumn = `uraian`

func (ctx paketRepository) IsPaketExistsByIndex(paket models.Paket) (models.Paket, bool, error) {
	var args []interface{}
	var result models.Paket

	query := `
	SELECT paket.id, paket.` + defineColumn + `, idcorporate, corporate.uraian
	FROM paket
	INNER JOIN corporate ON paket.idcorporate = corporate.id
	WHERE paket.deleted_at IS NULL`

	if paket.Id != constants.EMPTY_VALUE_INT {
		query += ` AND paket.id = ? `
		args = append(args, paket.Id)
	}

	if paket.Uraian != constants.EMPTY_VALUE {
		query += ` AND paket.uraian = ? `
		args = append(args, paket.Uraian)
	}

	query += ` ORDER BY id ASC`

	query = ReplaceSQL(query, "?")

	rows, err := ctx.RepoDB.DB.Query(query, args...)

	if err != nil {
		return result, constants.FALSE_VALUE, err
	}
	defer rows.Close()

	data, err := PaketDto(rows)
	if err != nil {
		return result, constants.FALSE_VALUE, err
	}

	if len(data) == constants.EMPTY_VALUE_INT {
		return result, constants.FALSE_VALUE, err
	}
	return data[0], constants.TRUE_VALUE, err
}

// Get All Paket
func (ctx paketRepository) GetListPaket() ([]models.Paket, error) {
	var result []models.Paket

	query := `SELECT paket.id, paket.` + defineColumn + `, idcorporate, corporate.uraian
	FROM paket
	INNER JOIN corporate ON paket.idcorporate = corporate.id
	WHERE paket.deleted_at IS NULL AND corporate.deleted_at IS NULL
	ORDER BY paket.id ASC`

	rows, err := ctx.RepoDB.DB.Query(query)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	data, err := PaketDto(rows)

	if err != nil {
		log.Println("Error querying GetPaket : ", err)
	}

	return data, err
}

// Insert Paket
func (ctx paketRepository) InsertPaket(paket models.RequestAddPaket) (id int, err error) {
	query := `
	INSERT INTO paket
	(` + defineColumn + `, idcorporate, created_at) 
	VALUES 
	(?, ?, ?)
	RETURNING id`
	query = ReplaceSQL(query, "?")
	err = ctx.RepoDB.DB.QueryRow(query, &paket.Uraian, &paket.IdCorporate, time.Now()).Scan(&id)

	if err != nil {
		log.Println("Error querying InsertPaket : ", err)
	}

	return
}

// Update Paket
func (ctx paketRepository) UpdatePaket(paket models.RequestUpdatePaket) (id int, err error) {
	query := `
	UPDATE paket SET 
		uraian=?, idcorporate=?, updated_at=?
	WHERE id=? 
	RETURNING id
	`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, paket.Uraian, paket.IdCorporate, time.Now(), paket.Id).Scan(&id)

	if err != nil {
		log.Println("Error querying UpdatePaket : ", err)
	}

	return
}

// Delete Paket
func (ctx paketRepository) DeletePaket(paket models.RequestDeletePaket) (id int, err error) {
	query := `
		UPDATE paket SET deleted_at=? WHERE id=? RETURNING id
	`

	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, time.Now(), paket.Id).Scan(&id)

	if err != nil {
		log.Println("Error querying DeletePaket : ", err)
	}

	return
}

func PaketDto(rows *sql.Rows) (result []models.Paket, err error) {
	for rows.Next() {
		var paket models.Paket
		err = rows.Scan(&paket.Id, &paket.Uraian, &paket.IdCorporate, &paket.NamaCorporate)
		if err != nil {
			return
		}
		result = append(result, paket)
	}

	return
}
