package user

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/maulanar/kms/app"
	"github.com/maulanar/kms/src/orang"
	"github.com/maulanar/kms/src/tag"
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

// UseCaseHandler provides a convenient interface for User use case, use UseCase to access UseCaseHandler.
type UseCaseHandler struct {
	User

	// injectable dependencies
	Ctx   *app.Ctx   `json:"-" db:"-" gorm:"-"`
	Query url.Values `json:"-" db:"-" gorm:"-"`
}

// Async return UseCaseHandler with async process.
func (u UseCaseHandler) Async(ctx app.Ctx, query ...url.Values) UseCaseHandler {
	ctx.IsAsync = true
	return UseCase(ctx, query...)
}

// GetByID returns the User data for the specified ID.
func (u UseCaseHandler) GetByID(id string) (User, error) {
	res := User{}

	// check permission
	err := u.Ctx.ValidatePermission("user.detail")
	if err != nil {
		return res, err
	}

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

	return res, err
}

// Get returns the list of User data.
func (u UseCaseHandler) Get() (app.ListModel, error) {
	res := app.ListModel{}

	// check permission
	err := u.Ctx.ValidatePermission("user.list")
	if err != nil {
		return res, err
	}
	// get from cache and return if exists
	cacheKey := u.EndPoint() + "?" + u.Query.Encode()
	// err = app.Cache().Get(cacheKey, &res)
	// if err == nil {
	// 	return res, err
	// }

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
		err = app.Query().PaginationInfo(tx, &User{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	// return data count if $per_page set to 0
	if res.PageContext.PerPage == 0 {
		return res, err
	}

	// find data
	data, err := app.Query().Find(tx, &User{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	res.SetData(data, u.Query)

	// save to cache and return if exists
	app.Cache().Set(cacheKey, res)
	return res, err
}

// Create creates a new data User with specified parameters.
func (u UseCaseHandler) Create(p *ParamCreate) error {
	// check permission
	err := u.Ctx.ValidatePermission("user.create")
	if err != nil {
		return err
	}

	p.Ctx = u.Ctx

	// validate param
	err = u.Ctx.ValidateParam(p)
	if err != nil {
		return err
	}

	//validasi lanjutan
	if p.OrangNik.Valid {
		if !p.Username.Valid || p.Username.String == "" {
			p.Username = p.OrangNik
		}
	}

	//enc password
	if p.Password.Valid {
		enc, err := app.Crypto().Encrypt(p.Password.String)
		if err != nil {
			return err
		}
		p.Password.Set(enc)
	}

	// copy same field to param.UsecaseHandler.Model & set default value for undefined field
	err = copySameField(p, &p.UseCaseHandler.User)
	if err != nil {
		return err
	}
	err = p.setDefaultValue(User{})
	if err != nil {
		return err
	}

	//Insert to Orang
	org := orang.ParamCreate{}
	org.Nama = p.OrangNama
	org.NamaPanggilan = p.OrangNama
	org.Nip = p.OrangNip
	org.Nik = p.OrangNik
	org.TempatLahir = p.OrangTempatLahir
	org.TglLahir = p.OrangTglLahir
	org.JenisKelamin = p.OrangJenisKelamin
	org.Alamat = p.OrangAlamat
	org.Email = p.OrangEmail
	org.Telp = p.OrangTelp
	org.Jabatan = p.OrangJabatan
	org.FotoID = p.OrangFotoID
	org.UnitKerja = p.OrangUnitKerja
	org.UserLevel = p.OrangUserLevel
	org.StatusLevel = p.OrangStatusLevel

	err = orang.UseCase(*u.Ctx).Create(&org)
	if err != nil {
		return err
	}

	p.OrangId.Set(org.ID.Int64)
	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// save data to db
	err = tx.Model(&p).Create(p).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	//relation folowing tags
	if len(p.FollowingTags) > 0 {
		for _, ft := range p.FollowingTags {
			//validation
			tg, err := tag.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(ft.HastagID.Int64)))
			if err != nil {
				return err
			}
			ft.UserID.Set(u.Ctx.User.ID)
			ft.HastagID.Set(tg.ID.Int64)
			err = tx.Create(&ft).Error
			if err != nil {
				return err
			}
		}
	}

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint())

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("POST", "create", strconv.Itoa(int(p.ID.Int64)), p)
	return nil
}

func copySameField(source interface{}, dest *User) error {
	sourceValue := reflect.ValueOf(source)
	destValue := reflect.ValueOf(dest).Elem()

	if sourceValue.Kind() != reflect.Ptr || sourceValue.IsNil() {
		return fmt.Errorf("source must be a non-nil pointer to a struct")
	}

	sourceValue = sourceValue.Elem()
	if sourceValue.Kind() != reflect.Struct {
		return fmt.Errorf("source must be a non-nil pointer to a struct")
	}

	for i := 0; i < sourceValue.NumField(); i++ {
		sourceField := sourceValue.Field(i)
		destField := destValue.FieldByName(sourceValue.Type().Field(i).Name)

		if destField.IsValid() && sourceField.Type() == destField.Type() {
			destField.Set(sourceField)
		}
	}

	return nil
}

// UpdateByID updates the User data for the specified ID with specified parameters.
func (u UseCaseHandler) UpdateByID(id string, p *ParamUpdate) error {

	// check permission
	err := u.Ctx.ValidatePermission("user.edit")
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

	//validasi lanjutan
	if p.OrangNik.Valid {
		if !p.Username.Valid || p.Username.String == "" {
			p.Username = p.OrangNik
		}
	}

	// set default value for undefined field
	err = p.setDefaultValue(old)
	if err != nil {
		return err
	}

	//enc password
	if p.Password.Valid {
		enc, err := app.Crypto().Encrypt(p.Password.String)
		if err != nil {
			return err
		}
		p.Password.Set(enc)
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// update data on the db
	err = tx.Model(&p).Where("id_user = ?", old.ID).Updates(p).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	//relation folowing tags
	if len(p.FollowingTags) > 0 {
		tx.Where("id_user", u.Ctx.User.ID).Delete(&FollowdHastag{})
		for _, ft := range p.FollowingTags {
			//validation
			tg, err := tag.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(ft.HastagID.Int64)))
			if err != nil {
				return err
			}
			ft.UserID.Set(u.Ctx.User.ID)
			ft.HastagID.Set(tg.ID.Int64)
			err = tx.Create(&ft).Error
			if err != nil {
				return err
			}
		}
	}

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint(), strconv.Itoa(int(old.ID.Int64)))

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("PUT", "By Sistem", strconv.Itoa(int(old.ID.Int64)), old)
	return nil
}

// PartiallyUpdateByID updates the User data for the specified ID with specified parameters.
func (u UseCaseHandler) PartiallyUpdateByID(id string, p *ParamPartiallyUpdate) error {

	// check permission
	err := u.Ctx.ValidatePermission("user.edit")
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

	//validasi lanjutan
	if p.OrangNik.Valid {
		if !p.Username.Valid || p.Username.String == "" {
			p.Username = p.OrangNik
		}
	}

	// set default value for undefined field
	err = p.setDefaultValue(old)
	if err != nil {
		return err
	}

	//enc password
	if p.Password.Valid {
		enc, err := app.Crypto().Encrypt(p.Password.String)
		if err != nil {
			return err
		}
		p.Password.Set(enc)
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// update data on the db
	err = tx.Model(&p).Where("id_user = ?", old.ID).Updates(p).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	//relation folowing tags
	if len(p.FollowingTags) > 0 {
		tx.Where("id_user", u.Ctx.User.ID).Delete(&FollowdHastag{})
		for _, ft := range p.FollowingTags {
			//validation
			tg, err := tag.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(ft.HastagID.Int64)))
			if err != nil {
				return err
			}
			ft.UserID.Set(u.Ctx.User.ID)
			ft.HastagID.Set(tg.ID.Int64)
			err = tx.Create(&ft).Error
			if err != nil {
				return err
			}
		}
	}

	//update orang
	op := orang.ParamPartiallyUpdate{}
	op.ID = p.OrangId
	op.Nama = p.OrangNama
	op.NamaPanggilan = p.OrangNamaPanggilan
	op.Nip = p.OrangNip
	op.Nik = p.OrangNik
	op.TempatLahir = p.OrangTempatLahir
	op.TglLahir = p.OrangTglLahir
	op.JenisKelamin = p.OrangJenisKelamin
	op.Alamat = p.OrangAlamat
	op.Email = p.OrangEmail
	op.Telp = p.OrangTelp
	op.Jabatan = p.OrangJabatan
	op.FotoID = p.OrangFotoID
	op.FotoUrl = p.OrangFotoUrl
	op.FotoNama = p.OrangFotoNama
	op.UnitKerja = p.OrangUnitKerja
	op.UserLevel = p.OrangUserLevel
	op.StatusLevel = p.OrangStatusLevel
	if old.OrangId.Valid {
		err = orang.UseCase(*u.Ctx).PartiallyUpdateByID(strconv.Itoa(int(old.OrangId.Int64)), &op)
		if err != nil {
			return err
		}
	}

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint(), strconv.Itoa(int(old.ID.Int64)))

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("PATCH", "By Sistem", strconv.Itoa(int(old.ID.Int64)), old)
	return nil
}

// DeleteByID deletes the User data for the specified ID.
func (u UseCaseHandler) DeleteByID(id string, p *ParamDelete) error {

	// check permission
	err := u.Ctx.ValidatePermission("user.delete")
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

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// update data on the db
	err = tx.Model(&p).Where("id_user = ?", old.ID).Update("deleted_at", time.Now().UTC()).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint(), strconv.Itoa(int(old.ID.Int64)))

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("DELETE", "By Sistem", strconv.Itoa(int(old.ID.Int64)), old)
	return nil
}

// setDefaultValue set default value of undefined field when create or update User data.
func (p *UseCaseHandler) setDefaultValue(old User) error {
	if old.ID.Valid {
		p.ID = old.ID
	}

	if p.User.Username.Valid && (p.User.Username != old.Username) {
		//check ke db, pastikan tidak duplikat
		var count int64
		tx, err := p.Ctx.DB()
		if err != nil {
			return app.Error().New(http.StatusInternalServerError, err.Error())
		}

		tx.Model(p.User).Where("username = ?", p.User.Username.String).Where("deleted_at IS NULL").Count(&count)
		if count > 1 {
			return app.Error().New(http.StatusBadRequest, app.Translator().Trans(p.Ctx.Lang, "unique", map[string]string{
				"attribute": "username",
				"value":     p.User.Username.String,
			}))
		}
	}

	return nil
}
