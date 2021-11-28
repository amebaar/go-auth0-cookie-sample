package main

import (
	"encoding/gob"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go-auth0-cookie-sample/api/app/adapter/web/controller"
	"go-auth0-cookie-sample/api/app/infrastructure/web"
	"go-auth0-cookie-sample/api/auth/authenticator"
	authMiddleware "go-auth0-cookie-sample/api/auth/middleware"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}
	gob.Register(map[string]interface{}{})

	// DI
	au, err := authenticator.New()
	if err != nil {
		panic(err)
	}
	en := authMiddleware.NewEnforcer()
	ac := controller.NewAuthController(au)
	pc := controller.NewPolicyController(en)
	uc := controller.NewUserController()

	web.Start(ac, pc, uc, en)
}
