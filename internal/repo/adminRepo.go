package repo

import (
	"fmt"
	"govno/internal/db"
	"govno/internal/models"
	"log"
)

type AdminRepo struct{}

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

func (r *AdminRepo) SaveToken(adminID uint, token string) error {
	DB := db.GetDB()

	// Удаляем все существующие токены для этого администратора
	_, err := DB.Exec("DELETE FROM jwt_tokens WHERE admin_id = $1", adminID)
	if err != nil {
		log.Printf("Ошибка при удалении старых токенов: %v", err)
		return fmt.Errorf("не удалось удалить старые токены: %w", err)
	}

	// Сохраняем новый токен
	_, err = DB.Exec(
		"INSERT INTO jwt_tokens (token, admin_id, time, is_expired) VALUES ($1, $2, NOW(), false)",
		token,
		adminID,
	)
	if err != nil {
		log.Printf("Ошибка при сохранении токена: %v", err)
		return fmt.Errorf("не удалось сохранить токен: %w", err)
	}

	return nil
}

func (r *AdminRepo) GetAdminByCredentials(login, password string) (*models.Admin, error) {
	DB := db.GetDB()
	admin := &models.Admin{}

	err := DB.QueryRow(
		"SELECT id, login, name FROM admins WHERE login = $1 AND password = $2",
		login,
		password,
	).Scan(&admin.Id, &admin.Login, &admin.Name)

	if err != nil {
		return nil, err
	}

	return admin, nil
}

// метод для проверки истечения срока токена
func (r *AdminRepo) IsTokenExpired(token string) (bool, error) {
	DB := db.GetDB()
	var isExpired bool

	err := DB.QueryRow(
		"SELECT is_expired FROM jwt_tokens WHERE token = $1",
		token,
	).Scan(&isExpired)

	if err != nil {
		return true, fmt.Errorf("токен не найден: %w", err)
	}

	return isExpired, nil
}

// метод для принудительного удаления токена (например, при логауте)
func (r *AdminRepo) InvalidateToken(token string) error {
	DB := db.GetDB()

	_, err := DB.Exec(
		"UPDATE jwt_tokens SET is_expired = true WHERE token = $1",
		token,
	)

	return err
}
