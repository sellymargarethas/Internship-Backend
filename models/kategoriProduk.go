package models

type KategoriProduk struct {
	Id         int    `json:"id"`
	Uraian     string `json:"uraian" validate:"required"`
	ServiceFee int    `json:"serviceFee" validate:"required"`
}

type RequestAddKategoriProduk struct {
	Uraian     string `json:"uraian" validate:"required"`
	ServiceFee int    `json:"serviceFee" validate:"required"`
	CreatedAt  string `json:"createdAt"`
}

type RequestUpdateKategoriProduk struct {
	Id         int    `json:"id" validate:"required"`
	Uraian     string `json:"uraian" validate:"required"`
	ServiceFee int    `json:"serviceFee" validate:"required"`
	UpdatedAt  string `json:"updatedAt"`
}

type RequestDeleteKategoriProduk struct {
	Id        int    `json:"id" validate:"required"`
	DeletedAt string `json:"deletedAt"`
}

type KategoriProdukArray struct {
	KategoriProdukArray []KategoriProduk `json:"kategori_produk_array"`
}
