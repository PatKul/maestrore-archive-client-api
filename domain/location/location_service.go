package location

import (
	"maestrore/domain/location/data"
	"maestrore/domain/location/dto"
)

type LocationService struct {
	repository *data.LocationRepository
	encoder    *LocationDtoEncoder
}

func NewLocationService(repository *data.LocationRepository) *LocationService {
	encoder := NewLocationDtoEncoder()

	return &LocationService{
		repository: repository,
		encoder:    encoder,
	}
}

/**
 * Find all locations
 * @return []data.LocationListData
 * @return error
 */
func (s *LocationService) FindAll() ([]dto.LocationListDto, error) {
	locations, error := s.repository.FindAll()
	if error != nil {
		return nil, error
	}

	list := s.encoder.EncodeListData(&locations)

	return *list, nil
}
