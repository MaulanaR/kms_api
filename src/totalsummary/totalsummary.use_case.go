package totalsummary

import (
	"net/http"
	"net/url"

	"github.com/maulanar/kms/app"
)

func UseCase(ctx app.Ctx, query ...url.Values) UseCaseHandler {
	u := UseCaseHandler{
		Ctx:   &ctx,
		Query: url.Values{},
	}
	if len(query) > 0 {
		u.Query = query[0]
	}
	return u
}

type UseCaseHandler struct {
	TotalSummary

	Ctx   *app.Ctx   `json:"-" db:"-" gorm:"-"`
	Query url.Values `json:"-" db:"-" gorm:"-"`
}

func (u UseCaseHandler) Async(ctx app.Ctx, query ...url.Values) UseCaseHandler {
	ctx.IsAsync = true
	return UseCase(ctx, query...)
}

func (u UseCaseHandler) Get() (TotalSummary, error) {
	res := TotalSummary{}

	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	err = app.Query().First(tx, &res, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	return res, err
}
