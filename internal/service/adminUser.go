package service

import (
	"context"
	"errors"
	"log"
	"ordernationn/internal/domain"
)

type AdminUserService interface {
	AdminLogin(context.Context, string, string) (domain.AdminUser, error)
}

type adminUserSerivce struct {
	repo domain.AdminUserRepository
}

func NewAdminUserService(r domain.AdminUserRepository) AdminUserService {
	return &adminUserSerivce{
		repo: r,
	}
}

func (s *adminUserSerivce) AdminLogin(ctx context.Context, email string, password string) (domain.AdminUser, error) {
	if email == "" || password == "" {
		return domain.AdminUser{}, errors.New("email and password are required")
	}
	user, err := s.repo.Login(ctx, email, password)
	if user.ID == 0 {
		log.Println(err)
		return domain.AdminUser{}, errors.New("no admin found")
	}
	return user, nil
}
