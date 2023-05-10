package models

type ListBarangResponse struct {
	ResponseCode          string   `json:"responseCode"`
	ResponseMessage       string   `json:"responseMessage"`
	ResponseMessageDetail string   `json:"responseMessageDetail"`
	ListData              []Barang `json:ListData`
}

type ErrorResponse struct {
	ResponseCode          string `json:"responseCode"`
	ResponseMessage       string `json:"responseMessage"`
	ResponseMessageDetail string `json:"responseMessageDetail"`
}

type SignatureResponse struct {
	ResponseCode          string      `json:"responseCode"`
	ResponseMessage       string      `json:"responseMessage"`
	ResponseMessageDetail string      `json:"responseMessageDetail"`
	Data                  interface{} `json:"Data"`
}

type RegisterResponse struct {
	ResponseCode          string      `json:"responseCode"`
	ResponseMessage       string      `json:"responseMessage"`
	ResponseMessageDetail string      `json:"responseMessageDetail"`
	DataAccount           interface{} `json:DataAccount`
}
