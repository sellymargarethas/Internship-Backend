package deviceService

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/services"
	. "Internship-Backend/utils"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type deviceService struct {
	Service services.UsecaseService
}

func NewDeviceService(service services.UsecaseService) deviceService {
	return deviceService{
		Service: service,
	}
}

func (svc deviceService) GetListDevice(ctx echo.Context) (err error) {
	var result models.ResponseList
	var request models.RequestList

	err = BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data GetListDevice : ", err.Error())
		result = ResponseListJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", constants.EMPTY_VALUE_INT, err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}
	if request.HirarkiId == constants.EMPTY_VALUE {
		request.HirarkiId = "1000000001"
	}

	var orderby string

	if request.OrderBy == "id" {
		orderby = "device.id"
	}
	if request.OrderBy == "deviceId" {
		orderby = "device.device_id"
	}
	if request.OrderBy == "idCorporate" {
		orderby = "device.idcorporate"
	}
	if request.OrderBy == "namaCorporate" {
		orderby = "corporate.uraian"
	}
	if request.OrderBy == "mid" {
		orderby = "device.mid"
	}
	if request.OrderBy == "tid" {
		orderby = "device.tid"
	}
	if request.OrderBy == "dkiTid" {
		orderby = "device.dkitid"
	}
	if request.OrderBy == "dkiMid" {
		orderby = "device.dkimid"
	}
	if request.OrderBy == "merchantKey" {
		orderby = "device.merchantkey"
	}
	if request.OrderBy == "CSKName" {
		orderby = "coresettlementkeys.name"
	}
	if request.OrderBy == "jenisDevice" {
		orderby = "device.jenisdevice"
	}

	if request.OrderBy == constants.EMPTY_VALUE {
		orderby = "device.id"
	}

	request.OrderBy = orderby

	resdevice, err := svc.Service.DeviceRepo.GetListDevice(request)
	if err != nil {
		log.Println("Error Get Device- Get Device : ", err)
		result = ResponseListJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Device", constants.EMPTY_VALUE_INT, nil)

		return ctx.JSON(http.StatusBadRequest, result)
	}

	page := request.Page - 1
	limitbawah := page * request.Limit
	limitatas := request.Limit * request.Page

	if limitatas > len(resdevice) {
		limitatas = len(resdevice)
	}

	if limitbawah > len(resdevice) {
		log.Println("Error GetListDevice - GetListDevice : Limit Melebihi Jumlah Data")
		result = ResponseListJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Limit Melebihi Jumlah Data", constants.EMPTY_VALUE_INT, err)
		return ctx.JSON(http.StatusOK, result)
	}

	sliceDevice := resdevice[limitbawah:limitatas]
	log.Println("Reponse Get Device- Get Device")
	result = ResponseListJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get Device", len(resdevice), sliceDevice)
	return ctx.JSON(http.StatusOK, result)
}

func (svc deviceService) GetSingleDevice(ctx echo.Context) (err error) {
	var result models.Response
	var request models.RequestGetDevice
	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data: GetSingleDevice", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}
	device := models.Device{
		Id: request.Id,
	}
	resDevice, err := svc.Service.DeviceRepo.GetSingleDevice(device)

	if err != nil {
		log.Println("Error Get Device- Get Device : ", err.Error())
		if err == sql.ErrNoRows {
			log.Println("error Get Device- Get Device:", err)
			result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Device Not Found", nil)

			return ctx.JSON(http.StatusBadRequest, result)

		}
		log.Println("error Get Device- Get Device:", err)
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Device", nil)

		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("Reponse Get Device- Get Device: ", request.Id)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get Device", resDevice)

	return ctx.JSON(http.StatusOK, result)
}

func (svc deviceService) InsertDevice(ctx echo.Context) (err error) {
	var result models.Response
	var request models.RequestAddDevice
	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data: InsertDevice", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}
	device := models.Device{
		DeviceId: request.DeviceId,
	}
	checkdevice, err := svc.Service.DeviceRepo.GetSingleDevice(device)
	if err != nil {
		log.Println("Error GetSingleDevice:InsertDevice ", err)

	}
	if checkdevice.DeviceId != constants.EMPTY_VALUE {
		log.Println("Reponse ERROR GetSingleDevice -  InsertDevice: Device Id Already Registered", err)
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Device Id Already Registered", checkdevice.DeviceId)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	corporatecheck := models.Corporate{
		Id: request.IdCorporate,
	}
	resCorporate, exist, err := svc.Service.CorporateRepo.IsCorporateExistsByIndex(corporatecheck)
	if err != nil {
		log.Println("Error IsCorporateExistsByIndex  :  InsertDevice ", err.Error())
	}
	if !exist {
		log.Println("Reponse ERROR IsCorporateExistsByIndex -  InsertDevice: Corporate ID Not Found", err)
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "CORPORATE ID NOT FOUND", resCorporate)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	coresettlementcheck := models.CoreSettlementKeys{
		Id: request.IdCorporateMerchantkey,
	}

	coresettlement, err := svc.Service.CoreSettlementRepo.IsCoresettlementkeysExistByIndex(coresettlementcheck)
	if err != nil {
		log.Println("Error IsCoresettlementkeysExistByIndex  :  InsertDevice ", err.Error())
	}
	if !coresettlement {
		log.Println("Error Insert device- Insert device : id merchant key not found ", err)
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert device: Id merchant key not Found", nil)

		return ctx.JSON(http.StatusBadRequest, result)
	}
	resInsertDevice, err := svc.Service.DeviceRepo.InsertDevice(request)
	if err != nil {
		log.Println("Error Insert device- Insert device : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert device", err.Error())

		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("Reponse Insert device- Insert device: Success")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert device", resInsertDevice)

	return ctx.JSON(http.StatusOK, result)

}

func (svc deviceService) UpdateDevice(ctx echo.Context) (err error) {
	var result models.Response
	var request models.RequestUpdateDevice
	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data: UpdateDevice", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}
	device := models.Device{
		Id: request.Id,
	}
	checkdevice, err := svc.Service.DeviceRepo.GetSingleDevice(device)
	if err != nil {
		log.Println("Error GetSingleDevice:UpdateDevice ", err)

	}
	if checkdevice.DeviceId == constants.EMPTY_VALUE {
		log.Println("Reponse ERROR GetSingleDevice -  UpdateDevice: Device Not Found", err)
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Device Not Found", checkdevice.DeviceId)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	corporatecheck := models.Corporate{
		Id: request.IdCorporate,
	}
	resCorporate, exist, err := svc.Service.CorporateRepo.IsCorporateExistsByIndex(corporatecheck)
	if err != nil {
		log.Println("Error IsCorporateExistsByIndex  :  InsertDevice ", err.Error())
	}
	if !exist {
		log.Println("Reponse ERROR IsCorporateExistsByIndex -  InsertDevice: Corporate ID Not Found", err)
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "CORPORATE ID NOT FOUND", resCorporate)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	coresettlementcheck := models.CoreSettlementKeys{
		Id: request.IdCorporateMerchantkey,
	}

	coresettlement, err := svc.Service.CoreSettlementRepo.IsCoresettlementkeysExistByIndex(coresettlementcheck)
	if err != nil {
		log.Println("Error IsCoresettlementkeysExistByIndex  :  EditDevice ", err.Error())
	}
	if !coresettlement {
		log.Println("Error Update device- Update device : id merchant key not found ", err)
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Failed Insert device: Id merchant key not Found", nil)

		return ctx.JSON(http.StatusBadRequest, result)
	}

	resupdatedevice, err := svc.Service.DeviceRepo.EditDevice(request)
	if err != nil {
		log.Println("Error UpdateDevice ", err, resupdatedevice)
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "FAILED Edit device", err)

		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("Reponse edit device- edit device")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Edit device", resupdatedevice)

	return ctx.JSON(http.StatusOK, result)
}

func (svc deviceService) DeleteDevice(ctx echo.Context) (err error) {
	var result models.Response
	var request models.RequestDeleteDevice
	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data:  DeleteDevice", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data: DeleteDevice", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}
	device := models.Device{
		Id: request.Id,
	}
	checkdevice, err := svc.Service.DeviceRepo.GetSingleDevice(device)
	if err != nil {
		log.Println("Error GetSingleDevice:DeleteDevice ", err)

	}
	if checkdevice.DeviceId == constants.EMPTY_VALUE {
		log.Println("Reponse ERROR GetSingleDevice -  InsertDevice: Device Not Found", err)
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Device Not Found", checkdevice.DeviceId)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	resdevice, err := svc.Service.DeviceRepo.DeleteDevice(request)
	if err != nil {
		log.Println("Error Delete Device", err, resdevice)
		return err
	}
	log.Println("Reponse Delete device- Delete device")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Delete device", resdevice)

	return ctx.JSON(http.StatusOK, result)

}
