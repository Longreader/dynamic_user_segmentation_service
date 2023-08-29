package service

import (
	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/repository"
)

type AuditService struct {
	repo repository.AuditInterface
}

func NewAuditService(repo repository.AuditInterface) *AuditService {
	return &AuditService{repo: repo}
}

func (a *AuditService) SendAuditInformation(date string) (string, error) {
	return a.repo.SendAuditInformation(date)
}
