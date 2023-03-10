package transactionService

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/services"
	"log"
	"net/http"

	. "Internship-Backend/utils"

	"github.com/labstack/echo"
)

type listtransactionService struct {
	Service services.UsecaseService
}

func NewListTransactionService(service services.UsecaseService) listtransactionService {
	return listtransactionService{
		Service: service,
	}
}

func (svc listtransactionService) CountTransaction(ctx echo.Context) error {
	var result models.Response
	var request models.RequestListTrx

	err := BindValidateStruct(ctx, &request)
	if err != nil {
		log.Println("Error Validate Data CountTransaction : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data: Count Transaction", constants.EMPTY_VALUE_INT)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	request.Corporate = request.Corporate + "%"
	count, err := svc.Service.ListTransactionRepo.CountListTransaction(request)
	if err != nil {
		log.Println("Error CountTransaction - CountTransaction : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Count Transaction", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse CountTransaction - CountTransaction : ", count)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Count Transaction", count)
	return ctx.JSON(http.StatusOK, result)
}
func (svc listtransactionService) CountTransactionSettled(ctx echo.Context) error {
	var result models.Response
	var request models.RequestListTrx

	err := BindValidateStruct(ctx, &request)
	if err != nil {
		log.Println("Error Validate Data CountTransactionSettled : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data: Count Transaction", constants.EMPTY_VALUE_INT)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	request.Corporate = request.Corporate + "%"
	count, err := svc.Service.ListTransactionRepo.CountListTransactionSettled(request)
	if err != nil {
		log.Println("Error CountTransactionSettled - CountTransactionSettled : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Count Transaction", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse CountTransactionSettled - CountTransactionSettled : ", count)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Count Transaction", count)
	return ctx.JSON(http.StatusOK, result)
}

// Get All Transaction
func (svc listtransactionService) GetListTransaction(ctx echo.Context) error {
	var result models.ResponseList
	var request models.RequestListTrx

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data GetListTransaction : ", err.Error())
		result = ResponseListJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", constants.EMPTY_VALUE_INT, err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}
	var orderby string

	if request.Pagination.OrderBy == "id" {
		orderby = "trx_v2.id"
	}

	if request.Pagination.OrderBy == "nomorHeader" {
		orderby = "trx_v2.noheader"
	}
	if request.Pagination.OrderBy == "merchantNoRef" {
		orderby = "trx_v2.merchantnoref"
	}
	if request.Pagination.OrderBy == "createdAt" {
		orderby = "trx_v2.created_at"
	}
	if request.Pagination.OrderBy == "metodePembayaran" {
		orderby = "trx_v2.payment_category_name"
	}
	if request.Pagination.OrderBy == "acquiring" {
		orderby = "trx_v2.payment_method_name"
	}
	if request.Pagination.OrderBy == "settlementDestination" {
		orderby = "trx_v2.settlement_dest"
	}
	if request.Pagination.OrderBy == "statusTrx" {
		orderby = "trx_v2.payment_status"
	}
	if request.Pagination.OrderBy == "statusSettlement" {
		orderby = "trx_v2.status_settlement"
	}
	if request.Pagination.OrderBy == "resNoRef" {
		orderby = "trx_v2.response_noref"
	}
	if request.Pagination.OrderBy == "cid" {
		orderby = "trx_v2.corporate_cid"
	}
	if request.Pagination.OrderBy == "corporate" {
		orderby = "trx_v2.corporate_name"
	}
	if request.Pagination.OrderBy == "deviceId" {
		orderby = "trx_v2.device_id"
	}
	if request.Pagination.OrderBy == "mid" {
		orderby = "trx_v2.bank_mid"
	}
	if request.Pagination.OrderBy == "tid" {
		orderby = "trx_v2.bank_tid"
	}
	if request.Pagination.OrderBy == "cardPan" {
		orderby = "trx_v2.card_pan"
	}
	if request.Pagination.OrderBy == "cardType" {
		orderby = "trx_v2.card_type"
	}
	if request.Pagination.OrderBy == "hargaJual" {
		orderby = "trx_v2.payment_amount"
	}
	if request.Pagination.OrderBy == "potongan" {
		orderby = "trx_v2.payment_disc"
	}
	if request.Pagination.OrderBy == "kodePromo" {
		orderby = "trx_v2.payment_promo_code"
	}
	if request.Pagination.OrderBy == "mdr" {
		orderby = "trx_v2.payment_mdr"
	}
	if request.Pagination.OrderBy == "serviceFee" {
		orderby = "trx_v2.service_fee"
	}
	if request.Pagination.OrderBy == "paymentFee" {
		orderby = "trx_v2.payment_fee"
	}
	if request.Pagination.OrderBy == "vendorFee" {
		orderby = "trx_v2.vendor_fee"
	}

	if request.Pagination.OrderBy == constants.EMPTY_VALUE {
		orderby = "trx_v2.id"
	}

	request.Pagination.OrderBy = orderby

	// fmt.Println(request)
	paginations := models.PageOffsetLimit{
		Offset: request.Pagination.Limit * (request.Pagination.Page - 1),
		Limit:  request.Pagination.Limit,
	}

	resTransaction, err := svc.Service.ListTransactionRepo.GetListTransaction(request, paginations)

	if err != nil {
		log.Println("Error GetListTransaction - GetListTransaction: ", err.Error())
		result = ResponseListJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get List Transaction", constants.EMPTY_VALUE_INT, err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Success GetListTransaction")
	result = ResponseListJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get Data Transaction", len(resTransaction), resTransaction)
	return ctx.JSON(http.StatusOK, result)
}

// Get All Transaction
func (svc listtransactionService) GetListTransactionSettled(ctx echo.Context) error {
	var result models.ResponseList
	var request models.RequestListTrx

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data GetListTransaction : ", err.Error())
		result = ResponseListJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", constants.EMPTY_VALUE_INT, err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}
	var orderby string

	if request.Pagination.OrderBy == "id" {
		orderby = "trx_v2.id"
	}

	if request.Pagination.OrderBy == "nomorHeader" {
		orderby = "trx_v2.noheader"
	}
	if request.Pagination.OrderBy == "merchantNoRef" {
		orderby = "trx_v2.merchantnoref"
	}
	if request.Pagination.OrderBy == "createdAt" {
		orderby = "trx_v2.created_at"
	}
	if request.Pagination.OrderBy == "metodePembayaran" {
		orderby = "trx_v2.payment_category_name"
	}
	if request.Pagination.OrderBy == "acquiring" {
		orderby = "trx_v2.payment_method_name"
	}
	if request.Pagination.OrderBy == "settlementDestination" {
		orderby = "trx_v2.settlement_dest"
	}
	if request.Pagination.OrderBy == "statusTrx" {
		orderby = "trx_v2.payment_status"
	}
	if request.Pagination.OrderBy == "statusSettlement" {
		orderby = "trx_v2.status_settlement"
	}
	if request.Pagination.OrderBy == "resNoRef" {
		orderby = "trx_v2.response_noref"
	}
	if request.Pagination.OrderBy == "cid" {
		orderby = "trx_v2.corporate_cid"
	}
	if request.Pagination.OrderBy == "corporate" {
		orderby = "trx_v2.corporate_name"
	}
	if request.Pagination.OrderBy == "deviceId" {
		orderby = "trx_v2.device_id"
	}
	if request.Pagination.OrderBy == "mid" {
		orderby = "trx_v2.bank_mid"
	}
	if request.Pagination.OrderBy == "tid" {
		orderby = "trx_v2.bank_tid"
	}
	if request.Pagination.OrderBy == "cardPan" {
		orderby = "trx_v2.card_pan"
	}
	if request.Pagination.OrderBy == "cardType" {
		orderby = "trx_v2.card_type"
	}
	if request.Pagination.OrderBy == "hargaJual" {
		orderby = "trx_v2.payment_amount"
	}
	if request.Pagination.OrderBy == "potongan" {
		orderby = "trx_v2.payment_disc"
	}
	if request.Pagination.OrderBy == "kodePromo" {
		orderby = "trx_v2.payment_promo_code"
	}
	if request.Pagination.OrderBy == "mdr" {
		orderby = "trx_v2.payment_mdr"
	}
	if request.Pagination.OrderBy == "serviceFee" {
		orderby = "trx_v2.service_fee"
	}
	if request.Pagination.OrderBy == "paymentFee" {
		orderby = "trx_v2.payment_fee"
	}
	if request.Pagination.OrderBy == "vendorFee" {
		orderby = "trx_v2.vendor_fee"
	}

	if request.Pagination.OrderBy == constants.EMPTY_VALUE {
		orderby = "trx_v2.id"
	}

	request.Pagination.OrderBy = orderby

	// fmt.Println(request)
	paginations := models.PageOffsetLimit{
		Offset: request.Pagination.Limit * (request.Pagination.Page - 1),
		Limit:  request.Pagination.Limit,
	}

	resTransaction, err := svc.Service.ListTransactionRepo.GetListTransactionSettled(request, paginations)

	if err != nil {
		log.Println("Error GetListTransaction - GetListTransaction: ", err.Error())
		result = ResponseListJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get List Transaction", constants.EMPTY_VALUE_INT, err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Success GetListTransaction")
	result = ResponseListJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get Data Transaction", len(resTransaction), resTransaction)
	return ctx.JSON(http.StatusOK, result)
}

func (svc listtransactionService) GetSingleTransaction(ctx echo.Context) error {
	var result models.Response
	var request models.RequestTrx

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data GetSingleTransaction : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	_, exists, err := svc.Service.ListTransactionRepo.IsListTrxExistsByIndex(request)

	if err != nil {
		log.Println("Error GetSingleTransaction - IsListTransactionExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Transaction", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error GetSingleTransaction - IsListTransactionExistsByIndex : Id Transaction Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Transaction Not Found", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	resTransaction, err := svc.Service.ListTransactionRepo.GetSingleTransaction(request.IdTrx)

	if err != nil {
		log.Println("Error GetSingleTransaction - GetSingleTransaction : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Data Transaction", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse GetSingleTransaction - GetSingleTransaction : ", resTransaction)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get Data Transaction", resTransaction)
	return ctx.JSON(http.StatusOK, result)
}

func (svc listtransactionService) GetSingleDetailProduk(ctx echo.Context) error {
	var result models.Response
	var request models.RequestTrx

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data GetSingleTransaction : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	_, exists, err := svc.Service.ListTransactionRepo.IsListTrxExistsByIndex(request)

	if err != nil {
		log.Println("Error GetSingleTransaction - IsListTransactionExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Transaction", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error GetSingleTransaction - IsListTransactionExistsByIndex : Id Transaction Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Transaction Not Found", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	resTransaction, err := svc.Service.ListTransactionRepo.GetSingleProdukTransaction(request.IdTrx)

	if err != nil {
		log.Println("Error GetSingleTransaction - GetSingleProdukTransaction : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Data Transaction", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse GetSingleTransaction - GetSingleProdukTransaction : ", resTransaction)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get Data Transaction", resTransaction)
	return ctx.JSON(http.StatusOK, result)
}
