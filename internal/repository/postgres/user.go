package postgres

import (
	"database/sql"
	"dev_toolbox_auth_service/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *domain.User) error {
	query := `
        INSERT INTO users (email, password_hash, full_name)
        VALUES ($1, $2, $3)
        RETURNING id`

	return r.db.QueryRow(query, user.Email, user.PasswordHash, user.FullName).Scan(&user.ID)
}
