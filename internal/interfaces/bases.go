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
	Status2xx_Response        // The classic response that you love
	Location           string `json:"location"` // URL or place where your creation can be found
}

// Response for 202 status code
type Status202Accepted_Response struct {
	Status2xx_Response        // The classic response that you love
	RequestId          string `json:"request_id"` // ID with wich you can follow the process of your request
}

// Response for 203 status code
type Status203NonAI_Response struct {
	Status2xx_Response // The classic response that you love

	// Here you can put info about the third source involved
	Source struct {
		Name        string `json:"name"`        // Source name
		Description string `json:"description"` // Source description
		Source      string `json:"source"`      // Source url
	} `json:"source"`
}

// Response for 207 status code
type Status207MultiStatus_Response struct {
	Status2xx_Response                // The classic response that you love
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

// Default response for 4xx status code. Specific response for 400 Bad request, 401 Unauthorized, 402 Payment required and 403 Forbidden state codes
type GenericErrorResponse struct {
	BaseResponse
	Errors []interface{} `json:"errorMessage"`
}

type Status404NotFound_Response struct {
	GenericErrorResponse
	NotFoundResource string // Name of not found resource
}

type Status405MethodNotAllowed_Response struct {
	GenericErrorResponse
	AllowedMethods []string // List of allowed method
}

type Status406NotAcceptable_Response struct {
	GenericErrorResponse
	AllowedRepresentations []string // List of allowed representations
}

type Status407ProxyAuthenticationRequired_Response struct {
	GenericErrorResponse
	AuthenticationType string // Auth type
	Realm              string // Domain to which auth is requested
}

type Status408RequestTimeout_Response struct {
	GenericErrorResponse
	TimeWaiting string // Time waiting
}

type Status409Conflict_Response struct {
	GenericErrorResponse
	ConflictResource string // Conflicting resource name
	ConflictId       string // Personalized conflict ID
}

type Status410Gone_Response struct {
	GenericErrorResponse
	GoneResource string // Gone resource name
	Reason       string // Reason why the resource is gone
}

type Status411LengthRequired_Response struct {
	GenericErrorResponse
	RequiredHeaders string // Content-Length header
}
