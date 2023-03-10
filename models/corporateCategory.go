package models

type CorporateCategory struct {
	IdCategory int    `json:"idCategory"`
	Kode       string `json:"kode"`
	Uraian     string `json:"uraian"`
}

type RequestUpdateCorporateCategory struct {
	IdCategory int    `json:"idCategory" validate:"required"`
	Kode       string `json:"kode" validate:"required"`
	Uraian     string `json:"uraian" validate:"required"`
}

type RequestAddCorporateCategory struct {
	Kode   string `json:"kode" validate:"required"`
	Uraian string `json:"uraian" validate:"required"`
}

type DeleteCorporateCategory struct {
	IdCategory int `json:"idCategory" validate:"required"`
}
