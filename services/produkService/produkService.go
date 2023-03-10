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

type produkService struct {
	Service services.UsecaseService
}

func NewProdukService(service services.UsecaseService) produkService {
	return produkService{
		Service: service,
	}
}

// Get All Produk
func (svc produkService) GetListProduk(ctx echo.Context) error {
	var result models.ResponseList
	var request models.RequestList

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data GetListProduk : ", err.Error())
		result = ResponseListJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", constants.EMPTY_VALUE_INT, err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var orderby string

	if request.OrderBy == "id" {
		orderby = "produk.id"
	}

	if request.OrderBy == "kodeProduk" {
		orderby = "produk.kodeproduk"
	}
	if request.OrderBy == "nama" {
		orderby = "produk.nama"
	}
	if request.OrderBy == "hargaJual" {
		orderby = "produk.hargajual"
	}
	if request.OrderBy == "hargaRombongan" {
		orderby = "produk.hargarombongan"
	}
	if request.OrderBy == "minRombongan" {
		orderby = "produk.minrombongan"
	}
	if request.OrderBy == "gambar" {
		orderby = "produk.gambar"
	}
	if request.OrderBy == "namaCorporate" {
		orderby = "corporate.uraian"
	}
	if request.OrderBy == "namaSatuanProduk" {
		orderby = "satuanproduk.uraian"
	}
	if request.OrderBy == "currentStok" {
		orderby = "produk.currentstok"
	}
	if request.OrderBy == "jenis" {
		orderby = "produk.jenis"
	}
	if request.OrderBy == "statusPaket" {
		orderby = "produk.statuspaket"
	}
	if request.OrderBy == "statusStok" {
		orderby = "produk.statusstok"
	}

	if request.OrderBy == constants.EMPTY_VALUE {
		orderby = "produk.id"
	}

	request.OrderBy = orderby

	resProduk, err := svc.Service.ProdukRepo.GetListProduk(request)

	if err != nil {
		log.Println("Error GetListProduk - GetListProduk: ", err.Error())
		result = ResponseListJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Data Produk", constants.EMPTY_VALUE_INT, err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	page := request.Page - 1
	limitbawah := page * request.Limit
	limitatas := request.Limit * request.Page

	if limitatas > len(resProduk) {
		limitatas = len(resProduk)
	}

	if limitbawah > len(resProduk) {
		log.Println("Error GetListProduk - GetListProduk : Limit Melebihi Jumlah Data")
		result = ResponseListJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Limit Melebihi Jumlah Data", constants.EMPTY_VALUE_INT, err)
		return ctx.JSON(http.StatusOK, result)
	}

	sliceProduk := resProduk[limitbawah:limitatas]

	log.Println("Reponse GetListProduk - GetListProduk : ", resProduk)
	result = ResponseListJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get Data Produk", len(resProduk), sliceProduk)
	return ctx.JSON(http.StatusOK, result)
}

func (svc produkService) GetSingleProduk(ctx echo.Context) error {
	var result models.Response
	var request models.Produk

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data GetSingleProduk : ", err)
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	_, exists, err := svc.Service.ProdukRepo.IsProdukExistsByIndex(request)

	if err != nil {
		log.Println("Error GetSingleProduk - IsProdukExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Produk", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error GetSingleProduk - IsProdukExistsByIndex : Id Produk Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Produk Not Found", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	resProduk, err := svc.Service.ProdukRepo.GetSingleProduk(request.Id)

	if err != nil {
		log.Println("Error GetSingleProduk - GetSingleProduk : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Data Produk", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse GetSingleProduk - GetSingleProduk : ", resProduk)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get Data Produk", resProduk)
	return ctx.JSON(http.StatusOK, result)
}

// Insert Produk
func (svc produkService) InsertProduk(ctx echo.Context) error {
	var result models.Response
	var request models.RequestAddProduk
	var kodeproduk models.Produk

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data InsertProduk : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	kodeproduk.KodeProduk = request.KodeProduk
	_, exists, err := svc.Service.ProdukRepo.IsProdukExistsByIndex(kodeproduk)
	if err != nil {
		log.Println("Error InsertProduk - IsPaketExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Kode Produk", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if exists {
		log.Println("Error InsertProduk - IsPaketExistsByIndex : Kode Produk Already Exists")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_FOUND_CODE, "Kode Produk Already Exists", err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	produk := models.RequestAddProduk{
		KodeProduk:       request.KodeProduk,
		Nama:             request.Nama,
		HargaJual:        request.HargaJual,
		Gambar:           request.Gambar,
		Jenis:            request.Jenis,
		IdCorporate:      request.IdCorporate,
		IdSatuanProduk:   request.IdSatuanProduk,
		IdKategoriProduk: request.IdKategoriProduk,
		HargaRombongan:   request.HargaRombongan,
		MinRombongan:     request.MinRombongan,
		StatusStok:       request.StatusStok,
		CurrentStok:      request.CurrentStok,
		StatusPaket:      request.StatusStok,
		CreatedAt:        TimeStampNow(),
	}

	idproduk, err := svc.Service.ProdukRepo.InsertProduk(produk)

	if err != nil {
		log.Println("Error InsertProduk - InsertProduk : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert Produk", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse InsertProduk - InsertProduk : ", idproduk)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert Produk", idproduk)
	return ctx.JSON(http.StatusOK, result)
}

// Update  Produk
func (svc produkService) UpdateProduk(ctx echo.Context) error {
	var result models.Response
	var request models.RequestUpdateProduk
	var produk models.Produk

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data UpdateProduk : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	produk.Id = request.Id
	_, exists, err := svc.Service.ProdukRepo.IsProdukExistsByIndex(produk)

	if err != nil {
		log.Println("Error UpdateProduk - IsProdukExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Produk", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error UpdateProduk - IsProdukExistsByIndex : Id Produk Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Produk Not Found", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	idproduk, err := svc.Service.ProdukRepo.UpdateProduk(request)

	if err != nil {
		log.Println("Error UpdateProduk- UpdateProduk : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Update Produk", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse UpdateProduk - UpdateProduk : ", idproduk)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Update Produk", idproduk)
	return ctx.JSON(http.StatusOK, result)
}

// Delete  Produk
func (svc produkService) DeleteProduk(ctx echo.Context) error {
	var result models.Response
	var request models.RequestDeleteProduk
	var produk models.Produk

	err := BindValidateStruct(ctx, &request)
	if err != nil {
		log.Println("Error Validate Data DeleteProduk : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	produk.Id = request.Id
	_, exists, err := svc.Service.ProdukRepo.IsProdukExistsByIndex(produk)

	if err != nil {
		log.Println("Error DeleteProduk - IsProdukExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Produk", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error DeleteProduk - IsProdukExistsByIndex : Id Produk Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Produk Not Found", err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	idProduk, err := svc.Service.ProdukRepo.DeleteProduk(request)

	if err != nil {
		log.Println("Error DeleteProduk - DeleteProduk : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Delete Produk", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse DeleteProduk - DeleteProduk : ", idProduk)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Delete Produk", idProduk)
	return ctx.JSON(http.StatusOK, result)
}
