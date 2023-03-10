package dashboardService

import (
	"Internship-Backend/constants"
	"Internship-Backend/models"
	"Internship-Backend/services"
	. "Internship-Backend/utils"
	"database/sql"
	"fmt"
	"time"

	"log"
	"net/http"

	"github.com/labstack/echo"
)

type dashboardSummaryService struct {
	Service services.UsecaseService
}

func NewDashboardSummaryService(service services.UsecaseService) dashboardSummaryService {
	return dashboardSummaryService{
		Service: service,
	}
}

func (svc dashboardSummaryService) GetDailyRevenue(ctx echo.Context) error {
	var result models.Response
	var request models.RequestRevenue
	timeNow := time.Now()

	err := BindValidateStruct(ctx, &request)
	log.Println("bind request daily revenue")
	fmt.Println("Bind request time start :", timeNow)
	if err != nil {
		log.Println("Error Bind Data GetDailyRevenue : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Bind Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}

	err = DBTransaction(svc.Service.RepoDB, func(tx *sql.Tx) error {
		if request.Corporate == constants.EMPTY_VALUE {
			request.Corporate = "1000000001"
		}
		request.Corporate = request.Corporate + "%"
		// fmt.Println(request.Corporate)
		resdashboard, err := svc.Service.DashboardRepo.GetDailyRevenue(request, tx)
		log.Println("get daily revenue")
		// fmt.Println(resdashboard)
		elapsed := time.Since(timeNow).Milliseconds()
		fmt.Println("GetDailyRevenue - GetDailyRevenue elapsed time millisecond :", elapsed)
		if err != nil {
			log.Println("Error GetDailyRevenue : GetDailyRevenue", err)

			result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "ERROR GetDailyRevenue", constants.EMPTY_VALUE_INT)
			return err
		}
		log.Println("finished getting daily revenue")
		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success GetDailyRevenue", resdashboard)
		return err
	})

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result)
	}

	return ctx.JSON(http.StatusOK, result)
}

func (svc dashboardSummaryService) GetWeeklyRevenue(ctx echo.Context) error {
	var result models.Response
	var request models.RequestRevenue
	timeNow := time.Now()

	err := BindValidateStruct(ctx, &request)
	if err != nil {
		log.Println("Error Bind Data GetWeeklyRevenue : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Bind Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}
	fmt.Println("Bind request time start :", timeNow)

	err = DBTransaction(svc.Service.RepoDB, func(tx *sql.Tx) error {
		if request.Corporate == constants.EMPTY_VALUE {
			request.Corporate = "1000000001"
		}
		request.Corporate = request.Corporate + "%"
		resdashboard, err := svc.Service.DashboardRepo.GetWeeklyRevenue(request, tx)
		if err != nil {
			log.Println("Error GetWeeklyRevenue : GetWeeklyRevenue", err)

			result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "ERROR GetWeeklyRevenue", constants.EMPTY_VALUE_INT)
			return err
		}
		elapsed := time.Since(timeNow).Milliseconds()
		fmt.Println("GetWeeklyRevenue - GetWeeklyRevenue elapsed time millisecond :", elapsed)

		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success GetWeeklyRevenue", resdashboard)
		return err
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result)
	}

	return ctx.JSON(http.StatusOK, result)
}

func (svc dashboardSummaryService) GetMonthlyRevenue(ctx echo.Context) error {
	var result models.Response
	var request models.RequestRevenue
	timeNow := time.Now()

	err := BindValidateStruct(ctx, &request)
	if err != nil {
		log.Println("Error Bind Data GetMonthlyRevenue : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Bind Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}
	fmt.Println("Bind request time start :", timeNow)

	if request.Corporate == constants.EMPTY_VALUE {
		request.Corporate = "1000000001"
	}
	request.Corporate = request.Corporate + "%"

	err = DBTransaction(svc.Service.RepoDB, func(tx *sql.Tx) error {

		resdashboard, err := svc.Service.DashboardRepo.GetMonthlyRevenue(request, tx)
		if err != nil {
			log.Println("Error GetMonthlyRevenue : GetMonthlyRevenue", err)

			result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "ERROR GetMonthlyRevenue", constants.EMPTY_VALUE_INT)
			return err
		}
		elapsed := time.Since(timeNow).Milliseconds()
		fmt.Println("GetMonthlyRevenue - GetMonthlyRevenue elapsed time millisecond :", elapsed)

		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success GetMonthlyRevenue", resdashboard)
		return err
	})

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result)
	}

	return ctx.JSON(http.StatusOK, result)
}

func (svc dashboardSummaryService) GetYearlyRevenue(ctx echo.Context) error {
	var result models.Response
	var request models.RequestRevenue
	timeNow := time.Now()

	err := BindValidateStruct(ctx, &request)
	if err != nil {
		log.Println("Error Bind Data GetYearlyRevenue : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Bind Data", err.Error())
		return ctx.JSON(http.StatusBadRequest, result)
	}
	fmt.Println("Bind request time start :", timeNow)

	err = DBTransaction(svc.Service.RepoDB, func(tx *sql.Tx) error {
		if request.Corporate == constants.EMPTY_VALUE {
			request.Corporate = "1000000001"
		}
		request.Corporate = request.Corporate + "%"
		resdashboard, err := svc.Service.DashboardRepo.GetYearlyRevenue(request, tx)
		if err != nil {
			log.Println("Error GetYearlyRevenue : GetYearlyRevenue", err)

			result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "ERROR GetYearlyRevenue", constants.EMPTY_VALUE_INT)
			return err
		}
		elapsed := time.Since(timeNow).Milliseconds()
		fmt.Println("GetYearlyRevenue` - GetYearlyRevenue` elapsed time millisecond :", elapsed)

		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success GetYearlyRevenue", resdashboard)
		return err
	})
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result)
	}

	return ctx.JSON(http.StatusOK, result)
}
func (svc dashboardSummaryService) SummaryDashboard(ctx echo.Context) error {
	var result models.Response
	var request models.RequestSummaryRevenue
	timeNow := time.Now()
	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data:  SummaryDashboard ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), constants.EMPTY_VALUE_INT)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	if request.Corporate == constants.EMPTY_VALUE {
		request.Corporate = "1000000001"
	}
	// fmt.Println(request)
	request.Corporate = request.Corporate + "%"
	err := DBTransaction(svc.Service.RepoDB, func(tx *sql.Tx) error {

		resdashboard, err := svc.Service.DashboardRepo.SummaryDashboard(request, tx)
		if err != nil {
			log.Println("Error Get SummaryDashboard : SummaryDashboard", err)

			result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "ERROR Get SummaryDashboard", constants.EMPTY_VALUE_INT)
			return err
		}
		elapsed := time.Since(timeNow).Milliseconds()
		fmt.Println("Get SummaryDashboard elapsed time:", elapsed)
		// fmt.Println(resdashboard)
		res := models.ResponseOmsetDashboard{
			Omset:      resdashboard.TotalBruto,
			ServiceFee: resdashboard.TotalServiceFee,
			TrxCount:   resdashboard.JumlahTrx,
		}
		// fmt.Println(res)
		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get SummaryDashboard", res)
		return err
	})

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result)

	}

	return ctx.JSON(http.StatusOK, result)
}

func (svc dashboardSummaryService) TrafficGrossTrxHourly(ctx echo.Context) error {
	var result models.Response
	var request models.RequestTrafficGrossTrx
	timeNow := time.Now()
	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data: TrafficGrossTrxHourly ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), constants.EMPTY_VALUE_INT)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	if request.Corporate == constants.EMPTY_VALUE {
		request.Corporate = "1000000001"
	}
	request.Corporate = request.Corporate + "%"
	resdashboard, err := svc.Service.DashboardRepo.TrafficGrossTrxHourly(request)
	if err != nil {
		log.Println("Error Get TrafficGrossTrxHourly : TrafficGrossTrxHourly", err)
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "ERROR Get TrafficGrossTrxHourly", constants.EMPTY_VALUE_INT)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	elapsed := time.Since(timeNow).Milliseconds()
	fmt.Println("TrafficGrossTrxHourly elapsed time: ", elapsed)
	log.Println("Reponse GetTrafficGrossTrxHourly ", "Success Get Data")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get TrafficGrossTrxHourly", resdashboard)
	return ctx.JSON(http.StatusOK, result)
}

func (svc dashboardSummaryService) GrossTrxWeeklyByPayMethod(ctx echo.Context) error {
	var result models.Response
	var request models.RequestRevenue
	timeNow := time.Now()
	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data: GrossTrxWeeklyByPayMethod ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), constants.EMPTY_VALUE_INT)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	if request.Corporate == constants.EMPTY_VALUE {
		request.Corporate = "1000000001"
	}
	request.Corporate = request.Corporate + "%"
	err := DBTransaction(svc.Service.RepoDB, func(tx *sql.Tx) error {

		resdashboard, err := svc.Service.DashboardRepo.GrossTrxWeeklyByPayMethod(request, tx)
		if err != nil {
			log.Println("Error Get GrossTrxWeeklyByPayMethod : GrossTrxWeeklyByPayMethod", err)

			result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "ERROR Get GrossTrxWeeklyByPayMethod", constants.EMPTY_VALUE_INT)
			return err
		}
		elapsed := time.Since(timeNow).Milliseconds()
		fmt.Println("GrossTrxWeeklyByPayMethod elapsed: ", elapsed)
		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get GrossTrxWeeklyByPayMethod", resdashboard)
		return err
	})

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result)

	}

	return ctx.JSON(http.StatusOK, result)
}

func (svc dashboardSummaryService) RevenueByCorporateId(ctx echo.Context) error {
	var result models.Response
	var request models.RequestSummaryRevenue
	timeNow := time.Now()
	if err := BindValidateStruct(ctx, &request); err != nil {

		log.Println("Error Validate Data: RevenueByCorporateId ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), constants.EMPTY_VALUE_INT)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	request.Corporate = request.Corporate + "%"
	err := DBTransaction(svc.Service.RepoDB, func(tx *sql.Tx) error {

		resdashboard, err := svc.Service.DashboardRepo.RevenueByCorporateId(request, tx)
		if err != nil {
			log.Println("Error Get RevenueByCorporateId : RevenueByCorporateId", err)

			result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "ERROR Get RevenueByCorporateId", constants.EMPTY_VALUE_INT)
			return err
		}
		elapsed := time.Since(timeNow).Milliseconds()
		fmt.Println("RevenueByCorporateId elapsed time: ", elapsed)
		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Get RevenueByCorporateId", resdashboard)
		return err
	})

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, result)

	}

	return ctx.JSON(http.StatusOK, result)
}
