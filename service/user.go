package service

import (
	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/models"
	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/repository"
)

type UserService struct {
	repo repository.UserInterface
}

func NewUserService(repo repository.UserInterface) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) CreateUser(usr models.User) (int, error) {
	return u.repo.CreateUser(usr)
}

func (u *UserService) GetUser(usr models.User) (int, error) {
	return u.repo.GetUser(usr)
}

func (u *UserService) DeleteUser(usr models.User) error {
	return u.repo.DeleteUser(usr)
}
