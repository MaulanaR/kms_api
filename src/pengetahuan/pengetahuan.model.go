package pengetahuan

import (
	"github.com/maulanar/kms/app"
	"github.com/maulanar/kms/src/tpengetahuanrelation"
)

// Pengetahuan is the main model of Pengetahuan data. It provides a convenient interface for app.ModelInterface
type Pengetahuan struct {
	app.Model
	ID                     app.NullInt64                                      `json:"id"                       db:"m.id_pengetahuan"                    gorm:"column:id_pengetahuan;primaryKey"`
	JenisPengetahuanID     app.NullInt64                                      `json:"jenis_pengetahuan.id"     db:"m.id_jenis_pengetahuan"              gorm:"column:id_jenis_pengetahuan"`
	JenisPengtahuanNama    app.NullText                                       `json:"jenis_pengetahuan.nama"   db:"sjp.nama_jenis_pengetahuan"          gorm:"-"`
	LingkupPengetahuanID   app.NullInt64                                      `json:"lingkup_pengetahuan.id"   db:"m.id_lingkup_pengetahuan"            gorm:"column:id_lingkup_pengetahuan"`
	LingkupPengetahuanNama app.NullText                                       `json:"lingkup_pengetahuan.nama" db:"lp.nama_lingkup_pengetahuan"         gorm:"-"`
	StatusPengetahuanID    app.NullInt64                                      `json:"status_pengetahuan.id"    db:"m.id_status_pengetahuan"             gorm:"column:id_status_pengetahuan"`
	StatusPengetahuanNama  app.NullText                                       `json:"status_pengetahuan.nama"  db:"status.nama_status_pengetahuan"      gorm:"-"`
	Judul                  app.NullText                                       `json:"judul"                    db:"m.judul"                             gorm:"column:judul"`
	Ringkasan              app.NullText                                       `json:"ringkasan"                db:"m.ringkasan"                         gorm:"column:ringkasan"`
	Thumbnail              app.NullString                                     `json:"thumbnail"                db:"m.thumbnail"                         gorm:"column:thumbnail"`
	Penulis1ID             app.NullInt64                                      `json:"penulis_1.id"             db:"m.penulis_1"                         gorm:"column:penulis_1"`
	Penulis1Nama           app.NullString                                     `json:"penulis_1.nama"           db:"p1.nama"                             gorm:"-"`
	Penulis1Jabatan        app.NullString                                     `json:"penulis_1.jabatan"        db:"p1.jabatan"                          gorm:"-"`
	Penulis1Foto           app.NullString                                     `json:"penulis_1.foto"           db:"p1.foto"                             gorm:"-"`
	Penulis2ID             app.NullInt64                                      `json:"penulis_2.id"             db:"m.penulis_2"                         gorm:"column:penulis_2"`
	Penulis2Nama           app.NullString                                     `json:"penulis_2.nama"           db:"p2.nama"                             gorm:"-"`
	Penulis2Jabatan        app.NullString                                     `json:"penulis_2.jabatan"        db:"p2.jabatan"                          gorm:"-"`
	Penulis2Foto           app.NullString                                     `json:"penulis_2.foto"           db:"p2.foto"                             gorm:"-"`
	Penulis3ID             app.NullInt64                                      `json:"penulis_3.id"             db:"m.penulis_3"                         gorm:"column:penulis_3"`
	Penulis3Nama           app.NullString                                     `json:"penulis_3.nama"           db:"p3.nama"                             gorm:"-"`
	Penulis3Jabatan        app.NullString                                     `json:"penulis_3.jabatan"        db:"p3.jabatan"                          gorm:"-"`
	Penulis3Foto           app.NullString                                     `json:"penulis_3.foto"           db:"p3.foto"                             gorm:"-"`
	CreatedBy              app.NullInt64                                      `json:"created_by.id"            db:"m.created_by"                        gorm:"column:created_by"`
	CreatedByUsername      app.NullString                                     `json:"created_by.username"      db:"cbuser.username"                     gorm:"-"`
	CreatedByNip           app.NullString                                     `json:"created_by.nip"           db:"cbuser.nip"                          gorm:"-"`
	CreatedByJabatan       app.NullString                                     `json:"created_by.jabatan"       db:"cbuser.jabatan"                      gorm:"-"`
	UpdatedBy              app.NullInt64                                      `json:"updated_by.id"            db:"m.updated_by"                        gorm:"column:updated_by"`
	UpdatedByUsername      app.NullString                                     `json:"updated_by.username"      db:"ubuser.username"                     gorm:"-"`
	UpdatedByNip           app.NullString                                     `json:"updated_by.nip"           db:"ubuser.nip"                          gorm:"-"`
	UpdatedByJabatan       app.NullString                                     `json:"updated_by.jabatan"       db:"ubuser.jabatan"                      gorm:"-"`
	DeletedBy              app.NullInt64                                      `json:"deleted_by.id"            db:"m.deleted_by"                        gorm:"column:deleted_by"`
	DeletedByUsername      app.NullString                                     `json:"deleted_by.username"      db:"dbuser.username"                     gorm:"-"`
	DeletedByNip           app.NullString                                     `json:"deleted_by.nip"           db:"dbuser.nip"                          gorm:"-"`
	DeletedByJabatan       app.NullString                                     `json:"deleted_by.jabatan"       db:"dbuser.jabatan"                      gorm:"-"`
	CreatedAt              app.NullDateTime                                   `json:"created_at"               db:"m.created_at"                        gorm:"column:created_at"`
	UpdatedAt              app.NullDateTime                                   `json:"updated_at"               db:"m.updated_at"                        gorm:"column:updated_at"`
	DeletedAt              app.NullDateTime                                   `json:"deleted_at"               db:"m.deleted_at,hide"                   gorm:"column:deleted_at"`
	Akademi                []tpengetahuanrelation.TPengetahuanAkademi         `json:"akademi"                  db:"tpa.id_pengetahuan=id"               gorm:"-"`
	PenulisExternal        []tpengetahuanrelation.TPengetahuanPenulisExternal `json:"penulis_external"         db:"tppenulisexternal.id_pengetahuan=id" gorm:"-"`
	Tag                    []tpengetahuanrelation.TPengetahuanTag             `json:"tag"                      db:"tptag.id_pengetahuan=id"             gorm:"-"`
	Referensi              []tpengetahuanrelation.TPengetahuanReferensi       `json:"referensi"                db:"tpref.id_pengetahuan=id"             gorm:"-"`
	Kompetensi             []tpengetahuanrelation.TPengetahuanKompetensi      `json:"kompetensi"               db:"tpkompetensi.id_pengetahuan=id"      gorm:"-"`
	//jenis
	//tugas
	Tujuan         app.NullText `json:"tujuan"                   db:"tptugas.tujuan"                      gorm:"-"`
	DasarHukum     app.NullText `json:"dasar_hukum"              db:"tptugas.dasar_hukum"                 gorm:"-"`
	ProsesBisnis   app.NullText `json:"proses_bisnis"            db:"tptugas.proses_bisnis"               gorm:"-"`
	RumusanMasalah app.NullText `json:"rumusan_masalah"          db:"tptugas.rumusan_masalah"             gorm:"-"`
	PenyebabTemuan app.NullText `json:"penyebab_temuan"          db:"tptugas.penyebab_temuan"             gorm:"-"`
	Keahlian       app.NullText `json:"keahlian"                 db:"tptugas.keahlian"                    gorm:"-"`
	KebutuhanData  app.NullText `json:"kebutuhan_data"           db:"tptugas.kebutuhan_data"              gorm:"-"`
	TenagaAhli     app.NullText `json:"tenaga_ahli"              db:"tptugas.tenaga_ahli"                 gorm:"-"`
	Pedoman        app.NullText `json:"pedoman"                  db:"tptugas.pedoman"                     gorm:"-"`
	//Kiat
	Masalah        app.NullText `json:"masalah"                  db:"tpkiat.masalah"                      gorm:"-"`
	Dampak         app.NullText `json:"dampak"                   db:"tpkiat.dampak"                       gorm:"-"`
	Penyebab       app.NullText `json:"penyebab"                 db:"tpkiat.penyebab"                     gorm:"-"`
	Solusi         app.NullText `json:"solusi"                   db:"tpkiat.solusi"                       gorm:"-"`
	HasilPerbaikan app.NullText `json:"hasil_perbaikan"          db:"tpkiat.hasil_perbaikan"              gorm:"-"`
	//kapitalisasi
	Diskusi app.NullText `json:"diskusi"                  db:"tpk.diskusi"                         gorm:"-"`
	Kapus   app.NullText `json:"kapus"                    db:"tpk.kapus"                           gorm:"-"`
}

// EndPoint returns the Pengetahuan end point, it used for cache key, etc.
func (Pengetahuan) EndPoint() string {
	return "pengetahuan"
}

// TableVersion returns the versions of the Pengetahuan table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Pengetahuan) TableVersion() string {
	return "28.07.011152"
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
	m.AddRelation("left", "m_user", "ubuser", []map[string]any{{"column1": "ubuser.id_user", "column2": "m.updated_by"}})
	m.AddRelation("left", "m_user", "dbuser", []map[string]any{{"column1": "dbuser.id_user", "column2": "m.deleted_by"}})

	//jenis pengetahuan
	m.AddRelation("left", "m_jenis_pengetahuan", "sjp", []map[string]any{{"column1": "sjp.id_jenis_pengetahuan", "column2": "m.id_jenis_pengetahuan"}})

	//lingkup pengetahuan
	m.AddRelation("left", "m_lingkup_pengetahuan", "lp", []map[string]any{{"column1": "lp.id_lingkup_pengetahuan", "column2": "m.id_lingkup_pengetahuan"}})

	//status pengetahuan
	m.AddRelation("left", "m_status_pengetahuan", "status", []map[string]any{{"column1": "status.id_status_pengetahuan", "column2": "m.id_status_pengetahuan"}})

	//penulis 1
	m.AddRelation("left", "m_orang", "p1", []map[string]any{{"column1": "p1.id_orang", "column2": "m.penulis_1"}})

	//penulis 2
	m.AddRelation("left", "m_orang", "p2", []map[string]any{{"column1": "p2.id_orang", "column2": "m.penulis_2"}})

	//penulis 3
	m.AddRelation("left", "m_orang", "p3", []map[string]any{{"column1": "p3.id_orang", "column2": "m.penulis_3"}})

	//tugas
	m.AddRelation("left", "t_pengetahuan_tugas", "tptugas", []map[string]any{{"column1": "tptugas.id_pengetahuan", "column2": "m.id_pengetahuan"}})

	//kiat
	m.AddRelation("left", "t_pengetahuan_kiat", "tpkiat", []map[string]any{{"column1": "tpkiat.id_pengetahuan", "column2": "m.id_pengetahuan"}})

	//kapitalisasi
	m.AddRelation("left", "t_pengetahuan_kapitalisasi", "tpk", []map[string]any{{"column1": "tpk.id_pengetahuan", "column2": "m.id_pengetahuan"}})

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
