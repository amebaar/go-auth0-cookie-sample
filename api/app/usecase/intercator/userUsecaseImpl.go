package intercator

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-auth0-cookie-sample/api/app/domain/model"
	"go-auth0-cookie-sample/api/app/domain/repository"
	"go-auth0-cookie-sample/api/auth/session"
)

type userController struct {
	repo repository.UserRepository
}

func (c *userController) List(ctx echo.Context) (model.Users, error) {
	company, err := session.GetUserCompany(ctx)
	if err != nil {
		return nil, fmt.Errorf("faield to get user company info: %+v", err)
	}

	return c.repo.GetByCompany(
		ctx.Request().Context(),
		company,
	)
}
