package usersservice

import (
	config "Internship-Backend/config"
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/services"
	. "Internship-Backend/utils"
	"log"
	"net/http"

	"strconv"

	"github.com/labstack/echo"
	// "golang.org/x/crypto/bcrypt"
)

type usersService struct {
	Service services.UsecaseService
}

func NewUsersService(service services.UsecaseService) usersService {
	return usersService{
		Service: service,
	}
}

func (svc usersService) InsertUser(ctx echo.Context) error {
	var result models.Response
	var request models.RequestAddUser

	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data:  InsertUser", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}
	users := models.Users{
		Username: request.Username,
	}

	exist, err := svc.Service.UserRepo.GetUserByIndex(users)
	if err != nil {
		log.Println("Error checking Getsingleuser-username: InsertUser")
	}
	if len(exist) > 0 {
		log.Println("Username Already Exist: InsertUser", err)
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_FOUND_CODE, "Username Already Exist", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	users = models.Users{
		Email: request.Email,
	}

	exist, err = svc.Service.UserRepo.GetUserByIndex(users)
	if err != nil {
		log.Println("Error checking GetsingleUser-email: InsertUser")
	}
	if len(exist) > 0 {
		log.Println("Email Already Exist: InsertUser", err)
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_FOUND_CODE, "Email Already Exist", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	// hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	// if err != nil {
	// 	log.Println("error generate hash password", err)
	// }
	// request.Password = string(hash)
	// hashcred, err := bcrypt.GenerateFromPassword([]byte(request.Username), bcrypt.MinCost)
	// if err != nil {
	// 	log.Println("error generate hash enccredential", err)
	// }
	// request.EncCredentials = string(hashcred)

	hash := Encrypt(config.E_KEY, request.Password)
	// if err != nil {
	// 	log.Println("error generate hash password", err)
	// }
	request.Password = string(hash)
	// hashcred, err := bcrypt.GenerateFromPassword([]byte(request.Username), bcrypt.MinCost)
	// if err != nil {
	// 	log.Println("error generate hash enccredential", err)
	// }
	// hashcred := Encrypt(config.E_KEY, request.Username)
	// request.EncCredentials = string(hashcred)

	resInsertUser, err := svc.Service.UserRepo.InsertUser(request)
	if err != nil {
		log.Println("Error Insert User- InsertUser : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert User", err.Error())

		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("Reponse Insert User- InsertUser")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert User", resInsertUser)

	strid := strconv.Itoa(resInsertUser)

	encryptstrid := Encrypt(config.E_KEY, strid)
	//hash encrypted text using sha1
	hashedenc := Hash(encryptstrid)
	enchash := Encrypt(config.E_KEY, hashedenc)
	encwithhash := encryptstrid + "|" + enchash

	resUpdate, err := svc.Service.UserRepo.UpdateEnc(resInsertUser, encwithhash)

	if err != nil {
		log.Println("Error Insert User- UpdateEnc : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert User", err.Error())

		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("Reponse Insert User- UpdateEnc")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert User", resUpdate)

	return ctx.JSON(http.StatusOK, result)

}

func (svc usersService) GetAllUser(ctx echo.Context) (err error) {
	var result models.ResponseList
	var request models.RequestList

	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data: GetAllUser", err.Error())
		result = ResponseListJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), constants.EMPTY_VALUE_INT, nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var orderby string
	if request.HirarkiId == constants.EMPTY_VALUE {
		request.HirarkiId = "1000000001"
	}
	if request.OrderBy == "id" {
		orderby = "users.id"
	}
	if request.OrderBy == "nama" {
		orderby = "users.nama"
	}
	if request.OrderBy == "username" {
		orderby = "users.username"
	}
	if request.OrderBy == "jenisUser" {
		orderby = "jenisuser.uraian"
	}
	if request.OrderBy == "namaCorporate" {
		orderby = "corporate.uraian"
	}
	if request.OrderBy == "email" {
		orderby = "users.email"
	}
	if request.OrderBy == "encCredentials" {
		orderby = "users.enccredential"
	}
	if request.OrderBy == "role" {
		orderby = "roles.rolename"
	}
	if request.OrderBy == constants.EMPTY_VALUE {
		orderby = "users.id"
	}

	request.OrderBy = orderby

	resuser, err := svc.Service.UserRepo.GetAllUser(request)

	if err != nil {
		log.Println("Error Get User- Get User : ", err.Error())
		result = ResponseListJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get User", constants.EMPTY_VALUE_INT, nil)

		return ctx.JSON(http.StatusBadRequest, result)
	}
	limitbawah := (request.Page - 1) * request.Limit
	limitatas := request.Limit * request.Page

	if len(resuser) < limitatas {
		limitatas = len(resuser)
	}
	if limitbawah > len(resuser) {
		log.Println("Error Get Single Core Settlement Keys : GetListCoresettlementKeys: limit melebihi jumlah data", err)
		result = ResponseListJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Limit melebihi jumlah data", len(resuser), err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	userslice := resuser[limitbawah:limitatas]
	if len(userslice) == constants.EMPTY_VALUE_INT {
		log.Println("Error Get Single Core Settlement Keys : GetListCoresettlementKeys: limit melebihi jumlah data", err)
		result = ResponseListJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Data Tidak Ditemukan", len(resuser), err)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse Get List User- Get List User")
	result = ResponseListJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get User", len(resuser), userslice)

	return ctx.JSON(http.StatusOK, result)
}

func (svc usersService) GetUser(ctx echo.Context) (err error) {
	var result models.Response
	var request models.Users

	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	resGetUser, err := svc.Service.UserRepo.GetUserByIndex(request)
	if err != nil {

		log.Println("Error Get User- Get User : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get User", nil)

		return ctx.JSON(http.StatusBadRequest, result)
	}

	if len(resGetUser) == 0 {
		log.Println("Error Get User- Get User : User Not Found ", err)
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "User Not Found", nil)

		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("Reponse Get User- Get User")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get User", resGetUser)

	return ctx.JSON(http.StatusOK, result)

}

func (svc usersService) EditUser(ctx echo.Context) error {
	var result models.Response
	var request models.RequestUpdateUser
	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}
	// users := models.Users{
	// 	Id: request.Id,
	// }
	// exist, err := svc.Service.UserRepo.GetUserByIndex(users)
	// if err != nil {
	// 	log.Println("Error Edit User- Edit User : User Not Exist", exist)
	// 	result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "User Not Exist", nil)
	// 	return ctx.JSON(http.StatusBadRequest, result)
	// }
	// if len(exist) == constants.EMPTY_VALUE_INT {
	// 	log.Println("Error Edit User- Edit User: User Not Exist", exist)
	// 	result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "User Not Exist", nil)
	// 	return ctx.JSON(http.StatusBadRequest, result)
	// }
	// fmt.Println(request)

	// hashcred, err := bcrypt.GenerateFromPassword([]byte(request.Username), bcrypt.MinCost)
	// if err != nil {
	// 	log.Println("error generate hash enccredential", err)
	// }

	// hashcred := Encrypt(config.E_KEY, request.Username)
	// request.EncCredentials = string(hashcred)
	strid := strconv.Itoa(request.Id)

	encryptstrid := Encrypt(config.E_KEY, strid)
	//hash encrypted text using sha1
	hashedenc := Hash(encryptstrid)
	enchash := Encrypt(config.E_KEY, hashedenc)
	encwithhash := encryptstrid + "|" + enchash

	request.EncCredentials = encwithhash
	resUpdate, err := svc.Service.UserRepo.UpdateUser(request)
	if err != nil {
		log.Println("Error Edit User- Edit User : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Edit User", nil)

		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("Reponse Edit User- Edit User", result, err)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Edit User", resUpdate)

	return ctx.JSON(http.StatusOK, result)

}
func (svc usersService) EditUserPassword(ctx echo.Context) error {
	var result models.Response
	var request models.RequestUpdatePassword
	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}
	users := models.Users{
		Id: request.Id,
	}
	exist, err := svc.Service.UserRepo.GetUserByIndex(users)
	if err != nil {
		log.Println("Error Edit User- Edit User : User Not Exist", exist)
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "User Not Exist", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	if len(exist) == constants.EMPTY_VALUE_INT {
		log.Println("Error Edit User- Edit User: User Not Exist", exist)
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "User Not Exist", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	// hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	// if err != nil {
	// 	log.Println("error generate hash password", err)
	// }
	hash := Encrypt(config.E_KEY, request.Password)
	request.Password = string(hash)

	resUpdate, err := svc.Service.UserRepo.UpdateUserPassword(request)
	if err != nil {
		log.Println("Error Edit User- Edit User : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Edit User", nil)

		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("Reponse Edit User- Edit User", result, err)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Edit User", resUpdate)

	return ctx.JSON(http.StatusOK, result)

}

func (svc usersService) DeleteUser(ctx echo.Context) (err error) {
	var result models.Response
	var request models.RequestDeleteUser
	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}
	users := models.Users{
		Id: request.Id,
	}
	exist, err := svc.Service.UserRepo.GetUserByIndex(users)
	if err != nil {
		log.Println("Error Checking User-DeleteUser : User Not Exist", nil)
	}
	if len(exist) == constants.EMPTY_VALUE_INT {
		log.Println("Error Delete User- Delete User : User Not Exist", nil)
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "User Not Exist", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	insertuser, err := svc.Service.UserRepo.DeleteUser(request)
	if err != nil {
		log.Println("Error Delete User- Delete User : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Delete User", nil)

		return ctx.JSON(http.StatusBadRequest, result)
	}

	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Delete User", insertuser)
	log.Println("Reponse Delete User- Delete User", result)

	return ctx.JSON(http.StatusOK, result)

}

func (svc usersService) GetJenisUser(ctx echo.Context) (err error) {
	var result models.Response
	jenisuser, err := svc.Service.UserRepo.GetJenisUser()

	if err != nil {
		log.Println("Error Get Jenis User- Get Jenis User : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get User", nil)

		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("Reponse Get Jenis User- Get Jenis User")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get User", jenisuser)

	return ctx.JSON(http.StatusOK, result)
}
