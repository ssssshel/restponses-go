package restponses

import (
	"github.com/ssssshel/restponses-go/internal/interfaces"
	"github.com/ssssshel/restponses-go/internal/methods"
	"github.com/ssssshel/restponses-go/internal/states"
)

// 100s

func Response1xxInformative(statusCode StatusCode1xx, input *BaseInput) *interfaces.GenericInformativeResponse {
	var defaultContent states.HttpStatus

	switch statusCode {
	case Status100Continue:
		defaultContent = states.DefaultStatesContent[states.Status100Continue]
	case Status101SwitchingProtocols:
		defaultContent = states.DefaultStatesContent[states.Status101SwitchingProtocols]
	case Status102Processing:
		defaultContent = states.DefaultStatesContent[states.Status102Processing]
	case Status103EarlyHints:
		defaultContent = states.DefaultStatesContent[states.Status103Checkpoint]
	default:
		return nil
	}

	res := &interfaces.GenericInformativeResponse{
		HttpStatus: &defaultContent.Code,
		BaseResponse: &interfaces.BaseResponse{
			ServerMessage:     methods.DefaultStringReplacer(input.ServerMessage, defaultContent.Message),
			Detail:            methods.DefaultStringReplacer(input.Detail, defaultContent.Details),
			ConsultedResource: input.ConsultedResource,
		},
	}

	return res
}

// 200s

func Response2xxSuccessfull(statusCode StatusCode2xx, input *BaseSuccessfulInput) *interfaces.GenericSuccessfullResponse {
	var defaultContent states.HttpStatus

	switch statusCode {
	case Status200Ok:
		defaultContent = states.DefaultStatesContent[states.Status200Ok]
	case Status201Continue:
		defaultContent = states.DefaultStatesContent[states.Status201Created]
	case Status202Accepted:
		defaultContent = states.DefaultStatesContent[states.Status202Accepted]
	case Status203NonAuthoritativeInfo:
		defaultContent = states.DefaultStatesContent[states.Status203NonAuthoritativeInformation]
	case Status204NoContent:
		defaultContent = states.DefaultStatesContent[states.Status204NoContent]
	case Status205ResetContent:
		defaultContent = states.DefaultStatesContent[states.Status205ResetContent]
	case Status206PartialContent:
		defaultContent = states.DefaultStatesContent[states.Status206PartialContent]
	case Status207MultiStatus:
		defaultContent = states.DefaultStatesContent[states.Status207MultiStatus]
	case Status208AlreadyReported:
		defaultContent = states.DefaultStatesContent[states.Status208AlreadyReported]
	case Status226IMUsed:
		defaultContent = states.DefaultStatesContent[states.Status226IMUsed]
	default:
		return nil
	}

	res := &interfaces.GenericSuccessfullResponse{
		HttpStatus: &defaultContent.Code,
		BaseResponse: &interfaces.BaseResponse{
			ServerMessage:     methods.DefaultStringReplacer(input.ServerMessage, defaultContent.Message),
			Detail:            methods.DefaultStringReplacer(input.Detail, defaultContent.Details),
			ConsultedResource: input.ConsultedResource,
		},
		Data: input.Data,
		SuccessErrorProps: &interfaces.SuccessErrorProps{
			Success: true,
			Error:   false,
		},
	}

	if input.StatusOptions != nil {
		input.StatusOptions(res)
	}

	return res
}

// 300s

func Response3xxRedirection(statusCode StatusCode3xx, input *BaseRedirectionInput) *interfaces.GenericRedirectionResponse {
	var defaultContent states.HttpStatus

	switch statusCode {
	case Status300MultipleChoices:
		defaultContent = states.DefaultStatesContent[states.Status300MultipleChoices]
	case Status301MovedPermanently:
		defaultContent = states.DefaultStatesContent[states.Status301MovedPermanently]
	case Status302Found:
		defaultContent = states.DefaultStatesContent[states.Status302Found]
	case Status303SeeOther:
		defaultContent = states.DefaultStatesContent[states.Status303SeeOther]
	case Status304NotModified:
		defaultContent = states.DefaultStatesContent[states.Status304NotModified]
	case Status305UseProxy:
		defaultContent = states.DefaultStatesContent[states.Status305UseProxy]
	case Status307TemporaryRedirect:
		defaultContent = states.DefaultStatesContent[states.Status307TemporaryRedirect]
	case Status308PermanentRedirect:
		defaultContent = states.DefaultStatesContent[states.Status308PermanentRedirect]
	}

	res := &interfaces.GenericRedirectionResponse{
		HttpStatus: &defaultContent.Code,
		BaseResponse: &interfaces.BaseResponse{
			ServerMessage:     methods.DefaultStringReplacer(input.ServerMessage, defaultContent.Message),
			Detail:            methods.DefaultStringReplacer(input.Detail, defaultContent.Details),
			ConsultedResource: input.Detail,
		},
	}

	if input.StatusOptions != nil {
		input.StatusOptions(res)
	}

	return res
}

// 400s

func Response4xxClientError(statusCode StatusCode4xx, input *BaseClientErrorInput) *interfaces.GenericErrorResponse {
	var defaultContent states.HttpStatus

	switch statusCode {
	case Status400BadRequest:
		defaultContent = states.DefaultStatesContent[states.Status400BadRequest]
	case Status401Unauthorized:
		defaultContent = states.DefaultStatesContent[states.Status401Unauthorized]
	case Status402PaymentRequired:
		defaultContent = states.DefaultStatesContent[states.Status402PaymentRequired]
	case Status403Forbidden:
		defaultContent = states.DefaultStatesContent[states.Status403Forbidden]
	case Status404NotFound:
		defaultContent = states.DefaultStatesContent[states.Status404NotFound]
	case Status405MethodNotAllowed:
		defaultContent = states.DefaultStatesContent[states.Status405MethodNotAllowed]
	case Status406NotAcceptable:
		defaultContent = states.DefaultStatesContent[states.Status406NotAcceptable]
	case Status407ProxyAuthenticationRequired:
		defaultContent = states.DefaultStatesContent[states.Status407ProxyAuthenticationRequired]
	case Status408RequestTimeout:
		defaultContent = states.DefaultStatesContent[states.Status408RequestTimeout]
	case Status409Conflict:
		defaultContent = states.DefaultStatesContent[states.Status409Conflict]
	case Status410Gone:
		defaultContent = states.DefaultStatesContent[states.Status410Gone]
	case Status411LengthRequired:
		defaultContent = states.DefaultStatesContent[states.Status411LengthRequired]
	case Status412PreconditionFailed:
		defaultContent = states.DefaultStatesContent[states.Status412PreconditionFailed]
	case Status413PayloadTooLarge:
		defaultContent = states.DefaultStatesContent[states.Status413PayloadTooLarge]
	case Status414RequestUriTooLong:
		defaultContent = states.DefaultStatesContent[states.Status414RequestUriTooLong]
	case Status415UnsupportedMediaType:
		defaultContent = states.DefaultStatesContent[states.Status415UnsupportedMediaType]
	case Status416RequestRangeNotSatisfiable:
		defaultContent = states.DefaultStatesContent[states.Status416RequestRangeNotSatisfiable]
	case Status417ExpectationFailed:
		defaultContent = states.DefaultStatesContent[states.Status417ExpectationFailed]
	case Status418Teapot:
		defaultContent = states.DefaultStatesContent[states.Status418Teapot]
	case Status421MisdirectedRequest:
		defaultContent = states.DefaultStatesContent[states.Status421MisdirectedRequest]
	case Status422UnprocessableEntity:
		defaultContent = states.DefaultStatesContent[states.Status422UnprocessableEntity]
	case Status423Locked:
		defaultContent = states.DefaultStatesContent[states.Status423Locked]
	case Status424FailedDependency:
		defaultContent = states.DefaultStatesContent[states.Status424FailedDependency]
	case Status425Unassigned:
		defaultContent = states.DefaultStatesContent[states.Status425Unassigned]
	case Status426UpgradeRequired:
		defaultContent = states.DefaultStatesContent[states.Status426UpgradeRequired]
	case Status428PreconditionRequired:
		defaultContent = states.DefaultStatesContent[states.Status428PreconditionRequired]
	case Status429TooManyRequests:
		defaultContent = states.DefaultStatesContent[states.Status429TooManyRequests]
	case Status431RequestHeaderFieldsTooLarge:
		defaultContent = states.DefaultStatesContent[states.Status431RequestHeaderFieldsTooLarge]
	case Status451UnavailableForLegalReasons:
		defaultContent = states.DefaultStatesContent[states.Status451UnavailableForLegalReasons]
	}

	res := &interfaces.GenericErrorResponse{
		HttpStatus: &defaultContent.Code,
		BaseResponse: &interfaces.BaseResponse{
			ServerMessage:     methods.DefaultStringReplacer(input.ServerMessage, defaultContent.Message),
			Detail:            methods.DefaultStringReplacer(input.Detail, defaultContent.Details),
			ConsultedResource: input.ConsultedResource,
		},

		Errors:           input.Errors,
		ErrorName:        input.ErrorName,
		ErrorCode:        input.ErrorCode,
		ErrorDescription: input.ErrorDescription,

		ErrorDetails: &interfaces.ErrorDetails{},
		SuccessErrorProps: &interfaces.SuccessErrorProps{
			Success: false,
			Error:   true,
		},
	}

	if input.StatusOptions != nil {
		input.StatusOptions(res)
	}

	return res
}

// 500s

func Response5xxServerError(statusCode StatusCode5xx, input *BaseServerErrorInput) *interfaces.GenericErrorResponse {
	var defaultContent states.HttpStatus

	switch statusCode {
	case Status500InternalServerError:
		defaultContent = states.DefaultStatesContent[states.Status500InternalServerError]
	case Status501NotImplemented:
		defaultContent = states.DefaultStatesContent[states.Status501NotImplemented]
	case Status502BadGateway:
		defaultContent = states.DefaultStatesContent[states.Status502BadGateway]
	case Status503ServiceUnavailable:
		defaultContent = states.DefaultStatesContent[states.Status503ServiceUnavailable]
	case Status504GatewayTimeout:
		defaultContent = states.DefaultStatesContent[states.Status504GatewayTimeout]
	case Status505HttpVersionNotSupported:
		defaultContent = states.DefaultStatesContent[states.Status505HTTPVersionNotSupported]
	case Status506VariantAlsoNegotiates:
		defaultContent = states.DefaultStatesContent[states.Status506VariantAlsoNegotiates]
	case Status507InsufficientStorage:
		defaultContent = states.DefaultStatesContent[states.Status507InsufficientStorage]
	case Status508LoopDetected:
		defaultContent = states.DefaultStatesContent[states.Status508LoopDetected]
	case Status509Unassigned:
		defaultContent = states.DefaultStatesContent[states.Status509BandwidthLimitExceeded]
	case Status510NotExtended:
		defaultContent = states.DefaultStatesContent[states.Status510NotExtended]
	case Status511NetworkAuthenticationRequired:
		defaultContent = states.DefaultStatesContent[states.Status511NetworkAuthenticationRequired]
	case Status521WebServerIsDown:
		defaultContent = states.DefaultStatesContent[states.Status521WebServerIsDown]
	}

	res := &interfaces.GenericErrorResponse{
		HttpStatus: &defaultContent.Code,
		BaseResponse: &interfaces.BaseResponse{
			ServerMessage:     methods.DefaultStringReplacer(input.ServerMessage, defaultContent.Message),
			Detail:            methods.DefaultStringReplacer(input.Detail, defaultContent.Details),
			ConsultedResource: input.ConsultedResource,
		},

		Errors:           input.Errors,
		ErrorName:        input.ErrorName,
		ErrorCode:        input.ErrorCode,
		ErrorDescription: input.ErrorDescription,

		ErrorDetails: &interfaces.ErrorDetails{},
		SuccessErrorProps: &interfaces.SuccessErrorProps{
			Success: false,
			Error:   true,
		},
	}

	return res
}
