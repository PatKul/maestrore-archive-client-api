package domain

import (
	"encoding/json"
	"maestrore/core"
	"maestrore/core/httpclient"
	"net/http"
)

/**
 * Payload of encrypted value
 * @param value string
 */
type EncodedDto struct {
	Value string `json:"value"`
}

/**
 * This default route handler will handle encryption of
 * a given value
 */
type DefaultRouteHandler struct {
	router    *http.ServeMux
	encryptor *core.Encryptor
}

func NewDefaultRouteHandler(router *http.ServeMux, encryptor *core.Encryptor) *DefaultRouteHandler {
	return &DefaultRouteHandler{router: router, encryptor: encryptor}
}

func (h *DefaultRouteHandler) RegisterRoute() {
	h.router.HandleFunc("POST /encode", h.Encode)
}

func (h *DefaultRouteHandler) Encode(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var dto EncodedDto
	error := decoder.Decode(&dto)
	if error != nil {
		httpclient.ErrorHandler(w, error)
		return
	}

	encodedValue, err := h.encryptor.Encrypt(dto.Value)
	if err != nil {
		httpclient.ErrorHandler(w, err)
		return
	}

	httpclient.OkJsonResponse(w, encodedValue, "Value encoded successfully")
}
