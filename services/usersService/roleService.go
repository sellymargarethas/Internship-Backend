package usersservice

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/services"
	"log"
	"net/http"

	. "Internship-Backend/utils"

	"github.com/labstack/echo"
)

type roleService struct {
	Service services.UsecaseService
}

func NewRoleService(service services.UsecaseService) roleService {
	return roleService{
		Service: service,
	}
}

func (svc roleService) AddUserRole(ctx echo.Context) error {
	var result models.Response
	var request models.RequestAddUsersRole

	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data:  InsertUserRole", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	err := svc.Service.RoleRepo.AddUserRole(request)
	if err != nil {
		log.Println("Error Insert User Role- InsertUserRole : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert User Role", err.Error())

		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("Reponse Insert User Role- InsertUser")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert User Role", request.Role)

	return ctx.JSON(http.StatusOK, result)
}

func (svc roleService) GetListRole(ctx echo.Context) error {
	var result models.Response
	var request models.RequestListRole

	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data:  ListUserRole-requestlist", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}
	if request.HirarkiId == constants.EMPTY_VALUE {
		request.HirarkiId = "1000000001"
	}
	if request.OrderBy == constants.EMPTY_VALUE {
		request.OrderBy = "roles.id"
	}
	if request.OrderBy == "idCorporate" {
		request.OrderBy = "roles.idcorporate"
	}
	if request.OrderBy == "namaCorporate" {
		request.OrderBy = "corporate.uraian"
	}
	if request.OrderBy == "roles" {
		request.OrderBy = "roles.rolename"
	}
	if request.OrderBy == "id" {
		request.OrderBy = "roles.id"
	}

	res, err := svc.Service.RoleRepo.GetListRole(request)
	if err != nil {
		log.Println("Error Get List User Role- Get ListUserRole : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get List User Role", err.Error())

		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("Reponse Get List Role- Get List Role ")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get List role", res)

	return ctx.JSON(http.StatusOK, result)
}

func (svc roleService) GetSingleRole(ctx echo.Context) error {
	var result models.Response
	var request models.RequestSingleRole

	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data:  SingleUserRole-requestSingle", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}
	res, err := svc.Service.RoleRepo.GetSingleRole(request)
	if err != nil {
		log.Println("Error Get Single User Role- Get SingleUserRole : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Single User Role", err.Error())

		return ctx.JSON(http.StatusBadRequest, result)
	}
	res.Privilege, err = svc.Service.RoleRepo.GetSingleRoleDetails(request)
	if err != nil {
		log.Println("Error Get Single User Role Details- Get SingleUserRole Details : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Single User Role Details", err.Error())

		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("Reponse Get Single Role- Get Single Role ")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get Single role", res)

	return ctx.JSON(http.StatusOK, result)
}
func (svc roleService) UpdateUserRole(ctx echo.Context) error {
	var result models.Response
	var request models.RequestUpdateRole
	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data:  UpdateUserRole", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}
	// fmt.Println(request)
	status, err := svc.Service.RoleRepo.UpdateRole(request)
	if err != nil {
		log.Println("Error Update User Role- UpdateUserRole : ", err.Error())
		result = ResponseJSON(status, constants.FAILED_CODE, "Failed Update User Role", err.Error())

		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("Reponse Update User Role- UpdateUser")
	result = ResponseJSON(status, constants.SUCCESS_CODE, "Success Update User Role", request.Role)

	return ctx.JSON(http.StatusOK, result)
}
func (svc roleService) DeleteUserRole(ctx echo.Context) error {
	var result models.Response
	var request models.RequestSingleRole
	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data:  DeleteUserRole", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}
	status, err := svc.Service.RoleRepo.DeleteRole(request)
	if err != nil {
		log.Println("Error Delete User Role- DeleteUserRole : ", err.Error())
		result = ResponseJSON(status, constants.FAILED_CODE, "Failed Delete User Role", err.Error())

		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("Reponse Delete User Role- DeleteUser")
	result = ResponseJSON(status, constants.SUCCESS_CODE, "Success Delete User ROle", nil)

	return ctx.JSON(http.StatusOK, result)
}

func (svc roleService) GetAllRoleMenu(ctx echo.Context) error {
	var result models.Response

	rolelist, err := svc.Service.RoleRepo.GetAllRoleMenu()
	if err != nil {
		log.Println("Error Get User Role- Role Menu : ", err.Error())
		result = ResponseJSON(constants.TRUE_VALUE, constants.FAILED_CODE, "Failed Get User Role Menu", err.Error())

		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("Reponse Get User Role- Role Menu")
	result = ResponseJSON(constants.FALSE_VALUE, constants.SUCCESS_CODE, "Success Get User ROle Menu", rolelist)

	return ctx.JSON(http.StatusOK, result)
}

func (svc roleService) GetAllRoleTask(ctx echo.Context) error {
	var result models.Response

	roletask, err := svc.Service.RoleRepo.GetAllRoleTask()
	// fmt.Println(roletask)
	if err != nil {
		log.Println("Error Get User Role- RoleTask : ", err.Error())
		result = ResponseJSON(constants.TRUE_VALUE, constants.FAILED_CODE, "Failed Get User Role Task", err.Error())

		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("Reponse Get User Role- RoleTask")
	result = ResponseJSON(constants.FALSE_VALUE, constants.SUCCESS_CODE, "Success Get User Role Task", roletask)

	return ctx.JSON(http.StatusOK, result)
}
