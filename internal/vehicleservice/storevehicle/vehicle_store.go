package store

import "database/sql"

type Store struct {
	store *sql.DB
}
