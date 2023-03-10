package corporaterepository

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/repositories"
	"log"

	. "Internship-Backend/utils"

	"database/sql"
	"fmt"
	"time"
)

type corporateRepository struct {
	RepoDB repositories.Repository
}

func NewCorporateRepository(repoDB repositories.Repository) corporateRepository {
	return corporateRepository{
		RepoDB: repoDB,
	}
}

// const defineColumn = `uraian, cid, alamat, telepon, parent_cid, hirarki_id, level, gambar, ispercentage, idcorporatecategory, iplocalserver, idkota` // 12 fields
const defineColumn = `uraian, cid, alamat, telepon, hirarki_id, level, gambar, ispercentage, idcorporatecategory, iplocalserver, idkota` // 12 fields

func (ctx corporateRepository) IsCorporateExistsByIndex(corporate models.Corporate) (models.Corporate, bool, error) {
	var args []interface{}
	var result models.Corporate

	query := `
	SELECT corporate.id, corporate.` + defineColumn + `, (serviceFee::money::numeric::integer), corporate.nama_kota, corporate.nama_provinsi
	FROM corporate
	WHERE corporate.deleted_at IS NULL`

	if corporate.Id != constants.EMPTY_VALUE_INT {
		query += ` AND corporate.id = ? `
		args = append(args, corporate.Id)
	}

	if corporate.CID != constants.EMPTY_VALUE {
		query += ` AND corporate.cid = ? `
		args = append(args, corporate.CID)
	}

	query += ` ORDER BY hirarki_id ASC`

	query = ReplaceSQL(query, "?")

	rows, err := ctx.RepoDB.DB.Query(query, args...)

	if err != nil {
		return result, constants.FALSE_VALUE, err
	}
	defer rows.Close()

	data, err := corporate2Dto(rows)
	if err != nil {
		return result, constants.FALSE_VALUE, err
	}

	if len(data) == constants.EMPTY_VALUE_INT {
		return result, constants.FALSE_VALUE, err
	}
	return data[0], constants.TRUE_VALUE, err
}

func (ctx corporateRepository) GetListCorporate(request models.RequestList) ([]models.Corporate, error) {
	var args []interface{}

	query := `SELECT corporate.id, corporate.uraian, corporate.cid, corporate.alamat, corporate.telepon, corporate.hirarki_id, corporate.level, corporate.gambar, corporate.ispercentage, corporate.idcorporatecategory, corporate.iplocalserver, corporate.idkota, (corporate.serviceFee::money::numeric::integer), corporate.nama_kota, corporate.nama_provinsi, corporatecategory.uraian
	FROM corporate
	INNER JOIN corporatecategory ON corporatecategory.id = corporate.idcorporatecategory
	WHERE corporate.deleted_at IS NULL AND corporatecategory.deleted_at IS NULL
	AND corporate.hirarki_id LIKE ?|| '%'`
	args = append(args, request.HirarkiId)

	if request.Keyword != "" {
		query += ` AND (
		CAST(corporate.id AS TEXT) ILIKE '%' || ? || '%' OR
		corporate.uraian ILIKE '%' || ? || '%' OR
		corporate.cid ILIKE '%' || ? || '%' OR
		corporate.alamat ILIKE '%' || ? || '%' OR
		corporate.telepon ILIKE '%' || ? || '%' OR
		corporate.hirarki_id ILIKE '%' || ? || '%' OR
		CAST(corporate.level AS TEXT) ILIKE '%' || ? || '%' OR
		corporate.gambar ILIKE '%' || ? || '%' OR
		corporatecategory.uraian ILIKE '%' || ? || '%' OR
		corporate.iplocalserver ILIKE '%' || ? || '%' OR
		corporate.nama_kota ILIKE '%' || ? || '%' OR
		corporate.nama_provinsi ILIKE '%' || ? || '%' OR
		CAST(corporate.ispercentage AS TEXT) ILIKE '%' || ? || '%')`
		args = append(args, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword)
	}

	// if request.Order == "DESC" {
	// 	query += " ORDER BY corporate.id DESC"
	// }

	// if request.Order == "ASC" {
	// 	query += " ORDER BY corporate.id ASC"
	// }

	orderby := fmt.Sprintf(request.OrderBy)
	order := fmt.Sprintf(request.Order)

	query += ` ORDER BY ` + orderby + ` ` + order

	query = ReplaceSQL(query, "?")
	rows, err := ctx.RepoDB.DB.Query(query, args...)
	// fmt.Println(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := corporateDto(rows)

	if err != nil {
		log.Println("Error querying GetAllCorporate : ", err)
	}

	return data, err
}

func (ctx corporateRepository) GetSingleCorporate(corporate models.Corporate) ([]models.Corporate, error) {
	var args []interface{}

	var query = `SELECT corporate.id, corporate.` + defineColumn + `, (serviceFee::money::numeric::integer), corporate.nama_kota, corporate.nama_provinsi, corporatecategory.uraian
	FROM corporate
	INNER JOIN corporatecategory ON corporatecategory.id = corporate.idcorporatecategory
	WHERE corporate.deleted_at IS NULL`

	if corporate.Id != constants.EMPTY_VALUE_INT {
		query += ` AND corporate.id = ? `
		args = append(args, corporate.Id)
	}

	if corporate.CID != constants.EMPTY_VALUE {
		query += ` AND cid = ? `
		args = append(args, corporate.CID)
	}

	if corporate.Uraian != constants.EMPTY_VALUE {
		query += ` AND corporate.uraian ILIKE '%' || ? || '%' `
		args = append(args, corporate.Uraian)
	}

	if corporate.HirarkiId != constants.EMPTY_VALUE {
		query += ` AND hirarki_id ILIKE ? || '%'`
		args = append(args, corporate.HirarkiId)
	}

	query += ` ORDER BY hirarki_id ASC `
	query = ReplaceSQL(query, "?")

	rows, err := ctx.RepoDB.DB.Query(query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := corporateDto(rows)
	if err != nil {
		log.Println("Error querying GetSingleCorporate : ", err)
	}

	return data, err
}

func (ctx corporateRepository) InsertCorporate(corporate models.RequestAddCorporate) (id int, err error) {
	var cid string
	var hirarkiID string

	cid = fmt.Sprintf(" CONCAT(?::varchar,LPAD(nextval('cid_level_%d_sequence')::varchar,9,'0')) ", corporate.Level)
	hirarkiID = fmt.Sprintf(" CONCAT(?::varchar,'/',?::varchar,LPAD(currval('cid_level_%d_sequence')::varchar,9,'0')) ", corporate.Level)

	query := `INSERT INTO corporate 
	(` + defineColumn + `, serviceFee, nama_kota, nama_provinsi, created_at) 
	VALUES 
	(?,` + cid + `, ?, ?,` + hirarkiID + `, ?, ?, ?, ?, ?, ?, ?,
		(SELECT kota.uraian FROM kota WHERE id = ?), 
		(SELECT provinsi.uraian FROM provinsi INNER JOIN kota ON kota.idprovinsi = provinsi.id WHERE kota.id = ?), ?)
	RETURNING id`

	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, &corporate.Uraian, &corporate.Level, &corporate.Alamat, &corporate.Telepon, &corporate.HirarkiId, &corporate.Level, &corporate.Level, &corporate.Gambar, &corporate.IsPercentage, &corporate.IdCorporateCategory, &corporate.IpLocalServer, &corporate.IdKota, &corporate.ServiceFee, &corporate.IdKota, &corporate.IdKota, time.Now()).Scan(&id)

	if err != nil {
		log.Println("Error querying InsertCorporate : ", err)
	}

	return
}

func (ctx corporateRepository) UpdateCorporate(corporate models.RequestUpdateCorporate) (id int, err error) {
	query := `
	UPDATE corporate SET 
		uraian=?, alamat=?, telepon=?, iplocalserver=?, serviceFee=?, idkota=?, nama_kota=(SELECT kota.uraian FROM kota WHERE id = ?), nama_provinsi=(SELECT provinsi.uraian FROM provinsi INNER JOIN kota ON kota.idprovinsi = provinsi.id WHERE kota.id = ?), gambar=?, ispercentage=?, idcorporatecategory=?, updated_at=?
	WHERE id=?
	RETURNING id
	`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, corporate.Uraian, corporate.Alamat, corporate.Telepon, corporate.IpLocalServer, corporate.ServiceFee, corporate.IdKota, corporate.IdKota, corporate.IdKota, corporate.Gambar, corporate.IsPercentage, corporate.IdCorporateCategory, time.Now(), corporate.Id).Scan(&id)

	if err != nil {
		log.Println("Error querying UpdateCorporate : ", err)
	}

	return
}

func (ctx corporateRepository) DeleteCorporate(corporate models.RequestDeleteCorporate, tx *sql.Tx) (id int, cid string, err error) {
	query := `
		UPDATE corporate SET 
			deleted_at=? WHERE id=? 
		RETURNING id, cid
	`

	query = ReplaceSQL(query, "?")
	err = ctx.RepoDB.DB.QueryRow(query, time.Now(), corporate.Id).Scan(&id, &cid)

	if err != nil {
		log.Println("Error querying DeleteCorporate : ", err)
	}

	return
}

func (ctx corporateRepository) DeleteParentCorporate(cid string, tx *sql.Tx) (cid2 string, err error) {
	query := `
		UPDATE corporate SET 
			deleted_at=? WHERE corporate.hirarki_id ILIKE '%' || ? || '%'
		RETURNING id
	`

	query = ReplaceSQL(query, "?")
	err = ctx.RepoDB.DB.QueryRow(query, time.Now(), cid).Scan(&cid2)

	if err != nil {
		log.Println("Error querying DeleteParentCorporate : ", err)
	}

	return
}

func corporateDto(rows *sql.Rows) (result []models.Corporate, err error) {
	for rows.Next() {
		var corporate models.Corporate
		err = rows.Scan(&corporate.Id, &corporate.Uraian, &corporate.CID, &corporate.Alamat, &corporate.Telepon, &corporate.HirarkiId, &corporate.Level, &corporate.Gambar, &corporate.IsPercentage, &corporate.IdCorporateCategory, &corporate.IpLocalServer, &corporate.IdKota, &corporate.ServiceFee, &corporate.NamaKota, &corporate.NamaProvinsi, &corporate.NamaCorporateCategory)
		if err != nil {
			return
		}
		result = append(result, corporate)
	}

	return
}

func corporate2Dto(rows *sql.Rows) (result []models.Corporate, err error) {
	for rows.Next() {
		var corporate models.Corporate
		err = rows.Scan(&corporate.Id, &corporate.Uraian, &corporate.CID, &corporate.Alamat, &corporate.Telepon, &corporate.HirarkiId, &corporate.Level, &corporate.Gambar, &corporate.IsPercentage, &corporate.IdCorporateCategory, &corporate.IpLocalServer, &corporate.IdKota, &corporate.ServiceFee, &corporate.NamaKota, &corporate.NamaProvinsi)
		if err != nil {
			return
		}
		result = append(result, corporate)
	}

	return
}
