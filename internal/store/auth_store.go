package store

import (
	"database/sql"
	"proyecto/internal/model"
	"time"

	"github.com/google/uuid"
)

type storeAuth struct {
	db *sql.DB
}

type AuthStore interface {
	Exist(email string) (bool, error)
	CreateUser(user *model.User, code string, createdAt time.Time, expiresAt time.Time) error
	GenerateCode(id uuid.UUID, code string, createdAt time.Time, expiresAt time.Time) error
}

func NewAuthStore(db *sql.DB) AuthStore {
	return &storeAuth{db: db}
}

func (s *storeAuth) CreateEmailVerification(id uuid.UUID, code string, created_at time.Time, expires_at time.Time) error {
	q := `INSERT INTO email_verifications (id, code, created_at, expires_at) VALUES ($1, $2, $3, $4)`
	_, err := s.db.Exec(q, id, code, created_at, expires_at)
	return err
}

func (s *storeAuth) CreateUser(user *model.User, code string, createdAt time.Time, expiresAt time.Time) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	qUser := `INSERT INTO users (id, email, hashed_password) VALUES ($1, $2, $3)`
	_, err = tx.Exec(qUser, user.ID, user.Email, user.HashedPassword)
	if err != nil {
		tx.Rollback()
		return err
	}

	qCode := `INSERT INTO email_verifications (id, code, created_at, expires_at) VALUES ($1, $2, $3, $4)`
	_, err = tx.Exec(qCode, user.ID, code, createdAt, expiresAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s *storeAuth) Exist(email string) (bool, error) {
	q := `SELECT 1 FROM users WHERE email=$1`
	row := s.db.QueryRow(q, email)
	var exists int
	err := row.Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *storeAuth) GenerateCode(id uuid.UUID, code string, createdAt time.Time, expiresAt time.Time) error {
	q := `UPDATE email_verifications SET code=$1, created_at=$2, expires_at=$3 WHERE id=$4`
	_, err := s.db.Exec(q, code, createdAt, expiresAt, id)
	return err
}
