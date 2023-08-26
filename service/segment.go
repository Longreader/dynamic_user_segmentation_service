package service

import (
	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/models"
	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/repository"
)

type SegmentService struct {
	repo repository.SegmentInterface
}

func NewSegmentService(repo repository.SegmentInterface) *SegmentService {
	return &SegmentService{repo: repo}
}

func (s *SegmentService) CreateSegment(sgmt models.Segment) (int, error) {
	return s.repo.CreateSegment(sgmt)
}

func (s *SegmentService) GetSegment(sgmt models.Segment) (string, error) {
	return s.repo.GetSegment(sgmt)
}

func (s *SegmentService) DeleteSegment(sgmt models.Segment) error {
	return s.repo.DeleteSegment(sgmt)
}
