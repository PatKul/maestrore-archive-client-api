package location

import (
	"database/sql"
	"net/http"

	"maestrore/core/httpclient"
	"maestrore/domain/location/data"
)

type RouteHandler struct {
	router  *http.ServeMux
	service *LocationService
}

func NewRouteHandler(router *http.ServeMux, db *sql.DB) *RouteHandler {
	repository := data.NewLocationRepository(db)
	service := NewLocationService(repository)

	return &RouteHandler{router: router, service: service}
}

/**
 * Register all location routes
 */
func (handler *RouteHandler) RegisterRoute() {
	handler.router.HandleFunc("GET /location", handler.FindAll)
}

/**
 * Get all locations
 * @param w http.ResponseWriter
 * @param r *http.Request
 */
func (handler *RouteHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	locations, err := handler.service.FindAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpclient.OkJsonResponse(w, locations, "Locations successfully retrieved")
}
