package data

import (
	"maestrore/domain/sales/dto"
)

type SalesDataEncoder struct {
}

func NewSalesEncoder() *SalesDataEncoder {
	encoder := SalesDataEncoder{}

	return &encoder
}

/**
 * Encode client query to sql where clause query
 */
func (e *SalesDataEncoder) EncodeQuery(query *dto.QueryPayload) string {
	var whereClause string

	if query.SearchField != "" && query.SearchValue != "" {
		whereClause += e.EncodeSearchValue(query.SearchField, query.SearchValue)
	}

	dateRangeQueryClause := e.EncodeDateRange(query.StartDate.Format("2006-01-02"), query.EndDate.Format("2006-01-02"))

	if whereClause != "" && dateRangeQueryClause != "" {
		whereClause += " AND "
	}

	whereClause += dateRangeQueryClause

	return whereClause
}

/**
 * Encode search value
 * NB: Currently only note and comment can be searched since most of our client
 * queries are based on these fields
 * @param searchField string
 * @param searchValue string
 * @return string
 */
func (e *SalesDataEncoder) EncodeSearchValue(searchField string, searchValue string) string {
	if searchField == "sales_note" {
		return "s.note LIKE '%" + searchValue + "%'"
	} else if searchField == "entry_comment" {
		return "se.comment LIKE '%" + searchValue + "%'"
	}

	return ""
}

/**
 * Encode date range
 * @param startDate string
 * @param endDate string
 * @return string
 */
func (e *SalesDataEncoder) EncodeDateRange(startDate string, endDate string) string {
	var whereClause string

	if startDate != "" {
		whereClause += "sales_date >= '" + startDate + " 00:00:00'"
	}

	if endDate != "" {
		if whereClause != "" {
			whereClause += " AND "
		}
		whereClause += "sales_date <= '" + endDate + " 23:59:59'"
	}

	return whereClause
}
