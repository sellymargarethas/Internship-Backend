package models

type SatuanProduk struct {
	Id     int    `json:"id"`
	Uraian string `json:"uraian"`
}

type RequestAddSatuanProduk struct {
	Id        int    `json:"id"`
	Uraian    string `json:"uraian" validate:"required"`
	CreatedAt string `json:"createdAt"`
}

type RequestUpdateSatuanProduk struct {
	Id        int    `json:"id" validate:"required"`
	Uraian    string `json:"uraian" validate:"required"`
	UpdatedAt string `json:"updatedAt"`
}

type RequestDeleteSatuanProduk struct {
	Id        int    `json:"id" validate:"required"`
	DeletedAt string `json:"deletedAt"`
}

type SatuanProdukArray struct {
	SatuanProdukArray []SatuanProduk `json:"satuan_produk"`
}
