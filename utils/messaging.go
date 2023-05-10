package utils

import (
	"go-minimarket/models"
	"strconv"
)

func ResponseError(code int, statusX string, msg string) (int, models.ErrorResponse) {
	return code, models.ErrorResponse{
		ResponseCode:          strconv.Itoa(code),
		ResponseMessage:       statusX,
		ResponseMessageDetail: msg,
	}
}
