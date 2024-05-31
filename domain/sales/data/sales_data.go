package data

import (
	"database/sql"
)

type SalesListData struct {
	Id                    int
	LocationTransactionId int
	LocationName          sql.NullString
	CustomerId            int
	TillNumber            sql.NullInt64
	SalesDate             []uint8
	NetAmount             sql.NullFloat64
	AmountPaid            float64
	CommissionDue         float64
	Note                  sql.NullString
}

type SalesDetailData struct {
	Id                    int
	LocationTransactionId int
	LocationName          sql.NullString
	CustomerId            int
	TillNumber            sql.NullInt64
	SalesDate             []uint8
	NetAmount             sql.NullFloat64
	AmountPaid            float64
	CommissionDue         float64
	Note                  sql.NullString
	Entries               []SalesDetailEntryData
	Tenders               []SalesDetailTenderEntryData
}

type SalesDetailEntryData struct {
	Id              int
	InventoryId     int
	InventoryName   string
	Comment         sql.NullString
	Quantity        int
	CostPrice       float64
	SalesPrice      float64
	TotalCostPrice  float64
	TotalSalesPrice float64
}

type SalesDetailTenderEntryData struct {
	Id         int
	TenderName string
	Amount     float64
}
