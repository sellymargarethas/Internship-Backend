package models

type KategoriPembayaran struct {
	Id     int    `json:"id"`
	Uraian string `json:"uraian"`
}

type RequestAddKategoriPembayaran struct {
	Uraian    string `json:"uraian" validate:"required"`
	CreatedAt string `json:"createdAt"`
}

type RequestUpdateKategoriPembayaran struct {
	Id        int    `json:"id" validate:"required"`
	Uraian    string `json:"uraian" validate:"required"`
	UpdatedAt string `json:"updatedAt"`
}

type RequestDeleteKategoriPembayaran struct {
	Id        int    `json:"id" validate:"required"`
	DeletedAt string `json:"deletedAt"`
}

type KategoriPembayaranArray struct {
	KategoriPembayaranArray []KategoriPembayaran `json:"kategoriPembayaran"`
}
type KategoriPembayaranId struct {
	IdKategoriPembayaran int `json:"kategoriPembayaran"`
}
