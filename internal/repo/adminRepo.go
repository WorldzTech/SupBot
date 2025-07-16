package repo

import (
	"govno/internal/db"
	"govno/internal/models"
)

type AdminRepo struct {}

func (r *AdminRepo) GetAdminList() ([]*models.Admin, error) {
	DB := db.GetDB()
	rows, err := DB.Query("SELECT id, name, login FROM admins")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var response []*models.Admin
	for rows.Next() {
		var admin models.Admin
		err := rows.Scan(&admin.Id, &admin.Name, &admin.Login)
		if err != nil {
			return nil, err
		}

		response = append(response, &admin)
	}

	return response, nil
}