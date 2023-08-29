package repository

import (
	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/models"
	"github.com/jmoiron/sqlx"
)

type SegmentInterface interface {
	CreateSegment(sgmt models.Segment) (int, error)
	GetSegment(sgmt models.Segment) (int, error)
	DeleteSegment(sgmt models.Segment) error
}

type UserInterface interface {
	CreateUser(usr models.User) (int, error)
	GetUser(usr models.User) (int, error)
	DeleteUser(usr models.User) error
}

type ComparisonInterface interface {
	SetUserSegments(sl models.UserSetSegment) error
	GetActiveSegmnents(u models.User) ([]string, error)
}

type AuditInterface interface {
	SendAuditInformation(date string) (string, error)
}

type Repository struct {
	SegmentInterface
	UserInterface
	ComparisonInterface
	AuditInterface
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		SegmentInterface:    NewSegmentPostgres(db),
		UserInterface:       NewUserPostgres(db),
		ComparisonInterface: NewComparisonPostgres(db),
		AuditInterface:      NewAuditStorage(),
	}
}
