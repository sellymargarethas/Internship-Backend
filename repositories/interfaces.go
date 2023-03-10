package repositories

import (
	"Internship-Backend/models"
	"database/sql"
)

type CorporateRepository interface {
	IsCorporateExistsByIndex(corporate models.Corporate) (models.Corporate, bool, error)
	GetListCorporate(request models.RequestList) ([]models.Corporate, error)
	GetSingleCorporate(corporate models.Corporate) ([]models.Corporate, error)
	InsertCorporate(corporate models.RequestAddCorporate) (int, error)
	UpdateCorporate(corporate models.RequestUpdateCorporate) (id int, err error)
	DeleteCorporate(corporate models.RequestDeleteCorporate, tx *sql.Tx) (id int, cid string, err error)
	DeleteParentCorporate(cid string, tx *sql.Tx) (cid2 string, err error)
}
type CorporatePaymentRepository interface {
	InsertCorporatePayment(corporatePayment models.RequestAddCorporatePayment) (result int, err error)
	GetListCorporatePayment(request models.RequestList) ([]models.ViewCorporatePayment, error)
	GetCorporatePaymentByIdCorporate(id int) ([]models.MetodePembayaranDetails, error)

	IsCorporatePaymentExistsByIndex(idCorporate int) (int, bool, error)
	UpdateCorporatePayment(corporatePayment models.RequestUpdateCorporatePayment, tx *sql.Tx) (status bool, err error)
	DeleteCorporatePaymentDetails(id int, tx *sql.Tx) (result bool, err error)
}

type CorporateCategoryRepository interface {
	GetCorporateCategoryList() (list []models.CorporateCategory, err error)
	InsertCorporateCategory(data models.RequestAddCorporateCategory) (id int, err error)
	EditCorporateCategory(data models.RequestUpdateCorporateCategory) (id int, err error)
	DeleteCorporateCategory(data models.DeleteCorporateCategory) (id int, err error)
	IsCorporateCategoryExistbyIndex(category models.CorporateCategory) (status bool, err error)
}

type KategoriPembayaranRepository interface {
	GetListKategoriPembayaran() ([]models.KategoriPembayaran, error)
	InsertKategoriPembayaran(kategoripembayaran models.RequestAddKategoriPembayaran) (int, error)
	UpdateKategoriPembayaran(pembayaran models.RequestUpdateKategoriPembayaran) (id int, err error)
	DeleteKategoriPembayaran(pembayaran models.RequestDeleteKategoriPembayaran) (id int, err error)
	IsKategoriPembayaranExistsByIndex(kategoripembayaran models.KategoriPembayaran) (models.KategoriPembayaran, bool, error)
	GetListKategoriPembayaranId() (result []models.KategoriPembayaranId, err error)
}

type MetodePembayaranRepository interface {
	GetListMetodePembayaran() ([]models.MetodePembayaran, error)
	InsertMetodePembayaran(metodepembayaran models.RequestAddMetodePembayaran) (id int, err error)
	UpdateMetodePembayaran(pembayaran models.RequestUpdateMetodePembayaran) (id int, err error)
	DeleteMetodePembayaran(pembayaran models.RequestDeleteMetodePembayaran) (id int, err error)
	IsMetodePembayaranExistsByIndex(metodepembayaran models.MetodePembayaran) (models.MetodePembayaran, bool, error)
}

type VendorPembayaranRepository interface {
	GetListVendorPembayaran() ([]models.VendorPembayaran, error)
	InsertVendorPembayaran(vendorpembayaran models.RequestAddVendorPembayaran) (id int, err error)
	UpdateVendorPembayaran(pembayaran models.RequestUpdateVendorPembayaran) (id int, err error)
	DeleteVendorPembayaran(pembayaran models.RequestDeleteVendorPembayaran) (id int, err error)
	IsVendorPembayaranExistsByIndex(vendorpembayaran models.VendorPembayaran) (models.VendorPembayaran, bool, error)
}

type CoreSettlementRepository interface {
	GetListCoreSettlementKeys(request models.RequestList) (coresettlementkeys []models.ResponseGetCoreSettlement, err error)
	GetSingleCoreSettlementKeys(id int) (list models.ResponseGetCoreSettlement, err error)
	InsertCoreSettlementKey(coresettlement models.RequestAddCoreSettlement) (id int, err error)
	UpdateCoreSettlementKey(coresettlement models.RequestUpdateCoreSettlement) (id int, err error)
	DeleteCoreSettlementKey(idcoresettlement models.RequestDeleteCoreSettlement) (bool, error)
	IsCoresettlementkeysExistByIndex(csk models.CoreSettlementKeys) (bool, error)
	GetListCSKId() (result []models.SettlementDestination, err error)
	GetStrAggCSKName(request models.RequestSummarySettled) (result string, err error)
}

type UsersRepository interface {
	InsertUser(user models.RequestAddUser) (id int, err error)
	GetAllUser(request models.RequestList) (users []models.ResponseUser, err error)
	GetUserByIndex(user models.Users) (response []models.ResponseUser, err error)
	UpdateUser(user models.RequestUpdateUser) (id int, err error)
	DeleteUser(user models.RequestDeleteUser) (id int, err error)
	GetJenisUser() (jenis []models.ResponseJenisUser, err error)
	UpdateUserPassword(request models.RequestUpdatePassword) (id int, err error)
	UpdateEnc(id int, enc string) (id2 int, err error)
}

type DeviceRepository interface {
	GetListDevice(request models.RequestList) (devices []models.ResponseDevice, err error)
	InsertDevice(device models.RequestAddDevice) (id int, err error)
	EditDevice(device models.RequestUpdateDevice) (id int, err error)
	DeleteDevice(device models.RequestDeleteDevice) (id int, err error)
	GetSingleDevice(device models.Device) (data models.ResponseSingleDevice, err error)
}

type BlktgRepository interface {
	GetBlktgPembayaran(id models.RequestGetBlktgPembayaranByIdUser) (response []models.ResponseBlktgPembayaran, err error)
	InsertBlktgPembayaran(request models.RequestUpdateBlktgPembayaran) (status bool, err error)

	DeleteBlktgPembayaran(id int, tx *sql.Tx) (result bool, err error)
	IsBlkgPembayaranExistsByIndex(blktgPembayaran models.BlktgPembayaran) (models.BlktgPembayaran, bool, error)

	// UpdateBlktgPembayaran(request models.RequestUpdateBlktgPembayaran, tx *sql.Tx) (status bool, err error)
	UpdateBlktgPembayaran(request models.RequestUpdateBlktgPembayaran, tx *sql.Tx) (status bool, err error)
}
type ProvinsiRepository interface {
	GetListProvinsi() ([]models.Provinsi, error)
	InsertProvinsi(provinsi models.RequestAddProvinsi) (id int, err error)
	UpdateProvinsi(pembayaran models.RequestUpdateProvinsi) (id int, err error)
	DeleteProvinsi(pembayaran models.RequestDeleteProvinsi) (id int, err error)
	IsProvinsiExistsByIndex(provinsi models.Provinsi) (models.Provinsi, bool, error)
}

type KotaRepository interface {
	GetListKota() ([]models.Kota, error)
	InsertKota(kota models.RequestAddKota) (id int, err error)
	UpdateKota(kota models.RequestUpdateKota) (id int, err error)
	DeleteKota(pembayaran models.RequestDeleteKota) (id int, err error)
	IsKotaExistsByIndex(kota models.Kota) (models.Kota, bool, error)
}

type ProdukRepository interface {
	GetListProduk(request models.RequestList) ([]models.Produk, error)
	GetSingleProduk(idproduk int) (data models.Produk, err error)
	InsertProduk(produk models.RequestAddProduk) (id int, err error)
	UpdateProduk(produk models.RequestUpdateProduk) (id int, err error)
	DeleteProduk(produk models.RequestDeleteProduk) (id int, err error)
	IsProdukExistsByIndex(produk models.Produk) (models.Produk, bool, error)
}

type SatuanProdukRepository interface {
	GetListSatuanProduk() ([]models.SatuanProduk, error)
	InsertSatuanProduk(satuanproduk models.RequestAddSatuanProduk) (id int, err error)
	UpdateSatuanProduk(satuanproduk models.RequestUpdateSatuanProduk) (id int, err error)
	DeleteSatuanProduk(satuanproduk models.RequestDeleteSatuanProduk) (id int, err error)
	IsSatuanProdukExistsByIndex(satuanproduk models.SatuanProduk) (models.SatuanProduk, bool, error)
}

type KategoriProdukRepository interface {
	GetListKategoriProduk() ([]models.KategoriProduk, error)
	InsertKategoriProduk(kategoriproduk models.RequestAddKategoriProduk) (id int, err error)
	UpdateKategoriProduk(kategoriproduk models.RequestUpdateKategoriProduk) (id int, err error)
	DeleteKategoriProduk(kategoriproduk models.RequestDeleteKategoriProduk) (id int, err error)
	IsKategoriProdukExistsByIndex(kategoriproduk models.KategoriProduk) (models.KategoriProduk, bool, error)
}

type LoginRepository interface {
	CheckLogin(data models.Login) (hash string, jenis int, err error)
	LoginReturn(data models.Login) (user models.LoginResponse, err error)
}

type PaketRepository interface {
	GetListPaket() ([]models.Paket, error)
	InsertPaket(paket models.RequestAddPaket) (id int, err error)
	UpdatePaket(paket models.RequestUpdatePaket) (id int, err error)
	DeletePaket(paket models.RequestDeletePaket) (id int, err error)
	IsPaketExistsByIndex(paket models.Paket) (models.Paket, bool, error)
}

type ProdukPaketRepository interface {
	GetListProdukPaket() ([]models.ProdukPaket, error)
	InsertProdukPaket(produkpaket models.RequestAddProdukPaket) (id int, err error)
	UpdateProdukPaket(produkpaket models.RequestUpdateProdukPaket) (id int, err error)
	DeleteProdukPaket(produkpaket models.RequestDeleteProdukPaket) (id int, err error)
	IsProdukPaketExistsByIndex(produkpaket models.ProdukPaket) (models.ProdukPaket, bool, error)
	GetCountList() (count int, err error)
}

type SuccessSummaryRepository interface {
	// ViewSuccessSummary(request models.RequestSummarySuccess, tx *sql.Tx) (response models.ScanSummary, err error)
	ViewSuccessSummary(request models.RequestSummarySuccess, tx *sql.Tx) (temp string, err error)
	ScanResultSuccess(temp string, tx *sql.Tx) (response models.ScanSummary, err error)
	// ViewSuccessSummary(request models.RequestSummarySuccess, tx *sql.Tx) (response models.ResponseSummarySuccessTrx, err error)
	// ViewPaymentMethodSuccessSummary(request models.RequestSummarySuccess, tx *sql.Tx) (response models.AllSummary, err error)
	// ViewSettledSummary(request models.RequestSummarySettled, tx *sql.Tx) (response models.AllSummary, err error)
	// ViewSettledSummary(request models.RequestSummarySettled, tx *sql.Tx) (response models.ResponseSummarySettledTrx, err error)
	// ViewSettledSummary(request models.RequestSummarySettled, tx *sql.Tx) (temp string, err error)
	ViewSettledSummary(request models.RequestSummarySettled, tx *sql.Tx) (response models.ScanSettlement, err error)
	// ScanResultSettled(temp string, tx *sql.Tx) (response models.ScanSettlement, err error)
}

type ListTransactionRepository interface {
	IsListTrxExistsByIndex(listtrx models.RequestTrx) (models.ResponseListTrx, bool, error)
	GetListTransaction(request models.RequestListTrx, page models.PageOffsetLimit) ([]models.ResponseListTrx, error)
	GetListTransactionSettled(request models.RequestListTrx, page models.PageOffsetLimit) ([]models.ResponseListTrx, error)
	GetSingleTransaction(idtrx int) (onetrx models.ResponseOneTrx, err error)
	GetSingleProdukTransaction(idtrx int) (detailproduktrx models.ResponseDetailProductTrx, err error)
	CountListTransaction(request models.RequestListTrx) (count int, err error)
	CountListTransactionSettled(request models.RequestListTrx) (count int, err error)
}

type DashboardRepository interface {
	GetDailyRevenue(request models.RequestRevenue, tx *sql.Tx) (data []models.ResponseDailyRevenue, err error)
	GetWeeklyRevenue(request models.RequestRevenue, tx *sql.Tx) (data []models.ResponseWeeklyRevenue, err error)
	GetMonthlyRevenue(request models.RequestRevenue, tx *sql.Tx) (data []models.ResponseMonthlyRevenue, err error)
	GetYearlyRevenue(request models.RequestRevenue, tx *sql.Tx) (data []models.ResponseYearlyRevenue, err error)
	SummaryDashboard(request models.RequestSummaryRevenue, tx *sql.Tx) (response models.ScanSummary, err error)
	TrafficGrossTrxHourly(request models.RequestTrafficGrossTrx) (response []models.ResponseSummaryOmset, err error)
	GrossTrxWeeklyByPayMethod(request models.RequestRevenue, tx *sql.Tx) (data []models.ResponseDashboardPaymentCategory, err error)
	RevenueByCorporateId(request models.RequestSummaryRevenue, tx *sql.Tx) (data []models.ResponseSummaryOmset, err error)
}

type RoleRepository interface {
	AddUserRole(request models.RequestAddUsersRole) (err error)
	GetListRole(request models.RequestListRole) (response []models.ResponseListRole, err error)
	GetSingleRoleDetails(request models.RequestSingleRole) (response []models.RoleDetails, err error)
	GetSingleRole(request models.RequestSingleRole) (response models.ResponseRoleDetails, err error)
	UpdateRole(request models.RequestUpdateRole) (status bool, err error)
	DeleteRole(request models.RequestSingleRole) (status bool, err error)
	GetRoleIdByIdUser(id int) (role int, err error)
	GetAllRoleMenu() (response []models.RoleMenu, err error)
	GetAllRoleTask() (response []models.RoleTask, err error)
}
