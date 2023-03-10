package produkrepository

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/repositories"
	. "Internship-Backend/utils"
	"database/sql"
	"log"
	"time"
)

type satuanprodukRepository struct {
	RepoDB repositories.Repository
}

func NewSatuanProdukRepository(repoDB repositories.Repository) satuanprodukRepository {
	return satuanprodukRepository{
		RepoDB: repoDB,
	}
}

const defineColumn3 = `uraian`

func (ctx satuanprodukRepository) IsSatuanProdukExistsByIndex(satuanproduk models.SatuanProduk) (models.SatuanProduk, bool, error) {
	var args []interface{}
	var result models.SatuanProduk

	query := `
	SELECT id, ` + defineColumn3 + `
	FROM satuanproduk
	WHERE deleted_at IS NULL`

	if satuanproduk.Id != constants.EMPTY_VALUE_INT {
		query += ` AND satuanproduk.id = ? `
		args = append(args, satuanproduk.Id)
	}

	if satuanproduk.Uraian != constants.EMPTY_VALUE {
		query += ` AND satuanproduk.uraian = ? `
		args = append(args, satuanproduk.Uraian)
	}

	query += ` ORDER BY id ASC`

	query = ReplaceSQL(query, "?")

	rows, err := ctx.RepoDB.DB.Query(query, args...)

	if err != nil {
		return result, constants.FALSE_VALUE, err
	}
	defer rows.Close()

	data, err := SatuanProdukDto(rows)
	if err != nil {
		return result, constants.FALSE_VALUE, err
	}

	if len(data) == constants.EMPTY_VALUE_INT {
		return result, constants.FALSE_VALUE, err
	}
	return data[0], constants.TRUE_VALUE, err
}

// Get All Satuan Produk
func (ctx satuanprodukRepository) GetListSatuanProduk() ([]models.SatuanProduk, error) {
	var result []models.SatuanProduk

	query := `SELECT id, ` + defineColumn3 + `
	FROM satuanproduk
	WHERE deleted_at IS NULL
	ORDER BY id ASC`

	rows, err := ctx.RepoDB.DB.Query(query)

	if err != nil {
		return result, err
	}
	defer rows.Close()

	data, err := SatuanProdukDto(rows)

	if err != nil {
		log.Println("Error querying GetSatuanProduk : ", err)
	}

	return data, err
}

// Insert Satuan Produk
func (ctx satuanprodukRepository) InsertSatuanProduk(satuanproduk models.RequestAddSatuanProduk) (id int, err error) {
	query := `
	INSERT INTO satuanproduk
	(` + defineColumn3 + `, created_at) 
	VALUES 
	(?, ?) returning id`
	query = ReplaceSQL(query, "?")
	err = ctx.RepoDB.DB.QueryRow(query, &satuanproduk.Uraian, time.Now()).Scan(&id)

	if err != nil {
		log.Println("Error querying InsertSatuanProduk : ", err)
	}

	return
}

// Update Satuan Produk
func (ctx satuanprodukRepository) UpdateSatuanProduk(satuanproduk models.RequestUpdateSatuanProduk) (id int, err error) {
	query := `
	UPDATE satuanproduk SET 
		uraian=?, updated_at=?
	WHERE id=? RETURNING id
	`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, satuanproduk.Uraian, time.Now(), satuanproduk.Id).Scan(&id)

	if err != nil {
		log.Println("Error querying UpdateSatuanProduk : ", err)
	}

	return
}

// Delete Satuan Produk
func (ctx satuanprodukRepository) DeleteSatuanProduk(satuanproduk models.RequestDeleteSatuanProduk) (id int, err error) {
	query := `
		UPDATE satuanproduk SET deleted_at=? WHERE id=? RETURNING id
	`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, time.Now(), satuanproduk.Id).Scan(&id)

	if err != nil {
		log.Println("Error querying DeleteSatuanProduk : ", err)
	}

	return
}

func SatuanProdukDto(rows *sql.Rows) ([]models.SatuanProduk, error) {
	var result []models.SatuanProduk
	for rows.Next() {
		var satuanproduk models.SatuanProduk
		err := rows.Scan(&satuanproduk.Id, &satuanproduk.Uraian)
		if err != nil {
			return result, err
		}
		result = append(result, satuanproduk)
	}

	return result, nil
}
