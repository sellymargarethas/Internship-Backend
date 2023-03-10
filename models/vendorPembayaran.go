package models

type VendorPembayaran struct {
	Id     int    `json:"id"`
	Uraian string `json:"uraian"`
}

type RequestAddVendorPembayaran struct {
	Uraian    string `json:"uraian" validate:"required"`
	CreatedAt string `json:"createdAt"`
}

type RequestUpdateVendorPembayaran struct {
	Id        int    `json:"id" validate:"required"`
	Uraian    string `json:"uraian" validate:"required"`
	UpdatedAt string `json:"updatedAt"`
}

type RequestDeleteVendorPembayaran struct {
	Id        int    `json:"id" validate:"required"`
	DeletedAt string `json:"deletedAt"`
}

type VendorPembayaranArray struct {
	VendorPembayaranArray []VendorPembayaran `json:"vendorPembayaran"`
}
