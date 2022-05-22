package models

import(

)

type SuccessResponse struct{
	StatusCode int `json:"status_code"`
	Message string `json:"message"`	
}

type ErrorResponse struct{
	StatusCode int `json:"status_code"`
	Error string `json:"error"`
}

