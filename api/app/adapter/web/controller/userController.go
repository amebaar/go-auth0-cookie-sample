package controller

import (
	"github.com/labstack/echo/v4"
	"go-auth0-cookie-sample/api/app/usecase/inputport"
	"net/http"
)

type UserController interface {
	List(ctx echo.Context) error
	Create(ctx echo.Context) error
	Detail(ctx echo.Context) error
	Save(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type userController struct {
	userUsecase inputport.UserUsecase
}

func NewUserController() UserController {
	return &userController{}
}

func (u userController) List(ctx echo.Context) error {
	users, err := u.userUsecase.List(ctx)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return ctx.JSON(http.StatusOK, users)
}

func (u userController) Create(ctx echo.Context) error {
	return ctx.String(http.StatusOK, ctx.Get("cid").(string))
}

func (u *userController) Detail(ctx echo.Context) error {
	return ctx.String(http.StatusOK, ctx.Get("cid").(string))
}

func (u userController) Save(ctx echo.Context) error {
	return ctx.String(http.StatusOK, ctx.Get("cid").(string))
}

func (u userController) Delete(ctx echo.Context) error {
	return ctx.String(http.StatusOK, ctx.Get("cid").(string))
}
