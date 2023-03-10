package pembayaranService

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/services"
	. "Internship-Backend/utils"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type metodepembayaranService struct {
	Service services.UsecaseService
}

func NewMetodePembayaranService(service services.UsecaseService) metodepembayaranService {
	return metodepembayaranService{
		Service: service,
	}
}

// Get All Metode Pembayaran
func (svc metodepembayaranService) GetListMetodePembayaran(ctx echo.Context) error {
	var result models.Response
	resMetodePembayaran, err := svc.Service.MetodePembayaranRepo.GetListMetodePembayaran()

	if err != nil {
		log.Println("Error GetListMetodePembayaran - GetListMetodePembayaran : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Data Metode Pembayaran", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse GetListMetodePembayaran - GetListMetodePembayaran : ", resMetodePembayaran)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get Data Metode Pembayaran", resMetodePembayaran)
	return ctx.JSON(http.StatusOK, result)
}

// Insert Metode Pembayaran
func (svc metodepembayaranService) InsertMetodePembayaran(ctx echo.Context) error {
	var result models.Response
	var request models.RequestAddMetodePembayaran
	var uraian models.MetodePembayaran

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data InsertMetodePembayaran : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	uraian.Uraian = request.Uraian
	_, exists, err := svc.Service.MetodePembayaranRepo.IsMetodePembayaranExistsByIndex(uraian)

	if err != nil {
		log.Println("Error InsertMetodePembayaran - IsMetodePembayaranExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Uraian", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if exists {
		log.Println("Error InsertMetodePembayaran - IsMetodePembayaranExistsByIndex : Uraian Already Exists")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_FOUND_CODE, "Uraian Already Exists", err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	metodepembayaran := models.RequestAddMetodePembayaran{
		Uraian:               request.Uraian,
		IdKategoriPembayaran: request.IdKategoriPembayaran,
		CreatedAt:            TimeStampNow(),
	}

	idmetodepembayaran, err := svc.Service.MetodePembayaranRepo.InsertMetodePembayaran(metodepembayaran)

	if err != nil {
		log.Println("Error InsertMetodePembayaran- InsertMetodePembayaran : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert Metode Pembayaran", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("\nReponse InsertMetodePembayaran - InsertMetodePembayaran : ", idmetodepembayaran)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert Metode Pembayaran", idmetodepembayaran)
	return ctx.JSON(http.StatusOK, result)
}

// Update Metode Pembayaran
func (svc metodepembayaranService) UpdateMetodePembayaran(ctx echo.Context) error {
	var result models.Response
	var request models.RequestUpdateMetodePembayaran
	var metodepembayaran models.MetodePembayaran

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data UpdateMetodePembayaran : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	metodepembayaran.Id = request.Id
	_, exists, err := svc.Service.MetodePembayaranRepo.IsMetodePembayaranExistsByIndex(metodepembayaran)

	if err != nil {
		log.Println("Error UpdateMetodePembayaran - IsMetodePembayaranExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Metode Pembayaran", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error UpdateMetodePembayaran - IsMetodePembayaranExistsByIndex : Id Metode Pembayaran Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Metode Pembayaran Not Found", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	idmetodepembayaran, err := svc.Service.MetodePembayaranRepo.UpdateMetodePembayaran(request)

	if err != nil {
		log.Println("Error UpdateMetodePembayaran - UpdateMetodePembayaran : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Update Metode Pembayaran", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse UpdateMetodePembayaran - UpdateMetodePembayaran : ", idmetodepembayaran)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Update Metode Pembayaran", idmetodepembayaran)
	return ctx.JSON(http.StatusOK, result)
}

// Delete Metode Pembayaran
func (svc metodepembayaranService) DeleteMetodePembayaran(ctx echo.Context) error {
	var result models.Response
	var request models.RequestDeleteMetodePembayaran
	var metodepembayaran models.MetodePembayaran

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data DeleteMetodePembayaran : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	metodepembayaran.Id = request.Id
	_, exists, err := svc.Service.MetodePembayaranRepo.IsMetodePembayaranExistsByIndex(metodepembayaran)

	if err != nil {
		log.Println("Error DeleteMetodePembayaran - IsMetodePembayaranExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Metode Pembayaran", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error DeleteMetodePembayaran - IsMetodePembayaranExistsByIndex : Id Metode Pembayaran Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Metode Pembayaran Not Found", err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	idmetodepembayaran, err := svc.Service.MetodePembayaranRepo.DeleteMetodePembayaran(request)

	if err != nil {
		log.Println("Error DeleteMetodePembayaran - DeleteMetodePembayaran : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Delete Metode Pembayaran", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse DeleteMetodePembayaran - DeleteMetodePembayaran : ", idmetodepembayaran)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Delete Metode Pembayaran", idmetodepembayaran)
	return ctx.JSON(http.StatusOK, result)
}
