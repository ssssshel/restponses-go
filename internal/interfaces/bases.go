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

// Base of all responses. Also default response for 1xx state codes
type BaseResponse struct {
	HttpStatus        states.StatusCode `json:"status"`
	ServerMessage     string            `json:"serverMessage"`
	Details           string            `json:"details,omitempty"`           // Datailed info of your response
	ConsultedResource string            `json:"consultedResource,omitempty"` // Name/URL of the consulted resource
}

/*
# Default response for 2xx status code.
 - Specific response for: 200 Ok | 204 No Content | 205 Reset Content | 206 Partial Content | 208 Already Reported | 226 IM Used
*/
type Status2xx_Response struct {
	*BaseResponse
	Data interface{} `json:"data"` // Anything you want to return, don't use for 204 and 205 codes
	*SuccessErrorProps
}

// Response for 201 status code
type Status201Created_Response struct {
	*Status2xx_Response        // The classic response that you love ❤️
	Location            string `json:"location,omitempty"` // URL or place where your creation can be found
}

// Response for 202 status code
type Status202Accepted_Response struct {
	*Status2xx_Response        // The classic response that you love ❤️
	RequestId           string `json:"requestId,omitempty"` // ID with wich you can follow the process of your request
}

// Response for 203 status code
type Status203NonAI_Response struct {
	*Status2xx_Response // The classic response that you love ❤️
	// Here you can put info about the third source involved
	Source *Source `json:"source,omitempty"`
}

type Source struct {
	Name        string `json:"name,omitempty"`        // Source name
	Description string `json:"description,omitempty"` // Source description
	Source      string `json:"source,omitempty"`      // Source url
}

// Response for 207 status code
type Status207MultiStatus_Response struct {
	*Status2xx_Response                 // The classic response that you love ❤️
	States              []*BaseResponse `json:"states,omitempty"` // Here you can put info about the multiple states (Status, server_message and resource)
}

/*
# Default response for 3xx status code.
 - Specific response for: 304 Not modified
*/
type Status3xx_Response struct {
	*BaseResponse
	// SuccessErrorProps
}

// Response for 300 status code
type Status300MultipleChoices_Response struct {
	*Status3xx_Response
	Options []interface{} `json:"options,omitempty"` // Here you can put the options available for the requested resource
}

// Response for 301 status code
type Status301MovedPermanently_Response struct {
	*Status3xx_Response
	Sources *Sources `json:"sources,omitempty"`
}

type Sources struct {
	OldSouce  string `json:"oldSource,omitempty"` // Old URL or resource
	NewSource string `json:"newSource,omitempty"` // New URL or resource
}

// Response for 302 status code
type Status302Found_Response struct {
	*Status3xx_Response
	RedirectUrl string `json:"redirectUrl,omitempty"` // Temporary redirect URL
}

// Response for 303 status code
type Status303SeeOther_Response struct {
	*Status3xx_Response
	RedirectUrl string `json:"redirectUrl"` // URL to which the GET request should be made
}

// Response for 305 status code
type Status305UseProxy_Response struct {
	*Status3xx_Response
	ProxyUrl string `json:"proxyUrl"` // Proxy URL to redirect
}

// Response for 307 status code
type Status307TemporaryRedirect_Response struct {
	*Status3xx_Response
	RedirectUrl string `json:"redirectUrl"` // Temporary redirect URL that should be queried with the same HTTP method
}

// Response for 308 status code
type Status308PermanentRedirect_Response struct {
	*Status3xx_Response
	RedirectUrl string `json:"redirectUrl"` // Permanent redirect URL that should be queried with tha same HTTP method
}

/*
# Default response for 4xx and 5xx status code.
*/
type GenericErrorResponse struct {
	*BaseResponse
	Errors []interface{} `json:"errors"` // List []*ErrorDetail or representation of errors
	*SuccessErrorProps
}
