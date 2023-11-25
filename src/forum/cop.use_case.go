package forum

import (
	"encoding/json"
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
	Forum

	Ctx   *app.Ctx   `json:"-" db:"-" gorm:"-"`
	Query url.Values `json:"-" db:"-" gorm:"-"`
}

func (u UseCaseHandler) Async(ctx app.Ctx, query ...url.Values) UseCaseHandler {
	ctx.IsAsync = true
	return UseCase(ctx, query...)
}

// Get returns the list of Pengetahuan data.
func (u UseCaseHandler) GetSearch() (app.ListModel, error) {
	res := app.ListModel{}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	u.Query.Add("$is_disable_pagination", "true")
	// set pagination info
	res.Count,
		res.PageContext.Page,
		res.PageContext.PerPage,
		res.PageContext.PageCount,
		err = app.Query().PaginationInfo(tx, &Forum{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	// return data count if $per_page set to 0
	if res.PageContext.PerPage == 0 {
		return res, err
	}

	// find data
	data, err := app.Query().Find(tx, &Forum{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	res.SetData(data, u.Query)
	res.Count = int64(len(data))
	if u.Query.Has("levenshtein.keyword.$eq") {
		newData := []map[string]any{}
		keyword := u.Query.Get("levenshtein.keyword.$eq")
		// do for levenshtein
		listJudul := []string{}
		for _, v := range data {
			_, ok := v["judul"].(string)
			if ok {
				_, ok2 := v["deskripsi"].(string)
				if ok2 {
					listJudul = append(listJudul, v["judul"].(string)+" "+v["deskripsi"].(string))
				} else {
					listJudul = append(listJudul, v["judul"].(string))
				}
			} else {
				listJudul = append(listJudul, "")
			}
		}
		rnk := app.FindSimilarStrings(keyword, listJudul, 2)
		for i, _ := range rnk {
			data[i]["levenshtein.keyword"] = keyword

			newData = append(newData, data[i])
		}
		res.SetData(newData, u.Query)
		res.Count = int64(len(newData))
	}

	return res, err
}

func (u UseCaseHandler) GetByID(id string) (Forum, error) {
	res := Forum{}

	err := u.Ctx.ValidatePermission("forum.detail")
	if err != nil {
		return res, err
	}

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

	//get is liked & is disliked
	tx.Raw("SELECT CASE WHEN COUNT(*) > 0 THEN 1 ELSE 0 END FROM t_like WHERE id_forum = ? and id_user = ?", id, u.Ctx.User.ID).Scan(&res.IsLiked)
	tx.Raw("SELECT CASE WHEN COUNT(*) > 0 THEN 1 ELSE 0 END FROM t_dislike WHERE id_forum = ? and id_user = ?", id, u.Ctx.User.ID).Scan(&res.IsDisliked)

	//update count view
	tx.Exec("UPDATE t_forum SET count_view = count_view + 1 WHERE id = ?", id)
	return res, err
}

func (u UseCaseHandler) Get() (app.ListModel, error) {
	res := app.ListModel{}

	err := u.Ctx.ValidatePermission("forum.list")
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
		err = app.Query().PaginationInfo(tx, &Forum{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	if res.PageContext.PerPage == 0 {
		return res, err
	}

	data, err := app.Query().Find(tx, &Forum{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	jData, err := json.Marshal(data)
	if err != nil {
		return res, err
	}

	sData := []Forum{}
	err = json.Unmarshal([]byte(jData), &sData)
	if err != nil {
		return res, err
	}
	for k, d := range sData {
		var isLiked bool
		var isDisliked bool
		//get is liked & is disliked
		tx.Raw("SELECT CASE WHEN COUNT(*) > 0 THEN 1 ELSE 0 END FROM t_like WHERE id_forum = ? and id_user = ?", d.ID, u.Ctx.User.ID).Scan(&isLiked)
		tx.Raw("SELECT CASE WHEN COUNT(*) > 0 THEN 1 ELSE 0 END FROM t_dislike WHERE id_forum = ? and id_user = ?", d.ID, u.Ctx.User.ID).Scan(&isDisliked)

		sData[k].IsLiked.Set(isLiked)
		sData[k].IsDisliked.Set(isDisliked)
	}

	j2Data, err := json.Marshal(sData)
	if err != nil {
		return res, err
	}

	s2Data := []map[string]any{}
	err = json.Unmarshal([]byte(j2Data), &s2Data)
	if err != nil {
		return res, err
	}

	res.SetData(s2Data, u.Query)

	app.Cache().Set(cacheKey, res)
	return res, err
}

func (u UseCaseHandler) Create(p *ParamCreate) error {

	err := u.Ctx.ValidatePermission("forum.create")
	if err != nil {
		return err
	}

	err = u.Ctx.ValidateParam(p)
	if err != nil {
		return err
	}

	err = p.setDefaultValue(Forum{})
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

	err := u.Ctx.ValidatePermission("forum.edit")
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

	err := u.Ctx.ValidatePermission("forum.edit")
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

	err := u.Ctx.ValidatePermission("forum.delete")
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

func (u *UseCaseHandler) setDefaultValue(old Forum) error {
	if old.ID.Valid {
		u.ID = old.ID
	}

	if u.Ctx.Action.Method == "POST" {
		u.CreatedBy.Set(u.Ctx.User.ID)
	}

	if u.GambarID.Valid {
		//validasi
		_, err := attachment.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(u.GambarID.Int64)))
		if err != nil {
			return err
		}
	}

	if u.DokumenID.Valid {
		//validasi
		_, err := attachment.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(u.DokumenID.Int64)))
		if err != nil {
			return err
		}
	}

	return nil
}
