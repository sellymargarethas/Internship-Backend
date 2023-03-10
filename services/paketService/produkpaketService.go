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

type produkpaketService struct {
	Service services.UsecaseService
}

func NewProdukPaketService(service services.UsecaseService) produkpaketService {
	return produkpaketService{
		Service: service,
	}
}

// Get All Produk Paket
func (svc produkpaketService) GetListProdukPaket(ctx echo.Context) error {
	var result models.ResponseList
	resProdukPaket, err := svc.Service.ProdukPaketRepo.GetListProdukPaket()
	if err != nil {
		log.Println("Error GetListProdukPaket - GetListProdukPaket : ", err.Error())
		result = ResponseListJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Data Produk Paket", constants.EMPTY_VALUE_INT, err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	count, err := svc.Service.ProdukPaketRepo.GetCountList()
	if err != nil {
		log.Println("Error GetListProdukPaket - GetCountList : ", err.Error())
		result = ResponseListJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Data Produk Paket", constants.EMPTY_VALUE_INT, err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("Reponse GetListProdukPaket - GetListProdukPaket : ", resProdukPaket)
	result = ResponseListJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get Data Produk Paket", count, resProdukPaket)
	return ctx.JSON(http.StatusOK, result)
}

// Insert Produk Paket
func (svc produkpaketService) InsertProdukPaket(ctx echo.Context) error {
	var result models.Response
	var request models.RequestAddProdukPaket

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data InsertProdukPaket : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	newprodukpaket, err := svc.Service.ProdukPaketRepo.InsertProdukPaket(request)

	if err != nil {
		log.Println("Error InsertProdukPaket- InsertProdukPaket : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert Produk Paket", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse InsertProdukPaket - InsertProdukPaket : ", newprodukpaket)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert Produk Paket", newprodukpaket)
	return ctx.JSON(http.StatusOK, result)
}

// Update Produk Paket
func (svc produkpaketService) UpdateProdukPaket(ctx echo.Context) error {
	var result models.Response
	var request models.RequestUpdateProdukPaket
	var produkpaket models.ProdukPaket

	err := BindValidateStruct(ctx, &request)
	if err != nil {
		log.Println("Error Validate Data UpdateProdukPaket : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	produkpaket.Id = request.Id
	_, exists, err := svc.Service.ProdukPaketRepo.IsProdukPaketExistsByIndex(produkpaket)

	if err != nil {
		log.Println("Error UpdateProdukPaket - IsProdukPaketExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Produk Paket", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error UpdateProdukPaket - IsProdukPaketExistsByIndex : Id Produk Paket Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Produk Paket Not Found", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	idProdukPaket, err := svc.Service.ProdukPaketRepo.UpdateProdukPaket(request)

	if err != nil {
		log.Println("Error UpdateProdukPaket - UpdateProdukPaket : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Update Produk Paket", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse UpdateProdukPaket - UpdateProdukPaket : ", idProdukPaket)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Update Produk Paket", idProdukPaket)
	return ctx.JSON(http.StatusOK, result)
}

// Delete Produk Paket
func (svc produkpaketService) DeleteProdukPaket(ctx echo.Context) error {
	var result models.Response
	var request models.RequestDeleteProdukPaket
	var produkpaket models.ProdukPaket

	err := BindValidateStruct(ctx, &request)
	if err != nil {
		log.Println("Error Validate Data DeleteProdukPaket : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	produkpaket.Id = request.Id
	_, exists, err := svc.Service.ProdukPaketRepo.IsProdukPaketExistsByIndex(produkpaket)

	if err != nil {
		log.Println("Error DeleteProdukPaket - IsProdukPaketExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Produk Paket", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error DeleteProdukPaket - IsProdukPaketExistsByIndex : Id Produk Paket Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Produk Paket Not Found", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	idprodukpaket, err := svc.Service.ProdukPaketRepo.DeleteProdukPaket(request)

	if err != nil {
		log.Println("Error DeleteProdukPaket - DeleteProdukPaket : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Delete Produk Paket", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse DeleteProdukPaket - DeleteProdukPaket : ", idprodukpaket)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Delete Produk Paket", idprodukpaket)
	return ctx.JSON(http.StatusOK, result)
}
