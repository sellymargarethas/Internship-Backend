package produkService

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/services"
	. "Internship-Backend/utils"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type kategoriprodukService struct {
	Service services.UsecaseService
}

func NewKategoriProdukService(service services.UsecaseService) kategoriprodukService {
	return kategoriprodukService{
		Service: service,
	}
}

// Get All Kategori Produk
func (svc kategoriprodukService) GetListKategoriProduk(ctx echo.Context) error {
	var result models.Response
	resKategoriProduk, err := svc.Service.KategoriProdukRepo.GetListKategoriProduk()

	if err != nil {
		log.Println("Error GetListKategoriProduk - GetListKategoriProduk : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Data Kategori Produk", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse GetListKategoriProduk - GetListKategoriProduk : ", resKategoriProduk)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get Data Kategori Produk", resKategoriProduk)
	return ctx.JSON(http.StatusOK, result)
}

// Insert Kategori Produk
func (svc kategoriprodukService) InsertKategoriProduk(ctx echo.Context) error {
	var result models.Response
	var request models.RequestAddKategoriProduk
	var uraian models.KategoriProduk

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data InsertKategoriProduk : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	uraian.Uraian = request.Uraian
	_, exists, err := svc.Service.KategoriProdukRepo.IsKategoriProdukExistsByIndex(uraian)

	if err != nil {
		log.Println("Error InsertKategoriProduk - IsKategoriProdukExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Uraian", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if exists {
		log.Println("Error InsertKategoriProduk - IsKategoriProdukExistsByIndex : Uraian Already Exists")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_FOUND_CODE, "Uraian Already Exists", err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	kategoriproduk := models.RequestAddKategoriProduk{
		Uraian:     request.Uraian,
		ServiceFee: request.ServiceFee,
		CreatedAt:  TimeStampNow(),
	}

	idkategoriproduk, err := svc.Service.KategoriProdukRepo.InsertKategoriProduk(kategoriproduk)

	if err != nil {
		log.Println("Error InsertKategoriProduk- InsertKategoriProduk : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert Kategori Produk", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse InsertKategoriProduk - InsertKategoriProduk : ", idkategoriproduk)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert Kategori Produk", idkategoriproduk)
	return ctx.JSON(http.StatusOK, result)
}

// Update Kategori Produk
func (svc kategoriprodukService) UpdateKategoriProduk(ctx echo.Context) error {
	var result models.Response
	var request models.RequestUpdateKategoriProduk
	var kategoriproduk models.KategoriProduk

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data UpdateKategoriProduk : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	kategoriproduk.Id = request.Id
	_, exists, err := svc.Service.KategoriProdukRepo.IsKategoriProdukExistsByIndex(kategoriproduk)

	if err != nil {
		log.Println("Error UpdateKategoriProduk - IsKategoriProdukExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Kategori Produk", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error UpdateKategoriProduk - IsKategoriProdukExistsByIndex : Id Kategori Produk Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Kategori Produk Not Found", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	idkategoriproduk, err := svc.Service.KategoriProdukRepo.UpdateKategoriProduk(request)

	if err != nil {
		log.Println("Error UpdateKategoriProduk - UpdateKategoriProduk : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Update Kategori Produk", err.Error())

		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse UpdateKategoriProduk - UpdateKategoriProduk : ", idkategoriproduk)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Update Kategori Produk", idkategoriproduk)
	return ctx.JSON(http.StatusOK, result)
}

// Delete Kategori Produk
func (svc kategoriprodukService) DeleteKategoriProduk(ctx echo.Context) error {
	var result models.Response
	var request models.RequestDeleteKategoriProduk
	var kategoriproduk models.KategoriProduk

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data DeleteKategoriProduk : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	kategoriproduk.Id = request.Id
	_, exists, err := svc.Service.KategoriProdukRepo.IsKategoriProdukExistsByIndex(kategoriproduk)

	if err != nil {
		log.Println("Error DeleteKategoriProduk - IsKategoriProdukExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Kategori Produk", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error DeleteKategoriProduk - IsKategoriProdukExistsByIndex : Id Kategori Produk Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Kategori Produk Not Found", err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	idkategoriproduk, err := svc.Service.KategoriProdukRepo.DeleteKategoriProduk(request)

	if err != nil {
		log.Println("Error DeleteKategoriProduk - DeleteKategoriProduk : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Delete Kategori Produk", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse DeleteKategoriProduk - DeleteKategoriProduk : ", idkategoriproduk)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Delete Kategori Produk", idkategoriproduk)
	return ctx.JSON(http.StatusOK, result)
}
