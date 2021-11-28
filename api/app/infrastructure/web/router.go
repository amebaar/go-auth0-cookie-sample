package web

import (
	"github.com/boj/redistore"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"go-auth0-cookie-sample/api/app/adapter/web/controller"
	auth "go-auth0-cookie-sample/api/auth/middleware"
)

func Start(
	authController controller.AuthController,
	policyController controller.PolicyController,
	userController controller.UserController,
	enforcer *auth.Enforcer,
) {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Session Setup
	store, err := redistore.NewRediStore(10, "tcp", "redis:6379", "", []byte("secret"))
	if err != nil {
		panic(err)
	}
	defer store.Close()
	e.Use(session.Middleware(sessions.Store(store)))

	e.POST("/policies/init", func(c echo.Context) error { return policyController.Init(c) })
	e.GET("/policies", func(c echo.Context) error { return policyController.List(c) })

	e.GET("/login", func(c echo.Context) error { return authController.Login(c) })
	e.GET("/callback", func(c echo.Context) error { return authController.Callback(c) })
	e.GET("/logout", func(c echo.Context) error { return authController.Logout(c) })

	user := e.Group("/users")
	user.Use(enforcer.Enforce)

	user.GET("", func(c echo.Context) error { return userController.List(c) })
	user.POST("", func(c echo.Context) error { return userController.Create(c) })

	user.GET("/:id", func(c echo.Context) error { return userController.Detail(c) })
	user.PUT("/:id", func(c echo.Context) error { return userController.Save(c) })
	user.DELETE("/:id", func(c echo.Context) error { return userController.Delete(c) })

	me := e.Group("/me")
	me.Use(enforcer.Enforce)

	me.GET("", func(c echo.Context) error { return userController.Detail(c) })
	me.PUT("", func(c echo.Context) error { return userController.Save(c) })

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
