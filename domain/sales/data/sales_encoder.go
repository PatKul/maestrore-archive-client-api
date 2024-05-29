package data

import (
	"maestrore/domain/sales/dto"
)

type SalesEncoder struct {
}

func NewSalesEncoder() *SalesEncoder {
	encoder := SalesEncoder{}

	return &encoder
}

/**
 * Encode client query to sql where clause query
 */
func (encoder *SalesEncoder) EncodeQuery(query *dto.QueryPayload) string {
	var whereClause string

	if query.SearchField != "" && query.SearchValue != "" {
		if query.SearchField == "sales_note" {
			whereClause = "s.note LIKE '%" + query.SearchValue + "%'"
		} else if query.SearchField == "entry_comment" {
			whereClause = "se.comment LIKE '%" + query.SearchValue + "%'"
		}
	}

	if !query.StartDate.IsZero() {
		if whereClause != "" {
			whereClause += " AND "
		}

		whereClause += "sales_date >= '" + query.StartDate.Format("2006-01-02") + " 00:00:00'"
	}

	if !query.EndDate.IsZero() {
		if whereClause != "" {
			whereClause += " AND "
		}

		whereClause += "sales_date <= '" + query.EndDate.Format("2006-01-02") + " 23:59:59'"
	}

	return whereClause
}
