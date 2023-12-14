package middleware

import (
	"net/http"
	"path"
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
	if auth.IsNeedValidate(ctx) {
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
	} else {
		bearerToken := strings.Split(c.Get("Authorization"), " ")
		if len(bearerToken) > 1 {
			token.AccessToken.Set(bearerToken[1])
		}

		if token.AccessToken.String != "" {
			//check access token
			tx, err := ctx.DB()
			if err != nil {
				return app.Error().New(http.StatusInternalServerError, err.Error())
			}

			err = tx.Table("access_token").
				Where("expired_at > ?", time.Now()).
				Take(&token).Error

			if err == nil {
				// get user
				err = tx.Table("m_user").
					Where("id_user", token.UserId.Int64).
					Where("deleted_at IS NULL").
					Take(&user).Error
			}

			ctx.Token = token
			ctx.User = user
		}

	}

	return c.Next()
}

func (auth *authHandler) IsNeedValidate(ctx *app.Ctx) bool {
	urlPath := auth.FiberCtx.Path()
	method := auth.FiberCtx.Method()
	cleanedPath := path.Clean(urlPath)

	ctx.Action.Method = method

	// Split the cleaned path into segments
	segments := strings.Split(cleanedPath, "/")

	if len(segments) >= 3 {
		if segments[1] != "api" || segments[2] == "docs" || segments[2] == "storages" {
			ctx.Action.EndPoint = segments[3]
			return false
		}

		if len(segments) >= 4 {
			ctx.Action.EndPoint = segments[3]
			if len(segments) >= 5 {
				ctx.Action.DataID = segments[4]
			}

			//get tanpa login
			if segments[3] == "pengetahuan" && method == "GET" {
				return false
			} else if segments[3] == "komentar" && method == "GET" {
				return false
			} else if segments[3] == "search_pengetahuan" && method == "GET" {
				return false
			} else if segments[3] == "events" && method == "GET" {
				return false
			} else if segments[3] == "event_materi" && method == "GET" {
				return false
			} else if segments[3] == "leader_talk" && method == "GET" {
				return false
			} else if segments[3] == "search_forum" && method == "GET" {
				return false
			} else if segments[3] == "forum" && method == "GET" {
				return false
			} else if segments[3] == "user" && (method == "GET" || method == "POST") {
				return false
			} else if segments[3] == "jenis_pengetahuan" && method == "GET" {
				return false
			} else if segments[3] == "kelompok_dokumen" && method == "GET" {
				return false
			} else if segments[3] == "attachment" && method == "GET" {
				return false
			} else if segments[3] == "like" && method == "GET" {
				return false
			} else if segments[3] == "dislike" && method == "GET" {
				return false
			} else if segments[3] == "library_cafe" && method == "GET" {
				return false
			} else if segments[3] == "kategori_pengetahuan" && method == "GET" {
				return false
			} else if segments[3] == "dokumen" && method == "GET" {
				return false
			} else if segments[3] == "kategori_buku" && method == "GET" {
				return false
			} else if segments[3] == "elibrary" && method == "GET" {
				return false
			} else if (segments[3] == "slider_pengetahuan" || segments[3] == "mix_slider") && method == "GET" {
				return false
			} else if segments[3] == "advis_kategori" && method == "GET" {
				return false
			} else if (segments[3] == "advis_kategori" || segments[3] == "advis_sub_kategori" || segments[3] == "advis_sumber_data") && method == "GET" {
				return false
			}

			if (segments[3] == "advis_analytics" || segments[3] == "advis_list_data") && method == "GET" && (ctx.Action.DataID == "" || ctx.Action.DataID == "template_csv") {
				return false
			}

			switch segments[3] {
			case "login", "version", "docs", "caches":
				return false
			}
		}
	}
	return true
}
