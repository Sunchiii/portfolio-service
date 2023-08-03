package utils

import (
    "encoding/json"
    "net/http"
)

// JSONResponse represents a JSON response.
type JSONResponse struct {
    Data interface{} `json:"data,omitempty"`
    Meta interface{} `json:"meta,omitempty"`
}

// NewJSONResponse creates a new JSONResponse instance.
func NewJSONResponse(data interface{}, meta interface{}) *JSONResponse {
    return &JSONResponse{
        Data: data,
        Meta: meta,
    }
}

// WriteJSONResponse writes a JSON response to the HTTP response writer.
func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}, meta interface{}) {
    jsonResponse := NewJSONResponse(data, meta)
    jsonBytes, err := json.Marshal(jsonResponse)
    if err != nil {
        WriteErrorResponse(w, InternalServerError("Error encoding JSON response"))
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    w.Write(jsonBytes)
}

// WriteErrorResponse writes an error response to the HTTP response writer.
func WriteErrorResponse(w http.ResponseWriter, err *HTTPError) {
    errorResponse := NewErrorResponse(err.Message)
    jsonBytes, _ := json.Marshal(errorResponse)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(err.Status)
    w.Write(jsonBytes)
}
