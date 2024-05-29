package sales

import (
	"maestrore/core/httpclient"
	"maestrore/domain/sales/data"
	"maestrore/domain/sales/dto"
)

type SalesService struct {
	repository *data.SalesRepository
}

func NewSalesService(repository *data.SalesRepository) *SalesService {
	return &SalesService{repository: repository}
}

/**
 * Find sales matching query and return paginated result
 * @param query QueryPayload
 * @return *httpclient.HttpPaginatedDto[SalesListDto]
 * @return error
 */
func (s *SalesService) FindPagedSales(query *dto.QueryPayload) (*httpclient.HttpPaginatedDto[dto.SalesListDto], error) {
	count, error := s.repository.Count(query)
	if error != nil {
		return nil, error
	}

	sales, error := s.repository.FindPagedResult(query)
	if error != nil {
		return nil, error
	}

	var list []dto.SalesListDto = []dto.SalesListDto{}
	for _, sale := range sales {
		list = append(list, dto.SalesListDto{
			Id:                    sale.Id,
			LocationTransactionId: sale.LocationTransactionId,
			LocationName:          sale.LocationName,
			TillNumber:            sale.TillNumber,
			SalesDate:             sale.SalesDate,
			NetAmount:             sale.NetAmount,
			AmountPaid:            sale.AmountPaid,
			CommissionDue:         sale.CommissionDue,
		})
	}

	return &httpclient.HttpPaginatedDto[dto.SalesListDto]{
		Page:   query.Page,
		Limit:  query.Limit,
		Result: list,
		Total:  count,
	}, nil
}

/**
 * Get sales by id
 * @param id int
 * @return dto.SalesDetailDto
 * @return error
 */
func (s *SalesService) FindById(id string) (dto.SalesDetailDto, error) {
	sale, error := s.repository.FindById(id)
	if error != nil {
		return dto.SalesDetailDto{}, error
	}

	var entries []dto.SalesDetailEntryDto = []dto.SalesDetailEntryDto{}
	for _, entry := range sale.Entries {
		comment := ""
		if entry.Comment.Valid {
			comment = entry.Comment.String
		}

		entries = append(entries, dto.SalesDetailEntryDto{
			Id:              entry.Id,
			InventoryId:     entry.InventoryId,
			InventoryName:   entry.InventoryName,
			Comment:         comment,
			Quantity:        entry.Quantity,
			CostPrice:       entry.CostPrice,
			SalesPrice:      entry.SalesPrice,
			TotalCostPrice:  entry.TotalCostPrice,
			TotalSalesPrice: entry.TotalSalesPrice,
		})
	}

	var tenders []dto.SalesDetailTenderEntryDto = []dto.SalesDetailTenderEntryDto{}
	for _, tender := range sale.Tenders {
		tenders = append(tenders, dto.SalesDetailTenderEntryDto{
			Id:         tender.Id,
			TenderName: tender.TenderName,
			Amount:     tender.Amount,
		})
	}

	note := ""
	if sale.Note.Valid {
		note = sale.Note.String
	}
	return dto.SalesDetailDto{
		Id:                    sale.Id,
		LocationTransactionId: sale.LocationTransactionId,
		LocationName:          sale.LocationName,
		TillNumber:            sale.TillNumber,
		SalesDate:             sale.SalesDate,
		NetAmount:             sale.NetAmount,
		AmountPaid:            sale.AmountPaid,
		CommissionDue:         sale.CommissionDue,
		Note:                  note,
		Entries:               entries,
		Tenders:               tenders,
	}, nil
}
