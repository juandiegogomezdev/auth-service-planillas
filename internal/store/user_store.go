package store

import (
	"database/sql"
	"proyecto/internal/dto"
	"proyecto/internal/model"
)

type storeUser struct {
	db *sql.DB
}

type UserStore interface {
	Create(user *model.User) error
	GetAll() ([]*dto.UsersQuery, error)
}

func NewUserStore(db *sql.DB) UserStore {
	return &storeUser{db: db}
}

func (s *storeUser) Create(user *model.User) error {
	q := `INSERT INTO users (id, email, hashed_password) VALUES ($1, $2, $3)`
	_, err := s.db.Exec(q, user.ID, user.Email, user.HashedPassword)
	return err
}

func (s *storeUser) GetAll() ([]*dto.UsersQuery, error) {
	q := `SELECT id, email, created_at FROM users`
	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize as an empty (non-nil) slice so JSON encodes it as [] instead of null when empty
	users := make([]*dto.UsersQuery, 0)
	for rows.Next() {
		var user dto.UsersQuery
		if err := rows.Scan(&user.ID, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
