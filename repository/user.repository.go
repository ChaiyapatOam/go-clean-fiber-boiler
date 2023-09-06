package repository

import (
	"database/sql"

	"github.com/chaiyapatoam/go-clean-fiber-boiler/domain"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewuserRepository(db *sqlx.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) Create(user *domain.User) error {
	_, err := repo.db.NamedExec("INSERT INTO user VALUES (:id, :email, :first_name, :last_name, :profile_url, :created_at)", user)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) Get(id string) (*domain.User, error) {
	user := domain.User{}
	err := repo.db.Get(&user, "SELECT * FROM user WHERE id = ? LIMIT 1", id)
	
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}
