package models

type Kota struct {
	Id           int    `json:"id"`
	Uraian       string `json:"uraian"`
	IdProvinsi   int    `json:"idProvinsi"`
	NamaProvinsi string `json:"namaProvinsi"`
}

type RequestAddKota struct {
	Uraian     string `json:"uraian" validate:"required"`
	IdProvinsi int    `json:"idProvinsi" validate:"required"`
	CreatedAt  string `json:"createdAt"`
}

type RequestUpdateKota struct {
	Id         int    `json:"id" validate:"required"`
	Uraian     string `json:"uraian" validate:"required"`
	IdProvinsi int    `json:"idProvinsi" validate:"required"`
	DeletedAt  string `json:"deletedAt"`
}

type RequestDeleteKota struct {
	Id        int    `json:"id" validate:"required"`
	DeletedAt string `json:"deletedAt"`
}

type KotaArray struct {
	KotaArray []Kota `json:"kota"`
}
