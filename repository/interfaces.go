package repository

import (
	"context"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/model"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, user model.User) (*model.User, error)
	Update(ctx context.Context, userId int, user generated.ProfileUpdateParams) error
	UpdateLoginSuccess(ctx context.Context, userId int) error
	FindBy(ctx context.Context, column string, value any) (*model.User, error)
}
