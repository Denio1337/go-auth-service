package response

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SuccessResponse(data any) *Response {
	return &Response{
		Success: true,
		Message: "",
		Data:    data,
	}
}

func ErrorResponse(message string) *Response {
	return &Response{
		Success: false,
		Message: message,
		Data:    nil,
	}
}
