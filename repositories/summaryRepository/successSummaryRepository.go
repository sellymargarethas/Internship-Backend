package summaryrepository

import (
	"Internship-Backend/models"
	"Internship-Backend/repositories"
	"log"
	"strings"

	. "Internship-Backend/utils"
	"database/sql"
	"fmt"
)

type successSummaryRepository struct {
	RepoDB repositories.Repository
}

func NewSuccessSummaryRepository(repoDB repositories.Repository) successSummaryRepository {
	return successSummaryRepository{
		RepoDB: repoDB,
	}
}
func (ctx successSummaryRepository) ViewSuccessSummary(request models.RequestSummarySuccess, tx *sql.Tx) (temp string, err error) {
	var args []interface{}

	query := `SELECT public.fs_summary_trx(?, ?, ?);`
	args = append(args, &request.Corporate, &request.StartDate, &request.EndDate)
	query = ReplaceSQL(query, "?")

	res, err := tx.Query(query, args...)
	if err != nil {
		// tx.Rollback()
		return
	}
	// var temp string
	res.Close()
	for res.Next() {
		err = res.Scan(&temp)
		// fmt.Println(temp)
		if err != nil {
			log.Print(err)
		}

	}

	return
}
func (ctx successSummaryRepository) ScanResultSuccess(temp string, tx *sql.Tx) (response models.ScanSummary, err error) {
	err = tx.QueryRow(`FETCH ALL IN "summaryCard";`).Scan(&response.JumlahTrx, &response.TotalBruto, &response.TotalServiceFee, &response.TotalMDR, &response.TotalPaymentFee, &response.TotalVendorFee, &response.CountTunai, &response.CountLainnya, &response.CountCC, &response.CountDebit, &response.CountQRIS, &response.CountBrizzi, &response.CountEMoney, &response.CountTapCash, &response.CountFlazz, &response.CountJakcard, &response.CountVA, &response.CountBiller, &response.TotalTunai, &response.TotalLain, &response.TotalCC, &response.TotalDebit, &response.TotalQRIS, &response.TotalBrizzi, &response.TotalEMoney, &response.TotalTapCash, &response.TotalFlazz, &response.TotalJakcard, &response.TotalVA, &response.TotalBiller, &response.SFeeTunai, &response.SFeeLainnya, &response.SFeeCC, &response.SFeeDebit, &response.SFeeQRIS, &response.SFeeBrizzi, &response.SFeeEMoney, &response.SFeeTapCash, &response.SFeeFlazz, &response.SFeeJakcard, &response.SFeeVA, &response.SFeeBiller, &response.MdrTunai, &response.MdrLainnya, &response.MdrCC, &response.MdrDebit, &response.MdrQRIS, &response.MdrBrizzi, &response.MdrEMoney, &response.MdrTapCash, &response.MdrFlazz, &response.MdrJakcard, &response.MdrVA, &response.MdrBiller, &response.PayfeeTunai, &response.PayfeeLainnya, &response.PayfeeCC, &response.PayfeeDebit, &response.PayfeeQRIS, &response.PayfeeBrizzi, &response.PayfeeEMoney, &response.PayfeeTapCash, &response.PayfeeFlazz, &response.PayfeeJakcard, &response.PayfeeVA, &response.PayfeeBiller, &response.VendorfeeTunai, &response.VendorfeeLainnya, &response.VendorfeeCC, &response.VendorfeeDebit, &response.VendorfeeQRIS, &response.VendorfeeBrizzi, &response.VendorfeeEMoney, &response.VendorfeeTapCash, &response.VendorfeeFlazz, &response.VendorfeeJakcard, &response.VendorfeeVA, &response.VendorfeeBiller)
	if err != nil {
		// tx.Rollback()
		log.Println("error scan result success: ", err.Error())
		return
	}
	return
}

// func (ctx successSummaryRepository) ViewSettledSummary(request models.RequestSummarySettled, tx *sql.Tx) (temp string, err error) {
// 	var args []interface{}

// 	query := `SELECT public.fs_settlement_trx(?, ?, ?, ARRAY[`
// 	args = append(args, &request.Corporate, &request.StartDate, &request.EndDate)
// 	for _, row := range request.SettlementDestination {
// 		query += "?,"
// 		args = append(args, row.IdCoreSettlement)
// 	}
// 	query = strings.TrimRight(query, ",")

// 	query += `]::varchar[])`

// 	query = ReplaceSQL(query, "?")

// 	res, err := tx.Query(query, args...)

// 	if err != nil {
// 		fmt.Println("cek error 1: ", err)
// 	}
// 	res.Close()
// 	for res.Next() {
// 		err = res.Scan(&temp)
// 		if err != nil {
// 			fmt.Print(err)
// 		}
// 	}
// 	return
// }

func (ctx successSummaryRepository) ViewSettledSummary(request models.RequestSummarySettled, tx *sql.Tx) (response models.ScanSettlement, err error) {
	var args []interface{}

	query := `
	SELECT 
		COALESCE(SUM(transaction_count), 0::numeric) transaction_count, 
		COALESCE(SUM(transaction_sum), 0::numeric) transaction_sum, 
		COALESCE(SUM(servicefee_sum), 0::numeric) servicefee_sum, 
		COALESCE(SUM(mdr_sum), 0::numeric) mdr_sum,
		COALESCE(SUM(payfee_sum), 0::numeric) payfee_sum, 
		COALESCE(SUM(vendorfee_sum), 0::numeric) vendorfee_sum, 
		COALESCE(SUM(tunai_count), 0::numeric) tunai_count, 
		COALESCE(SUM(lainnya_count), 0::numeric) lainnya_count, COALESCE(SUM(cc_count), 0::numeric) cc_count, COALESCE(SUM(debit_count), 0::numeric) debit_count, COALESCE(SUM(qris_count), 0::numeric) qris_count, COALESCE(SUM(brizzi_count), 0::numeric) brizzi_count, COALESCE(SUM(emoney_count), 0::numeric) emoney_count, COALESCE(SUM(tapcash_count), 0::numeric) tapcash_count, COALESCE(SUM(flazz_count), 0::numeric) flazz_count, COALESCE(SUM(jakcard_count), 0::numeric) jakcard_count, COALESCE(SUM(va_count), 0::numeric) va_count, COALESCE(SUM(biller_count), 0::numeric) biller_count,
		COALESCE(SUM(tunai_sum), 0::numeric) tunai_sum, COALESCE(SUM(lainnya_sum), 0::numeric) lainnya_sum, COALESCE(SUM(cc_sum), 0::numeric) cc_sum, COALESCE(SUM(debit_sum), 0::numeric) debit_sum, COALESCE(SUM(qris_sum), 0::numeric) qris_sum, COALESCE(SUM(brizzi_sum), 0::numeric) brizzi_sum, COALESCE(SUM(emoney_sum), 0::numeric) emoney_sum, COALESCE(SUM(tapcash_sum), 0::numeric) tapcash_sum, COALESCE(SUM(flazz_sum), 0::numeric) flazz_sum, COALESCE(SUM(jakcard_sum), 0::numeric) jakcard_sum, COALESCE(SUM(va_sum), 0::numeric) va_sum, COALESCE(SUM(biller_sum), 0::numeric) biller_sum,
		COALESCE(SUM(tunai_sfee), 0::numeric) tunai_sfee, COALESCE(SUM(lainnya_sfee), 0::numeric) lainnya_sfee, COALESCE(SUM(cc_sfee) , 0::numeric)cc_sfee, COALESCE(SUM(debit_sfee), 0::numeric) debit_sfee, COALESCE(SUM(qris_sfee), 0::numeric) qris_sfee, COALESCE(SUM(brizzi_sfee), 0::numeric) brizzi_sfee, COALESCE(SUM(emoney_sfee), 0::numeric) emoney_sfee, COALESCE(SUM(tapcash_sfee), 0::numeric) tapcash_sfee, COALESCE(SUM(flazz_sfee), 0::numeric) flazz_sfee, COALESCE(SUM(jakcard_sfee), 0::numeric) jakcard_sfee, COALESCE(SUM(va_sfee), 0::numeric) va_sfee, COALESCE(SUM(biller_sfee), 0::numeric) biller_sfee,
		COALESCE(SUM(tunai_mdr), 0::numeric) tunai_mdr, COALESCE(SUM(lainnya_mdr), 0::numeric) lainnya_mdr, COALESCE(SUM(cc_mdr), 0::numeric) cc_mdr, COALESCE(SUM(debit_mdr), 0::numeric) debit_mdr, COALESCE(SUM(qris_mdr), 0::numeric) qris_mdr, COALESCE(SUM(brizzi_mdr), 0::numeric) brizzi_mdr, COALESCE(SUM(emoney_mdr), 0::numeric) emoney_mdr, COALESCE(SUM(tapcash_mdr), 0::numeric) tapcash_mdr, COALESCE(SUM(flazz_mdr), 0::numeric) flazz_mdr, COALESCE(SUM(jakcard_mdr), 0::numeric) jakcard_mdr, COALESCE(SUM(va_mdr), 0::numeric) va_mdr, COALESCE(SUM(biller_mdr), 0::numeric) biller_mdr,
		COALESCE(SUM(tunai_payfee), 0::numeric) tunai_payfee, COALESCE(SUM(lainnya_payfee), 0::numeric) lainnya_payfee, COALESCE(SUM(cc_payfee), 0::numeric) cc_payfee, COALESCE(SUM(debit_payfee), 0::numeric) debit_payfee, COALESCE(SUM(qris_payfee), 0::numeric) qris_payfee, COALESCE(SUM(brizzi_payfee), 0::numeric) brizzi_payfee, COALESCE(SUM(emoney_payfee), 0::numeric) emoney_payfee, COALESCE(SUM(tapcash_payfee), 0::numeric) tapcash_payfee, COALESCE(SUM(flazz_payfee), 0::numeric) flazz_payfee, COALESCE(SUM(jakcard_payfee), 0::numeric) jakcard_payfee, COALESCE(SUM(va_payfee), 0::numeric) va_payfee, COALESCE(SUM(biller_payfee), 0::numeric) biller_payfee,
		COALESCE(SUM(tunai_vendorfee), 0::numeric) tunai_vendorfee, COALESCE(SUM(lainnya_vendorfee), 0::numeric) lainnya_vendorfee, COALESCE(SUM(cc_vendorfee), 0::numeric) cc_vendorfee, COALESCE(SUM(debit_vendorfee), 0::numeric) debit_vendorfee, COALESCE(SUM(qris_vendorfee), 0::numeric) qris_vendorfee, COALESCE(SUM(brizzi_vendorfee), 0::numeric) brizzi_vendorfee, COALESCE(SUM(emoney_vendorfee), 0::numeric) emoney_vendorfee, COALESCE(SUM(tapcash_vendorfee), 0::numeric) tapcash_vendorfee, COALESCE(SUM(flazz_vendorfee), 0::numeric) flazz_vendorfee, COALESCE(SUM(jakcard_vendorfee), 0::numeric) jakcard_vendorfee, COALESCE(SUM(va_vendorfee), 0::numeric) va_vendorfee, COALESCE(SUM(biller_vendorfee), 0::numeric) biller_vendorfee
	FROM vw_settlement_trx_materialized
	WHERE
		hirarki_id LIKE ?||'%' AND
		date_settled_at >= ? and date_settled_at <= ?`
	args = append(args, &request.Corporate, &request.StartDate, &request.EndDate)

	if request.KategoriPembayaran != nil {
		query += ` AND kategoripembayaran=ANY(ARRAY[?])`
		args = append(args, &request.KategoriPembayaran)
	}
	if request.SettlementDestination != nil {
		query += ` AND settlement_destination=ANY(ARRAY[`
	}

	for _, row := range request.SettlementDestination {
		query += "?,"
		args = append(args, row.IdCoreSettlement)
	}

	if request.SettlementDestination != nil {
		query = strings.TrimRight(query, ",")
		query += `]::varchar[])`
	}

	query = ReplaceSQL(query, "?")
	fmt.Println(query)
	// res, err := tx.Query(query, args...)
	err = ctx.RepoDB.DB.QueryRow(query, args...).Scan(&response.JumlahTrx, &response.TotalBruto, &response.TotalServiceFee, &response.TotalMDR, &response.TotalPaymentFee, &response.TotalVendorFee, &response.CountTunai, &response.CountLainnya, &response.CountCC, &response.CountDebit, &response.CountQRIS, &response.CountBrizzi, &response.CountEMoney, &response.CountTapCash, &response.CountFlazz, &response.CountJakcard, &response.CountVA, &response.CountBiller, &response.TotalTunai, &response.TotalLain, &response.TotalCC, &response.TotalDebit, &response.TotalQRIS, &response.TotalBrizzi, &response.TotalEMoney, &response.TotalTapCash, &response.TotalFlazz, &response.TotalJakcard, &response.TotalVA, &response.TotalBiller, &response.SFeeTunai, &response.SFeeLainnya, &response.SFeeCC, &response.SFeeDebit, &response.SFeeQRIS, &response.SFeeBrizzi, &response.SFeeEMoney, &response.SFeeTapCash, &response.SFeeFlazz, &response.SFeeJakcard, &response.SFeeVA, &response.SFeeBiller, &response.MdrTunai, &response.MdrLainnya, &response.MdrCC, &response.MdrDebit, &response.MdrQRIS, &response.MdrBrizzi, &response.MdrEMoney, &response.MdrTapCash, &response.MdrFlazz, &response.MdrJakcard, &response.MdrVA, &response.MdrBiller, &response.PayfeeTunai, &response.PayfeeLainnya, &response.PayfeeCC, &response.PayfeeDebit, &response.PayfeeQRIS, &response.PayfeeBrizzi, &response.PayfeeEMoney, &response.PayfeeTapCash, &response.PayfeeFlazz, &response.PayfeeJakcard, &response.PayfeeVA, &response.PayfeeBiller, &response.VendorfeeTunai, &response.VendorfeeLainnya, &response.VendorfeeCC, &response.VendorfeeDebit, &response.VendorfeeQRIS, &response.VendorfeeBrizzi, &response.VendorfeeEMoney, &response.VendorfeeTapCash, &response.VendorfeeFlazz, &response.VendorfeeJakcard, &response.VendorfeeVA, &response.VendorfeeBiller)
	if err != nil {
		fmt.Println("cek error 1: ", err)
	}
	return
}

// func (ctx successSummaryRepository) ScanResultSettled(temp string, tx *sql.Tx) (response models.ScanSettlement, err error) {
// 	err = Scan(&response.JumlahTrx, &response.TotalBruto, &response.TotalServiceFee, &response.TotalMDR, &response.TotalPaymentFee, &response.TotalVendorFee, &response.CountTunai, &response.CountLainnya, &response.CountCC, &response.CountDebit, &response.CountQRIS, &response.CountBrizzi, &response.CountEMoney, &response.CountTapCash, &response.CountFlazz, &response.CountJakcard, &response.CountVA, &response.CountBiller, &response.TotalTunai, &response.TotalLain, &response.TotalCC, &response.TotalDebit, &response.TotalQRIS, &response.TotalBrizzi, &response.TotalEMoney, &response.TotalTapCash, &response.TotalFlazz, &response.TotalJakcard, &response.TotalVA, &response.TotalBiller, &response.SFeeTunai, &response.SFeeLainnya, &response.SFeeCC, &response.SFeeDebit, &response.SFeeQRIS, &response.SFeeBrizzi, &response.SFeeEMoney, &response.SFeeTapCash, &response.SFeeFlazz, &response.SFeeJakcard, &response.SFeeVA, &response.SFeeBiller, &response.MdrTunai, &response.MdrLainnya, &response.MdrCC, &response.MdrDebit, &response.MdrQRIS, &response.MdrBrizzi, &response.MdrEMoney, &response.MdrTapCash, &response.MdrFlazz, &response.MdrJakcard, &response.MdrVA, &response.MdrBiller, &response.PayfeeTunai, &response.PayfeeLainnya, &response.PayfeeCC, &response.PayfeeDebit, &response.PayfeeQRIS, &response.PayfeeBrizzi, &response.PayfeeEMoney, &response.PayfeeTapCash, &response.PayfeeFlazz, &response.PayfeeJakcard, &response.PayfeeVA, &response.PayfeeBiller, &response.VendorfeeTunai, &response.VendorfeeLainnya, &response.VendorfeeCC, &response.VendorfeeDebit, &response.VendorfeeQRIS, &response.VendorfeeBrizzi, &response.VendorfeeEMoney, &response.VendorfeeTapCash, &response.VendorfeeFlazz, &response.VendorfeeJakcard, &response.VendorfeeVA, &response.VendorfeeBiller)

// 	if err != nil {
// 		fmt.Println("Err ScanResultSettled : ", err.Error())
// 	}

// 	return
// }
