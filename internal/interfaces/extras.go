package interfaces

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

// Response for 207 status code
type Status207MultiStatus_Response struct {
	*Status2xx_Response                // The classic response that you love ❤️
	States              []*BasicStatus `json:"states,omitempty"` // Here you can put info about the multiple states (Status, server_message and resource)
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
