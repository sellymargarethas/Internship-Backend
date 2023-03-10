package corporateservice

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/services"
	. "Internship-Backend/utils"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type corporatePaymentService struct {
	Service services.UsecaseService
}

func NewCorporatePaymentService(service services.UsecaseService) corporatePaymentService {
	return corporatePaymentService{
		Service: service,
	}
}

func (svc corporatePaymentService) GetAllCorporatePayment(ctx echo.Context) (err error) {
	var result models.ResponseList
	var request models.RequestList
	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data: GetAllCorporatePayment", err.Error())
		result = ResponseListJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), constants.EMPTY_VALUE_INT, nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	if request.HirarkiId == constants.EMPTY_VALUE {
		request.HirarkiId = "1000000001"
	}

	var orderby string

	if request.OrderBy == "idCorporate" {
		orderby = "corporate.id"
	}
	if request.OrderBy == "CID" {
		orderby = "corporate.cid"
	}
	if request.OrderBy == "uraian" {
		orderby = "corporate.uraian"
	}
	if request.OrderBy == "kota" {
		orderby = "corporate.nama_kota"
	}
	if request.OrderBy == "jmlMetodePemabayaran" {
		orderby = "(SELECT COUNT(idmetodepembayaran) FROM corporatepayment WHERE corporatepayment.idcorporate=(corporate.id::int))"
	}

	if request.OrderBy == constants.EMPTY_VALUE {
		orderby = "corporate.id"
	}

	request.OrderBy = orderby
	// fmt.Println(request)

	resCorporatePayment, err := svc.Service.CorporatePaymentRepo.GetListCorporatePayment(request)
	if err != nil {
		log.Println("Error GET CorporatePayment :  GetAllCorporatePayment ", err.Error())
		result = ResponseListJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Corporate Payment", constants.EMPTY_VALUE_INT, nil)

		return ctx.JSON(http.StatusBadRequest, result)
	}
	limitbawah := (request.Page - 1) * request.Limit
	limitatas := request.Limit * request.Page

	if len(resCorporatePayment) < limitatas {
		limitatas = len(resCorporatePayment)
	}
	if limitbawah > len(resCorporatePayment) {
		log.Println("Error Get Single Core Settlement Keys : GetListCoresettlementKeys: limit melebihi jumlah data", err)
		result = ResponseListJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "ERROR Get Coresettlementkeys: Limit melebihi jumlah data", len(resCorporatePayment), err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	corporatepaymentslice := resCorporatePayment[limitbawah:limitatas]
	if len(corporatepaymentslice) == constants.EMPTY_VALUE_INT {
		log.Println("Error Get Single Core Settlement Keys : GetListCoresettlementKeys: limit melebihi jumlah data", err)
		result = ResponseListJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Data Tidak Ditemukan", len(resCorporatePayment), err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse GetAllCorporatePayment ", "GetAllCorporatePayment", err)
	result = ResponseListJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, constants.EMPTY_VALUE, len(resCorporatePayment), corporatepaymentslice)
	return ctx.JSON(http.StatusOK, result)
}

func (svc corporatePaymentService) GetCorporatePaymentByIdCorporate(ctx echo.Context) (err error) {
	var result models.Response
	var request models.RequestGetCorporatePaymentByID

	var response models.ResponseCorporatePaymentDetails

	if err := BindValidateStruct(ctx, &request); err != nil {
		log.Println("Error Validate Data: GetCorporatePaymentByIdCorporate", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	corporate := models.Corporate{
		Id: request.IdCorporate,
	}

	resCorporate, exist, err := svc.Service.CorporateRepo.IsCorporateExistsByIndex(corporate)
	if err != nil {
		log.Println("Error IsCorporateExistsByIndex  :  GetCorporatePaymentByIdCorporate ", err.Error())
	}

	if !exist {
		log.Println("Reponse ERROR GetListTransaction: IsCorporateExistsByIndex -  GetCorporatePaymentByIdCorporate", err)
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "CORPORATE ID NOT FOUND", resCorporate)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	resCorporatePayment, err := svc.Service.CorporatePaymentRepo.GetCorporatePaymentByIdCorporate(request.IdCorporate)
	if err != nil {
		log.Println("Reponse ERROR GetListTransaction ", "ERROR Get Transaction", err)
		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "CORPORATE ID NOT FOUND", err)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	response.IdCorporate = resCorporate.Id
	res := models.ResponseGetOne{
		IdCorporate:             request.IdCorporate,
		MetodePembayaranDetails: resCorporatePayment,
	}

	log.Println("Reponse GetListTransaction ", "Success GetCorporatePayment", res)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get One Corporate Payment", res)
	return ctx.JSON(http.StatusOK, result)
}

func (svc corporatePaymentService) InsertCorporatePayment(ctx echo.Context) (err error) {
	var result models.Response
	var request models.RequestAddCorporatePayment

	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data: InsertCorporatePayment", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	corporate := models.Corporate{
		Id: request.IdCorporate,
	}

	resCorporate, exist, err := svc.Service.CorporateRepo.IsCorporateExistsByIndex(corporate)
	if err != nil {
		log.Println("Error IsCorporateExistsByIndex  :  InsertCorporatePayment ", err.Error())
	}

	if !exist {
		log.Println("Reponse ERROR InsertCorporatePayment: IsCorporateExistsByIndex -  GetCorporatePaymentByIdCorporate", err)
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "CORPORATE ID NOT FOUND", resCorporate)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	resInsert, err := svc.Service.CorporatePaymentRepo.InsertCorporatePayment(request)
	fmt.Println(resInsert)
	if err != nil {
		log.Println("Error InsertCorporatePayment- InsertCorporatePayment : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert Corporate Payment", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("Reponse InsertCorporatePayment- InsertCorporatePayment")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert Corporate Payment", request.IdCorporate)
	return ctx.JSON(http.StatusOK, result)
}

func (svc corporatePaymentService) UpdateCorporatePayment(ctx echo.Context) (err error) {
	var result models.Response
	var request models.RequestUpdateCorporatePayment
	if err := BindValidateStruct(ctx, &request); err != nil {
		log.Println("Error Validate Data", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	err = DBTransaction(svc.Service.RepoDB, func(tx *sql.Tx) error {
		if len(request.MetodePembayaranDetails) == constants.EMPTY_VALUE_INT {
			resDelete, err := svc.Service.CorporatePaymentRepo.DeleteCorporatePaymentDetails(request.IdCorporate, tx)
			fmt.Println(resDelete)
			if err != nil {
				log.Println("Error UpdateCorporatePayment- DeleteCorporatePaymentDetails : ", err.Error())
				result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Update Corporate Payment", err)
				return ctx.JSON(http.StatusBadRequest, result)
			}
			log.Println("Reponse UpdateCorporatePayment- UpdateCorporatePayment")
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Corporate Payment Berhasil dihapus", request.IdCorporate)
			return ctx.JSON(http.StatusOK, result)
		}

		resDelete, err := svc.Service.CorporatePaymentRepo.DeleteCorporatePaymentDetails(request.IdCorporate, tx)
		fmt.Println(resDelete)
		if err != nil {
			log.Println("Error UpdateCorporatePayment- DeleteCorporatePaymentDetails : ", err.Error())
			result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Update Corporate Payment", err)
			return ctx.JSON(http.StatusBadRequest, result)
		}

		resInsert, err := svc.Service.CorporatePaymentRepo.UpdateCorporatePayment(request, tx)
		fmt.Println(resInsert)
		if err != nil {
			log.Println("Error UpdateCorporatePayment- InsertCorporatePaymentDetails : ", err.Error())
			result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Update Corporate Payment", err)
			return ctx.JSON(http.StatusBadRequest, result)
		}
		log.Println("Reponse UpdateCorporatePayment- UpdateCorporatePayment")
		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Update Corporate Payment", request.IdCorporate)
		return ctx.JSON(http.StatusOK, result)
	})
	return nil
}
