package repository

import (
	"context"
	"go-auth0-cookie-sample/api/app/domain/model"
)

type UserRepository interface {
	GetByCompany(ctx context.Context, company *model.Company) (model.Users, error)
	GetByID(ctx context.Context, user *model.User) (model.Users, error)
	Create(ctx context.Context, user *model.User) error
}
