package models

type GetBookListResponses struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Amount int32  `json:"amount"`
}

type CreateBulkBook struct {
	Book string
	File string
}
