package interfaces

import "github.com/ssssshel/restponses-go/internal/states"

// Success and Error properties
type SuccessErrorProps struct {
	Success bool `json:"success,omitempty"`
	Error   bool `json:"error,omitempty"`
}

// Base of all responses
type GeneralInput struct {
	// GENERAL | Default response for 1xx

	HttpStatus        states.StatusCode `json:"status"`
	ServerMessage     string            `json:"serverMessage"`
	Details           string            `json:"details,omitempty"`           // Datailed info of your response
	ConsultedResource string            `json:"consultedResource,omitempty"` // Name/URL of the consulted resource

	//200s

	Data      interface{} `json:"data,omitempty"`      // 2xx ONLY | Anything you want to return, don't use for 204 and 205 states
	Location  string      `json:"location,omitempty"`  // 201 ONLY | URL or place where your creation can be found
	RequestId string      `json:"requestId,omitempty"` // 202 ONLY | ID with wich you can follow the process of your request
	// 203 ONLY | Here you can put info about the third source involved
	Source struct {
		Name        string `json:"name,omitempty"`        // Source name
		Description string `json:"description,omitempty"` // Source description
		Source      string `json:"source,omitempty"`      // Source url
	} `json:"source,omitempty"`
	States []BaseResponse `json:"states,omitempty"` // 207 ONLY | Here you can put info about the multiple states (Status, server_message and resource)

	// 300s

	Options []interface{} `json:"options,omitempty"` // 300 ONLY | Here you can put the options available for the requested resource
	// 301 ONLY
	Sources struct {
		OldSouce  string `json:"oldSource,omitempty"` // Old URL or resource
		NewSource string `json:"newSource,omitempty"` // New URL or resource
	} `json:"sources,omitempty"`
	/*
		Redirect for multiple states:
		 - 302 => Temporary redirect URL
		 - 303 => URL to which the GET request should be made
		 - 307 => Temporary redirect URL that should be queried with the same HTTP method
		 - 308 => Permanent redirect URL that should be queried with tha same HTTP method
	*/
	RedirectUrl string `json:"redirectUrl,omitempty"`
	ProxyUrl    string `json:"proxyUrl,omitempty"` // 305 ONLY | Proxy URL to redirect

	// ERRORS

	Errors []struct {
		// GENERIC | Default specs for 4xx and 5xx

		Name        string `json:"name,omitempty"`        // Error name | e.g., "Resource api/potato not found"
		Code        string `json:"code,omitempty"`        // Error code | e.g., "ERR_NOT_FOUND"
		Description string `json:"description,omitempty"` // Error description | e.g., "The requested resource could not be found"

		// 400s

		NotFoundResource       string   `json:"notFoundResource,omitempty"`       // 404 ONLY | Name of not found resource | e.g., "api/potato/user"
		AllowedMethods         []string `json:"allowedMethods,omitempty"`         // 405 ONLY | List of allowed method | e.g., ["GET", "PUT"]
		AllowedRepresentations []string `json:"allowedRepresentations,omitempty"` // 406 ONLY | List of allowed representations | e.g., ["image/png", "application/json"]
		AuthenticationType     string   `json:"authenticationType,omitempty"`     // 407 ONLY | Auth type | e.g., "Bearer"
		Realm                  string   `json:"realm,omitempty"`                  // 407 ONLY | Domain to which auth was requested | e.g., "proxy.potato.com"
		TimeWaiting            string   `json:"timeWaiting,omitempty"`            // 408 ONLY | Time waiting | e.g., "30s"
		ConflictResource       string   `json:"conflictResource,omitempty"`       // 409 ONLY | Conflicting resource name | e.g., "Potato user"
		ConflictId             string   `json:"conflictId,omitempty"`             // 409 ONLY | Personalized conflict ID | e.g., "POT_USER_ERR_H345"
		GoneResource           string   `json:"goneResource,omitempty"`           // 410 ONLY | Gone resource name | e.g., "Potato user"
		Reason                 string   `json:"reason,omitempty"`                 // 410 ONLY | Reason why the resource is gone | e.g., "Deleted"
		RequiredHeader         string   `json:"requiredHeader,omitempty"`         // 411 ONLY | Content-Length header: "Content-Lenght"
		MaxAllowedSize         string   `json:"maxAllowedSize,omitempty"`         // 413 ONLY | Max allowed payload size | e.g., "10MB"
		MaxAllowedLength       uint     `json:"maxAllowedLength,omitempty"`       // 414 ONLY | Max allowed URI length | e.g., 3500
		SupportedMediaTypes    []string `json:"supportedMediaTypes,omitempty"`    // 415 ONLY | List of supported media types | e.g., ["image/jpg", "audio/*"]
		RequestedContentRange  string   `json:"requestedContentRange,omitempty"`  // 416 ONLY | Range requested by client | e.g., */500
		SupportedContentRange  string   `json:"supportedContentRange,omitempty"`  // 416 ONLY | Range supported by server | e.g., */100
		LockedResource         string   `json:"lockedResource,omitempty"`         // 423 ONLY | Locked resource name/URL
		FailedDependency       string   `json:"failedDependency,omitempty"`       // 424 ONLY | Failed dependency name
	}
}

// 100s

// Base of all responses. Also default response for 1xx state codes
type BaseResponse struct {
	HttpStatus        states.StatusCode `json:"status"`
	ServerMessage     string            `json:"serverMessage"`
	Details           string            `json:"details,omitempty"`           // Datailed info of your response
	ConsultedResource string            `json:"consultedResource,omitempty"` // Name/URL of the consulted resource
}

// 200s

type GenericSuccessfullResponse struct {
	*BaseResponse
	Data interface{} `json:"data"` // Anything you want to return, don't use for 204 and 205 codes

	Location  string         `json:"location,omitempty"`  // 201 ONLY | URL or place where your creation can be found
	RequestId string         `json:"requestId,omitempty"` // 202 ONLY | ID with wich you can follow the process of your request
	Source    *Source        `json:"source,omitempty"`    // 203 ONLY | Here you can put info about the third source involved
	States    []*BasicStatus `json:"states,omitempty"`    // 207 ONLY | Here you can put info about the multiple states (Status, server_message and resource)

	*SuccessErrorProps
}

type Source struct {
	Name        string `json:"name,omitempty"`        // Source name
	Description string `json:"description,omitempty"` // Source description
	Source      string `json:"source,omitempty"`      // Source url
}

type BasicStatus struct {
	HttpStatus    states.StatusCode `json:"status"`
	ServerMessage string            `json:"serverMessage"`
	Details       string            `json:"details,omitempty"`
}

// 300s

type GenericRedirectionResponse struct {
	*BaseResponse

	Options []interface{} `json:"options,omitempty"` // 300 ONLY | Here you can put the options available for the requested resource
	Sources *Sources      `json:"sources,omitempty"` // 301 ONLY
	/*
		Redirect for multiple states:
		 - 302 => Temporary redirect URL
		 - 303 => URL to which the GET request should be made
		 - 307 => Temporary redirect URL that should be queried with the same HTTP method
		 - 308 => Permanent redirect URL that should be queried with tha same HTTP method
	*/
	RedirectUrl string `json:"redirectUrl,omitempty"`
	ProxyUrl    string `json:"proxyUrl,omitempty"` // 305 ONLY | Proxy URL to redirect
}

type Sources struct {
	OldSouce  string `json:"oldSource,omitempty"` // Old URL or resource
	NewSource string `json:"newSource,omitempty"` // New URL or resource
}

// 400s 500s

/*
# Default response for 4xx and 5xx status code.
*/
type GenericErrorResponse struct {
	*BaseResponse
	ErrorDetails *ErrorDetails
	Errors       interface{} `json:"errors,omitempty"` // List []*ErrorDetail or representation of errors
	*SuccessErrorProps
}

type ErrorDetails struct {
	// GENERIC | Default specs for 4xx and 5xx

	Name        string `json:"name,omitempty"`        // Error name | e.g., "Resource api/potato not found"
	Code        string `json:"code,omitempty"`        // Error code | e.g., "ERR_NOT_FOUND"
	Description string `json:"description,omitempty"` // Error description | e.g., "The requested resource could not be found"

	// 400s

	NotFoundResource       string   `json:"notFoundResource,omitempty"`       // 404 ONLY | Name of not found resource | e.g., "api/potato/user"
	AllowedMethods         []string `json:"allowedMethods,omitempty"`         // 405 ONLY | List of allowed method | e.g., ["GET", "PUT"]
	AllowedRepresentations []string `json:"allowedRepresentations,omitempty"` // 406 ONLY | List of allowed representations | e.g., ["image/png", "application/json"]
	AuthenticationType     string   `json:"authenticationType,omitempty"`     // 407 ONLY | Auth type | e.g., "Bearer"
	Realm                  string   `json:"realm,omitempty"`                  // 407 ONLY | Domain to which auth was requested | e.g., "proxy.potato.com"
	TimeWaiting            string   `json:"timeWaiting,omitempty"`            // 408 ONLY | Time waiting | e.g., "30s"
	ConflictResource       string   `json:"conflictResource,omitempty"`       // 409 ONLY | Conflicting resource name | e.g., "Potato user"
	ConflictId             string   `json:"conflictId,omitempty"`             // 409 ONLY | Personalized conflict ID | e.g., "POT_USER_ERR_H345"
	GoneResource           string   `json:"goneResource,omitempty"`           // 410 ONLY | Gone resource name | e.g., "Potato user"
	Reason                 string   `json:"reason,omitempty"`                 // 410 ONLY | Reason why the resource is gone | e.g., "Deleted"
	RequiredHeader         string   `json:"requiredHeader,omitempty"`         // 411 ONLY | Content-Length header: "Content-Length"
	MaxAllowedSize         string   `json:"maxAllowedSize,omitempty"`         // 413 ONLY | Max allowed payload size | e.g., "10MB"
	MaxAllowedLength       uint     `json:"maxAllowedLength,omitempty"`       // 414 ONLY | Max allowed URI length | e.g., 3500
	SupportedMediaTypes    []string `json:"supportedMediaTypes,omitempty"`    // 415 ONLY | List of supported media types | e.g., ["image/jpg", "audio/*"]
	RequestedContentRange  string   `json:"requestedContentRange,omitempty"`  // 416 ONLY | Range requested by client | e.g., */500
	SupportedContentRange  string   `json:"supportedContentRange,omitempty"`  // 416 ONLY | Range supported by server | e.g., */100
	LockedResource         string   `json:"lockedResource,omitempty"`         // 423 ONLY | Locked resource name/URL
	FailedDependency       string   `json:"failedDependency,omitempty"`       // 424 ONLY | Failed dependency name
}
