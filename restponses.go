package restponses

import (
	"github.com/ssssshel/restponses-go/internal/interfaces"
	"github.com/ssssshel/restponses-go/internal/methods"
	"github.com/ssssshel/restponses-go/internal/states"
)

type method int

const (
	Get method = iota
	Post
	Delete
	Put
)

type Ff struct {
	Mes string
}

// 100s

func Response100Continue(resource string, props ...methods.BaseResponseOpt) *interfaces.BaseResponse {

	defaultContent := states.DefaultStatesContent[states.Status100Continue]

	res := &interfaces.BaseResponse{
		HttpStatus:        defaultContent.Code,
		ServerMessage:     defaultContent.Message,
		ConsultedResource: resource,
	}

	for _, prop := range props {
		prop(res)
	}
	return res
}

func Response101SwitchingProtocols(resource string, props ...methods.BaseResponseOpt) *interfaces.BaseResponse {

	defaultContent := states.DefaultStatesContent[states.Status101SwitchingProtocols]

	res := &interfaces.BaseResponse{
		HttpStatus:        defaultContent.Code,
		ServerMessage:     defaultContent.Message,
		ConsultedResource: resource,
	}

	for _, prop := range props {
		prop(res)
	}
	return res
}

func Response102Processing(resource string, props ...methods.BaseResponseOpt) *interfaces.BaseResponse {

	defaultContent := states.DefaultStatesContent[states.Status102Processing]

	res := &interfaces.BaseResponse{
		HttpStatus:        defaultContent.Code,
		ServerMessage:     defaultContent.Message,
		ConsultedResource: resource,
	}

	for _, prop := range props {
		prop(res)
	}
	return res
}

func Response103Checkpoint(resource string, props ...methods.BaseResponseOpt) *interfaces.BaseResponse {

	defaultContent := states.DefaultStatesContent[states.Status103Checkpoint]

	res := &interfaces.BaseResponse{
		HttpStatus:        defaultContent.Code,
		ServerMessage:     defaultContent.Message,
		ConsultedResource: resource,
	}

	for _, prop := range props {
		prop(res)
	}
	return res
}

// 200s

func Response200Ok(resource string, props []methods.Status2xxResponseOpt, props2 []methods.BaseResponseOpt) *interfaces.Status2xx_Response {

	defaultContent := states.DefaultStatesContent[states.Status200Ok]

	baseResponse := &interfaces.BaseResponse{
		HttpStatus:        defaultContent.Code,
		ServerMessage:     defaultContent.Message,
		ConsultedResource: resource,
	}

	res := &interfaces.Status2xx_Response{
		BaseResponse: *baseResponse,
		SuccessErrorProps: interfaces.SuccessErrorProps{
			Success: true,
			Error:   false,
		},
	}

	for _, prop := range props {
		prop(res)
	}
	return res
}
