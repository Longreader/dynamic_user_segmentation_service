package repository

import (
	"fmt"

	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/models"
	"github.com/jmoiron/sqlx"
)

type SegmentPostgres struct {
	db *sqlx.DB
}

func NewSegmentPostgres(db *sqlx.DB) *SegmentPostgres {
	return &SegmentPostgres{db: db}
}

func (s *SegmentPostgres) CreateSegment(sgmt models.Segment) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (segment) values ($1) RETURNING id", segmentsTable)
	row := s.db.QueryRow(query, sgmt.Segment)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *SegmentPostgres) GetSegment(sgmt models.Segment) (int, error) {
	var sgmtIDOut int
	query := fmt.Sprintf("SELECT id FROM %s WHERE segment=$1", segmentsTable)
	row := s.db.QueryRow(query, sgmt.Segment)
	if err := row.Scan(&sgmtIDOut); err != nil {
		return 0, err
	}
	return sgmtIDOut, nil
}

func (s *SegmentPostgres) DeleteSegment(sgmt models.Segment) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE segment=$1", segmentsTable)
	_ = s.db.QueryRow(query, sgmt.Segment)
	return nil
}
