package repository

import (
	"api-egressos/model"
	"context"
)

type ProfileRepo interface {
	GetProfiles(ctx context.Context) ([]*model.Profile, error)
}
