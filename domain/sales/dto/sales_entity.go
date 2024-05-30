package dto

import (
	"net/url"
	"strconv"
	"time"
)

type QueryPayload struct {
	SearchField string    `json:"search_field"`
	SearchValue string    `json:"search_value"`
	LocationId  string    `json:"location_id"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Page        int       `json:"page"`
	Limit       int       `json:"limit"`
	SortBy      string    `json:"sort_by"`
	OrderBy     string    `json:"order_by"`
}

func NewQueryPayload(params url.Values) *QueryPayload {
	startDate, err := time.Parse("2006-01-02", params.Get("start_date"))
	if err != nil {
		startDate = time.Now()
	}

	endDate, err := time.Parse("2006-01-02", params.Get("end_date"))
	if err != nil {
		endDate = time.Now()
	}

	page := 1
	limit := 10

	pageStr := params.Get("page")
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	limitStr := params.Get("limit")
	if limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	}

	return &QueryPayload{
		SearchField: params.Get("search_field"),
		SearchValue: params.Get("search_value"),
		StartDate:   startDate,
		EndDate:     endDate,
		Page:        page,
		Limit:       limit,
	}
}

type SalesListDto struct {
	Id                    int       `json:"id"`
	LocationTransactionId int       `json:"location_transaction_id"`
	LocationName          string    `json:"location_name"`
	TillNumber            int       `json:"till_number"`
	SalesDate             time.Time `json:"sales_date"`
	NetAmount             float64   `json:"net_amount"`
	AmountPaid            float64   `json:"amount_paid"`
	CommissionDue         float64   `json:"commission_due"`
	Note                  string    `json:"note"`
}

type SalesDetailDto struct {
	Id                    int                         `json:"id"`
	LocationTransactionId int                         `json:"location_transaction_id"`
	LocationName          string                      `json:"location_name"`
	TillNumber            int                         `json:"till_number"`
	SalesDate             time.Time                   `json:"sales_date"`
	NetAmount             float64                     `json:"net_amount"`
	AmountPaid            float64                     `json:"amount_paid"`
	CommissionDue         float64                     `json:"commission_due"`
	Note                  string                      `json:"note"`
	Entries               []SalesDetailEntryDto       `json:"entries"`
	Tenders               []SalesDetailTenderEntryDto `json:"tenders"`
}

type SalesDetailEntryDto struct {
	Id              int     `json:"id"`
	InventoryId     int     `json:"inventory_id"`
	InventoryName   string  `json:"inventory_name"`
	Comment         string  `json:"comment"`
	Quantity        int     `json:"quantity"`
	CostPrice       float64 `json:"cost_price"`
	SalesPrice      float64 `json:"sales_price"`
	TotalCostPrice  float64 `json:"total_cost_price"`
	TotalSalesPrice float64 `json:"total_sales_price"`
}

type SalesDetailTenderEntryDto struct {
	Id         int     `json:"id"`
	TenderName string  `json:"tender_name"`
	Amount     float64 `json:"amount"`
}
