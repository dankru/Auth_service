package pg_repo

import (
	"database/sql"
	"github.com/dankru/Auth_service/internal/domain"
)

type Tokens struct {
	db *sql.DB
}

func NewTokens(db *sql.DB) *Tokens {
	return &Tokens{db}
}

func (r *Tokens) Create(token domain.RefreshSession) error {
	return nil
}

func (r *Tokens) Get(token string) (domain.RefreshSession, error) {
	return domain.RefreshSession{}, nil
}
