package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-auth0-cookie-sample/auth/policy"
	"net/http"
)

func HasPermissionFuncBuilder(operation string, resource string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		profile := sessions.Default(ctx).Get("profile")
		if profile == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized) // 401
		}
		castedProfile := profile.(map[string]interface{})
		roles := castedProfile["http://schemas.microsoft.com/ws/2008/06/identity/claims/role"].([]interface{})

		hasGrants := false
		for _, r := range roles {
			if policy.GetManager().HasPermission(r.(string), operation, resource) {
				hasGrants = true
				break
			}
		}
		if hasGrants {
			ctx.Next()
		} else {
			ctx.AbortWithStatus(http.StatusForbidden) // 403
		}
	}
}
