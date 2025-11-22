package dto

type ErrorObj struct {
	Code  		ErrorCode		`json:"code"`
	Message		string			`json:"message"`
}

type ErrorResponse struct {
	Error 		ErrorObj		`json:"error"`
}