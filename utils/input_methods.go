package utils

import (
	"github.com/ssssshel/restponses-go/internal/interfaces"
	"github.com/ssssshel/restponses-go/internal/methods"
)

func WMessage(message string) methods.BaseResponseOpt {
	return func(r *interfaces.BaseResponse) {
		r.ServerMessage = message
	}
}

func WDetails(details string) methods.BaseResponseOpt {
	return func(r *interfaces.BaseResponse) {
		r.Details = details
	}
}

func WResourceName(resource string) methods.BaseResponseOpt {
	return func(r *interfaces.BaseResponse) {
		r.ConsultedResource = resource
	}
}

func WData(data interface{}) methods.Status2xxResponseOpt {
	return func(r *interfaces.Status2xx_Response) {
		r.Data = data
	}
}
