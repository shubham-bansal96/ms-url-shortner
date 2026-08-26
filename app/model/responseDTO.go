package model

type ResponseDTO struct {
	Data  interface{} `json:"data"`
	Error *Error      `json:"error"`
}

type Error struct {
	ErrorCode    *int    `json:"code,omitempty"`
	ErrorMessage *string `json:"message,omitempty"`
}

// NewError creates an Error with the given HTTP status code and message
func NewError(code int, message string) *Error {
	return &Error{ErrorCode: &code, ErrorMessage: &message}
}
