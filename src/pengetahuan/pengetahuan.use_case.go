package pengetahuan

import (
	"encoding/json"
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
	"github.com/maulanar/kms/src/pedoman"
	"github.com/maulanar/kms/src/referensi"
	"github.com/maulanar/kms/src/statuspengetahuan"
	"github.com/maulanar/kms/src/subjenispengetahuan"
	"github.com/maulanar/kms/src/tag"
	"github.com/maulanar/kms/src/tpengetahuanrelation"

	"github.com/lithammer/fuzzysearch/fuzzy"
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
	//cacheKey := u.EndPoint() + "." + id
	// app.Cache().Get(cacheKey, &res)
	// if res.ID.Valid {
	// 	return res, err
	// }

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

	//get is liked & is disliked
	tx.Raw("SELECT CASE WHEN COUNT(*) > 0 THEN 1 ELSE 0 END FROM t_like WHERE id_pengetahuan = ? and id_user = ?", id, u.Ctx.User.ID).Scan(&res.IsLiked)
	tx.Raw("SELECT CASE WHEN COUNT(*) > 0 THEN 1 ELSE 0 END FROM t_dislike WHERE id_pengetahuan = ? and id_user = ?", id, u.Ctx.User.ID).Scan(&res.IsDisliked)

	//update count view
	tx.Exec("UPDATE t_pengetahuan SET count_view = count_view + 1 WHERE id_pengetahuan = ?", id)

	// save to cache and return if exists
	// app.Cache().Set(cacheKey, res)
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

	for k, d := range data {
		var isLiked bool
		var isDisliked bool
		//get is liked & is disliked
		tx.Raw("SELECT CASE WHEN COUNT(*) > 0 THEN 1 ELSE 0 END FROM t_like WHERE id_pengetahuan = ? and id_user = ?", d["id"].(int32), u.Ctx.User.ID).Scan(&isLiked)
		tx.Raw("SELECT CASE WHEN COUNT(*) > 0 THEN 1 ELSE 0 END FROM t_dislike WHERE id_pengetahuan = ? and id_user = ?", d["id"].(int32), u.Ctx.User.ID).Scan(&isDisliked)

		data[k]["is_liked"] = isLiked
		data[k]["is_disliked"] = isDisliked
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

	p.StatistikView.Set(0)

	//cek by jenis
	jenis, err := jenispengetahuan.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(p.JenisPengetahuanID.Int64)))
	if err != nil {
		return err
	}

	//validasi LingkupPengetahuan
	_, err = lingkuppengetahuan.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(p.LingkupPengetahuanID.Int64)))
	if err != nil {
		return err
	}

	//validasi StatusPengetahuan
	if !p.StatusPengetahuanID.Valid {
		p.StatusPengetahuanID.Set(1)
	}

	_, err = statuspengetahuan.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(p.StatusPengetahuanID.Int64)))
	if err != nil {
		return err
	}

	//validasi penulis (orang)
	if p.Penulis1ID.Valid {
		_, err = orang.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(p.Penulis1ID.Int64)))
		if err != nil {
			return err
		}
	}

	if p.Penulis2ID.Valid {
		_, err = orang.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(p.Penulis2ID.Int64)))
		if err != nil {
			return err
		}
	}

	if p.Penulis3ID.Valid {
		_, err = orang.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(p.Penulis3ID.Int64)))
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
		for i, ref := range p.Referensi {
			//validasi
			_, err := referensi.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(ref.ReferensiID.Int64)))
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
		for i, ref := range p.Tag {
			//validasi
			_, err := tag.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(ref.TagID.Int64)))
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
		for i, ref := range p.Kompetensi {
			//validasi
			_, err := kompetensi.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(ref.KompetensiID.Int64)))
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
		for i, ref := range p.Dokumen {
			//validasi
			_, err := attachment.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(ref.AttachmentID.Int64)))
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

	//validasi tenaga ahli
	if len(p.TenagaAhli) > 0 {
		for i, ref := range p.TenagaAhli {
			//validasi
			_, err := orang.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(ref.TenagaAhliID.Int64)))
			if err != nil {
				return err
			}
			p.TenagaAhli[i].PengetahuanID.Set(p.ID.Int64)
		}

		err = tx.Create(&p.TenagaAhli).Error
		if err != nil {
			return err
		}
	}

	//validasi pedoman
	if len(p.Pedoman) > 0 {
		for i, v := range p.Pedoman {
			//validasi
			_, err := pedoman.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(v.PedomanID.Int64)))
			if err != nil {
				return err
			}
			p.Pedoman[i].PengetahuanID.Set(p.ID.Int64)
		}

		err = tx.Create(&p.Pedoman).Error
		if err != nil {
			return err
		}
	}

	//NOTE :
	// 1 : Tugas (Panduan Penugasan)
	// 2 : KIAT
	// 3 : Kapitalisasi / Analytic Today
	// 4 : Resensi
	// 5 : Aksi Perubahan
	// 6 : PKS (Pelatihan Kantor Sendiri)
	// 7 : Karya Tulis
	// 8 : Newsletter LC
	if jenis.ID.Int64 == 1 {
		//tugas
		rel := tpengetahuanrelation.TPengetahuanTugas{}
		rel.PengetahuanID.Set(p.ID.Int64)
		rel.Tujuan = p.Tujuan
		rel.DasarHukum = p.DasarHukum
		rel.ProsesBisnis = p.ProsesBisnis
		rel.RumusanMasalah = p.RumusanMasalah
		rel.RisikoObjetPengawasan = p.RisikoObjetPengawasan
		rel.MetodePengawasan = p.MetodePengawasan
		rel.TemuanMaterial = p.TemuanMaterial
		rel.KeahlianDibutuhkan = p.KeahlianDibutuhkan
		rel.DataDigunakan = p.DataDigunakan

		err = tx.Create(&rel).Error
		if err != nil {
			return err
		}
	} else if jenis.ID.Int64 == 2 {
		rel := tpengetahuanrelation.TPengetahuanKiat{}
		rel.PengetahuanID.Set(p.ID.Int64)
		rel.Masalah = p.Masalah
		rel.Dampak = p.Dampak
		rel.Penyebab = p.Penyebab
		rel.Solusi = p.Solusi
		rel.SyaratHasil = p.SyaratHasil

		err = tx.Create(&rel).Error
		if err != nil {
			return err
		}
	} else if jenis.ID.Int64 == 3 {
		rel := tpengetahuanrelation.TPengetahuanKapitalisasi{}
		rel.PengetahuanID.Set(p.ID.Int64)
		rel.LatarBelakang = p.LatarBelakang
		rel.PenelitianTerdahulu = p.PenelitianTerdahulu
		rel.Hipotesis = p.Hipotesis
		rel.Pengujian = p.Pengujian
		rel.Pembahasan = p.Pembahasan
		rel.KesimpulanRekomendasi = p.KesimpulanRekomendasi

		err = tx.Create(&rel).Error
		if err != nil {
			return err
		}
	} else if jenis.ID.Int64 == 4 {
		rel := tpengetahuanrelation.TPengetahuanResensi{}
		rel.PengetahuanID.Set(p.ID.Int64)
		rel.JumlahHalaman = p.JumlahHalaman
		rel.TahunTerbit = p.TahunTerbit
		rel.LatarBelakang = p.LatarBelakang
		rel.PenelitianTerdahulu = p.PenelitianTerdahulu
		rel.LessonLearned = p.LessonLearned

		err = tx.Create(&rel).Error
		if err != nil {
			return err
		}

		if len(p.Penerbit) > 0 {
			for i, _ := range p.Penerbit {
				//validasi
				p.Penerbit[i].PengetahuanID.Set(p.ID.Int64)
			}

			err = tx.Create(&p.Penerbit).Error
			if err != nil {
				return err
			}
		}

		if len(p.Narasumber) > 0 {
			for i, _ := range p.Narasumber {
				//validasi
				p.Narasumber[i].PengetahuanID.Set(p.ID.Int64)
			}

			err = tx.Create(&p.Narasumber).Error
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

	if u.Ctx.Action.Method == "POST" {
		u.CreatedBy.Set(u.Ctx.User.ID)
	}

	if u.Ctx.Action.Method == "PUT" || u.Ctx.Action.Method == "PATCH" {
		u.UpdatedBy.Set(u.Ctx.User.ID)
	}

	if u.Ctx.Action.Method == "DELETE" {
		u.DeletedBy.Set(u.Ctx.User.ID)
	}

	return nil
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
	res.Count = int64(len(data))
	if u.Query.Has("levenshtein.keyword.$eq") {
		newData := []map[string]any{}
		keyword := u.Query.Get("levenshtein.keyword.$eq")
		// do for levenshtein
		listJudul := []string{}
		for _, v := range data {
			_, ok := v["judul"].(string)
			if ok {
				listJudul = append(listJudul, v["judul"].(string))
			} else {
				listJudul = append(listJudul, "")
			}
		}
		rnk := fuzzy.RankFindFold(keyword, listJudul)
		for _, v := range rnk {
			var pr float64 = 0
			pr = float64(v.Distance) / float64(len(v.Target)) * 100
			data[v.OriginalIndex]["levenshtein.keyword"] = keyword
			data[v.OriginalIndex]["levenshtein.distance"] = v.Distance
			data[v.OriginalIndex]["levenshtein.percentage"] = int64(pr)

			newData = append(newData, data[v.OriginalIndex])
		}
		res.SetData(newData, u.Query)
		res.Count = int64(len(newData))
	}

	return res, err
}

// Get returns the list of Pengetahuan data per jenis.
func (u UseCaseHandler) GetSlider() ([]Pengetahuan, error) {
	res := []Pengetahuan{}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// find data
	//cek by subjenis
	subjenis := []subjenispengetahuan.SubjenisPengetahuan{}
	err = tx.Model(&subjenispengetahuan.SubjenisPengetahuan{}).Find(&subjenis).Error
	if err != nil {
		return res, err
	}

	for _, sjp := range subjenis {
		p := []Pengetahuan{}
		q := url.Values{}
		q.Add("subjenis_pengetahuan.id", strconv.Itoa(int(sjp.ID.Int64)))
		q.Add("$per_page", strconv.Itoa(1))
		r, err := app.Query().Find(tx, &Pengetahuan{}, q)
		if err != nil {
			return res, err
		}
		b, err := json.Marshal(r)
		if err == nil {
			err = json.Unmarshal(b, &p)
			if err == nil {
				for _, v := range p {
					res = append(res, v)
				}
			}
		}
	}

	return res, err
}
