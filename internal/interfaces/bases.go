package interfaces

// Success and Error properties
type SuccessErrorProps struct {
	Success bool `json:"success,omitempty"`
	Error   bool `json:"error,omitempty"`
}

// Base of all responses. Also default response for 1xx state codes
type BaseResponse struct {
	HttpStatus        uint16 `json:"status"`
	ServerMessage     string `json:"server_message"`
	Details           string `json:"details,omitempty"`            // Datailed info of your response
	ConsultedResource string `json:"consulted_resource,omitempty"` // Name/URL of the consulted resource
}

// Default response for 2xx status code. Specific response for 200 Ok, 204 No content, 205 Reset content, 206 Partial content, 208 Already reported and 226 IM used state codes
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

// Default response for 3xx status code. Specific response for 304 Not modified status code
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

// Default response for 4xx and 5xx status code.
/*
Specific response for 400 Bad request, 401 Unauthorized, 402 Payment required, 403 Forbidden, 412 Precondition failed, 417 Expectation failed, 421 Misdirected request, 422 Unprocessable entity, 425 Unassigned, 426 Upgrade required, 428 Precondition required, 429 Too many requests, 431 Request header fields too large and 451 Unavailable for legal reasons state codes*/
type GenericErrorResponse struct {
	BaseResponse
	Errors []ErrorDetail `json:"errors"` // List of errors
	SuccessErrorProps
}

type ErrorDetail struct {
	Name        string `json:"name"`           // Error name e.g., "Resource api/potato not found"
	Code        string `json:"code,omitempty"` // Error code e.g., "ERR_NOT_FOUND"
	Description string `json:"description"`    // Error description
}

// Response for 404 status code
type Status404NotFound_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		NotFoundResource string `json:"not_found_resource"` // Name of not found resource
	}
	SuccessErrorProps
}

// Response for 405 status code
type Status405MethodNotAllowed_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		AllowedMethods []string `json:"allowed_methods"` // List of allowed method
	}
	SuccessErrorProps
}

// Response for 406 status code
type Status406NotAcceptable_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		AllowedRepresentations []string `json:"allowed_representations"` // List of allowed representations
	}
	SuccessErrorProps
}

// Response for 407 status code
type Status407ProxyAuthenticationRequired_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		AuthenticationType string `json:"authentication_type"` // Auth type
		Realm              string `json:"realm"`               // Domain to which auth was requested
	}
	SuccessErrorProps
}

// Response for 408 status code
type Status408RequestTimeout_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		TimeWaiting string `json:"time_waiting"` // Time waiting
	}
	SuccessErrorProps
}

// Response for 409 status code
type Status409Conflict_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		ConflictResource string `json:"conflict_resource"` // Conflicting resource name
		ConflictId       string `json:"conflict_id"`       // Personalized conflict ID
	}
	SuccessErrorProps
}

// Response for 410 status code
type Status410Gone_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		GoneResource string `json:"gone_resource"` // Gone resource name
		Reason       string `json:"reason"`        // Reason why the resource is gone
	}
	SuccessErrorProps
}

// Response for 411 status code
type Status411LengthRequired_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		RequiredHeader string `json:"required_header"` // Content-Length header
	}
	SuccessErrorProps
}

// Response for 413 status code
type Status413PayloadTooLarge_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		MaxAllowedSize string `json:"max_allowed_size"` // Max allowed payload size
	}
	SuccessErrorProps
}

// Response for 414 status code
type Status414RequestUriTooLong_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		MaxAllowedLength uint `json:"max_allowed_length"` // Max allowed URI length | e.g.,
	}
	SuccessErrorProps
}

// Response for 415 status code
type Status415UnsupportedMediaType_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		SupportedMediaTypes []string `json:"supported_media_types"` // List of supported media types
	}
	SuccessErrorProps
}

// Response for 416 status code
type Status416RequestRangeNotSatisfiable_Response struct {
	BaseResponse
	Errors []struct {
		ErrorDetail
		RequestedContentRange string `json:"requested_content_range"` // Range requested by client
		SupportedContentRange string `json:"supported_content_range"` // Range supported by server
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
