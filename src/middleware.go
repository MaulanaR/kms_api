package src

import (
	"github.com/maulanar/kms/app"
	"github.com/maulanar/kms/middleware"
)

func Middleware() *middlewareUtil {
	if mdlwr == nil {
		mdlwr = &middlewareUtil{}
		mdlwr.Configure()
		mdlwr.isConfigured = true
	}
	return mdlwr
}

var mdlwr *middlewareUtil

type middlewareUtil struct {
	isConfigured bool
}

func (*middlewareUtil) Configure() {
	app.Server().AddMiddleware(middleware.Ctx().New)
	app.Server().AddMiddleware(middleware.DB().New)
	app.Server().AddMiddleware(middleware.Auth().ValidateAuth)
}
