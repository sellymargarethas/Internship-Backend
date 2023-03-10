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

type kategoriprodukRepository struct {
	RepoDB repositories.Repository
}

func NewKategoriProdukRepository(repoDB repositories.Repository) kategoriprodukRepository {
	return kategoriprodukRepository{
		RepoDB: repoDB,
	}
}

const defineColumn2 = `uraian`

func (ctx kategoriprodukRepository) IsKategoriProdukExistsByIndex(kategoriproduk models.KategoriProduk) (models.KategoriProduk, bool, error) {
	var args []interface{}
	var result models.KategoriProduk

	query := `
	SELECT id, ` + defineColumn2 + `, (serviceFee::money::numeric::integer)
	FROM kategoriproduk
	WHERE deleted_at IS NULL`

	if kategoriproduk.Id != constants.EMPTY_VALUE_INT {
		query += ` AND kategoriproduk.id = ? `
		args = append(args, kategoriproduk.Id)
	}

	if kategoriproduk.Uraian != constants.EMPTY_VALUE {
		query += ` AND kategoriproduk.uraian = ? `
		args = append(args, kategoriproduk.Uraian)
	}

	query += ` ORDER BY id ASC`

	query = ReplaceSQL(query, "?")

	rows, err := ctx.RepoDB.DB.Query(query, args...)

	if err != nil {
		return result, constants.FALSE_VALUE, err
	}
	defer rows.Close()

	data, err := KategoriProdukDto(rows)
	if err != nil {
		return result, constants.FALSE_VALUE, err
	}

	if len(data) == constants.EMPTY_VALUE_INT {
		return result, constants.FALSE_VALUE, err
	}
	return data[0], constants.TRUE_VALUE, err
}

// Get All Kategori Produk
func (ctx kategoriprodukRepository) GetListKategoriProduk() ([]models.KategoriProduk, error) {
	var result []models.KategoriProduk

	query := `SELECT id, ` + defineColumn2 + `, (serviceFee::money::numeric::integer)
	FROM kategoriproduk
	WHERE deleted_at IS NULL
	ORDER BY id ASC`
	query = ReplaceSQL(query, "?")

	rows, err := ctx.RepoDB.DB.Query(query)

	if err != nil {
		return result, err
	}
	defer rows.Close()

	data, err := KategoriProdukDto(rows)

	if err != nil {
		log.Println("Error querying GetKategoriProduk : ", err)
	}

	return data, err
}

// Insert Kategori Produk
func (ctx kategoriprodukRepository) InsertKategoriProduk(kategoriproduk models.RequestAddKategoriProduk) (id int, err error) {
	query := `
	INSERT INTO kategoriproduk
	(` + defineColumn2 + `, serviceFee, created_at) 
	VALUES 
	(?, ?, ?) returning id`

	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, &kategoriproduk.Uraian, &kategoriproduk.ServiceFee, time.Now()).Scan(&id)

	if err != nil {
		log.Println("Error querying InsertKategoriProduk : ", err)
	}

	return
}

// Update Kategori Produk
func (ctx kategoriprodukRepository) UpdateKategoriProduk(kategoriproduk models.RequestUpdateKategoriProduk) (id int, err error) {
	query := `
	UPDATE kategoriproduk SET 
		uraian=?, serviceFee=?, updated_at=?
	WHERE id=? RETURNING id
	`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, kategoriproduk.Uraian, kategoriproduk.ServiceFee, time.Now(), kategoriproduk.Id).Scan(&id)

	if err != nil {
		log.Println("Error querying UpdateKategoriProduk : ", err)
	}

	return
}

// Delete Kategori Produk
func (ctx kategoriprodukRepository) DeleteKategoriProduk(kategoriproduk models.RequestDeleteKategoriProduk) (id int, err error) {
	query := `
		UPDATE kategoriproduk SET deleted_at=? WHERE id=? RETURNING id
	`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, time.Now(), kategoriproduk.Id).Scan(&id)

	if err != nil {
		log.Println("Error querying DeleteKategoriProduk : ", err)
	}

	return
}

func KategoriProdukDto(rows *sql.Rows) (result []models.KategoriProduk, err error) {
	for rows.Next() {
		var kategoriproduk models.KategoriProduk
		err = rows.Scan(&kategoriproduk.Id, &kategoriproduk.Uraian, &kategoriproduk.ServiceFee)
		if err != nil {
			return
		}
		result = append(result, kategoriproduk)
	}

	return
}
