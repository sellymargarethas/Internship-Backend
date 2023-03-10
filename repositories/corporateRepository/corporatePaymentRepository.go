package corporaterepository

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/repositories"
	. "Internship-Backend/utils"
	"fmt"

	"database/sql"
	"log"
)

type corporatePaymentRepository struct {
	RepoDB repositories.Repository
}

func NewCorporatePaymentRepository(repoDB repositories.Repository) corporatePaymentRepository {
	return corporatePaymentRepository{
		RepoDB: repoDB,
	}
}

func (ctx corporatePaymentRepository) InsertCorporatePayment(corporatePayment models.RequestAddCorporatePayment) (result int, err error) {
	vals := []interface{}{}
	sqlStatement := `INSERT INTO corporatepayment(idcorporate, idmetodepembayaran, samnum) VALUES`
	for _, row := range corporatePayment.MetodePembayaranDetails {
		sqlStatement += "(?,?,?),"
		vals = append(vals, corporatePayment.IdCorporate, row.IdCorporatePayment, row.Samnum)
	}
	sqlStatement = sqlStatement[0 : len(sqlStatement)-1]
	sqlStatement = ReplaceSQL(sqlStatement, "?")

	stmt, _ := ctx.RepoDB.DB.Prepare(sqlStatement)

	_, err = stmt.Exec(vals...)
	if err != nil {
		return
	}
	return
}

func (ctx corporatePaymentRepository) GetListCorporatePayment(request models.RequestList) (result []models.ViewCorporatePayment, err error) {
	var args []interface{}
	query := `SELECT corporate.id, corporate.cid,corporate.uraian, corporate.nama_kota, 
	(SELECT COUNT(idmetodepembayaran) FROM corporatepayment WHERE corporatepayment.idcorporate=(corporate.id::int)) 
	FROM corporate
	WHERE corporate.deleted_at IS NULL
	AND corporate.hirarki_id LIKE ?||'%'
	`
	args = append(args, request.HirarkiId)
	if request.Keyword != "" {
		query += ` AND (CAST(corporate.id AS TEXT) ILIKE '%' || ? || '%' 
		OR corporate.cid ILIKE '%' || ? || '%' 
		OR corporate.uraian ILIKE '%' || ? || '%' 
		OR corporate.nama_kota ILIKE '%' || ? || '%'
		OR CAST((SELECT COUNT(idmetodepembayaran) 
		FROM corporatepayment 
		WHERE corporatepayment.idcorporate=(corporate.id::int)) AS TEXT) ILIKE '%' || ? || '%')
		`
		args = append(args, request.Keyword, request.Keyword, request.Keyword, request.Keyword, request.Keyword)
	}
	// if request.Order == "ASC" {
	// 	query += ` ORDER BY corporate.id ASC `
	// }
	// if request.Order == "DESC" {
	// 	query += ` ORDER BY corporate.id DESC `
	// }

	orderby := fmt.Sprintf(request.OrderBy)
	order := fmt.Sprintf(request.Order)

	query += ` ORDER BY ` + orderby + ` ` + order

	query = ReplaceSQL(query, "?")
	// fmt.Println(query)
	rows, err := ctx.RepoDB.DB.Query(query, args...)

	if err != nil {
		log.Println("error GetCorporatePayment Repository: ", err)
	}
	for rows.Next() {
		var corporate models.ViewCorporatePayment
		rows.Scan(&corporate.IdCorporate, &corporate.CID, &corporate.Uraian, &corporate.Kota, &corporate.JmlMetodePembayaran)
		result = append(result, corporate)
	}

	return
}

func (ctx corporatePaymentRepository) GetCorporatePaymentByIdCorporate(id int) (corporatepayment []models.MetodePembayaranDetails, err error) {
	query := `
	SELECT idmetodepembayaran, metodepembayaran.uraian, samnum
	FROM corporatepayment
	INNER JOIN metodepembayaran on metodepembayaran.id=corporatepayment.idmetodepembayaran
	WHERE corporatepayment.idcorporate = ?
	ORDER BY idcorporate ASC`
	query = ReplaceSQL(query, "?")
	rows, err := ctx.RepoDB.DB.Query(query, id)
	if err != nil {
		log.Println("ERROR: query: GetCorporatePaymentByIdCorporate", err)
	}

	for rows.Next() {
		var list models.MetodePembayaranDetails
		rows.Scan(&list.IdMetodePembayaran, &list.Uraian, &list.Samnum)
		corporatepayment = append(corporatepayment, list)
	}
	return
}

func (ctx corporatePaymentRepository) DeleteCorporatePaymentDetails(id int, tx *sql.Tx) (result bool, err error) {
	query := `
	DELETE
	FROM corporatepayment
	WHERE idcorporate=?`
	query = ReplaceSQL(query, "?")

	_, err = ctx.RepoDB.DB.Query(query, id)
	if err != nil {
		log.Println("ERROR InsertCorporateCategory Repository", err)
	}

	return
}

func (ctx corporatePaymentRepository) UpdateCorporatePayment(corporatePayment models.RequestUpdateCorporatePayment, tx *sql.Tx) (status bool, err error) {
	vals := []interface{}{}
	sqlStatement := `INSERT INTO corporatepayment(idcorporate, idmetodepembayaran, samnum) VALUES`
	for _, row := range corporatePayment.MetodePembayaranDetails {
		sqlStatement += "(?,?,?),"
		vals = append(vals, corporatePayment.IdCorporate, row.IdCorporatePayment, row.Samnum)
	}
	sqlStatement = sqlStatement[0 : len(sqlStatement)-1]
	sqlStatement = ReplaceSQL(sqlStatement, "?")

	stmt, _ := ctx.RepoDB.DB.Prepare(sqlStatement)

	_, err = stmt.Exec(vals...)
	if err != nil {
		return constants.FALSE_VALUE, err
	}
	return constants.TRUE_VALUE, nil
}

func (ctx corporatePaymentRepository) IsCorporatePaymentExistsByIndex(idCorporate int) (int, bool, error) {
	var value int

	query := `
	SELECT idcorporate
	FROM corporatepayment WHERE idcorporate = ?`

	query = ReplaceSQL(query, "?")

	err := ctx.RepoDB.DB.QueryRow(query, idCorporate).Scan(&value)

	if err != nil {
		if err == sql.ErrNoRows {
			return value, constants.FALSE_VALUE, err
		}
		return value, constants.FALSE_VALUE, err
	}
	return value, constants.TRUE_VALUE, err
}
