package models

type GetListReq struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
