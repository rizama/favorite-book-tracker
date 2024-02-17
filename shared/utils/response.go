package utils

import "github.com/rizama/favorite-book-tracker/delivery/http/dto/response"

func ConstructorResponseError(statuCode int, errorMsg string) (response.ResponseErrorDTO, int) {
	resp := response.ResponseErrorDTO{
		StatusCode: statuCode,
		Error:      errorMsg,
	}

	return resp, statuCode
}

func ConstructorResponseSuccess(statuCode int, msg string, data any) response.ResponseSuccessDTO {
	resp := response.ResponseSuccessDTO{
		StatusCode: statuCode,
		Message:    msg,
		Data:       data,
	}

	return resp
}
