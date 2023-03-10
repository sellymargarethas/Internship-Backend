package models

type MetodePembayaran struct {
	Id                     int    `json:"id"`
	Uraian                 string `json:"uraian"`
	IdKategoriPembayaran   int    `json:"idKategoriPembayaran"`
	NamaKategoriPembayaran string `json:"namaKategoriPembayaran"`
}

type RequestAddMetodePembayaran struct {
	Uraian               string `json:"uraian" validate:"required"`
	IdKategoriPembayaran int    `json:"idKategoriPembayaran" validate:"required"`
	CreatedAt            string `json:"createdAt"`
}

type RequestUpdateMetodePembayaran struct {
	Id                   int    `json:"id" validate:"required"`
	Uraian               string `json:"uraian" validate:"required"`
	IdKategoriPembayaran int    `json:"idKategoriPembayaran" validate:"required"`
	UpdatedAt            string `json:"updatedAt"`
}

type RequestDeleteMetodePembayaran struct {
	Id        int    `json:"id" validate:"required"`
	DeletedAt string `json:"deletedAt"`
}

type MetodePembayaranArray struct {
	MetodePembayaranArray []MetodePembayaran `json:"metodePembayaran"`
}
