package models

type Paket struct {
	Id            int    `json:"id"`
	Uraian        string `json:"uraian" validate:"required"`
	IdCorporate   int    `json:"idCorporate" validate:"required"`
	NamaCorporate string `json:"namaCorporate"`
}

type RequestAddPaket struct {
	Uraian      string `json:"uraian" validate:"required"`
	IdCorporate int    `json:"idCorporate" validate:"required"`
	CreatedAt   string `json:"createdAt"`
}

type RequestUpdatePaket struct {
	Id          int    `json:"id"`
	Uraian      string `json:"uraian" validate:"required"`
	IdCorporate int    `json:"idCorporate" validate:"required"`
	UpdatedAt   string `json:"updatedAt"`
}

type RequestDeletePaket struct {
	Id        int    `json:"id" validate:"required"`
	DeletedAt string `json:"deletedAt"`
}

type PaketArray struct {
	PaketArray []Paket `json:"paket"`
}
