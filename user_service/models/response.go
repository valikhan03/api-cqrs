package models

type ErrorResponse struct{
	StatusCode int `json:"status_code"`
	Error string `json:"error"`
}
