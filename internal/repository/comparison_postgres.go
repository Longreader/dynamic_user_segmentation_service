package repository

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/models"
	"github.com/jmoiron/sqlx"
)

type ComparisonPostgres struct {
	db *sqlx.DB
}

func NewComparisonPostgres(db *sqlx.DB) *ComparisonPostgres {
	return &ComparisonPostgres{db: db}
}

func (c *ComparisonPostgres) CreateSegmentIfNotExist(seg string) (int, error) {
	// Проверка существования сегмента
	var sgmtIDOut int
	query := fmt.Sprintf("SELECT id FROM %s WHERE segment=$1", segmentsTable)
	row := c.db.QueryRow(query, seg)
	if err := row.Scan(&sgmtIDOut); err != nil {
		if err.Error() == "sql: no rows in result set" {
			// Создание сегмента
			query := fmt.Sprintf("INSERT INTO %s (segment) values ($1) RETURNING id", segmentsTable)
			row := c.db.QueryRow(query, seg)
			if err := row.Scan(&sgmtIDOut); err != nil {
				// Если при создании произошла ошибка
				// прекращаем исполнение программы
				return 0, err
			} else {
				// Объект создан. Возвращаем ID
				return sgmtIDOut, nil
			}
		} else {
			// Ошибка не в существующем сегменте
			return 0, err
		}
	} else {
		// Сегмент существует
		return sgmtIDOut, nil
	}
}

func (c *ComparisonPostgres) GetSegmentIfExist(seg string) (int, error) {
	// Проверка существования сегмента
	var sgmtIDOut int
	query := fmt.Sprintf("SELECT id FROM %s WHERE segment=$1", segmentsTable)
	row := c.db.QueryRow(query, seg)
	if err := row.Scan(&sgmtIDOut); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return 0, nil
		}
		return 0, err
	} else {
		// Сегмент существует
		return sgmtIDOut, nil
	}
}

func (c *ComparisonPostgres) GetSegmentNameIfExist(seg int) (string, error) {
	// Проверка существования сегмента
	var sgmtOut string
	query := fmt.Sprintf("SELECT segment FROM %s WHERE id=$1", segmentsTable)
	row := c.db.QueryRow(query, seg)
	if err := row.Scan(&sgmtOut); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return "", nil
		}
		return "", err
	} else {
		// Сегмент существует
		return sgmtOut, nil
	}
}

func (c *ComparisonPostgres) CreateUserIfNotExist(usr int) (int, error) {
	var usrIDOut int
	query := fmt.Sprintf("SELECT id FROM %s WHERE user_id=$1", usersTable)
	row := c.db.QueryRow(query, usr)
	if err := row.Scan(&usrIDOut); err != nil {
		if err.Error() == "sql: no rows in result set" {
			query := fmt.Sprintf("INSERT INTO %s (user_id) values ($1) RETURNING id", usersTable)
			row := c.db.QueryRow(query, usr)
			if err := row.Scan(&usrIDOut); err != nil {
				return 0, err
			} else {
				return usrIDOut, nil
			}
		} else {
			// Ошибка не в отсутствии User
			return 0, err
		}
	} else {
		return usrIDOut, nil
	}
}

func (c *ComparisonPostgres) GetUserIfExist(usr int) (int, error) {
	var usrIDOut int
	query := fmt.Sprintf("SELECT id FROM %s WHERE user_id=$1", usersTable)
	row := c.db.QueryRow(query, usr)
	if err := row.Scan(&usrIDOut); err != nil {
		return 0, err
	} else {
		return usrIDOut, nil
	}
}

func (c *ComparisonPostgres) CreateComparisonIfNotExist(seg int, usr int) (int, error) {
	var comparisonID int
	query := fmt.Sprintf("SELECT id FROM %s WHERE segment_id=$1 AND user_id=$2", comparisonTable)
	row := c.db.QueryRow(query, seg, usr)
	if err := row.Scan(&comparisonID); err != nil {
		if err.Error() == "sql: no rows in result set" {
			query := fmt.Sprintf("INSERT INTO %s (user_id, segment_id) values ($1, $2) RETURNING id", comparisonTable)
			row := c.db.QueryRow(query, usr, seg)
			if err := row.Scan(&comparisonID); err != nil {
				// Ошибка создания
				return 0, err
			} else {
				// Запись создана
				return comparisonID, nil
			}
		} else {
			// Ошибка не связанная с отсутствием записи
			return 0, err
		}
	}
	// Запись была создана ранее
	return comparisonID, nil
}

func (c *ComparisonPostgres) GetComparisonIfExist(seg int, usr int) (int, error) {
	// Проверка существования сегмента
	var compIDOut int
	query := fmt.Sprintf(`SELECT id FROM %s WHERE user_id=$1 AND segment_id=$2`, comparisonTable)
	err := c.db.Get(&compIDOut, query, usr, seg)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return 0, nil
		}
		return 0, err
	} else {
		// Сегмент существует
		return compIDOut, nil
	}
}

func (c *ComparisonPostgres) DeleteComparisonIfExist(com int, seg int, usr int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1 AND segment_id=$2 AND user_id=$3`, comparisonTable)
	fmt.Println("we gere")
	res, err := c.db.Exec(query, com, seg, usr)
	fmt.Println(res, err)
	return err
}

func (c *ComparisonPostgres) SetUserSegments(uss models.UserSetSegment) error {

	var toSet []string
	var toDelete []string
	var userToSet int
	var userID int
	var setID int
	var delID int
	var compID int
	var err error

	toSet = uss.SegmentsSet
	toDelete = uss.SegmentsDelete
	userToSet = uss.UserId

	userID, err = c.CreateUserIfNotExist(userToSet)
	if err != nil {
		return err
	}

	log.Debug("CreateUserIfNotExist work")

	if len(toSet) != 0 {
		for _, seg := range toSet {
			setID, err = c.CreateSegmentIfNotExist(seg)
			if err != nil {
				return err
			}
			log.Debug("CreateSegmentIfNotExist work")

			_, err = c.CreateComparisonIfNotExist(setID, userID)
			if err != nil {
				return err
			}
			log.Debug("CreateComparisonIfNotExist work")

		}
	}
	if len(toDelete) != 0 {
		for _, seg := range toDelete {
			delID, err = c.GetSegmentIfExist(seg)
			if err != nil {
				return err
			}

			compID, err = c.GetComparisonIfExist(delID, userID)
			if err != nil {
				return err
			}

			log.Info("DELETE", delID, compID, userToSet)

			err = c.DeleteComparisonIfExist(compID, delID, userID)
			if err != nil {
				return err
			}
			log.Debug("CreateComparisonIfNotExist work")
		}
	}
	return nil
}

func (c *ComparisonPostgres) GetActiveSegmnents(usr models.User) ([]string, error) {

	var usrID int
	var sgmtID int
	// var sgmt string
	var sgmts []string
	var err error

	usrID, err = c.GetUserIfExist(usr.UserId)

	if err != nil {
		return make([]string, 0), nil
	}

	query := fmt.Sprintf(`SELECT segment_id FROM %s WHERE user_id=$1`, comparisonTable)

	rows, err := c.db.Query(query, usrID)

	if err != nil {
		return make([]string, 0), err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&sgmtID); err != nil {
			return sgmts, err
		}
		sgmt, err := c.GetSegmentNameIfExist(sgmtID)

		if err != nil {
			return sgmts, err
		}

		sgmts = append(sgmts, sgmt)
	}
	if err = rows.Err(); err != nil {
		return sgmts, err
	}
	return sgmts, err
}
