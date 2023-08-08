package pengetahuan

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/maulanar/kms/app"
	"github.com/maulanar/kms/src/attachment"
	"github.com/maulanar/kms/src/jenispengetahuan"
	"github.com/maulanar/kms/src/kompetensi"
	"github.com/maulanar/kms/src/lingkuppengetahuan"
	"github.com/maulanar/kms/src/orang"
	"github.com/maulanar/kms/src/referensi"
	"github.com/maulanar/kms/src/statuspengetahuan"
	"github.com/maulanar/kms/src/tag"
	"github.com/maulanar/kms/src/tpengetahuanrelation"
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

// UseCaseHandler provides a convenient interface for Pengetahuan use case, use UseCase to access UseCaseHandler.
type UseCaseHandler struct {
	Pengetahuan

	// injectable dependencies
	Ctx   *app.Ctx   `json:"-" db:"-" gorm:"-"`
	Query url.Values `json:"-" db:"-" gorm:"-"`
}

// Async return UseCaseHandler with async process.
func (u UseCaseHandler) Async(ctx app.Ctx, query ...url.Values) UseCaseHandler {
	ctx.IsAsync = true
	return UseCase(ctx, query...)
}

// GetByID returns the Pengetahuan data for the specified ID.
func (u UseCaseHandler) GetByID(id string) (Pengetahuan, error) {
	res := Pengetahuan{}

	// check permission
	err := u.Ctx.ValidatePermission("pengetahuan.detail")
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

// Get returns the list of Pengetahuan data.
func (u UseCaseHandler) Get() (app.ListModel, error) {
	res := app.ListModel{}

	// check permission
	err := u.Ctx.ValidatePermission("pengetahuan.list")
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
		err = app.Query().PaginationInfo(tx, &Pengetahuan{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	// return data count if $per_page set to 0
	if res.PageContext.PerPage == 0 {
		return res, err
	}

	// find data
	data, err := app.Query().Find(tx, &Pengetahuan{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	res.SetData(data, u.Query)

	// save to cache and return if exists
	app.Cache().Set(cacheKey, res)
	return res, err
}

// Create creates a new data Pengetahuan with specified parameters.
func (u UseCaseHandler) Create(p *ParamCreate) error {

	// check permission
	err := u.Ctx.ValidatePermission("pengetahuan.create")
	if err != nil {
		return err
	}

	// validate param
	err = u.Ctx.ValidateParam(p)
	if err != nil {
		return err
	}

	// set default value for undefined field
	err = p.setDefaultValue(Pengetahuan{})
	if err != nil {
		return err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	//cek by jenis
	jp := jenispengetahuan.UseCaseHandler{
		Ctx:   u.Ctx,
		Query: url.Values{},
	}

	jenis, err := jp.GetByID(strconv.Itoa(int(p.JenisPengetahuanID.Int64)))
	if err != nil {
		return err
	}

	//validasi LingkupPengetahuan
	lp := lingkuppengetahuan.UseCaseHandler{
		Ctx:   u.Ctx,
		Query: url.Values{},
	}

	_, err = lp.GetByID(strconv.Itoa(int(p.LingkupPengetahuanID.Int64)))
	if err != nil {
		return err
	}

	//validasi StatusPengetahuan
	sp := statuspengetahuan.UseCaseHandler{
		Ctx:   u.Ctx,
		Query: url.Values{},
	}

	_, err = sp.GetByID(strconv.Itoa(int(p.StatusPengetahuanID.Int64)))
	if err != nil {
		return err
	}

	org := orang.UseCaseHandler{
		Ctx:   u.Ctx,
		Query: url.Values{},
	}
	//validasi penulis (orang)
	if p.Penulis1ID.Valid {
		_, err = org.GetByID(strconv.Itoa(int(p.Penulis1ID.Int64)))
		if err != nil {
			return err
		}
	}

	if p.Penulis2ID.Valid {
		_, err = org.GetByID(strconv.Itoa(int(p.Penulis2ID.Int64)))
		if err != nil {
			return err
		}
	}

	if p.Penulis3ID.Valid {
		_, err = org.GetByID(strconv.Itoa(int(p.Penulis3ID.Int64)))
		if err != nil {
			return err
		}
	}

	// save data to db to get ID
	err = tx.Model(&p).Create(&p).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	//RELATION

	//validasi referensi
	if len(p.Referensi) > 0 {
		rf := referensi.UseCaseHandler{
			Ctx:   u.Ctx,
			Query: url.Values{},
		}

		for i, ref := range p.Referensi {
			//validasi
			_, err := rf.GetByID(strconv.Itoa(int(ref.ReferensiID.Int64)))
			if err != nil {
				return err
			}
			p.Referensi[i].PengetahuanID.Set(p.ID.Int64)
		}

		err = tx.Create(&p.Referensi).Error
		if err != nil {
			return err
		}
	}

	//validasi penulis
	// if len(p.PenulisExternal) > 0 {
	// 	for i, _ := range p.PenulisExternal {
	// 		p.PenulisExternal[i].PengetahuanID.Set(p.ID.Int64)
	// 	}

	// 	err = tx.Create(&p.PenulisExternal).Error
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	//validasi hastag
	if len(p.Tag) > 0 {
		tg := tag.UseCaseHandler{
			Ctx:   u.Ctx,
			Query: url.Values{},
		}

		for i, ref := range p.Tag {
			//validasi
			_, err := tg.GetByID(strconv.Itoa(int(ref.TagID.Int64)))
			if err != nil {
				return err
			}
			p.Tag[i].PengetahuanID.Set(p.ID.Int64)
		}

		err = tx.Create(&p.Tag).Error
		if err != nil {
			return err
		}
	}

	//validasi kompetensi
	if len(p.Kompetensi) > 0 {
		kpt := kompetensi.UseCaseHandler{
			Ctx:   u.Ctx,
			Query: url.Values{},
		}
		for i, ref := range p.Kompetensi {
			//validasi
			_, err := kpt.GetByID(strconv.Itoa(int(ref.KompetensiID.Int64)))
			if err != nil {
				return err
			}
			p.Kompetensi[i].PengetahuanID.Set(p.ID.Int64)
		}

		err = tx.Create(&p.Kompetensi).Error
		if err != nil {
			return err
		}
	}

	//validasi dokumen
	if len(p.Dokumen) > 0 {
		doku := attachment.UseCaseHandler{
			Ctx:   u.Ctx,
			Query: url.Values{},
		}
		for i, ref := range p.Dokumen {
			//validasi
			_, err := doku.GetByID(strconv.Itoa(int(ref.AttachmentID.Int64)))
			if err != nil {
				return err
			}
			p.Dokumen[i].PengetahuanID.Set(p.ID.Int64)
		}

		err = tx.Create(&p.Dokumen).Error
		if err != nil {
			return err
		}
	}

	if jenis.Nama.String == "Tugas" || jenis.Nama.String == "Tugas (Panduan Penugasan)" {
		relTugas := tpengetahuanrelation.TPengetahuanTugas{}
		relTugas.PengetahuanID.Set(p.ID.Int64)
		relTugas.Tujuan = p.Tujuan
		relTugas.DasarHukum = p.DasarHukum
		relTugas.ProsesBisnis = p.ProsesBisnis
		relTugas.RumusanMasalah = p.RumusanMasalah
		relTugas.PenyebabTemuan = p.PenyebabTemuan
		relTugas.Keahlian = p.Keahlian
		relTugas.KebutuhanData = p.KebutuhanData

		err = tx.Create(&relTugas).Error
		if err != nil {
			return err
		}
	}

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint())

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("POST", "create", strconv.Itoa(int(p.ID.Int64)), p)
	return nil
}

// UpdateByID updates the Pengetahuan data for the specified ID with specified parameters.
func (u UseCaseHandler) UpdateByID(id string, p *ParamUpdate) error {

	// check permission
	err := u.Ctx.ValidatePermission("pengetahuan.edit")
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
	err = tx.Model(&p).Where("id_pengetahuan = ?", old.ID).Updates(p).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint(), strconv.Itoa(int(old.ID.Int64)))

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("PUT", "By Sistem", strconv.Itoa(int(old.ID.Int64)), old)
	return nil
}

// PartiallyUpdateByID updates the Pengetahuan data for the specified ID with specified parameters.
func (u UseCaseHandler) PartiallyUpdateByID(id string, p *ParamPartiallyUpdate) error {

	// check permission
	err := u.Ctx.ValidatePermission("pengetahuan.edit")
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
	err = tx.Model(&p).Where("id_pengetahuan = ?", old.ID).Updates(p).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint(), strconv.Itoa(int(old.ID.Int64)))

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("PATCH", "By Sistem", strconv.Itoa(int(old.ID.Int64)), old)
	return nil
}

// DeleteByID deletes the Pengetahuan data for the specified ID.
func (u UseCaseHandler) DeleteByID(id string, p *ParamDelete) error {

	// check permission
	err := u.Ctx.ValidatePermission("pengetahuan.delete")
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
	err = tx.Model(&p).Where("id_pengetahuan = ?", old.ID).Update("deleted_at", time.Now().UTC()).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint(), strconv.Itoa(int(old.ID.Int64)))

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("DELETE", "By Sistem", strconv.Itoa(int(old.ID.Int64)), old)
	return nil
}

// setDefaultValue set default value of undefined field when create or update Pengetahuan data.
func (u *UseCaseHandler) setDefaultValue(old Pengetahuan) error {
	if old.ID.Valid {
		u.ID = old.ID
	}

	return nil
}
