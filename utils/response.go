package utils

type SuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewSuccessResponse(message string, data any) SuccessResponse {
	repsonse := SuccessResponse{Status: "success", Message: message, Data: data}
	return repsonse
}
