package controller

import (
	"encoding/base64"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"go-auth0-cookie-sample/api/auth/authenticator"
	"math/rand"
	"net/http"
	"net/url"
	"os"
)

type AuthController interface {
	Login(ctx echo.Context) error
	Callback(ctx echo.Context) error
	Logout(ctx echo.Context) error
}

type authController struct {
	authenticator *authenticator.Authenticator
}

func NewAuthController(authenticator *authenticator.Authenticator) AuthController {
	return &authController{
		authenticator: authenticator,
	}
}

func (c *authController) Login(ctx echo.Context) error {
	state, err := generateRandomState()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	// Save the state inside the session. [COOKIE]
	sess, _ := session.Get("session", ctx)
	sess.Values["state"] = state
	if err := sess.Save(ctx.Request(), ctx.Response()); err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.Redirect(http.StatusTemporaryRedirect, c.authenticator.AuthCodeURL(state))
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}

func (c *authController) Logout(ctx echo.Context) error {
	logoutUrl, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/v2/logout")
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	returnTo, err := url.Parse("http://localhost:3000")
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	logoutUrl.RawQuery = parameters.Encode()

	sess, _ := session.Get("session", ctx)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	sess.Values = nil
	sess.Save(ctx.Request(), ctx.Response())
	return ctx.Redirect(http.StatusTemporaryRedirect, logoutUrl.String())
}

func (c *authController) Callback(ctx echo.Context) error {
	sess, err := session.Get("session", ctx)
	if err != nil || ctx.QueryParam("state") != sess.Values["state"] {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("Invalid state parameter. %#v, %#v", ctx.QueryParam("state"), sess.Values["state"]))
	}

	// Exchange an authorization code for a token.
	token, err := c.authenticator.Exchange(ctx.Request().Context(), ctx.QueryParam("code"))
	if err != nil {
		return ctx.String(http.StatusUnauthorized, "Failed to convert an authorization code into a token.")
	}

	idToken, err := c.authenticator.VerifyIDToken(ctx.Request().Context(), token)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to verify ID Token.")
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to get claims: "+err.Error())
	}

	sess.Values["access_token"] = token.AccessToken
	sess.Values["profile"] = profile
	if err := sess.Save(ctx.Request(), ctx.Response()); err != nil {
		ctx.Logger().Errorf("%+v", err)
		return ctx.String(http.StatusInternalServerError, "Failed to save session: "+err.Error())
	}

	// Redirect to logged in page.
	return ctx.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000")
}
