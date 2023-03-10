package paketService

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/services"
	. "Internship-Backend/utils"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type paketService struct {
	Service services.UsecaseService
}

func NewPaketService(service services.UsecaseService) paketService {
	return paketService{
		Service: service,
	}
}

// Get All Paket
func (svc paketService) GetListPaket(ctx echo.Context) error {
	var result models.Response
	resPaket, err := svc.Service.PaketRepo.GetListPaket()

	if err != nil {
		log.Println("Error GetListPaket - GetListPaket : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Data Paket", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse GetListPaket - GetListPaket : ", resPaket)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get Data Paket", resPaket)
	return ctx.JSON(http.StatusOK, result)
}

// Insert Paket
func (svc paketService) InsertPaket(ctx echo.Context) error {
	var result models.Response
	var request models.RequestAddPaket
	var uraian models.Paket

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data InsertPaket : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	uraian.Uraian = request.Uraian
	_, exists, err := svc.Service.PaketRepo.IsPaketExistsByIndex(uraian)

	if err != nil {
		log.Println("Error InsertPaket - IsPaketExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Uraian", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if exists {
		log.Println("Error InsertPaket - IsPaketExistsByIndex : Uraian Already Exists")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_FOUND_CODE, "Uraian Already Exists", err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	paket := models.RequestAddPaket{
		Uraian:      request.Uraian,
		IdCorporate: request.IdCorporate,
		CreatedAt:   TimeStampNow(),
	}

	newpaket, err := svc.Service.PaketRepo.InsertPaket(paket)

	if err != nil {
		log.Println("Error InsertPaket - InsertPaket : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert Paket", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse InsertPaket - InsertPaket : ", newpaket)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert Paket", newpaket)
	return ctx.JSON(http.StatusOK, result)
}

// Update Paket
func (svc paketService) UpdatePaket(ctx echo.Context) error {
	var result models.Response
	var request models.RequestUpdatePaket
	var paket models.Paket

	err := BindValidateStruct(ctx, &request)
	if err != nil {
		log.Println("Error Validate Data UpdatePaket : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	paket.Id = request.Id
	_, exists, err := svc.Service.PaketRepo.IsPaketExistsByIndex(paket)

	if err != nil {
		log.Println("Error UpdatePaket - IsPaketExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Paket", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error UpdatePaket - IsPaketExistsByIndex : Id Paket Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Paket Not Found", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	idpaket, err := svc.Service.PaketRepo.UpdatePaket(request)

	if err != nil {
		log.Println("Error UpdatePaket - UpdatePaket : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Update Paket", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse UpdatePaket - UpdatePaket : ", idpaket)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Update Paket", idpaket)
	return ctx.JSON(http.StatusOK, result)
}

// Delete Paket
func (svc paketService) DeletePaket(ctx echo.Context) error {
	var result models.Response
	var request models.RequestDeletePaket
	var paket models.Paket

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data DeletePaket : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	paket.Id = request.Id
	_, exists, err := svc.Service.PaketRepo.IsPaketExistsByIndex(paket)

	if err != nil {
		log.Println("Error DeletePaket - IsPaketExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Paket", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error DeletePaket - IsPaketExistsByIndex : Id Paket Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Paket Not Found", err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	idPaket, err := svc.Service.PaketRepo.DeletePaket(request)

	if err != nil {
		log.Println("Error DeletePaket - DeletePaket : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Delete Paket", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse DeletePaket - DeletePaket : ", idPaket)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Delete Paket", idPaket)
	return ctx.JSON(http.StatusOK, result)
}
