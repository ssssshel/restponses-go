package restponses

import (
	"github.com/ssssshel/restponses-go/internal/interfaces"
	"github.com/ssssshel/restponses-go/internal/methods"
	"github.com/ssssshel/restponses-go/internal/states"
)

// 100s

func Response1xxInformative(statusCode StatusCode1xx, serverMessage, details, consultedResource string) *interfaces.BaseResponse {
	var defaultContent states.HttpStatus

	switch statusCode {
	case Status100:
		defaultContent = states.DefaultStatesContent[states.Status100Continue]
	case Status101:
		defaultContent = states.DefaultStatesContent[states.Status101SwitchingProtocols]
	case Status102:
		defaultContent = states.DefaultStatesContent[states.Status102Processing]
	case Status103:
		defaultContent = states.DefaultStatesContent[states.Status103Checkpoint]
	default:
		return nil
	}

	res := &interfaces.BaseResponse{
		HttpStatus:        defaultContent.Code,
		ServerMessage:     methods.DefaultStringReplacer(serverMessage, defaultContent.Message),
		Details:           methods.DefaultStringReplacer(details, defaultContent.Details),
		ConsultedResource: consultedResource,
	}

	return res
}

// 200s

func Response2xxSuccessfull(statusCode StatusCode2xx, serverMessage, details, consultedResource string, data interface{}, statusOptions methods.Response2xxOpt) *interfaces.GenericSuccessfullResponse {
	var defaultContent states.HttpStatus

	switch statusCode {
	case Status200:
		defaultContent = states.DefaultStatesContent[states.Status200Ok]
	case Status201:
		defaultContent = states.DefaultStatesContent[states.Status201Created]
	case Status202:
		defaultContent = states.DefaultStatesContent[states.Status202Accepted]
	case Status203:
		defaultContent = states.DefaultStatesContent[states.Status203NonAuthoritativeInformation]
	case Status204:
		defaultContent = states.DefaultStatesContent[states.Status204NoContent]
	case Status205:
		defaultContent = states.DefaultStatesContent[states.Status205ResetContent]
	case Status206:
		defaultContent = states.DefaultStatesContent[states.Status206PartialContent]
	case Status207:
		defaultContent = states.DefaultStatesContent[states.Status207MultiStatus]
	case Status208:
		defaultContent = states.DefaultStatesContent[states.Status208AlreadyReported]
	case Status226:
		defaultContent = states.DefaultStatesContent[states.Status226IMUsed]
	default:
		return nil
	}

	res := &interfaces.GenericSuccessfullResponse{
		BaseResponse: &interfaces.BaseResponse{
			HttpStatus:        defaultContent.Code,
			ServerMessage:     methods.DefaultStringReplacer(serverMessage, defaultContent.Message),
			Details:           methods.DefaultStringReplacer(details, defaultContent.Details),
			ConsultedResource: consultedResource,
		},
		Data: data,
		SuccessErrorProps: &interfaces.SuccessErrorProps{
			Success: true,
			Error:   false,
		},
	}

	if statusOptions != nil {
		statusOptions(res)
	}

	return res
}

// 300s

func Response3xxRedirection(statusCode StatusCode3xx, serverMessage, details, consultedResource string, statusOptions methods.Response3xxOpt) *interfaces.GenericRedirectionResponse {
	var defaultContent states.HttpStatus

	switch statusCode {
	case Status300:
		defaultContent = states.DefaultStatesContent[states.Status300MultipleChoices]
	case Status301:
		defaultContent = states.DefaultStatesContent[states.Status301MovedPermanently]
	case Status302:
		defaultContent = states.DefaultStatesContent[states.Status302Found]
	case Status303:
		defaultContent = states.DefaultStatesContent[states.Status303SeeOther]
	case Status304:
		defaultContent = states.DefaultStatesContent[states.Status304NotModified]
	case Status305:
		defaultContent = states.DefaultStatesContent[states.Status305UseProxy]
	case Status307:
		defaultContent = states.DefaultStatesContent[states.Status307TemporaryRedirect]
	case Status308:
		defaultContent = states.DefaultStatesContent[states.Status308PermanentRedirect]
	}

	res := &interfaces.GenericRedirectionResponse{
		BaseResponse: &interfaces.BaseResponse{
			HttpStatus:        defaultContent.Code,
			ServerMessage:     methods.DefaultStringReplacer(serverMessage, defaultContent.Message),
			Details:           methods.DefaultStringReplacer(details, defaultContent.Details),
			ConsultedResource: consultedResource,
		},
	}

	if statusOptions != nil {
		statusOptions(res)
	}

	return res
}

// 400s

func Response4xxClientError(statusCode StatusCode4xx, serverMessage, details, consultedResource string, arbitraryErrors interface{}, statusOptions methods.ResponseErrorOpt) *interfaces.GenericErrorResponse {
	var defaultContent states.HttpStatus

	switch statusCode {
	case Status400:
		defaultContent = states.DefaultStatesContent[states.Status400BadRequest]
	case Status401:
		defaultContent = states.DefaultStatesContent[states.Status401Unauthorized]
	case Status402:
		defaultContent = states.DefaultStatesContent[states.Status402PaymentRequired]
	case Status403:
		defaultContent = states.DefaultStatesContent[states.Status403Forbidden]
	case Status404:
		defaultContent = states.DefaultStatesContent[states.Status404NotFound]
	case Status405:
		defaultContent = states.DefaultStatesContent[states.Status405MethodNotAllowed]
	case Status406:
		defaultContent = states.DefaultStatesContent[states.Status406NotAcceptable]
	case Status407:
		defaultContent = states.DefaultStatesContent[states.Status407ProxyAuthenticationRequired]
	case Status408:
		defaultContent = states.DefaultStatesContent[states.Status408RequestTimeout]
	case Status409:
		defaultContent = states.DefaultStatesContent[states.Status409Conflict]
	case Status410:
		defaultContent = states.DefaultStatesContent[states.Status410Gone]
	case Status411:
		defaultContent = states.DefaultStatesContent[states.Status411LengthRequired]
	case Status412:
		defaultContent = states.DefaultStatesContent[states.Status412PreconditionFailed]
	case Status413:
		defaultContent = states.DefaultStatesContent[states.Status413PayloadTooLarge]
	case Status414:
		defaultContent = states.DefaultStatesContent[states.Status414RequestUriTooLong]
	case Status415:
		defaultContent = states.DefaultStatesContent[states.Status415UnsupportedMediaType]
	case Status416:
		defaultContent = states.DefaultStatesContent[states.Status416RequestRangeNotSatisfiable]
	case Status417:
		defaultContent = states.DefaultStatesContent[states.Status417ExpectationFailed]
	case Status418:
		defaultContent = states.DefaultStatesContent[states.Status418Teapot]
	case Status421:
		defaultContent = states.DefaultStatesContent[states.Status421MisdirectedRequest]
	case Status422:
		defaultContent = states.DefaultStatesContent[states.Status422UnprocessableEntity]
	case Status423:
		defaultContent = states.DefaultStatesContent[states.Status423Locked]
	case Status424:
		defaultContent = states.DefaultStatesContent[states.Status424FailedDependency]
	case Status425:
		defaultContent = states.DefaultStatesContent[states.Status425Unassigned]
	case Status426:
		defaultContent = states.DefaultStatesContent[states.Status426UpgradeRequired]
	case Status428:
		defaultContent = states.DefaultStatesContent[states.Status428PreconditionRequired]
	case Status429:
		defaultContent = states.DefaultStatesContent[states.Status429TooManyRequests]
	case Status431:
		defaultContent = states.DefaultStatesContent[states.Status431RequestHeaderFieldsTooLarge]
	case Status451:
		defaultContent = states.DefaultStatesContent[states.Status451UnavailableForLegalReasons]
	}

	res := &interfaces.GenericErrorResponse{
		BaseResponse: &interfaces.BaseResponse{
			HttpStatus:        defaultContent.Code,
			ServerMessage:     methods.DefaultStringReplacer(serverMessage, defaultContent.Message),
			Details:           methods.DefaultStringReplacer(details, defaultContent.Details),
			ConsultedResource: consultedResource,
		},

		Errors: arbitraryErrors,

		SuccessErrorProps: &interfaces.SuccessErrorProps{
			Success: false,
			Error:   true,
		},
	}

	if statusOptions != nil {
		statusOptions(res)
	}

	return res
}

// 500s

func Response5xxServerError(statusCode StatusCode5xx, serverMessage, details, consultedResource string, arbitraryErrors interface{}) *interfaces.GenericErrorResponse {
	var defaultContent states.HttpStatus

	switch statusCode {
	case Status500:
		defaultContent = states.DefaultStatesContent[states.Status500InternalServerError]
	case Status501:
		defaultContent = states.DefaultStatesContent[states.Status501NotImplemented]
	case Status502:
		defaultContent = states.DefaultStatesContent[states.Status502BadGateway]
	case Status503:
		defaultContent = states.DefaultStatesContent[states.Status503ServiceUnavailable]
	case Status504:
		defaultContent = states.DefaultStatesContent[states.Status504GatewayTimeout]
	case Status505:
		defaultContent = states.DefaultStatesContent[states.Status505HTTPVersionNotSupported]
	case Status506:
		defaultContent = states.DefaultStatesContent[states.Status506VariantAlsoNegotiates]
	case Status507:
		defaultContent = states.DefaultStatesContent[states.Status507InsufficientStorage]
	case Status508:
		defaultContent = states.DefaultStatesContent[states.Status508LoopDetected]
	case Status509:
		defaultContent = states.DefaultStatesContent[states.Status509BandwidthLimitExceeded]
	case Status510:
		defaultContent = states.DefaultStatesContent[states.Status510NotExtended]
	case Status511:
		defaultContent = states.DefaultStatesContent[states.Status511NetworkAuthenticationRequired]
	case Status521:
		defaultContent = states.DefaultStatesContent[states.Status521WebServerIsDown]
	}

	res := &interfaces.GenericErrorResponse{
		BaseResponse: &interfaces.BaseResponse{
			HttpStatus:        defaultContent.Code,
			ServerMessage:     methods.DefaultStringReplacer(serverMessage, defaultContent.Message),
			Details:           methods.DefaultStringReplacer(details, defaultContent.Details),
			ConsultedResource: consultedResource,
		},

		Errors: arbitraryErrors,

		SuccessErrorProps: &interfaces.SuccessErrorProps{
			Success: false,
			Error:   true,
		},
	}

	return res
}
