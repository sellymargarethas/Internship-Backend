package models

type RequestSummaryRevenue struct {
	StartDate string `json:"startDate" validate:"required"`
	EndDate   string `json:"endDate" validate:"required"`
	Corporate string `json:"corporate"`
}
type RequestTrafficGrossTrx struct {
	Date      string `json:"date" validate:"required"`
	Corporate string `json:"corporate"`
}

type ResponseOmsetDashboard struct {
	Omset      float64 `json:"omset"`
	ServiceFee float64 `json:"serviceFee"`
	TrxCount   int     `json:"trxCount"`
}

type ResponseSummaryOmset struct {
	Date     int     `json:"date"`
	Omset    float64 `json:"omset"`
	TrxCount int     `json:"trxCount"`
}

type ResponseDashboardPaymentCategory struct {
	Date    string  `json:"date"`
	Day     string  `json:"day"`
	Tunai   float64 `json:"tunai"`
	Lainnya float64 `json:"lain"`
	CC      float64 `json:"cc"`
	Debit   float64 `json:"debit"`
	QRIS    float64 `json:"qris"`
	Brizzi  float64 `json:"brizzi"`
	EMoney  float64 `json:"eMoney"`
	TapCash float64 `json:"tapCash"`
	Flazz   float64 `json:"flazz"`
	Jakcard float64 `json:"jakcard"`
	VA      float64 `json:"va"`
	Biller  float64 `json:"biller"`
}

type RequestRevenue struct {
	Corporate string `json:"corporate"`
}

type ResponseDailyRevenue struct {
	Day              string  `json:"day"`
	TransactionCount int     `json:"transaction_count"`
	SumTransactions  float64 `json:"sum_transactions"`
}

type ResponseWeeklyRevenue struct {
	Week             string  `json:"week"`
	TransactionCount int     `json:"transaction_count"`
	SumTransactions  float64 `json:"sum_transactions"`
}

type ResponseMonthlyRevenue struct {
	Month            string  `json:"month"`
	TransactionCount int     `json:"transaction_count"`
	SumTransactions  float64 `json:"sum_transactions"`
}

type ResponseYearlyRevenue struct {
	Year             string  `json:"year"`
	TransactionCount int     `json:"transaction_count"`
	SumTransactions  float64 `json:"sum_transactions"`
}
