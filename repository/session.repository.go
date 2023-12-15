package repository

import (
	"database/sql"

	"github.com/chaiyapatoam/go-clean-fiber-boiler/domain"
	"github.com/jmoiron/sqlx"
)

type sessionRepository struct {
	db *sqlx.DB
}

func NewSessionRepository(db *sqlx.DB) domain.SessionRepository {
	return &sessionRepository{db: db}
}

func (repo *sessionRepository) Create(session *domain.Session) error {
	_, err := repo.db.NamedExec("INSERT INTO session (id, user_id, ip_address, expired_at, created_at) VALUES (:id, :user_id, :ip_address, :expired_at, :created_at)", session)
	if err != nil {
		return err
	}
	return nil
}

func (repo *sessionRepository) Get(ssid string) (*domain.Session, error) {
	session := domain.Session{}
	err := repo.db.Get(&session, "SELECT * FROM session WHERE id = ? LIMIT 1", ssid)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &session, nil
}

func (repo *sessionRepository) Delete(id string) error {
	_, err := repo.db.Exec("DELETE FROM session WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
