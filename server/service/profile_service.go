package service

import (
	"api-egressos/model"
	"api-egressos/repository"
	profile "api-egressos/repository/sql"
	"context"
)

type ProfileService struct {
	repo repository.ProfileRepo
}

func NewProfileService() *ProfileService {
	return &ProfileService{
		repo: profile.NewSQLProfileRepository(),
	}
}

func (p *ProfileService) GetProfiles(ctx context.Context) ([]*model.Profile, error) {
	profiles, err := p.repo.GetProfiles(ctx)
	return profiles, err
}
