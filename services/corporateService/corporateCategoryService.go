package corporateservice

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/services"
	. "Internship-Backend/utils"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type corporateCategoryService struct {
	Service services.UsecaseService
}

func NewCorporateCategoryService(service services.UsecaseService) corporateCategoryService {
	return corporateCategoryService{
		Service: service,
	}
}

func (svc corporateCategoryService) GetCorporateCategory(ctx echo.Context) (err error) {
	var result models.Response
	rescorporatecategory, err := svc.Service.CorporateCategoryRepo.GetCorporateCategoryList()
	if err != nil {
		log.Println("Error GET corporate Category : GET corporate Category ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed get corporate Category", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("Reponse get corporate Category")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get corporate Category", rescorporatecategory)
	return ctx.JSON(http.StatusOK, result)

}

func (svc corporateCategoryService) InsertCorporateCategory(ctx echo.Context) (err error) {
	var result models.Response
	var request models.RequestAddCorporateCategory
	// var response models.ResponseCorporateCategory

	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data: InsertCorporateCategory", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}
	category := models.CorporateCategory{
		Kode: request.Kode,
	}

	exist, err := svc.Service.CorporateCategoryRepo.IsCorporateCategoryExistbyIndex(category)
	if err != nil {
		log.Println("Error IsCorporateCategoryExistbyIndex : InsertCorporateCategory ", err.Error())
	}
	if exist == constants.TRUE_VALUE {
		log.Println("Error Insert Corporate Category- IsCorporateCategoryExistbyIndex : ", err)
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert Corporate Category: Category Already Exist", err)

		return ctx.JSON(http.StatusBadRequest, result)
	}

	corporatecategory, err := svc.Service.CorporateCategoryRepo.InsertCorporateCategory(request)
	if err != nil {
		log.Println("Error Insert Corporate Category- Insert Corporate Category : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert Corporate Category", err.Error())

		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse Insert Corporate Category- Insert Corporate Category")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert Corporate Category", corporatecategory)

	return ctx.JSON(http.StatusOK, result)
}

func (svc corporateCategoryService) UpdateCorporateCategory(ctx echo.Context) (err error) {
	var result models.Response
	var request models.RequestUpdateCorporateCategory

	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data: UpdateCorporateCategory", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	category := models.CorporateCategory{
		IdCategory: request.IdCategory,
	}
	exist, err := svc.Service.CorporateCategoryRepo.IsCorporateCategoryExistbyIndex(category)
	if err != nil {
		log.Println("Error IsCorporateCategoryExistbyIndex : UpdateCorporateCategory ", err.Error())
	}
	if !exist {
		log.Println("Error Update Corporate Category- Corporate Category Not Found : ", err)
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Update Corporate Category: Corporate Category Not Found", err)

		return ctx.JSON(http.StatusBadRequest, result)
	}
	corporatecategory, err := svc.Service.CorporateCategoryRepo.EditCorporateCategory(request)
	if err != nil {
		log.Println("Error Update Corporate Category- Update Corporate Category : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Update Corporate Category", err.Error())

		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse Update Corporate Category- Update Corporate Category")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Update Corporate Category", corporatecategory)

	return ctx.JSON(http.StatusOK, result)

}

func (svc corporateCategoryService) DeleteCorporateCategory(ctx echo.Context) (err error) {
	var result models.Response
	var request models.DeleteCorporateCategory

	if err := BindValidateStruct(ctx, &request); err != nil {
		log.Println("Error Validate Data: DeleteCorporateCategory", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}
	category := models.CorporateCategory{
		IdCategory: request.IdCategory,
	}
	exist, err := svc.Service.CorporateCategoryRepo.IsCorporateCategoryExistbyIndex(category)
	if err != nil {
		log.Println("Error IsCorporateCategoryExistbyIndex : DeleteCorporateCategory ", err.Error())
	}
	if !exist {
		log.Println("Error Delete Corporate Category- Corporate Category Not Found : ", err)
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Delete Corporate Category: Corporate Category Not Found", err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	corporatecategory, err := svc.Service.CorporateCategoryRepo.DeleteCorporateCategory(request)
	if err != nil {
		log.Println("Error Delete Corporate Category- Delete Corporate Category : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Delete Corporate Category", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse Delete Corporate Category- Delete Corporate Category")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Delete Corporate Category", corporatecategory)
	return ctx.JSON(http.StatusOK, result)
}
