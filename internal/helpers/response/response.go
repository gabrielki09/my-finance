package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponseData struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Error   any    `json:"error"`
}

type SuccessResponseData struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Data    any    `json:"data"`
}

func WriteJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func ErrorResponse(message string, data any) ErrorResponseData {
	return ErrorResponseData{
		Status:  false,
		Message: message,
		Error:   data,
	}
}

func SuccessResponse(message string, data any) SuccessResponseData {
	return SuccessResponseData{
		Status:  true,
		Message: message,
		Data:    data,
	}
}
