package models

import "time"

type BlktgPembayaran struct {
	Id                   int       `json:"id"`
	IdUser               int       `json:"idUser"`
	IdKategoriPembayaran int       `json:"idKategoriPembayaran"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}
type RequestAddBlktgPembayaranArray struct {
	IdUser                    int                 `json:"idUser" validate:"required"`
	RequestAddBlktgPembayaran []IdBlktgPembayaran `json:"requestAddBlktgPembayaran" validate:"dive,required"`
}

type RequestUpdateBlktgPembayaran struct {
	IdUser               int                 `json:"idUser" validate:"required"`
	IdKategoriPembayaran []IdBlktgPembayaran `json:"idKategoriPembayaran" validate:"required,dive"`
}

type RequestDeleteBlktgPembayaran struct {
	IdUser int `json:"idUser" validate:"required"`
}

type IdBlktgPembayaran struct {
	IdKategoriPembayaran int `json:"idKategoriPembayaran" validate:"required"`
}

type RequestGetBlktgPembayaranByIdUser struct {
	IdUser int `json:"idUser" validate:"required"`
}

type ResponseListBlktgPembayaran struct {
	IdUser          int                       `json:"idUser"`
	BlktgPembayaran []ResponseBlktgPembayaran `json:"blktgPembayaran"`
}

type ResponseBlktgPembayaran struct {
	Idkategoripembayaran int    `json:"idKategoriPembayaran"`
	KategoriPembayaran   string `json:"kategoriPembayaran"`
	Status               bool   `json:"status"`
}
