package models

type ViewCorporatePayment struct {
	IdCorporate         int    `json:"idCorporate"`
	CID                 int    `json:"CID"`
	Uraian              string `json:"uraian"`
	Kota                string `json:"kota"`
	JmlMetodePembayaran int    `json:"jmlMetodePemabayaran"`
}
type MetodePembayaranDetails struct {
	IdMetodePembayaran int    `json:"idMetodePembayaran"`
	Uraian             string `json:"uraian"`
	Samnum             int    `json:"samnum"`
}

type ViewCorporatePaymentDetails struct {
	IdCorporate      int                       `json:"idCorporate"`
	CID              int                       `json:"CID"`
	Uraian           string                    `json:"uraian"`
	Kota             string                    `json:"kota"`
	MetodePembayaran []MetodePembayaranDetails `json:"metodePembayaran"`
}

type RequestAddCorporatePayment struct {
	IdCorporate             int                       `json:"idCorporate" validate:"required"`
	MetodePembayaranDetails []DetailsCorporatePayment `json:"metodePembayaran" validate:"required,dive"`
}

type RequestAddCorporatePaymentArray struct {
	RequestAddCorporatePayment []RequestAddCorporatePayment `json:"requestAddCorporatePayment" validate:"required,dive,required"`
}

type RequestGetCorporatePaymentByID struct {
	IdCorporate int `json:"idCorporate" validate:"required"`
}

type RequestDeleteCorporatePayment struct {
	IdCorporate             int                      `json:"idCorporate" validate:"required"`
	IdMetodePembayaranArray []CorporatePaymentDelete `json:"idMetodePembayaranArray" validate:"required,dive"`
}

type CorporatePaymentDelete struct {
	IdMetodePembayaran int `json:"idMetodePembayaran" validate:"required"`
}

type ResponseCorporatePaymentDetails struct {
	IdCorporate int `json:"idCorporate"`
}

type DetailsCorporatePayment struct {
	IdCorporatePayment int  `json:"idMetodePembayaran" validate:"required"`
	Samnum             *int `json:"samnum"`
}

type ResponseGetOne struct {
	IdCorporate             int                       `json:"idCorporate"`
	MetodePembayaranDetails []MetodePembayaranDetails `json:"metodePembayaran"`
}

type RequestUpdateCorporatePayment struct {
	IdCorporate             int                       `json:"idCorporate" validate:"required"`
	MetodePembayaranDetails []DetailsCorporatePayment `json:"metodePembayaran" validate:"required,dive"`
}
