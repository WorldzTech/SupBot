package db

import (
	"context"
	"govno/internal/models"
)

type DB interface {
	GetAdminByLogpass(ctx context.Context, filter models.AdminModelLogpassFilter) ([]models.AdminModel, error)
}
