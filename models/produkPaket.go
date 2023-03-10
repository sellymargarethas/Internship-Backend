package models

type ProdukPaket struct {
	Id         int    `json:"id"`
	IdPaket    int    `json:"idPaket" validate:"required"`
	NamaPaket  string `json:"namaPaket"`
	IdProduk   int    `json:"idProduk" validate:"required"`
	NamaProduk string `json:"namaProduk"`
	HargaJual  int    `json:"hargaJual"`
}

type RequestAddProdukPaket struct {
	IdPaket    int    `json:"idPaket" validate:"required"`
	IdProduk   int    `json:"idProduk" validate:"required"`
	NamaProduk string `json:"namaProduk"`
	HargaJual  int    `json:"hargaJual"`
	CreatedAt  string `json:"createdAt"`
}

type RequestUpdateProdukPaket struct {
	Id         int    `json:"id"`
	IdPaket    int    `json:"idPaket" validate:"required"`
	IdProduk   int    `json:"idProduk" validate:"required"`
	NamaProduk string `json:"namaProduk"`
	HargaJual  int    `json:"hargaJual"`
	UpdatedAt  string `json:"updatedAt"`
}

type RequestDeleteProdukPaket struct {
	Id        int    `json:"id"`
	DeletedAt string `json:"deletedAt"`
}

type ProdukPaketArray struct {
	ProdukPaketArray []ProdukPaket `json:"produkPaketArray"`
}
