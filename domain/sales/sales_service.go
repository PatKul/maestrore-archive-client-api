package sales

import (
	"maestrore/core/httpclient"
	"maestrore/domain/sales/data"
	"maestrore/domain/sales/dto"
)

type SalesService struct {
	repository *data.SalesRepository
	encoder    *SalesDtoEncoder
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

	var list = s.encoder.EncodeListData(&sales)

	return &httpclient.HttpPaginatedDto[dto.SalesListDto]{
		Page:   query.Page,
		Limit:  query.Limit,
		Result: *list,
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

	dto := s.encoder.EncodeDetailData(&sale)

	return *dto, nil

}
