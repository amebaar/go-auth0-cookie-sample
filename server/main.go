package main

import (
	"github.com/gin-contrib/sessions/redis"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"go-auth0-cookie-sample/auth/authenticator"
	"go-auth0-cookie-sample/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	store, err := redis.NewStore(10, "tcp", "redis:6379", "", []byte("secret"))
	if err != nil {
		log.Fatalf("Failed to initialize session storage: %v", err)
	}
	/* fixme: 適切なCookie設定を行うこと
	store.Options(
		sessions.Options{
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})

	*/

	rtr := router.New(auth, store)

	log.Print("Server listening on http://localhost:8080/")
	if err := http.ListenAndServe("0.0.0.0:8080", rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
