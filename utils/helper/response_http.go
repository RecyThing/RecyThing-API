package helper

type ErrorResponseJson struct {
	Message string `json:"message"`
}

type SuccessResponseJson struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func FailedResponse(message string) ErrorResponseJson {
	return ErrorResponseJson{
		Message: message,
	}
}

func SuccessResponse(message string) SuccessResponseJson {
	return SuccessResponseJson{
		Message: message,
	}
}

func SuccessWithDataResponse(message string, data interface{}) SuccessResponseJson {
	return SuccessResponseJson{
		Message: message,
		Data:    data,
	}
}
