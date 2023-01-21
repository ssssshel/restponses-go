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

// 200s

func Status201Opt(location string) methods.Response2xxOpt {
	return func(r *interfaces.GenericSuccessfullResponse) {
		r.Location = location
	}
}

func Status202Opt(requestId string) methods.Response2xxOpt {
	return func(r *interfaces.GenericSuccessfullResponse) {
		r.RequestId = requestId
	}
}

func Status203Opt(source *interfaces.Source) methods.Response2xxOpt {
	return func(r *interfaces.GenericSuccessfullResponse) {
		r.Source = source
	}
}

func Status207Opt(states []*interfaces.BasicStatus) methods.Response2xxOpt {
	return func(r *interfaces.GenericSuccessfullResponse) {
		r.States = states
	}
}

// 300s

func Status300Opt(options []interface{}) methods.Response3xxOpt {
	return func(r *interfaces.GenericRedirectionResponse) {
		r.Options = options
	}
}

func Status301Opt(sources *interfaces.Sources) methods.Response3xxOpt {
	return func(r *interfaces.GenericRedirectionResponse) {
		r.Sources = sources
	}
}

func Status302Opt(redirectUrl string) methods.Response3xxOpt {
	return func(r *interfaces.GenericRedirectionResponse) {
		r.RedirectUrl = redirectUrl
	}
}

func Status303Opt(redirectUrl string) methods.Response3xxOpt {
	return func(r *interfaces.GenericRedirectionResponse) {
		r.RedirectUrl = redirectUrl

	}
}

func Status305Opt(proxyUrl string) methods.Response3xxOpt {
	return func(r *interfaces.GenericRedirectionResponse) {
		r.ProxyUrl = proxyUrl
	}
}

func Status307Opt(redirectUrl string) methods.Response3xxOpt {
	return func(r *interfaces.GenericRedirectionResponse) {
		r.RedirectUrl = redirectUrl
	}
}

func Status308Opt(redirectUrl string) methods.Response3xxOpt {
	return func(r *interfaces.GenericRedirectionResponse) {
		r.RedirectUrl = redirectUrl
	}
}
