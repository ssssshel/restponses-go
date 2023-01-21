package restponses

import (
	"github.com/ssssshel/restponses-go/internal/interfaces"
	"github.com/ssssshel/restponses-go/internal/methods"
	"github.com/ssssshel/restponses-go/internal/states"
)

/*
You can use these functions as parameters:
  - WMessage() => Allows to add "serverMessage" field in your response
  - WDetails() => Allows to add "details" field in your response
  - WResourceName() => Allows to add "consultedResource" field in your response
*/
func Response100Continue(serverMessage, details, consultedResource string) *interfaces.BaseResponse {

	defaultContent := states.DefaultStatesContent[states.Status100Continue]

	res := &interfaces.BaseResponse{
		HttpStatus:        defaultContent.Code,
		ServerMessage:     methods.DefaultStringReplacer(serverMessage, defaultContent.Message),
		Details:           methods.DefaultStringReplacer(details, defaultContent.Details),
		ConsultedResource: consultedResource,
	}

	return res
}

/*
You can use these functions as parameters:
  - WMessage() => Allows to add "serverMessage" field in your response
  - WDetails() => Allows to add "details" field in your response
  - WResourceName() => Allows to add "consultedResource" field in your response
*/
func Response101SwitchingProtocols(serverMessage, details, consultedResource string) *interfaces.BaseResponse {

	defaultContent := states.DefaultStatesContent[states.Status101SwitchingProtocols]

	res := &interfaces.BaseResponse{
		HttpStatus:        defaultContent.Code,
		ServerMessage:     methods.DefaultStringReplacer(serverMessage, defaultContent.Message),
		Details:           methods.DefaultStringReplacer(details, defaultContent.Details),
		ConsultedResource: consultedResource,
	}

	return res
}

/*
You can use these functions as parameters:
  - WMessage() => Allows to add "serverMessage" field in your response
  - WDetails() => Allows to add "details" field in your response
  - WResourceName() => Allows to add "consultedResource" field in your response
*/
func Response102Processing(serverMessage, details, consultedResource string) *interfaces.BaseResponse {

	defaultContent := states.DefaultStatesContent[states.Status102Processing]

	res := &interfaces.BaseResponse{
		HttpStatus:        defaultContent.Code,
		ServerMessage:     methods.DefaultStringReplacer(serverMessage, defaultContent.Message),
		Details:           methods.DefaultStringReplacer(details, defaultContent.Details),
		ConsultedResource: consultedResource,
	}

	return res
}

/*
You can use these functions as parameters:
  - WMessage() => Allows to add "serverMessage" field in your response
  - WDetails() => Allows to add "details" field in your response
  - WResourceName() => Allows to add "consultedResource" field in your response
*/
func Response103Checkpoint(serverMessage, details, consultedResource string) *interfaces.BaseResponse {

	defaultContent := states.DefaultStatesContent[states.Status103Checkpoint]

	res := &interfaces.BaseResponse{
		HttpStatus:        defaultContent.Code,
		ServerMessage:     methods.DefaultStringReplacer(serverMessage, defaultContent.Message),
		Details:           methods.DefaultStringReplacer(details, defaultContent.Details),
		ConsultedResource: consultedResource,
	}

	return res
}

// 200s

/*
You can use these functions as parameters:
  - WMessage() => Allows to add "serverMessage" field in your response
  - WDetails() => Allows to add "details" field in your response
  - WResourceName() => Allows to add "consultedResource" field in your response
*/
func Response200Ok(serverMessage, details, consultedResource string, data interface{}) *interfaces.Status2xx_Response {

	defaultContent := states.DefaultStatesContent[states.Status200Ok]

	res := &interfaces.Status2xx_Response{
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

	return res
}

func Response201Created(serverMessage, details, consultedResource string, data interface{}, location string) *interfaces.Status201Created_Response {

	defaultContent := states.DefaultStatesContent[states.Status201Created]

	res := &interfaces.Status201Created_Response{
		Status2xx_Response: &interfaces.Status2xx_Response{
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
		},
		Location: location,
	}

	return res
}

func Response202Accepted(serverMessage, details, consultedResource string, data interface{}, requestId string) *interfaces.Status202Accepted_Response {

	defaultContent := states.DefaultStatesContent[states.Status202Accepted]

	res := &interfaces.Status202Accepted_Response{
		Status2xx_Response: &interfaces.Status2xx_Response{
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
		},
		RequestId: requestId,
	}

	return res
}

func Response203NonAuthoritativeInformation(serverMessage, details, consultedResource string, data interface{}, sourceName, sourceDescription, sourceUrl string) *interfaces.Status203NonAI_Response {

	defaultContent := states.DefaultStatesContent[states.Status203NonAuthoritativeInformation]

	res := &interfaces.Status203NonAI_Response{
		Status2xx_Response: &interfaces.Status2xx_Response{
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
		},
		Source: &interfaces.Source{
			Name:        sourceName,
			Description: sourceDescription,
			Source:      sourceUrl,
		},
	}

	return res
}

func Response204NoContent(serverMessage, details, consultedResource string) *interfaces.Status2xx_Response {

	defaultContent := states.DefaultStatesContent[states.Status204NoContent]

	res := &interfaces.Status2xx_Response{
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

	return res
}

func Response205ResetContent(serverMessage, details, consultedResource string, data interface{}) *interfaces.Status2xx_Response {

	defaultContent := states.DefaultStatesContent[states.Status205ResetContent]

	res := &interfaces.Status2xx_Response{
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

	return res
}

func Response206PartialContent(serverMessage, details, consultedResource string, data interface{}) *interfaces.Status2xx_Response {

	defaultContent := states.DefaultStatesContent[states.Status206PartialContent]

	res := &interfaces.Status2xx_Response{
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

	return res
}

// func Response207MultiStatus(serverMessage, details, consultedResource string, data interface{}, multipleStates []*interfaces.BaseResponse) *interfaces.Status207MultiStatus_Response {

// 	defaultContent := states.DefaultStatesContent[states.Status207MultiStatus]

// 	res := &interfaces.Status207MultiStatus_Response{
// 		Status2xx_Response: &interfaces.Status2xx_Response{
// 			BaseResponse: &interfaces.BaseResponse{
// 				HttpStatus:        defaultContent.Code,
// 				ServerMessage:     methods.DefaultStringReplacer(serverMessage, defaultContent.Message),
// 				Details:           methods.DefaultStringReplacer(details, defaultContent.Details),
// 				ConsultedResource: consultedResource,
// 			},
// 			Data: data,
// 			SuccessErrorProps: &interfaces.SuccessErrorProps{
// 				Success: true,
// 				Error:   false,
// 			},
// 		},
// 		States: multipleStates,
// 	}

// 	return res
// }

func Response208AlreadyReported(serverMessage, details, consultedResource string, data interface{}) *interfaces.Status2xx_Response {

	defaultContent := states.DefaultStatesContent[states.Status208AlreadyReported]

	res := &interfaces.Status2xx_Response{
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

	return res
}

func Response226IMUsed(serverMessage, details, consultedResource string, data interface{}) *interfaces.Status2xx_Response {

	defaultContent := states.DefaultStatesContent[states.Status226IMUsed]

	res := &interfaces.Status2xx_Response{
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

	return res
}
