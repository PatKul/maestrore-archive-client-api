package data

import (
	"database/sql"
	"time"
)

type SalesListData struct {
	Id                    int
	LocationTransactionId int
	LocationName          string
	CustomerId            int
	TillNumber            int
	SalesDate             time.Time
	NetAmount             float64
	AmountPaid            float64
	CommissionDue         float64
	Note                  sql.NullString
}

type SalesDetailData struct {
	Id                    int
	LocationTransactionId int
	LocationName          string
	CustomerId            int
	TillNumber            int
	SalesDate             time.Time
	NetAmount             float64
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
