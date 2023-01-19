package interfaces

import "github.com/ssssshel/restponses-go/internal/states"

// Success and Error properties
type SuccessErrorProps struct {
	Success bool `json:"success,omitempty"`
	Error   bool `json:"error,omitempty"`
}

type ErrorDetail struct {
	Name        string `json:"name"`                  // Error name | e.g., "Resource api/potato not found"
	Code        string `json:"code,omitempty"`        // Error code | e.g., "ERR_NOT_FOUND"
	Description string `json:"description,omitempty"` // Error description | e.g., "The requested resource could not be found"
}

// Base of all responses
type BaseResponse struct {

	// GENERAL | Default response for 1xx and 2xx

	HttpStatus        states.StatusCode `json:"status"`
	ServerMessage     string            `json:"server_message"`
	Details           string            `json:"details,omitempty"`            // Datailed info of your response
	ConsultedResource string            `json:"consulted_resource,omitempty"` // Name/URL of the consulted resource

	//200s

	Data      interface{} `json:"data,omitempty"`       // 2xx ONLY | Anything you want to return, don't use for 204 and 205 states
	Location  string      `json:"location,omitempty"`   // 201 ONLY | URL or place where your creation can be found
	RequestId string      `json:"request_id,omitempty"` // 202 ONLY | ID with wich you can follow the process of your request
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
		OldSouce  string `json:"old_source,omitempty"` // Old URL or resource
		NewSource string `json:"new_source,omitempty"` // New URL or resource
	} `json:"sources,omitempty"`
	/*
		Redirect for multiple states:
		 - 302 => Temporary redirect URL
		 - 303 => URL to which the GET request should be made
		 - 307 => Temporary redirect URL that should be queried with the same HTTP method
		 - 308 => Permanent redirect URL that should be queried with tha same HTTP method
	*/
	RedirectUrl string `json:"redirect_url,omitempty"`
	ProxyUrl    string `json:"proxy_url,omitempty"` // 305 ONLY | Proxy URL to redirect

	// ERRORS

	Errors []struct {
		// GENERIC | Default specs for 4xx and 5xx

		Name        string `json:"name,omitempty"`        // Error name | e.g., "Resource api/potato not found"
		Code        string `json:"code,omitempty"`        // Error code | e.g., "ERR_NOT_FOUND"
		Description string `json:"description,omitempty"` // Error description | e.g., "The requested resource could not be found"

		// 400s

		NotFoundResource       string   `json:"not_found_resource"`      // 404 ONLY | Name of not found resource | e.g., "api/potato/user"
		AllowedMethods         []string `json:"allowed_methods"`         // 405 ONLY | List of allowed method | e.g., ["GET", "PUT"]
		AllowedRepresentations []string `json:"allowed_representations"` // 406 ONLY | List of allowed representations | e.g., ["image/png", "application/json"]
		AuthenticationType     string   `json:"authentication_type"`     // 407 ONLY | Auth type | e.g., "Bearer"
		Realm                  string   `json:"realm"`                   // 407 ONLY | Domain to which auth was requested | e.g., "proxy.potato.com"
		TimeWaiting            string   `json:"time_waiting"`            // 408 ONLY | Time waiting | e.g., "30s"
		ConflictResource       string   `json:"conflict_resource"`       // 409 ONLY | Conflicting resource name | e.g., "Potato user"
		ConflictId             string   `json:"conflict_id,omitempty"`   // 409 ONLY | Personalized conflict ID | e.g., "POT_USER_ERR_H345"
		GoneResource           string   `json:"gone_resource"`           // 410 ONLY | Gone resource name | e.g., "Potato user"
		Reason                 string   `json:"reason"`                  // 410 ONLY | Reason why the resource is gone | e.g., "Deleted"
		RequiredHeader         string   `json:"required_header"`         // 411 ONLY | Content-Length header: "Content-Lenght"
		MaxAllowedSize         string   `json:"max_allowed_size"`        // 413 ONLY | Max allowed payload size | e.g., "10MB"

	} `json:"errors,omitempty"` // 4xx and 5xx ONLY

	SuccessErrorProps
}

/*
# Default response for 2xx status code.
 - Specific response for: 200 Ok | 204 No Content | 205 Reset Content | 206 Partial Content | 208 Already Reported | 226 IM Used
*/
type Status2xx_Response struct {
	BaseResponse
	Data interface{} `json:"data,omitempty"` // Anything you want to return, don't use for 204 and 205 codes
	SuccessErrorProps
}

// Response for 201 status code
type Status201Created_Response struct {
	Status2xx_Response        // The classic response that you love ❤️
	Location           string `json:"location"` // URL or place where your creation can be found
}

// Response for 202 status code
type Status202Accepted_Response struct {
	Status2xx_Response        // The classic response that you love ❤️
	RequestId          string `json:"request_id"` // ID with wich you can follow the process of your request
}

// Response for 203 status code
type Status203NonAI_Response struct {
	Status2xx_Response // The classic response that you love ❤️

	// Here you can put info about the third source involved
	Source struct {
		Name        string `json:"name"`        // Source name
		Description string `json:"description"` // Source description
		Source      string `json:"source"`      // Source url
	} `json:"source"`
}

// Response for 207 status code
type Status207MultiStatus_Response struct {
	Status2xx_Response                // The classic response that you love ❤️
	States             []BaseResponse `json:"states"` // Here you can put info about the multiple states (Status, server_message and resource)
}

/*
# Default response for 3xx status code.
 - Specific response for: 304 Not modified
*/
type Status3xx_Response struct {
	BaseResponse
	// SuccessErrorProps
}

// Response for 300 status code
type Status300MultipleChoices_Response struct {
	Status3xx_Response
	Options []interface{} `json:"options"` // Here you can put the options available for the requested resource
}

// Response for 301 status code
type Status301MovedPermanently_Response struct {
	Status3xx_Response
	Sources struct {
		OldSouce  string `json:"old_source"` // Old URL or resource
		NewSource string `json:"new_source"` // New URL or resource
	} `json:"sources"`
}

// Response for 302 status code
type Status302Found_Response struct {
	Status3xx_Response
	RedirectUrl string `json:"redirect_url"` // Temporary redirect URL
}

// Response for 303 status code
type Status303SeeOther_Response struct {
	Status3xx_Response
	RedirectUrl string `json:"redirect_url"` // URL to which the GET request should be made
}

// Response for 305 status code
type Status305UseProxy_Response struct {
	Status3xx_Response
	ProxyUrl string `json:"proxy_url"` // Proxy URL to redirect
}

// Response for 307 status code
type Status307TemporaryRedirect_Response struct {
	Status3xx_Response
	RedirectUrl string `json:"redirect_url"` // Temporary redirect URL that should be queried with the same HTTP method
}

// Response for 308 status code
type Status308PermanentRedirect_Response struct {
	Status3xx_Response
	RedirectUrl string `json:"redirect_url"` // Permanent redirect URL that should be queried with tha same HTTP method
}

/*
# Default response for 4xx and 5xx status code.
 - Specific response for: 400 Bad Request | 401 Unauthorized | 402 Payment Required | 403 Forbidden | 412 Precondition Failed | 417 Expectation Failed | 421 Misdirected Request | 422 Unprocessable Entity | 425 Unassigned | 426 Upgrade Required | 428 Precondition Required | 429 Too Many Requests | 431 Request Header Fields Too Large | 451 Unavailable For Legal Rasons
 - Specific response for any 5xx status
*/
type GenericErrorResponse struct {
	BaseResponse
	Errors []ErrorDetail `json:"errors"` // List of errors
	SuccessErrorProps
}

// Response for 404 status code
type Status404NotFound_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		NotFoundResource string `json:"not_found_resource"` // Name of not found resource | e.g., "api/potato/user"
	}
	SuccessErrorProps
}

// Response for 405 status code
type Status405MethodNotAllowed_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		AllowedMethods []string `json:"allowed_methods"` // List of allowed method | e.g., ["GET", "PUT"]
	}
	SuccessErrorProps
}

// Response for 406 status code
type Status406NotAcceptable_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		AllowedRepresentations []string `json:"allowed_representations"` // List of allowed representations | e.g., ["image/png", "application/json"]
	}
	SuccessErrorProps
}

// Response for 407 status code
type Status407ProxyAuthenticationRequired_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		AuthenticationType string `json:"authentication_type"` // Auth type | e.g., "Bearer"
		Realm              string `json:"realm"`               // Domain to which auth was requested | e.g., "proxy.potato.com"
	}
	SuccessErrorProps
}

// Response for 408 status code
type Status408RequestTimeout_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		TimeWaiting string `json:"time_waiting"` // Time waiting | e.g., "30s"
	}
	SuccessErrorProps
}

// Response for 409 status code
type Status409Conflict_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		ConflictResource string `json:"conflict_resource"`     // Conflicting resource name | e.g., "Potato user"
		ConflictId       string `json:"conflict_id,omitempty"` // Personalized conflict ID | e.g., "POT_USER_ERR_H345"
	}
	SuccessErrorProps
}

// Response for 410 status code
type Status410Gone_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		GoneResource string `json:"gone_resource"` // Gone resource name | e.g., "Potato user"
		Reason       string `json:"reason"`        // Reason why the resource is gone | e.g., "Deleted"
	}
	SuccessErrorProps
}

// Response for 411 status code
type Status411LengthRequired_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		RequiredHeader string `json:"required_header"` // Content-Length header: "Content-Lenght"
	}
	SuccessErrorProps
}

// Response for 413 status code
type Status413PayloadTooLarge_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		MaxAllowedSize string `json:"max_allowed_size"` // Max allowed payload size | e.g., "10MB"
	}
	SuccessErrorProps
}

// Response for 414 status code
type Status414RequestUriTooLong_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		MaxAllowedLength uint `json:"max_allowed_length"` // Max allowed URI length | e.g., 3500
	}
	SuccessErrorProps
}

// Response for 415 status code
type Status415UnsupportedMediaType_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		SupportedMediaTypes []string `json:"supported_media_types"` // List of supported media types | e.g., ["image/jpg", "audio/*"]
	}
	SuccessErrorProps
}

// Response for 416 status code
type Status416RequestRangeNotSatisfiable_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		RequestedContentRange string `json:"requested_content_range"` // Range requested by client | e.g., */500
		SupportedContentRange string `json:"supported_content_range"` // Range supported by server | e.g., */100
	}
	SuccessErrorProps
}

// Response for 423 status code
type Status423Locked_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		LockedResource string `json:"locked_resource"` // Locked resource name/URL
	}
	SuccessErrorProps
}

// Response for 424 status code
type Status424FailedDependency struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		FailedDependency string `json:"failed_dependency"` // Failed dependency name
	}
	SuccessErrorProps
}
