package wilayahService

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/services"
	. "Internship-Backend/utils"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type provinsiService struct {
	Service services.UsecaseService
}

func NewProvinsiService(service services.UsecaseService) provinsiService {
	return provinsiService{
		Service: service,
	}
}

// Get All Provinsi
func (svc provinsiService) GetListProvinsi(ctx echo.Context) error {
	var result models.Response
	resProvinsi, err := svc.Service.ProvinsiRepo.GetListProvinsi()

	if err != nil {
		log.Println("Error GetListProvinsi - GetListProvinsi : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Data Provinsi", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse GetListProvinsi - GetListProvinsi : ", resProvinsi)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get Data Provinsi", resProvinsi)
	return ctx.JSON(http.StatusOK, result)
}

// Insert Provinsi
func (svc provinsiService) InsertProvinsi(ctx echo.Context) error {
	var result models.Response
	var request models.RequestAddProvinsi
	var uraian models.Provinsi

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data InsertProvinsi : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	uraian.Uraian = request.Uraian
	_, exists, err := svc.Service.ProvinsiRepo.IsProvinsiExistsByIndex(uraian)

	if err != nil {
		log.Println("Error InsertProvinsi - IsProvinsiExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Uraian", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if exists {
		log.Println("Error InsertProvinsi - IsProvinsiExistsByIndex : Uraian Already Exists")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_FOUND_CODE, "Uraian Already Exists", err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	provinsi := models.RequestAddProvinsi{
		Uraian:    request.Uraian,
		CreatedAt: TimeStampNow(),
	}

	idprovinsi, err := svc.Service.ProvinsiRepo.InsertProvinsi(provinsi)

	if err != nil {
		log.Println("Error InsertProvinsi - InsertProvinsi : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert Provinsi", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse InsertProvinsi - InsertProvinsi : ", idprovinsi)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert Provinsi", idprovinsi)
	return ctx.JSON(http.StatusOK, result)
}

// Update Provinsi
func (svc provinsiService) UpdateProvinsi(ctx echo.Context) error {
	var result models.Response
	var request models.RequestUpdateProvinsi
	var provinsi models.Provinsi

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data UpdateProvinsi : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	provinsi.Id = request.Id
	_, exists, err := svc.Service.ProvinsiRepo.IsProvinsiExistsByIndex(provinsi)

	if err != nil {
		log.Println("Error UpdateProvinsi - IsProvinsiExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Provinsi", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error UpdateProvinsi - IsProvinsiExistsByIndex : Id Provinsi Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Provinsi Not Found", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	idprovinsi, err := svc.Service.ProvinsiRepo.UpdateProvinsi(request)

	if err != nil {
		log.Println("Error UpdateProvinsi - UpdateProvinsi : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Update Provinsi", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse UpdateProvinsi - UpdateProvinsi : ", idprovinsi)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Update Provinsi", idprovinsi)
	return ctx.JSON(http.StatusOK, result)
}

// Delete Provinsi
func (svc provinsiService) DeleteProvinsi(ctx echo.Context) error {
	var result models.Response
	var request models.RequestDeleteProvinsi
	var provinsi models.Provinsi

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data DeleteProvinsi : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	provinsi.Id = request.Id
	_, exists, err := svc.Service.ProvinsiRepo.IsProvinsiExistsByIndex(provinsi)

	if err != nil {
		log.Println("Error DeleteProvinsi - IsProvinsiExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Provinsi", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error DeleteProvinsi - IsProvinsiExistsByIndex : Id Provinsi Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Provinsi Not Found", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	idprovinsi, err := svc.Service.ProvinsiRepo.DeleteProvinsi(request)

	if err != nil {
		log.Println("Error DeleteProvinsi - DeleteProvinsi : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Delete Provinsi", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse DeleteProvinsi - DeleteProvinsi : ", idprovinsi)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Delete Provinsi", idprovinsi)
	return ctx.JSON(http.StatusOK, result)
}
