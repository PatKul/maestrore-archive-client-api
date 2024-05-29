package location

import (
	"maestrore/domain/location/data"
)

type LocationService struct {
	repository *data.LocationRepository
}

func NewLocationService(repository *data.LocationRepository) *LocationService {
	return &LocationService{repository: repository}
}

/**
 * Find all locations
 * @return []data.LocationListData
 * @return error
 */
func (s *LocationService) FindAll() ([]data.LocationListData, error) {
	locations, error := s.repository.FindAll()
	if error != nil {
		return nil, error
	}

	return locations, nil
}
