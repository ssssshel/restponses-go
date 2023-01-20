package utils

import (
	"github.com/ssssshel/restponses-go/internal/interfaces"
	"github.com/ssssshel/restponses-go/internal/methods"
)

// Allows to add "serverMessage" field in your response
func WMessage(serverMessage string) methods.BaseResponseOpt {
	return func(r *interfaces.GeneralInput) {
		r.ServerMessage = serverMessage
	}
}

// Allows to add "details" field in your response
func WDetails(details string) methods.BaseResponseOpt {
	return func(r *interfaces.GeneralInput) {
		r.Details = details
	}
}

// Allows to add "resourceName" field in your response
func WResourceName(resource string) methods.BaseResponseOpt {
	return func(r *interfaces.GeneralInput) {
		r.ConsultedResource = resource
	}
}

// Allows to add "data" field in your response
func WData(data interface{}) methods.BaseResponseOpt {
	return func(r *interfaces.GeneralInput) {
		r.Data = data
	}
}
