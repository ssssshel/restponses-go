package restponses

type StatusCode1xx uint16
type StatusCode2xx uint16
type StatusCode3xx uint16
type StatusCode4xx uint16
type StatusCode5xx uint16

const (
	Status100 StatusCode1xx = 100
	Status101 StatusCode1xx = 101
	Status102 StatusCode1xx = 102
	Status103 StatusCode1xx = 103
)

const (
	Status200 StatusCode2xx = 200
	Status201 StatusCode2xx = 201
	Status202 StatusCode2xx = 202
	Status203 StatusCode2xx = 203
	Status204 StatusCode2xx = 204
	Status205 StatusCode2xx = 205
	Status206 StatusCode2xx = 206
	Status207 StatusCode2xx = 207
	Status208 StatusCode2xx = 208
	Status226 StatusCode2xx = 226
)

const (
	Status300 StatusCode3xx = 300
	Status301 StatusCode3xx = 301
	Status302 StatusCode3xx = 302
	Status303 StatusCode3xx = 303
	Status304 StatusCode3xx = 304
	Status305 StatusCode3xx = 305
	Status307 StatusCode3xx = 307
	Status308 StatusCode3xx = 309
)

const (
	Status400 StatusCode4xx = 400
	Status401 StatusCode4xx = 401
	Status402 StatusCode4xx = 402
	Status403 StatusCode4xx = 403
	Status404 StatusCode4xx = 404
	Status405 StatusCode4xx = 405
	Status406 StatusCode4xx = 406
	Status407 StatusCode4xx = 407
	Status408 StatusCode4xx = 408
	Status409 StatusCode4xx = 409
	Status410 StatusCode4xx = 410
	Status411 StatusCode4xx = 411
	Status412 StatusCode4xx = 412
	Status413 StatusCode4xx = 413
	Status414 StatusCode4xx = 414
	Status415 StatusCode4xx = 415
	Status416 StatusCode4xx = 416
	Status417 StatusCode4xx = 417
	Status418 StatusCode4xx = 418
	Status421 StatusCode4xx = 421
	Status422 StatusCode4xx = 422
	Status423 StatusCode4xx = 423
	Status424 StatusCode4xx = 424
	Status425 StatusCode4xx = 425
	Status426 StatusCode4xx = 426
	Status428 StatusCode4xx = 428
	Status429 StatusCode4xx = 429
	Status431 StatusCode4xx = 431
	Status451 StatusCode4xx = 451
)

const (
	Status500 StatusCode5xx = 500
	Status501 StatusCode5xx = 501
	Status502 StatusCode5xx = 502
	Status503 StatusCode5xx = 503
	Status504 StatusCode5xx = 504
	Status505 StatusCode5xx = 505
	Status506 StatusCode5xx = 506
	Status507 StatusCode5xx = 507
	Status508 StatusCode5xx = 508
	Status509 StatusCode5xx = 509
	Status510 StatusCode5xx = 510
	Status511 StatusCode5xx = 511
	Status521 StatusCode5xx = 521
)
