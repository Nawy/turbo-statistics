package utils

import (
	"encoding/json"
	"net/http"
)

// HTTPMethod is special type for control http method names
type HTTPMethod string

const (
	POST   = HTTPMethod("POST")
	GET    = HTTPMethod("GET")
	DELETE = HTTPMethod("DELETE")
	PUT    = HTTPMethod("PUT")
)

type errorMessageJSON struct {
	Message string `json:"message"`
}

func IsHttpMethod(r *http.Request, method HTTPMethod) bool {
	return (string(method) == r.Method)
}

func WriteBadGateway(w http.ResponseWriter, message string) {
	WriteJsonResponse(w, &errorMessageJSON{message}, 400)
}

func WriteJsonResponse(w http.ResponseWriter, response interface{}, statusCode int) {
	responseJSON, err := json.Marshal(response)

	if err != nil {
		panic("Cannot create json response")
	}

	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
}
