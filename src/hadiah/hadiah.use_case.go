package hadiah

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/maulanar/kms/app"
	"github.com/maulanar/kms/src/attachment"
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
	Hadiah

	Ctx   *app.Ctx   `json:"-" db:"-" gorm:"-"`
	Query url.Values `json:"-" db:"-" gorm:"-"`
}

func (u UseCaseHandler) Async(ctx app.Ctx, query ...url.Values) UseCaseHandler {
	ctx.IsAsync = true
	return UseCase(ctx, query...)
}

func (u UseCaseHandler) GetByID(id string) (Hadiah, error) {
	res := Hadiah{}

	err := u.Ctx.ValidatePermission("hadiah.detail")
	if err != nil {
		return res, err
	}

	cacheKey := u.EndPoint() + "." + id
	app.Cache().Get(cacheKey, &res)
	if res.ID.Valid {
		return res, err
	}

	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	key := "id"

	u.Query.Add(key, id)
	err = app.Query().First(tx, &res, u.Query)
	if err != nil {
		return res, u.Ctx.NotFoundError(err, u.EndPoint(), key, id)
	}

	app.Cache().Set(cacheKey, res)
	return res, err
}

func (u UseCaseHandler) Get() (app.ListModel, error) {
	res := app.ListModel{}

	err := u.Ctx.ValidatePermission("hadiah.list")
	if err != nil {
		return res, err
	}

	cacheKey := u.EndPoint() + "?" + u.Query.Encode()
	err = app.Cache().Get(cacheKey, &res)
	if err == nil {
		return res, err
	}

	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	res.Count,
		res.PageContext.Page,
		res.PageContext.PerPage,
		res.PageContext.PageCount,
		err = app.Query().PaginationInfo(tx, &Hadiah{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	if res.PageContext.PerPage == 0 {
		return res, err
	}

	data, err := app.Query().Find(tx, &Hadiah{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	res.SetData(data, u.Query)

	app.Cache().Set(cacheKey, res)
	return res, err
}

func (u UseCaseHandler) Create(p *ParamCreate) error {

	err := u.Ctx.ValidatePermission("hadiah.create")
	if err != nil {
		return err
	}

	err = u.Ctx.ValidateParam(p)
	if err != nil {
		return err
	}

	err = p.setDefaultValue(Hadiah{})
	if err != nil {
		return err
	}

	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	err = tx.Model(&p).Create(&p).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	app.Cache().Invalidate(u.EndPoint())

	go u.Ctx.Hook("POST", "create", strconv.Itoa(int(p.ID.Int64)), p)
	return nil
}

func (u UseCaseHandler) UpdateByID(id string, p *ParamUpdate) error {

	err := u.Ctx.ValidatePermission("hadiah.edit")
	if err != nil {
		return err
	}

	err = u.Ctx.ValidateParam(p)
	if err != nil {
		return err
	}

	old, err := u.GetByID(id)
	if err != nil {
		return err
	}

	err = p.setDefaultValue(old)
	if err != nil {
		return err
	}

	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	err = tx.Model(&p).Where("id = ?", old.ID).Updates(p).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	app.Cache().Invalidate(u.EndPoint(), strconv.Itoa(int(old.ID.Int64)))

	go u.Ctx.Hook("PUT", "By Sistem", strconv.Itoa(int(old.ID.Int64)), old)
	return nil
}

func (u UseCaseHandler) PartiallyUpdateByID(id string, p *ParamPartiallyUpdate) error {

	err := u.Ctx.ValidatePermission("hadiah.edit")
	if err != nil {
		return err
	}

	err = u.Ctx.ValidateParam(p)
	if err != nil {
		return err
	}

	old, err := u.GetByID(id)
	if err != nil {
		return err
	}

	err = p.setDefaultValue(old)
	if err != nil {
		return err
	}

	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	err = tx.Model(&p).Where("id = ?", old.ID).Updates(p).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	app.Cache().Invalidate(u.EndPoint(), strconv.Itoa(int(old.ID.Int64)))

	go u.Ctx.Hook("PATCH", "By Sistem", strconv.Itoa(int(old.ID.Int64)), old)
	return nil
}

func (u UseCaseHandler) DeleteByID(id string, p *ParamDelete) error {

	err := u.Ctx.ValidatePermission("hadiah.delete")
	if err != nil {
		return err
	}

	err = u.Ctx.ValidateParam(p)
	if err != nil {
		return err
	}

	old, err := u.GetByID(id)
	if err != nil {
		return err
	}

	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	err = tx.Model(&p).Where("id = ?", old.ID).Update("deleted_at", time.Now().UTC()).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	app.Cache().Invalidate(u.EndPoint(), strconv.Itoa(int(old.ID.Int64)))

	go u.Ctx.Hook("DELETE", "By Sistem", strconv.Itoa(int(old.ID.Int64)), old)
	return nil
}

func (u *UseCaseHandler) setDefaultValue(old Hadiah) error {
	if old.ID.Valid {
		u.ID = old.ID
	}

	if u.Ctx.Action.Method == "POST" {
		u.CreatedBy.Set(u.Ctx.User.ID)
	}

	if !u.IsActive.Valid {
		u.IsActive.Set(true)
	}

	if u.GambarID.Valid {
		//validasi
		_, err := attachment.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(u.GambarID.Int64)))
		if err != nil {
			return err
		}
	}

	if u.Ctx.Action.Method == "PUT" || u.Ctx.Action.Method == "PATCH" {
		u.UpdatedBy.Set(u.Ctx.User.ID)
	}

	if u.Ctx.Action.Method == "DELETE" {
		u.DeletedBy.Set(u.Ctx.User.ID)
	}
	return nil
}

func (u UseCaseHandler) GetRanking() (app.ListModel, error) {
	res := app.ListModel{}

	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	res.Count,
		res.PageContext.Page,
		res.PageContext.PerPage,
		res.PageContext.PageCount,
		err = app.Query().PaginationInfo(tx, &RankingPoint{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	if res.PageContext.PerPage == 0 {
		return res, err
	}

	data, err := app.Query().Find(tx, &RankingPoint{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	res.SetData(data, u.Query)

	return res, err
}
