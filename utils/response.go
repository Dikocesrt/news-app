package utils

type BaseSuccessResponse struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Data any `json:"data"`
}

type BaseErrorResponse struct {
	Status bool `json:"status"`
	Message string `json:"message"`
}

func NewBaseSuccessResponse(message string, data any) BaseSuccessResponse {
	return BaseSuccessResponse{
		Status: true,
		Message: message,
		Data: data,
	}
}

func NewBaseErrorResponse(message string) BaseErrorResponse {
	return BaseErrorResponse{
		Status: false,
		Message: message,
	}
}