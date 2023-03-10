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

type produkpaketRepository struct {
	RepoDB repositories.Repository
}

func NewProdukPaketRepository(repoDB repositories.Repository) produkpaketRepository {
	return produkpaketRepository{
		RepoDB: repoDB,
	}
}

func (ctx produkpaketRepository) IsProdukPaketExistsByIndex(produkpaket models.ProdukPaket) (models.ProdukPaket, bool, error) {
	var args []interface{}
	var result models.ProdukPaket

	query := `
	SELECT produkpaket.id, produkpaket.idpaket, paket.uraian, produkpaket.idproduk, produkpaket.namaproduk, (produkpaket.hargajual::money::numeric::integer)
	FROM produkpaket
	INNER JOIN paket ON paket.id = produkpaket.idpaket
	WHERE produkpaket.deleted_at IS NULL`

	if produkpaket.Id != constants.EMPTY_VALUE_INT {
		query += ` AND produkpaket.id = ? `
		args = append(args, produkpaket.Id)
	}

	query += ` ORDER BY id ASC`

	query = ReplaceSQL(query, "?")

	rows, err := ctx.RepoDB.DB.Query(query, args...)

	if err != nil {
		return result, constants.FALSE_VALUE, err
	}
	defer rows.Close()

	data, err := ProdukPaketDto(rows)
	if err != nil {
		return result, constants.FALSE_VALUE, err
	}

	if len(data) == constants.EMPTY_VALUE_INT {
		return result, constants.FALSE_VALUE, err
	}
	return data[0], constants.TRUE_VALUE, err
}

// Get All Produk Paket
func (ctx produkpaketRepository) GetListProdukPaket() ([]models.ProdukPaket, error) {
	var result []models.ProdukPaket

	query := `SELECT produkpaket.id, produkpaket.idpaket, paket.uraian, produkpaket.idproduk, produkpaket.namaproduk, (produkpaket.hargajual::money::numeric::integer)
	FROM produkpaket
	INNER JOIN paket ON paket.id = produkpaket.idpaket
	WHERE produkpaket.deleted_at IS NULL AND paket.deleted_at IS NULL
	ORDER BY produkpaket.id ASC`

	rows, err := ctx.RepoDB.DB.Query(query)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	data, _ := ProdukPaketDto(rows)

	if err != nil {
		log.Println("Error querying GetProdukPaket : ", err)
	}

	return data, err
}

func (ctx produkpaketRepository) GetCountList() (count int, err error) {
	query := `SELECT COUNT(produkpaket.id)
	FROM produkpaket
	INNER JOIN paket ON paket.id = produkpaket.idpaket
	WHERE produkpaket.deleted_at IS NULL AND paket.deleted_at IS NULL`

	err = ctx.RepoDB.DB.QueryRow(query).Scan(&count)
	if err != nil {
		log.Println("Error querying GetCountList : ", err)
	}

	return
}

// Insert Produk Paket
func (ctx produkpaketRepository) InsertProdukPaket(produkpaket models.RequestAddProdukPaket) (id int, err error) {
	query := `
	INSERT INTO produkpaket
	(idpaket, idproduk, namaproduk, hargajual, created_at) 
	VALUES 
	(?, ?, (SELECT nama FROM produk WHERE id=?), (SELECT hargajual FROM produk WHERE id=?), ?) returning id`
	query = ReplaceSQL(query, "?")
	err = ctx.RepoDB.DB.QueryRow(query, &produkpaket.IdPaket, &produkpaket.IdProduk, &produkpaket.IdProduk, &produkpaket.IdProduk, time.Now()).Scan(&id)

	if err != nil {
		log.Println("Error querying InsertProdukPaket : ", err)
	}

	return
}

// Update Produk Paket
func (ctx produkpaketRepository) UpdateProdukPaket(produkpaket models.RequestUpdateProdukPaket) (id int, err error) {
	query := `
	UPDATE produkpaket SET 
		idpaket=?, idproduk=?, namaproduk=(SELECT nama FROM produk WHERE id=?), hargajual=(SELECT hargajual FROM produk WHERE id=?), updated_at=?
	WHERE id=? RETURNING id
	`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, produkpaket.IdPaket, produkpaket.IdProduk, produkpaket.IdProduk, produkpaket.IdProduk, time.Now(), produkpaket.Id).Scan(&id)

	if err != nil {
		log.Println("Error querying UpdateProdukPaket : ", err)
	}

	return
}

// Delete Produk Paket
func (ctx produkpaketRepository) DeleteProdukPaket(produkpaket models.RequestDeleteProdukPaket) (id int, err error) {
	query := `
		UPDATE produkpaket SET deleted_at=? WHERE id=? RETURNING id
	`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, time.Now(), produkpaket.Id).Scan(&id)

	if err != nil {
		log.Println("Error querying DeleteProdukPaket : ", err)
	}

	return
}

func ProdukPaketDto(rows *sql.Rows) (result []models.ProdukPaket, err error) {
	for rows.Next() {
		var produkpaket models.ProdukPaket
		err = rows.Scan(&produkpaket.Id, &produkpaket.IdPaket, &produkpaket.NamaPaket, &produkpaket.IdProduk, &produkpaket.NamaProduk, &produkpaket.HargaJual)
		if err != nil {
			return
		}
		result = append(result, produkpaket)
	}

	return
}
