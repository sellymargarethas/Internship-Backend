package produkrepository

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/repositories"
	. "Internship-Backend/utils"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type produkRepository struct {
	RepoDB repositories.Repository
}

func NewProdukRepository(repoDB repositories.Repository) produkRepository {
	return produkRepository{
		RepoDB: repoDB,
	}
}

const defineColumn = `uraian`

func (ctx produkRepository) IsProdukExistsByIndex(produk models.Produk) (models.Produk, bool, error) {
	var args []interface{}
	var result models.Produk

	query := `
	SELECT produk.id, kodeproduk, nama, (hargajual::money::numeric::integer), produk.gambar, jenis, produk.idcorporate, corporate.uraian, produk.idsatuanproduk, satuanproduk.uraian, produk.idkategoriproduk, kategoriproduk.uraian, (hargarombongan::money::numeric::integer), minrombongan, statusstok, currentstok, statuspaket
	FROM produk
	INNER JOIN corporate ON produk.idcorporate=corporate.id
	INNER JOIN satuanproduk ON produk.idsatuanproduk=satuanproduk.id
	INNER JOIN kategoriproduk ON produk.idkategoriproduk=kategoriproduk.id
	WHERE produk.deleted_at IS NULL`

	if produk.Id != constants.EMPTY_VALUE_INT {
		query += ` AND produk.id = ?`
		args = append(args, produk.Id)
	}

	if produk.KodeProduk != constants.EMPTY_VALUE {
		query += ` AND produk.kodeproduk = ?`
		args = append(args, produk.KodeProduk)
	}

	if produk.Nama != constants.EMPTY_VALUE {
		query += ` AND produk.nama = ?`
		args = append(args, produk.Nama)
	}

	query += ` ORDER BY id ASC`

	query = ReplaceSQL(query, "?")

	rows, err := ctx.RepoDB.DB.Query(query, args...)

	if err != nil {
		return result, constants.FALSE_VALUE, err
	}
	defer rows.Close()

	data, err := ProdukDto(rows)
	if err != nil {
		return result, constants.FALSE_VALUE, err
	}

	if len(data) == constants.EMPTY_VALUE_INT {
		return result, constants.FALSE_VALUE, err
	}
	return data[0], constants.TRUE_VALUE, err
}

// Get All Produk
func (ctx produkRepository) GetListProduk(request models.RequestList) ([]models.Produk, error) {
	var args []interface{}

	query := `SELECT produk.id, kodeproduk, nama, (hargajual::money::numeric::integer), produk.gambar, jenis, produk.idcorporate, corporate.uraian, produk.idsatuanproduk, satuanproduk.uraian, produk.idkategoriproduk, kategoriproduk.uraian, (hargarombongan::money::numeric::integer), minrombongan, statusstok, currentstok, statuspaket
	FROM produk
	INNER JOIN corporate ON produk.idcorporate=corporate.id
	INNER JOIN satuanproduk ON produk.idsatuanproduk=satuanproduk.id
	INNER JOIN kategoriproduk ON produk.idkategoriproduk=kategoriproduk.id
	WHERE produk.deleted_at IS NULL AND corporate.deleted_at IS NULL AND satuanproduk.deleted_at IS NULL AND kategoriproduk.deleted_at IS NULL`

	if request.Keyword != "" {
		query += ` AND 
		(CAST(produk.id AS TEXT) ILIKE '%' || ? || '%' OR
		produk.kodeproduk ILIKE '%' || ? || '%' OR
		produk.nama ILIKE '%' || ? || '%' OR
		CAST(produk.hargajual AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(produk.hargarombongan AS TEXT) ILIKE '%' || ? || '%' OR	
		CAST(produk.minrombongan AS TEXT) ILIKE '%' || ? || '%' OR
		produk.gambar ILIKE '%' || ? || '%' OR
		CAST(produk.jenis AS TEXT) ILIKE '%' || ? || '%' OR
		corporate.uraian ILIKE '%' || ? || '%' OR
		satuanproduk.uraian ILIKE '%' || ? || '%' OR
		CAST(produk.statuspaket AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(produk.statusstok AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(produk.currentstok AS TEXT) ILIKE '%' || ? || '%')`
		args = append(args, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword)
	}

	// if request.Order == "DESC" {
	// 	query += " ORDER BY produk.id DESC"
	// }

	// if request.Order == "ASC" {
	// 	query += " ORDER BY produk.id ASC"
	// }

	orderby := fmt.Sprintf(request.OrderBy)
	order := fmt.Sprintf(request.Order)

	query += ` ORDER BY ` + orderby + ` ` + order

	query = ReplaceSQL(query, "?")
	rows, err := ctx.RepoDB.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := ProdukDto(rows)

	if err != nil {
		log.Println("Error querying GetProduk : ", err)
	}

	return data, err
}

func (ctx produkRepository) GetSingleProduk(idproduk int) (data models.Produk, err error) {
	var query = `SELECT produk.id, kodeproduk, nama, (hargajual::money::numeric::integer), produk.gambar, jenis, produk.idcorporate, corporate.uraian, produk.idsatuanproduk, satuanproduk.uraian, produk.idkategoriproduk, kategoriproduk.uraian, (hargarombongan::money::numeric::integer), minrombongan, statusstok, currentstok, statuspaket
	FROM produk
	INNER JOIN corporate ON produk.idcorporate=corporate.id
	INNER JOIN satuanproduk ON produk.idsatuanproduk=satuanproduk.id
	INNER JOIN kategoriproduk ON produk.idkategoriproduk=kategoriproduk.id
	WHERE produk.deleted_at IS NULL AND produk.id=?`

	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, idproduk).Scan(&data.Id, &data.KodeProduk, &data.Nama, &data.HargaJual, &data.Gambar, &data.Jenis, &data.IdCorporate, &data.NamaCorporate, &data.IdSatuanProduk, &data.NamaSatuanProduk, &data.IdKategoriProduk, &data.NamaKategoriProduk, &data.HargaRombongan, &data.MinRombongan, &data.StatusStok, &data.CurrentStok, &data.StatusPaket)

	if err != nil {
		log.Println("Error querying GetSingleProduk: ", err)
	}

	return data, err
}

// Insert Produk
func (ctx produkRepository) InsertProduk(produk models.RequestAddProduk) (id int, err error) {
	query := `
	INSERT INTO produk (kodeproduk, nama, hargajual, gambar, jenis, idcorporate, idsatuanproduk, idkategoriproduk, hargarombongan, minrombongan, statusstok, currentstok, statuspaket, created_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING id`
	query = ReplaceSQL(query, "?")
	err = ctx.RepoDB.DB.QueryRow(query, &produk.KodeProduk, &produk.Nama, &produk.HargaJual, &produk.Gambar, &produk.Jenis, &produk.IdCorporate, &produk.IdSatuanProduk, &produk.IdKategoriProduk, &produk.HargaRombongan, &produk.MinRombongan, &produk.StatusStok, &produk.CurrentStok, &produk.StatusPaket, time.Now()).Scan(&id)

	if err != nil {
		log.Println("Error querying InsertProduk : ", err)
	}

	return
}

// Update Produk
func (ctx produkRepository) UpdateProduk(produk models.RequestUpdateProduk) (id int, err error) {
	query := `
	UPDATE produk SET 
		kodeproduk=?, nama=?, hargajual=?, gambar=?, jenis=?, idcorporate=?, idsatuanproduk=?, idkategoriproduk=?, hargarombongan=?, minrombongan=?, statusstok=?, currentstok=?, statuspaket=?, updated_at=?
	WHERE id=? RETURNING id
	`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, &produk.KodeProduk, &produk.Nama, &produk.HargaJual, &produk.Gambar, &produk.Jenis, &produk.IdCorporate, &produk.IdSatuanProduk, &produk.IdKategoriProduk, &produk.HargaRombongan, &produk.MinRombongan, &produk.StatusStok, &produk.CurrentStok, &produk.StatusPaket, time.Now(), produk.Id).Scan(&id)

	if err != nil {
		log.Println("Error querying UpdateProduk : ", err)
	}

	return
}

// Delete Produk
func (ctx produkRepository) DeleteProduk(produk models.RequestDeleteProduk) (id int, err error) {
	query := `
		UPDATE produk SET deleted_at=? WHERE id=? RETURNING id
	`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, time.Now(), produk.Id).Scan(&id)

	if err != nil {
		log.Println("Error querying UpdateProduk : ", err)
	}

	return
}

func ProdukDto(rows *sql.Rows) (result []models.Produk, err error) {
	for rows.Next() {
		var produk models.Produk
		err = rows.Scan(&produk.Id, &produk.KodeProduk, &produk.Nama, &produk.HargaJual, &produk.Gambar, &produk.Jenis, &produk.IdCorporate, &produk.NamaCorporate, &produk.IdSatuanProduk, &produk.NamaSatuanProduk, &produk.IdKategoriProduk, &produk.NamaKategoriProduk, &produk.HargaRombongan, &produk.MinRombongan, &produk.StatusStok, &produk.CurrentStok, &produk.StatusPaket)
		if err != nil {
			return
		}
		result = append(result, produk)
	}

	return
}
