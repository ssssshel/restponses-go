package utils

// Error details
/*
 - Specific response for: 400 Bad Request | 401 Unauthorized | 402 Payment Required | 403 Forbidden | 412 Precondition Failed | 417 Expectation Failed | 421 Misdirected Request | 422 Unprocessable Entity | 425 Unassigned | 426 Upgrade Required | 428 Precondition Required | 429 Too Many Requests | 431 Request Header Fields Too Large | 451 Unavailable For Legal Rasons
 - Specific response for any 5xx status
*/
type ErrorGenericDetail struct {
	Name        string `json:"name"`                  // Error name | e.g., "Resource api/potato not found"
	Code        string `json:"code,omitempty"`        // Error code | e.g., "ERR_NOT_FOUND"
	Description string `json:"description,omitempty"` // Error description | e.g., "The requested resource could not be found"
}

type Error404Detail struct {
	*ErrorGenericDetail
	NotFoundResource string `json:"not_found_resource"` // Name of not found resource | e.g., "api/potato/user"
}

type Error405Detail struct {
	*ErrorGenericDetail
	AllowedMethods []string `json:"allowed_methods"` // List of allowed method | e.g., ["GET", "PUT"]
}

type Error406Detail struct {
	*ErrorGenericDetail
	AllowedRepresentations []string `json:"allowed_representations"` // List of allowed representations | e.g., ["image/png", "application/json"]
}

type Error407Detail struct {
	*ErrorGenericDetail
	AuthenticationType string `json:"authentication_type"` // Auth type | e.g., "Bearer"
	Realm              string `json:"realm"`               // Domain to which auth was requested | e.g., "proxy.potato.com"
}

type Error408Detail struct {
	*ErrorGenericDetail
	TimeWaiting string `json:"time_waiting"` // Time waiting | e.g., "30s"
}

type Error409Detail struct {
	*ErrorGenericDetail
	ConflictResource string `json:"conflict_resource"`     // Conflicting resource name | e.g., "Potato user"
	ConflictId       string `json:"conflict_id,omitempty"` // Personalized conflict ID | e.g., "POT_USER_ERR_H345"
}

type Error410Detail struct {
	*ErrorGenericDetail
	GoneResource string `json:"gone_resource"` // Gone resource name | e.g., "Potato user"
	Reason       string `json:"reason"`        // Reason why the resource is gone | e.g., "Deleted"
}

type Error411Detail struct {
	*ErrorGenericDetail
	RequiredHeader string `json:"required_header"` // Content-Length header: "Content-Lenght"
}

type Error413Detail struct {
	*ErrorGenericDetail
	MaxAllowedSize string `json:"max_allowed_size"` // Max allowed payload size | e.g., "10MB"
}

type Error414Detail struct {
	*ErrorGenericDetail
	MaxAllowedLength uint `json:"max_allowed_length"` // Max allowed URI length | e.g., 3500
}

type Error415Detail struct {
	*ErrorGenericDetail
	SupportedMediaTypes []string `json:"supported_media_types"` // List of supported media types | e.g., ["image/jpg", "audio/*"]
}

type Error416Detail struct {
	*ErrorGenericDetail
	RequestedContentRange string `json:"requested_content_range"` // Range requested by client | e.g., */500
	SupportedContentRange string `json:"supported_content_range"` // Range supported by server | e.g., */100
}

type Error423Detail struct {
	*ErrorGenericDetail
	LockedResource string `json:"locked_resource"` // Locked resource name/URL
}

type Error424Detail struct {
	*ErrorGenericDetail
	FailedDependency string `json:"failed_dependency"` // Failed dependency name
}
