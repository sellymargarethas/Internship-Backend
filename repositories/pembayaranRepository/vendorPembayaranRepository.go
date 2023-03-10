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

type vendorpembayaranRepository struct {
	RepoDB repositories.Repository
}

func NewVendorPembayaranRepository(repoDB repositories.Repository) vendorpembayaranRepository {
	return vendorpembayaranRepository{
		RepoDB: repoDB,
	}
}

const defineColumn2 = `uraian`

func (ctx vendorpembayaranRepository) IsVendorPembayaranExistsByIndex(vendorpembayaran models.VendorPembayaran) (models.VendorPembayaran, bool, error) {
	var args []interface{}
	var result models.VendorPembayaran

	query := `
	SELECT id, ` + defineColumn2 + `
	FROM vendorpembayaran
	WHERE deleted_at IS NULL`

	if vendorpembayaran.Id != constants.EMPTY_VALUE_INT {
		query += ` AND vendorpembayaran.id = ? `
		args = append(args, vendorpembayaran.Id)
	}

	if vendorpembayaran.Uraian != constants.EMPTY_VALUE {
		query += ` AND vendorpembayaran.uraian = ? `
		args = append(args, vendorpembayaran.Uraian)
	}

	query += ` ORDER BY id ASC`

	query = ReplaceSQL(query, "?")

	rows, err := ctx.RepoDB.DB.Query(query, args...)

	if err != nil {
		return result, constants.FALSE_VALUE, err
	}
	defer rows.Close()

	data, err := VendorPembayaranDto(rows)
	if err != nil {
		return result, constants.FALSE_VALUE, err
	}

	if len(data) == constants.EMPTY_VALUE_INT {
		return result, constants.FALSE_VALUE, err
	}
	return data[0], constants.TRUE_VALUE, err
}

// Get All Vendor Pembayaran
func (ctx vendorpembayaranRepository) GetListVendorPembayaran() ([]models.VendorPembayaran, error) {
	var result []models.VendorPembayaran

	query := `SELECT id, ` + defineColumn2 + `
	FROM vendorpembayaran
	WHERE deleted_at IS NULL
	ORDER BY id ASC`

	rows, err := ctx.RepoDB.DB.Query(query)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	data, err := VendorPembayaranDto(rows)

	if err != nil {
		log.Println("Error querying GetVendorPembayaran : ", err)
	}

	return data, err
}

// Insert Vendor Pembayaran
func (ctx vendorpembayaranRepository) InsertVendorPembayaran(vendorpembayaran models.RequestAddVendorPembayaran) (id int, err error) {
	query := `
	INSERT INTO vendorpembayaran
	(` + defineColumn2 + `, created_at) 
	VALUES 
	(?, ?) returning id`
	query = ReplaceSQL(query, "?")
	err = ctx.RepoDB.DB.QueryRow(query, &vendorpembayaran.Uraian, time.Now()).Scan(&id)

	if err != nil {
		log.Println("Error querying InsertVendorPembayaran : ", err)
	}

	return
}

// Update Vendor Pembayaran
func (ctx vendorpembayaranRepository) UpdateVendorPembayaran(pembayaran models.RequestUpdateVendorPembayaran) (id int, err error) {
	query := `
	UPDATE vendorpembayaran SET 
		uraian=?, updated_at=?
	WHERE id=? RETURNING id
	`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, pembayaran.Uraian, time.Now(), pembayaran.Id).Scan(&id)

	if err != nil {
		log.Println("Error querying UpdateVendorPembayaran : ", err)
	}

	return
}

// Delete Vendor Pembayaran
func (ctx vendorpembayaranRepository) DeleteVendorPembayaran(pembayaran models.RequestDeleteVendorPembayaran) (id int, err error) {
	query := `
		UPDATE vendorpembayaran SET deleted_at=? WHERE id=? RETURNING id
	`
	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, time.Now(), pembayaran.Id).Scan(&id)

	if err != nil {
		log.Println("Error querying UpdateVendorPembayaran : ", err)
	}

	return
}

func VendorPembayaranDto(rows *sql.Rows) (result []models.VendorPembayaran, err error) {
	for rows.Next() {
		var pembayaran models.VendorPembayaran
		err = rows.Scan(&pembayaran.Id, &pembayaran.Uraian)
		if err != nil {
			return
		}
		result = append(result, pembayaran)
	}

	return
}
