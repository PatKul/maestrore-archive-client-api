package sales

import (
	"database/sql"
	"maestrore/core/httpclient"
	"maestrore/domain/sales/data"
	"maestrore/domain/sales/dto"
	"net/http"
)

type RouteHandler struct {
	router  *http.ServeMux
	service *SalesService
}

func NewRouteHandler(router *http.ServeMux, db *sql.DB) *RouteHandler {
	repository := data.NewSalesRepository(db)
	service := NewSalesService(repository)

	return &RouteHandler{router: router, service: service}
}

/**
 * Register all sales routes
 */
func (handler *RouteHandler) RegisterRoute() {
	handler.router.HandleFunc("GET /sale", handler.FindPagedSales)
	handler.router.HandleFunc("GET /sale/{id}", handler.GetById)
}

/**
 * Get paginated sales
 * @param w http.ResponseWriter
 * @param r *http.Request
 */
func (handler *RouteHandler) FindPagedSales(w http.ResponseWriter, r *http.Request) {
	queryPayload := dto.NewQueryPayload(r.URL.Query())
	pagedSales, err := handler.service.FindPagedSales(queryPayload)

	if err != nil {
		httpclient.ErrorHandler(w, err)
		return
	}

	httpclient.OkJsonResponse(w, pagedSales, "Sales successfully retrieved")
}

/**
 * Get sales by id
 * @param w http.ResponseWriter
 * @param r *http.Request
 */
func (handler *RouteHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	sale, err := handler.service.FindById(id)

	if err != nil {
		httpclient.ErrorHandler(w, err)
		return
	}

	httpclient.OkJsonResponse(w, sale, "Sales successfully retrieved")
}
