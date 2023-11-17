package helper

type ErrorResponseJson struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type SuccessResponseJson struct {
	Status    bool        `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Pagnation interface{} `json:"pagnation,omiempty"`
}

func ErrorResponse(message string) ErrorResponseJson {
	return ErrorResponseJson{
		Status:  false,
		Message: message,
	}
}

func SuccessResponse(message string) SuccessResponseJson {
	return SuccessResponseJson{
		Status:  true,
		Message: message,
	}
}

func SuccessWithDataResponse(message string, data interface{}) SuccessResponseJson {
	return SuccessResponseJson{
		Status:  true,
		Message: message,
		Data:    data,
	}
}
func SuccessWithPagnationAndDataResponse(message string, data interface{}, pagnation interface{}) SuccessResponseJson {
	return SuccessResponseJson{
		Status:  true,
		Message: message,
		Data:    data,
		Pagnation: pagnation,
	}
}
