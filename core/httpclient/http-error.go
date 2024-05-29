package httpclient

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"maestrore/core"
)

type HttpError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewHttpError(message string, code int) HttpError {
	return HttpError{Message: message, Code: code}
}

func ErrorHandler(w http.ResponseWriter, error error) {
	slog.Error(error.Error())

	if errors.Is(error, core.ErrorNotFound) {
		NotFoundHandler(w, error)
		return
	}

	if errors.Is(error, core.ErrorBadRequest) {
		BadRequestHandler(w, error)
		return
	}

	if errors.Is(error, core.ErrorConflict) {
		ConflictHandler(w, error)
		return
	}

	InternalServerErrorHandler(w, error)
	return
}

func InternalServerErrorHandler(w http.ResponseWriter, error error) {
	w.WriteHeader(http.StatusInternalServerError)

	payload := EncodeJsonResponse(error.Error(), http.StatusInternalServerError)
	w.Write(payload)
}

func BadRequestHandler(w http.ResponseWriter, error error) {
	w.WriteHeader(http.StatusBadRequest)

	payload := EncodeJsonResponse(error.Error(), http.StatusBadRequest)
	w.Write(payload)
}

func NotFoundHandler(w http.ResponseWriter, error error) {
	w.WriteHeader(http.StatusNotFound)

	payload := EncodeJsonResponse(error.Error(), http.StatusNotFound)
	w.Write(payload)
}

func UnauthorizedHandler(w http.ResponseWriter, error error) {
	w.WriteHeader(http.StatusUnauthorized)

	payload := EncodeJsonResponse(error.Error(), http.StatusUnauthorized)
	w.Write(payload)
}

func ForbiddenHandler(w http.ResponseWriter, error error) {
	w.WriteHeader(http.StatusForbidden)

	payload := EncodeJsonResponse(error.Error(), http.StatusForbidden)
	w.Write(payload)
}

func ConflictHandler(w http.ResponseWriter, error error) {
	w.WriteHeader(http.StatusConflict)

	payload := EncodeJsonResponse(error.Error(), http.StatusConflict)
	w.Write(payload)
}

func EncodeJsonResponse(message string, code int) []byte {
	return []byte(fmt.Sprintf(`{"message":"%s","code":%d}`, message, code))
}
