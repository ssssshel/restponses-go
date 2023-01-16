package internal_interfaces

type BaseResponse struct {
	HttpStatus    uint16 `json:"httpStatus"`
	ServerMessage string `json:"serverMessage"`
	Success       bool   `json:"success"`
	Error         bool   `json:"error"`
}

type GenericSuccessResponse struct {
	BaseResponse
	Payload     interface{} `json:"payload"`
	MoreDetails string      `json:"moreDetails"`
}

type GenericErrorResponse struct {
	BaseResponse
	ErrorMessage string `json:"errorMessage"`
}
