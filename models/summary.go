package models

import "time"

type ResponseSummarySuccessTrx struct {
	StartDate       string  `json:"startDate"`
	EndDate         string  `json:"endDate"`
	Corporate       string  `json:"corporate"`
	Pembayaran      string  `json:"pembayaran"`
	Status          string  `json:"status"`
	JumlahTrx       int     `json:"jumlahTrx"`
	TotalTunai      float64 `json:"totalTunai"`
	TotalLain       float64 `json:"totalLain"`
	TotalCC         float64 `json:"totalCC"`
	TotalDebit      float64 `json:"totalDebit"`
	TotalQRIS       float64 `json:"totalQRIS"`
	TotalBrizzi     float64 `json:"totalBrizzi"`
	TotalEMoney     float64 `json:"totalEMoney"`
	TotalTapCash    float64 `json:"totalTapCash"`
	TotalFlazz      float64 `json:"totalFlazz"`
	TotalJakcard    float64 `json:"totalJakcard"`
	TotalVA         float64 `json:"totalVA"`
	TotalBiller     float64 `json:"totalBiller"`
	TotalBruto      float64 `json:"totalBruto"`
	TotalMDR        float64 `json:"totalMDR"`
	TotalPotongan   float64 `json:"totalPotongan"`
	TotalServiceFee float64 `json:"totalServiceFee"`
	TotalPaymentFee float64 `json:"totalPaymentFee"`
	TotalVendorFee  float64 `json:"totalVendorFee"`
	TotalNett       float64 `json:"totalNett"`
}
type RequestSuccess struct {
	StartDate          string `json:"startDate" validate:"required"`
	EndDate            string `json:"endDate" validate:"required"`
	KategoriPembayaran string `json:"kategoriPembayaran"`
	Corporate          string `json:"corporate"`
}

type RequestSummarySuccess struct {
	StartDate string `json:"startDate" validate:"required"`
	EndDate   string `json:"endDate" validate:"required"`
	Corporate string `json:"corporate"`
}

type ResponseSummarySettledTrx struct {
	StartDate             string  `json:"startDate"`
	EndDate               string  `json:"endDate"`
	Corporate             string  `json:"corporate"`
	Pembayaran            string  `json:"pembayaran"`
	Status                string  `json:"status"`
	SettlementDestination string  `json:"settlementDestination"`
	JumlahTrx             int     `json:"jumlahTrx"`
	TotalBruto            float64 `json:"totalBruto"`
	TotalMDR              float64 `json:"totalMDR"`
	TotalServiceFee       float64 `json:"totalServiceFee"`
	TotalPaymentFee       float64 `json:"totalPaymentFee"`
	TotalVendorFee        float64 `json:"totalVendorFee"`
	TotalTunai            float64 `json:"totalTunai"`
	TotalLain             float64 `json:"totalLain"`
	TotalCC               float64 `json:"totalCC"`
	TotalDebit            float64 `json:"totalDebit"`
	TotalQRIS             float64 `json:"totalQRIS"`
	TotalBrizzi           float64 `json:"totalBrizzi"`
	TotalEMoney           float64 `json:"totalEMoney"`
	TotalTapCash          float64 `json:"totalTapCash"`
	TotalFlazz            float64 `json:"totalFlazz"`
	TotalJakcard          float64 `json:"totalJakcard"`
	TotalVA               float64 `json:"totalVA"`
	TotalBiller           float64 `json:"totalBiller"`
	TotalPotongan         float64 `json:"totalPotongan"`
	TotalNett             float64 `json:"totalNett"`
}

type RequestSummarySettled struct {
	StartDate             string                  `json:"startDate" validate:"required"`
	EndDate               string                  `json:"endDate" validate:"required"`
	KategoriPembayaran    []KategoriPembayaranId  `json:"kategoriPembayaran"`
	Corporate             string                  `json:"corporate"`
	SettlementDestination []SettlementDestination `json:"settlementDestination"`
}

type RequestSettled struct {
	StartDate             string                  `json:"startDate" validate:"required"`
	EndDate               string                  `json:"endDate" validate:"required"`
	KategoriPembayaran    string                  `json:"kategoriPembayaran"`
	Corporate             string                  `json:"corporate"`
	SettlementDestination []SettlementDestination `json:"settlementDestination"`
}

type SettlementDestination struct {
	IdCoreSettlement string `json:"idCoreSettlement"`
}

type ScanSettlement struct {
	StartDate             string  `json:"startDate"`
	EndDate               string  `json:"endDate"`
	Corporate             string  `json:"corporate"`
	Pembayaran            string  `json:"pembayaran"`
	Status                string  `json:"status"`
	SettlementDestination string  `json:"settlementDestination"`
	JumlahTrx             int     `json:"jumlahTrx"`
	TotalTunai            float64 `json:"totalTunai"`
	TotalLain             float64 `json:"totalLain"`
	TotalCC               float64 `json:"totalCC"`
	TotalDebit            float64 `json:"totalDebit"`
	TotalQRIS             float64 `json:"totalQRIS"`
	TotalBrizzi           float64 `json:"totalBrizzi"`
	TotalEMoney           float64 `json:"totalEMoney"`
	TotalTapCash          float64 `json:"totalTapCash"`
	TotalFlazz            float64 `json:"totalFlazz"`
	TotalJakcard          float64 `json:"totalJakcard"`
	TotalVA               float64 `json:"totalVA"`
	TotalBiller           float64 `json:"totalBiller"`
	TotalBruto            float64 `json:"totalBruto"`
	TotalMDR              float64 `json:"totalMDR"`
	TotalPotongan         float64 `json:"totalPotongan"`
	TotalServiceFee       float64 `json:"totalServiceFee"`
	TotalPaymentFee       float64 `json:"totalPaymentFee"`
	TotalVendorFee        float64 `json:"totalVendorFee"`
	TotalNett             float64 `json:"totalNett"`
	CountTunai            int     `json:"countTunai"`
	CountLainnya          int     `json:"countLainnya"`
	CountCC               int     `json:"countCC"`
	CountDebit            int     `json:"countDebit"`
	CountQRIS             int     `json:"countQRIS"`
	CountBrizzi           int     `json:"countBrizzi"`
	CountEMoney           int     `json:"countEMoney"`
	CountTapCash          int     `json:"countTapCash"`
	CountFlazz            int     `json:"countFlazz"`
	CountJakcard          int     `json:"countJakcard"`
	CountVA               int     `json:"countVa"`
	CountBiller           int     `json:"countBiller"`
	SFeeTunai             float64 `json:"sFeeTunai"`
	SFeeLainnya           float64 `json:"sFeeLainnya"`
	SFeeCC                float64 `json:"sFeeCC"`
	SFeeDebit             float64 `json:"sFeeDebit"`
	SFeeQRIS              float64 `json:"sFeeQRIS"`
	SFeeBrizzi            float64 `json:"sFeeBrizzi"`
	SFeeEMoney            float64 `json:"sFeeEMoney"`
	SFeeTapCash           float64 `json:"sFeeTapCash"`
	SFeeFlazz             float64 `json:"sFeeFlazz"`
	SFeeJakcard           float64 `json:"sFeeJakcard"`
	SFeeVA                float64 `json:"sFeeVA"`
	SFeeBiller            float64 `json:"sFeeBiller"`
	MdrTunai              float64 `json:"mdrTunai"`
	MdrLainnya            float64 `json:"mdrLainnya"`
	MdrCC                 float64 `json:"mdrCC"`
	MdrDebit              float64 `json:"mdrDebit"`
	MdrQRIS               float64 `json:"mdrQRIS"`
	MdrBrizzi             float64 `json:"mdrBrizzi"`
	MdrEMoney             float64 `json:"mdrEMoney"`
	MdrTapCash            float64 `json:"mdrTapCash"`
	MdrFlazz              float64 `json:"mdrFlazz"`
	MdrJakcard            float64 `json:"mdrJakcard"`
	MdrVA                 float64 `json:"mdrVA"`
	MdrBiller             float64 `json:"mdrBiller"`
	PayfeeTunai           float64 `json:"payfeeTunai"`
	PayfeeLainnya         float64 `json:"payfeeLainnya"`
	PayfeeCC              float64 `json:"payfeeCC"`
	PayfeeDebit           float64 `json:"payfeeDebit"`
	PayfeeQRIS            float64 `json:"payfeeQRIS"`
	PayfeeBrizzi          float64 `json:"payfeeBrizzi"`
	PayfeeEMoney          float64 `json:"payfeeEMoney"`
	PayfeeTapCash         float64 `json:"payfeeTapCash"`
	PayfeeFlazz           float64 `json:"payfeeFlazz"`
	PayfeeJakcard         float64 `json:"payfeeJakcard"`
	PayfeeVA              float64 `json:"payfeeVA"`
	PayfeeBiller          float64 `json:"payfeeBiller"`
	VendorfeeTunai        float64 `json:"vendorfeeTunai"`
	VendorfeeLainnya      float64 `json:"vendorfeeLainnya"`
	VendorfeeCC           float64 `json:"vendorfeeCC"`
	VendorfeeDebit        float64 `json:"vendorfeeDebit"`
	VendorfeeQRIS         float64 `json:"vendorfeeQRIS"`
	VendorfeeBrizzi       float64 `json:"vendorfeeBrizzi"`
	VendorfeeEMoney       float64 `json:"vendorfeeEMoney"`
	VendorfeeTapCash      float64 `json:"vendorfeeTapCash"`
	VendorfeeFlazz        float64 `json:"vendorfeeFlazz"`
	VendorfeeJakcard      float64 `json:"vendorfeeJakcard"`
	VendorfeeVA           float64 `json:"vendorfeeVA"`
	VendorfeeBiller       float64 `json:"vendorfeeBiller"`
}

type ScanSummary struct {
	DateCreatedAt    time.Time `json:"dateCreatedAt"`
	Date             int       `json:"date"`
	DateString       string    `json:"dateString"`
	JumlahTrx        int       `json:"jumlahTrx"`
	TotalTunai       float64   `json:"totalTunai"`
	TotalLain        float64   `json:"totalLain"`
	TotalCC          float64   `json:"totalCC"`
	TotalDebit       float64   `json:"totalDebit"`
	TotalQRIS        float64   `json:"totalQRIS"`
	TotalBrizzi      float64   `json:"totalBrizzi"`
	TotalEMoney      float64   `json:"totalEMoney"`
	TotalTapCash     float64   `json:"totalTapCash"`
	TotalFlazz       float64   `json:"totalFlazz"`
	TotalJakcard     float64   `json:"totalJakcard"`
	TotalVA          float64   `json:"totalVA"`
	TotalBiller      float64   `json:"totalBiller"`
	TotalBruto       float64   `json:"totalBruto"`
	TotalMDR         float64   `json:"totalMDR"`
	TotalPotongan    float64   `json:"totalPotongan"`
	TotalServiceFee  float64   `json:"totalServiceFee"`
	TotalPaymentFee  float64   `json:"totalPaymentFee"`
	TotalVendorFee   float64   `json:"totalVendorFee"`
	TotalNett        float64   `json:"totalNett"`
	CountTunai       int       `json:"countTunai"`
	CountLainnya     int       `json:"countLainnya"`
	CountCC          int       `json:"countCC"`
	CountDebit       int       `json:"countDebit"`
	CountQRIS        int       `json:"countQRIS"`
	CountBrizzi      int       `json:"countBrizzi"`
	CountEMoney      int       `json:"countEMoney"`
	CountTapCash     int       `json:"countTapCash"`
	CountFlazz       int       `json:"countFlazz"`
	CountJakcard     int       `json:"countJakcard"`
	CountVA          int       `json:"countVa"`
	CountBiller      int       `json:"countBiller"`
	SFeeTunai        float64   `json:"sFeeTunai"`
	SFeeLainnya      float64   `json:"sFeeLainnya"`
	SFeeCC           float64   `json:"sFeeCC"`
	SFeeDebit        float64   `json:"sFeeDebit"`
	SFeeQRIS         float64   `json:"sFeeQRIS"`
	SFeeBrizzi       float64   `json:"sFeeBrizzi"`
	SFeeEMoney       float64   `json:"sFeeEMoney"`
	SFeeTapCash      float64   `json:"sFeeTapCash"`
	SFeeFlazz        float64   `json:"sFeeFlazz"`
	SFeeJakcard      float64   `json:"sFeeJakcard"`
	SFeeVA           float64   `json:"sFeeVA"`
	SFeeBiller       float64   `json:"sFeeBiller"`
	MdrTunai         float64   `json:"mdrTunai"`
	MdrLainnya       float64   `json:"mdrLainnya"`
	MdrCC            float64   `json:"mdrCC"`
	MdrDebit         float64   `json:"mdrDebit"`
	MdrQRIS          float64   `json:"mdrQRIS"`
	MdrBrizzi        float64   `json:"mdrBrizzi"`
	MdrEMoney        float64   `json:"mdrEMoney"`
	MdrTapCash       float64   `json:"mdrTapCash"`
	MdrFlazz         float64   `json:"mdrFlazz"`
	MdrJakcard       float64   `json:"mdrJakcard"`
	MdrVA            float64   `json:"mdrVA"`
	MdrBiller        float64   `json:"mdrBiller"`
	PayfeeTunai      float64   `json:"payfeeTunai"`
	PayfeeLainnya    float64   `json:"payfeeLainnya"`
	PayfeeCC         float64   `json:"payfeeCC"`
	PayfeeDebit      float64   `json:"payfeeDebit"`
	PayfeeQRIS       float64   `json:"payfeeQRIS"`
	PayfeeBrizzi     float64   `json:"payfeeBrizzi"`
	PayfeeEMoney     float64   `json:"payfeeEMoney"`
	PayfeeTapCash    float64   `json:"payfeeTapCash"`
	PayfeeFlazz      float64   `json:"payfeeFlazz"`
	PayfeeJakcard    float64   `json:"payfeeJakcard"`
	PayfeeVA         float64   `json:"payfeeVA"`
	PayfeeBiller     float64   `json:"payfeeBiller"`
	VendorfeeTunai   float64   `json:"vendorfeeTunai"`
	VendorfeeLainnya float64   `json:"vendorfeeLainnya"`
	VendorfeeCC      float64   `json:"vendorfeeCC"`
	VendorfeeDebit   float64   `json:"vendorfeeDebit"`
	VendorfeeQRIS    float64   `json:"vendorfeeQRIS"`
	VendorfeeBrizzi  float64   `json:"vendorfeeBrizzi"`
	VendorfeeEMoney  float64   `json:"vendorfeeEMoney"`
	VendorfeeTapCash float64   `json:"vendorfeeTapCash"`
	VendorfeeFlazz   float64   `json:"vendorfeeFlazz"`
	VendorfeeJakcard float64   `json:"vendorfeeJakcard"`
	VendorfeeVA      float64   `json:"vendorfeeVA"`
	VendorfeeBiller  float64   `json:"vendorfeeBiller"`
}
