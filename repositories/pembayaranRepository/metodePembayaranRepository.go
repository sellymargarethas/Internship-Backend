package pembayaranrepository

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/repositories"
	. "Internship-Backend/utils"
	"database/sql"
	"log"
	"time"
)

type metodepembayaranRepository struct {
	RepoDB repositories.Repository
}

func NewMetodePembayaranRepository(repoDB repositories.Repository) metodepembayaranRepository {
	return metodepembayaranRepository{
		RepoDB: repoDB,
	}
}

func (ctx metodepembayaranRepository) IsMetodePembayaranExistsByIndex(metodepembayaran models.MetodePembayaran) (models.MetodePembayaran, bool, error) {
	var args []interface{}
	var result models.MetodePembayaran

	query := `
	SELECT metodepembayaran.id, metodepembayaran.uraian, idkategoripembayaran, kategoripembayaran.uraian
	FROM metodepembayaran
	INNER JOIN kategoripembayaran on kategoripembayaran.id=metodepembayaran.idkategoripembayaran
	WHERE metodepembayaran.deleted_at IS NULL`

	if metodepembayaran.Id != constants.EMPTY_VALUE_INT {
		query += ` AND metodepembayaran.id = ? `
		args = append(args, metodepembayaran.Id)
	}

	if metodepembayaran.Uraian != constants.EMPTY_VALUE {
		query += ` AND metodepembayaran.uraian = ? `
		args = append(args, metodepembayaran.Uraian)
	}

	query += ` ORDER BY id ASC`

	query = ReplaceSQL(query, "?")

	rows, err := ctx.RepoDB.DB.Query(query, args...)

	if err != nil {
		return result, constants.FALSE_VALUE, err
	}
	defer rows.Close()

	data, err := MetodePembayaranDto(rows)
	if err != nil {
		return result, constants.FALSE_VALUE, err
	}

	if len(data) == constants.EMPTY_VALUE_INT {
		return result, constants.FALSE_VALUE, err
	}
	return data[0], constants.TRUE_VALUE, err
}

// Get All Metode Pembayaran
func (ctx metodepembayaranRepository) GetListMetodePembayaran() ([]models.MetodePembayaran, error) {
	var result []models.MetodePembayaran

	query := `SELECT metodepembayaran.id, metodepembayaran.uraian, idkategoripembayaran, kategoripembayaran.uraian
	FROM metodepembayaran
	INNER JOIN kategoripembayaran on kategoripembayaran.id=metodepembayaran.idkategoripembayaran
	WHERE metodepembayaran.deleted_at IS NULL AND kategoripembayaran.deleted_at IS NULL
	ORDER BY metodepembayaran.id ASC`

	rows, err := ctx.RepoDB.DB.Query(query)

	if err != nil {
		return result, err
	}
	defer rows.Close()

	data, err := MetodePembayaranDto(rows)

	if err != nil {
		log.Println("Error querying GetMetodePembayaran : ", err)
	}
	return data, err
}

// Insert Metode Pembayaran
func (ctx metodepembayaranRepository) InsertMetodePembayaran(metodepembayaran models.RequestAddMetodePembayaran) (id int, err error) {
	query := `
	INSERT INTO metodepembayaran (uraian, idkategoripembayaran, created_at) 
	VALUES (?, ?, ?) RETURNING id`

	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, &metodepembayaran.Uraian, metodepembayaran.IdKategoriPembayaran, time.Now()).Scan(&id)

	if err != nil {
		log.Println("Error querying InsertMetodePembayaran : ", err)
	}

	return
}

// Update Metode Pembayaran
func (ctx metodepembayaranRepository) UpdateMetodePembayaran(pembayaran models.RequestUpdateMetodePembayaran) (id int, err error) {
	query := `
	UPDATE metodepembayaran SET 
		uraian=?, idkategoripembayaran=?, updated_at=?
	WHERE id=? RETURNING id
	`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, pembayaran.Uraian, pembayaran.IdKategoriPembayaran, time.Now(), pembayaran.Id).Scan(&id)

	if err != nil {
		log.Println("Error querying UpdateMetodePembayaran : ", err)
	}

	return
}

// Delete Metode Pembayaran
func (ctx metodepembayaranRepository) DeleteMetodePembayaran(pembayaran models.RequestDeleteMetodePembayaran) (id int, err error) {
	query := `
		UPDATE metodepembayaran SET deleted_at=? WHERE id=? RETURNING id
	`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, time.Now(), pembayaran.Id).Scan(&id)

	if err != nil {
		log.Println("Error querying DeleteMetodePembayaran : ", err)
	}

	return
}

func MetodePembayaranDto(rows *sql.Rows) (result []models.MetodePembayaran, err error) {
	for rows.Next() {
		var pembayaran models.MetodePembayaran
		err = rows.Scan(&pembayaran.Id, &pembayaran.Uraian, &pembayaran.IdKategoriPembayaran, &pembayaran.NamaKategoriPembayaran)
		if err != nil {
			return
		}
		result = append(result, pembayaran)
	}

	return
}
