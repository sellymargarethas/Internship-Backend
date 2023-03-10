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

type vendorpembayaranService struct {
	Service services.UsecaseService
}

func NewVendorPembayaranService(service services.UsecaseService) vendorpembayaranService {
	return vendorpembayaranService{
		Service: service,
	}
}

// Get All Vendor Pembayaran
func (svc vendorpembayaranService) GetListVendorPembayaran(ctx echo.Context) error {
	var result models.Response
	resVendorPembayaran, err := svc.Service.VendorPembayaranRepo.GetListVendorPembayaran()

	if err != nil {
		log.Println("Error GetListVendorPembayaran - GetListVendorPembayaran : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Data Vendor Pembayaran", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse GetListVendorPembayaran - GetListVendorPembayaran : ", resVendorPembayaran)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get Data Vendor Pembayaran", resVendorPembayaran)
	return ctx.JSON(http.StatusOK, result)
}

// Insert Vendor Pembayaran
func (svc vendorpembayaranService) InsertVendorPembayaran(ctx echo.Context) error {
	var result models.Response
	var request models.RequestAddVendorPembayaran
	var uraian models.VendorPembayaran

	err := BindValidateStruct(ctx, &request)

	if err != nil {

		log.Println("Error Validate Data InsertVendorPembayaran : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	uraian.Uraian = request.Uraian
	_, exists, err := svc.Service.VendorPembayaranRepo.IsVendorPembayaranExistsByIndex(uraian)

	if err != nil {
		log.Println("Error InsertVendorPembayaran - IsVendorPembayaranExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Uraian", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if exists {
		log.Println("Error InsertVendorPembayaran - IsVendorPembayaranExistsByIndex : Uraian Already Exists")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_FOUND_CODE, "Uraian Already Exists", err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	vendorpembayaran := models.RequestAddVendorPembayaran{
		Uraian:    request.Uraian,
		CreatedAt: TimeStampNow(),
	}

	idvendorpembayaran, err := svc.Service.VendorPembayaranRepo.InsertVendorPembayaran(vendorpembayaran)

	if err != nil {
		log.Println("Error InsertVendorPembayaran- InsertVendorPembayaran : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert Vendor Pembayaran", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse  InsertVendorPembayaran - InsertVendorPembayaran : ", idvendorpembayaran)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert Vendor Pembayaran", idvendorpembayaran)
	return ctx.JSON(http.StatusOK, result)
}

// Update Vendor Pembayaran
func (svc vendorpembayaranService) UpdateVendorPembayaran(ctx echo.Context) error {
	var result models.Response
	var request models.RequestUpdateVendorPembayaran
	var vendorpembayaran models.VendorPembayaran

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data UpdateVendorPembayaran : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	vendorpembayaran.Id = request.Id
	_, exists, err := svc.Service.VendorPembayaranRepo.IsVendorPembayaranExistsByIndex(vendorpembayaran)

	if err != nil {
		log.Println("Error UpdateVendorPembayaran - IsVendorPembayaranExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Vendor Pembayaran", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error UpdateVendorPembayaran - IsVendorPembayaranExistsByIndex : Id Vendor Pembayaran Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Vendor Pembayaran Not Found", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	idvendorpembayaran, err := svc.Service.VendorPembayaranRepo.UpdateVendorPembayaran(request)

	if err != nil {
		log.Println("Error UpdateVendorPembayaran - UpdateVendorPembayaran : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Update Vendor Pembayaran", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse UpdateVendorPembayaran- UpdateVendorPembayaran : ", idvendorpembayaran)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Update Vendor", idvendorpembayaran)
	return ctx.JSON(http.StatusOK, result)
}

// Delete Vendor Pembayaran
func (svc vendorpembayaranService) DeleteVendorPembayaran(ctx echo.Context) error {
	var result models.Response
	var request models.RequestDeleteVendorPembayaran
	var vendorpembayaran models.VendorPembayaran

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data DeleteVendorPembayaran : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	vendorpembayaran.Id = request.Id
	_, exists, err := svc.Service.VendorPembayaranRepo.IsVendorPembayaranExistsByIndex(vendorpembayaran)

	if err != nil {
		log.Println("Error DeleteVendorPembayaran - IsVendorPembayaranExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Vendor Pembayaran", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error DeleteVendorPembayaran - IsVendorPembayaranExistsByIndex : Id Vendor Pembayaran Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Vendor Pembayaran Not Found", err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	idvendorpembayaran, err := svc.Service.VendorPembayaranRepo.DeleteVendorPembayaran(request)

	if err != nil {
		log.Println("Error DeleteVendorPembayaran - DeleteVendorPembayaran : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Delete Vendor Pembayaran", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse DeleteVendorPembayaran - DeleteVendorPembayaran : ", idvendorpembayaran)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Delete Vendor Pembayaran", idvendorpembayaran)
	return ctx.JSON(http.StatusOK, result)
}
