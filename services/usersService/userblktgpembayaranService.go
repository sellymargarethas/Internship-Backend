package usersservice

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

type blktgService struct {
	Service services.UsecaseService
}

func NewblktgService(service services.UsecaseService) blktgService {
	return blktgService{
		Service: service,
	}
}
func (svc blktgService) GetBlktgPembayaranByIdUser(ctx echo.Context) (err error) {
	var result models.Response
	var request models.RequestGetBlktgPembayaranByIdUser

	if err := BindValidateStruct(ctx, &request); err != nil {
		log.Println("Error Validate Data: GetBlktgPembayaranByIdUser", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	resBlktgPembayaran, err := svc.Service.BlktgRepo.GetBlktgPembayaran(request)
	if err != nil {
		log.Println("Error Get BlktgPembayaran: GetBlktgPembayaranByIdUser", err)
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, constants.EMPTY_VALUE, err)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	res := models.ResponseListBlktgPembayaran{
		IdUser:          request.IdUser,
		BlktgPembayaran: resBlktgPembayaran,
	}
	log.Println("Reponse Get blacklisted paymentkategori ", "Success Get Blacklisted Paymentkategori", res)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, constants.EMPTY_VALUE, res)
	return ctx.JSON(http.StatusOK, result)
}

func (svc blktgService) UpdateBlktgPembayaran(ctx echo.Context) (err error) {
	var result models.Response
	var request models.RequestUpdateBlktgPembayaran
	var iduser models.BlktgPembayaran

	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data: UpdateBlktgPembayaran", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	iduser.IdUser = request.IdUser

	_, exists, err := svc.Service.BlktgRepo.IsBlkgPembayaranExistsByIndex(iduser)
	if err != nil {
		log.Println("Error UpdateBlktgPembayaran - IsBlkgPembayaranExistsByIndex : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Parent Corporate", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if !exists {
		resBlktgPembayaran, err := svc.Service.BlktgRepo.InsertBlktgPembayaran(request)
		fmt.Println(resBlktgPembayaran)
		if err != nil {
			log.Println("Error UpdateBlktgPembayaran - InsertBlktgPembayaran : ", err.Error())
			result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert Blacklisted Payment", nil)

			return ctx.JSON(http.StatusBadRequest, result)
		}
		log.Println("Reponse UpdateBlktgPembayaran - InsertBlktgPembayaran")
		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert Blaclisted Payment", request.IdUser)
	}

	err = DBTransaction(svc.Service.RepoDB, func(tx *sql.Tx) error {
		if len(request.IdKategoriPembayaran) == constants.EMPTY_VALUE_INT {
			resDelete, err := svc.Service.BlktgRepo.DeleteBlktgPembayaran(request.IdUser, tx)
			fmt.Println(resDelete)
			if err != nil {
				log.Println("Error UpdateBlktgPembayaran - DeleteBlktgPembayaran : ", err.Error())
				result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Update Corporate Payment", err)
				return ctx.JSON(http.StatusBadRequest, result)
			}
			log.Println("Reponse UpdateBlktgPembayaran - UpdateBlktgPembayaran")
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "User Blacklist Berhasil dihapus", request.IdUser)
			return ctx.JSON(http.StatusOK, result)
		}

		resDelete, err := svc.Service.BlktgRepo.DeleteBlktgPembayaran(request.IdUser, tx)
		fmt.Println(resDelete)
		if err != nil {
			log.Println("Error UpdateBlktgPembayaran - DeleteBlktgPembayaran : ", err.Error())
			result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Update UserBlktgPembayaran", err)
			return ctx.JSON(http.StatusBadRequest, result)
		}

		resUpdate, err := svc.Service.BlktgRepo.UpdateBlktgPembayaran(request, tx)
		fmt.Println(resUpdate)
		if err != nil {
			log.Println("Error UpdateBlktgPembayaran - UpdateBlktgPembayaran : ", err.Error())
			result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Update UserBlktgPembayaran", err)
			return ctx.JSON(http.StatusBadRequest, result)
		}
		log.Println("Reponse UpdateBlktgPembayaran - UpdateBlktgPembayaran")
		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Update UserBlktgPembayaran", request.IdUser)
		return ctx.JSON(http.StatusOK, result)
	})

	return nil

}
