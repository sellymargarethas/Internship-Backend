package transactionrepository

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/repositories"
	. "Internship-Backend/utils"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type listtransactionRepository struct {
	RepoDB repositories.Repository
}

func NewListTransactionRepository(repoDB repositories.Repository) listtransactionRepository {
	return listtransactionRepository{
		RepoDB: repoDB,
	}
}

func (ctx listtransactionRepository) IsListTrxExistsByIndex(listtrx models.RequestTrx) (models.ResponseListTrx, bool, error) {
	var args []interface{}
	var result models.ResponseListTrx

	query := `
	SELECT id, noheader, merchantnoref, payment_method_name, payment_category_name, settlement_dest, payment_status, payment_status_desc, status_settlement, response_noref, corporate_cid, device_id, corporate_name, bank_mid, bank_tid, card_pan, card_type, COALESCE(payment_amount, 0), COALESCE(payment_disc, 0), payment_promo_code, COALESCE(payment_mdr, 0), COALESCE(service_fee, 0), COALESCE(payment_fee, 0), COALESCE(vendor_fee, 0), created_at
	FROM trx_v2 WHERE`

	if listtrx.IdTrx != constants.EMPTY_VALUE_INT {
		query += ` trx_v2.id = ? `
		args = append(args, listtrx.IdTrx)
	}

	query += ` ORDER BY id ASC`

	query = ReplaceSQL(query, "?")

	rows, err := ctx.RepoDB.DB.Query(query, args...)

	if err != nil {
		return result, constants.FALSE_VALUE, err
	}
	defer rows.Close()

	data, err := ListTrxDto(rows)
	if err != nil {
		return result, constants.FALSE_VALUE, err
	}

	if len(data) == constants.EMPTY_VALUE_INT {
		return result, constants.FALSE_VALUE, err
	}
	return data[0], constants.TRUE_VALUE, err
}

// Get All Transaction
func (ctx listtransactionRepository) GetListTransaction(request models.RequestListTrx, page models.PageOffsetLimit) ([]models.ResponseListTrx, error) {
	var args []interface{}

	query := `SELECT id, noheader, merchantnoref, payment_method_name, payment_category_name, settlement_dest, payment_status, payment_status_desc, status_settlement, response_noref, corporate_cid, device_id, corporate_name, bank_mid, bank_tid, card_pan, card_type, COALESCE(payment_amount, 0) AS payment_amount, COALESCE(payment_disc, 0) AS payment_disc, payment_promo_code, COALESCE(payment_mdr, 0) AS payment_mdr, COALESCE(service_fee, 0) AS service_fee, COALESCE(payment_fee, 0) AS payment_fee, COALESCE(vendor_fee, 0) AS vendor_fee, (created_at::text)
	FROM trx_v2 WHERE transaction_date >= ?`

	args = append(args, request.StartDate)

	if request.EndDate != constants.EMPTY_VALUE {
		query += ` AND transaction_date <= ?`
		args = append(args, request.EndDate)
	}

	if request.StatusTrx != constants.EMPTY_VALUE_INT {
		query += ` AND payment_status = ?`
		args = append(args, request.StatusTrx)
	}

	if request.KategoriPembayaran != constants.EMPTY_VALUE {
		query += ` AND payment_category_name = ?`
		args = append(args, request.KategoriPembayaran)
	}

	if request.Corporate != constants.EMPTY_VALUE {
		query += ` AND corporate_hirarki LIKE ? || '%'`
		args = append(args, request.Corporate)
	}

	if len(request.SattlementDestination) != constants.EMPTY_VALUE_INT {
		query += ` AND settlement_dest IN(`
		for _, row := range request.SattlementDestination {
			query += "?,"
			args = append(args, row.SattlementDestination)
		}
		query = strings.TrimRight(query, ",")
		query += ")"
	}

	if request.Pagination.Keyword != "" {
		query += ` AND (CAST(trx_v2.id AS TEXT) ILIKE '%' || ? || '%' OR
		trx_v2.noheader ILIKE '%' || ? || '%' OR
		trx_v2.merchantnoref ILIKE '%' || ? || '%' OR
		trx_v2.payment_method_name ILIKE '%' || ? || '%' OR
		trx_v2.payment_category_name ILIKE '%' || ? || '%' OR
		trx_v2.settlement_dest ILIKE '%' || ? || '%' OR
		CAST(trx_v2.payment_status AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(trx_v2.status_settlement AS TEXT) ILIKE '%' || ? || '%' OR
		trx_v2.response_noref ILIKE '%' || ? || '%' OR
		trx_v2.corporate_hirarki ILIKE '%' || ? || '%' OR
		trx_v2.device_id ILIKE '%' || ? || '%' OR
		trx_v2.corporate_name ILIKE '%' || ? || '%' OR
		trx_v2.bank_mid ILIKE '%' || ? || '%' OR
		trx_v2.bank_tid ILIKE '%' || ? || '%' OR
		trx_v2.card_pan ILIKE '%' || ? || '%' OR
		trx_v2.card_type ILIKE '%' || ? || '%' OR
		CAST(trx_v2.payment_amount AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(trx_v2.payment_disc AS TEXT) ILIKE '%' || ? || '%' OR
		trx_v2.payment_promo_code ILIKE '%' || ? || '%' OR
		CAST(trx_v2.payment_mdr AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(trx_v2.service_fee AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(trx_v2.payment_fee AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(trx_v2.vendor_fee AS TEXT) ILIKE '%' || ? || '%')`
		args = append(args, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword)
	}

	// if request.Pagination.Order == "DESC" {
	// 	query += " ORDER BY trx_v2.id DESC"
	// }

	// if request.Pagination.Order == "ASC" {
	// 	query += " ORDER BY trx_v2.id ASC"
	// }

	orderby := fmt.Sprintf(request.Pagination.OrderBy)
	order := fmt.Sprintf(request.Pagination.Order)

	query += ` ORDER BY ` + orderby + ` ` + order

	query += ` LIMIT ? OFFSET ?`
	args = append(args, page.Limit, page.Offset)

	query = ReplaceSQL(query, "?")
	// fmt.Println(query)
	rows, err := ctx.RepoDB.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := ListTrxDto(rows)

	if err != nil {
		log.Println("Error querying GetListTransaction : ", err)
	}

	return data, err
}

// Get All Transaction
func (ctx listtransactionRepository) GetListTransactionSettled(request models.RequestListTrx, page models.PageOffsetLimit) ([]models.ResponseListTrx, error) {
	var args []interface{}

	query := `SELECT id, noheader, merchantnoref, payment_method_name, payment_category_name, settlement_dest, payment_status, payment_status_desc, status_settlement, response_noref, corporate_cid, device_id, corporate_name, bank_mid, bank_tid, card_pan, card_type, COALESCE(payment_amount, 0) AS payment_amount, COALESCE(payment_disc, 0) AS payment_disc, payment_promo_code, COALESCE(payment_mdr, 0) AS payment_mdr, COALESCE(service_fee, 0) AS service_fee, COALESCE(payment_fee, 0) AS payment_fee, COALESCE(vendor_fee, 0) AS vendor_fee, (created_at::text)
	FROM trx_v2 WHERE COALESCE(date(settled_at), date(updated_at), date(created_at)) >= ?`

	args = append(args, request.StartDate)

	if request.EndDate != constants.EMPTY_VALUE {
		query += ` AND COALESCE(date(settled_at), date(updated_at), date(created_at)) <= ?`
		args = append(args, request.EndDate)
	}

	if request.StatusTrx != constants.EMPTY_VALUE_INT {
		query += ` AND payment_status = ?`
		args = append(args, request.StatusTrx)
	}

	if request.KategoriPembayaran != constants.EMPTY_VALUE {
		query += ` AND payment_category_name = ?`
		args = append(args, request.KategoriPembayaran)
	}

	if request.Corporate != constants.EMPTY_VALUE {
		query += ` AND corporate_hirarki LIKE ? || '%'`
		args = append(args, request.Corporate)
	}

	if len(request.SattlementDestination) != constants.EMPTY_VALUE_INT {
		query += ` AND settlement_dest IN(`
		for _, row := range request.SattlementDestination {
			query += "?,"
			args = append(args, row.SattlementDestination)
		}
		query = strings.TrimRight(query, ",")
		query += ")"
	}

	if request.Pagination.Keyword != "" {
		query += ` AND (CAST(trx_v2.id AS TEXT) ILIKE '%' || ? || '%' OR
		trx_v2.noheader ILIKE '%' || ? || '%' OR
		trx_v2.merchantnoref ILIKE '%' || ? || '%' OR
		trx_v2.payment_method_name ILIKE '%' || ? || '%' OR
		trx_v2.payment_category_name ILIKE '%' || ? || '%' OR
		trx_v2.settlement_dest ILIKE '%' || ? || '%' OR
		CAST(trx_v2.payment_status AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(trx_v2.status_settlement AS TEXT) ILIKE '%' || ? || '%' OR
		trx_v2.response_noref ILIKE '%' || ? || '%' OR
		trx_v2.corporate_hirarki ILIKE '%' || ? || '%' OR
		trx_v2.device_id ILIKE '%' || ? || '%' OR
		trx_v2.corporate_name ILIKE '%' || ? || '%' OR
		trx_v2.bank_mid ILIKE '%' || ? || '%' OR
		trx_v2.bank_tid ILIKE '%' || ? || '%' OR
		trx_v2.card_pan ILIKE '%' || ? || '%' OR
		trx_v2.card_type ILIKE '%' || ? || '%' OR
		CAST(trx_v2.payment_amount AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(trx_v2.payment_disc AS TEXT) ILIKE '%' || ? || '%' OR
		trx_v2.payment_promo_code ILIKE '%' || ? || '%' OR
		CAST(trx_v2.payment_mdr AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(trx_v2.service_fee AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(trx_v2.payment_fee AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(trx_v2.vendor_fee AS TEXT) ILIKE '%' || ? || '%')`
		args = append(args, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword)
	}

	// if request.Pagination.Order == "DESC" {
	// 	query += " ORDER BY trx_v2.id DESC"
	// }

	// if request.Pagination.Order == "ASC" {
	// 	query += " ORDER BY trx_v2.id ASC"
	// }

	orderby := fmt.Sprintf(request.Pagination.OrderBy)
	order := fmt.Sprintf(request.Pagination.Order)

	query += ` ORDER BY ` + orderby + ` ` + order

	query += ` LIMIT ? OFFSET ?`
	args = append(args, page.Limit, page.Offset)

	query = ReplaceSQL(query, "?")
	// fmt.Println(query)
	rows, err := ctx.RepoDB.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := ListTrxDto(rows)

	if err != nil {
		log.Println("Error querying GetListTransaction : ", err)
	}

	return data, err
}

func (ctx listtransactionRepository) CountListTransaction(request models.RequestListTrx) (count int, err error) {
	var args []interface{}

	query := `SELECT COUNT(id)
	FROM trx_v2 WHERE transaction_date >= ?`

	args = append(args, request.StartDate)

	if request.EndDate != constants.EMPTY_VALUE {
		query += ` AND transaction_date <= ?`
		args = append(args, request.EndDate)
	}

	if request.StatusTrx != constants.EMPTY_VALUE_INT {
		query += ` AND payment_status = ?`
		args = append(args, request.StatusTrx)
	}

	if request.KategoriPembayaran != constants.EMPTY_VALUE {
		query += ` AND payment_category_name = ?`
		args = append(args, request.KategoriPembayaran)
	}

	if request.Corporate != constants.EMPTY_VALUE {
		query += ` AND corporate_hirarki LIKE ? || '%'`
		args = append(args, request.Corporate)
	}

	if len(request.SattlementDestination) != constants.EMPTY_VALUE_INT {
		query += ` AND settlement_dest IN(`
		for _, row := range request.SattlementDestination {
			query += "?,"
			args = append(args, row.SattlementDestination)
		}
		query = strings.TrimRight(query, ",")
		query += ")"
	}

	if request.Pagination.Keyword != "" {
		query += ` AND (CAST(trx_v2.id AS TEXT) ILIKE '%' || ? || '%' OR
		trx_v2.noheader ILIKE '%' || ? || '%' OR
		trx_v2.merchantnoref ILIKE '%' || ? || '%' OR
		trx_v2.payment_method_name ILIKE '%' || ? || '%' OR
		trx_v2.payment_category_name ILIKE '%' || ? || '%' OR
		trx_v2.settlement_dest ILIKE '%' || ? || '%' OR
		CAST(trx_v2.payment_status AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(trx_v2.status_settlement AS TEXT) ILIKE '%' || ? || '%' OR
		trx_v2.response_noref ILIKE '%' || ? || '%' OR
		trx_v2.corporate_hirarki ILIKE '%' || ? || '%' OR
		trx_v2.device_id ILIKE '%' || ? || '%' OR
		trx_v2.corporate_name ILIKE '%' || ? || '%' OR
		trx_v2.bank_mid ILIKE '%' || ? || '%' OR
		trx_v2.bank_tid ILIKE '%' || ? || '%' OR
		trx_v2.card_pan ILIKE '%' || ? || '%' OR
		trx_v2.card_type ILIKE '%' || ? || '%' OR
		CAST(trx_v2.payment_amount AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(trx_v2.payment_disc AS TEXT) ILIKE '%' || ? || '%' OR
		trx_v2.payment_promo_code ILIKE '%' || ? || '%' OR
		CAST(trx_v2.payment_mdr AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(trx_v2.service_fee AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(trx_v2.payment_fee AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(trx_v2.vendor_fee AS TEXT) ILIKE '%' || ? || '%')`
		args = append(args, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword)
	}

	query = ReplaceSQL(query, "?")
	// fmt.Println(query)
	err = ctx.RepoDB.DB.QueryRow(query, args...).Scan(&count)

	if err != nil {
		log.Println("Error querying COUNTGetListTransaction : ", err)
	}

	return
}

func (ctx listtransactionRepository) CountListTransactionSettled(request models.RequestListTrx) (count int, err error) {
	var args []interface{}

	query := `SELECT COUNT(1)
	FROM trx_v2 WHERE COALESCE(date(settled_at), date(updated_at), date(created_at)) >= ?`

	args = append(args, request.StartDate)

	if request.EndDate != constants.EMPTY_VALUE {
		query += ` AND COALESCE(date(settled_at), date(updated_at), date(created_at)) <= ?`
		args = append(args, request.EndDate)
	}

	if request.StatusTrx != constants.EMPTY_VALUE_INT {
		query += ` AND payment_status = ?`
		args = append(args, request.StatusTrx)
	}

	if request.KategoriPembayaran != constants.EMPTY_VALUE {
		query += ` AND payment_category_name = ?`
		args = append(args, request.KategoriPembayaran)
	}

	if request.Corporate != constants.EMPTY_VALUE {
		query += ` AND corporate_hirarki LIKE ? || '%'`
		args = append(args, request.Corporate)
	}

	if len(request.SattlementDestination) != constants.EMPTY_VALUE_INT {
		query += ` AND settlement_dest IN(`
		for _, row := range request.SattlementDestination {
			query += "?,"
			args = append(args, row.SattlementDestination)
		}
		query = strings.TrimRight(query, ",")
		query += ")"
	}

	if request.Pagination.Keyword != "" {
		query += ` AND (CAST(trx_v2.id AS TEXT) ILIKE '%' || ? || '%' OR
		trx_v2.noheader ILIKE '%' || ? || '%' OR
		trx_v2.merchantnoref ILIKE '%' || ? || '%' OR
		trx_v2.payment_method_name ILIKE '%' || ? || '%' OR
		trx_v2.payment_category_name ILIKE '%' || ? || '%' OR
		trx_v2.settlement_dest ILIKE '%' || ? || '%' OR
		CAST(trx_v2.payment_status AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(trx_v2.status_settlement AS TEXT) ILIKE '%' || ? || '%' OR
		trx_v2.response_noref ILIKE '%' || ? || '%' OR
		trx_v2.corporate_hirarki ILIKE '%' || ? || '%' OR
		trx_v2.device_id ILIKE '%' || ? || '%' OR
		trx_v2.corporate_name ILIKE '%' || ? || '%' OR
		trx_v2.bank_mid ILIKE '%' || ? || '%' OR
		trx_v2.bank_tid ILIKE '%' || ? || '%' OR
		trx_v2.card_pan ILIKE '%' || ? || '%' OR
		trx_v2.card_type ILIKE '%' || ? || '%' OR
		CAST(trx_v2.payment_amount AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(trx_v2.payment_disc AS TEXT) ILIKE '%' || ? || '%' OR
		trx_v2.payment_promo_code ILIKE '%' || ? || '%' OR
		CAST(trx_v2.payment_mdr AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(trx_v2.service_fee AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(trx_v2.payment_fee AS TEXT) ILIKE '%' || ? || '%' OR
		CAST(trx_v2.vendor_fee AS TEXT) ILIKE '%' || ? || '%')`
		args = append(args, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword, request.Pagination.Keyword)
	}

	query = ReplaceSQL(query, "?")
	// fmt.Println(query)
	err = ctx.RepoDB.DB.QueryRow(query, args...).Scan(&count)

	if err != nil {
		log.Println("Error querying COUNTGetListTransaction : ", err)
	}

	return
}

func (ctx listtransactionRepository) GetSingleTransaction(idtrx int) (onetrx models.ResponseOneTrx, err error) {
	query := `SELECT id, noheader, merchantnoref, COALESCE(payment_amount, 0), 'PENDING' AS "Status Awal Transaction", payment_status, payment_status_desc, payment_category_name, response_noref, corporate_cid, corporate_name, device_id, bank_mid, bank_tid, card_pan, card_type, COALESCE(payment_amount, 0), COALESCE(payment_mdr, 0), COALESCE(payment_disc, 0), COALESCE(service_fee, 0), COALESCE(payment_disc, 0), (created_at::text)
	FROM trx_v2 WHERE trx_v2.id = ?`

	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, idtrx).Scan(&onetrx.ID, &onetrx.NomorHeader, &onetrx.MerchantNoRef, &onetrx.TerimaTunai, &onetrx.StatusAwalTrx, &onetrx.StatusTrx, &onetrx.StatusTrxDesc, &onetrx.MetodePembayaran, &onetrx.ResNoRef, &onetrx.CID, &onetrx.Corporate, &onetrx.DeviceId, &onetrx.MID, &onetrx.TID, &onetrx.CardPan, &onetrx.CardType, &onetrx.HargaJual, &onetrx.PaymentMDR, &onetrx.PaymentDisc, &onetrx.ServiceFee, &onetrx.Potongan, &onetrx.CreatedAt)

	if err != nil {
		return
	}

	return
}

func (ctx listtransactionRepository) GetSingleProdukTransaction(idtrx int) (detailproduktrx models.ResponseDetailProductTrx, err error) {
	var query = `
	SELECT 
		kodeproduk,
		namaproduk,
		kategoriproduk.uraian,
		CASE 
			WHEN trxdetail.idpaket = 0 THEN '-'
			ELSE (SELECT uraian FROM paket INNER JOIN trxdetail ON trxdetail.idpaket=paket.id WHERE trxdetail.idpaket=paket.id)
		END AS paket,
		COALESCE((hargajual::money::numeric::float), 0),
		COALESCE((diskon::money::numeric::float), 0),
		kuantitas,
		uraiansatuan,
		COALESCE((trxdetail.serviceFee::money::numeric::float), 0)
	FROM trxdetail
	INNER JOIN kategoriproduk ON trxdetail.idkategoriproduk=kategoriproduk.id
	WHERE trxdetail.idtrx=?`

	query = ReplaceSQL(query, "?")

	err = ctx.RepoDB.DB.QueryRow(query, idtrx).Scan(&detailproduktrx.KodeProduk, &detailproduktrx.NamaProduk, &detailproduktrx.Kategori, &detailproduktrx.Paket, &detailproduktrx.HargaJual, &detailproduktrx.Diskon, &detailproduktrx.Qty, &detailproduktrx.Satuan, &detailproduktrx.ServiceFee)

	if err != nil {
		return
	}

	return
}

func ListTrxDto(rows *sql.Rows) (result []models.ResponseListTrx, err error) {
	for rows.Next() {
		var listtrx models.ResponseListTrx
		err = rows.Scan(&listtrx.ID, &listtrx.NomorHeader, &listtrx.MerchantNoRef, &listtrx.Acquiring, &listtrx.MetodePembayaran, &listtrx.SettlementDestination, &listtrx.StatusTrx, &listtrx.StatusTrxDesc, &listtrx.StatusSettlement, &listtrx.ResNoRef, &listtrx.CID, &listtrx.DeviceId, &listtrx.Corporate, &listtrx.MID, &listtrx.TID, &listtrx.CardPan, &listtrx.CardType, &listtrx.HargaJual, &listtrx.Potongan, &listtrx.KodePromo, &listtrx.MDR, &listtrx.ServiceFee, &listtrx.PaymentFee, &listtrx.VendorFee, &listtrx.CreatedAt)
		if err != nil {
			return
		}
		result = append(result, listtrx)
	}

	return
}
