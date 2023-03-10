package models

import "time"

type Device struct {
	Id          int       `json:"id"`
	DeviceId    string    `json:"deviceId"`
	Tid         string    `json:"tid"`
	IdCorporate int       `json:"idCorporate"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
	Mid         string    `json:"mid"`
	AquiringMid string    `json:"aquiringMid"`
	AquiringTid string    `json:"aquiringTid"`
	Merchantkey string    `json:"merchantKey"`
	DkiMid      string    `json:"dkiMid"`
	DkiTid      string    `json:"dkiTid"`
	Tokenfcm    string    `json:"tokenFcm"`
	JenisDevice string    `json:"jenisDevice"`
}

type RequestAddDevice struct {
	DeviceId               string  `json:"deviceId" validate:"required"`
	IdCorporate            int     `json:"idCorporate" validate:"required"`
	Tid                    *string `json:"tid"`
	Mid                    *string `json:"mid"`
	DkiTid                 string  `json:"dkiTid"`
	DkiMid                 string  `json:"dkiMid"`
	AquiringMid            *string `json:"aquiringMid"`
	AquiringTid            *string `json:"aquiringTid"`
	IdCorporateMerchantkey int     `json:"idCorporateMerchantKey" validate:"required"` //id dari idcorporate yg ada di coresettlementkey
	Tokenfcm               *string `json:"tokenFcm"`
	JenisDevice            string  `json:"jenisDevice"`
}
type RequestUpdateDevice struct {
	Id                     int     `json:"id"`
	DeviceId               string  `json:"deviceId" validate:"required"`
	IdCorporate            int     `json:"idCorporate" validate:"required"`
	Tid                    *string `json:"tid"`
	Mid                    *string `json:"mid"`
	DkiTid                 string  `json:"dkiTid"`
	DkiMid                 string  `json:"dkiMid"`
	AquiringMid            *string `json:"aquiringMid"`
	AquiringTid            *string `json:"aquiringTid"`
	IdCorporateMerchantkey int     `json:"idCorporateMerchantKey" validate:"required"` //id dari idcorporate yg ada di coresettlementkey
	Tokenfcm               *string `json:"tokenFcm"`
	JenisDevice            string  `json:"jenisDevice"`
}

type RequestDeleteDevice struct {
	Id int `json:"id" validate:"required"`
}

type RequestGetDevice struct {
	Id int `json:"id" validate:"required"`
}

type ResponseDevice struct {
	Id            int     `json:"id"`
	DeviceId      string  `json:"deviceId"`
	IdCorporate   int     `json:"idCorporate"`
	NamaCorporate string  `json:"namaCorporate"`
	Mid           *string `json:"mid"`
	Tid           *string `json:"tid"`
	DkiTid        *string `json:"dkiTid"`
	DkiMid        *string `json:"dkiMid"`
	AquiringMid   *string `json:"aquiringMid"`
	AquiringTid   *string `json:"aquiringTid"`
	Tokenfcm      *string `json:"tokenFcm"`
	MerchantKey   string  `json:"merchantKey"` //ambil id dari coresettlementkeys, select value nya
	CSKName       string  `json:"CSKName"`     //nama corporate berdasarkan coresettlementkeys -nya
	JenisDevice   string  `json:"jenisDevice"` //name nya coresettlementkeys
}

type ResponseSingleDevice struct {
	Id            int     `json:"id"`
	DeviceId      string  `json:"deviceId"`
	IdCorporate   int     `json:"idCorporate"`
	NamaCorporate string  `json:"namaCorporate"`
	Mid           *string `json:"mid"`
	Tid           *string `json:"tid"`
	DkiTid        *string `json:"dkiTid"`
	DkiMid        *string `json:"dkiMid"`
	AquiringMid   *string `json:"aquiringMid"`
	AquiringTid   *string `json:"aquiringTid"`
	Tokenfcm      *string `json:"tokenFcm"`
	MerchantKey   int     `json:"merchantKey"` //ambil id dari coresettlementkeys, select value nya
	CSKName       string  `json:"CSKName"`     //nama corporate berdasarkan coresettlementkeys -nya
	JenisDevice   string  `json:"jenisDevice"` //name nya coresettlementkeys
}
