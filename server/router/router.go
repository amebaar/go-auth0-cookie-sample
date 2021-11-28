package router

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions/redis"
	"go-auth0-cookie-sample/auth/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"go-auth0-cookie-sample/auth/authenticator"
	"go-auth0-cookie-sample/handler"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator, store redis.Store) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	router.Use(sessions.Sessions("auth-session", store))
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000", // fixme: 正しいドメインを指定すること
		},
		AllowMethods: []string{
			"GET",
		},
		AllowCredentials: true,
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	router.GET("/login", handler.Login(auth))
	router.GET("/callback", handler.Callback(auth))
	router.GET("/me", middleware.IsAuthenticated, handler.User)
	router.GET("/logout", handler.Logout)
	router.GET("/read", middleware.HasPermissionFuncBuilder("read", "sample"), handler.User)
	router.GET("/create", middleware.HasPermissionFuncBuilder("create", "sample"), handler.User)

	return router
}
