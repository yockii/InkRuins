package domain

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewResponse(code int, message string, data any) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func NewSuccessResponse(data any) *Response {
	return NewResponse(0, "success", data)
}

func NewErrorResponse(code int, message string) *Response {
	return NewResponse(code, message, nil)
}
