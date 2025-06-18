package repository

import (
	"fmt"
	"math"

	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u *UserPostgres) CountUsers() (int, error) {
	var count int
	query := fmt.Sprintf(`SELECT count(*) FROM  %s`, usersTable)
	row := u.db.QueryRow(query)
	if err := row.Scan(&count); err != nil {
		fmt.Println(count)
		return 0, err
	}
	return count, nil
}

func (u *UserPostgres) GetRandomUsers(count int) ([]int, error) {

	var usrs []int
	var usrID int

	query := fmt.Sprintf(`SELECT user_id FROM %s ORDER BY RANDOM() LIMIT $1;`, usersTable)

	rows, err := u.db.Query(query, count)

	if err != nil {
		return make([]int, 0), err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&usrID); err != nil {
			return usrs, err
		}

		usrs = append(usrs, usrID)
	}
	if err = rows.Err(); err != nil {
		return usrs, err
	}
	return usrs, err
}

func (u *UserPostgres) CreateUser(usr models.User) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (user_id) values ($1) RETURNING id`, usersTable)
	row := u.db.QueryRow(query, usr.UserID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (u *UserPostgres) GetUser(usr models.User) (int, error) {
	var usrIdOut int
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE user_id=$1", usersTable)
	row := u.db.QueryRow(query, usr.UserID)
	if err := row.Scan(&usrIdOut); err != nil {
		return 0, err
	}
	return usrIdOut, nil
}

func (u *UserPostgres) DeleteUser(usr models.User) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1", usersTable)
	_ = u.db.QueryRow(query, usr.UserID)
	return nil
}

func (u *UserPostgres) GetRandUsers(s models.Segment) ([]int, error) {

	var usrs []int
	var countUsrs int
	var total float64
	var err error

	countUsrs, err = u.CountUsers()
	if err != nil {
		return make([]int, 0), err
	}

	total = float64(countUsrs) * (float64(s.Percent) / float64(100))

	usrs, err = u.GetRandomUsers(int(math.Round(total)))

	if err != nil {
		return make([]int, 0), err
	}

	return usrs, nil
}
