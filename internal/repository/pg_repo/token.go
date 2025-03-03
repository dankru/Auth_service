package pg_repo

import (
	"database/sql"
	"github.com/dankru/Auth_service/internal/domain"
)

type TokensRepository struct {
	db *sql.DB
}

func NewTokensRepository(db *sql.DB) *TokensRepository {
	return &TokensRepository{db}
}

func (r *TokensRepository) Create(token domain.RefreshSession) error {
	_, err := r.db.Exec("INSERT INTO refresh_tokens (user_id, token, expires_at) values ($1, $2, $3)",
		token.UserID, token.Token, token.ExpiresAt)

	return err
}

func (r *TokensRepository) Get(token string) (domain.RefreshSession, error) {
	return domain.RefreshSession{}, nil
}
