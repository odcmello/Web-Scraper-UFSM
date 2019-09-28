package repository

import "api-egressos/model"

type ProfileRepo interface {
	GetProfiles() ([]*model.Profile, error)
}
