package coresettlementService

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/services"
	. "Internship-Backend/utils"

	"log"
	"net/http"

	"github.com/labstack/echo"
)

type coresettlementService struct {
	Service services.UsecaseService
}

func NewCoreSettlementService(service services.UsecaseService) coresettlementService {
	return coresettlementService{
		Service: service,
	}
}

func (svc coresettlementService) InsertCoreSettlementKey(ctx echo.Context) (err error) {
	var result models.Response
	var request models.RequestAddCoreSettlement

	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data: InsertCoreSettlementKey", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	resCoreSettlement, err := svc.Service.CoreSettlementRepo.InsertCoreSettlementKey(request)
	if err != nil {
		log.Println("Error insert core settlement key- insert core settlement key : ", err.Error(), nil)
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert Corporate", err.Error())

		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse insert core settlement key- insert core settlement key", resCoreSettlement, err)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert Core Settlement Keys", resCoreSettlement)
	return ctx.JSON(http.StatusOK, result)
}

func (svc coresettlementService) UpdateCoreSettlementKey(ctx echo.Context) (err error) {
	var result models.Response
	var request models.RequestUpdateCoreSettlement

	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data: UpdateCoreSettlementKey", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	isExist := models.CoreSettlementKeys{
		Id: request.Id,
	}
	check, err := svc.Service.CoreSettlementRepo.IsCoresettlementkeysExistByIndex(isExist)

	if err != nil {
		log.Println("Error IsCoresettlementkeysExistById: UpdateCoreSettlementKey", err.Error())
	}
	if !check {
		log.Println("Error Update core settlement keys- Update core settlement keys: ID not found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.SYSTEM_ERROR_CODE, "ID Not Found", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	resUpdateCoreSettlement, err := svc.Service.CoreSettlementRepo.UpdateCoreSettlementKey(request)
	if err != nil {
		log.Println("Error Update core settlement keys- Update core settlement keys : ", err.Error(), resUpdateCoreSettlement)
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "FailedUpdate Core Settlement Keys", nil)

		return ctx.JSON(http.StatusBadRequest, result)
	}
	if resUpdateCoreSettlement == constants.EMPTY_VALUE_INT {
		log.Println("Error Update core settlement keys- Update core settlement keys : ", resUpdateCoreSettlement)
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "FailedUpdate Core Settlement Keys", nil)

		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse Update core settlement keys- Update core settlement keys", err)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Update core settlement keys", resUpdateCoreSettlement)

	return ctx.JSON(http.StatusOK, result)

}

func (svc coresettlementService) DeleteCoreSettlementKey(ctx echo.Context) (err error) {
	var result models.Response
	var request models.RequestDeleteCoreSettlement
	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data: DeleteCoreSettlementKey", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	isExist := models.CoreSettlementKeys{
		Id: request.Id,
	}
	check, err := svc.Service.CoreSettlementRepo.IsCoresettlementkeysExistByIndex(isExist)
	if err != nil {
		log.Println("Error IsCoresettlementkeysExistById: DeleteCoreSettlementKey", err.Error())
	}

	if !check {
		log.Println("Error Delete core settlement keys- Delete core settlement keys: ID not found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.SYSTEM_ERROR_CODE, "ID Not Found", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	resDeleteCoreSettlement, err := svc.Service.CoreSettlementRepo.DeleteCoreSettlementKey(request)
	if err != nil {
		log.Println("Error Delete core settlement keys- Delete core settlement keys : ", err.Error(), resDeleteCoreSettlement)
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert Core Settlement Keys", nil)

		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse Delete core settlement keys- Delete core settlement keys", err)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Delete core settlement keys", resDeleteCoreSettlement)

	return ctx.JSON(http.StatusOK, result)

}

func (svc coresettlementService) GetSingleCoreSettlementKeys(ctx echo.Context) error {
	var result models.Response
	var request models.ResponseCoreSettlement
	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	data, err := svc.Service.CoreSettlementRepo.GetSingleCoreSettlementKeys(request.Id)

	if err != nil {
		log.Println("Error Get Single Core Settlement Keys : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("Reponse Get Single Core Settlement Keys ", "Success Get Data")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get Data", data)
	return ctx.JSON(http.StatusOK, result)
}

func (svc coresettlementService) GetAllCoreSettlementKeys(ctx echo.Context) error {
	var result models.ResponseList
	var request models.RequestList
	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data: GetAllCoreSettlementKeys ", err.Error())
		result = ResponseListJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), constants.EMPTY_VALUE_INT, nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var orderby string

	if request.OrderBy == "id" {
		orderby = "coresettlementkeys.id"
	}

	if request.OrderBy == "value" {
		orderby = "coresettlementkeys.value"
	}
	if request.OrderBy == "name" {
		orderby = "coresettlementkeys.name"
	}
	if request.OrderBy == "workerUsername" {
		orderby = "coresettlementkeys.workerusername"
	}

	if request.OrderBy == constants.EMPTY_VALUE {
		orderby = "coresettlementkeys.id"
	}

	request.OrderBy = orderby

	rescoresettlement, err := svc.Service.CoreSettlementRepo.GetListCoreSettlementKeys(request)
	if err != nil {
		log.Println("Error Get Single Core Settlement Keys : GetListCoresettlementKeys", err)
		result = ResponseListJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "ERROR Get Coresettlementkeys", constants.EMPTY_VALUE_INT, nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	limitbawah := (request.Page - 1) * request.Limit
	limitatas := request.Limit * request.Page

	if len(rescoresettlement) < limitatas {
		limitatas = len(rescoresettlement)
	}
	if limitbawah > len(rescoresettlement) {
		log.Println("Error Get Single Core Settlement Keys : GetListCoresettlementKeys: limit melebihi jumlah data", err)
		result = ResponseListJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "ERROR Get Coresettlementkeys: Limit melebihi jumlah data", len(rescoresettlement), err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	coresettlementslice := rescoresettlement[limitbawah:limitatas]

	if len(coresettlementslice) == constants.EMPTY_VALUE_INT {
		log.Println("Error Get Single Core Settlement Keys : GetListCoresettlementKeys: limit melebihi jumlah data", err)
		result = ResponseListJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Data Tidak Ditemukan", len(rescoresettlement), err)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("Reponse Get Single Core Settlement Keys ", "Success Get Data")
	result = ResponseListJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get Coresettlementkeys", len(rescoresettlement), coresettlementslice)
	return ctx.JSON(http.StatusOK, result)
}
