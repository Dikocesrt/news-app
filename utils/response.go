package utils

import "test-indonesia-cakap-digital/entities"

type BaseSuccessResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data any `json:"data"`
}

type BaseMetadataSuccessResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Metadata entities.Metadata `json:"metadata"`
	Data any `json:"data"`
}

type BaseErrorResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

func NewBaseSuccessResponse(message string, data any) BaseSuccessResponse {
	return BaseSuccessResponse{
		Status: "success",
		Message: message,
		Data: data,
	}
}

func NewBaseMetadataSuccessResponse(message string, metadata entities.Metadata, data any) BaseMetadataSuccessResponse {
	return BaseMetadataSuccessResponse{
		Status: "success",
		Message: message,
		Metadata: metadata,
		Data: data,
	}
}

func NewBaseErrorResponse(message string) BaseErrorResponse {
	return BaseErrorResponse{
		Status: "fail",
		Message: message,
	}
}