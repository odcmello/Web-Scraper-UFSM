package profile

import (
	"api-egressos/database"
	"api-egressos/model"
	"api-egressos/repository"

	"github.com/jmoiron/sqlx"
)

func NewSQLProfileRepository() repository.ProfileRepo {
	return &pRepository{}
}

type pRepository struct {
	Db *sqlx.DB
}

func (p *pRepository) GetProfiles() ([]*model.Profile, error) {
	rows, err := database.Db.Conn.Query("SELECT * FROM egressos")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	profiles := make([]*model.Profile, 0)
	for rows.Next() {
		var profile model.Profile
		if err = rows.Scan(&profile.ID, &profile.Name, &profile.JobTitle, &profile.Company, &profile.Location, &profile.URL); err != nil {
			return nil, err
		}
		profiles = append(profiles, &profile)
	}

	return profiles, nil
}
