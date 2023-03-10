package models

type RequestListTrx struct {
	StartDate             string                  `json:"startDate" validate:"required"`
	EndDate               string                  `json:"endDate"`
	StatusTrx             int                     `json:"statusTrx"`
	KategoriPembayaran    string                  `json:"kategoriPembayaran"`
	Corporate             string                  `json:"corporate"`
	SattlementDestination []SattlementDestination `json:"sattlementDestination"`
	Pagination            RequestList             `json:"pagination"`
}

type ResponseListTrx struct {
	ID                    int      `json:"id"`
	NomorHeader           string   `json:"nomorHeader"`
	MerchantNoRef         string   `json:"merchantNoRef"`
	MetodePembayaran      string   `json:"metodePembayaran"` //kategori pembayaran
	Acquiring             string   `json:"acquiring"`        //metode pembayaran
	SettlementDestination string   `json:"settlementDestination"`
	StatusTrx             int      `json:"statusTrx"`
	StatusTrxDesc         string   `json:"statusTrxDesc"`
	StatusSettlement      *bool    `json:"statusSettlement"`
	ResNoRef              *string  `json:"resNoRef"`
	CID                   string   `json:"cid"`
	Corporate             string   `json:"corporate"`
	DeviceId              string   `json:"deviceId"`
	MID                   *string  `json:"mid"`
	TID                   *string  `json:"tid"`
	CardPan               *string  `json:"cardPan"`
	CardType              *string  `json:"cardType"`
	HargaJual             float64  `json:"hargaJual"`
	Potongan              int      `json:"potongan"`
	KodePromo             *string  `json:"kodePromo"`
	MDR                   *float64 `json:"mdr"`
	ServiceFee            float64  `json:"serviceFee"`
	PaymentFee            *int     `json:"paymentFee"`
	VendorFee             *int     `json:"vendorFee"`
	CreatedAt             string   `json:"createdAt"`
}

type ResponseOneTrx struct {
	ID               int      `json:"id"`
	NomorHeader      string   `json:"nomorHeader"`
	MerchantNoRef    string   `json:"merchantNoRef"`
	TerimaTunai      float64  `json:"terimaTunai"`
	StatusAwalTrx    string   `json:"statusAwalTrx"`
	StatusTrx        int      `json:"statusTrx"`
	StatusTrxDesc    string   `json:"statusTrxDesc"`
	MetodePembayaran string   `json:"metodePembayaran"`
	ResNoRef         *string  `json:"resNoRef"`
	CID              string   `json:"cid"`
	Corporate        string   `json:"corporate"`
	DeviceId         string   `json:"deviceId"`
	MID              *string  `json:"mid"`
	TID              *string  `json:"tid"`
	CardPan          *string  `json:"cardPan"`
	CardType         *string  `json:"cardType"`
	HargaJual        float64  `json:"hargaJual"`
	PaymentMDR       float64  `json:"paymentMDR"`
	PaymentDisc      float64  `json:"paymentDisc"`
	ServiceFee       float64  `json:"serviceFee"`
	Potongan         *float64 `json:"potongan"`
	CreatedAt        string   `json:"createdAt"`
}

type RequestTrx struct {
	IdTrx int `json:"idTrx" validate:"required"`
}

type ResponseDetailProductTrx struct {
	KodeProduk string  `json:"kodeProduk"`
	NamaProduk string  `json:"namaProduk"`
	Kategori   string  `json:"kategori"`
	Paket      *string `json:"paket"`
	HargaJual  float64 `json:"hargaJual"`
	Diskon     float64 `json:"diskon"`
	Qty        int     `json:"qty"`
	Satuan     string  `json:"satuan"`
	ServiceFee float64 `json:"serviceFee"`
}

type SattlementDestination struct {
	SattlementDestination string `json:"sattlementDestination"`
}
