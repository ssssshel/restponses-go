package restponses

type StatusCode1xx uint16
type StatusCode2xx uint16
type StatusCode3xx uint16
type StatusCode4xx uint16
type StatusCode5xx uint16

const (
	Status100Continue           StatusCode1xx = 100
	Status101SwitchingProtocols StatusCode1xx = 101
	Status102Processing         StatusCode1xx = 102
	Status103EarlyHints         StatusCode1xx = 103
)

const (
	Status200Ok                   StatusCode2xx = 200
	Status201Continue             StatusCode2xx = 201
	Status202Accepted             StatusCode2xx = 202
	Status203NonAuthoritativeInfo StatusCode2xx = 203
	Status204NoContent            StatusCode2xx = 204
	Status205ResetContent         StatusCode2xx = 205
	Status206PartialContent       StatusCode2xx = 206
	Status207MultiStatus          StatusCode2xx = 207
	Status208AlreadyReported      StatusCode2xx = 208
	Status226IMUsed               StatusCode2xx = 226
)

const (
	Status300MultipleChoices   StatusCode3xx = 300
	Status301MovedPermanently  StatusCode3xx = 301
	Status302Found             StatusCode3xx = 302
	Status303SeeOther          StatusCode3xx = 303
	Status304NotModified       StatusCode3xx = 304
	Status305UseProxy          StatusCode3xx = 305
	Status307TemporaryRedirect StatusCode3xx = 307
	Status308PermanentRedirect StatusCode3xx = 308
)

const (
	Status400BadRequest                  StatusCode4xx = 400
	Status401Unauthorized                StatusCode4xx = 401
	Status402PaymentRequired             StatusCode4xx = 402
	Status403Forbidden                   StatusCode4xx = 403
	Status404NotFound                    StatusCode4xx = 404
	Status405MethodNotAllowed            StatusCode4xx = 405
	Status406NotAcceptable               StatusCode4xx = 406
	Status407ProxyAuthenticationRequired StatusCode4xx = 407
	Status408RequestTimeout              StatusCode4xx = 408
	Status409Conflict                    StatusCode4xx = 409
	Status410Gone                        StatusCode4xx = 410
	Status411LengthRequired              StatusCode4xx = 411
	Status412PreconditionFailed          StatusCode4xx = 412
	Status413PayloadTooLarge             StatusCode4xx = 413
	Status414RequestUriTooLong           StatusCode4xx = 414
	Status415UnsupportedMediaType        StatusCode4xx = 415
	Status416RequestRangeNotSatisfiable  StatusCode4xx = 416
	Status417ExpectationFailed           StatusCode4xx = 417
	Status418Teapot                      StatusCode4xx = 418
	Status421MisdirectedRequest          StatusCode4xx = 421
	Status422UnprocessableEntity         StatusCode4xx = 422
	Status423Locked                      StatusCode4xx = 423
	Status424FailedDependency            StatusCode4xx = 424
	Status425Unassigned                  StatusCode4xx = 425
	Status426UpgradeRequired             StatusCode4xx = 426
	Status428PreconditionRequired        StatusCode4xx = 428
	Status429TooManyRequests             StatusCode4xx = 429
	Status431RequestHeaderFieldsTooLarge StatusCode4xx = 431
	Status451UnavailableForLegalReasons  StatusCode4xx = 451
)

const (
	Status500InternalServerError           StatusCode5xx = 500
	Status501NotImplemented                StatusCode5xx = 501
	Status502BadGateway                    StatusCode5xx = 502
	Status503ServiceUnavailable            StatusCode5xx = 503
	Status504GatewayTimeout                StatusCode5xx = 504
	Status505HttpVersionNotSupported       StatusCode5xx = 505
	Status506VariantAlsoNegotiates         StatusCode5xx = 506
	Status507InsufficientStorage           StatusCode5xx = 507
	Status508LoopDetected                  StatusCode5xx = 508
	Status509Unassigned                    StatusCode5xx = 509
	Status510NotExtended                   StatusCode5xx = 510
	Status511NetworkAuthenticationRequired StatusCode5xx = 511
	Status521WebServerIsDown               StatusCode5xx = 521
)
