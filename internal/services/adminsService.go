package services

import (
	"govno/internal/models"
	"govno/internal/repo"
)

type AdminService struct {
	Model models.Admin
	Repo repo.AdminRepo
}

func (s *AdminService) GetAdminsList() ([]*models.Admin, error) {
	return s.Repo.GetAdminList()
}