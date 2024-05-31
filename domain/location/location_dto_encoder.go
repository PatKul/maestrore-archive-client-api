package location

import (
	"maestrore/domain/location/data"
	"maestrore/domain/location/dto"
)

type LocationDtoEncoder struct {
}

func NewLocationDtoEncoder() *LocationDtoEncoder {
	encoder := LocationDtoEncoder{}

	return &encoder
}

/**
 * Encode list of sales data from database to list of sales dto
 * @param data *[]data.LocationListData
 * @return *[]dto.LocationListDto
 */
func (e *LocationDtoEncoder) EncodeListData(data *[]data.LocationListData) *[]dto.LocationListDto {
	var list []dto.LocationListDto = []dto.LocationListDto{}
	for _, location := range *data {

		list = append(list, dto.LocationListDto{
			Id:   location.Id,
			Name: location.Name,
		})
	}

	return &list
}
