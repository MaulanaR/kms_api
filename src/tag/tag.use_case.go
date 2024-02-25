package tag

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/maulanar/kms/app"
	"github.com/maulanar/kms/src/kompetensi"
	"grest.dev/grest"
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

// UseCaseHandler provides a convenient interface for Tag use case, use UseCase to access UseCaseHandler.
type UseCaseHandler struct {
	Tag

	// injectable dependencies
	Ctx   *app.Ctx   `json:"-" db:"-" gorm:"-"`
	Query url.Values `json:"-" db:"-" gorm:"-"`
}

// Async return UseCaseHandler with async process.
func (u UseCaseHandler) Async(ctx app.Ctx, query ...url.Values) UseCaseHandler {
	ctx.IsAsync = true
	return UseCase(ctx, query...)
}

// GetByID returns the Tag data for the specified ID.
func (u UseCaseHandler) GetByID(id string) (Tag, error) {
	res := Tag{}

	// check permission
	err := u.Ctx.ValidatePermission("tag.detail")
	if err != nil {
		return res, err
	}

	// get from cache and return if exists
	cacheKey := u.EndPoint() + "." + id
	app.Cache().Get(cacheKey, &res)
	if res.ID.Valid {
		return res, err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// get from db
	key := "id"

	u.Query.Add(key, id)
	err = app.Query().First(tx, &res, u.Query)
	if err != nil {
		return res, u.Ctx.NotFoundError(err, u.EndPoint(), key, id)
	}

	// save to cache and return if exists
	app.Cache().Set(cacheKey, res)
	return res, err
}

// Get returns the list of Tag data.
func (u UseCaseHandler) Get() (app.ListModel, error) {
	res := app.ListModel{}

	// check permission
	err := u.Ctx.ValidatePermission("tag.list")
	if err != nil {
		return res, err
	}

	u.GetDataFromAPI()

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
		err = app.Query().PaginationInfo(tx, &Tag{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	// return data count if $per_page set to 0
	if res.PageContext.PerPage == 0 {
		return res, err
	}

	// find data
	data, err := app.Query().Find(tx, &Tag{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	res.SetData(data, u.Query)

	return res, err
}

// Create creates a new data Tag with specified parameters.
func (u UseCaseHandler) Create(p *ParamCreate) error {

	// check permission
	err := u.Ctx.ValidatePermission("tag.create")
	if err != nil {
		return err
	}

	// validate param
	err = u.Ctx.ValidateParam(p)
	if err != nil {
		return err
	}

	// set default value for undefined field
	err = p.setDefaultValue(Tag{})
	if err != nil {
		return err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// save data to db
	err = tx.Model(&p).Create(&p).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint())

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("POST", "create", strconv.Itoa(int(p.ID.Int64)), p)
	return nil
}

// UpdateByID updates the Tag data for the specified ID with specified parameters.
func (u UseCaseHandler) UpdateByID(id string, p *ParamUpdate) error {

	// check permission
	err := u.Ctx.ValidatePermission("tag.edit")
	if err != nil {
		return err
	}

	// validate param
	err = u.Ctx.ValidateParam(p)
	if err != nil {
		return err
	}

	// get previous data
	old, err := u.GetByID(id)
	if err != nil {
		return err
	}

	if old.Jenis.String == "stara" {
		return app.Error().New(http.StatusForbidden, "Data berasal dari API STARA, untuk mengubahnya silahkan melalui di API STARA.")
	}

	// set default value for undefined field
	err = p.setDefaultValue(old)
	if err != nil {
		return err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// update data on the db
	err = tx.Model(&p).Where("id_tag = ?", old.ID).Updates(p).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint(), strconv.Itoa(int(old.ID.Int64)))

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("PUT", "By Sistem", strconv.Itoa(int(old.ID.Int64)), old)
	return nil
}

// PartiallyUpdateByID updates the Tag data for the specified ID with specified parameters.
func (u UseCaseHandler) PartiallyUpdateByID(id string, p *ParamPartiallyUpdate) error {

	// check permission
	err := u.Ctx.ValidatePermission("tag.edit")
	if err != nil {
		return err
	}

	// validate param
	err = u.Ctx.ValidateParam(p)
	if err != nil {
		return err
	}

	// get previous data
	old, err := u.GetByID(id)
	if err != nil {
		return err
	}

	if old.Jenis.String == "stara" {
		return app.Error().New(http.StatusForbidden, "Data berasal dari API STARA, untuk mengubahnya silahkan melalui di API STARA.")
	}

	// set default value for undefined field
	err = p.setDefaultValue(old)
	if err != nil {
		return err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// update data on the db
	err = tx.Model(&p).Where("id_tag = ?", old.ID).Updates(p).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint(), strconv.Itoa(int(old.ID.Int64)))

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("PATCH", "By Sistem", strconv.Itoa(int(old.ID.Int64)), old)
	return nil
}

// DeleteByID deletes the Tag data for the specified ID.
func (u UseCaseHandler) DeleteByID(id string, p *ParamDelete) error {

	// check permission
	err := u.Ctx.ValidatePermission("tag.delete")
	if err != nil {
		return err
	}

	// validate param
	err = u.Ctx.ValidateParam(p)
	if err != nil {
		return err
	}

	// get previous data
	old, err := u.GetByID(id)
	if err != nil {
		return err
	}

	if old.Jenis.String == "stara" {
		return app.Error().New(http.StatusForbidden, "Data berasal dari API STARA, untuk menghapusnya silahkan melalui di API STARA.")
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// update data on the db
	err = tx.Model(&p).Where("id_tag = ?", old.ID).Update("deleted_at", time.Now().UTC()).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint(), strconv.Itoa(int(old.ID.Int64)))

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("DELETE", "By Sistem", strconv.Itoa(int(old.ID.Int64)), old)
	return nil
}

// setDefaultValue set default value of undefined field when create or update Tag data.
func (u *UseCaseHandler) setDefaultValue(old Tag) error {
	if old.ID.Valid {
		u.ID = old.ID
	}

	return nil
}

func (u UseCaseHandler) GetDataFromAPI() error {
	//LOGIN TO API STARA
	endPoint := "https://api-stara.bpkp.go.id/api/auth/login"
	body := map[string]any{
		"username": "eko r prastiawan",
		"password": "jlnias&7",
	}

	c := app.HttpClient("POST", endPoint)
	// c.AddHeader(app.AuthHeaderKey, "Bearer "+u.Ctx.Auth.AccessToken)
	c.Debug()
	c.AddJsonBody(body)
	_, err := c.Send()
	if err != nil {
		return err
	}
	staraAuth := kompetensi.LoginStara{}
	err = grest.NewJSON(c.BodyResponse).ToFlat().Unmarshal(&staraAuth)
	if err != nil {
		return err
	}
	// c.UnmarshalJson(&staraAuth)

	//GET DATA FROM API STARA
	c2 := app.HttpClient("GET", "https://api-stara.bpkp.go.id/api/kompetensi")
	c2.AddHeader("Authorization", "Bearer "+staraAuth.Token.String)
	c2.Debug()
	_, err = c2.Send()
	if err != nil {
		return err
	}
	resC2 := make(map[string]interface{})
	c2.UnmarshalJson(&resC2)

	kompentensi := []kompetensi.Kompetensi{}
	err = grest.NewJSON(resC2["data"]).ToFlat().Unmarshal(&kompentensi)
	if err != nil {
		return err
	}

	var allIDMappings []int64

	// Iterasi melalui slice Kompentensi
	for _, k := range kompentensi {
		allIDMappings = append(allIDMappings, k.IDMapping.Int64)
	}

	tx, err := u.Ctx.DB()
	if err != nil {
		return err
	}

	//hapus data di database jika id_mapping tidak ada di list ini
	tx.Not("id_mapping =?", allIDMappings).Where("jenis = ?", "stara").Delete(Tag{})

	//Update / Insert data
	for _, k := range kompentensi {
		dt := Tag{}
		result := tx.Where("id_mapping = ?", k.IDMapping).Where("jenis = ?", "stara").First(&dt)
		if result.RowsAffected < 1 {
			//insert
			dt.Nama = k.NamaKompetensiSDM
			dt.StaraID = k.IDMapping
			dt.Jenis.Set("stara")
			err = tx.Create(&dt).Error
			if err != nil {
				return err
			}
		} else {
			//update
			if dt.Nama != k.NamaKompetensiSDM {
				dt.Nama = k.NamaKompetensiSDM
				dt.Jenis.Set("stara")
				err = tx.Where("id_tag = ?", dt.ID).Updates(&dt).Error
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
