package corporateservice

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/services"
	. "Internship-Backend/utils"
	"database/sql"
	"fmt"

	"log"
	"net/http"

	"github.com/labstack/echo"
)

type corporateService struct {
	Service services.UsecaseService
}

func NewCorporateService(service services.UsecaseService) corporateService {
	return corporateService{
		Service: service,
	}
}

// Get All Corporate
func (svc corporateService) GetListCorporate(ctx echo.Context) error {
	var result models.ResponseList
	var request models.RequestList

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data GetListCorporate : ", err.Error())
		result = ResponseListJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", constants.EMPTY_VALUE_INT, err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var orderby string

	if request.OrderBy == "id" {
		orderby = "corporate.id"
	}
	if request.OrderBy == "cid" {
		orderby = "corporate.cid"
	}
	if request.OrderBy == "uraian" {
		orderby = "corporate.uraian"
	}
	if request.OrderBy == "namaKota" {
		orderby = "corporate.nama_kota"
	}
	if request.OrderBy == "idKota" {
		orderby = "corporate.idkota"
	}
	if request.OrderBy == "namaProvinsi" {
		orderby = "corporate.nama_provinsi"
	}
	if request.OrderBy == "alamat" {
		orderby = "corporate.alamat"
	}
	if request.OrderBy == "telepon" {
		orderby = "corporate.telepon"
	}
	if request.OrderBy == "level" {
		orderby = "corporate.level"
	}
	if request.OrderBy == "gambar" {
		orderby = "corporate.gambar"
	}
	if request.OrderBy == "hirarkiId" {
		orderby = "corporate.hirarki_id"
	}
	if request.OrderBy == "ipLocalServer" {
		orderby = "corporate.iplocalserver"
	}
	if request.OrderBy == "isPercentage" {
		orderby = "corporate.ispercentage"
	}
	if request.OrderBy == "serviceFee" {
		orderby = "corporate.servicefee"
	}
	if request.OrderBy == "corporateCategory" {
		orderby = "corporate.idcorporatecategory"
	}
	if request.OrderBy == "namaCorporateCategory" {
		orderby = "corporatecategory.uraian"
	}

	if request.OrderBy == constants.EMPTY_VALUE {
		orderby = "corporate.id"
	}

	request.OrderBy = orderby

	resCorporate, err := svc.Service.CorporateRepo.GetListCorporate(request)
	if err != nil {
		log.Println("Error GetListCorporate - GetListCorporate : ", err.Error())
		result = ResponseListJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Data Corporate", constants.EMPTY_VALUE_INT, err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	page := request.Page - 1
	limitbawah := page * request.Limit
	limitatas := request.Limit * request.Page

	if limitatas > len(resCorporate) {
		limitatas = len(resCorporate)
	}

	if limitbawah > len(resCorporate) {
		log.Println("Error GetListCorporate - GetListCorporate : Limit Melebihi Jumlah Data")
		result = ResponseListJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Limit Melebihi Jumlah Data", constants.EMPTY_VALUE_INT, err)
		return ctx.JSON(http.StatusOK, result)
	}

	sliceCorporate := resCorporate[limitbawah:limitatas]

	log.Println("Success GetListCorporate - GetListCorporate : ", sliceCorporate)
	result = ResponseListJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get Data Corporate", len(resCorporate), sliceCorporate)
	return ctx.JSON(http.StatusOK, result)
}

// Get One Corporate
func (svc corporateService) GetSingleCorporate(ctx echo.Context) error {
	var result models.Response
	var request models.Corporate

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data GetSingleCorporate : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	_, exists, err := svc.Service.CorporateRepo.IsCorporateExistsByIndex(request)

	if err != nil {
		log.Println("Error GetSingleCorporate - IsCorporateExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Corporate", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error GetSingleCorporate - IsCorporateExistsByIndex : Id Corporate Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Corporate Not Found", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	resCorporate, err := svc.Service.CorporateRepo.GetSingleCorporate(request)

	if err != nil {
		log.Println("Error GetSingleCorporate - GetSingleCorporate : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Data Corporate", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse GetSingleCorporate - GetSingleCorporate : ", resCorporate)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get Data Corporate", resCorporate)
	return ctx.JSON(http.StatusOK, result)
}

// Insert Corporate
func (svc corporateService) InsertCorporate(ctx echo.Context) error {
	var result models.Response
	var parentcid models.Corporate
	var request models.RequestAddCorporate

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data InsertCorporate : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	parentcid.CID = request.ParentCID

	resParentCorporate, exists, err := svc.Service.CorporateRepo.IsCorporateExistsByIndex(parentcid)
	if err != nil {
		log.Println("Error InsertCorporate - IsCorporateExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Parent Corporate", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error InsertCorporate - IsCorporateExistsByIndex : Parent CID Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Parent CID Not Found", err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	corporate := models.RequestAddCorporate{
		Uraian:              request.Uraian,
		Alamat:              request.Alamat,
		Telepon:             request.Telepon,
		ParentCID:           request.ParentCID,
		HirarkiId:           resParentCorporate.HirarkiId,
		Level:               resParentCorporate.Level + 1,
		Gambar:              request.Gambar,
		IsPercentage:        request.IsPercentage,
		IdCorporateCategory: request.IdCorporateCategory,
		IpLocalServer:       request.IpLocalServer,
		IdKota:              request.IdKota,
		ServiceFee:          request.ServiceFee,
		CreatedAt:           TimeStampNow(),
	}

	idcorporate, err := svc.Service.CorporateRepo.InsertCorporate(corporate)

	if err != nil {
		log.Println("Error InsertCorporate - InsertCorporate : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert Corporate-", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse InsertCorporate - InsertCorporate : ", idcorporate)

	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert Corporate", idcorporate)
	return ctx.JSON(http.StatusOK, result)
}

// Update Corporate
func (svc corporateService) UpdateCorporate(ctx echo.Context) error {
	var result models.Response
	var request models.RequestUpdateCorporate
	var corporate models.Corporate

	err := BindValidateStruct(ctx, &request)
	if err != nil {
		log.Println("Error Validate Data UpdateCorporate : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	corporate.Id = request.Id
	_, exists, err := svc.Service.CorporateRepo.IsCorporateExistsByIndex(corporate)

	if err != nil {
		log.Println("Error UpdateCorporate - IsCorporateExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Corporate", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error UpdateCorporate - IsCorporateExistsByIndex : Id Corporate Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Corporate Not Found", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	idcorporate, err := svc.Service.CorporateRepo.UpdateCorporate(request)

	if err != nil {
		log.Println("Error UpdateCorporate - UpdateCorporate : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Update Corporate", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse UpdateCorporate - UpdateCorporate : ", idcorporate)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Update Corporate", idcorporate)
	return ctx.JSON(http.StatusOK, result)
}

// Delete Corporate
func (svc corporateService) DeleteCorporate(ctx echo.Context) error {
	var result models.Response
	var request models.RequestDeleteCorporate
	var corporate models.Corporate

	err := BindValidateStruct(ctx, &request)

	if err != nil {
		log.Println("Error Validate Data DeleteCorporate : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	corporate.Id = request.Id
	_, exists, err := svc.Service.CorporateRepo.IsCorporateExistsByIndex(corporate)

	if err != nil {
		log.Println("Error DeleteCorporate - IsCorporateExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Id Corporate", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		log.Println("Error DeleteCorporate - IsCorporateExistsByIndex : Id Corporate Not Found")
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Id Corporate Not Found", err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	err = DBTransaction(svc.Service.RepoDB, func(tx *sql.Tx) error {
		idCorporate, cid, err := svc.Service.CorporateRepo.DeleteCorporate(request, tx)
		if err != nil {
			log.Println("Error DeleteCorporate - DeleteCorporate : ", err.Error())
			result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Delete Corporate", nil)
			return ctx.JSON(http.StatusBadRequest, result)
		}

		var corp models.Corporate
		corp.ParentCID = cid
		_, exists, err := svc.Service.CorporateRepo.IsCorporateExistsByIndex(corp)

		if err != nil {
			log.Println("Error DeleteCorporate - IsCorporateExistsByIndex : ", err.Error())
		}

		if exists {
			idparentCorporate, err := svc.Service.CorporateRepo.DeleteParentCorporate(cid, tx)
			fmt.Println(idparentCorporate)

			if err != nil {
				log.Println("Error DeleteCorporate - DeleteParentCorporate : ", err.Error())
				result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Delete Corporate", nil)
				return ctx.JSON(http.StatusBadRequest, result)
			}

			log.Println("Reponse DeleteCorporate - DeleteParentCorporate : ", idCorporate)
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Delete Corporate", idCorporate)
			return ctx.JSON(http.StatusOK, result)
		}
		log.Println("Reponse DeleteCorporate - DeleteParentCorporate : ", idCorporate)
		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Delete Corporate", idCorporate)
		return ctx.JSON(http.StatusOK, result)
	})
	if err != nil {
		log.Println("Error DeleteCorporate - DeleteParentCorporate : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Delete Corporate", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	return nil
}
