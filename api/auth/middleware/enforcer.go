package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type Enforcer struct {
	enforcer *casbin.Enforcer
}

func NewEnforcer() *Enforcer {
	adaptor, _ := gormadapter.NewAdapter("mysql", "docker:docker@tcp(db:3306)/")
	enforcer, _ := casbin.NewEnforcer("conf/rbac_model.conf", adaptor)
	enforcer.AddNamedDomainMatchingFunc("g", "KeyMatch2", util.KeyMatch2)
	enforcer.BuildRoleLinks()
	enforcer.LoadPolicy()
	return &Enforcer{enforcer: enforcer}
}

func (e *Enforcer) Enforcer() *casbin.Enforcer {
	return e.enforcer
}

func (e *Enforcer) Enforce(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		sess, err := session.Get("session", ctx)
		if err != nil {
			return echo.ErrInternalServerError
		}
		profile := sess.Values["profile"]
		if profile == nil {
			return echo.ErrUnauthorized
		}
		castedProfile := profile.(map[string]interface{})
		roles := castedProfile["http://schemas.microsoft.com/ws/2008/06/identity/claims/role"].([]interface{})
		cid := castedProfile["http://schemas.xmlsoap.org/claims/Group"].(string)

		method := ctx.Request().Method
		path := ctx.Request().URL.Path
		ctx.Logger().Infof("%v, %s, %s, %s", roles, "cid"+cid, path, method)

		hasGrants := false
		for _, r := range roles {
			result, _ := e.enforcer.Enforce(r, "cid"+cid, path, method)
			if result {
				ctx.Logger().Infof("%s: %s, %s, %s, %s", result, r, "cid"+cid, path, method)
				hasGrants = true
			}
		}
		if hasGrants {
			ctx.Set("cid", cid)
			return next(ctx)
		}

		return echo.ErrForbidden // 403
	}
}
