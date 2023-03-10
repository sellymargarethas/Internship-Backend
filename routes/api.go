package routes

import (
	"Internship-Backend/config"
	"Internship-Backend/constants"
	"Internship-Backend/services"
	"net/http"

	"github.com/labstack/echo/middleware"

	"Internship-Backend/services/coresettlementService"
	corporateservice "Internship-Backend/services/corporateService"

	dashboardService "Internship-Backend/services/dashboardService"
	"Internship-Backend/services/deviceService"
	"Internship-Backend/services/paketService"
	"Internship-Backend/services/pembayaranService"
	"Internship-Backend/services/produkService"
	summaryservice "Internship-Backend/services/summaryService"
	"Internship-Backend/services/transactionService"
	usersservice "Internship-Backend/services/usersService"
	"Internship-Backend/services/wilayahService"

	"github.com/labstack/echo"
)

// Routes API
func RoutesApi(e echo.Echo, usecaseSvc services.UsecaseService) {
	public := e.Group("")

	private := e.Group("")
	private.Use(middleware.JWT([]byte(config.GetEnv("JWT_KEY"))))
	private.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: constants.TRUE_VALUE,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	corporateGroup := private.Group("/corporate")
	corporateSvc := corporateservice.NewCorporateService(usecaseSvc)
	corporateGroup.POST("/list", corporateSvc.GetListCorporate)
	corporateGroup.POST("/get", corporateSvc.GetSingleCorporate)
	corporateGroup.POST("/add", corporateSvc.InsertCorporate)
	corporateGroup.POST("/edit", corporateSvc.UpdateCorporate)
	corporateGroup.POST("/remove", corporateSvc.DeleteCorporate)

	corporateCategoryGroup := private.Group("/corpcategory")
	corporateCategorySvc := corporateservice.NewCorporateCategoryService(usecaseSvc)
	corporateCategoryGroup.GET("/list", corporateCategorySvc.GetCorporateCategory)
	corporateCategoryGroup.POST("/add", corporateCategorySvc.InsertCorporateCategory)
	corporateCategoryGroup.POST("/edit", corporateCategorySvc.UpdateCorporateCategory)
	corporateCategoryGroup.POST("/remove", corporateCategorySvc.DeleteCorporateCategory)

	corporatePaymentGroup := private.Group("/corppayment")
	corporatePaymentSvc := corporateservice.NewCorporatePaymentService(usecaseSvc)
	corporatePaymentGroup.POST("/list", corporatePaymentSvc.GetAllCorporatePayment)
	corporatePaymentGroup.POST("/get", corporatePaymentSvc.GetCorporatePaymentByIdCorporate)
	corporatePaymentGroup.POST("/add", corporatePaymentSvc.InsertCorporatePayment)
	corporatePaymentGroup.POST("/edit", corporatePaymentSvc.UpdateCorporatePayment)

	coresettlementGroup := private.Group("/core")
	coresettlementSvc := coresettlementService.NewCoreSettlementService(usecaseSvc)
	coresettlementGroup.POST("/list", coresettlementSvc.GetAllCoreSettlementKeys)
	coresettlementGroup.POST("/add", coresettlementSvc.InsertCoreSettlementKey)
	coresettlementGroup.POST("/remove", coresettlementSvc.DeleteCoreSettlementKey)
	coresettlementGroup.POST("/edit", coresettlementSvc.UpdateCoreSettlementKey)
	coresettlementGroup.POST("/get", coresettlementSvc.GetSingleCoreSettlementKeys)

	userGroup := private.Group("/user")
	userSvc := usersservice.NewUsersService(usecaseSvc)
	userGroup.POST("/add", userSvc.InsertUser)
	userGroup.POST("/list", userSvc.GetAllUser)
	userGroup.POST("/edit", userSvc.EditUser)
	userGroup.POST("/remove", userSvc.DeleteUser)
	userGroup.POST("/get", userSvc.GetUser)
	userGroup.GET("/jenis", userSvc.GetJenisUser)
	userGroup.POST("/updatepassword", userSvc.EditUserPassword)

	deviceGroup := private.Group("/device")
	deviceSvc := deviceService.NewDeviceService(usecaseSvc)
	deviceGroup.POST("/list", deviceSvc.GetListDevice)
	deviceGroup.POST("/add", deviceSvc.InsertDevice)
	deviceGroup.POST("/edit", deviceSvc.UpdateDevice)
	deviceGroup.POST("/remove", deviceSvc.DeleteDevice)
	deviceGroup.POST("/get", deviceSvc.GetSingleDevice)

	kategoripembayaranGroup := private.Group("/kategoripembayaran")
	kategoripembayaranSvc := pembayaranService.NewKategoriPembayaranService(usecaseSvc)
	kategoripembayaranGroup.GET("/list", kategoripembayaranSvc.GetListKategoriPembayaran)
	kategoripembayaranGroup.POST("/add", kategoripembayaranSvc.InsertKategoriPembayaran)
	kategoripembayaranGroup.POST("/edit", kategoripembayaranSvc.UpdateKategoriPembayaran)
	kategoripembayaranGroup.POST("/remove", kategoripembayaranSvc.DeleteKategoriPembayaran)

	metodepembayaranGroup := private.Group("/metodepembayaran")
	metodepembayaranSvc := pembayaranService.NewMetodePembayaranService(usecaseSvc)
	metodepembayaranGroup.GET("/list", metodepembayaranSvc.GetListMetodePembayaran)
	metodepembayaranGroup.POST("/add", metodepembayaranSvc.InsertMetodePembayaran)
	metodepembayaranGroup.POST("/edit", metodepembayaranSvc.UpdateMetodePembayaran)
	metodepembayaranGroup.POST("/remove", metodepembayaranSvc.DeleteMetodePembayaran)

	vendorpembayaranGroup := private.Group("/vendorpembayaran")
	vendorpembayaranSvc := pembayaranService.NewVendorPembayaranService(usecaseSvc)
	vendorpembayaranGroup.GET("/list", vendorpembayaranSvc.GetListVendorPembayaran)
	vendorpembayaranGroup.POST("/add", vendorpembayaranSvc.InsertVendorPembayaran)
	vendorpembayaranGroup.POST("/edit", vendorpembayaranSvc.UpdateVendorPembayaran)
	vendorpembayaranGroup.POST("/remove", vendorpembayaranSvc.DeleteVendorPembayaran)

	blktgpembayaranGroup := private.Group("/bl")
	blktgpembayaranSvc := usersservice.NewblktgService(usecaseSvc)
	// blktgpembayaranGroup.POST("/add", blktgpembayaranSvc.InsertBlktgPembayaran)
	blktgpembayaranGroup.POST("/list", blktgpembayaranSvc.GetBlktgPembayaranByIdUser)
	blktgpembayaranGroup.POST("/edit", blktgpembayaranSvc.UpdateBlktgPembayaran)

	provinsiGroup := private.Group("/provinsi")
	provinsiSvc := wilayahService.NewProvinsiService(usecaseSvc)
	provinsiGroup.GET("/list", provinsiSvc.GetListProvinsi)
	provinsiGroup.POST("/add", provinsiSvc.InsertProvinsi)
	provinsiGroup.POST("/edit", provinsiSvc.UpdateProvinsi)
	provinsiGroup.POST("/remove", provinsiSvc.DeleteProvinsi)

	kotaGroup := private.Group("/kota")
	kotaSvc := wilayahService.NewKotaService(usecaseSvc)
	kotaGroup.GET("/list", kotaSvc.GetListKota)
	kotaGroup.POST("/add", kotaSvc.InsertKota)
	kotaGroup.POST("/edit", kotaSvc.UpdateKota)
	kotaGroup.POST("/remove", kotaSvc.DeleteKota)

	satuanprodukGroup := private.Group("/satuanproduk")
	satuanprodukSvc := produkService.NewSatuanProdukService(usecaseSvc)
	satuanprodukGroup.GET("/list", satuanprodukSvc.GetListSatuanProduk)
	satuanprodukGroup.POST("/add", satuanprodukSvc.InsertSatuanProduk)
	satuanprodukGroup.POST("/edit", satuanprodukSvc.UpdateSatuanProduk)
	satuanprodukGroup.POST("/remove", satuanprodukSvc.DeleteSatuanProduk)

	kategoriprodukGroup := private.Group("/kategoriproduk")
	kategoriprodukSvc := produkService.NewKategoriProdukService(usecaseSvc)
	kategoriprodukGroup.GET("/list", kategoriprodukSvc.GetListKategoriProduk)
	kategoriprodukGroup.POST("/add", kategoriprodukSvc.InsertKategoriProduk)
	kategoriprodukGroup.POST("/edit", kategoriprodukSvc.UpdateKategoriProduk)
	kategoriprodukGroup.POST("/remove", kategoriprodukSvc.DeleteKategoriProduk)

	produkGroup := private.Group("/produk")
	produkSvc := produkService.NewProdukService(usecaseSvc)
	produkGroup.POST("/list", produkSvc.GetListProduk)
	produkGroup.POST("/add", produkSvc.InsertProduk)
	produkGroup.POST("/edit", produkSvc.UpdateProduk)
	produkGroup.POST("/remove", produkSvc.DeleteProduk)
	produkGroup.POST("/get", produkSvc.GetSingleProduk)

	loginGroup := public.Group("/login")
	loginSvc := usersservice.NewloginService(usecaseSvc)
	loginGroup.POST("/login", loginSvc.Login)

	paketGroup := private.Group("/paket")
	paketSvc := paketService.NewPaketService(usecaseSvc)
	paketGroup.GET("/list", paketSvc.GetListPaket)
	paketGroup.POST("/add", paketSvc.InsertPaket)
	paketGroup.POST("/edit", paketSvc.UpdatePaket)
	paketGroup.POST("/remove", paketSvc.DeletePaket)

	produkpaketGroup := private.Group("/produkpaket")
	produkpaketSvc := paketService.NewProdukPaketService(usecaseSvc)
	produkpaketGroup.GET("/list", produkpaketSvc.GetListProdukPaket)
	produkpaketGroup.POST("/add", produkpaketSvc.InsertProdukPaket)
	produkpaketGroup.POST("/edit", produkpaketSvc.UpdateProdukPaket)
	produkpaketGroup.POST("/remove", produkpaketSvc.DeleteProdukPaket)

	listtransactionGroup := private.Group("/listtrx")
	listtransactionSvc := transactionService.NewListTransactionService(usecaseSvc)
	listtransactionGroup.POST("/list", listtransactionSvc.GetListTransaction)
	listtransactionGroup.POST("/listsettled", listtransactionSvc.GetListTransactionSettled)
	listtransactionGroup.POST("/get", listtransactionSvc.GetSingleTransaction)
	listtransactionGroup.POST("/getproduk", listtransactionSvc.GetSingleDetailProduk)
	listtransactionGroup.POST("/count", listtransactionSvc.CountTransaction)
	listtransactionGroup.POST("/countsettled", listtransactionSvc.CountTransactionSettled)

	successsummaryGroup := private.Group("/summary")
	successsumarySvc := summaryservice.NewSuccessSummaryService(usecaseSvc)
	successsummaryGroup.POST("/success", successsumarySvc.GetSuccessSummary)
	successsummaryGroup.POST("/settled", successsumarySvc.GetSettledSummary)

	dashboardGroup := private.Group("/dashboard")
	dashboardSvc := dashboardService.NewDashboardSummaryService(usecaseSvc)
	dashboardGroup.POST("/dailyrevenue", dashboardSvc.GetDailyRevenue)
	dashboardGroup.POST("/weeklyrevenue", dashboardSvc.GetWeeklyRevenue)
	dashboardGroup.POST("/monthlyrevenue", dashboardSvc.GetMonthlyRevenue)
	dashboardGroup.POST("/yearlyrevenue", dashboardSvc.GetYearlyRevenue)
	dashboardGroup.POST("/header", dashboardSvc.SummaryDashboard)
	dashboardGroup.POST("/hourly", dashboardSvc.TrafficGrossTrxHourly)
	dashboardGroup.POST("/paymethod", dashboardSvc.GrossTrxWeeklyByPayMethod)
	dashboardGroup.POST("/revenue", dashboardSvc.RevenueByCorporateId)

	roleGroup := private.Group("/role")
	roleSvc := usersservice.NewRoleService(usecaseSvc)
	roleGroup.POST("/add", roleSvc.AddUserRole)
	roleGroup.POST("/list", roleSvc.GetListRole)
	roleGroup.POST("/get", roleSvc.GetSingleRole)
	roleGroup.POST("/edit", roleSvc.UpdateUserRole)
	roleGroup.POST("/delete", roleSvc.DeleteUserRole)
	roleGroup.GET("/menu", roleSvc.GetAllRoleMenu)
	roleGroup.GET("/task", roleSvc.GetAllRoleTask)
}
