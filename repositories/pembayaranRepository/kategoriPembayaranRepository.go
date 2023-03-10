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

type kategoripembayaranRepository struct {
	RepoDB repositories.Repository
}

func NewKategoriPembayaranRepository(repoDB repositories.Repository) kategoripembayaranRepository {
	return kategoripembayaranRepository{
		RepoDB: repoDB,
	}
}

const defineColumn = `uraian`

func (ctx kategoripembayaranRepository) IsKategoriPembayaranExistsByIndex(kategoripembayaran models.KategoriPembayaran) (models.KategoriPembayaran, bool, error) {
	var args []interface{}
	var result models.KategoriPembayaran

	query := `
	SELECT id, ` + defineColumn + `
	FROM kategoripembayaran
	WHERE deleted_at IS NULL`

	if kategoripembayaran.Id != constants.EMPTY_VALUE_INT {
		query += ` AND kategoripembayaran.id = ? `
		args = append(args, kategoripembayaran.Id)
	}

	if kategoripembayaran.Uraian != constants.EMPTY_VALUE {
		query += ` AND kategoripembayaran.uraian = ? `
		args = append(args, kategoripembayaran.Uraian)
	}

	query += ` ORDER BY id ASC`

	query = ReplaceSQL(query, "?")

	rows, err := ctx.RepoDB.DB.Query(query, args...)

	if err != nil {
		return result, constants.FALSE_VALUE, err
	}
	defer rows.Close()

	data, err := KategoriPembayaranDto(rows)
	if err != nil {
		return result, constants.FALSE_VALUE, err
	}

	if len(data) == constants.EMPTY_VALUE_INT {
		return result, constants.FALSE_VALUE, err
	}
	return data[0], constants.TRUE_VALUE, err
}

// Get All Kategori Pembayaran
func (ctx kategoripembayaranRepository) GetListKategoriPembayaran() ([]models.KategoriPembayaran, error) {
	var result []models.KategoriPembayaran

	query := `SELECT id, ` + defineColumn + `
	FROM kategoripembayaran
	WHERE deleted_at IS NULL
	ORDER BY id ASC`

	rows, err := ctx.RepoDB.DB.Query(query)

	if err != nil {
		return result, err
	}
	defer rows.Close()

	data, err := KategoriPembayaranDto(rows)

	if err != nil {
		log.Println("Error querying GetKategoriPembayaran : ", err)
	}
	return data, err
}

// Insert Kategori Pembayaran
func (ctx kategoripembayaranRepository) InsertKategoriPembayaran(kategoripembayaran models.RequestAddKategoriPembayaran) (id int, err error) {

	query := `INSERT INTO kategoripembayaran (` + defineColumn + `, created_at) 
	VALUES (?, ?) RETURNING id`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, &kategoripembayaran.Uraian, time.Now()).Scan(&id)

	if err != nil {
		log.Println("Error querying InsertKategoriPembayaran : ", err)
	}

	return
}

// Update Kategori Pembayaran
func (ctx kategoripembayaranRepository) UpdateKategoriPembayaran(pembayaran models.RequestUpdateKategoriPembayaran) (id int, err error) {
	query := `
	UPDATE kategoripembayaran SET 
		uraian=?, updated_at=?
	WHERE id=? RETURNING id`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, pembayaran.Uraian, time.Now(), pembayaran.Id).Scan(&id)

	if err != nil {
		log.Println("Error querying UpdateKategoriPembayaran : ", err)
	}

	return
}

// Delete Kategori Pembayaran
func (ctx kategoripembayaranRepository) DeleteKategoriPembayaran(pembayaran models.RequestDeleteKategoriPembayaran) (id int, err error) {
	query := `UPDATE kategoripembayaran SET deleted_at=? WHERE id=? RETURNING id`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, time.Now(), pembayaran.Id).Scan(&id)

	if err != nil {
		log.Println("Error querying DeleteKategoriPembayaran : ", err)
	}

	return
}

func KategoriPembayaranDto(rows *sql.Rows) (result []models.KategoriPembayaran, err error) {
	for rows.Next() {
		var pembayaran models.KategoriPembayaran
		err = rows.Scan(&pembayaran.Id, &pembayaran.Uraian)
		if err != nil {
			return
		}
		result = append(result, pembayaran)
	}

	return
}

func (ctx kategoripembayaranRepository) GetListKategoriPembayaranId() (result []models.KategoriPembayaranId, err error) {
	query := `SELECT id from kategoripembayaran`
	rows, err := ctx.RepoDB.DB.Query(query)
	for rows.Next() {
		var data models.KategoriPembayaranId
		rows.Scan(&data.IdKategoriPembayaran)
		result = append(result, data)
	}
	return
}
