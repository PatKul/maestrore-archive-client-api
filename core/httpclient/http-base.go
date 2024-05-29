package httpclient

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

/**
 * HttpResponseDto
 */
type HttpResponseDto[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message"`
}

func NewHttpResponseDto[T any](data T, message string) HttpResponseDto[T] {
	return HttpResponseDto[T]{Data: data, Message: message}
}

/**
 * HttpListResponseDto
 */
type HttpListResponseDto[T any] struct {
	Data    []T    `json:"data"`
	Message string `json:"message"`
}

func NewHttpListResponseDto[T any](data []T, message string) HttpListResponseDto[T] {
	return HttpListResponseDto[T]{Data: data, Message: message}
}

type HttpPaginatedDto[T any] struct {
	Page   int `json:"page"`
	Limit  int `json:"limit"`
	Result []T `json:"result"`
	Total  int `json:"total"`
}

func NewHttpPaginatedDto[T any](page int, limit int, result []T, total int) HttpPaginatedDto[T] {
	return HttpPaginatedDto[T]{Page: page, Limit: limit, Result: result, Total: total}
}

func OkJsonResponse[T any](w http.ResponseWriter, data T, message string) {
	response := NewHttpResponseDto[T](data, message)

	json, error := json.Marshal(response)
	if error != nil {
		slog.Error(error.Error())
		InternalServerErrorHandler(w, error)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)

}
