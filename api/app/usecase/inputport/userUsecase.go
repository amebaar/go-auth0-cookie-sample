package inputport

import (
	"github.com/labstack/echo/v4"
	"go-auth0-cookie-sample/api/app/domain/model"
)

type UserUsecase interface {
	List(ctx echo.Context) (model.Users, error)
}
