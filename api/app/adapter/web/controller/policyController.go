package controller

import (
	"github.com/labstack/echo/v4"
	authMiddleware "go-auth0-cookie-sample/api/auth/middleware"
	"net/http"
)

type PolicyController interface {
	Init(ctx echo.Context) error
	List(ctx echo.Context) error
}

type policyController struct {
	enforcer *authMiddleware.Enforcer
}

func NewPolicyController(enforcerMiddleware *authMiddleware.Enforcer) PolicyController {
	return &policyController{enforcer: enforcerMiddleware}
}

func (c *policyController) Init(ctx echo.Context) error {
	e := c.enforcer.Enforcer()

	companies := []string{"cid1", "cid2"}
	for _, comp := range companies {
		e.AddPolicy("admin", comp, "/users/:id", "*")
		e.AddPolicy("admin", comp, "/users", "*")
		e.AddPolicy("admin", comp, "/me", "GET")
		e.AddPolicy("admin", comp, "/me", "PUT")

		e.AddPolicy("viewer", comp, "/users", "GET")
		e.AddPolicy("viewer", comp, "/me", "GET")
		e.AddPolicy("viewer", comp, "/me", "PUT")
	}

	return e.SavePolicy()
}

func (c *policyController) List(ctx echo.Context) error {
	e := c.enforcer.Enforcer()
	return ctx.JSON(http.StatusOK, e.GetPolicy())
}
