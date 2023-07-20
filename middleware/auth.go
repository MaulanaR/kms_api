package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/maulanar/kms/app"
)

func Auth() *authHandler {
	if auth == nil {
		auth = &authHandler{}
	}
	return auth
}

var auth *authHandler

type authHandler struct {
	FiberCtx *fiber.Ctx
}

func (auth *authHandler) ValidateAuth(c *fiber.Ctx) error {
	auth.FiberCtx = c
	ctx, ok := c.Locals(app.CtxKey).(*app.Ctx)
	if !ok {
		return app.Error().New(http.StatusInternalServerError, "ctx is not found")
	}

	token := app.Token{}
	user := app.User{}
	if auth.IsNeedValidate() {
		bearerToken := strings.Split(c.Get("Authorization"), " ")
		if len(bearerToken) > 1 {
			token.AccessToken.Set(bearerToken[1])
		} else {
			return app.Error().New(http.StatusUnauthorized, "Access token is require.")
		}

		if token.AccessToken.String == "" {
			return app.Error().New(http.StatusUnauthorized, "Access token is invalid.")
		}

		//check access token
		tx, err := ctx.DB()
		if err != nil {
			return app.Error().New(http.StatusInternalServerError, err.Error())
		}

		err = tx.Table("access_token").
			Where("expired_at > ?", time.Now()).
			Take(&token).Error

		if err != nil {
			return app.Error().New(http.StatusUnauthorized, "Access token is invalid/ expired.")
		}

		// get user
		err = tx.Table("m_user").
			Where("id_user", token.UserId.Int64).
			Where("deleted_at IS NULL").
			Take(&user).Error

		if err != nil {
			return app.Error().New(http.StatusUnauthorized, "User Not Found.")
		}

		ctx.Token = token
		ctx.User = user
	}

	return c.Next()
}

func (auth *authHandler) IsNeedValidate() bool {
	path := auth.FiberCtx.Path()
	// method := auth.FiberCtx.Method()

	// legacy
	switch path {
	case "/api/login", "/api/version":
		return false
	}

	// // check login, login pake jwt
	// if (path == "/api/v1/session" || path == "/api/v1/session_std") && method != "PATCH" {
	// 	return false
	// }

	// proxy pake auth dari url tujuan
	// if strings.HasPrefix(path, "/api/v2") {
	// 	return false
	// }
	return true
}
