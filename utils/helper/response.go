package helper

import "recything/utils/pagination"

type ErrorResponseJson struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type SuccessResponseJson struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type SuccessResponseJsonWithPagenation struct {
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
}

type SuccessResponseJsonWithPagination struct {
	Status      bool        `json:"status"`
	Message     string      `json:"message"`
	DataMessage string      `json:"data_message,omitempty"`
	Data        interface{} `json:"data,omitempty"`
	Pagination  interface{} `json:"pagination,omitempty"`
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

func SuccessWithPagnation(message string, data interface{}, pagnation interface{}) SuccessResponseJsonWithPagenation {
	return SuccessResponseJsonWithPagenation{
		Status:     true,
		Message:    message,
		Data:       data,
		Pagination: pagnation,
	}
}

func SuccessWithPaginationAndDataResponse(message string, DataMessage string, data interface{}, paginationInfo pagination.PageInfo, totalData int) SuccessResponseJsonWithPagination {
	paginationMessage := pagination.PaginationMessage(paginationInfo, totalData)
	return SuccessResponseJsonWithPagination{
		Status:      true,
		Message:     message,
		DataMessage: paginationMessage,
		Data:        data,
		Pagination:  paginationInfo,
	}
}
