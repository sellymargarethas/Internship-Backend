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

type kotaService struct {
	Service services.UsecaseService
}

func NewKotaService(service services.UsecaseService) kotaService {
	return kotaService{
		Service: service,
	}
}

// Get All Kota
func (svc kotaService) GetListKota(ctx echo.Context) error {
	var result models.Response
	resKota, err := svc.Service.KotaRepo.GetListKota()

	if err != nil {
		log.Println("Error GetListKota - GetListKota : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Data Kota", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse GetListKota - GetListKota : ", resKota)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get Data Kota", resKota)
	return ctx.JSON(http.StatusOK, result)
}

// Insert Kota
func (svc kotaService) InsertKota(ctx echo.Context) error {
	var result models.Response
	var request models.RequestAddKota
	var uraian models.Kota

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data InsertKota : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	uraian.Uraian = request.Uraian
	_, exists, err := svc.Service.KotaRepo.IsKotaExistsByIndex(uraian)

	if err != nil {
		log.Println("Error InsertKota - IsKotaExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Uraian", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if exists {
		log.Println("Error InsertKota - IsKotaExistsByIndex : Uraian Already Exists")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_FOUND_CODE, "Uraian Already Exists", err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	kota := models.RequestAddKota{
		Uraian:     request.Uraian,
		IdProvinsi: request.IdProvinsi,
		CreatedAt:  TimeStampNow(),
	}

	idkota, err := svc.Service.KotaRepo.InsertKota(kota)

	if err != nil {
		log.Println("Error InsertKota - InsertKota : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert Kota", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse InsertKota - InsertKota : ", idkota)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert Kota", idkota)
	return ctx.JSON(http.StatusOK, result)
}

// Update Kota
func (svc kotaService) UpdateKota(ctx echo.Context) error {
	var result models.Response
	var request models.RequestUpdateKota
	var kota models.Kota

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data UpdateKota : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	kota.Id = request.Id
	_, exists, err := svc.Service.KotaRepo.IsKotaExistsByIndex(kota)

	if err != nil {
		log.Println("Error UpdateKota - IsKotaExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Kota", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error UpdateKota - IsKotaExistsByIndex : Id Kota Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Kota Not Found", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	idkota, err := svc.Service.KotaRepo.UpdateKota(request)

	if err != nil {
		log.Println("Error UpdateKota- UpdateKota : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Update Kota", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse UpdateKota - UpdateKota : ", idkota)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Update Kota", idkota)
	return ctx.JSON(http.StatusOK, result)
}

// Delete Kota
func (svc kotaService) DeleteKota(ctx echo.Context) error {
	var result models.Response
	var request models.RequestDeleteKota
	var kota models.Kota

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data DeleteKota : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	kota.Id = request.Id
	_, exists, err := svc.Service.KotaRepo.IsKotaExistsByIndex(kota)

	if err != nil {
		log.Println("Error DeleteKota - IsKotaExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Kota", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error DeleteKota - IsKotaExistsByIndex : Id Kota Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Kota Not Found", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	idkota, err := svc.Service.KotaRepo.DeleteKota(request)

	if err != nil {
		log.Println("Error DeleteKota - DeleteKota : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Delete Kota", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse DeleteKota - DeleteKota : ", idkota)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Delete Kota", idkota)
	return ctx.JSON(http.StatusOK, result)
}
