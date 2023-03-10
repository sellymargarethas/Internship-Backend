package models

type Corporate struct {
	Id                    int     `json:"id"`
	CID                   string  `json:"cid"`
	Uraian                string  `json:"uraian"`
	IdKota                *int    `json:"idKota"`
	Alamat                string  `json:"alamat"`
	Telepon               string  `json:"telepon"`
	ParentCID             string  `json:"parentCID"`
	Level                 int     `json:"level"`
	Gambar                *string `json:"gambar"`
	HirarkiId             string  `json:"hirarkiId"`
	IpLocalServer         *string `json:"ipLocalServer"`
	ServiceFee            int     `json:"serviceFee"`
	IsPercentage          bool    `json:"isPercentage"`
	IdCorporateCategory   *int    `json:"corporateCategory"`
	NamaCorporateCategory string  `json:"namaCorporateCategory"`
	NamaKota              string  `json:"namaKota"`
	NamaProvinsi          string  `json:"namaProvinsi"`
}

type RequestAddCorporate struct {
	Uraian              string  `json:"uraian" validate:"required"`
	CID                 string  `json:"cid"`
	Alamat              string  `json:"alamat"`
	Telepon             string  `json:"telepon"`
	ParentCID           string  `json:"parentCID" validate:"required"`
	HirarkiId           string  `json:"hirarkiId"`
	Level               int     `json:"level"`
	Gambar              *string `json:"gambar"`
	IsPercentage        bool    `json:"isPercentage"`
	IdCorporateCategory int     `json:"corporateCategory" validate:"required"`
	IpLocalServer       *string `json:"ipLocalServer"`
	IdKota              int     `json:"idKota" validate:"required"`
	ServiceFee          int     `json:"serviceFee" validate:"required"`
	CreatedAt           string  `json:"createdAt"`
}

type RequestUpdateCorporate struct {
	Id                  int     `json:"id" validate:"required"`
	Uraian              string  `json:"uraian" validate:"required"`
	Alamat              string  `json:"alamat"`
	Telepon             string  `json:"telepon"`
	IpLocalServer       *string `json:"ipLocalServer"`
	ServiceFee          int     `json:"serviceFee" validate:"required"`
	IdKota              int     `json:"idKota" validate:"required"`
	Gambar              *string `json:"gambar"`
	IsPercentage        bool    `json:"isPercentage"`
	IdCorporateCategory int     `json:"corporateCategory" validate:"required"`
	UpdatedAt           string  `json:"updatedAt"`
}

type RequestDeleteCorporate struct {
	Id        int    `json:"id" validate:"required"`
	DeletedAt string `json:"deletedAt"`
}

type CorporateArray struct {
	CorporateArray []Corporate `json:"corporate"`
}
