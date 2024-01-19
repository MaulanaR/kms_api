package dislike

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/maulanar/kms/app"
	"github.com/maulanar/kms/src/forum"
	"github.com/maulanar/kms/src/leadertalk"
	"github.com/maulanar/kms/src/librarycafe"
	"github.com/maulanar/kms/src/notifikasi"
	"github.com/maulanar/kms/src/pengetahuan"
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
	Dislike

	Ctx   *app.Ctx   `json:"-" db:"-" gorm:"-"`
	Query url.Values `json:"-" db:"-" gorm:"-"`
}

func (u UseCaseHandler) Async(ctx app.Ctx, query ...url.Values) UseCaseHandler {
	ctx.IsAsync = true
	return UseCase(ctx, query...)
}

func (u UseCaseHandler) GetByID(id string) (Dislike, error) {
	res := Dislike{}

	err := u.Ctx.ValidatePermission("dislike.detail")
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

	err := u.Ctx.ValidatePermission("dislike.list")
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
		err = app.Query().PaginationInfo(tx, &Dislike{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	if res.PageContext.PerPage == 0 {
		return res, err
	}

	data, err := app.Query().Find(tx, &Dislike{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	res.SetData(data, u.Query)

	app.Cache().Set(cacheKey, res)
	return res, err
}

func (u UseCaseHandler) Create(p *ParamCreate) error {

	err := u.Ctx.ValidatePermission("dislike.create")
	if err != nil {
		return err
	}

	err = u.Ctx.ValidateParam(p)
	if err != nil {
		return err
	}

	err = p.setDefaultValue(Dislike{})
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

	err := u.Ctx.ValidatePermission("dislike.edit")
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

	err := u.Ctx.ValidatePermission("dislike.edit")
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

	err := u.Ctx.ValidatePermission("dislike.delete")
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

func (u *UseCaseHandler) setDefaultValue(old Dislike) error {
	if old.ID.Valid {
		u.ID = old.ID
	}

	if old.UserID.Valid {
		u.UserID = old.UserID
	}

	if !u.UserID.Valid {
		u.UserID.Set(u.Ctx.User.ID)
	}

	//validasi
	if u.PengetahuanID.Valid {
		_, err := pengetahuan.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(u.PengetahuanID.Int64)))
		if err != nil {
			return err
		}
	}

	//validasi
	if u.ForumID.Valid {
		_, err := forum.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(u.ForumID.Int64)))
		if err != nil {
			return err
		}
	}

	//validasi
	if u.LeaderTalkID.Valid {
		_, err := leadertalk.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(u.LeaderTalkID.Int64)))
		if err != nil {
			return err
		}
	}

	//validasi
	if u.LibraryCafeID.Valid {
		_, err := librarycafe.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(u.LibraryCafeID.Int64)))
		if err != nil {
			return err
		}
	}

	return nil
}

func (u UseCaseHandler) UpdateByPengetahuanID(id string, p *ParamUpdate) error {
	//validasi
	pgth, err := pengetahuan.UseCase(*u.Ctx).GetByID(id)
	if err != nil {
		return err
	}

	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	//make sure that person havent doing this action before
	var exist int64
	tx.Model(&Dislike{}).
		Where("id_pengetahuan = ?", pgth.ID.Int64).
		Where("id_user = ?", u.Ctx.User.ID).
		Count(&exist)

	//delete like action if exists
	var likeExist int64 = 0
	tx.Table("t_like").
		Where("id_pengetahuan = ?", pgth.ID.Int64).
		Where("id_user = ?", u.Ctx.User.ID).
		Count(&likeExist)
	if likeExist > 0 {
		//delete like
		tx.Table("t_like").
			Where("id_pengetahuan = ?", pgth.ID.Int64).
			Where("id_user = ?", u.Ctx.User.ID).
			Delete("t_like")
	}

	if exist < 1 {
		//Set param
		p.PengetahuanID.Set(pgth.ID.Int64)
		p.UserID.Set(u.Ctx.User.ID)
		p.CreatedAt.Set(time.Now())

		err = tx.Model(&p).Create(&p).Error
		if err != nil {
			return app.Error().New(http.StatusInternalServerError, err.Error())
		}

		go notifikasi.UseCase(*u.Ctx).SaveNotif("Dislike Baru", u.Ctx.User.OrangNama+" baru saja dislike atas postingan anda.", pgth.CreatedBy.Int64, pgth.EndPoint(), pgth.ID.Int64, p)
	}

	app.Cache().Invalidate(u.EndPoint())

	go u.Ctx.Hook("POST", "create", strconv.Itoa(int(p.ID.Int64)), p)
	return nil
}

func (u UseCaseHandler) UpdateByforumID(id string, p *ParamUpdate) error {
	//validasi
	pgth, err := forum.UseCase(*u.Ctx).GetByID(id)
	if err != nil {
		return err
	}

	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	//make sure that person havent doing this action before
	var exist int64
	tx.Model(&Dislike{}).
		Where("id_forum = ?", pgth.ID.Int64).
		Where("id_user = ?", u.Ctx.User.ID).
		Count(&exist)

	//delete like action if exists
	var likeExist int64 = 0
	tx.Table("t_like").
		Where("id_forum = ?", pgth.ID.Int64).
		Where("id_user = ?", u.Ctx.User.ID).
		Count(&likeExist)
	if likeExist > 0 {
		//delete like
		tx.Table("t_like").
			Where("id_forum = ?", pgth.ID.Int64).
			Where("id_user = ?", u.Ctx.User.ID).
			Delete("t_like")
	}

	if exist < 1 {
		//Set param
		p.ForumID.Set(pgth.ID.Int64)
		p.UserID.Set(u.Ctx.User.ID)
		p.CreatedAt.Set(time.Now())

		err = tx.Model(&p).Create(&p).Error
		if err != nil {
			return app.Error().New(http.StatusInternalServerError, err.Error())
		}

		go notifikasi.UseCase(*u.Ctx).SaveNotif("Dislike Baru", u.Ctx.User.OrangNama+" baru saja dislike atas postingan anda.", pgth.CreatedBy.Int64, pgth.EndPoint(), pgth.ID.Int64, p)

	}

	app.Cache().Invalidate(u.EndPoint())

	go u.Ctx.Hook("POST", "create", strconv.Itoa(int(p.ID.Int64)), p)
	return nil
}

func (u UseCaseHandler) UpdateByLeaderTalkID(id string, p *ParamUpdate) error {
	//validasi
	pgth, err := leadertalk.UseCase(*u.Ctx).GetByID(id)
	if err != nil {
		return err
	}

	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	//make sure that person havent doing this action before
	var exist int64
	tx.Model(&Dislike{}).
		Where("id_leader_talk = ?", pgth.ID.Int64).
		Where("id_user = ?", u.Ctx.User.ID).
		Count(&exist)

	//delete like action if exists
	var likeExist int64 = 0
	tx.Table("t_like").
		Where("id_leader_talk = ?", pgth.ID.Int64).
		Where("id_user = ?", u.Ctx.User.ID).
		Count(&likeExist)
	if likeExist > 0 {
		//delete like
		tx.Table("t_like").
			Where("id_leader_talk = ?", pgth.ID.Int64).
			Where("id_user = ?", u.Ctx.User.ID).
			Delete("t_like")
	}

	if exist < 1 {
		//Set param
		p.LeaderTalkID.Set(pgth.ID.Int64)
		p.UserID.Set(u.Ctx.User.ID)
		p.CreatedAt.Set(time.Now())

		err = tx.Model(&p).Create(&p).Error
		if err != nil {
			return app.Error().New(http.StatusInternalServerError, err.Error())
		}
	}

	app.Cache().Invalidate(u.EndPoint())

	go u.Ctx.Hook("POST", "create", strconv.Itoa(int(p.ID.Int64)), p)
	return nil
}

func (u UseCaseHandler) UpdateByLibraryCafeID(id string, p *ParamUpdate) error {
	//validasi
	pgth, err := librarycafe.UseCase(*u.Ctx).GetByID(id)
	if err != nil {
		return err
	}

	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	//make sure that person havent doing this action before
	var exist int64
	tx.Model(&Dislike{}).
		Where("id_library_cafe = ?", pgth.ID.Int64).
		Where("id_user = ?", u.Ctx.User.ID).
		Count(&exist)

	//delete like action if exists
	var likeExist int64 = 0
	tx.Table("t_like").
		Where("id_library_cafe = ?", pgth.ID.Int64).
		Where("id_user = ?", u.Ctx.User.ID).
		Count(&likeExist)
	if likeExist > 0 {
		//delete like
		tx.Table("t_like").
			Where("id_library_cafe = ?", pgth.ID.Int64).
			Where("id_user = ?", u.Ctx.User.ID).
			Delete("t_like")
	}

	if exist < 1 {
		//Set param
		p.LibraryCafeID.Set(pgth.ID.Int64)
		p.UserID.Set(u.Ctx.User.ID)
		p.CreatedAt.Set(time.Now())

		err = tx.Model(&p).Create(&p).Error
		if err != nil {
			return app.Error().New(http.StatusInternalServerError, err.Error())
		}
	}

	app.Cache().Invalidate(u.EndPoint())

	go u.Ctx.Hook("POST", "create", strconv.Itoa(int(p.ID.Int64)), p)
	return nil
}
