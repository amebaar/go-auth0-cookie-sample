package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController interface {
	List(ctx echo.Context) error
	Create(ctx echo.Context) error
	Detail(ctx echo.Context) error
	Save(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type userController struct{}

func NewUserController() UserController {
	return &userController{}
}

func (u userController) List(ctx echo.Context) error {
	return ctx.String(http.StatusOK, ctx.Get("cid").(string))
}

func (u userController) Create(ctx echo.Context) error {
	return ctx.String(http.StatusOK, ctx.Get("cid").(string))
}

func (u userController) Detail(ctx echo.Context) error {
	return ctx.String(http.StatusOK, ctx.Get("cid").(string))
}

func (u userController) Save(ctx echo.Context) error {
	return ctx.String(http.StatusOK, ctx.Get("cid").(string))
}

func (u userController) Delete(ctx echo.Context) error {
	return ctx.String(http.StatusOK, ctx.Get("cid").(string))
}
