package utils

type Params203 struct {
	Name        string // Source name
	Description string // Source description
	Source      string // Source url
}

type Params207 struct {
	HttpStatus    uint16 // Status of the state
	ServerMessage string // Message of the state
	Detail        string //Detail of the state
}

type Params301 struct {
	OldSouce  string // Old URL or resource
	NewSource string // New URL or resource
}
