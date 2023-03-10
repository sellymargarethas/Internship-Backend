package app

import (
	"database/sql"

	"Internship-Backend/repositories"
	"Internship-Backend/repositories/coresettlementRepository"
	summaryRepository "Internship-Backend/repositories/summaryRepository"

	corporateRepository "Internship-Backend/repositories/corporateRepository"
	"Internship-Backend/repositories/deviceRepository"

	paketrepository "Internship-Backend/repositories/paketRepository"
	pembayaranrepository "Internship-Backend/repositories/pembayaranRepository"
	produkrepository "Internship-Backend/repositories/produkRepository"

	dashboardRepository "Internship-Backend/repositories/dashboardRepository"
	transactionRepository "Internship-Backend/repositories/transactionRepository"
	usersrepository "Internship-Backend/repositories/usersRepository"
	"Internship-Backend/repositories/wilayahRepository"

	"Internship-Backend/services"
)

func SetupApp(DB *sql.DB, repo repositories.Repository) services.UsecaseService {
	corporateRepo := corporateRepository.NewCorporateRepository(repo)
	kategoripembayaranRepo := pembayaranrepository.NewKategoriPembayaranRepository(repo)
	metodepembayaranRepo := pembayaranrepository.NewMetodePembayaranRepository(repo)
	vendorpembayaranRepo := pembayaranrepository.NewVendorPembayaranRepository(repo)
	coresettlementRepo := coresettlementRepository.NewCoreSettlementRepository(repo)
	usersRepo := usersrepository.NewUsersRepository(repo)
	deviceRepo := deviceRepository.NewDeviceRepository(repo)
	blktgpembayaranRepo := usersrepository.NewBlktgsRepository(repo)
	provinsiRepository := wilayahRepository.NewProvinsiRepository(repo)
	produkRepository := produkrepository.NewProdukRepository(repo)
	loginRepository := usersrepository.NewloginRepository(repo)
	paketRepository := paketrepository.NewPaketRepository(repo)
	corporatePaymentRepository := corporateRepository.NewCorporatePaymentRepository(repo)
	corporateCategoryRepository := corporateRepository.NewCorporateCategoryRepository(repo)
	produkPaketRepository := paketrepository.NewProdukPaketRepository(repo)
	satuanProdukRepository := produkrepository.NewSatuanProdukRepository(repo)
	kategoriProdukRepository := produkrepository.NewKategoriProdukRepository(repo)
	kotaRepository := wilayahRepository.NewKotaRepository(repo)
	listTransactionRepository := transactionRepository.NewListTransactionRepository(repo)
	successSummaryRepository := summaryRepository.NewSuccessSummaryRepository(repo)
	dashboardRepository := dashboardRepository.NewDashboardRepository(repo)
	roleRepository := usersrepository.NewRoleRepository(repo)

	usecaseSvc := services.NewUsecaseService(
		DB, corporateRepo, kategoripembayaranRepo, coresettlementRepo, usersRepo, deviceRepo, blktgpembayaranRepo, provinsiRepository, produkRepository, satuanProdukRepository, kategoriProdukRepository, loginRepository, paketRepository, produkPaketRepository, corporatePaymentRepository, corporateCategoryRepository, metodepembayaranRepo, vendorpembayaranRepo, kotaRepository, successSummaryRepository, listTransactionRepository, dashboardRepository, roleRepository,
	)

	return usecaseSvc
}
