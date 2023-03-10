package services

import (
	"database/sql"

	"Internship-Backend/repositories"
)

type UsecaseService struct {
	RepoDB                 *sql.DB
	CorporateRepo          repositories.CorporateRepository
	KategoriPembayaranRepo repositories.KategoriPembayaranRepository
	CoreSettlementRepo     repositories.CoreSettlementRepository
	UserRepo               repositories.UsersRepository
	DeviceRepo             repositories.DeviceRepository
	BlktgRepo              repositories.BlktgRepository
	ProvinsiRepo           repositories.ProvinsiRepository
	ProdukRepo             repositories.ProdukRepository
	SatuanProdukRepo       repositories.SatuanProdukRepository
	KategoriProdukRepo     repositories.KategoriProdukRepository
	LoginRepo              repositories.LoginRepository
	PaketRepo              repositories.PaketRepository
	ProdukPaketRepo        repositories.ProdukPaketRepository
	CorporatePaymentRepo   repositories.CorporatePaymentRepository
	CorporateCategoryRepo  repositories.CorporateCategoryRepository
	MetodePembayaranRepo   repositories.MetodePembayaranRepository
	VendorPembayaranRepo   repositories.VendorPembayaranRepository
	KotaRepo               repositories.KotaRepository

	SuccessSummaryRepo repositories.SuccessSummaryRepository

	ListTransactionRepo repositories.ListTransactionRepository
	DashboardRepo       repositories.DashboardRepository
	RoleRepo            repositories.RoleRepository
}

func NewUsecaseService(
	repoDB *sql.DB,
	CorporateRepo repositories.CorporateRepository,
	KategoriPembayaranRepo repositories.KategoriPembayaranRepository,
	CoreSettlementRepo repositories.CoreSettlementRepository,
	UserRepo repositories.UsersRepository,
	DeviceRepo repositories.DeviceRepository,
	BlktgRepo repositories.BlktgRepository,
	ProvinsiRepo repositories.ProvinsiRepository,
	ProdukRepo repositories.ProdukRepository,
	SatuanProdukRepo repositories.SatuanProdukRepository,
	KategoriProdukRepo repositories.KategoriProdukRepository,
	LoginRepo repositories.LoginRepository,
	PaketRepo repositories.PaketRepository,
	ProdukPaketRepo repositories.ProdukPaketRepository,
	CorporatePaymentRepo repositories.CorporatePaymentRepository,
	CorporateCategoryRepo repositories.CorporateCategoryRepository,
	MetodePembayaranRepo repositories.MetodePembayaranRepository,
	VendorPembayaranRepo repositories.VendorPembayaranRepository,
	KotaRepo repositories.KotaRepository,

	SuccessSummaryRepo repositories.SuccessSummaryRepository,

	ListTransactionRepo repositories.ListTransactionRepository,
	DashboardRepo repositories.DashboardRepository,
	RoleRepo repositories.RoleRepository,

) UsecaseService {
	return UsecaseService{
		RepoDB:                 repoDB,
		CorporateRepo:          CorporateRepo,
		KategoriPembayaranRepo: KategoriPembayaranRepo,
		CoreSettlementRepo:     CoreSettlementRepo,
		UserRepo:               UserRepo,
		DeviceRepo:             DeviceRepo,
		BlktgRepo:              BlktgRepo,
		ProvinsiRepo:           ProvinsiRepo,
		ProdukRepo:             ProdukRepo,
		SatuanProdukRepo:       SatuanProdukRepo,
		KategoriProdukRepo:     KategoriProdukRepo,
		LoginRepo:              LoginRepo,
		PaketRepo:              PaketRepo,
		ProdukPaketRepo:        ProdukPaketRepo,
		CorporatePaymentRepo:   CorporatePaymentRepo,
		CorporateCategoryRepo:  CorporateCategoryRepo,
		MetodePembayaranRepo:   MetodePembayaranRepo,
		VendorPembayaranRepo:   VendorPembayaranRepo,
		KotaRepo:               KotaRepo,

		SuccessSummaryRepo: SuccessSummaryRepo,

		ListTransactionRepo: ListTransactionRepo,
		DashboardRepo:       DashboardRepo,
		RoleRepo:            RoleRepo,
	}
}
