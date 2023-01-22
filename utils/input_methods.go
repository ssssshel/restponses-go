package utils

import (
	"github.com/ssssshel/restponses-go/internal/interfaces"
	"github.com/ssssshel/restponses-go/internal/methods"
)

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

// 400s

func Status404Opt(notFoundResourceName string) methods.ResponseErrorOpt {
	return func(r *interfaces.GenericErrorResponse) {
		r.ErrorDetails.NotFoundResource = notFoundResourceName
	}
}

func Status405Opt(allowedMethods []string) methods.ResponseErrorOpt {
	return func(r *interfaces.GenericErrorResponse) {
		r.ErrorDetails.AllowedMethods = allowedMethods
	}
}

func Status406Opt(allowedRepresentations []string) methods.ResponseErrorOpt {
	return func(r *interfaces.GenericErrorResponse) {
		r.ErrorDetails.AllowedRepresentations = allowedRepresentations
	}
}

func Status407Opt(authenticationType, realm string) methods.ResponseErrorOpt {
	return func(r *interfaces.GenericErrorResponse) {
		r.ErrorDetails.AuthenticationType = authenticationType
		r.ErrorDetails.Realm = realm
	}
}

func Status408Opt(timeWaitng string) methods.ResponseErrorOpt {
	return func(r *interfaces.GenericErrorResponse) {
		r.ErrorDetails.TimeWaiting = timeWaitng
	}
}

func Status409Opt(conflictResource, conflictId string) methods.ResponseErrorOpt {
	return func(r *interfaces.GenericErrorResponse) {
		r.ErrorDetails.ConflictResource = conflictResource
		r.ErrorDetails.ConflictId = conflictId
	}
}

func Status410Opt(goneResource, reason string) methods.ResponseErrorOpt {
	return func(r *interfaces.GenericErrorResponse) {
		r.ErrorDetails.GoneResource = goneResource
		r.ErrorDetails.Reason = reason
	}
}

func Status411Opt() methods.ResponseErrorOpt {
	return func(r *interfaces.GenericErrorResponse) {
		r.ErrorDetails.RequiredHeader = "Content-Length"
	}
}

func Status413Opt(maxAllowedSize string) methods.ResponseErrorOpt {
	return func(r *interfaces.GenericErrorResponse) {
		r.ErrorDetails.MaxAllowedSize = maxAllowedSize
	}
}

func Status414Opt(maxAllowedLength uint) methods.ResponseErrorOpt {
	return func(r *interfaces.GenericErrorResponse) {
		r.ErrorDetails.MaxAllowedLength = maxAllowedLength
	}
}

func Status415Opt(supportedMediaTypes []string) methods.ResponseErrorOpt {
	return func(r *interfaces.GenericErrorResponse) {
		r.ErrorDetails.SupportedMediaTypes = supportedMediaTypes
	}
}

func Status416Opt(requestedContentRange, supportedContentRange string) methods.ResponseErrorOpt {
	return func(r *interfaces.GenericErrorResponse) {
		r.ErrorDetails.RequestedContentRange = requestedContentRange
		r.ErrorDetails.SupportedContentRange = supportedContentRange
	}
}

func Status423Opt(lockedResource string) methods.ResponseErrorOpt {
	return func(r *interfaces.GenericErrorResponse) {
		r.ErrorDetails.LockedResource = lockedResource
	}
}

func Status424Opt(failedDependency string) methods.ResponseErrorOpt {
	return func(r *interfaces.GenericErrorResponse) {
		r.ErrorDetails.FailedDependency = failedDependency
	}
}
