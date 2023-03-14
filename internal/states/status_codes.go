package states

type StatusCode uint16

type HttpStatus struct {
	Code    StatusCode
	Message string
	Details string
}

const (
	Status100 StatusCode = 100
	Status101 StatusCode = 101
	Status102 StatusCode = 102
	Status103 StatusCode = 103

	Status200 StatusCode = 200
	Status201 StatusCode = 201
	Status202 StatusCode = 202
	Status203 StatusCode = 203
	Status204 StatusCode = 204
	Status205 StatusCode = 205
	Status206 StatusCode = 206
	Status207 StatusCode = 207
	Status208 StatusCode = 208
	Status226 StatusCode = 226

	Status300 StatusCode = 300
	Status301 StatusCode = 301
	Status302 StatusCode = 302
	Status303 StatusCode = 303
	Status304 StatusCode = 304
	Status305 StatusCode = 305
	Status307 StatusCode = 307
	Status308 StatusCode = 308

	Status400 StatusCode = 400
	Status401 StatusCode = 401
	Status402 StatusCode = 402
	Status403 StatusCode = 403
	Status404 StatusCode = 404
	Status405 StatusCode = 405
	Status406 StatusCode = 406
	Status407 StatusCode = 407
	Status408 StatusCode = 408
	Status409 StatusCode = 409
	Status410 StatusCode = 410
	Status411 StatusCode = 411
	Status412 StatusCode = 412
	Status413 StatusCode = 413
	Status414 StatusCode = 414
	Status415 StatusCode = 415
	Status416 StatusCode = 416
	Status417 StatusCode = 417
	Status418 StatusCode = 418
	Status421 StatusCode = 421
	Status422 StatusCode = 422
	Status423 StatusCode = 423
	Status424 StatusCode = 424
	Status425 StatusCode = 425
	Status426 StatusCode = 426
	Status428 StatusCode = 428
	Status429 StatusCode = 429
	Status431 StatusCode = 431
	Status451 StatusCode = 451

	Status500 StatusCode = 500
	Status501 StatusCode = 501
	Status502 StatusCode = 502
	Status503 StatusCode = 503
	Status504 StatusCode = 504
	Status505 StatusCode = 505
	Status506 StatusCode = 506
	Status507 StatusCode = 507
	Status508 StatusCode = 508
	Status509 StatusCode = 509
	Status510 StatusCode = 510
	Status511 StatusCode = 511
	Status521 StatusCode = 521
)

const (
	Status100Continue = iota
	Status101SwitchingProtocols
	Status102Processing
	Status103Checkpoint

	Status200Ok
	Status201Created
	Status202Accepted
	Status203NonAuthoritativeInformation
	Status204NoContent
	Status205ResetContent
	Status206PartialContent
	Status207MultiStatus
	Status208AlreadyReported
	Status226IMUsed

	Status300MultipleChoices
	Status301MovedPermanently
	Status302Found
	Status303SeeOther
	Status304NotModified
	Status305UseProxy
	Status307TemporaryRedirect
	Status308PermanentRedirect

	Status400BadRequest
	Status401Unauthorized
	Status402PaymentRequired
	Status403Forbidden
	Status404NotFound
	Status405MethodNotAllowed
	Status406NotAcceptable
	Status407ProxyAuthenticationRequired
	Status408RequestTimeout
	Status409Conflict
	Status410Gone
	Status411LengthRequired
	Status412PreconditionFailed
	Status413PayloadTooLarge
	Status414RequestUriTooLong
	Status415UnsupportedMediaType
	Status416RequestRangeNotSatisfiable
	Status417ExpectationFailed
	Status418Teapot
	Status421MisdirectedRequest
	Status422UnprocessableEntity
	Status423Locked
	Status424FailedDependency
	Status425Unassigned
	Status426UpgradeRequired
	Status428PreconditionRequired
	Status429TooManyRequests
	Status431RequestHeaderFieldsTooLarge
	Status451UnavailableForLegalReasons

	Status500InternalServerError
	Status501NotImplemented
	Status502BadGateway
	Status503ServiceUnavailable
	Status504GatewayTimeout
	Status505HTTPVersionNotSupported
	Status506VariantAlsoNegotiates
	Status507InsufficientStorage
	Status508LoopDetected
	Status509BandwidthLimitExceeded
	Status510NotExtended
	Status511NetworkAuthenticationRequired
	Status521WebServerIsDown
)

var DefaultStatesContent = [...]HttpStatus{
	Status100Continue:           {Status100, "Continue", "Continue with the request"},
	Status101SwitchingProtocols: {Status101, "Switching Protocols", ""},
	Status102Processing:         {Status102, "Processing", ""},
	Status103Checkpoint:         {Status103, "Checkpoint", ""},

	Status200Ok:                          {Status200, "Ok", "The request has been successfully processed"},
	Status201Created:                     {Status201, "Created", "Resource successfully created"},
	Status202Accepted:                    {Status202, "Accepted", "Request accepted for processing"},
	Status203NonAuthoritativeInformation: {Status203, "Non-authoritative information", "Non-authoritative information returned"},
	Status204NoContent:                   {Status204, "No content", "No content to return"},
	Status205ResetContent:                {Status205, "Reset Content", "The request has been successfully processed but there is no content to return"},
	Status206PartialContent:              {Status206, "Partial Content", "Partial content returned"},
	Status207MultiStatus:                 {Status207, "Multi-Status", "Multi-status response"},
	Status208AlreadyReported:             {Status208, "Already Reported", "Resource already reported"},
	Status226IMUsed:                      {Status226, "IM Used", ""},

	Status300MultipleChoices:   {Status300, "Multiple Choices", ""},
	Status301MovedPermanently:  {Status301, "Moved Permanently", ""},
	Status302Found:             {Status302, "Found", ""},
	Status303SeeOther:          {Status303, "See other", ""},
	Status304NotModified:       {Status304, "Not Modified", ""},
	Status305UseProxy:          {Status305, "Use Proxy", ""},
	Status307TemporaryRedirect: {Status307, "Temporary Redirect", ""},
	Status308PermanentRedirect: {Status308, "Permanent Redirect", ""},

	Status400BadRequest:                  {Status400, "Bad Request", ""},
	Status401Unauthorized:                {Status401, "Unauthorized", ""},
	Status402PaymentRequired:             {Status402, "Payment Required", ""},
	Status403Forbidden:                   {Status403, "Forbidden", ""},
	Status404NotFound:                    {Status404, "Not Found", ""},
	Status405MethodNotAllowed:            {Status405, "Method Not Allowed", ""},
	Status406NotAcceptable:               {Status406, "Not Acceptable", ""},
	Status407ProxyAuthenticationRequired: {Status407, "Proxy Authentication Required", ""},
	Status408RequestTimeout:              {Status408, "Request Timeout", ""},
	Status409Conflict:                    {Status409, "Conflict", ""},
	Status410Gone:                        {Status410, "Gone", ""},
	Status411LengthRequired:              {Status411, "Length Required", ""},
	Status412PreconditionFailed:          {Status412, "Precondition Failed", ""},
	Status413PayloadTooLarge:             {Status413, "Payload Too Large", ""},
	Status414RequestUriTooLong:           {Status414, "Request-URI Too Long", ""},
	Status415UnsupportedMediaType:        {Status415, "Unsupported Media Type", ""},
	Status416RequestRangeNotSatisfiable:  {Status416, "Range Not Satisfiable", ""},
	Status417ExpectationFailed:           {Status417, "Expectation Failed", ""},
	Status418Teapot:                      {Status418, "I'm a Teapot", ""},
	Status421MisdirectedRequest:          {Status421, "Misdirected Request", ""},
	Status422UnprocessableEntity:         {Status422, "Unprocessable Entity", ""},
	Status423Locked:                      {Status423, "Locked", ""},
	Status424FailedDependency:            {Status424, "Failed Dependency", ""},
	Status425Unassigned:                  {Status425, "Too Early", ""},
	Status426UpgradeRequired:             {Status426, "Upgrade Required", ""},
	Status428PreconditionRequired:        {Status428, "Precondition Required", ""},
	Status429TooManyRequests:             {Status429, "Too Many Requests", ""},
	Status431RequestHeaderFieldsTooLarge: {Status431, "Request Header Fields Too Large", ""},
	Status451UnavailableForLegalReasons:  {Status451, "Unavailable For Legal Reasons", ""},

	Status500InternalServerError:           {Status500, "Internal Server Error", ""},
	Status501NotImplemented:                {Status501, "Not Implemented", ""},
	Status502BadGateway:                    {Status502, "Bad Gateway", ""},
	Status503ServiceUnavailable:            {Status503, "Service Unavailable", ""},
	Status504GatewayTimeout:                {Status504, "Gateway Timeout", ""},
	Status505HTTPVersionNotSupported:       {Status505, "HTTP Version Not Supported", ""},
	Status506VariantAlsoNegotiates:         {Status506, "Variant Also Negotiates", ""},
	Status507InsufficientStorage:           {Status507, "Insufficient Storage", ""},
	Status508LoopDetected:                  {Status508, "Loop Detected", ""},
	Status509BandwidthLimitExceeded:        {Status509, "Bandwidth Limit Exceeded", ""},
	Status510NotExtended:                   {Status510, "Not Extended", ""},
	Status511NetworkAuthenticationRequired: {Status511, "Network Authentication Required", ""},
	Status521WebServerIsDown:               {Status521, "Web Server Is Down", ""},
}
