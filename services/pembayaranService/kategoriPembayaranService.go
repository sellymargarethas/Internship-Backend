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

type kategoripembayaranService struct {
	Service services.UsecaseService
}

func NewKategoriPembayaranService(service services.UsecaseService) kategoripembayaranService {
	return kategoripembayaranService{
		Service: service,
	}
}

// Get All Kategori Pembayaran
func (svc kategoripembayaranService) GetListKategoriPembayaran(ctx echo.Context) error {
	var result models.Response
	resKategoriPembayaran, err := svc.Service.KategoriPembayaranRepo.GetListKategoriPembayaran()

	if err != nil {
		log.Println("Error GetListKategoriPembayaran - GetListKategoriPembayaran: ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Data Kategori Pembayaran", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse GetListKategoriPembayaran - GetListKategoriPembayaran : ", resKategoriPembayaran)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get Data Kategori Pembayaran", resKategoriPembayaran)
	return ctx.JSON(http.StatusOK, result)
}

// Insert Kategori Pembayaran
func (svc kategoripembayaranService) InsertKategoriPembayaran(ctx echo.Context) error {
	var result models.Response
	var request models.RequestAddKategoriPembayaran
	var uraian models.KategoriPembayaran

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data InsertKategoriPembayaran : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	uraian.Uraian = request.Uraian
	_, exists, err := svc.Service.KategoriPembayaranRepo.IsKategoriPembayaranExistsByIndex(uraian)

	if err != nil {
		log.Println("Error InsertKategoriPembayaran - IsKategoriPembayaranExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Uraian", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if exists {
		log.Println("Error InsertKategoriPembayaran - IsKategoriPembayaranExistsByIndex : Uraian Already Exists")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_FOUND_CODE, "Uraian Already Exists", err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	kategoripembayaran := models.RequestAddKategoriPembayaran{
		Uraian:    request.Uraian,
		CreatedAt: TimeStampNow(),
	}

	idkategoripembayaran, err := svc.Service.KategoriPembayaranRepo.InsertKategoriPembayaran(kategoripembayaran)

	if err != nil {
		log.Println("Error InsertKategoriPembayaran- InsertKategoriPembayaran : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert Kategori Pembayaran", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse InsertKategori - InsertKategori : ", idkategoripembayaran)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert Kategori Pembayaran", idkategoripembayaran)
	return ctx.JSON(http.StatusOK, result)
}

// Update Kategori Pembayaran
func (svc kategoripembayaranService) UpdateKategoriPembayaran(ctx echo.Context) error {
	var result models.Response
	var request models.RequestUpdateKategoriPembayaran
	var kategoripembayaran models.KategoriPembayaran

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data UpdateKategoriPembayaran : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	kategoripembayaran.Id = request.Id
	_, exists, err := svc.Service.KategoriPembayaranRepo.IsKategoriPembayaranExistsByIndex(kategoripembayaran)

	if err != nil {
		log.Println("Error UpdateKategoriPembayaran - IsKategoriPembayaranExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Kategori Pembayaran", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error UpdateKategoriPembayaran - IsKategoriPembayaranExistsByIndex : Id KategoriPembayaran Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Kategori Pembayaran Not Found", err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	idkategoripembayaran, err := svc.Service.KategoriPembayaranRepo.UpdateKategoriPembayaran(request)

	if err != nil {
		log.Println("Error UpdateKategoriPembayaran - UpdateKategoriPembayaran : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Update Kategori Pembayaran", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse UpdateKategoriPembayaran - UpdateKategoriPembayaran : ", idkategoripembayaran)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Update Kategori Pembayaran", idkategoripembayaran)
	return ctx.JSON(http.StatusOK, result)
}

// Delete Kategori Pembayaran
func (svc kategoripembayaranService) DeleteKategoriPembayaran(ctx echo.Context) error {
	var result models.Response
	var request models.RequestDeleteKategoriPembayaran
	var kategoripembayaran models.KategoriPembayaran

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data DeleteKategoriPembayaran : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	kategoripembayaran.Id = request.Id
	_, exists, err := svc.Service.KategoriPembayaranRepo.IsKategoriPembayaranExistsByIndex(kategoripembayaran)

	if err != nil {
		log.Println("Error DeleteKategoriPembayaran - IsKategoriPembayaranExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Kategori Pembayaran", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error DeleteKategoriPembayaran - IsKategoriPembayaranExistsByIndex : Id Kategori Pembayaran Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Kategori Pembayaran Not Found", err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	idkategoripembayaran, err := svc.Service.KategoriPembayaranRepo.DeleteKategoriPembayaran(request)

	if err != nil {
		log.Println("Error DeleteKategoriPembayaran- DeleteKategoriPembayaran : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Delete Kategori Pembayaran", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse DeleteKategoriPembayaran- DeleteKategoriPembayaran : ", idkategoripembayaran)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Delete Kategori Pembayaran", idkategoripembayaran)
	return ctx.JSON(http.StatusOK, result)
}
