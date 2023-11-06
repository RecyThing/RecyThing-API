package helper

type ErrorResponseJson struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SuccessResponseJson struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ErrorResponse(message string) ErrorResponseJson {
	return ErrorResponseJson{
		Status : "failed",
		Message: message,
	}
}

func SuccessResponse(message string) SuccessResponseJson {
	return SuccessResponseJson{
		Status : "success",
		Message: message,
	}
}

func SuccessWithDataResponse(message string, data interface{}) SuccessResponseJson {
	return SuccessResponseJson{
		Status : "success",
		Message: message,
		Data:    data,
	}
}
