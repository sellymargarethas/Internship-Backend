package usersservice

import (
	"Internship-Backend/config"
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/services"
	. "Internship-Backend/utils"
	"log"
	"net/http"

	// "golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo"
)

type loginService struct {
	Service services.UsecaseService
}

func NewloginService(service services.UsecaseService) loginService {
	return loginService{
		Service: service,
	}
}

func (svc loginService) Login(ctx echo.Context) error {
	var result models.Response
	var request models.RequestLogin

	if err := BindValidateStruct(ctx, &request); err != nil {
		log.Println("Error Validate Login: Login", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	reqLogin := models.Login{
		Username: request.Username,
		Password: request.Password,
	}

	passUser, jenisUser, err := svc.Service.LoginRepo.CheckLogin(reqLogin)
	if err != nil {
		log.Println("Error Login: Check login", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "User Not Exist", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	if jenisUser != constants.USER_DASHBOARD {
		log.Println("Error Login: Check login: Unauthorized User", err)
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Wrong Username/Password ", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if passUser == constants.EMPTY_VALUE {
		log.Println("Error Login: User Not Found", err)
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Wrong Username/Password ", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	// err = bcrypt.CompareHashAndPassword([]byte(passUser), []byte(request.Password))
	// if err != nil {
	// 	log.Println("ERROR Login: Password ", err.Error())
	// 	result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, "Wrong Username/Password Salah", nil)
	// 	return ctx.JSON(http.StatusBadRequest, result)
	// }

	password := Decrypt(config.E_KEY, passUser)
	// fmt.Println(password)

	if password == request.Password {
		token, err := GenerateToken(reqLogin)
		if err != nil {
			log.Println("Error Generating Token: Login", err.Error())
			result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Error Generating Token", nil)
			return ctx.JSON(http.StatusBadRequest, result)
		}

		resLogin, err := svc.Service.LoginRepo.LoginReturn(reqLogin)
		if err != nil {
			log.Println("Error Get Login Return: Login", err.Error())
			result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Error Getting Information", nil)
			return ctx.JSON(http.StatusBadRequest, result)
		}
		idrole, err := svc.Service.RoleRepo.GetRoleIdByIdUser(resLogin.Id)
		if err != nil {
			// log.Println("ID ROLE NOT FOUND", err.Error())
			log.Println("Error Get Login Return: ID ROLE NOT FOUND", err.Error())
			result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "ID ROLE NOT FOUND", nil)
			return ctx.JSON(http.StatusBadRequest, result)
		}

		resrole := models.RequestSingleRole{
			IdRole: idrole,
		}
		roles, err := svc.Service.RoleRepo.GetSingleRole(resrole)
		if err != nil {
			log.Println("Error Get Login Return: ROLE NOT FOUND", err.Error())
			result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "ROLE NOT FOUND", nil)
			return ctx.JSON(http.StatusBadRequest, result)
		}
		privilege, err := svc.Service.RoleRepo.GetSingleRoleDetails(resrole)
		if err != nil {
			log.Println("Error Get Login Return: ROLE DETAILS NOT FOUND", err.Error())
			result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "ROLE DETAILS NOT FOUND", nil)
			return ctx.JSON(http.StatusBadRequest, result)
		}
		resLogin.Privilege = privilege
		resLogin.Role = roles.Role

		resLogin.Token = token
		log.Println("Reponse Login ", "Success Login", err)
		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Berhasil Login", resLogin)
		return ctx.JSON(http.StatusOK, result)
	}
	if password != request.Password {

		log.Println("Error Login: Wrong Password", err)
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Wrong Username/Password ", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	return nil
}
