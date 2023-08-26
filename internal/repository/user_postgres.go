package repository

import (
	"fmt"

	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u *UserPostgres) CreateUser(usr models.User) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (user_id) values ($1) RETURNING id`, usersTable)
	row := u.db.QueryRow(query, usr.UserId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (u *UserPostgres) GetUser(usr models.User) (int, error) {
	var usrIdOut int
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE user_id=$1", usersTable)
	row := u.db.QueryRow(query, usr.UserId)
	if err := row.Scan(&usrIdOut); err != nil {
		return 0, err
	}
	return usrIdOut, nil
}

func (u *UserPostgres) DeleteUser(usr models.User) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1", usersTable)
	_ = u.db.QueryRow(query, usr.UserId)
	return nil
}
