package pengetahuan

import (
	"github.com/maulanar/kms/app"
	"github.com/maulanar/kms/src/tpengetahuanrelation"
)

// Pengetahuan is the main model of Pengetahuan data. It provides a convenient interface for app.ModelInterface
type Pengetahuan struct {
	app.Model
	//for leveinsthein method
	LevenshteinKeyword    app.NullString  `json:"levenshtein.keyword"          db:"-"                                                                                  gorm:"-"`
	LevenshteinDistance   app.NullInt64   `json:"levenshtein.distance"         db:"-"                                                                                  gorm:"-"`
	LevenshteinPercentage app.NullFloat64 `json:"levenshtein.percentage"       db:"-"                                                                                  gorm:"-"`
	//common data
	ID                       app.NullInt64  `json:"id"                           db:"m.id_pengetahuan"                                                                   gorm:"column:id_pengetahuan;primaryKey"`
	JenisPengetahuanID       app.NullInt64  `json:"jenis_pengetahuan.id"         db:"m.id_jenis_pengetahuan"                                                             gorm:"column:id_jenis_pengetahuan"`
	JenisPengtahuanNama      app.NullText   `json:"jenis_pengetahuan.nama"       db:"jp.nama_jenis_pengetahuan"                                                          gorm:"-"`
	SubJenisPengetahuanID    app.NullInt64  `json:"subjenis_pengetahuan.id"      db:"m.id_subjenis_pengetahuan"                                                          gorm:"column:id_subjenis_pengetahuan"`
	SubJenisPengtahuanNama   app.NullText   `json:"subjenis_pengetahuan.nama"    db:"sjp.nama_subjenis_pengetahuan"                                                      gorm:"-"`
	SubJenisPengtahuanIsShow app.NullBool   `json:"subjenis_pengetahuan.is_show" db:"sjp.is_show"                                                                        gorm:"-"`
	LingkupPengetahuanID     app.NullInt64  `json:"lingkup_pengetahuan.id"       db:"m.id_lingkup_pengetahuan"                                                           gorm:"column:id_lingkup_pengetahuan"`
	LingkupPengetahuanNama   app.NullText   `json:"lingkup_pengetahuan.nama"     db:"lp.nama_lingkup_pengetahuan"                                                        gorm:"-"`
	StatusPengetahuanID      app.NullInt64  `json:"status_pengetahuan.id"        db:"m.id_status_pengetahuan"                                                            gorm:"column:id_status_pengetahuan"`
	StatusPengetahuanNama    app.NullText   `json:"status_pengetahuan.nama"      db:"status.nama_status_pengetahuan"                                                     gorm:"-"`
	Judul                    app.NullText   `json:"judul"                        db:"m.judul"                                                                            gorm:"column:judul"`
	Ringkasan                app.NullText   `json:"ringkasan"                    db:"m.ringkasan"                                                                        gorm:"column:ringkasan"`
	ThumbnailID              app.NullInt64  `json:"thumbnail.id"                 db:"m.thumbnail"                                                                        gorm:"column:thumbnail"`
	ThumbnailName            app.NullString `json:"thumbnail.nama"               db:"attachment.filename"                                                                gorm:"-"`
	ThumbnailUrl             app.NullString `json:"thumbnail.url"                db:"attachment.url"                                                                     gorm:"-"`
	Penulis1ID               app.NullInt64  `json:"penulis_1.id"                 db:"m.penulis_1"                                                                        gorm:"column:penulis_1"`
	Penulis1Nama             app.NullString `json:"penulis_1.nama"               db:"p1.nama"                                                                            gorm:"-"`
	Penulis1Jabatan          app.NullString `json:"penulis_1.jabatan"            db:"p1.jabatan"                                                                         gorm:"-"`
	Penulis1Foto             app.NullString `json:"penulis_1.foto.id"            db:"p1.foto"                                                                            gorm:"-"`
	Penulis1Url              app.NullString `json:"penulis_1.foto.url"           db:"p1attachment.url"                                                                   gorm:"-"`
	Penulis1Filename         app.NullString `json:"penulis_1.foto.nama"          db:"p1attachment.filename"                                                              gorm:"-"`
	Penulis2ID               app.NullInt64  `json:"penulis_2.id"                 db:"m.penulis_2"                                                                        gorm:"column:penulis_2"`
	Penulis2Nama             app.NullString `json:"penulis_2.nama"               db:"p2.nama"                                                                            gorm:"-"`
	Penulis2Jabatan          app.NullString `json:"penulis_2.jabatan"            db:"p2.jabatan"                                                                         gorm:"-"`
	Penulis2Foto             app.NullString `json:"penulis_2.foto.id"            db:"p2.foto"                                                                            gorm:"-"`
	Penulis2Url              app.NullString `json:"penulis_2.foto.url"           db:"p2attachment.url"                                                                   gorm:"-"`
	Penulis2Filename         app.NullString `json:"penulis_2.foto.nama"          db:"p2attachment.filename"                                                              gorm:"-"`
	Penulis3ID               app.NullInt64  `json:"penulis_3.id"                 db:"m.penulis_3"                                                                        gorm:"column:penulis_3"`
	Penulis3Nama             app.NullString `json:"penulis_3.nama"               db:"p3.nama"                                                                            gorm:"-"`
	Penulis3Jabatan          app.NullString `json:"penulis_3.jabatan"            db:"p3.jabatan"                                                                         gorm:"-"`
	Penulis3Foto             app.NullString `json:"penulis_3.foto.id"            db:"p3.foto"                                                                            gorm:"-"`
	Penulis3Url              app.NullString `json:"penulis_3.foto.url"           db:"p3attachment.url"                                                                   gorm:"-"`
	Penulis3Filename         app.NullString `json:"penulis_3.foto.nama"          db:"p3attachment.filename"                                                              gorm:"-"`
	CreatedBy                app.NullInt64  `json:"created_by.id"                db:"m.created_by"                                                                       gorm:"column:created_by"`
	CreatedByUsername        app.NullString `json:"created_by.username"          db:"cbuser.username"                                                                    gorm:"-"`
	CreatedByOrangFotoID     app.NullInt64  `json:"created_by.foto.id"           db:"cbo.foto"                                                                           gorm:"-"`
	CreatedByOrangFotoUrl    app.NullString `json:"created_by.foto.url"          db:"cbatt.url"                                                                          gorm:"-"`
	CreatedByOrangFotoNama   app.NullString `json:"created_by.foto.nama"         db:"cbatt.filename"                                                                     gorm:"-"`

	UpdatedBy         app.NullInt64    `json:"updated_by.id"                db:"m.updated_by"                                                                       gorm:"column:updated_by"`
	UpdatedByUsername app.NullString   `json:"updated_by.username"          db:"ubuser.username"                                                                    gorm:"-"`
	DeletedBy         app.NullInt64    `json:"deleted_by.id"                db:"m.deleted_by"                                                                       gorm:"column:deleted_by"`
	DeletedByUsername app.NullString   `json:"deleted_by.username"          db:"dbuser.username"                                                                    gorm:"-"`
	CreatedAt         app.NullDateTime `json:"created_at"                   db:"m.created_at"                                                                       gorm:"column:created_at"`
	UpdatedAt         app.NullDateTime `json:"updated_at"                   db:"m.updated_at"                                                                       gorm:"column:updated_at"`
	DeletedAt         app.NullDateTime `json:"deleted_at"                   db:"m.deleted_at,hide"                                                                  gorm:"column:deleted_at"`

	//statistik View, like, dislike, komentar
	StatistikView     app.NullInt64 `json:"statistik.view"               db:"(CASE WHEN m.count_view > 0 THEN m.count_view ELSE 0 END)"                          gorm:"column:count_view;default:0"`
	StatistikLike     app.NullInt64 `json:"statistik.like"               db:"(SELECT COUNT(*) FROM t_like WHERE t_like.id_pengetahuan=m.id_pengetahuan)"         gorm:"-"`
	StatistikDislike  app.NullInt64 `json:"statistik.dislike"            db:"(SELECT COUNT(*) FROM t_dislike WHERE t_dislike.id_pengetahuan=m.id_pengetahuan)"   gorm:"-"`
	StatistikKomentar app.NullInt64 `json:"statistik.komentar"           db:"(SELECT COUNT(*) FROM t_komentar WHERE t_komentar.id_pengetahuan=m.id_pengetahuan)" gorm:"-"`
	IsLiked           app.NullBool  `json:"is_liked"                     db:"-"                                                                                  gorm:"-"`
	IsDisliked        app.NullBool  `json:"is_disliked"                  db:"-"                                                                                  gorm:"-"`

	Akademi         []tpengetahuanrelation.TPengetahuanAkademi         `json:"akademi"                      db:"pengetahuan.id={id}"                                                                gorm:"-"`
	PenulisExternal []tpengetahuanrelation.TPengetahuanPenulisExternal `json:"penulis_external"             db:"pengetahuan.id={id}"                                                                gorm:"-"`
	Tag             []tpengetahuanrelation.TPengetahuanTag             `json:"tag"                          db:"pengetahuan.id={id}"                                                                gorm:"-"`
	Referensi       []tpengetahuanrelation.TPengetahuanReferensi       `json:"referensi"                    db:"pengetahuan.id={id}"                                                                gorm:"-"`
	Kompetensi      []tpengetahuanrelation.TPengetahuanKompetensi      `json:"kompetensi"                   db:"pengetahuan.id={id}"                                                                gorm:"-"`
	Dokumen         []tpengetahuanrelation.TPengetahuanDokumen         `json:"dokumen"                      db:"pengetahuan.id={id}"                                                                gorm:"-"`
	TenagaAhli      []tpengetahuanrelation.TPengetahuanTenagaAhli      `json:"tenaga_ahli"                  db:"pengetahuan.id={id}"                                                                gorm:"-"`
	Pedoman         []tpengetahuanrelation.TPengetahuanPedoman         `json:"pedoman"                      db:"pengetahuan.id={id}"                                                                gorm:"-"`
	//jenis

	//tugas
	Tujuan                app.NullText `json:"tujuan"                       db:"tptugas.tujuan"                                                                     gorm:"-"`
	DasarHukum            app.NullText `json:"dasar_hukum"                  db:"tptugas.dasar_hukum"                                                                gorm:"-"`
	ProsesBisnis          app.NullText `json:"proses_bisnis"                db:"tptugas.proses_bisnis"                                                              gorm:"-"`
	RumusanMasalah        app.NullText `json:"rumusan_masalah"              db:"tptugas.rumusan_masalah"                                                            gorm:"-"`
	RisikoObjetPengawasan app.NullText `json:"risiko_objek_pengawasan"      db:"tptugas.risiko_objek_pengawasan"                                                    gorm:"-"`
	MetodePengawasan      app.NullText `json:"metode_pengawasan"            db:"tptugas.metode_pengawasan"                                                          gorm:"-"`
	TemuanMaterial        app.NullText `json:"temuan_material"              db:"tptugas.temuan_material"                                                            gorm:"-"`
	KeahlianDibutuhkan    app.NullText `json:"keahlian_dibutuhkan"          db:"tptugas.keahlian_dibutuhkan"                                                        gorm:"-"`
	DataDigunakan         app.NullText `json:"data_digunakan"               db:"tptugas.data_digunakan"                                                             gorm:"-"`

	//Kiat
	Masalah     app.NullText `json:"masalah"                      db:"tpkiat.masalah"                                                                     gorm:"-"`
	Dampak      app.NullText `json:"dampak"                       db:"tpkiat.dampak"                                                                      gorm:"-"`
	Penyebab    app.NullText `json:"penyebab"                     db:"tpkiat.penyebab"                                                                    gorm:"-"`
	Solusi      app.NullText `json:"solusi"                       db:"tpkiat.solusi"                                                                      gorm:"-"`
	SyaratHasil app.NullText `json:"syarat_hasil"                 db:"tpkiat.syarat_hasil"                                                                gorm:"-"`

	//kapitalisasi
	LatarBelakang         app.NullText `json:"latar_belakang"               db:"COALESCE(tpk.latar_belakang, tp_resensi.latar_belakang)"                            gorm:"-"`
	PenelitianTerdahulu   app.NullText `json:"penelitian_terdahulu"         db:"COALESCE(tpk.penelitian_terdahulu,tp_resensi.penelitian_terdahulu)"                 gorm:"-"`
	Hipotesis             app.NullText `json:"hipotesis"                    db:"tpk.hipotesis"                                                                      gorm:"-"`
	Pengujian             app.NullText `json:"pengujian"                    db:"tpk.pengujian"                                                                      gorm:"-"`
	Pembahasan            app.NullText `json:"pembahasan"                   db:"tpk.pembahasan"                                                                     gorm:"-"`
	KesimpulanRekomendasi app.NullText `json:"kesimpulan_rekomendasi"       db:"tpk.kesimpulan_rekomendasi"                                                         gorm:"-"`

	//resensi
	Narasumber    []tpengetahuanrelation.TPengetahuanNarsum   `json:"narasumber"                   db:"pengetahuan.id={id}"                                                                gorm:"-"`
	JumlahHalaman app.NullInt64                               `json:"jumlah_halaman"               db:"tp_resensi.jumlah_halaman"                                                          gorm:"-"`
	Penerbit      []tpengetahuanrelation.TPengetahuanPenerbit `json:"penerbit"                     db:"pengetahuan.id={id}"                                                                gorm:"-"`
	TahunTerbit   app.NullInt64                               `json:"tahun_terbit"                 db:"tp_resensi.tahun_terbit"                                                            gorm:"-"`
	LessonLearned app.NullText                                `json:"lesson_learned"               db:"tp_resensi.lesson_learned"                                                          gorm:"-"`
}

// EndPoint returns the Pengetahuan end point, it used for cache key, etc.
func (Pengetahuan) EndPoint() string {
	return "pengetahuan"
}

// TableVersion returns the versions of the Pengetahuan table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Pengetahuan) TableVersion() string {
	return "28.07.061152"
}

// TableName returns the name of the Pengetahuan table in the database.
func (Pengetahuan) TableName() string {
	return "t_pengetahuan"
}

// TableAliasName returns the table alias name of the Pengetahuan table, used for querying.
func (Pengetahuan) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Pengetahuan data in the database, used for querying.
func (m *Pengetahuan) GetRelations() map[string]map[string]any {
	//created by, updated by , deleted by
	m.AddRelation("left", "m_user", "cbuser", []map[string]any{{"column1": "cbuser.id_user", "column2": "m.created_by"}})
	m.AddRelation("left", "m_orang", "cbo", []map[string]any{{"column1": "cbo.id_orang", "column2": "cbuser.id_orang"}})
	m.AddRelation("left", "m_attachments", "cbatt", []map[string]any{{"column1": "cbatt.id", "column2": "cbo.foto"}})

	m.AddRelation("left", "m_user", "ubuser", []map[string]any{{"column1": "ubuser.id_user", "column2": "m.updated_by"}})
	m.AddRelation("left", "m_user", "dbuser", []map[string]any{{"column1": "dbuser.id_user", "column2": "m.deleted_by"}})

	//attachment (thumbnail)
	m.AddRelation("left", "m_attachments", "attachment", []map[string]any{{"column1": "attachment.id", "column2": "m.thumbnail"}})

	//jenis pengetahuan
	m.AddRelation("left", "m_jenis_pengetahuan", "jp", []map[string]any{{"column1": "jp.id_jenis_pengetahuan", "column2": "m.id_jenis_pengetahuan"}})
	m.AddRelation("left", "m_subjenis_pengetahuan", "sjp", []map[string]any{{"column1": "sjp.id_subjenis_pengetahuan", "column2": "m.id_subjenis_pengetahuan"}})

	//lingkup pengetahuan
	m.AddRelation("left", "m_lingkup_pengetahuan", "lp", []map[string]any{{"column1": "lp.id_lingkup_pengetahuan", "column2": "m.id_lingkup_pengetahuan"}})

	//status pengetahuan
	m.AddRelation("left", "m_status_pengetahuan", "status", []map[string]any{{"column1": "status.id_status_pengetahuan", "column2": "m.id_status_pengetahuan"}})

	//penulis 1
	m.AddRelation("left", "m_orang", "p1", []map[string]any{{"column1": "p1.id_orang", "column2": "m.penulis_1"}})
	m.AddRelation("left", "m_attachments", "p1attachment", []map[string]any{{"column1": "p1attachment.id", "column2": "p1.foto"}})
	//penulis 2
	m.AddRelation("left", "m_orang", "p2", []map[string]any{{"column1": "p2.id_orang", "column2": "m.penulis_2"}})
	m.AddRelation("left", "m_attachments", "p2attachment", []map[string]any{{"column1": "p2attachment.id", "column2": "p2.foto"}})
	//penulis 3
	m.AddRelation("left", "m_orang", "p3", []map[string]any{{"column1": "p3.id_orang", "column2": "m.penulis_3"}})
	m.AddRelation("left", "m_attachments", "p3attachment", []map[string]any{{"column1": "p3attachment.id", "column2": "p3.foto"}})
	//tugas
	m.AddRelation("left", "t_pengetahuan_tugas", "tptugas", []map[string]any{{"column1": "tptugas.id_pengetahuan", "column2": "m.id_pengetahuan"}})

	//kiat
	m.AddRelation("left", "t_pengetahuan_kiat", "tpkiat", []map[string]any{{"column1": "tpkiat.id_pengetahuan", "column2": "m.id_pengetahuan"}})

	//kapitalisasi
	m.AddRelation("left", "t_pengetahuan_kapitalisasi", "tpk", []map[string]any{{"column1": "tpk.id_pengetahuan", "column2": "m.id_pengetahuan"}})

	//resensi
	m.AddRelation("left", "t_pengetahuan_resensi", "tp_resensi", []map[string]any{{"column1": "tp_resensi.id_pengetahuan", "column2": "m.id_pengetahuan"}})

	return m.Relations
}

// GetFilters returns the filter of the Pengetahuan data in the database, used for querying.
func (m *Pengetahuan) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Pengetahuan data in the database, used for querying.
func (m *Pengetahuan) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Pengetahuan data in the database, used for querying.
func (m *Pengetahuan) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Pengetahuan schema, used for querying.
func (m *Pengetahuan) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Pengetahuan schema in the open api documentation.
func (Pengetahuan) OpenAPISchemaName() string {
	return "Pengetahuan"
}

// GetOpenAPISchema returns the Open API Schema of the Pengetahuan in the open api documentation.
func (m *Pengetahuan) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type PengetahuanList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the PengetahuanList schema in the open api documentation.
func (PengetahuanList) OpenAPISchemaName() string {
	return "PengetahuanList"
}

// GetOpenAPISchema returns the Open API Schema of the PengetahuanList in the open api documentation.
func (p *PengetahuanList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Pengetahuan{})
}

// ParamCreate is the expected parameters for create a new Pengetahuan data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the Pengetahuan data.
type ParamUpdate struct {
	UseCaseHandler
}

// ParamPartiallyUpdate is the expected parameters for partially update the Pengetahuan data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
}

// ParamDelete is the expected parameters for delete the Pengetahuan data.
type ParamDelete struct {
	UseCaseHandler
}

type MixSlide struct {
	Pengetahuan []map[string]any `json:"pengetahuan" db:"-" gorm:"-"`
	LeaderTalk  []map[string]any `json:"leader_talk" db:"-" gorm:"-"`
	Events      []map[string]any `json:"event"       db:"-" gorm:"-"`
}

// Struct Search Pengetahuan
// SearchPengetahuan is the main model of SearchPengetahuan data. It provides a convenient interface for app.ModelInterface
type SearchPengetahuan struct {
	app.Model
	//for leveinsthein method
	LevenshteinKeyword    app.NullString  `json:"levenshtein.keyword"          db:"-"                                                                                  gorm:"-"`
	LevenshteinDistance   app.NullInt64   `json:"levenshtein.distance"         db:"-"                                                                                  gorm:"-"`
	LevenshteinPercentage app.NullFloat64 `json:"levenshtein.percentage"       db:"-"                                                                                  gorm:"-"`
	//common data
	ID            app.NullInt64  `json:"id"                           db:"m.id_pengetahuan"                                                                   gorm:"column:id_pengetahuan;primaryKey"`
	Judul         app.NullText   `json:"judul"                        db:"m.judul"                                                                            gorm:"column:judul"`
	Ringkasan     app.NullText   `json:"ringkasan"                    db:"m.ringkasan"                                                                        gorm:"column:ringkasan"`
	ThumbnailID   app.NullInt64  `json:"thumbnail.id"                 db:"m.thumbnail"                                                                        gorm:"column:thumbnail"`
	ThumbnailName app.NullString `json:"thumbnail.nama"               db:"attachment.filename"                                                                gorm:"-"`
	ThumbnailUrl  app.NullString `json:"thumbnail.url"                db:"attachment.url"                                                                     gorm:"-"`

	JenisPengetahuanID       app.NullInt64 `json:"jenis_pengetahuan.id"         db:"m.id_jenis_pengetahuan"                                                             gorm:"column:id_jenis_pengetahuan"`
	JenisPengtahuanNama      app.NullText  `json:"jenis_pengetahuan.nama"       db:"jp.nama_jenis_pengetahuan"                                                          gorm:"-"`
	SubJenisPengetahuanID    app.NullInt64 `json:"subjenis_pengetahuan.id"      db:"m.id_subjenis_pengetahuan"                                                          gorm:"column:id_subjenis_pengetahuan"`
	SubJenisPengtahuanNama   app.NullText  `json:"subjenis_pengetahuan.nama"    db:"sjp.nama_subjenis_pengetahuan"                                                      gorm:"-"`
	SubJenisPengtahuanIsShow app.NullBool  `json:"subjenis_pengetahuan.is_show" db:"sjp.is_show"                                                                        gorm:"-"`
	LingkupPengetahuanID     app.NullInt64 `json:"lingkup_pengetahuan.id"       db:"m.id_lingkup_pengetahuan"                                                           gorm:"column:id_lingkup_pengetahuan"`
	LingkupPengetahuanNama   app.NullText  `json:"lingkup_pengetahuan.nama"     db:"lp.nama_lingkup_pengetahuan"                                                        gorm:"-"`
	StatusPengetahuanID      app.NullInt64 `json:"status_pengetahuan.id"        db:"m.id_status_pengetahuan"                                                            gorm:"column:id_status_pengetahuan"`
	StatusPengetahuanNama    app.NullText  `json:"status_pengetahuan.nama"      db:"status.nama_status_pengetahuan"                                                     gorm:"-"`
}

// EndPoint returns the SearchPengetahuan end point, it used for cache key, etc.
func (SearchPengetahuan) EndPoint() string {
	return "search_pengetahuan"
}

// TableVersion returns the versions of the SearchPengetahuan table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (SearchPengetahuan) TableVersion() string {
	return "28.07.061152"
}

// TableName returns the name of the SearchPengetahuan table in the database.
func (SearchPengetahuan) TableName() string {
	return "t_pengetahuan"
}

// TableAliasName returns the table alias name of the SearchPengetahuan table, used for querying.
func (SearchPengetahuan) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the SearchPengetahuan data in the database, used for querying.
func (m *SearchPengetahuan) GetRelations() map[string]map[string]any {
	//jenis pengetahuan
	m.AddRelation("left", "m_jenis_pengetahuan", "jp", []map[string]any{{"column1": "jp.id_jenis_pengetahuan", "column2": "m.id_jenis_pengetahuan"}})
	m.AddRelation("left", "m_subjenis_pengetahuan", "sjp", []map[string]any{{"column1": "sjp.id_subjenis_pengetahuan", "column2": "m.id_subjenis_pengetahuan"}})

	//lingkup pengetahuan
	m.AddRelation("left", "m_lingkup_pengetahuan", "lp", []map[string]any{{"column1": "lp.id_lingkup_pengetahuan", "column2": "m.id_lingkup_pengetahuan"}})

	//status pengetahuan
	m.AddRelation("left", "m_status_pengetahuan", "status", []map[string]any{{"column1": "status.id_status_pengetahuan", "column2": "m.id_status_pengetahuan"}})

	//attachment (thumbnail)
	m.AddRelation("left", "m_attachments", "attachment", []map[string]any{{"column1": "attachment.id", "column2": "m.thumbnail"}})
	return m.Relations
}

// GetFilters returns the filter of the SearchPengetahuan data in the database, used for querying.
func (m *SearchPengetahuan) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the SearchPengetahuan data in the database, used for querying.
func (m *SearchPengetahuan) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the SearchPengetahuan data in the database, used for querying.
func (m *SearchPengetahuan) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the SearchPengetahuan schema, used for querying.
func (m *SearchPengetahuan) GetSchema() map[string]any {
	return m.SetSchema(m)
}

type SearchForum struct {
	app.Model
	ID             app.NullInt64  `json:"id"                   db:"m.id"                                                             gorm:"column:id;primaryKey"`
	Judul          app.NullText   `json:"judul"                db:"m.judul"                                                          gorm:"column:judul"`
	Deskripsi      app.NullText   `json:"ringkasan"            db:"m.deskripsi"                                                      gorm:"column:deskripsi"`
	GambarID       app.NullInt64  `json:"thumbnail.id"            db:"m.gambar_id"                                                      gorm:"column:gambar_id"`
	GambarFilename app.NullString `json:"thumbnail.filename"      db:"g.filename"                                                       gorm:"-"`
	GambarUrl      app.NullString `json:"thumbnail.url"           db:"g.url"                                                            gorm:"-"`
}

func (SearchForum) EndPoint() string {
	return "forum"
}

func (SearchForum) TableVersion() string {
	return "23.11.271152"
}

func (SearchForum) TableName() string {
	return "t_forum"
}

func (SearchForum) TableAliasName() string {
	return "m"
}

func (m *SearchForum) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_attachments", "g", []map[string]any{{"column1": "g.id", "column2": "m.gambar_id"}})

	return m.Relations
}

func (m *SearchForum) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *SearchForum) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *SearchForum) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *SearchForum) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (SearchForum) OpenAPISchemaName() string {
	return "SearchForum"
}

func (m *SearchForum) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}
