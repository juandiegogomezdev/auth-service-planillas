package storeauth

import (
	"proyecto/internal/authservice/modelauth"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type StoreAuth interface {
	GetByEmail(email string) (*modelauth.User, error)
	CreateUser(user *modelauth.User, code string, createdAt time.Time, expiresAt time.Time) error
	UpdateCode(id uuid.UUID, code string, createdAt time.Time, expiresAt time.Time) error
	ExistUser(email string) (bool, error)
	GetVerificationByEmail(email string) (VerificationInfo, error)
}

type store struct {
	db *sqlx.DB
}

func NewAuthStore(db *sqlx.DB) StoreAuth {
	return &store{db: db}
}

func (s *store) ExistUser(email string) (bool, error) {
	var exists bool
	q := `SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)`
	err := s.db.Get(&exists, q, email)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (s *store) CreateEmailVerification(id uuid.UUID, code string, created_at time.Time, expires_at time.Time) error {
	q := `INSERT INTO email_verifications (id, code, created_at, expires_at) VALUES ($1, $2, $3, $4)`
	_, err := s.db.Exec(q, id, code, created_at, expires_at)
	return err
}

func (s *store) CreateUser(user *modelauth.User, code string, createdAt time.Time, expiresAt time.Time) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	qUser := `INSERT INTO users (id, email, hashed_password, created_at) VALUES ($1, $2, $3, $4)`
	_, err = tx.Exec(qUser, user.ID, user.Email, user.HashedPassword, createdAt)
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

func (s *store) UpdateCode(id uuid.UUID, code string, createdAt time.Time, expiresAt time.Time) error {
	q := `UPDATE email_verifications SET code=$1, created_at=$2, expires_at=$3 WHERE id=$4`
	_, err := s.db.Exec(q, code, createdAt, expiresAt, id)
	return err
}

func (s *store) GetByEmail(email string) (*modelauth.User, error) {
	q := `SELECT id, email, hashed_password, created_at FROM users WHERE email=$1`
	row := s.db.QueryRow(q, email)

	var user modelauth.User
	err := row.Scan(&user.ID, &user.Email, &user.HashedPassword, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

type VerificationInfo struct {
	Code string
	ID   uuid.UUID
}

func (s *store) GetVerificationByEmail(email string) (VerificationInfo, error) {
	q := `SELECT code, id FROM email_verifications where id = (SELECT id FROM users WHERE email=$1)`
	row := s.db.QueryRow(q, email)

	var code string
	var id uuid.UUID
	err := row.Scan(&code, &id)
	if err != nil {
		return VerificationInfo{}, err
	}
	return VerificationInfo{Code: code, ID: id}, nil
}
