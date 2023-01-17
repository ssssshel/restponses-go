package interfaces

// Success and Error properties
type SuccessErrorProps struct {
	Success bool `json:"success,omitempty"`
	Error   bool `json:"error,omitempty"`
}

// Base of all responses. Also default response for 1xx, 204 & 205 status code
type BaseResponse struct {
	HttpStatus    uint16 `json:"status"`
	ServerMessage string `json:"server_message"`
	Resource      string `json:"resource,omitempty"`
}

// Default response for 2xx status code
type Status2xx_Response struct {
	BaseResponse
	Data    interface{} `json:"data"`              // Anything you want to return
	Details string      `json:"details,omitempty"` // Datailed info of your response
	SuccessErrorProps
}

// Response for 201 status code
type Status201Created_Response struct {
	Status2xx_Response        //The classic response that you love
	Location           string `json:"location"` //URL or place where your creation can be found
}

// Response for 202 status code
type Status202Accepted_Response struct {
	Status2xx_Response        //The classic response that you love
	RequestId          string `json:"request_id"` //ID with wich you can follow the process of your request
}

// Response for 203 status code
type Status203NonAI_Response struct {
	Status2xx_Response //The classic response that you love

	//Here you can put info about the third source involved
	Source struct {
		Name        string `json:"name"`        //Source name
		Description string `json:"description"` //Source description
		Source      string `json:"source"`      //Source url
	} `json:"source"`
}

// Response for 207 status code
type Status207MultiStatus_Response struct {
	Status2xx_Response                //The classic response that you love
	States             []BaseResponse `json:"states"` //Here you can put info about the multiple states (Status, server_message and resource)
}

// Default response for 3xx status code
type Status3xx_Response struct {
	BaseResponse
	Details string `json:"details,omitempty"` // Datailed info of your response
	SuccessErrorProps
}

type Status300MultipleChoices_Response struct {
	Status3xx_Response
	Options []interface{} //Here you can put the options available for the requested resource
}

type Status301MovedPermanently_Response struct {
	Status3xx_Response
	Sources struct {
		OldSouce  string
		NewSource string
	}
}

type GenericErrorResponse struct {
	BaseResponse
	ErrorMessage string `json:"errorMessage"`
}
