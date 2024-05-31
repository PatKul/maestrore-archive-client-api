package sales

import (
	"maestrore/domain/sales/data"
	"maestrore/domain/sales/dto"
)

type SalesDtoEncoder struct {
}

func NewSalesDtoEncoder() *SalesDtoEncoder {
	encoder := SalesDtoEncoder{}

	return &encoder
}

/**
 * Encode list of sales data from database to list of sales dto
 * @param data *[]data.SalesListData
 * @return *[]dto.SalesListDto
 */
func (e *SalesDtoEncoder) EncodeListData(data *[]data.SalesListData) *[]dto.SalesListDto {
	var list []dto.SalesListDto = []dto.SalesListDto{}
	for _, sale := range *data {
		tillNumber := 0
		if sale.TillNumber.Valid {
			tillNumber = int(sale.TillNumber.Int64)
		}

		netAmount := 0.0
		if sale.NetAmount.Valid {
			netAmount = sale.NetAmount.Float64
		}

		list = append(list, dto.SalesListDto{
			Id:                    sale.Id,
			LocationTransactionId: sale.LocationTransactionId,
			LocationName:          sale.LocationName.String,
			TillNumber:            tillNumber,
			SalesDate:             string(sale.SalesDate),
			NetAmount:             netAmount,
			AmountPaid:            sale.AmountPaid,
			CommissionDue:         sale.CommissionDue,
			Note:                  sale.Note.String,
		})
	}

	return &list
}

/**
 * Encode sales detail data from database to sales detail dto
 * @param sale *data.SalesDetailData
 * @return *dto.SalesDetailDto
 */
func (e *SalesDtoEncoder) EncodeDetailData(sale *data.SalesDetailData) *dto.SalesDetailDto {
	var entries []dto.SalesDetailEntryDto = []dto.SalesDetailEntryDto{}
	for _, entry := range sale.Entries {
		entries = append(entries, e.EncodeSalesDetailEntry(&entry))
	}

	var tenders []dto.SalesDetailTenderEntryDto = []dto.SalesDetailTenderEntryDto{}
	for _, tender := range sale.Tenders {
		tenders = append(tenders, e.EncodeSalesDetailTenderEntry(&tender))
	}

	note := ""
	if sale.Note.Valid {
		note = sale.Note.String
	}

	tillNumber := 0
	if sale.TillNumber.Valid {
		tillNumber = int(sale.TillNumber.Int64)
	}

	netAmount := 0.0
	if sale.NetAmount.Valid {
		netAmount = sale.NetAmount.Float64
	}

	return &dto.SalesDetailDto{
		Id:                    sale.Id,
		LocationTransactionId: sale.LocationTransactionId,
		LocationName:          sale.LocationName.String,
		TillNumber:            tillNumber,
		SalesDate:             string(sale.SalesDate),
		NetAmount:             netAmount,
		AmountPaid:            sale.AmountPaid,
		CommissionDue:         sale.CommissionDue,
		Note:                  note,
		Entries:               entries,
		Tenders:               tenders,
	}
}

/**
 * Encode sales detail entry data from database to sales detail entry dto
 * @param data *data.SalesDetailEntryData
 * @return dto.SalesDetailEntryDto
 */
func (e *SalesDtoEncoder) EncodeSalesDetailEntry(data *data.SalesDetailEntryData) dto.SalesDetailEntryDto {
	comment := ""
	if data.Comment.Valid {
		comment = data.Comment.String
	}

	return dto.SalesDetailEntryDto{
		Id:              data.Id,
		InventoryId:     data.InventoryId,
		InventoryName:   data.InventoryName,
		Comment:         comment,
		Quantity:        data.Quantity,
		CostPrice:       data.CostPrice,
		SalesPrice:      data.SalesPrice,
		TotalCostPrice:  data.TotalCostPrice,
		TotalSalesPrice: data.TotalSalesPrice,
	}
}

/**
 * Encode sales detail tender entry data from database to sales detail tender entry dto
 * @param data *data.SalesDetailTenderEntryData
 * @return dto.SalesDetailTenderEntryDto
 */
func (e *SalesDtoEncoder) EncodeSalesDetailTenderEntry(data *data.SalesDetailTenderEntryData) dto.SalesDetailTenderEntryDto {
	return dto.SalesDetailTenderEntryDto{
		Id:         data.Id,
		TenderName: data.TenderName,
		Amount:     data.Amount,
	}
}
