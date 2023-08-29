package service

import (
	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/models"
	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/repository"
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
	SetUserSegments(uss models.UserSetSegment) error
	GetActiveSegmnents(u models.User) ([]string, error)
}

type AuditInterface interface {
	SendAuditInformation(date string) (string, error)
}

type Service struct {
	SegmentInterface
	UserInterface
	ComparisonInterface
	AuditInterface
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		SegmentInterface:    NewSegmentService(repo.SegmentInterface),
		UserInterface:       NewUserService(repo.UserInterface),
		ComparisonInterface: NewComparisonService(repo.ComparisonInterface),
		AuditInterface:      NewAuditService(repo.AuditInterface),
	}
}
