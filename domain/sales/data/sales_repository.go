package data

import (
	"database/sql"
	"fmt"
	"strconv"

	"maestrore/domain/sales/dto"
)

type SalesRepository struct {
	db      *sql.DB
	encoder *SalesDataEncoder
}

func NewSalesRepository(db *sql.DB) *SalesRepository {
	repository := SalesRepository{
		db:      db,
		encoder: NewSalesEncoder(),
	}

	return &repository
}

/**
 * Count the number of sales
 * @param query QueryPayload
 * @return int
 * @return error
 */
func (r *SalesRepository) Count(query *dto.QueryPayload) (int, error) {
	creteria := r.encoder.EncodeQuery(query)

	whereClause := ""

	if creteria != "" {
		whereClause = "WHERE " + creteria
	}

	var count int
	sql := fmt.Sprintf("SELECT " +
		"COUNT(*) " +
		"FROM sales s " +
		"LEFT JOIN locations loc ON loc.id = s.location_id " +
		"LEFT JOIN sales_entries se ON se.sales_id = s.id " +
		whereClause)
	error := r.db.QueryRow(sql).Scan(&count)

	if error != nil {
		return 0, fmt.Errorf("failed to execute query: %s \nsql: %s", error.Error(), sql)
	}

	return count, nil
}

/**
 * Find sales with pagination
 * @param query SalesQueryDto
 * @return []SalesDto
 * @return error
 */
func (repository *SalesRepository) FindPagedResult(query *dto.QueryPayload) ([]SalesListData, error) {
	creteria := repository.encoder.EncodeQuery(query)

	whereClause := ""

	if creteria != "" {
		whereClause = "WHERE " + creteria
	}

	limitStr := strconv.Itoa(query.Limit)
	offsetStr := strconv.Itoa((query.Page - 1) * query.Limit)

	sql := fmt.Sprintf("SELECT "+
		"s.id, "+
		"s.location_transaction_id,"+
		"loc.name,"+
		"s.consumer_id,"+
		"s.till_number,"+
		"s.net_amount,"+
		"s.amount_paid,"+
		"s.commission_due,"+
		"s.note "+
		"FROM sales s LEFT JOIN locations loc ON loc.id = s.location_id "+
		"LEFT JOIN sales_entries se ON se.sales_id = s.id "+
		whereClause+
		" ORDER BY s.sales_date ASC "+
		"LIMIT %s OFFSET %s", limitStr, offsetStr)

	salesRows, error := repository.db.Query(sql)

	fmt.Print(sql)

	if error != nil {
		return nil, fmt.Errorf("failed to execute query: %s \n sql: %s", error.Error(), sql)
	}

	var sales []SalesListData
	for salesRows.Next() {
		var sale SalesListData

		error = salesRows.Scan(
			&sale.Id,
			&sale.LocationTransactionId,
			&sale.LocationName,
			&sale.CustomerId,
			&sale.TillNumber,
			&sale.NetAmount,
			&sale.AmountPaid,
			&sale.CommissionDue,
			&sale.Note)

		if error != nil {
			return nil, error
		}

		sales = append(sales, sale)
	}

	return sales, nil
}

func (repository *SalesRepository) FindById(id string) (SalesDetailData, error) {
	var sales SalesDetailData

	error := repository.db.QueryRow("SELECT "+
		"s.id, "+
		"s.location_transaction_id,"+
		"loc.name,"+
		"s.consumer_id,"+
		"s.till_number,"+
		"s.net_amount,"+
		"s.amount_paid,"+
		"s.commission_due,"+
		"s.note "+
		"FROM sales s LEFT JOIN locations loc ON loc.id = s.location_id "+
		"WHERE s.id = ?", id).
		Scan(
			&sales.Id,
			&sales.LocationTransactionId,
			&sales.LocationName,
			&sales.CustomerId,
			&sales.TillNumber,
			&sales.NetAmount,
			&sales.AmountPaid,
			&sales.CommissionDue,
			&sales.Note)

	if error != nil {
		return SalesDetailData{}, error
	}

	entries, error := repository.GetEntries(id)
	if error != nil {
		return SalesDetailData{}, error
	}

	tenders, error := repository.GetTenders(id)
	if error != nil {
		return SalesDetailData{}, error
	}

	sales.Entries = entries
	sales.Tenders = tenders

	return sales, nil

}

/**
 * Get sales entries under a given sales id
 * @param salesId int
 * @return []SalesEntryDto
 * @return error
 */
func (repository *SalesRepository) GetEntries(salesId string) ([]SalesDetailEntryData, error) {
	salesEntryRows, error := repository.db.Query("SELECT "+
		"se.id, "+
		"se.inventory_id, "+
		"inv.name, "+
		"se.comment, "+
		"se.quantity,"+
		"se.cost_price, "+
		"se.sales_price, "+
		"se.total_cost_price, "+
		"se.total_sales_price "+
		"FROM sales_entries se LEFT JOIN inventories inv ON inv.id = se.inventory_id "+
		"WHERE se.sales_id = ?", salesId)

	if error != nil {
		return nil, error
	}

	var salesEntries []SalesDetailEntryData
	for salesEntryRows.Next() {
		var entry SalesDetailEntryData

		error = salesEntryRows.Scan(
			&entry.Id,
			&entry.InventoryId,
			&entry.InventoryName,
			&entry.Comment,
			&entry.Quantity,
			&entry.CostPrice,
			&entry.SalesPrice,
			&entry.TotalCostPrice,
			&entry.TotalSalesPrice)

		if error != nil {
			return nil, error
		}

		salesEntries = append(salesEntries, entry)
	}

	return salesEntries, nil
}

func (repository *SalesRepository) GetTenders(salesId string) ([]SalesDetailTenderEntryData, error) {
	salesTenderRows, error := repository.db.Query("SELECT "+
		"te.id, "+
		"t.name, "+
		"te.amount "+
		"FROM tender_entries te "+
		"LEFT JOIN tenders t ON t.id = te.tender_id "+
		"WHERE transaction_type = 'S' AND te.transaction_id = ? ", salesId)

	if error != nil {
		return nil, error
	}

	var salesTenders []SalesDetailTenderEntryData
	for salesTenderRows.Next() {
		var tender SalesDetailTenderEntryData

		error = salesTenderRows.Scan(
			&tender.Id,
			&tender.TenderName,
			&tender.Amount)

		if error != nil {
			return nil, error
		}

		salesTenders = append(salesTenders, tender)
	}

	return salesTenders, nil
}
