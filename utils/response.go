package utils

import "makanan-app/entities"

type BaseSuccessResponse struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Data any `json:"data"`
}

type BaseMetadataSuccessResponse struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Metadata entities.Metadata `json:"metadata"`
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

func NewBaseMetadataSuccessResponse(message string, metadata entities.Metadata, data any) BaseMetadataSuccessResponse {
	return BaseMetadataSuccessResponse{
		Status: true,
		Message: message,
		Metadata: metadata,
		Data: data,
	}
}

func NewBaseErrorResponse(message string) BaseErrorResponse {
	return BaseErrorResponse{
		Status: false,
		Message: message,
	}
}