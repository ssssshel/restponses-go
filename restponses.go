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
	}

	res := &interfaces.GenericSuccessfullResponse{
		BaseResponse: &interfaces.BaseResponse{
			HttpStatus:        defaultContent.Code,
			ServerMessage:     methods.DefaultStringReplacer(serverMessage, defaultContent.Message),
			Details:           methods.DefaultStringReplacer(details, defaultContent.Details),
			ConsultedResource: consultedResource,
		},

		SuccessErrorProps: &interfaces.SuccessErrorProps{
			Success: true,
			Error:   false,
		},
	}

	// for _, opt := range statusOptions {
	// 	opt(res)
	// }

	statusOptions(res)

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

	statusOptions(res)

	return res
}

func Response300MultipleChoices(serverMessage, details, consultedResource string, options []interface{}) *interfaces.Status300MultipleChoices_Response {

	defaultContent := states.DefaultStatesContent[states.Status300MultipleChoices]

	res := &interfaces.Status300MultipleChoices_Response{
		Status3xx_Response: &interfaces.Status3xx_Response{
			BaseResponse: &interfaces.BaseResponse{
				HttpStatus:        defaultContent.Code,
				ServerMessage:     methods.DefaultStringReplacer(serverMessage, defaultContent.Message),
				Details:           methods.DefaultStringReplacer(details, defaultContent.Details),
				ConsultedResource: consultedResource,
			},
		},
		Options: options,
	}

	return res
}

func Response301MovedPermanently(serverMessage, details, consultedResource, oldSource, newSource string) *interfaces.Status301MovedPermanently_Response {

	defaultContent := states.DefaultStatesContent[states.Status301MovedPermanently]

	res := &interfaces.Status301MovedPermanently_Response{
		Status3xx_Response: &interfaces.Status3xx_Response{
			BaseResponse: &interfaces.BaseResponse{
				HttpStatus:        defaultContent.Code,
				ServerMessage:     methods.DefaultStringReplacer(serverMessage, defaultContent.Message),
				Details:           methods.DefaultStringReplacer(details, defaultContent.Details),
				ConsultedResource: consultedResource,
			},
		},
		Sources: &interfaces.Sources{
			OldSouce:  oldSource,
			NewSource: newSource,
		},
	}

	return res
}

func Response302Found(serverMessage, details, consultedResource, redirectUrl string) *interfaces.Status302Found_Response {

	defaultContent := states.DefaultStatesContent[states.Status302Found]

	res := &interfaces.Status302Found_Response{
		Status3xx_Response: &interfaces.Status3xx_Response{
			BaseResponse: &interfaces.BaseResponse{
				HttpStatus:        defaultContent.Code,
				ServerMessage:     methods.DefaultStringReplacer(serverMessage, defaultContent.Message),
				Details:           methods.DefaultStringReplacer(details, defaultContent.Details),
				ConsultedResource: consultedResource,
			},
		},
		RedirectUrl: redirectUrl,
	}

	return res
}

func Response303SeeOther(serverMessage, details, consultedResource, redirectUrl string) *interfaces.Status303SeeOther_Response {

	defaultContent := states.DefaultStatesContent[states.Status303SeeOther]

	res := &interfaces.Status303SeeOther_Response{
		Status3xx_Response: &interfaces.Status3xx_Response{
			BaseResponse: &interfaces.BaseResponse{
				HttpStatus:        defaultContent.Code,
				ServerMessage:     methods.DefaultStringReplacer(serverMessage, defaultContent.Message),
				Details:           methods.DefaultStringReplacer(details, defaultContent.Details),
				ConsultedResource: consultedResource,
			},
		},
		RedirectUrl: redirectUrl,
	}

	return res
}

func Response304NotModified(serverMessage, details, consultedResource string) *interfaces.Status3xx_Response {

	defaultContent := states.DefaultStatesContent[states.Status304NotModified]

	res := &interfaces.Status3xx_Response{
		BaseResponse: &interfaces.BaseResponse{
			HttpStatus:        defaultContent.Code,
			ServerMessage:     methods.DefaultStringReplacer(serverMessage, defaultContent.Message),
			Details:           methods.DefaultStringReplacer(details, defaultContent.Details),
			ConsultedResource: consultedResource,
		},
	}

	return res
}

func Response305UseProxy(serverMessage, details, consultedResource, proxyUrl string) *interfaces.Status305UseProxy_Response {

	defaultContent := states.DefaultStatesContent[states.Status305UseProxy]

	res := &interfaces.Status305UseProxy_Response{
		Status3xx_Response: &interfaces.Status3xx_Response{
			BaseResponse: &interfaces.BaseResponse{
				HttpStatus:        defaultContent.Code,
				ServerMessage:     methods.DefaultStringReplacer(serverMessage, defaultContent.Message),
				Details:           methods.DefaultStringReplacer(details, defaultContent.Details),
				ConsultedResource: consultedResource,
			},
		},
		ProxyUrl: proxyUrl,
	}

	return res
}

func Response307TemporaryRedirect(serverMessage, details, consultedResource, redirectUrl string) *interfaces.Status307TemporaryRedirect_Response {

	defaultContent := states.DefaultStatesContent[states.Status307TemporaryRedirect]

	res := &interfaces.Status307TemporaryRedirect_Response{
		Status3xx_Response: &interfaces.Status3xx_Response{
			BaseResponse: &interfaces.BaseResponse{
				HttpStatus:        defaultContent.Code,
				ServerMessage:     methods.DefaultStringReplacer(serverMessage, defaultContent.Message),
				Details:           methods.DefaultStringReplacer(details, defaultContent.Details),
				ConsultedResource: consultedResource,
			},
		},
		RedirectUrl: redirectUrl,
	}

	return res
}

func Response308PermanentRedirect(serverMessage, details, consultedResource, redirectUrl string) *interfaces.Status308PermanentRedirect_Response {

	defaultContent := states.DefaultStatesContent[states.Status308PermanentRedirect]

	res := &interfaces.Status308PermanentRedirect_Response{
		Status3xx_Response: &interfaces.Status3xx_Response{
			BaseResponse: &interfaces.BaseResponse{
				HttpStatus:        defaultContent.Code,
				ServerMessage:     methods.DefaultStringReplacer(serverMessage, defaultContent.Message),
				Details:           methods.DefaultStringReplacer(details, defaultContent.Details),
				ConsultedResource: consultedResource,
			},
		},
		RedirectUrl: redirectUrl,
	}

	return res
}

func Response400BadRequest(serverMessage, details, consultedResource string, errors []interface{}) *interfaces.GenericErrorResponse {

	defaultContent := states.DefaultStatesContent[states.Status400BadRequest]

	res := &interfaces.GenericErrorResponse{
		BaseResponse: &interfaces.BaseResponse{
			HttpStatus:        defaultContent.Code,
			ServerMessage:     methods.DefaultStringReplacer(serverMessage, defaultContent.Message),
			Details:           methods.DefaultStringReplacer(details, defaultContent.Details),
			ConsultedResource: consultedResource,
		},
		Errors: errors,
		SuccessErrorProps: &interfaces.SuccessErrorProps{
			Success: false,
			Error:   true,
		},
	}

	return res
}
