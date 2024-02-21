package utils

import (
	"fmt"

	"github.com/rizama/favorite-book-tracker/delivery/http/dto/response"
)

func ResponseError(statuCode int, errorMsg string) (response.ResponseErrorDTO, int) {
	resp := response.ResponseErrorDTO{
		StatusCode: statuCode,
		Error:      errorMsg,
	}

	return resp, statuCode
}

func ResponseSuccess(statuCode int, msg string, data any) response.ResponseSuccessDTO {
	resp := response.ResponseSuccessDTO{
		StatusCode: statuCode,
		Message:    msg,
		Data:       data,
	}

	return resp
}

type ErrorResponseMsg struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

func ErrValidationMsg(errs []ErrorResponseMsg) []string {
	errMsgs := make([]string, 0)
	for _, err := range errs {
		errMsgs = append(errMsgs, fmt.Sprintf(
			"[%s]: '%v' -> Needs to implement '%s'",
			err.FailedField,
			err.Value,
			err.Tag,
		))
	}

	return errMsgs
}
