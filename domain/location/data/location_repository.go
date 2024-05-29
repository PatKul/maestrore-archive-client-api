package data

import "database/sql"

type LocationRepository struct {
	db *sql.DB
}

func NewLocationRepository(db *sql.DB) *LocationRepository {
	repository := LocationRepository{
		db: db,
	}

	return &repository
}

/**
 * Find all locations
 * @return []LocationListData
 * @return error
 */
func (r *LocationRepository) FindAll() ([]LocationListData, error) {
	sql := "SELECT id, name FROM locations"
	rows, error := r.db.Query(sql)

	if error != nil {
		return nil, error
	}

	defer rows.Close()

	var locations []LocationListData

	for rows.Next() {
		var location LocationListData
		error := rows.Scan(&location.Id, &location.Name)

		if error != nil {
			return nil, error
		}

		locations = append(locations, location)
	}

	return locations, nil
}
