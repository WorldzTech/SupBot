package services

import (
	"govno/internal/models"
	"govno/internal/repo"
)

type AdminService struct {
	Repo repo.AdminRepo
}

func (s *AdminService) GetAdminsList() ([]*models.Admin, error) {
	return s.Repo.GetAdminList()
}

func (s *AdminService) Authenticate(login, password string) (*models.Admin, error) {
	admin, err := s.Repo.GetAdminByCredentials(login, password)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (s *AdminService) SaveToken(adminID uint, token string) error {
	return s.Repo.SaveToken(adminID, token)
}
