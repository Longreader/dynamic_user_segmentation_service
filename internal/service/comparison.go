package service

import (
	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/models"
	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/repository"
)

type ComparisonService struct {
	repo repository.ComparisonInterface
}

func NewComparisonService(repo repository.ComparisonInterface) *ComparisonService {
	return &ComparisonService{repo: repo}
}

func (c *ComparisonService) SetUserSegments(uss models.UserSetSegment) error {
	return c.repo.SetUserSegments(uss)
}

func (c *ComparisonService) GetActiveSegments(u models.User) ([]string, error) {
	return c.repo.GetActiveSegments(u)
}
