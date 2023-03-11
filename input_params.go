package restponses

import "github.com/ssssshel/restponses-go/internal/methods"

type BaseInput struct {
	ServerMessage     string
	Detail            string // Datailed info of your response
	ConsultedResource string // Name/URL of the consulted resource
}

type BaseSuccessfulInput struct {
	ServerMessage     string
	Detail            string      // Datailed info of your response
	ConsultedResource string      // Name/URL of the consulted resource
	Data              interface{} // Anything you want to return, don't use for 204 and 205 codes

	StatusOptions methods.Response2xxOpt
}

type BaseRedirectionInput struct {
	ServerMessage     string
	Detail            string // Datailed info of your response
	ConsultedResource string // Name/URL of the consulted resource

	StatusOptions methods.Response3xxOpt
}

type BaseClientErrorInput struct {
	ServerMessage     string
	Detail            string // Datailed info of your response
	ConsultedResource string // Name/URL of the consulted resource

	Errors           interface{}
	ErrorName        string // Error name | e.g., "Resource api/potato not found"
	ErrorCode        string // Error code | e.g., "ERR_NOT_FOUND"
	ErrorDescription string // Error description | e.g., "The requested resource could not be found"

	StatusOptions methods.Response4xxOpt
}

type BaseServerErrorInput struct {
	ServerMessage     string
	Detail            string // Datailed info of your response
	ConsultedResource string // Name/URL of the consulted resource

	Errors           interface{}
	ErrorName        string // Error name | e.g., "Resource api/potato not found"
	ErrorCode        string // Error code | e.g., "ERR_NOT_FOUND"
	ErrorDescription string // Error description | e.g., "The requested resource could not be found"
}
