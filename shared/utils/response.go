package utils

import "github.com/rizama/favorite-book-tracker/delivery/http/dto/response"

func ConstructorResponseError(statuCode int, errorMsg string) (response.ResponseErrorDTO, int) {
	resp := response.ResponseErrorDTO{
		StatusCode: statuCode,
		Error:      errorMsg,
	}

	return resp, statuCode
}
