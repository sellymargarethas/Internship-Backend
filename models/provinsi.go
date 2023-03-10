package models

type Provinsi struct {
	Id     int    `json:"id"`
	Uraian string `json:"uraian" validate:"required"`
}

type RequestAddProvinsi struct {
	Uraian    string `json:"uraian" validate:"required"`
	CreatedAt string `json:"createdAt"`
}

type RequestUpdateProvinsi struct {
	Id        int    `json:"id" validate:"required"`
	Uraian    string `json:"uraian" validate:"required"`
	UpdatedAt string `json:"updatedAt"`
}

type RequestDeleteProvinsi struct {
	Id        int    `json:"id" validate:"required"`
	DeletedAt string `json:"deletedAt"`
}

type ProvinsiArray struct {
	ProvinsiArray []Provinsi `json:"provinsi"`
}
