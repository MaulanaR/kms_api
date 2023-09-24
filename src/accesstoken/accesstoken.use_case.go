package accesstoken

import (
	"net/http"
	"net/url"
	"time"

	"github.com/maulanar/kms/app"
	"github.com/maulanar/kms/src/user"
)

// UseCase returns a UseCaseHandler for expected use case functional.
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

// UseCaseHandler provides a convenient interface for AccessToken use case, use UseCase to access UseCaseHandler.
type UseCaseHandler struct {
	AccessToken

	// injectable dependencies
	Ctx   *app.Ctx   `json:"-" db:"-" gorm:"-"`
	Query url.Values `json:"-" db:"-" gorm:"-"`
}

// Async return UseCaseHandler with async process.
func (u UseCaseHandler) Async(ctx app.Ctx, query ...url.Values) UseCaseHandler {
	ctx.IsAsync = true
	return UseCase(ctx, query...)
}

// GetByID returns the AccessToken data for the specified ID.
func (u UseCaseHandler) GetByID(id string) (AccessToken, error) {
	res := AccessToken{}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// get from db
	key := "access_token"
	if id == "" {
		id = u.Ctx.Token.AccessToken.String
	}
	u.Query.Add(key, id)
	err = app.Query().First(tx, &res, u.Query)
	if err != nil {
		return res, u.Ctx.NotFoundError(err, u.EndPoint(), key, id)
	}

	return res, err
}

// Get returns the list of AccessToken data.
func (u UseCaseHandler) Get() (app.ListModel, error) {
	res := app.ListModel{}

	// check permission
	err := u.Ctx.ValidatePermission("access_token.list")
	if err != nil {
		return res, err
	}
	// get from cache and return if exists
	cacheKey := u.EndPoint() + "?" + u.Query.Encode()
	err = app.Cache().Get(cacheKey, &res)
	if err == nil {
		return res, err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// set pagination info
	res.Count,
		res.PageContext.Page,
		res.PageContext.PerPage,
		res.PageContext.PageCount,
		err = app.Query().PaginationInfo(tx, &AccessToken{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	// return data count if $per_page set to 0
	if res.PageContext.PerPage == 0 {
		return res, err
	}

	// find data
	data, err := app.Query().Find(tx, &AccessToken{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	res.SetData(data, u.Query)

	return res, err
}

// Create creates a new data AccessToken with specified parameters.
func (u UseCaseHandler) Login(p *ParamCreate) (AccessToken, error) {
	res := AccessToken{}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// validasi data
	if p.Password.String == "" {
		return res, app.Error().New(http.StatusUnauthorized, "Password wajib diisikan", map[string]string{"Password": "Password is required"})
	}
	if p.Username.String == "" {
		return res, app.Error().New(http.StatusUnauthorized, "Username wajib diisikan", map[string]string{"Password": "Password is required"})
	}

	//query db
	dataUser := user.User{}
	err = tx.Table("m_user").Where("username", p.Username.String).First(&dataUser).Error
	if err != nil {
		return res, app.Error().New(http.StatusBadRequest, app.Translator().Trans(u.Ctx.Lang, "username_wrong", map[string]string{}))
	}

	//validate password
	decrypt, err := app.Crypto().Decrypt(dataUser.Password.String)
	if err != nil {
		return res, err
	}

	if p.Password.String != decrypt {
		return res, app.Error().New(http.StatusBadRequest, app.Translator().Trans(u.Ctx.Lang, "password_wrong", map[string]string{}))
	}

	//generate access token
	res.AccessToken.Set(app.Crypto().GenerateAccessToken(25))
	res.ExpiredAt.Set(time.Now().AddDate(0, 0, 1))
	res.UserId = dataUser.ID
	res.CreatedAt.Set(time.Now())

	err = tx.Create(&res).Error
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	return res, nil
}
