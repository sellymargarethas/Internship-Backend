package summaryservice

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/services"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	. "Internship-Backend/utils"

	"github.com/labstack/echo"
)

type successSummaryService struct {
	Service services.UsecaseService
}

func NewSuccessSummaryService(service services.UsecaseService) successSummaryService {
	return successSummaryService{
		Service: service,
	}
}

func (svc successSummaryService) GetSuccessSummary(ctx echo.Context) (err error) {
	var result models.Response
	var request models.RequestSuccess
	var req models.RequestSummarySuccess
	var rescorporate string
	timeNow := time.Now()

	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data:  SummaryDashboard ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), constants.EMPTY_VALUE_INT)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	// log.Println("bind request success summary")
	if request.Corporate == constants.EMPTY_VALUE {
		request.Corporate = "1000000001"

	}
	req.StartDate = request.StartDate
	req.EndDate = request.EndDate
	req.Corporate = request.Corporate

	_, _ = time.LoadLocation("Asia/Bangkok")
	dateValuestart, _ := time.Parse("2006-01-02", request.StartDate) // convert 'String' to 'Time' data type
	dateValueend, _ := time.Parse("2006-01-02", request.EndDate)     // output: 2021-01-31 00:00:00 +0000 UTC

	request.StartDate = dateValuestart.Format("02 January 2006")
	request.EndDate = dateValueend.Format("02 January 2006") // Format return a 'string' in your specified layout (YYYY-MM-DD
	// fmt.Println(loc)

	reqcorporate := models.Corporate{
		HirarkiId: request.Corporate,
	}
	// fmt.Println(reqcorporate)
	// fmt.Println("Start: getcorporatename - success summary time start:", timeNow)
	if request.Corporate != "1000000001" {
		getcorp, err := svc.Service.CorporateRepo.GetSingleCorporate(reqcorporate)
		// fmt.Println(len(getcorp))
		if len(getcorp) == 0 {
			rescorporate = "CORPORATE NOT FOUND"
		} else {

			if err != nil {
				log.Println("ERROR GETTING CORPORATE NAME- settled summary:", err)
				rescorporate = "CORPORATE NOT FOUND"
			}
			rescorporate = getcorp[0].Uraian
		}

	} else {
		rescorporate = "Semua Corporate"
	}
	elapsed := time.Since(timeNow).Milliseconds()
	fmt.Println("Getting corporatename: success summary elapsed time:", elapsed)
	// fmt.Println("nama corporate: ", rescorporate)

	req.Corporate = req.Corporate + "%"

	err = DBTransaction(svc.Service.RepoDB, func(tx *sql.Tx) error {

		// fmt.Println(req)
		// fmt.Println("Start view success summary time now: ", timeNow)
		temp, err := svc.Service.SuccessSummaryRepo.ViewSuccessSummary(req, tx)
		if err != nil {
			// tx.Rollback()
			log.Println("error view success summary", err)
		}
		elapsed := time.Since(timeNow).Milliseconds()
		fmt.Println("ViewSuccessSummary time elapsed miliseconds:", elapsed)

		timeNow := time.Now()
		// fmt.Println("Start view success summary time now: ", timeNow)
		ressummary, err := svc.Service.SuccessSummaryRepo.ScanResultSuccess(temp, tx)

		// log.Println("get success summary")
		if err != nil {
			log.Println("Error Get SummaryDashboard : SummaryDashboard - Payment Method")
			result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "ERROR Get SummaryDashboard", nil)
			// tx.Rollback()
			return err
		}
		elapsed = time.Since(timeNow).Milliseconds()
		fmt.Println("Get success summary: scan result success elapsed time: ", elapsed)
		// tx.Commit()
		if request.KategoriPembayaran == constants.EMPTY_VALUE {

			resultSuccess := models.ResponseSummarySuccessTrx{
				StartDate:       request.StartDate,
				EndDate:         request.EndDate,
				Corporate:       rescorporate,
				Pembayaran:      "ALL",
				Status:          "SUKSES",
				JumlahTrx:       ressummary.JumlahTrx,
				TotalTunai:      ressummary.TotalTunai,
				TotalLain:       ressummary.TotalLain,
				TotalCC:         ressummary.TotalCC,
				TotalDebit:      ressummary.TotalDebit,
				TotalQRIS:       ressummary.TotalQRIS,
				TotalBrizzi:     ressummary.TotalBrizzi,
				TotalEMoney:     ressummary.TotalEMoney,
				TotalTapCash:    ressummary.TotalTapCash,
				TotalFlazz:      ressummary.TotalFlazz,
				TotalJakcard:    ressummary.TotalJakcard,
				TotalVA:         ressummary.TotalVA,
				TotalBiller:     ressummary.TotalBiller,
				TotalBruto:      ressummary.TotalBruto,
				TotalMDR:        ressummary.TotalMDR,
				TotalPotongan:   ressummary.TotalPotongan,
				TotalServiceFee: ressummary.TotalServiceFee,
				TotalPaymentFee: ressummary.TotalPaymentFee,
				TotalVendorFee:  ressummary.TotalVendorFee,
				TotalNett:       ressummary.TotalBruto - ressummary.TotalMDR - ressummary.TotalPaymentFee - ressummary.TotalVendorFee - ressummary.TotalServiceFee,
			}

			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}

		if request.KategoriPembayaran == "TUNAI" {
			resultSuccess := models.ResponseSummarySuccessTrx{
				StartDate:       request.StartDate,
				EndDate:         request.EndDate,
				Corporate:       rescorporate,
				Pembayaran:      "TUNAI",
				Status:          "SUKSES",
				JumlahTrx:       ressummary.CountTunai,
				TotalTunai:      ressummary.TotalTunai,
				TotalLain:       0,
				TotalCC:         0,
				TotalDebit:      0,
				TotalQRIS:       0,
				TotalBrizzi:     0,
				TotalEMoney:     0,
				TotalTapCash:    0,
				TotalFlazz:      0,
				TotalJakcard:    0,
				TotalVA:         0,
				TotalBiller:     0,
				TotalBruto:      ressummary.TotalTunai,
				TotalMDR:        0,
				TotalPotongan:   0,
				TotalServiceFee: ressummary.SFeeTunai,
				TotalPaymentFee: ressummary.PayfeeTunai,
				TotalVendorFee:  ressummary.VendorfeeTunai,
				TotalNett:       ressummary.TotalTunai - ressummary.MdrTunai - ressummary.PayfeeTunai - ressummary.SFeeTunai - ressummary.VendorfeeTunai,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}
		if request.KategoriPembayaran == "LAINNYA" {
			resultSuccess := models.ResponseSummarySuccessTrx{
				StartDate:       request.StartDate,
				EndDate:         request.EndDate,
				Corporate:       rescorporate,
				Pembayaran:      "LAINNYA",
				Status:          "SUKSES",
				JumlahTrx:       ressummary.CountLainnya,
				TotalTunai:      0,
				TotalLain:       ressummary.TotalLain,
				TotalCC:         0,
				TotalDebit:      0,
				TotalQRIS:       0,
				TotalBrizzi:     0,
				TotalEMoney:     0,
				TotalTapCash:    0,
				TotalFlazz:      0,
				TotalJakcard:    0,
				TotalVA:         0,
				TotalBiller:     0,
				TotalBruto:      ressummary.TotalLain,
				TotalMDR:        ressummary.MdrLainnya,
				TotalPotongan:   0,
				TotalServiceFee: ressummary.SFeeLainnya,
				TotalPaymentFee: ressummary.PayfeeLainnya,
				TotalVendorFee:  ressummary.VendorfeeLainnya,
				TotalNett:       ressummary.TotalLain - ressummary.MdrLainnya - ressummary.PayfeeLainnya - ressummary.SFeeLainnya - ressummary.VendorfeeLainnya,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}
		if request.KategoriPembayaran == "CC" {
			resultSuccess := models.ResponseSummarySuccessTrx{
				StartDate:       request.StartDate,
				EndDate:         request.EndDate,
				Corporate:       rescorporate,
				Pembayaran:      "CC",
				Status:          "SUKSES",
				JumlahTrx:       ressummary.CountCC,
				TotalTunai:      0,
				TotalLain:       0,
				TotalCC:         ressummary.TotalCC,
				TotalDebit:      0,
				TotalQRIS:       0,
				TotalBrizzi:     0,
				TotalEMoney:     0,
				TotalTapCash:    0,
				TotalFlazz:      0,
				TotalJakcard:    0,
				TotalVA:         0,
				TotalBiller:     0,
				TotalBruto:      0,
				TotalMDR:        ressummary.MdrCC,
				TotalPotongan:   0,
				TotalServiceFee: ressummary.SFeeCC,
				TotalPaymentFee: ressummary.PayfeeCC,
				TotalVendorFee:  ressummary.VendorfeeCC,
				TotalNett:       ressummary.TotalCC - ressummary.MdrCC - ressummary.PayfeeCC - ressummary.SFeeCC - ressummary.VendorfeeCC,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}
		if request.KategoriPembayaran == "DEBIT" {
			resultSuccess := models.ResponseSummarySuccessTrx{
				StartDate:       request.StartDate,
				EndDate:         request.EndDate,
				Corporate:       rescorporate,
				Pembayaran:      "DEBIT",
				Status:          "SUKSES",
				JumlahTrx:       ressummary.CountDebit,
				TotalTunai:      0,
				TotalLain:       0,
				TotalCC:         0,
				TotalDebit:      ressummary.TotalDebit,
				TotalQRIS:       0,
				TotalBrizzi:     0,
				TotalEMoney:     0,
				TotalTapCash:    0,
				TotalFlazz:      0,
				TotalJakcard:    0,
				TotalVA:         0,
				TotalBiller:     0,
				TotalBruto:      ressummary.TotalDebit,
				TotalMDR:        ressummary.MdrDebit,
				TotalPotongan:   0,
				TotalServiceFee: ressummary.SFeeDebit,
				TotalPaymentFee: ressummary.PayfeeDebit,
				TotalVendorFee:  ressummary.VendorfeeDebit,
				TotalNett:       ressummary.TotalDebit - ressummary.MdrDebit - ressummary.PayfeeDebit - ressummary.SFeeDebit - ressummary.VendorfeeDebit,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}
		if request.KategoriPembayaran == "QRIS" {
			resultSuccess := models.ResponseSummarySuccessTrx{
				StartDate:       request.StartDate,
				EndDate:         request.EndDate,
				Corporate:       rescorporate,
				Pembayaran:      "QRIS",
				Status:          "SUKSES",
				JumlahTrx:       ressummary.CountQRIS,
				TotalTunai:      0,
				TotalLain:       0,
				TotalCC:         0,
				TotalDebit:      0,
				TotalQRIS:       ressummary.TotalQRIS,
				TotalBrizzi:     0,
				TotalEMoney:     0,
				TotalTapCash:    0,
				TotalFlazz:      0,
				TotalJakcard:    0,
				TotalVA:         0,
				TotalBiller:     0,
				TotalBruto:      ressummary.TotalQRIS,
				TotalMDR:        ressummary.MdrQRIS,
				TotalPotongan:   0,
				TotalServiceFee: ressummary.SFeeQRIS,
				TotalPaymentFee: ressummary.PayfeeQRIS,
				TotalVendorFee:  ressummary.VendorfeeQRIS,
				TotalNett:       ressummary.TotalQRIS - ressummary.MdrQRIS - ressummary.PayfeeQRIS - ressummary.SFeeQRIS - ressummary.VendorfeeQRIS,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}
		if request.KategoriPembayaran == "BRIZZI" {
			resultSuccess := models.ResponseSummarySuccessTrx{
				StartDate:       request.StartDate,
				EndDate:         request.EndDate,
				Corporate:       rescorporate,
				Pembayaran:      "BRIZZI",
				Status:          "SUKSES",
				JumlahTrx:       ressummary.CountBrizzi,
				TotalTunai:      0,
				TotalLain:       0,
				TotalCC:         0,
				TotalDebit:      0,
				TotalQRIS:       0,
				TotalBrizzi:     ressummary.TotalBrizzi,
				TotalEMoney:     0,
				TotalTapCash:    0,
				TotalFlazz:      0,
				TotalJakcard:    0,
				TotalVA:         0,
				TotalBiller:     0,
				TotalBruto:      ressummary.TotalBrizzi,
				TotalMDR:        ressummary.MdrBrizzi,
				TotalPotongan:   0,
				TotalServiceFee: ressummary.SFeeBrizzi,
				TotalPaymentFee: ressummary.PayfeeBrizzi,
				TotalVendorFee:  ressummary.VendorfeeBrizzi,
				TotalNett:       ressummary.TotalBrizzi - ressummary.MdrBrizzi - ressummary.PayfeeBrizzi - ressummary.SFeeBrizzi - ressummary.VendorfeeBrizzi,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}
		if request.KategoriPembayaran == "E-MONEY" {
			resultSuccess := models.ResponseSummarySuccessTrx{
				StartDate:       request.StartDate,
				EndDate:         request.EndDate,
				Corporate:       rescorporate,
				Pembayaran:      "E-MONEY",
				Status:          "SUKSES",
				JumlahTrx:       ressummary.CountEMoney,
				TotalTunai:      0,
				TotalLain:       0,
				TotalCC:         0,
				TotalDebit:      0,
				TotalQRIS:       0,
				TotalBrizzi:     0,
				TotalEMoney:     ressummary.TotalEMoney,
				TotalTapCash:    0,
				TotalFlazz:      0,
				TotalJakcard:    0,
				TotalVA:         0,
				TotalBiller:     0,
				TotalBruto:      ressummary.TotalEMoney,
				TotalMDR:        ressummary.MdrEMoney,
				TotalPotongan:   0,
				TotalServiceFee: ressummary.SFeeEMoney,
				TotalPaymentFee: ressummary.PayfeeEMoney,
				TotalVendorFee:  ressummary.VendorfeeEMoney,
				TotalNett:       ressummary.TotalEMoney - ressummary.MdrEMoney - ressummary.PayfeeEMoney - ressummary.SFeeEMoney - ressummary.VendorfeeEMoney,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}
		if request.KategoriPembayaran == "TAPCASH" {
			resultSuccess := models.ResponseSummarySuccessTrx{
				StartDate:       request.StartDate,
				EndDate:         request.EndDate,
				Corporate:       rescorporate,
				Pembayaran:      "TAPCASH",
				Status:          "SUKSES",
				JumlahTrx:       ressummary.CountTapCash,
				TotalTunai:      0,
				TotalLain:       0,
				TotalCC:         0,
				TotalDebit:      0,
				TotalQRIS:       0,
				TotalBrizzi:     0,
				TotalEMoney:     0,
				TotalTapCash:    ressummary.TotalTapCash,
				TotalFlazz:      0,
				TotalJakcard:    0,
				TotalVA:         0,
				TotalBiller:     0,
				TotalBruto:      ressummary.TotalTapCash,
				TotalMDR:        ressummary.MdrTapCash,
				TotalPotongan:   0,
				TotalServiceFee: ressummary.SFeeTapCash,
				TotalPaymentFee: ressummary.PayfeeTapCash,
				TotalVendorFee:  ressummary.VendorfeeTapCash,
				TotalNett:       ressummary.TotalTapCash - ressummary.MdrTapCash - ressummary.PayfeeTapCash - ressummary.SFeeTapCash - ressummary.VendorfeeTapCash,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}
		if request.KategoriPembayaran == "FLAZZ" {
			resultSuccess := models.ResponseSummarySuccessTrx{
				StartDate:       request.StartDate,
				EndDate:         request.EndDate,
				Corporate:       rescorporate,
				Pembayaran:      "FLAZZ",
				Status:          "SUKSES",
				JumlahTrx:       ressummary.CountFlazz,
				TotalTunai:      0,
				TotalLain:       0,
				TotalCC:         0,
				TotalDebit:      0,
				TotalQRIS:       0,
				TotalBrizzi:     0,
				TotalEMoney:     0,
				TotalTapCash:    0,
				TotalFlazz:      ressummary.TotalFlazz,
				TotalJakcard:    0,
				TotalVA:         0,
				TotalBiller:     0,
				TotalBruto:      ressummary.TotalFlazz,
				TotalMDR:        ressummary.MdrFlazz,
				TotalPotongan:   0,
				TotalServiceFee: ressummary.SFeeFlazz,
				TotalPaymentFee: ressummary.PayfeeFlazz,
				TotalVendorFee:  ressummary.VendorfeeFlazz,
				TotalNett:       ressummary.TotalFlazz - ressummary.MdrFlazz - ressummary.PayfeeFlazz - ressummary.SFeeFlazz - ressummary.VendorfeeFlazz,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}
		if request.KategoriPembayaran == "JAKCARD" {
			resultSuccess := models.ResponseSummarySuccessTrx{
				StartDate:       request.StartDate,
				EndDate:         request.EndDate,
				Corporate:       rescorporate,
				Pembayaran:      "JAKCARD",
				Status:          "SUKSES",
				JumlahTrx:       ressummary.CountJakcard,
				TotalTunai:      0,
				TotalLain:       0,
				TotalCC:         0,
				TotalDebit:      0,
				TotalQRIS:       0,
				TotalBrizzi:     0,
				TotalEMoney:     0,
				TotalTapCash:    0,
				TotalFlazz:      0,
				TotalJakcard:    ressummary.TotalJakcard,
				TotalVA:         0,
				TotalBiller:     0,
				TotalBruto:      ressummary.TotalJakcard,
				TotalMDR:        ressummary.MdrJakcard,
				TotalPotongan:   0,
				TotalServiceFee: ressummary.SFeeJakcard,
				TotalPaymentFee: ressummary.PayfeeJakcard,
				TotalVendorFee:  ressummary.VendorfeeJakcard,
				TotalNett:       ressummary.TotalJakcard - ressummary.MdrJakcard - ressummary.PayfeeJakcard - ressummary.SFeeJakcard - ressummary.VendorfeeJakcard,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}
		if request.KategoriPembayaran == "VA" {
			resultSuccess := models.ResponseSummarySuccessTrx{
				StartDate:       request.StartDate,
				EndDate:         request.EndDate,
				Corporate:       rescorporate,
				Pembayaran:      "VA",
				Status:          "SUKSES",
				JumlahTrx:       ressummary.CountVA,
				TotalTunai:      0,
				TotalLain:       0,
				TotalCC:         0,
				TotalDebit:      0,
				TotalQRIS:       0,
				TotalBrizzi:     0,
				TotalEMoney:     0,
				TotalTapCash:    0,
				TotalFlazz:      0,
				TotalJakcard:    0,
				TotalVA:         ressummary.TotalVA,
				TotalBiller:     0,
				TotalBruto:      0,
				TotalMDR:        ressummary.MdrVA,
				TotalPotongan:   0,
				TotalServiceFee: ressummary.SFeeVA,
				TotalPaymentFee: ressummary.PayfeeVA,
				TotalVendorFee:  ressummary.VendorfeeVA,
				TotalNett:       ressummary.TotalVA - ressummary.MdrVA - ressummary.PayfeeVA - ressummary.SFeeVA - ressummary.VendorfeeVA,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}
		if request.KategoriPembayaran == "BILLER" {
			resultSuccess := models.ResponseSummarySuccessTrx{
				StartDate:       request.StartDate,
				EndDate:         request.EndDate,
				Corporate:       rescorporate,
				Pembayaran:      "BILLER",
				Status:          "SUKSES",
				JumlahTrx:       ressummary.CountBiller,
				TotalTunai:      0,
				TotalLain:       0,
				TotalCC:         0,
				TotalDebit:      0,
				TotalQRIS:       0,
				TotalBrizzi:     0,
				TotalEMoney:     0,
				TotalTapCash:    0,
				TotalFlazz:      0,
				TotalJakcard:    0,
				TotalVA:         0,
				TotalBiller:     ressummary.TotalBiller,
				TotalBruto:      ressummary.TotalBiller,
				TotalMDR:        ressummary.MdrBiller,
				TotalPotongan:   0,
				TotalServiceFee: ressummary.SFeeBiller,
				TotalPaymentFee: ressummary.PayfeeBiller,
				TotalVendorFee:  ressummary.VendorfeeBiller,
				TotalNett:       ressummary.TotalBiller - ressummary.MdrBiller - ressummary.PayfeeBiller - ressummary.SFeeBiller - ressummary.VendorfeeBiller,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		} else {
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", "Invalid Request")
			return err
		}

		// result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", ressummary)
		// return err
	})
	if err != nil {
		log.Println("ERROR db transaction: summary success: ", err)

		return ctx.JSON(http.StatusOK, result)
	}
	log.Println("Success Get SummaryDashboard : SummaryDashboard")
	return ctx.JSON(http.StatusOK, result)

}

func (svc successSummaryService) GetSettledSummary(ctx echo.Context) (err error) {
	var result models.Response
	var request models.RequestSettled
	var req models.RequestSummarySettled
	var settlementDestination string
	var rescorporate string
	timeNow := time.Now()
	if err := BindValidateStruct(ctx, &request); err != nil {
		log.Println("Error Validate Data:  SettledSummaryDashboard ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), constants.EMPTY_VALUE_INT)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	fmt.Println(&request)
	log.Println("Bind request settled summary")
	fmt.Println("Bind request time start :", timeNow)

	if request.Corporate == constants.EMPTY_VALUE {
		request.Corporate = "1000000001"
	}
	req.StartDate = request.StartDate
	req.EndDate = request.EndDate
	req.Corporate = request.Corporate
	req.SettlementDestination = request.SettlementDestination

	//cek settlement destination
	if len(request.SettlementDestination) == constants.EMPTY_VALUE_INT {
		listCSK, err := svc.Service.CoreSettlementRepo.GetListCSKId()
		if err != nil {
			log.Println("Error getting all payment category", err)
		}
		req.SettlementDestination = listCSK
		settlementDestination = "ALL"
		elapsed := time.Since(timeNow).Milliseconds()
		fmt.Println("GetSettledSummary - GetListCSKId elapsed time millisecond :", elapsed)
	} else {
		settlementDestination, err = svc.Service.CoreSettlementRepo.GetStrAggCSKName(req)
		if err != nil {
			log.Println("Error getting all Settlement Destination", err)
		}

	}

	loc, _ := time.LoadLocation("Asia/Bangkok")
	dateValuestart, _ := time.Parse("2006-01-02", request.StartDate) // convert 'String' to 'Time' data type
	dateValueend, _ := time.Parse("2006-01-02", request.EndDate)     // output: 2021-01-31 00:00:00 +0000 UTC
	request.StartDate = dateValuestart.Format("02 January 2006")
	request.EndDate = dateValueend.Format("02 January 2006") // Format return a 'string' in your specified layout (YYYY-MM-DD)
	fmt.Println(loc)

	err = DBTransaction(svc.Service.RepoDB, func(tx *sql.Tx) error {
		if request.KategoriPembayaran == constants.EMPTY_VALUE {
			// kategoriPembayaran, err := svc.Service.KategoriPembayaranRepo.GetListKategoriPembayaranId()
			// if err != nil {
			// 	log.Println("Error getting all payment category", err)
			// }
			// req.KategoriPembayaran = kategoriPembayaran
			reqcorporate := models.Corporate{
				HirarkiId: request.Corporate,
			}
			ressummary, err := svc.Service.SuccessSummaryRepo.ViewSettledSummary(req, tx)

			if err != nil {
				log.Println("error view settled summary", err)
			}
			elapsed := time.Since(timeNow).Milliseconds()
			fmt.Println("GetSettledSummary - ViewSettledSummary elapsed time millisecond :", elapsed)

			elapsed2 := time.Since(timeNow).Milliseconds()
			fmt.Println("GetSettledSummary - ScanResultSettled elapsed time millisecond :", elapsed2)
			log.Println("view settled summary")
			if request.Corporate != "1000000001" {
				getcorp, err := svc.Service.CorporateRepo.GetSingleCorporate(reqcorporate)
				elapsed3 := time.Since(timeNow).Milliseconds()
				fmt.Println("GetSettledSummary - GetSingleCorporate elapsed time millisecond :", elapsed3)
				log.Println("get single corporate")
				if len(getcorp) == 0 {
					rescorporate = "CORPORATE NOT FOUND"
				} else {
					if err != nil {
						log.Println("ERROR GETTING CORPORATE NAME- settled summary:", err)
					}
					rescorporate = getcorp[0].Uraian
				}
			} else {
				rescorporate = "Semua Corporate"
			}
			if err != nil {
				log.Println("Error Get SettledSummaryDashboard : SettledSummaryDashboard - All payment method", err)

				result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "ERROR Get SettledSummaryDashboard", nil)
				return err
			}

			// ressummary.StartDate = request.StartDate
			// ressummary.EndDate = request.EndDate
			// ressummary.Pembayaran = "ALL"
			// ressummary.Status = "Sukses"
			// ressummary.SettlementDestination = settlementDestination
			// ressummary.TotalNett = ressummary.TotalBruto - ressummary.TotalMDR - ressummary.TotalPaymentFee - ressummary.TotalVendorFee - ressummary.TotalServiceFee - ressummary.TotalPaymentFee
			if request.KategoriPembayaran == constants.EMPTY_VALUE {
				resultSuccess := models.ResponseSummarySettledTrx{
					StartDate:             request.StartDate,
					EndDate:               request.EndDate,
					Corporate:             rescorporate,
					Pembayaran:            "ALL",
					Status:                "SUKSES",
					SettlementDestination: settlementDestination,
					JumlahTrx:             ressummary.JumlahTrx,
					TotalTunai:            ressummary.TotalTunai,
					TotalLain:             ressummary.TotalLain,
					TotalCC:               ressummary.TotalCC,
					TotalDebit:            ressummary.TotalDebit,
					TotalQRIS:             ressummary.TotalQRIS,
					TotalBrizzi:           ressummary.TotalBrizzi,
					TotalEMoney:           ressummary.TotalEMoney,
					TotalTapCash:          ressummary.TotalTapCash,
					TotalFlazz:            ressummary.TotalFlazz,
					TotalJakcard:          ressummary.TotalJakcard,
					TotalVA:               ressummary.TotalVA,
					TotalBiller:           ressummary.TotalBiller,
					TotalBruto:            ressummary.TotalBruto,
					TotalMDR:              ressummary.TotalMDR,
					TotalPotongan:         ressummary.TotalPotongan,
					TotalServiceFee:       ressummary.TotalServiceFee,
					TotalPaymentFee:       ressummary.TotalPaymentFee,
					TotalVendorFee:        ressummary.TotalVendorFee,
					TotalNett:             ressummary.TotalBruto - ressummary.TotalMDR - ressummary.TotalPaymentFee - ressummary.TotalServiceFee - ressummary.TotalVendorFee,
				}
				result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
				return err
			}

			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get settledSummaryDashboard", ressummary)
			return err
		}

		reqcategory := models.KategoriPembayaran{
			Uraian: request.KategoriPembayaran,
		}

		categoryid, _, err := svc.Service.KategoriPembayaranRepo.IsKategoriPembayaranExistsByIndex(reqcategory)

		if err != nil {
			log.Println("ERROR getting summary success:")
		}
		elapsed4 := time.Since(timeNow).Milliseconds()
		fmt.Println("GetSettledSummary - IsKategoriPembayaranExistsByIndex elapsed time millisecond :", elapsed4)
		fmt.Println(categoryid)

		// data := models.KategoriPembayaranId{
		// 	IdKategoriPembayaran: categoryid.Id,
		// }

		// req.KategoriPembayaran = append(req.KategoriPembayaran, data)

		ressummary, err := svc.Service.SuccessSummaryRepo.ViewSettledSummary(req, tx)

		if err != nil {
			log.Println("error view settled summary", err)
		}
		elapsed5 := time.Since(timeNow).Milliseconds()
		fmt.Println("GetSettledSummary - ViewSettledSummary (2) elapsed time millisecond :", elapsed5)

		elapsed6 := time.Since(timeNow).Milliseconds()
		fmt.Println("GetSettledSummary - ScanResultSettled (2) elapsed time millisecond :", elapsed6)
		log.Println("settled summary success")
		reqcorporate := models.Corporate{
			HirarkiId: request.Corporate,
		}
		if request.Corporate != "1000000001" {
			getcorp, err := svc.Service.CorporateRepo.GetSingleCorporate(reqcorporate)
			if err != nil {
				log.Println("ERROR GETTING CORPORATE NAME- settled summary:", err)
			}
			elapsed7 := time.Since(timeNow).Milliseconds()
			fmt.Println("GetSettledSummary - GetSingleCorporate (2) elapsed time millisecond :", elapsed7)
			rescorporate = getcorp[0].Uraian
		} else {
			rescorporate = "Semua Corporate"
		}
		if err != nil {
			log.Println("Error Get SettledSummaryDashboard-ViewSettledSummary", err)
			if request.KategoriPembayaran == constants.EMPTY_VALUE {
				request.KategoriPembayaran = "SEMUA"
			}
			resultSuccess := models.ResponseSummarySettledTrx{
				StartDate:             request.StartDate,
				EndDate:               request.EndDate,
				Corporate:             rescorporate,
				Pembayaran:            request.KategoriPembayaran,
				Status:                "SUKSES",
				SettlementDestination: settlementDestination,
				JumlahTrx:             0,
				TotalBruto:            0,
				TotalMDR:              0,
				TotalServiceFee:       0,
				TotalPaymentFee:       0,
				TotalVendorFee:        0,
				TotalTunai:            0,
				TotalLain:             0,
				TotalCC:               0,
				TotalDebit:            0,
				TotalQRIS:             0,
				TotalBrizzi:           0,
				TotalEMoney:           0,
				TotalTapCash:          0,
				TotalFlazz:            0,
				TotalJakcard:          0,
				TotalVA:               0,
				TotalBiller:           0,
				TotalPotongan:         0,
				TotalNett:             0,
			}

			result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "ERROR Get SettledSummaryDashboard", resultSuccess)
			return err
		}

		reqcorporate = models.Corporate{
			HirarkiId: request.Corporate,
		}

		// if request.Corporate != "1000000001" {
		// 	getcorp, err := svc.Service.CorporateRepo.GetSingleCorporate(reqcorporate)
		// 	if err != nil {
		// 		log.Println("ERROR GETTING CORPORATE NAME- settled summary:", err)
		// 	}
		// 	rescorporate = getcorp[0].Uraian
		// } else {
		// 	rescorporate = "Semua Corporate"
		// }

		// getcorp, err := svc.Service.CorporateRepo.GetSingleCorporate(reqcorporate)
		// if err != nil {
		// 	log.Println("ERROR GETTING CORPORATE NAME- settled summary:", err)
		// }

		// ressummary.StartDate = request.StartDate
		// ressummary.EndDate = request.EndDate
		// rescorporate = getcorp[0].Uraian
		// ressummary.Pembayaran = request.KategoriPembayaran
		// ressummary.Status = "Sukses"
		// ressummary.SettlementDestination = settlementDestination
		// ressummary.TotalNett = ressummary.TotalBruto - ressummary.TotalMDR - ressummary.TotalPaymentFee - ressummary.TotalVendorFee - ressummary.TotalServiceFee - ressummary.TotalPaymentFee

		if request.KategoriPembayaran == "TUNAI" {
			resultSuccess := models.ResponseSummarySettledTrx{
				StartDate:             request.StartDate,
				EndDate:               request.EndDate,
				Corporate:             rescorporate,
				Pembayaran:            "TUNAI",
				Status:                "SUKSES",
				SettlementDestination: settlementDestination,
				JumlahTrx:             ressummary.CountTunai,
				TotalTunai:            ressummary.TotalTunai,
				TotalLain:             0,
				TotalCC:               0,
				TotalDebit:            0,
				TotalQRIS:             0,
				TotalBrizzi:           0,
				TotalEMoney:           0,
				TotalTapCash:          0,
				TotalFlazz:            0,
				TotalJakcard:          0,
				TotalVA:               0,
				TotalBiller:           0,
				TotalBruto:            ressummary.TotalTunai,
				TotalMDR:              ressummary.MdrTunai,
				TotalPotongan:         0,
				TotalServiceFee:       ressummary.SFeeTunai,
				TotalPaymentFee:       ressummary.PayfeeTunai,
				TotalVendorFee:        ressummary.VendorfeeTunai,
				TotalNett:             ressummary.TotalTunai - ressummary.MdrTunai - ressummary.PayfeeTunai - ressummary.SFeeTunai - ressummary.VendorfeeTunai,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}

		if request.KategoriPembayaran == "LAINNYA" {
			resultSuccess := models.ResponseSummarySettledTrx{
				StartDate:             request.StartDate,
				EndDate:               request.EndDate,
				Corporate:             rescorporate,
				Pembayaran:            "LAINNYA",
				Status:                "SUKSES",
				SettlementDestination: settlementDestination,
				JumlahTrx:             ressummary.CountLainnya,
				TotalTunai:            0,
				TotalLain:             ressummary.TotalLain,
				TotalCC:               0,
				TotalDebit:            0,
				TotalQRIS:             0,
				TotalBrizzi:           0,
				TotalEMoney:           0,
				TotalTapCash:          0,
				TotalFlazz:            0,
				TotalJakcard:          0,
				TotalVA:               0,
				TotalBiller:           0,
				TotalBruto:            ressummary.TotalLain,
				TotalMDR:              ressummary.MdrLainnya,
				TotalPotongan:         0,
				TotalServiceFee:       ressummary.SFeeLainnya,
				TotalPaymentFee:       ressummary.PayfeeLainnya,
				TotalVendorFee:        ressummary.VendorfeeLainnya,
				TotalNett:             ressummary.TotalLain - ressummary.MdrLainnya - ressummary.PayfeeLainnya - ressummary.SFeeLainnya - ressummary.VendorfeeLainnya,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}

		if request.KategoriPembayaran == "CC" {
			resultSuccess := models.ResponseSummarySettledTrx{
				StartDate:             request.StartDate,
				EndDate:               request.EndDate,
				Corporate:             rescorporate,
				Pembayaran:            "CC",
				Status:                "SUKSES",
				SettlementDestination: settlementDestination,
				JumlahTrx:             ressummary.CountCC,
				TotalTunai:            0,
				TotalLain:             0,
				TotalCC:               ressummary.TotalCC,
				TotalDebit:            0,
				TotalQRIS:             0,
				TotalBrizzi:           0,
				TotalEMoney:           0,
				TotalTapCash:          0,
				TotalFlazz:            0,
				TotalJakcard:          0,
				TotalVA:               0,
				TotalBiller:           0,
				TotalBruto:            ressummary.TotalCC,
				TotalMDR:              ressummary.MdrCC,
				TotalPotongan:         0,
				TotalServiceFee:       ressummary.SFeeCC,
				TotalPaymentFee:       ressummary.PayfeeCC,
				TotalVendorFee:        ressummary.VendorfeeCC,
				TotalNett:             ressummary.TotalCC - ressummary.MdrCC - ressummary.PayfeeCC - ressummary.SFeeCC - ressummary.VendorfeeCC,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}

		if request.KategoriPembayaran == "DEBIT" {
			resultSuccess := models.ResponseSummarySettledTrx{
				StartDate:             request.StartDate,
				EndDate:               request.EndDate,
				Corporate:             rescorporate,
				Pembayaran:            "DEBIT",
				Status:                "SUKSES",
				SettlementDestination: settlementDestination,
				JumlahTrx:             ressummary.CountDebit,
				TotalTunai:            0,
				TotalLain:             0,
				TotalCC:               0,
				TotalDebit:            ressummary.TotalDebit,
				TotalQRIS:             0,
				TotalBrizzi:           0,
				TotalEMoney:           0,
				TotalTapCash:          0,
				TotalFlazz:            0,
				TotalJakcard:          0,
				TotalVA:               0,
				TotalBiller:           0,
				TotalBruto:            ressummary.TotalDebit,
				TotalMDR:              ressummary.MdrDebit,
				TotalPotongan:         0,
				TotalServiceFee:       ressummary.SFeeDebit,
				TotalPaymentFee:       ressummary.PayfeeDebit,
				TotalVendorFee:        ressummary.VendorfeeDebit,
				TotalNett:             ressummary.TotalDebit - ressummary.MdrDebit - ressummary.PayfeeDebit - ressummary.SFeeDebit - ressummary.VendorfeeDebit,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}

		if request.KategoriPembayaran == "QRIS" {
			resultSuccess := models.ResponseSummarySettledTrx{
				StartDate:             request.StartDate,
				EndDate:               request.EndDate,
				Corporate:             rescorporate,
				Pembayaran:            "QRIS",
				Status:                "SUKSES",
				SettlementDestination: settlementDestination,
				JumlahTrx:             ressummary.CountQRIS,
				TotalTunai:            0,
				TotalLain:             0,
				TotalCC:               0,
				TotalDebit:            0,
				TotalQRIS:             ressummary.TotalQRIS,
				TotalBrizzi:           0,
				TotalEMoney:           0,
				TotalTapCash:          0,
				TotalFlazz:            0,
				TotalJakcard:          0,
				TotalVA:               0,
				TotalBiller:           0,
				TotalBruto:            ressummary.TotalQRIS,
				TotalMDR:              ressummary.MdrQRIS,
				TotalPotongan:         0,
				TotalServiceFee:       ressummary.SFeeQRIS,
				TotalPaymentFee:       ressummary.PayfeeQRIS,
				TotalVendorFee:        ressummary.VendorfeeQRIS,
				TotalNett:             ressummary.TotalQRIS - ressummary.MdrQRIS - ressummary.PayfeeQRIS - ressummary.SFeeQRIS - ressummary.VendorfeeQRIS,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}

		if request.KategoriPembayaran == "BRIZZI" {
			resultSuccess := models.ResponseSummarySettledTrx{
				StartDate:             request.StartDate,
				EndDate:               request.EndDate,
				Corporate:             rescorporate,
				Pembayaran:            "BRIZZI",
				Status:                "SUKSES",
				SettlementDestination: settlementDestination,
				JumlahTrx:             ressummary.CountBrizzi,
				TotalTunai:            0,
				TotalLain:             0,
				TotalCC:               0,
				TotalDebit:            0,
				TotalQRIS:             0,
				TotalBrizzi:           ressummary.TotalBrizzi,
				TotalEMoney:           0,
				TotalTapCash:          0,
				TotalFlazz:            0,
				TotalJakcard:          0,
				TotalVA:               0,
				TotalBiller:           0,
				TotalBruto:            ressummary.TotalBrizzi,
				TotalMDR:              ressummary.MdrBrizzi,
				TotalPotongan:         0,
				TotalServiceFee:       ressummary.SFeeBrizzi,
				TotalPaymentFee:       ressummary.PayfeeBrizzi,
				TotalVendorFee:        ressummary.VendorfeeBrizzi,
				TotalNett:             ressummary.TotalBrizzi - ressummary.MdrBrizzi - ressummary.PayfeeBrizzi - ressummary.SFeeBrizzi - ressummary.VendorfeeBrizzi,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}

		if request.KategoriPembayaran == "E-MONEY" {
			resultSuccess := models.ResponseSummarySettledTrx{
				StartDate:             request.StartDate,
				EndDate:               request.EndDate,
				Corporate:             rescorporate,
				Pembayaran:            "E-MONEY",
				Status:                "SUKSES",
				SettlementDestination: settlementDestination,
				JumlahTrx:             ressummary.CountEMoney,
				TotalTunai:            0,
				TotalLain:             0,
				TotalCC:               0,
				TotalDebit:            0,
				TotalQRIS:             0,
				TotalBrizzi:           0,
				TotalEMoney:           ressummary.TotalEMoney,
				TotalTapCash:          0,
				TotalFlazz:            0,
				TotalJakcard:          0,
				TotalVA:               0,
				TotalBiller:           0,
				TotalBruto:            ressummary.TotalEMoney,
				TotalMDR:              ressummary.MdrEMoney,
				TotalPotongan:         0,
				TotalServiceFee:       ressummary.SFeeEMoney,
				TotalPaymentFee:       ressummary.PayfeeEMoney,
				TotalVendorFee:        ressummary.VendorfeeEMoney,
				TotalNett:             ressummary.TotalEMoney - ressummary.MdrEMoney - ressummary.PayfeeEMoney - ressummary.SFeeEMoney - ressummary.VendorfeeEMoney,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}

		if request.KategoriPembayaran == "TAPCASH" {
			resultSuccess := models.ResponseSummarySettledTrx{
				StartDate:             request.StartDate,
				EndDate:               request.EndDate,
				Corporate:             rescorporate,
				Pembayaran:            "TAPCASH",
				Status:                "SUKSES",
				SettlementDestination: settlementDestination,
				JumlahTrx:             ressummary.CountTapCash,
				TotalTunai:            0,
				TotalLain:             0,
				TotalCC:               0,
				TotalDebit:            0,
				TotalQRIS:             0,
				TotalBrizzi:           0,
				TotalEMoney:           0,
				TotalTapCash:          ressummary.TotalTapCash,
				TotalFlazz:            0,
				TotalJakcard:          0,
				TotalVA:               0,
				TotalBiller:           0,
				TotalBruto:            ressummary.TotalTapCash,
				TotalMDR:              ressummary.MdrTapCash,
				TotalPotongan:         0,
				TotalServiceFee:       ressummary.SFeeTapCash,
				TotalPaymentFee:       ressummary.PayfeeTapCash,
				TotalVendorFee:        ressummary.VendorfeeTapCash,
				TotalNett:             ressummary.TotalTapCash - ressummary.MdrTapCash - ressummary.PayfeeTapCash - ressummary.SFeeTapCash - ressummary.VendorfeeTapCash,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}

		if request.KategoriPembayaran == "FLAZZ" {
			resultSuccess := models.ResponseSummarySettledTrx{
				StartDate:             request.StartDate,
				EndDate:               request.EndDate,
				Corporate:             rescorporate,
				Pembayaran:            "FLAZZ",
				Status:                "SUKSES",
				SettlementDestination: settlementDestination,
				JumlahTrx:             ressummary.CountFlazz,
				TotalTunai:            0,
				TotalLain:             0,
				TotalCC:               0,
				TotalDebit:            0,
				TotalQRIS:             0,
				TotalBrizzi:           0,
				TotalEMoney:           0,
				TotalTapCash:          0,
				TotalFlazz:            ressummary.TotalFlazz,
				TotalJakcard:          0,
				TotalVA:               0,
				TotalBiller:           0,
				TotalBruto:            ressummary.TotalFlazz,
				TotalMDR:              ressummary.MdrFlazz,
				TotalPotongan:         0,
				TotalServiceFee:       ressummary.SFeeFlazz,
				TotalPaymentFee:       ressummary.PayfeeFlazz,
				TotalVendorFee:        ressummary.VendorfeeFlazz,
				TotalNett:             ressummary.TotalFlazz - ressummary.MdrFlazz - ressummary.PayfeeFlazz - ressummary.SFeeFlazz - ressummary.VendorfeeFlazz,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}

		if request.KategoriPembayaran == "JAKCARD" {
			resultSuccess := models.ResponseSummarySettledTrx{
				StartDate:             request.StartDate,
				EndDate:               request.EndDate,
				Corporate:             rescorporate,
				Pembayaran:            "JAKCARD",
				Status:                "SUKSES",
				SettlementDestination: settlementDestination,
				JumlahTrx:             ressummary.CountJakcard,
				TotalTunai:            0,
				TotalLain:             0,
				TotalCC:               0,
				TotalDebit:            0,
				TotalQRIS:             0,
				TotalBrizzi:           0,
				TotalEMoney:           0,
				TotalTapCash:          0,
				TotalFlazz:            0,
				TotalJakcard:          ressummary.TotalJakcard,
				TotalVA:               0,
				TotalBiller:           0,
				TotalBruto:            ressummary.TotalJakcard,
				TotalMDR:              ressummary.MdrJakcard,
				TotalPotongan:         0,
				TotalServiceFee:       ressummary.SFeeJakcard,
				TotalPaymentFee:       ressummary.PayfeeJakcard,
				TotalVendorFee:        ressummary.VendorfeeJakcard,
				TotalNett:             ressummary.TotalJakcard - ressummary.MdrJakcard - ressummary.PayfeeJakcard - ressummary.SFeeJakcard - ressummary.VendorfeeJakcard,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}

		if request.KategoriPembayaran == "VA" {
			resultSuccess := models.ResponseSummarySettledTrx{
				StartDate:             request.StartDate,
				EndDate:               request.EndDate,
				Corporate:             rescorporate,
				Pembayaran:            "VA",
				Status:                "SUKSES",
				SettlementDestination: settlementDestination,
				JumlahTrx:             ressummary.CountVA,
				TotalTunai:            0,
				TotalLain:             0,
				TotalCC:               0,
				TotalDebit:            0,
				TotalQRIS:             0,
				TotalBrizzi:           0,
				TotalEMoney:           0,
				TotalTapCash:          0,
				TotalFlazz:            0,
				TotalJakcard:          0,
				TotalVA:               ressummary.TotalVA,
				TotalBiller:           0,
				TotalBruto:            ressummary.TotalVA,
				TotalMDR:              ressummary.MdrVA,
				TotalPotongan:         0,
				TotalServiceFee:       ressummary.SFeeVA,
				TotalPaymentFee:       ressummary.PayfeeVA,
				TotalVendorFee:        ressummary.VendorfeeVA,
				TotalNett:             ressummary.TotalVA - ressummary.MdrVA - ressummary.PayfeeVA - ressummary.SFeeVA - ressummary.VendorfeeVA,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}

		if request.KategoriPembayaran == "BILLER" {
			resultSuccess := models.ResponseSummarySettledTrx{
				StartDate:             request.StartDate,
				EndDate:               request.EndDate,
				Corporate:             rescorporate,
				Pembayaran:            "BILLER",
				Status:                "SUKSES",
				SettlementDestination: settlementDestination,
				JumlahTrx:             ressummary.CountBiller,
				TotalTunai:            0,
				TotalLain:             0,
				TotalCC:               0,
				TotalDebit:            0,
				TotalQRIS:             0,
				TotalBrizzi:           0,
				TotalEMoney:           0,
				TotalTapCash:          0,
				TotalFlazz:            0,
				TotalJakcard:          0,
				TotalVA:               0,
				TotalBiller:           ressummary.TotalBiller,
				TotalBruto:            ressummary.TotalBiller,
				TotalMDR:              ressummary.MdrBiller,
				TotalPotongan:         0,
				TotalServiceFee:       ressummary.SFeeBiller,
				TotalPaymentFee:       ressummary.PayfeeBiller,
				TotalVendorFee:        ressummary.VendorfeeBiller,
				TotalNett:             ressummary.TotalBiller - ressummary.MdrBiller - ressummary.PayfeeBiller - ressummary.SFeeBiller - ressummary.VendorfeeBiller,
			}
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", resultSuccess)
			return err
		}
		log.Println("selesai settled summary")
		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", nil)
		return err
	})

	if err != nil {
		log.Println("ERROR db transaction: settled success: ", err)
		return ctx.JSON(http.StatusOK, result)
	}
	log.Println("SUCCESS Get SettledDashboard : SettledDashboard", err)
	return ctx.JSON(http.StatusOK, result)

}
