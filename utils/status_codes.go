package utils

type HttpStatus uint

const (
	Status100Continue           HttpStatus = 100
	Status101SwitchingProtocols HttpStatus = 101
	Status102Processing         HttpStatus = 102
	Status103Checkpoint         HttpStatus = 103

	Status200Ok                          HttpStatus = 200
	Status201Created                     HttpStatus = 201
	Status202Accepted                    HttpStatus = 202
	Status203NonAuthoritativeInformation HttpStatus = 203
	Status204NoContent                   HttpStatus = 204
	Status205ResetContent                HttpStatus = 205
	Status206PartialContent              HttpStatus = 206
	Status207MultiStatus                 HttpStatus = 207
	Status208AlreadyReported             HttpStatus = 208
	Status226IMUsed                      HttpStatus = 226

	Status300MultipleChoices   HttpStatus = 300
	Status301MovedPermanently  HttpStatus = 301
	Status302Found             HttpStatus = 302
	Status303SeeOther          HttpStatus = 303
	Status304NotModified       HttpStatus = 304
	Status305UseProxy          HttpStatus = 305
	Status307TemporaryRedirect HttpStatus = 307
	Status308PermanentRedirect HttpStatus = 309

	Status400BadRequest                  HttpStatus = 400
	Status401Unauthorized                HttpStatus = 401
	Status402PaymentRequired             HttpStatus = 402
	Status403Forbidden                   HttpStatus = 403
	Status404NotFound                    HttpStatus = 404
	Status405MethodNotAllowed            HttpStatus = 405
	Status406NotAcceptable               HttpStatus = 406
	Status407ProxyAuthenticationRequired HttpStatus = 407
	Status408RequestTimeout              HttpStatus = 408
	Status409Conflict                    HttpStatus = 409
	Status410Gone                        HttpStatus = 410
	Status411LengthRequired              HttpStatus = 411
	Status412PreconditionFailed          HttpStatus = 412
	Status413PayloadTooLarge             HttpStatus = 413
	Status414RequestUriTooLong           HttpStatus = 414
	Status415UnsupportedMediaType        HttpStatus = 415
	Status416RequestRangeNotSatisfiable  HttpStatus = 416
	Status417ExpectationFailed           HttpStatus = 417
	Status418Teapot                      HttpStatus = 418
	Status421MisdirectedRequest          HttpStatus = 421
	Status422UnprocessableEntity         HttpStatus = 422
	Status423Locked                      HttpStatus = 423
	Status424FailedDependency            HttpStatus = 424
	Status425Unassigned                  HttpStatus = 425
	Status426UpgradeRequired             HttpStatus = 426
	Status428PreconditionRequired        HttpStatus = 428
	Status429TooManyRequests             HttpStatus = 429
	Status431RequestHeaderFieldsTooLarge HttpStatus = 431
	Status451UnavailableForLegalReasons  HttpStatus = 451

	Status500InternalServerError           HttpStatus = 500
	Status501NotImplemented                HttpStatus = 501
	Status502BadGateway                    HttpStatus = 502
	Status503ServiceUnavailable            HttpStatus = 503
	Status504GatewayTimeout                HttpStatus = 504
	Status505HTTPVersionNotSupported       HttpStatus = 505
	Status506VariantAlsoNegotiates         HttpStatus = 506
	Status507InsufficientStorage           HttpStatus = 507
	Status508LoopDetected                  HttpStatus = 508
	Status509BandwidthLimitExceeded        HttpStatus = 509
	Status510NotExtended                   HttpStatus = 510
	Status511NetworkAuthenticationRequired HttpStatus = 511
	Status521WebServerIsDown               HttpStatus = 521
)
