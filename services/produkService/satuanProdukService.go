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

type satuanprodukService struct {
	Service services.UsecaseService
}

func NewSatuanProdukService(service services.UsecaseService) satuanprodukService {
	return satuanprodukService{
		Service: service,
	}
}

// Get All Satuan Produk
func (svc satuanprodukService) GetListSatuanProduk(ctx echo.Context) error {
	var result models.Response
	resSatuanProduk, err := svc.Service.SatuanProdukRepo.GetListSatuanProduk()

	if err != nil {
		log.Println("Error GetListSatuanProduk - GetListSatuanProduk : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Data Satuan Produk", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse GetListSatuanProduk - GetListSatuanProduk : ", resSatuanProduk)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get Data Satuan Produk", resSatuanProduk)
	return ctx.JSON(http.StatusOK, result)
}

// Insert Satuan Produk
func (svc satuanprodukService) InsertSatuanProduk(ctx echo.Context) error {
	var result models.Response
	var request models.RequestAddSatuanProduk
	var uraian models.SatuanProduk

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data InsertSatuanProduk : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	uraian.Uraian = request.Uraian
	_, exists, err := svc.Service.SatuanProdukRepo.IsSatuanProdukExistsByIndex(uraian)

	if err != nil {
		log.Println("Error InsertSatuanProduk - IsSatuanProdukExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Uraian", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if exists {
		log.Println("Error InsertSatuanProduk - IsSatuanProdukExistsByIndex : Uraian Already Exists")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_FOUND_CODE, "Uraian Already Exists", err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	satuanproduk := models.RequestAddSatuanProduk{
		Uraian:    request.Uraian,
		CreatedAt: TimeStampNow(),
	}

	idsatuanproduk, err := svc.Service.SatuanProdukRepo.InsertSatuanProduk(satuanproduk)

	if err != nil {
		log.Println("Error InsertSatuanProduk- InsertSatuanProduk : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert Satuan Produk", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse InsertSatuanProduk - InsertSatuanProduk : ", idsatuanproduk)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert Satuan Produk", idsatuanproduk)
	return ctx.JSON(http.StatusOK, result)
}

// Update Satuan Produk
func (svc satuanprodukService) UpdateSatuanProduk(ctx echo.Context) error {
	var result models.Response
	var request models.RequestUpdateSatuanProduk
	var satuanproduk models.SatuanProduk

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data UpdateSatuanProduk : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	satuanproduk.Id = request.Id
	_, exists, err := svc.Service.SatuanProdukRepo.IsSatuanProdukExistsByIndex(satuanproduk)

	if err != nil {
		log.Println("Error UpdateSatuanProduk - IsSatuanProdukExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Satuan Produk", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error UpdateSatuanProduk - IsSatuanProdukExistsByIndex : Id Satuan Produk Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Satuan Produk Not Found", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	idsatuanproduk, err := svc.Service.SatuanProdukRepo.UpdateSatuanProduk(request)

	if err != nil {
		log.Println("Error UpdateSatuanProduk - UpdateSatuanProduk : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Update Satuan Produk", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse UpdateSatuanProduk - UpdateSatuanProduk : ", idsatuanproduk)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Update Satuan Produk", idsatuanproduk)
	return ctx.JSON(http.StatusOK, result)
}

// Delete Satuan Produk
func (svc satuanprodukService) DeleteSatuanProduk(ctx echo.Context) error {
	var result models.Response
	var request models.RequestDeleteSatuanProduk
	var satuanproduk models.SatuanProduk

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data DeleteSatuanProduk : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	satuanproduk.Id = request.Id
	_, exists, err := svc.Service.SatuanProdukRepo.IsSatuanProdukExistsByIndex(satuanproduk)

	if err != nil {
		log.Println("Error DeleteSatuanProduk - IsSatuanProdukExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Satuan Produk", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error DeleteSatuanProduk - IsSatuanProdukExistsByIndex : Id Satuan Produk Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Satuan Produk Not Found", err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	idsatuanproduk, err := svc.Service.SatuanProdukRepo.DeleteSatuanProduk(request)

	if err != nil {
		log.Println("Error DeleteSatuanProduk - DeleteSatuanProduk : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Delete Satuan Produk", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse DeleteSatuanProduk - DeleteSatuanProduk : ", idsatuanproduk)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Delete Satuan Produk", idsatuanproduk)
	return ctx.JSON(http.StatusOK, result)
}
