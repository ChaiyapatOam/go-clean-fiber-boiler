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
	_, err := repo.db.NamedExec("INSERT INTO user (id, email, password, first_name, last_name, provider, profile_url, created_at) VALUES (:id, :email, password, :first_name, :last_name, provider, :profile_url, :created_at)", user)
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

func (repo *userRepository) FindByEmail(email string) (*domain.User, error) {
	user := domain.User{}
	err := repo.db.Get(&user, "SELECT * FROM user WHERE email = ? LIMIT 1", email)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) Update(user *domain.User) error {
	_, err := repo.db.NamedExec("UPDATE user SET first_name = :first_name, last_name = :last_name, phone = :phone, profile_url = :profile_url WHERE id = :id ", user)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) ChangePassword(userId, password string) error {
	_, err := repo.db.Exec("UPDATE user SET password = ? WHERE id = ? ", password, userId)
	if err != nil {
		return err
	}

	return nil
}
