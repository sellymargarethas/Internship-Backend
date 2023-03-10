package dashboardReposi

import (
	"Internship-Backend/models"
	"Internship-Backend/repositories"
	. "Internship-Backend/utils"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type dashboardRepository struct {
	RepoDB repositories.Repository
}

func NewDashboardRepository(repoDB repositories.Repository) dashboardRepository {
	return dashboardRepository{
		RepoDB: repoDB,
	}
}

func (ctx dashboardRepository) SummaryDashboard(request models.RequestSummaryRevenue, tx *sql.Tx) (response models.ScanSummary, err error) {
	res, err := tx.Query(`SELECT public.fs_summary_trx($1, $2, $3);`, &request.Corporate, &request.StartDate, &request.EndDate)
	if err != nil {
		// tx.Rollback()
		return
	}
	var temp string
	res.Close()
	for res.Next() {
		err = res.Scan(&temp)
		if err != nil {
			log.Print(err)
		}

	}
	err = tx.QueryRow(`FETCH ALL IN "summaryCard";`).Scan(&response.JumlahTrx, &response.TotalBruto, &response.TotalServiceFee, &response.TotalMDR, &response.TotalPaymentFee, &response.TotalVendorFee, &response.CountTunai, &response.CountLainnya, &response.CountCC, &response.CountDebit, &response.CountQRIS, &response.CountBrizzi, &response.CountEMoney, &response.CountTapCash, &response.CountFlazz, &response.CountJakcard, &response.CountVA, &response.CountBiller, &response.TotalTunai, &response.TotalLain, &response.TotalCC, &response.TotalDebit, &response.TotalQRIS, &response.TotalBrizzi, &response.TotalEMoney, &response.TotalTapCash, &response.TotalFlazz, &response.TotalJakcard, &response.TotalVA, &response.TotalBiller, &response.SFeeTunai, &response.SFeeLainnya, &response.SFeeCC, &response.SFeeDebit, &response.SFeeQRIS, &response.SFeeBrizzi, &response.SFeeEMoney, &response.SFeeTapCash, &response.SFeeFlazz, &response.SFeeJakcard, &response.SFeeVA, &response.SFeeBiller, &response.MdrTunai, &response.MdrLainnya, &response.MdrCC, &response.MdrDebit, &response.MdrQRIS, &response.MdrBrizzi, &response.MdrEMoney, &response.MdrTapCash, &response.MdrFlazz, &response.MdrJakcard, &response.MdrVA, &response.MdrBiller, &response.PayfeeTunai, &response.PayfeeLainnya, &response.PayfeeCC, &response.PayfeeDebit, &response.PayfeeQRIS, &response.PayfeeBrizzi, &response.PayfeeEMoney, &response.PayfeeTapCash, &response.PayfeeFlazz, &response.PayfeeJakcard, &response.PayfeeVA, &response.PayfeeBiller, &response.VendorfeeTunai, &response.VendorfeeLainnya, &response.VendorfeeCC, &response.VendorfeeDebit, &response.VendorfeeQRIS, &response.VendorfeeBrizzi, &response.VendorfeeEMoney, &response.VendorfeeTapCash, &response.VendorfeeFlazz, &response.VendorfeeJakcard, &response.VendorfeeVA, &response.VendorfeeBiller)
	// tx.Commit()

	return
}

func (ctx dashboardRepository) TrafficGrossTrxHourly(request models.RequestTrafficGrossTrx) (response []models.ResponseSummaryOmset, err error) {
	var args []interface{}

	query := `SELECT 
				extract(hour from created_at) as trx_hourcount,
				count(1) as count
			FROM 
				trx_v2
			WHERE 
				transaction_date = ?
				AND corporate_hirarki LIKE ?||'%'
				AND payment_status = 1
			GROUP BY trx_hourcount
			ORDER BY trx_hourcount
		`
	args = append(args, request.Date, request.Corporate)

	query = ReplaceSQL(query, "?")

	rows, err := ctx.RepoDB.DB.Query(query, args...)

	for rows.Next() {
		var list models.ResponseSummaryOmset
		rows.Scan(&list.Date, &list.TrxCount)
		response = append(response, list)
	}
	return
}

func (ctx dashboardRepository) GrossTrxWeeklyByPayMethod(request models.RequestRevenue, tx *sql.Tx) (data []models.ResponseDashboardPaymentCategory, err error) {

	query := `SELECT public.fs_summary_trx($1,  
	(current_date-7)::text, 
	(current_date)::text);`
	res, err := tx.Query(query, &request.Corporate)
	if err != nil {
		// tx.Rollback()
		return
	}
	var temp string
	res.Close()
	for res.Next() {
		err = res.Scan(&temp)
		if err != nil {
			fmt.Print(err)
		}
		// fmt.Println(temp)
	}

	rows, err := tx.Query(`FETCH ALL IN "daily";`)
	for rows.Next() {
		var response models.ScanSummary
		rows.Scan(&response.DateString, &response.JumlahTrx, &response.TotalBruto, &response.TotalServiceFee, &response.TotalMDR, &response.TotalPaymentFee, &response.TotalVendorFee, &response.CountTunai, &response.CountLainnya, &response.CountCC, &response.CountDebit, &response.CountQRIS, &response.CountBrizzi, &response.CountEMoney, &response.CountTapCash, &response.CountFlazz, &response.CountJakcard, &response.CountVA, &response.CountBiller, &response.TotalTunai, &response.TotalLain, &response.TotalCC, &response.TotalDebit, &response.TotalQRIS, &response.TotalBrizzi, &response.TotalEMoney, &response.TotalTapCash, &response.TotalFlazz, &response.TotalJakcard, &response.TotalVA, &response.TotalBiller, &response.SFeeTunai, &response.SFeeLainnya, &response.SFeeCC, &response.SFeeDebit, &response.SFeeQRIS, &response.SFeeBrizzi, &response.SFeeEMoney, &response.SFeeTapCash, &response.SFeeFlazz, &response.SFeeJakcard, &response.SFeeVA, &response.SFeeBiller, &response.MdrTunai, &response.MdrLainnya, &response.MdrCC, &response.MdrDebit, &response.MdrQRIS, &response.MdrBrizzi, &response.MdrEMoney, &response.MdrTapCash, &response.MdrFlazz, &response.MdrJakcard, &response.MdrVA, &response.MdrBiller, &response.PayfeeTunai, &response.PayfeeLainnya, &response.PayfeeCC, &response.PayfeeDebit, &response.PayfeeQRIS, &response.PayfeeBrizzi, &response.PayfeeEMoney, &response.PayfeeTapCash, &response.PayfeeFlazz, &response.PayfeeJakcard, &response.PayfeeVA, &response.PayfeeBiller, &response.VendorfeeTunai, &response.VendorfeeLainnya, &response.VendorfeeCC, &response.VendorfeeDebit, &response.VendorfeeQRIS, &response.VendorfeeBrizzi, &response.VendorfeeEMoney, &response.VendorfeeTapCash, &response.VendorfeeFlazz, &response.VendorfeeJakcard, &response.VendorfeeVA, &response.VendorfeeBiller)
		// fmt.Println(response)
		response.DateCreatedAt, _ = time.Parse("2006-01-02", response.DateString)
		YYYYMMDD := "2006-01-02"
		hasil := models.ResponseDashboardPaymentCategory{
			Date:    response.DateCreatedAt.Format(YYYYMMDD),
			Day:     response.DateCreatedAt.Weekday().String(),
			Tunai:   response.TotalBruto,
			Lainnya: response.TotalLain,
			CC:      response.TotalCC,
			Debit:   response.TotalDebit,
			QRIS:    response.TotalQRIS,
			Brizzi:  response.TotalBrizzi,
			EMoney:  response.TotalEMoney,
			TapCash: response.TotalTapCash,
			Flazz:   response.TotalFlazz,
			Jakcard: response.TotalJakcard,
			VA:      response.TotalVA,
			Biller:  response.TotalBiller,
		}
		data = append(data, hasil)

	}
	// tx.Commit()

	return
}
func (ctx dashboardRepository) RevenueByCorporateId(request models.RequestSummaryRevenue, tx *sql.Tx) (data []models.ResponseSummaryOmset, err error) {
	res, err := tx.Query(`SELECT public.fs_summary_trx($1, $2
	, $3);`, &request.Corporate, &request.StartDate, &request.EndDate)
	if err != nil {
		// tx.Rollback()
		return
	}
	fmt.Println(request)
	var temp string
	res.Close()
	for res.Next() {
		err = res.Scan(&temp)
		if err != nil {
			log.Println(err)
		}

	}

	rows, err := tx.Query(`FETCH ALL IN "daily";`)
	for rows.Next() {
		var response models.ScanSummary
		rows.Scan(&response.DateString, &response.JumlahTrx, &response.TotalBruto, &response.TotalServiceFee, &response.TotalMDR, &response.TotalPaymentFee, &response.TotalVendorFee, &response.CountTunai, &response.CountLainnya, &response.CountCC, &response.CountDebit, &response.CountQRIS, &response.CountBrizzi, &response.CountEMoney, &response.CountTapCash, &response.CountFlazz, &response.CountJakcard, &response.CountVA, &response.CountBiller, &response.TotalTunai, &response.TotalLain, &response.TotalCC, &response.TotalDebit, &response.TotalQRIS, &response.TotalBrizzi, &response.TotalEMoney, &response.TotalTapCash, &response.TotalFlazz, &response.TotalJakcard, &response.TotalVA, &response.TotalBiller, &response.SFeeTunai, &response.SFeeLainnya, &response.SFeeCC, &response.SFeeDebit, &response.SFeeQRIS, &response.SFeeBrizzi, &response.SFeeEMoney, &response.SFeeTapCash, &response.SFeeFlazz, &response.SFeeJakcard, &response.SFeeVA, &response.SFeeBiller, &response.MdrTunai, &response.MdrLainnya, &response.MdrCC, &response.MdrDebit, &response.MdrQRIS, &response.MdrBrizzi, &response.MdrEMoney, &response.MdrTapCash, &response.MdrFlazz, &response.MdrJakcard, &response.MdrVA, &response.MdrBiller, &response.PayfeeTunai, &response.PayfeeLainnya, &response.PayfeeCC, &response.PayfeeDebit, &response.PayfeeQRIS, &response.PayfeeBrizzi, &response.PayfeeEMoney, &response.PayfeeTapCash, &response.PayfeeFlazz, &response.PayfeeJakcard, &response.PayfeeVA, &response.PayfeeBiller, &response.VendorfeeTunai, &response.VendorfeeLainnya, &response.VendorfeeCC, &response.VendorfeeDebit, &response.VendorfeeQRIS, &response.VendorfeeBrizzi, &response.VendorfeeEMoney, &response.VendorfeeTapCash, &response.VendorfeeFlazz, &response.VendorfeeJakcard, &response.VendorfeeVA, &response.VendorfeeBiller)

		// date, _ := time.Parse(time.RFC3339, response.Date)
		response.DateCreatedAt, _ = time.Parse("2006-01-02", response.DateString)
		// YYYYMMDD := "2006-01-02"
		hasil := models.ResponseSummaryOmset{
			Date:     response.DateCreatedAt.Day(),
			Omset:    response.TotalBruto,
			TrxCount: response.JumlahTrx,
		}

		data = append(data, hasil)

	}

	// tx.Commit()

	return
}

func (ctx dashboardRepository) GetDailyRevenue(request models.RequestRevenue, tx *sql.Tx) (data []models.ResponseDailyRevenue, err error) {
	var args []interface{}
	query := `SELECT public.fs_summary_trx(
		?, 
		(select current_date - interval '1 month' - interval '1 day')::text,
		(select current_date)::text);`
	args = append(args, request.Corporate)
	query = ReplaceSQL(query, "?")

	res, err := tx.Query(query, args...)
	if err != nil {
		log.Println(err)
	}
	// fmt.Println(res)
	var temp string
	res.Close()
	for res.Next() {
		err = res.Scan(&temp)
		if err != nil {
			log.Println(err)
		}

	}

	rows, err := tx.Query(`FETCH ALL IN "daily";`)
	for rows.Next() {
		var response models.ScanSummary
		rows.Scan(&response.DateString, &response.JumlahTrx, &response.TotalBruto, &response.TotalServiceFee, &response.TotalMDR, &response.TotalPaymentFee, &response.TotalVendorFee, &response.CountTunai, &response.CountLainnya, &response.CountCC, &response.CountDebit, &response.CountQRIS, &response.CountBrizzi, &response.CountEMoney, &response.CountTapCash, &response.CountFlazz, &response.CountJakcard, &response.CountVA, &response.CountBiller, &response.TotalTunai, &response.TotalLain, &response.TotalCC, &response.TotalDebit, &response.TotalQRIS, &response.TotalBrizzi, &response.TotalEMoney, &response.TotalTapCash, &response.TotalFlazz, &response.TotalJakcard, &response.TotalVA, &response.TotalBiller, &response.SFeeTunai, &response.SFeeLainnya, &response.SFeeCC, &response.SFeeDebit, &response.SFeeQRIS, &response.SFeeBrizzi, &response.SFeeEMoney, &response.SFeeTapCash, &response.SFeeFlazz, &response.SFeeJakcard, &response.SFeeVA, &response.SFeeBiller, &response.MdrTunai, &response.MdrLainnya, &response.MdrCC, &response.MdrDebit, &response.MdrQRIS, &response.MdrBrizzi, &response.MdrEMoney, &response.MdrTapCash, &response.MdrFlazz, &response.MdrJakcard, &response.MdrVA, &response.MdrBiller, &response.PayfeeTunai, &response.PayfeeLainnya, &response.PayfeeCC, &response.PayfeeDebit, &response.PayfeeQRIS, &response.PayfeeBrizzi, &response.PayfeeEMoney, &response.PayfeeTapCash, &response.PayfeeFlazz, &response.PayfeeJakcard, &response.PayfeeVA, &response.PayfeeBiller, &response.VendorfeeTunai, &response.VendorfeeLainnya, &response.VendorfeeCC, &response.VendorfeeDebit, &response.VendorfeeQRIS, &response.VendorfeeBrizzi, &response.VendorfeeEMoney, &response.VendorfeeTapCash, &response.VendorfeeFlazz, &response.VendorfeeJakcard, &response.VendorfeeVA, &response.VendorfeeBiller)
		// response.DateCreatedAt, _ = time.Parse("2006-01-02", response.DateString)
		// YYYYMMDD := "2006-01-02"
		hasil := models.ResponseDailyRevenue{
			Day:              response.DateString,
			TransactionCount: response.JumlahTrx,
			SumTransactions:  response.TotalBruto,
		}
		data = append(data, hasil)
	}
	if err != nil {
		log.Println(err)
	}

	return
}

func (ctx dashboardRepository) GetWeeklyRevenue(request models.RequestRevenue, tx *sql.Tx) (data []models.ResponseWeeklyRevenue, err error) {
	var args []interface{}
	query := `SELECT public.fs_summary_trx(
		?, 
		(SELECT date_trunc('month', current_date)::date)::text,
		(select current_date)::text);`
	args = append(args, request.Corporate)
	query = ReplaceSQL(query, "?")

	res, err := tx.Query(query, args...)
	if err != nil {
		log.Println(err)
	}
	var temp string
	res.Close()
	for res.Next() {
		err = res.Scan(&temp)
		if err != nil {
			log.Println(err)
		}

	}

	rows, err := tx.Query(`FETCH ALL IN "weekly";`)
	for rows.Next() {
		var response models.ScanSummary
		rows.Scan(&response.Date, &response.JumlahTrx, &response.TotalBruto, &response.TotalServiceFee, &response.TotalMDR, &response.TotalPaymentFee, &response.TotalVendorFee, &response.CountTunai, &response.CountLainnya, &response.CountCC, &response.CountDebit, &response.CountQRIS, &response.CountBrizzi, &response.CountEMoney, &response.CountTapCash, &response.CountFlazz, &response.CountJakcard, &response.CountVA, &response.CountBiller, &response.TotalTunai, &response.TotalLain, &response.TotalCC, &response.TotalDebit, &response.TotalQRIS, &response.TotalBrizzi, &response.TotalEMoney, &response.TotalTapCash, &response.TotalFlazz, &response.TotalJakcard, &response.TotalVA, &response.TotalBiller, &response.SFeeTunai, &response.SFeeLainnya, &response.SFeeCC, &response.SFeeDebit, &response.SFeeQRIS, &response.SFeeBrizzi, &response.SFeeEMoney, &response.SFeeTapCash, &response.SFeeFlazz, &response.SFeeJakcard, &response.SFeeVA, &response.SFeeBiller, &response.MdrTunai, &response.MdrLainnya, &response.MdrCC, &response.MdrDebit, &response.MdrQRIS, &response.MdrBrizzi, &response.MdrEMoney, &response.MdrTapCash, &response.MdrFlazz, &response.MdrJakcard, &response.MdrVA, &response.MdrBiller, &response.PayfeeTunai, &response.PayfeeLainnya, &response.PayfeeCC, &response.PayfeeDebit, &response.PayfeeQRIS, &response.PayfeeBrizzi, &response.PayfeeEMoney, &response.PayfeeTapCash, &response.PayfeeFlazz, &response.PayfeeJakcard, &response.PayfeeVA, &response.PayfeeBiller, &response.VendorfeeTunai, &response.VendorfeeLainnya, &response.VendorfeeCC, &response.VendorfeeDebit, &response.VendorfeeQRIS, &response.VendorfeeBrizzi, &response.VendorfeeEMoney, &response.VendorfeeTapCash, &response.VendorfeeFlazz, &response.VendorfeeJakcard, &response.VendorfeeVA, &response.VendorfeeBiller)

		hasil := models.ResponseWeeklyRevenue{
			Week:             ToString(response.Date),
			TransactionCount: response.JumlahTrx,
			SumTransactions:  response.TotalBruto,
		}
		data = append(data, hasil)

	}
	if err != nil {
		log.Println(err)
	}

	return
}

func (ctx dashboardRepository) GetMonthlyRevenue(request models.RequestRevenue, tx *sql.Tx) (data []models.ResponseMonthlyRevenue, err error) {
	var args []interface{}
	query := `SELECT public.fs_summary_trx(
		?, 
		(select current_date - interval '1 month' - interval '1 day')::text,
		(select current_date)::text);`
	args = append(args, request.Corporate)
	query = ReplaceSQL(query, "?")

	res, err := tx.Query(query, args...)
	if err != nil {
		log.Println(err)
	}
	var temp string
	res.Close()
	for res.Next() {
		err = res.Scan(&temp)
		if err != nil {
			log.Println(err)
		}

	}

	rows, err := tx.Query(`FETCH ALL IN "monthly";`)
	for rows.Next() {
		var response models.ScanSummary
		rows.Scan(&response.Date, &response.JumlahTrx, &response.TotalBruto, &response.TotalServiceFee, &response.TotalMDR, &response.TotalPaymentFee, &response.TotalVendorFee, &response.CountTunai, &response.CountLainnya, &response.CountCC, &response.CountDebit, &response.CountQRIS, &response.CountBrizzi, &response.CountEMoney, &response.CountTapCash, &response.CountFlazz, &response.CountJakcard, &response.CountVA, &response.CountBiller, &response.TotalTunai, &response.TotalLain, &response.TotalCC, &response.TotalDebit, &response.TotalQRIS, &response.TotalBrizzi, &response.TotalEMoney, &response.TotalTapCash, &response.TotalFlazz, &response.TotalJakcard, &response.TotalVA, &response.TotalBiller, &response.SFeeTunai, &response.SFeeLainnya, &response.SFeeCC, &response.SFeeDebit, &response.SFeeQRIS, &response.SFeeBrizzi, &response.SFeeEMoney, &response.SFeeTapCash, &response.SFeeFlazz, &response.SFeeJakcard, &response.SFeeVA, &response.SFeeBiller, &response.MdrTunai, &response.MdrLainnya, &response.MdrCC, &response.MdrDebit, &response.MdrQRIS, &response.MdrBrizzi, &response.MdrEMoney, &response.MdrTapCash, &response.MdrFlazz, &response.MdrJakcard, &response.MdrVA, &response.MdrBiller, &response.PayfeeTunai, &response.PayfeeLainnya, &response.PayfeeCC, &response.PayfeeDebit, &response.PayfeeQRIS, &response.PayfeeBrizzi, &response.PayfeeEMoney, &response.PayfeeTapCash, &response.PayfeeFlazz, &response.PayfeeJakcard, &response.PayfeeVA, &response.PayfeeBiller, &response.VendorfeeTunai, &response.VendorfeeLainnya, &response.VendorfeeCC, &response.VendorfeeDebit, &response.VendorfeeQRIS, &response.VendorfeeBrizzi, &response.VendorfeeEMoney, &response.VendorfeeTapCash, &response.VendorfeeFlazz, &response.VendorfeeJakcard, &response.VendorfeeVA, &response.VendorfeeBiller)
		hasil := models.ResponseMonthlyRevenue{
			Month:            ToString(response.Date),
			TransactionCount: response.JumlahTrx,
			SumTransactions:  response.TotalBruto,
		}
		data = append(data, hasil)

	}
	if err != nil {
		log.Println(err)
	}

	return
}

func (ctx dashboardRepository) GetYearlyRevenue(request models.RequestRevenue, tx *sql.Tx) (data []models.ResponseYearlyRevenue, err error) {
	var args []interface{}
	query := `SELECT public.fs_summary_trx(
		?, 
		(select current_date - interval '1 month' - interval '1 day')::text,
		(select current_date)::text);`
	args = append(args, request.Corporate)
	query = ReplaceSQL(query, "?")

	res, err := tx.Query(query, args...)
	if err != nil {
		log.Println(err)
	}
	var temp string
	res.Close()
	for res.Next() {
		err = res.Scan(&temp)
		if err != nil {
			log.Println(err)
		}

	}

	rows, err := tx.Query(`FETCH ALL IN "yearly";`)
	for rows.Next() {
		var response models.ScanSummary
		rows.Scan(&response.Date, &response.JumlahTrx, &response.TotalBruto, &response.TotalServiceFee, &response.TotalMDR, &response.TotalPaymentFee, &response.TotalVendorFee, &response.CountTunai, &response.CountLainnya, &response.CountCC, &response.CountDebit, &response.CountQRIS, &response.CountBrizzi, &response.CountEMoney, &response.CountTapCash, &response.CountFlazz, &response.CountJakcard, &response.CountVA, &response.CountBiller, &response.TotalTunai, &response.TotalLain, &response.TotalCC, &response.TotalDebit, &response.TotalQRIS, &response.TotalBrizzi, &response.TotalEMoney, &response.TotalTapCash, &response.TotalFlazz, &response.TotalJakcard, &response.TotalVA, &response.TotalBiller, &response.SFeeTunai, &response.SFeeLainnya, &response.SFeeCC, &response.SFeeDebit, &response.SFeeQRIS, &response.SFeeBrizzi, &response.SFeeEMoney, &response.SFeeTapCash, &response.SFeeFlazz, &response.SFeeJakcard, &response.SFeeVA, &response.SFeeBiller, &response.MdrTunai, &response.MdrLainnya, &response.MdrCC, &response.MdrDebit, &response.MdrQRIS, &response.MdrBrizzi, &response.MdrEMoney, &response.MdrTapCash, &response.MdrFlazz, &response.MdrJakcard, &response.MdrVA, &response.MdrBiller, &response.PayfeeTunai, &response.PayfeeLainnya, &response.PayfeeCC, &response.PayfeeDebit, &response.PayfeeQRIS, &response.PayfeeBrizzi, &response.PayfeeEMoney, &response.PayfeeTapCash, &response.PayfeeFlazz, &response.PayfeeJakcard, &response.PayfeeVA, &response.PayfeeBiller, &response.VendorfeeTunai, &response.VendorfeeLainnya, &response.VendorfeeCC, &response.VendorfeeDebit, &response.VendorfeeQRIS, &response.VendorfeeBrizzi, &response.VendorfeeEMoney, &response.VendorfeeTapCash, &response.VendorfeeFlazz, &response.VendorfeeJakcard, &response.VendorfeeVA, &response.VendorfeeBiller)
		// year, _ := response.Date.ISOWeek()
		// fmt.Println(response.Date)
		hasil := models.ResponseYearlyRevenue{
			Year:             ToString(response.Date),
			TransactionCount: response.JumlahTrx,
			SumTransactions:  response.TotalBruto,
		}
		data = append(data, hasil)

	}
	if err != nil {
		log.Println(err)
	}

	return
}

func DailyRevenueDto(rows *sql.Rows) (result []models.ResponseDailyRevenue, err error) {
	for rows.Next() {
		var dailyrevenue models.ResponseDailyRevenue
		err = rows.Scan(&dailyrevenue.Day, &dailyrevenue.TransactionCount, &dailyrevenue.SumTransactions)
		if err != nil {
			return
		}
		result = append(result, dailyrevenue)
	}

	return
}

func WeeklyRevenueDto(rows *sql.Rows) (result []models.ResponseWeeklyRevenue, err error) {
	for rows.Next() {
		var weeklyrevenue models.ResponseWeeklyRevenue
		err = rows.Scan(&weeklyrevenue.Week, &weeklyrevenue.TransactionCount, &weeklyrevenue.SumTransactions)
		if err != nil {
			return
		}
		result = append(result, weeklyrevenue)
	}

	return
}

func MonthlyRevenueDto(rows *sql.Rows) (result []models.ResponseMonthlyRevenue, err error) {
	for rows.Next() {
		var monthlyrevenue models.ResponseMonthlyRevenue
		err = rows.Scan(&monthlyrevenue.Month, &monthlyrevenue.TransactionCount, &monthlyrevenue.SumTransactions)
		if err != nil {
			return
		}
		result = append(result, monthlyrevenue)
	}

	return
}

func YearlyRevenueDto(rows *sql.Rows) (result []models.ResponseYearlyRevenue, err error) {
	for rows.Next() {
		var yearlyrevenue models.ResponseYearlyRevenue
		err = rows.Scan(&yearlyrevenue.Year, &yearlyrevenue.TransactionCount, &yearlyrevenue.SumTransactions)
		if err != nil {
			return
		}
		result = append(result, yearlyrevenue)
	}

	return
}
