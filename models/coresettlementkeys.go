package models

import "time"

type CoreSettlementKeys struct {
	Id             int       `json:"id"`
	Value          string    `json:"value"`
	Name           string    `json:"name"`
	Created_at     time.Time `json:"createdAt"`
	Updated_at     time.Time `json:"updatedAt"`
	WorkerUsername *string   `json:"workerUsername"`
}

type ResponseGetCoreSettlement struct {
	Id             int     `json:"id"`
	Value          string  `json:"value"`
	Name           string  `json:"name"`
	WorkerUsername *string `json:"workerUsername"`
}

type RequestAddCoreSettlement struct {
	Value          string `json:"value" validate:"required"`
	Nama           string `json:"nama" validate:"required"`
	WorkerUsername string `json:"workerUsername"`
}

type RequestUpdateCoreSettlement struct {
	Id             int    `json:"id" validate:"required"`
	Value          string `json:"value" validate:"required"`
	Nama           string `json:"nama"`
	WorkerUsername string `json:"workerUsername"`
}

type RequestDeleteCoreSettlement struct {
	Id int `json:"id" validate:"required"`
}

type ResponseCoreSettlement struct {
	Id int `json:"id"`
}
