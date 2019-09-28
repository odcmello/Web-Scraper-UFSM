package service

import (
	"api-egressos/model"
	"api-egressos/repository"
	profile "api-egressos/repository/sql"
)

type ProfileService struct {
	repo repository.ProfileRepo
}

func NewProfileService() *ProfileService {
	return &ProfileService{
		repo: profile.NewSQLProfileRepository(),
	}
}

func (p *ProfileService) GetProfiles() ([]*model.Profile, error) {
	profiles, err := p.repo.GetProfiles()
	return profiles, err
}
