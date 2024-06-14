package pengetahuan

import (
	"bytes"
	"embed"
	"encoding/json"
	"html/template"
	"math"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"

	"github.com/maulanar/kms/app"
	"github.com/maulanar/kms/src/attachment"
	"github.com/maulanar/kms/src/event"
	"github.com/maulanar/kms/src/historypoint"
	"github.com/maulanar/kms/src/kompetensi"
	"github.com/maulanar/kms/src/leadertalk"
	"github.com/maulanar/kms/src/lingkuppengetahuan"
	"github.com/maulanar/kms/src/notifikasi"
	"github.com/maulanar/kms/src/orang"
	"github.com/maulanar/kms/src/pedoman"
	"github.com/maulanar/kms/src/referensi"
	"github.com/maulanar/kms/src/statuspengetahuan"
	"github.com/maulanar/kms/src/subjenispengetahuan"
	"github.com/maulanar/kms/src/tag"
	"github.com/maulanar/kms/src/tpengetahuanrelation"
	"github.com/maulanar/kms/src/user"
)

//go:embed template/*
var templatefs embed.FS

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
	app.Cache().Invalidate(u.EndPoint())
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
	// cacheKey := u.EndPoint() + "?" + u.Query.Encode()
	// err = app.Cache().Get(cacheKey, &res)
	// if err == nil {
	// 	return res, err
	// }

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	p := Pengetahuan{}
	if u.Query.Has("tag.id") || u.Query.Has("tag.id.$eq") {
		p.AddRelation("left", "t_pengetahuan_pengetahuan_tag", "tptag", []map[string]any{{"column1": "tptag.id_pengetahuan", "column2": "m.id_pengetahuan"}})
	}

	// set pagination info
	res.Count,
		res.PageContext.Page,
		res.PageContext.PerPage,
		res.PageContext.PageCount,
		err = app.Query().PaginationInfo(tx, &p, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	// return data count if $per_page set to 0
	if res.PageContext.PerPage == 0 {
		return res, err
	}

	// find data
	data, err := app.Query().Find(tx, &p, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	jData, err := json.Marshal(data)
	if err != nil {
		return res, err
	}

	sData := []Pengetahuan{}
	err = json.Unmarshal([]byte(jData), &sData)
	if err != nil {
		return res, err
	}
	for k, d := range sData {
		var isLiked bool
		var isDisliked bool
		//get is liked & is disliked
		tx.Raw("SELECT CASE WHEN COUNT(*) > 0 THEN 1 ELSE 0 END FROM t_like WHERE id_pengetahuan = ? and id_user = ?", d.ID, u.Ctx.User.ID).Scan(&isLiked)
		tx.Raw("SELECT CASE WHEN COUNT(*) > 0 THEN 1 ELSE 0 END FROM t_dislike WHERE id_pengetahuan = ? and id_user = ?", d.ID, u.Ctx.User.ID).Scan(&isDisliked)

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

	// save to cache and return if exists
	// app.Cache().Set(cacheKey, res)
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

	//cek by subjenis
	subjenis, err := subjenispengetahuan.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(p.SubJenisPengetahuanID.Int64)))
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

	//validasi referensi pengetahuan
	if len(p.ReferensiPengetahuan) > 0 {
		for i, ref := range p.ReferensiPengetahuan {
			//validasi
			var countValid int64 = 0
			err = tx.Model(&Pengetahuan{}).Where("id_pengetahuan", ref.RefID.Int64).Count(&countValid).Error
			if err != nil || countValid == 0 {
				return app.Error().New(http.StatusNotFound, u.Ctx.Trans("not_found",
					map[string]string{
						"entity": "Referensi Pengetahuan",
						"key":    "id",
						"value":  strconv.Itoa(int(ref.RefID.Int64)),
					},
				))
			}
			p.ReferensiPengetahuan[i].PengetahuanID.Set(p.ID.Int64)
		}

		err = tx.Create(&p.ReferensiPengetahuan).Error
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
			p.CreatedAt.Set(time.Now())
			p.CreatedBy.Set(u.Ctx.User.ID)
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
	if subjenis.ID.Int64 == 1 {
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
	} else if subjenis.ID.Int64 == 2 {
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
	} else if subjenis.ID.Int64 == 3 {
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
	} else if subjenis.ID.Int64 == 4 {
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

	//Adjustment point

	//ID SUBJENIS
	//KIAT = 2 => 3 poin
	//RESENSI = 4 => 1 poin
	//TUGAS = 1 => 4 poin
	//AKPER = 5 => 5 poin

	var poin int64 = 0
	if p.SubJenisPengetahuanID.Int64 == 2 {
		poin = 3
	} else if p.SubJenisPengetahuanID.Int64 == 4 {
		poin = 1
	} else if p.SubJenisPengetahuanID.Int64 == 1 {
		poin = 4
	} else if p.SubJenisPengetahuanID.Int64 == 5 {
		poin = 5
	}

	if poin > 0 {
		err = historypoint.UseCase(*u.Ctx).AddPoint(p.ID.Int64, poin)
		if err != nil {
			return err
		}
	}

	//Kirim Email broadcast
	if len(p.Tag) > 0 {
		u.SendBroadcast(p)
	}

	// save history (user activity), send webhook, etc
	go notifikasi.UseCase(*u.Ctx).Async(*u.Ctx).SaveNotif("Data Knowledge", "Data Knowledge berhasil ditambah", u.Ctx.User.ID, p.EndPoint(), p.ID.Int64, p)

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint())

	return nil
}

func (u UseCaseHandler) SendBroadcast(p *ParamCreate) error {
	//arr tag
	tags := []int64{}
	for _, t := range p.Tag {
		tags = append(tags, t.TagID.Int64)
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}

	//pengarang
	pengarang, err := user.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(p.CreatedBy.Int64)))
	if err != nil {
		return err
	}

	//cari siapa saja user yang mengikuti tag tsb
	followers := []user.FollowdHastag{}
	err = tx.Model(&user.FollowdHastag{}).Distinct("id_user").Where("id_tag", tags).Find(&followers).Error
	if err == nil {
		//KIRIM BROADCAST

		//prepare body email
		var tmpl, err = template.ParseFS(templatefs, "template/blast.html")
		if err != nil {
			return err
		}

		// emailPenerima := []string{}
		for _, f := range followers {
			flw, err := user.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(f.UserID.Int64)))
			if err == nil {
				// emailPenerima = append(emailPenerima, flw.OrangEmail.String)

				//kirim
				var bodyHtml bytes.Buffer
				var data = make(map[string]interface{})
				data["nama_penerima"] = flw.OrangNama.String
				data["date"] = p.CreatedAt.Time.Format("Monday, 2 January 2006")
				data["pengarang"] = pengarang.OrangNama.String
				data["judul"] = p.Judul.String
				data["link"] = "http://app.rampai.my.id/knowledge-detail?id=" + strconv.Itoa(int(p.ID.Int64)) + "&jenis=" + strconv.Itoa(int(p.JenisPengetahuanID.Int64)) + "&sub=" + strconv.Itoa(int(p.SubJenisPengetahuanID.Int64))
				err = tmpl.Execute(&bodyHtml, data)
				if err != nil {
					return err
				}

				//kirim
				if flw.OrangEmail.Valid {
					go app.SendMail(flw.OrangEmail.String, "Notifikasi Update KMS", bodyHtml.String())
					// if err != nil {
					// 	return err
					// }
				}
			}
		}
		return nil
	}

	return err
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

	p.ID = old.ID
	//cek by subjenis
	subjenis, err := subjenispengetahuan.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(old.SubJenisPengetahuanID.Int64)))
	if err != nil {
		return err
	}
	if p.SubJenisPengetahuanID.Valid {
		subjenis, err = subjenispengetahuan.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(p.SubJenisPengetahuanID.Int64)))
		if err != nil {
			return err
		}
	}

	//validasi LingkupPengetahuan
	if p.LingkupPengetahuanID.Valid {
		_, err = lingkuppengetahuan.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(p.LingkupPengetahuanID.Int64)))
		if err != nil {
			return err
		}
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

	// update data on the db
	err = tx.Model(&p).Where("id_pengetahuan = ?", old.ID).Updates(p).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}
	//RELATION

	//validasi referensi
	if len(p.Referensi) > 0 {
		//delete old data
		tx.Where("id_pengetahuan = ?", p.ID.Int64).Delete(&p.Referensi)
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

	//validasi referensi pengetahuan
	if len(p.ReferensiPengetahuan) > 0 {
		//delete old data
		tx.Where("id_pengetahuan = ?", p.ID.Int64).Delete(&p.ReferensiPengetahuan)
		for i, ref := range p.ReferensiPengetahuan {
			//validasi
			var countValid int64 = 0
			err = tx.Model(&Pengetahuan{}).Where("id_pengetahuan", ref.RefID.Int64).Count(&countValid).Error
			if err != nil || countValid == 0 {
				return app.Error().New(http.StatusNotFound, u.Ctx.Trans("not_found",
					map[string]string{
						"entity": "Referensi Pengetahuan",
						"key":    "id",
						"value":  strconv.Itoa(int(ref.RefID.Int64)),
					},
				))
			}
			p.ReferensiPengetahuan[i].PengetahuanID.Set(p.ID.Int64)
		}

		err = tx.Create(&p.ReferensiPengetahuan).Error
		if err != nil {
			return err
		}
	}

	//validasi hastag
	if len(p.Tag) > 0 {
		//delete old data
		tx.Where("id_pengetahuan = ?", p.ID.Int64).Delete(&p.Tag)
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
		//delete old data
		tx.Where("id_pengetahuan = ?", p.ID.Int64).Delete(&p.Kompetensi)
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
		//delete old data
		tx.Where("id_pengetahuan = ?", p.ID.Int64).Delete(&p.Dokumen)
		for i, ref := range p.Dokumen {
			//validasi
			_, err := attachment.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(ref.AttachmentID.Int64)))
			if err != nil {
				return err
			}
			p.Dokumen[i].PengetahuanID.Set(p.ID.Int64)
			p.CreatedAt.Set(time.Now())
			p.CreatedBy.Set(u.Ctx.User.ID)
		}

		err = tx.Create(&p.Dokumen).Error
		if err != nil {
			return err
		}
	}

	//validasi tenaga ahli
	if len(p.TenagaAhli) > 0 {
		//delete old data
		tx.Where("id_pengetahuan = ?", p.ID.Int64).Delete(&p.TenagaAhli)
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
		//delete old data
		tx.Where("id_pengetahuan = ?", p.ID.Int64).Delete(&p.Pedoman)
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
	if subjenis.ID.Int64 == 1 {
		//tugas
		rel := tpengetahuanrelation.TPengetahuanTugas{}
		rel.Tujuan = p.Tujuan
		rel.DasarHukum = p.DasarHukum
		rel.ProsesBisnis = p.ProsesBisnis
		rel.RumusanMasalah = p.RumusanMasalah
		rel.RisikoObjetPengawasan = p.RisikoObjetPengawasan
		rel.MetodePengawasan = p.MetodePengawasan
		rel.TemuanMaterial = p.TemuanMaterial
		rel.KeahlianDibutuhkan = p.KeahlianDibutuhkan
		rel.DataDigunakan = p.DataDigunakan

		err = tx.Where("id_pengetahuan = ?", p.ID.Int64).Updates(&rel).Error
		if err != nil {
			return err
		}
	} else if subjenis.ID.Int64 == 2 {
		rel := tpengetahuanrelation.TPengetahuanKiat{}
		rel.PengetahuanID.Set(p.ID.Int64)
		rel.Masalah = p.Masalah
		rel.Dampak = p.Dampak
		rel.Penyebab = p.Penyebab
		rel.Solusi = p.Solusi
		rel.SyaratHasil = p.SyaratHasil

		err = tx.Where("id_pengetahuan = ?", p.ID.Int64).Updates(&rel).Error
		if err != nil {
			return err
		}
	} else if subjenis.ID.Int64 == 3 {
		rel := tpengetahuanrelation.TPengetahuanKapitalisasi{}
		rel.PengetahuanID.Set(p.ID.Int64)
		rel.LatarBelakang = p.LatarBelakang
		rel.PenelitianTerdahulu = p.PenelitianTerdahulu
		rel.Hipotesis = p.Hipotesis
		rel.Pengujian = p.Pengujian
		rel.Pembahasan = p.Pembahasan
		rel.KesimpulanRekomendasi = p.KesimpulanRekomendasi

		err = tx.Where("id_pengetahuan = ?", p.ID.Int64).Updates(&rel).Error
		if err != nil {
			return err
		}
	} else if subjenis.ID.Int64 == 4 {
		rel := tpengetahuanrelation.TPengetahuanResensi{}
		rel.PengetahuanID.Set(p.ID.Int64)
		rel.JumlahHalaman = p.JumlahHalaman
		rel.TahunTerbit = p.TahunTerbit
		rel.LatarBelakang = p.LatarBelakang
		rel.PenelitianTerdahulu = p.PenelitianTerdahulu
		rel.LessonLearned = p.LessonLearned

		err = tx.Where("id_pengetahuan = ?", p.ID.Int64).Updates(&rel).Error
		if err != nil {
			return err
		}

		if len(p.Penerbit) > 0 {
			//delete old data
			tx.Where("id_pengetahuan = ?", p.ID.Int64).Delete(&p.Penerbit)
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
			//delete old data
			tx.Where("id_pengetahuan = ?", p.ID.Int64).Delete(&p.Narasumber)
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

	p.ID = old.ID
	//cek by subjenis
	subjenis, err := subjenispengetahuan.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(old.SubJenisPengetahuanID.Int64)))
	if err != nil {
		return err
	}
	if p.SubJenisPengetahuanID.Valid {
		subjenis, err = subjenispengetahuan.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(p.SubJenisPengetahuanID.Int64)))
		if err != nil {
			return err
		}
	}

	//validasi LingkupPengetahuan
	if p.LingkupPengetahuanID.Valid {
		_, err = lingkuppengetahuan.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(p.LingkupPengetahuanID.Int64)))
		if err != nil {
			return err
		}
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

	// update data on the db
	err = tx.Model(&p).Where("id_pengetahuan = ?", old.ID).Updates(p).Error
	if err != nil {
		return app.Error().New(http.StatusInternalServerError, err.Error())
	}
	//RELATION

	//validasi referensi
	if len(p.Referensi) > 0 {
		//delete old data
		tx.Where("id_pengetahuan = ?", p.ID.Int64).Delete(&p.Referensi)
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

	//validasi hastag
	if len(p.Tag) > 0 {
		//delete old data
		tx.Where("id_pengetahuan = ?", p.ID.Int64).Delete(&p.Tag)
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
		//delete old data
		tx.Where("id_pengetahuan = ?", p.ID.Int64).Delete(&p.Kompetensi)
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
		//delete old data
		tx.Where("id_pengetahuan = ?", p.ID.Int64).Delete(&p.Dokumen)
		for i, ref := range p.Dokumen {
			//validasi
			_, err := attachment.UseCase(*u.Ctx).GetByID(strconv.Itoa(int(ref.AttachmentID.Int64)))
			if err != nil {
				return err
			}
			p.Dokumen[i].PengetahuanID.Set(p.ID.Int64)
			p.CreatedAt.Set(time.Now())
			p.CreatedBy.Set(u.Ctx.User.ID)
		}

		err = tx.Create(&p.Dokumen).Error
		if err != nil {
			return err
		}
	}

	//validasi tenaga ahli
	if len(p.TenagaAhli) > 0 {
		//delete old data
		tx.Where("id_pengetahuan = ?", p.ID.Int64).Delete(&p.TenagaAhli)
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
		//delete old data
		tx.Where("id_pengetahuan = ?", p.ID.Int64).Delete(&p.Pedoman)
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
	if subjenis.ID.Int64 == 1 {
		//tugas
		rel := tpengetahuanrelation.TPengetahuanTugas{}
		rel.Tujuan = p.Tujuan
		rel.DasarHukum = p.DasarHukum
		rel.ProsesBisnis = p.ProsesBisnis
		rel.RumusanMasalah = p.RumusanMasalah
		rel.RisikoObjetPengawasan = p.RisikoObjetPengawasan
		rel.MetodePengawasan = p.MetodePengawasan
		rel.TemuanMaterial = p.TemuanMaterial
		rel.KeahlianDibutuhkan = p.KeahlianDibutuhkan
		rel.DataDigunakan = p.DataDigunakan

		err = tx.Where("id_pengetahuan = ?", p.ID.Int64).Updates(&rel).Error
		if err != nil {
			return err
		}
	} else if subjenis.ID.Int64 == 2 {
		rel := tpengetahuanrelation.TPengetahuanKiat{}
		rel.PengetahuanID.Set(p.ID.Int64)
		rel.Masalah = p.Masalah
		rel.Dampak = p.Dampak
		rel.Penyebab = p.Penyebab
		rel.Solusi = p.Solusi
		rel.SyaratHasil = p.SyaratHasil

		err = tx.Where("id_pengetahuan = ?", p.ID.Int64).Updates(&rel).Error
		if err != nil {
			return err
		}
	} else if subjenis.ID.Int64 == 3 {
		rel := tpengetahuanrelation.TPengetahuanKapitalisasi{}
		rel.PengetahuanID.Set(p.ID.Int64)
		rel.LatarBelakang = p.LatarBelakang
		rel.PenelitianTerdahulu = p.PenelitianTerdahulu
		rel.Hipotesis = p.Hipotesis
		rel.Pengujian = p.Pengujian
		rel.Pembahasan = p.Pembahasan
		rel.KesimpulanRekomendasi = p.KesimpulanRekomendasi

		err = tx.Where("id_pengetahuan = ?", p.ID.Int64).Updates(&rel).Error
		if err != nil {
			return err
		}
	} else if subjenis.ID.Int64 == 4 {
		rel := tpengetahuanrelation.TPengetahuanResensi{}
		rel.PengetahuanID.Set(p.ID.Int64)
		rel.JumlahHalaman = p.JumlahHalaman
		rel.TahunTerbit = p.TahunTerbit
		rel.LatarBelakang = p.LatarBelakang
		rel.PenelitianTerdahulu = p.PenelitianTerdahulu
		rel.LessonLearned = p.LessonLearned

		err = tx.Where("id_pengetahuan = ?", p.ID.Int64).Updates(&rel).Error
		if err != nil {
			return err
		}

		if len(p.Penerbit) > 0 {
			//delete old data
			tx.Where("id_pengetahuan = ?", p.ID.Int64).Delete(&p.Penerbit)
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
			//delete old data
			tx.Where("id_pengetahuan = ?", p.ID.Int64).Delete(&p.Narasumber)
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

	// find data pengetahuan
	data, err := app.Query().Find(tx, &SearchPengetahuan{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	res.SetData(data, u.Query)
	if u.Query.Has("levenshtein.keyword.$eq") {
		newData := []map[string]any{}
		keyword := u.Query.Get("levenshtein.keyword.$eq")
		// do for levenshtein
		listJudul := []string{}
		for _, v := range data {
			_, ok := v["judul"].(string)
			if ok {
				_, ok2 := v["ringkasan"].(string)
				if ok2 {
					// listJudul = append(listJudul, v["judul"].(string)+" "+app.RemoveHTMLTags(v["ringkasan"].(string)))
					listJudul = append(listJudul, v["judul"].(string))
				} else {
					listJudul = append(listJudul, v["judul"].(string))
				}
			} else {
				listJudul = append(listJudul, "")
			}
		}
		rnk := app.FindSimilarStrings(keyword, listJudul)
		sort.Slice(rnk, func(i, j int) bool {
			return rnk[i].Score > rnk[j].Score
		})

		r := 0
		for _, match := range rnk {
			//treshold = 0, jika score dibawah 0, maka tidak dianggap
			if match.Score > 0 {
				data[match.Index]["levenshtein.keyword"] = keyword
				data[match.Index]["rank"] = r
				data[match.Index]["tipe"] = "pengetahuan"
				newData = append(newData, data[match.Index])
				r++
			}
		}
		res.SetData(newData, u.Query)
	} else {
		//sisipkan tipe
		for i, _ := range res.Data {
			res.Data[i]["tipe"] = "pengetahuan"
		}
	}

	//data cop
	// find data
	data2, err := app.Query().Find(tx, &SearchForum{}, u.Query)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	if u.Query.Has("levenshtein.keyword.$eq") {
		newData := []map[string]any{}
		keyword := u.Query.Get("levenshtein.keyword.$eq")
		// do for levenshtein
		listJudul := []string{}
		for _, v := range data2 {
			_, ok := v["judul"].(string)
			if ok {
				_, ok2 := v["ringkasan"].(string)
				if ok2 {
					// listJudul = append(listJudul, v["judul"].(string)+" "+app.RemoveHTMLTags(v["ringkasan"].(string)))
					listJudul = append(listJudul, v["judul"].(string))
				} else {
					listJudul = append(listJudul, v["judul"].(string))
				}
			} else {
				listJudul = append(listJudul, "")
			}
		}
		rnk := app.FindSimilarStrings(keyword, listJudul)
		sort.Slice(rnk, func(i, j int) bool {
			return rnk[i].Score > rnk[j].Score
		})

		r := 0
		for _, match := range rnk {
			if match.Score > 0 {
				data2[match.Index]["levenshtein.keyword"] = keyword
				data2[match.Index]["rank"] = r
				data2[match.Index]["tipe"] = "cop"
				newData = append(newData, data2[match.Index])
				r++
			}
		}
		res.Data = append(res.Data, newData...)
	} else {
		for i, _ := range data2 {
			data2[i]["tipe"] = "cop"
		}
		res.Data = append(res.Data, data2...)
	}

	//order agar konsisten
	sortKey := "created_at"
	// Fungsi untuk membandingkan elemen-elemen berdasarkan kunci tertentu
	comparator := func(i, j int) bool {
		iValue, ok := res.Data[i][sortKey].(time.Time)
		if !ok {
			// Jika nilai tidak valid, anggap sebagai waktu paling awal
			return false
		}

		jValue, ok := res.Data[j][sortKey].(time.Time)
		if !ok {
			// Jika nilai tidak valid, anggap sebagai waktu paling awal
			return true
		}

		// Urutkan secara descending (terbaru ke terlama)
		return iValue.After(jValue)
	}

	if u.Query.Has("levenshtein.keyword.$eq") {
		sortKey = "rank"
		// Fungsi untuk membandingkan elemen-elemen berdasarkan kunci tertentu
		comparator = func(i, j int) bool {
			return res.Data[i][sortKey].(int) < res.Data[j][sortKey].(int)
		}
	}

	// Menggunakan sort.Slice untuk mengurutkan slice berdasarkan kunci
	sort.Slice(res.Data, comparator)

	//pagination
	perPage := 10
	if u.Query.Has("$per_page") {
		xperPage, err := strconv.Atoi(u.Query.Get("$per_page"))
		if err == nil {
			perPage = xperPage
		}
	}

	paging := 1
	if u.Query.Has("$page") {
		xpaging, err := strconv.Atoi(u.Query.Get("$page"))
		if err == nil {
			paging = xpaging
		}
	}

	totalData := len(res.Data)
	res.PageContext.Page = paging
	res.PageContext.PerPage = perPage
	res.PageContext.PageCount = int(math.Ceil(float64(totalData) / float64(perPage)))

	startIndex := (paging - 1) * perPage
	endIndex := paging * perPage
	if endIndex > totalData {
		endIndex = totalData
	}

	// Menampilkan data pada halaman saat ini
	dataPaging := []map[string]any{}
	for i := startIndex; i < endIndex; i++ {
		dataPaging = append(dataPaging, res.Data[i])
	}

	res.Data = dataPaging
	res.Count = int64(len(res.Data))
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
	// subjenis := []subjenispengetahuan.SubjenisPengetahuan{}
	// err = tx.Model(&subjenispengetahuan.SubjenisPengetahuan{}).Find(&subjenis).Error
	// if err != nil {
	// 	return res, err
	// }

	// for _, sjp := range subjenis {
	p := []Pengetahuan{}
	q := url.Values{}
	q.Add("subjenis_pengetahuan.id", "4")
	q.Add("$per_page", "20")
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
	// }

	return res, err
}

// get return 4 data of latest pengetahuan, 4 data of leader talk, 4 data of events
// Get returns the list of Pengetahuan data per jenis.
func (u UseCaseHandler) GetMixSlider() (MixSlide, error) {
	res := MixSlide{}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// find data pengetahuan
	q := url.Values{}
	q.Add("$per_page", "4")
	q.Add("$include", "all")
	dataPengetahuan, err := app.Query().Find(tx, &Pengetahuan{}, q)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	res.Pengetahuan = dataPengetahuan

	// find data leader talk
	dataLeaderTalk, err := app.Query().Find(tx, &leadertalk.LeaderTalk{}, q)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	res.LeaderTalk = dataLeaderTalk

	// find data events
	dataEvents, err := app.Query().Find(tx, &event.Event{}, q)
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}
	res.Events = dataEvents

	return res, err
}

// get return popular pengetahuan, based on higher like and view
func (u UseCaseHandler) GetPopuler() ([]Pengetahuan, error) {
	res := []Pengetahuan{}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.Error().New(http.StatusInternalServerError, err.Error())
	}

	// find data
	p := []Pengetahuan{}
	pFilter := Pengetahuan{}
	q := url.Values{}
	q.Add("$per_page", "20")

	pFilter.Sorts = append(pFilter.Sorts, map[string]any{"column": "(`statistik.view`)", "direction": "desc"})
	pFilter.Sorts = append(pFilter.Sorts, map[string]any{"column": "(`statistik.like`)", "direction": "desc"})

	r, err := app.Query().Find(tx, &pFilter, q)
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

	return res, err
}
