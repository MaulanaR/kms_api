package tpengetahuanrelation

import "github.com/maulanar/kms/app"

// TPengetahuanAkademi is the main model of TPengetahuanAkademi data. It provides a convenient interface for app.ModelInterface
type TPengetahuanAkademi struct {
	app.Model
	ID            app.NullInt64 `json:"-"              db:"tpa.id_pengetahuan_akademi" gorm:"column:id_pengetahuan_akademi;primaryKey"`
	PengetahuanID app.NullInt64 `json:"pengetahuan.id" db:"tpa.id_pengetahuan,hide"    gorm:"column:id_pengetahuan"`
	AkademiID     app.NullInt64 `json:"id"             db:"tpa.id_akademi"             gorm:"column:id_akademi"`
	AkademiNama   app.NullText  `json:"nama"           db:"tpk.nama_akademi"           gorm:"-"`
}

// EndPoint returns the TPengetahuanAkademi end point, it used for cache key, etc.
func (TPengetahuanAkademi) EndPoint() string {
	return "t_pengetahuan_akademi"
}

// TableVersion returns the versions of the TPengetahuanAkademi table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (TPengetahuanAkademi) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the TPengetahuanAkademi table in the database.
func (TPengetahuanAkademi) TableName() string {
	return "t_pengetahuan_akademi"
}

// TableAliasName returns the table alias name of the TPengetahuanAkademi table, used for querying.
func (TPengetahuanAkademi) TableAliasName() string {
	return "tpa"
}

// GetRelations returns the relations of the TPengetahuanAkademi data in the database, used for querying.
func (m *TPengetahuanAkademi) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_akademi", "tpk", []map[string]any{{"column1": "tpk.id_akademi", "column2": "tpa.id_akademi"}})

	return m.Relations
}

// GetFilters returns the filter of the TPengetahuanAkademi data in the database, used for querying.
func (m *TPengetahuanAkademi) GetFilters() []map[string]any {
	return m.Filters
}

// GetSorts returns the default sort of the TPengetahuanAkademi data in the database, used for querying.
func (m *TPengetahuanAkademi) GetSorts() []map[string]any {
	return m.Sorts
}

// GetFields returns list of the field of the TPengetahuanAkademi data in the database, used for querying.
func (m *TPengetahuanAkademi) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the TPengetahuanAkademi schema, used for querying.
func (m *TPengetahuanAkademi) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the TPengetahuanAkademi schema in the open api documentation.
func (TPengetahuanAkademi) OpenAPISchemaName() string {
	return "TPengetahuanAkademi"
}

// GetOpenAPISchema returns the Open API Schema of the TPengetahuanAkademi in the open api documentation.
func (m *TPengetahuanAkademi) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type TPengetahuanAkademiList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the TPengetahuanAkademiList schema in the open api documentation.
func (TPengetahuanAkademiList) OpenAPISchemaName() string {
	return "TPengetahuanAkademiList"
}

// GetOpenAPISchema returns the Open API Schema of the TPengetahuanAkademiList in the open api documentation.
func (p *TPengetahuanAkademiList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&TPengetahuanAkademi{})
}

// kapitalisasi
type TPengetahuanKapitalisasi struct {
	app.Model
	ID                    app.NullInt64 `json:"-"                      db:"tpk.id_pengetahuan_kapitalisasi" gorm:"column:id_pengetahuan_kapitalisasi;primaryKey"`
	PengetahuanID         app.NullInt64 `json:"pengetahuan.id"         db:"tpk.id_pengetahuan,hide"         gorm:"column:id_pengetahuan"`
	LatarBelakang         app.NullText  `json:"latar_belakang"         db:"tpk.latar_belakang"              gorm:"column:latar_belakang"`
	PenelitianTerdahulu   app.NullText  `json:"penelitian_terdahulu"   db:"tpk.penelitian_terdahulu"        gorm:"column:penelitian_terdahulu"`
	Hipotesis             app.NullText  `json:"hipotesis"              db:"tpk.hipotesis"                   gorm:"column:hipotesis"`
	Pengujian             app.NullText  `json:"pengujian"              db:"tpk.pengujian"                   gorm:"column:pengujian"`
	Pembahasan            app.NullText  `json:"pembahasan"             db:"tpk.pembahasan"                  gorm:"column:pembahasan"`
	KesimpulanRekomendasi app.NullText  `json:"kesimpulan_rekomendasi" db:"tpk.kesimpulan_rekomendasi"      gorm:"column:kesimpulan_rekomendasi"`
}

func (TPengetahuanKapitalisasi) EndPoint() string {
	return "t_pengetahuan_kapitalisasi"
}

func (TPengetahuanKapitalisasi) TableVersion() string {
	return "28.08.081152"
}

func (TPengetahuanKapitalisasi) TableName() string {
	return "t_pengetahuan_kapitalisasi"
}

func (TPengetahuanKapitalisasi) TableAliasName() string {
	return "tpk"
}

func (m *TPengetahuanKapitalisasi) GetRelations() map[string]map[string]any {
	return m.Relations
}

func (m *TPengetahuanKapitalisasi) GetFilters() []map[string]any {
	return m.Filters
}

func (m *TPengetahuanKapitalisasi) GetSorts() []map[string]any {
	return m.Sorts
}

func (m *TPengetahuanKapitalisasi) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *TPengetahuanKapitalisasi) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (TPengetahuanKapitalisasi) OpenAPISchemaName() string {
	return "TPengetahuanKapitalisasi"
}

func (m *TPengetahuanKapitalisasi) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type TPengetahuanKapitalisasiList struct {
	app.ListModel
}

func (TPengetahuanKapitalisasiList) OpenAPISchemaName() string {
	return "TPengetahuanKapitalisasiList"
}

func (p *TPengetahuanKapitalisasiList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&TPengetahuanKapitalisasi{})
}

// poengetahuan kiat
type TPengetahuanKiat struct {
	app.Model
	ID            app.NullInt64 `json:"-"              db:"tpkiat.id_kiat"             gorm:"column:id_kiat;primaryKey"`
	PengetahuanID app.NullInt64 `json:"pengetahuan.id" db:"tpkiat.id_pengetahuan,hide" gorm:"column:id_pengetahuan"`
	Masalah       app.NullText  `json:"masalah"        db:"tpkiat.masalah"             gorm:"column:masalah"`
	Dampak        app.NullText  `json:"dampak"         db:"tpkiat.dampak"              gorm:"column:dampak"`
	Penyebab      app.NullText  `json:"penyebab"       db:"tpkiat.penyebab"            gorm:"column:penyebab"`
	Solusi        app.NullText  `json:"solusi"         db:"tpkiat.solusi"              gorm:"column:solusi"`
	SyaratHasil   app.NullText  `json:"syarat_hasil"   db:"tpkiat.syarat_hasil"        gorm:"column:syarat_hasil"`
}

func (TPengetahuanKiat) EndPoint() string {
	return "t_pengetahuan_kiat"
}

func (TPengetahuanKiat) TableVersion() string {
	return "28.07.291152"
}

func (TPengetahuanKiat) TableName() string {
	return "t_pengetahuan_kiat"
}

func (TPengetahuanKiat) TableAliasName() string {
	return "tpkiat"
}

func (m *TPengetahuanKiat) GetRelations() map[string]map[string]any {

	return m.Relations
}

func (m *TPengetahuanKiat) GetFilters() []map[string]any {
	return m.Filters
}

func (m *TPengetahuanKiat) GetSorts() []map[string]any {
	return m.Sorts
}

func (m *TPengetahuanKiat) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *TPengetahuanKiat) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (TPengetahuanKiat) OpenAPISchemaName() string {
	return "TPengetahuanKiat"
}

func (m *TPengetahuanKiat) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type TPengetahuanKiatList struct {
	app.ListModel
}

func (TPengetahuanKiatList) OpenAPISchemaName() string {
	return "TPengetahuanKiatList"
}

func (p *TPengetahuanKiatList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&TPengetahuanKiat{})
}

// pengetahuan Kompetensi
type TPengetahuanKompetensi struct {
	app.Model
	ID             app.NullInt64 `json:"-"              db:"tpkompetensi.id_pengetahuan_kompetensi" gorm:"column:id_pengetahuan_kompetensi;primaryKey;auto_increment;autoIncrement;primary_key"`
	PengetahuanID  app.NullInt64 `json:"pengetahuan.id" db:"tpkompetensi.id_pengetahuan,hide"       gorm:"column:id_pengetahuan"`
	KompetensiID   app.NullInt64 `json:"id"             db:"tpkompetensi.id_kompetensi"             gorm:"column:id_kompetensi"`
	KompetensiNama app.NullText  `json:"nama"           db:"kom.nama_kompetensi"                    gorm:"-"`
}

func (TPengetahuanKompetensi) EndPoint() string {
	return "t_pengetahuan_kompetensi"
}

func (TPengetahuanKompetensi) TableVersion() string {
	return "28.06.291152"
}

func (TPengetahuanKompetensi) TableName() string {
	return "t_pengetahuan_kompetensi"
}

func (TPengetahuanKompetensi) TableAliasName() string {
	return "tpkompetensi"
}

func (m *TPengetahuanKompetensi) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_kompetensi", "kom", []map[string]any{{"column1": "kom.id_kompetensi", "column2": "tpkompetensi.id_kompetensi"}})

	return m.Relations
}

func (m *TPengetahuanKompetensi) GetFilters() []map[string]any {
	return m.Filters
}

func (m *TPengetahuanKompetensi) GetSorts() []map[string]any {
	return m.Sorts
}

func (m *TPengetahuanKompetensi) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *TPengetahuanKompetensi) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (TPengetahuanKompetensi) OpenAPISchemaName() string {
	return "TPengetahuanKompetensi"
}

func (m *TPengetahuanKompetensi) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type TPengetahuanKompetensiList struct {
	app.ListModel
}

func (TPengetahuanKompetensiList) OpenAPISchemaName() string {
	return "TPengetahuanKompetensiList"
}

func (p *TPengetahuanKompetensiList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&TPengetahuanKompetensi{})
}

// penulis eksternal
type TPengetahuanPenulisExternal struct {
	app.Model
	ID                  app.NullInt64  `json:"-"                     db:"tppenulisexternal.id_penulis_external"   gorm:"column:id_penulis_external;primaryKey"`
	PengetahuanID       app.NullInt64  `json:"pengetahuan.id"        db:"tppenulisexternal.id_pengetahuan,hide"   gorm:"column:id_pengetahuan"`
	NamaPenulisExternal app.NullString `json:"nama_penulis_external" db:"tppenulisexternal.nama_penulis_external" gorm:"column:nama_penulis_external"`
	Nik                 app.NullString `json:"nik"                   db:"tppenulisexternal.nik"                   gorm:"column:nik"`
	AsalInstansi        app.NullString `json:"asal_instansi"         db:"tppenulisexternal.asal_instansi"         gorm:"column:asal_instansi"`
}

func (TPengetahuanPenulisExternal) EndPoint() string {
	return "t_pengetahuan_penulis_external"
}

func (TPengetahuanPenulisExternal) TableVersion() string {
	return "28.06.291152"
}

func (TPengetahuanPenulisExternal) TableName() string {
	return "t_pengetahuan_penulis_external"
}

func (TPengetahuanPenulisExternal) TableAliasName() string {
	return "tppenulisexternal"
}

func (m *TPengetahuanPenulisExternal) GetRelations() map[string]map[string]any {

	return m.Relations
}

func (m *TPengetahuanPenulisExternal) GetFilters() []map[string]any {
	return m.Filters
}

func (m *TPengetahuanPenulisExternal) GetSorts() []map[string]any {
	return m.Sorts
}

func (m *TPengetahuanPenulisExternal) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *TPengetahuanPenulisExternal) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (TPengetahuanPenulisExternal) OpenAPISchemaName() string {
	return "TPengetahuanPenulisExternal"
}

func (m *TPengetahuanPenulisExternal) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type TPengetahuanPenulisExternalList struct {
	app.ListModel
}

func (TPengetahuanPenulisExternalList) OpenAPISchemaName() string {
	return "TPengetahuanPenulisExternalList"
}

func (p *TPengetahuanPenulisExternalList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&TPengetahuanPenulisExternal{})
}

// pengetahuan tag
type TPengetahuanTag struct {
	app.Model
	ID            app.NullInt64 `json:"-"              db:"tptag.id_pengetahuan_tag"  gorm:"column:id_pengetahuan_tag;primaryKey"`
	PengetahuanID app.NullInt64 `json:"pengetahuan.id" db:"tptag.id_pengetahuan,hide" gorm:"column:id_pengetahuan"`
	TagID         app.NullInt64 `json:"id"             db:"tptag.id_tag"              gorm:"column:id_tag"`
	TagNama       app.NullText  `json:"nama"           db:"mtg.nama_tag"              gorm:"-"`
}

func (TPengetahuanTag) EndPoint() string {
	return "t_pengetahuan_pengetahuan_tag"
}

func (TPengetahuanTag) TableVersion() string {
	return "28.06.291152"
}

func (TPengetahuanTag) TableName() string {
	return "t_pengetahuan_pengetahuan_tag"
}

func (TPengetahuanTag) TableAliasName() string {
	return "tptag"
}

func (m *TPengetahuanTag) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_tag", "mtg", []map[string]any{{"column1": "mtg.id_tag", "column2": "tptag.id_tag"}})

	return m.Relations
}

func (m *TPengetahuanTag) GetFilters() []map[string]any {
	return m.Filters
}

func (m *TPengetahuanTag) GetSorts() []map[string]any {
	return m.Sorts
}

func (m *TPengetahuanTag) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *TPengetahuanTag) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (TPengetahuanTag) OpenAPISchemaName() string {
	return "TPengetahuanTag"
}

func (m *TPengetahuanTag) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type TPengetahuanTagList struct {
	app.ListModel
}

func (TPengetahuanTagList) OpenAPISchemaName() string {
	return "TPengetahuanTagList"
}

func (p *TPengetahuanTagList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&TPengetahuanTag{})
}

// pengetahuan tugas
type TPengetahuanTugas struct {
	app.Model
	ID                    app.NullInt64 `json:"-"                       db:"tptugas.id_pengetahuan_tugas"    gorm:"column:id_pengetahuan_tugas;primaryKey"`
	PengetahuanID         app.NullInt64 `json:"pengetahuan.id"          db:"tptugas.id_pengetahuan,hide"     gorm:"column:id_pengetahuan"`
	Tujuan                app.NullText  `json:"tujuan"                  db:"tptugas.tujuan"                  gorm:"column:tujuan"`
	DasarHukum            app.NullText  `json:"dasar_hukum"             db:"tptugas.dasar_hukum"             gorm:"column:dasar_hukum"`
	ProsesBisnis          app.NullText  `json:"proses_bisnis"           db:"tptugas.proses_bisnis"           gorm:"column:proses_bisnis"`
	RumusanMasalah        app.NullText  `json:"rumusan_masalah"         db:"tptugas.rumusan_masalah"         gorm:"column:rumusan_masalah"`
	RisikoObjetPengawasan app.NullText  `json:"risiko_objek_pengawasan" db:"tptugas.risiko_objek_pengawasan" gorm:"column:risiko_objek_pengawasan"`
	MetodePengawasan      app.NullText  `json:"metode_pengawasan"       db:"tptugas.metode_pengawasan"       gorm:"column:metode_pengawasan"`
	TemuanMaterial        app.NullText  `json:"temuan_material"         db:"tptugas.temuan_material"         gorm:"column:temuan_material"`
	KeahlianDibutuhkan    app.NullText  `json:"keahlian_dibutuhkan"     db:"tptugas.keahlian_dibutuhkan"     gorm:"column:keahlian_dibutuhkan"`
	DataDigunakan         app.NullText  `json:"data_digunakan"          db:"tptugas.data_digunakan"          gorm:"column:data_digunakan"`
}

func (TPengetahuanTugas) EndPoint() string {
	return "t_pengetahuan_tugas"
}

func (TPengetahuanTugas) TableVersion() string {
	return "28.07.301152"
}

func (TPengetahuanTugas) TableName() string {
	return "t_pengetahuan_tugas"
}

func (TPengetahuanTugas) TableAliasName() string {
	return "tptugas"
}

func (m *TPengetahuanTugas) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_orang", "tno", []map[string]any{{"column1": "tno.id_orang", "column2": "tptugas.tenaga_ahli_id"}})
	return m.Relations
}

func (m *TPengetahuanTugas) GetFilters() []map[string]any {
	return m.Filters
}

func (m *TPengetahuanTugas) GetSorts() []map[string]any {
	return m.Sorts
}

func (m *TPengetahuanTugas) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *TPengetahuanTugas) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (TPengetahuanTugas) OpenAPISchemaName() string {
	return "TPengetahuanTugas"
}

func (m *TPengetahuanTugas) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type TPengetahuanTugasList struct {
	app.ListModel
}

func (TPengetahuanTugasList) OpenAPISchemaName() string {
	return "TPengetahuanTugasList"
}

func (p *TPengetahuanTugasList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&TPengetahuanTugas{})
}

// pengetahuan referensi
type TPengetahuanReferensi struct {
	app.Model
	ID            app.NullInt64 `json:"-"              db:"tpref.id_pengetahuan_referensi" gorm:"column:id_pengetahuan_referensi;primaryKey;auto_increment"`
	PengetahuanID app.NullInt64 `json:"pengetahuan.id" db:"tpref.id_pengetahuan,hide"      gorm:"column:id_pengetahuan"`
	ReferensiID   app.NullInt64 `json:"id"             db:"tpref.id_referensi"             gorm:"column:id_referensi"`
	ReferensiNama app.NullText  `json:"nama"           db:"mref.nama_referensi"            gorm:"-"`
}

func (TPengetahuanReferensi) EndPoint() string {
	return "t_pengetahuan_referensi"
}

func (TPengetahuanReferensi) TableVersion() string {
	return "28.07.291152"
}

func (TPengetahuanReferensi) TableName() string {
	return "t_pengetahuan_referensi"
}

func (TPengetahuanReferensi) TableAliasName() string {
	return "tpref"
}

func (m *TPengetahuanReferensi) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_referensi", "mref", []map[string]any{{"column1": "mref.id_referensi", "column2": "tpref.id_referensi"}})

	return m.Relations
}

func (m *TPengetahuanReferensi) GetFilters() []map[string]any {
	return m.Filters
}

func (m *TPengetahuanReferensi) GetSorts() []map[string]any {
	return m.Sorts
}

func (m *TPengetahuanReferensi) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *TPengetahuanReferensi) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (TPengetahuanReferensi) OpenAPISchemaName() string {
	return "TPengetahuanReferensi"
}

func (m *TPengetahuanReferensi) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type TPengetahuanReferensiList struct {
	app.ListModel
}

func (TPengetahuanReferensiList) OpenAPISchemaName() string {
	return "TPengetahuanReferensiList"
}

func (p *TPengetahuanReferensiList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&TPengetahuanReferensi{})
}

// tpengetahuan dokumen
type TPengetahuanDokumen struct {
	app.Model
	ID             app.NullInt64 `json:"-"              db:"tpdokumen.id_pengetahuan_referensi" gorm:"column:id_pengetahuan_referensi;primaryKey"`
	PengetahuanID  app.NullInt64 `json:"pengetahuan.id" db:"tpdokumen.id_pengetahuan,hide"      gorm:"column:id_pengetahuan"`
	AttachmentID   app.NullInt64 `json:"id"             db:"tpdokumen.id_attachment"            gorm:"column:id_attachment"`
	AttachmentNama app.NullText  `json:"nama"           db:"attachment.filename"                gorm:"-"`
	AttachmentUrl  app.NullText  `json:"url"            db:"attachment.url"                     gorm:"-"`
}

func (TPengetahuanDokumen) EndPoint() string {
	return "t_pengetahuan_dokumen"
}

func (TPengetahuanDokumen) TableVersion() string {
	return "28.06.291152"
}

func (TPengetahuanDokumen) TableName() string {
	return "t_pengetahuan_dokumen"
}

func (TPengetahuanDokumen) TableAliasName() string {
	return "tpdokumen"
}

func (m *TPengetahuanDokumen) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_attachments", "attachment", []map[string]any{{"column1": "attachment.id", "column2": "tpdokumen.id_attachment"}})

	return m.Relations
}

func (m *TPengetahuanDokumen) GetFilters() []map[string]any {
	return m.Filters
}

func (m *TPengetahuanDokumen) GetSorts() []map[string]any {
	return m.Sorts
}

func (m *TPengetahuanDokumen) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *TPengetahuanDokumen) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (TPengetahuanDokumen) OpenAPISchemaName() string {
	return "TPengetahuanDokumen"
}

func (m *TPengetahuanDokumen) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type TPengetahuanDokumenList struct {
	app.ListModel
}

func (TPengetahuanDokumenList) OpenAPISchemaName() string {
	return "TPengetahuanDokumenList"
}

func (p *TPengetahuanDokumenList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&TPengetahuanDokumen{})
}

// pengetahuan tenaga ahli
type TPengetahuanTenagaAhli struct {
	app.Model
	ID                app.NullInt64  `json:"-"              db:"tptenagaahli.id_pengetahuan_ta"   gorm:"column:id_pengetahuan_ta;primaryKey;auto_increment;AutoIncrement;"`
	PengetahuanID     app.NullInt64  `json:"pengetahuan.id" db:"tptenagaahli.id_pengetahuan,hide" gorm:"column:id_pengetahuan"`
	TenagaAhliID      app.NullInt64  `json:"id"             db:"tptenagaahli.tenaga_ahli_id"      gorm:"column:tenaga_ahli_id"`
	TenagaAhliNip     app.NullString `json:"nip"            db:"tno.nip"                          gorm:"-"`
	TenagaAhliNama    app.NullString `json:"nama_lengkap"   db:"tno.nama"                         gorm:"-"`
	TenagaAhliJabatan app.NullString `json:"jabatan"        db:"tno.jabatan"                      gorm:"-"`
	TenagaAhliEmail   app.NullString `json:"email"          db:"tno.email"                        gorm:"-"`
}

func (TPengetahuanTenagaAhli) EndPoint() string {
	return "t_pengetahuan_tenaga_ahli"
}

func (TPengetahuanTenagaAhli) TableVersion() string {
	return "28.07.301152"
}

func (TPengetahuanTenagaAhli) TableName() string {
	return "t_pengetahuan_tenaga_ahli"
}

func (TPengetahuanTenagaAhli) TableAliasName() string {
	return "tptenagaahli"
}

func (m *TPengetahuanTenagaAhli) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_orang", "tno", []map[string]any{{"column1": "tno.id_orang", "column2": "tptenagaahli.tenaga_ahli_id"}})
	return m.Relations
}

func (m *TPengetahuanTenagaAhli) GetFilters() []map[string]any {
	return m.Filters
}

func (m *TPengetahuanTenagaAhli) GetSorts() []map[string]any {
	return m.Sorts
}

func (m *TPengetahuanTenagaAhli) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *TPengetahuanTenagaAhli) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (TPengetahuanTenagaAhli) OpenAPISchemaName() string {
	return "TPengetahuanTenagaAhli"
}

func (m *TPengetahuanTenagaAhli) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type TPengetahuanTenagaAhliList struct {
	app.ListModel
}

func (TPengetahuanTenagaAhliList) OpenAPISchemaName() string {
	return "TPengetahuanTenagaAhliList"
}

func (p *TPengetahuanTenagaAhliList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&TPengetahuanTenagaAhli{})
}

// pengetahuan pedoman
type TPengetahuanPedoman struct {
	app.Model
	ID            app.NullInt64  `json:"-"              db:"tp_pedoman.id_pengetahuan_ta"   gorm:"column:id_pengetahuan_ta;primaryKey;auto_increment;AutoIncrement;"`
	PengetahuanID app.NullInt64  `json:"pengetahuan.id" db:"tp_pedoman.id_pengetahuan,hide" gorm:"column:id_pengetahuan"`
	PedomanID     app.NullText   `json:"id"             db:"tp_pedoman.pedoman_id"          gorm:"column:pedoman_id"`
	PedmoanNama   app.NullString `json:"nama"           db:"-"                              gorm:"-"`
	PedmoanData   app.NullString `json:"data"           db:"-"                              gorm:"-"`
}

func (TPengetahuanPedoman) EndPoint() string {
	return "t_pengetahuan_pedoman"
}

func (TPengetahuanPedoman) TableVersion() string {
	return "28.07.301152"
}

func (TPengetahuanPedoman) TableName() string {
	return "t_pengetahuan_pedoman"
}

func (TPengetahuanPedoman) TableAliasName() string {
	return "tp_pedoman"
}

func (m *TPengetahuanPedoman) GetRelations() map[string]map[string]any {
	return m.Relations
}

func (m *TPengetahuanPedoman) GetFilters() []map[string]any {
	return m.Filters
}

func (m *TPengetahuanPedoman) GetSorts() []map[string]any {
	return m.Sorts
}

func (m *TPengetahuanPedoman) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *TPengetahuanPedoman) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (TPengetahuanPedoman) OpenAPISchemaName() string {
	return "TPengetahuanPedoman"
}

func (m *TPengetahuanPedoman) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type TPengetahuanPedomanList struct {
	app.ListModel
}

func (TPengetahuanPedomanList) OpenAPISchemaName() string {
	return "TPengetahuanPedomanList"
}

func (p *TPengetahuanPedomanList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&TPengetahuanPedoman{})
}

// pengetahuan Narasumber
type TPengetahuanNarsum struct {
	app.Model
	ID             app.NullInt64  `json:"-"              db:"tp_narsum.id_pengetahuan_narsum" gorm:"column:id_pengetahuan_narsum;primaryKey;auto_increment;AutoIncrement;"`
	PengetahuanID  app.NullInt64  `json:"pengetahuan.id" db:"tp_narsum.id_pengetahuan,hide"   gorm:"column:id_pengetahuan"`
	NarasumberID   app.NullText   `json:"id"             db:"tp_narsum.id_narasumber"         gorm:"column:id_narasumber"`
	NarasumberNama app.NullString `json:"nama"           db:"narsum.nama"                     gorm:"-"`
}

func (TPengetahuanNarsum) EndPoint() string {
	return "t_pengetahuan_narsum"
}

func (TPengetahuanNarsum) TableVersion() string {
	return "28.07.301152"
}

func (TPengetahuanNarsum) TableName() string {
	return "t_pengetahuan_narsum"
}

func (TPengetahuanNarsum) TableAliasName() string {
	return "tp_narsum"
}

func (m *TPengetahuanNarsum) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_narasumber", "narsum", []map[string]any{{"column1": "narsum.id", "column2": "tp_narsum.id_narasumber"}})
	return m.Relations
}

func (m *TPengetahuanNarsum) GetFilters() []map[string]any {
	return m.Filters
}

func (m *TPengetahuanNarsum) GetSorts() []map[string]any {
	return m.Sorts
}

func (m *TPengetahuanNarsum) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *TPengetahuanNarsum) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (TPengetahuanNarsum) OpenAPISchemaName() string {
	return "TPengetahuanNarsum"
}

func (m *TPengetahuanNarsum) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type TPengetahuanNarsumList struct {
	app.ListModel
}

func (TPengetahuanNarsumList) OpenAPISchemaName() string {
	return "TPengetahuanNarsumList"
}

func (p *TPengetahuanNarsumList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&TPengetahuanNarsum{})
}

// pengetahuan Penerbit
type TPengetahuanPenerbit struct {
	app.Model
	ID            app.NullInt64  `json:"-"              db:"tp_penerbit.id_pengetahuan_penerbit" gorm:"column:id_pengetahuan_penerbit;primaryKey;auto_increment;AutoIncrement;"`
	PengetahuanID app.NullInt64  `json:"pengetahuan.id" db:"tp_penerbit.id_pengetahuan,hide"     gorm:"column:id_pengetahuan"`
	PenerbitID    app.NullText   `json:"id"             db:"tp_penerbit.id_penerbit"             gorm:"column:id_penerbit"`
	PenerbitNama  app.NullString `json:"nama"           db:"penerbit.nama"                       gorm:"-"`
}

func (TPengetahuanPenerbit) EndPoint() string {
	return "t_pengetahuan_penerbit"
}

func (TPengetahuanPenerbit) TableVersion() string {
	return "28.07.301152"
}

func (TPengetahuanPenerbit) TableName() string {
	return "t_pengetahuan_penerbit"
}

func (TPengetahuanPenerbit) TableAliasName() string {
	return "tp_penerbit"
}

func (m *TPengetahuanPenerbit) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_penerbit", "penerbit", []map[string]any{{"column1": "penerbit.id", "column2": "tp_penerbit.id_penerbit"}})
	return m.Relations
}

func (m *TPengetahuanPenerbit) GetFilters() []map[string]any {
	return m.Filters
}

func (m *TPengetahuanPenerbit) GetSorts() []map[string]any {
	return m.Sorts
}

func (m *TPengetahuanPenerbit) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *TPengetahuanPenerbit) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (TPengetahuanPenerbit) OpenAPISchemaName() string {
	return "TPengetahuanPenerbit"
}

func (m *TPengetahuanPenerbit) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type TPengetahuanPenerbitList struct {
	app.ListModel
}

func (TPengetahuanPenerbitList) OpenAPISchemaName() string {
	return "TPengetahuanPenerbitList"
}

func (p *TPengetahuanPenerbitList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&TPengetahuanPenerbit{})
}

// pengetahuan resensi
type TPengetahuanResensi struct {
	app.Model
	ID                  app.NullInt64          `json:"-"                    db:"tp_resensi.id_pengetahuan_narsum" gorm:"column:id_pengetahuan_narsum;primaryKey;auto_increment;AutoIncrement;"`
	PengetahuanID       app.NullInt64          `json:"pengetahuan.id"       db:"tp_resensi.id_pengetahuan,hide"   gorm:"column:id_pengetahuan"`
	Narasumber          []TPengetahuanNarsum   `json:"narasumber"           db:"pengetahuan.id={pengetahuan.id}"  gorm:"-"`
	JumlahHalaman       app.NullInt64          `json:"jumlah_halaman"       db:"tp_resensi.jumlah_halaman"        gorm:"column:jumlah_halaman"`
	Penerbit            []TPengetahuanPenerbit `json:"penerbit"             db:"pengetahuan.id={pengetahuan.id}"  gorm:"-"`
	TahunTerbit         app.NullInt64          `json:"tahun_terbit"         db:"tp_resensi.tahun_terbit"          gorm:"column:tahun_terbit"`
	LatarBelakang       app.NullText           `json:"latar_belakang"       db:"tp_resensi.latar_belakang"        gorm:"column:latar_belakang"`
	PenelitianTerdahulu app.NullText           `json:"penelitian_terdahulu" db:"tp_resensi.penelitian_terdahulu"  gorm:"column:penelitian_terdahulu"`
	LessonLearned       app.NullText           `json:"lesson_learned"       db:"tp_resensi.lesson_learned"        gorm:"column:lesson_learned"`
}

func (TPengetahuanResensi) EndPoint() string {
	return "t_pengetahuan_resensi"
}

func (TPengetahuanResensi) TableVersion() string {
	return "28.07.301152"
}

func (TPengetahuanResensi) TableName() string {
	return "t_pengetahuan_resensi"
}

func (TPengetahuanResensi) TableAliasName() string {
	return "tp_resensi"
}

func (m *TPengetahuanResensi) GetRelations() map[string]map[string]any {
	return m.Relations
}

func (m *TPengetahuanResensi) GetFilters() []map[string]any {
	return m.Filters
}

func (m *TPengetahuanResensi) GetSorts() []map[string]any {
	return m.Sorts
}

func (m *TPengetahuanResensi) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *TPengetahuanResensi) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (TPengetahuanResensi) OpenAPISchemaName() string {
	return "TPengetahuanResensi"
}

func (m *TPengetahuanResensi) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type TPengetahuanResensiList struct {
	app.ListModel
}

func (TPengetahuanResensiList) OpenAPISchemaName() string {
	return "TPengetahuanResensiList"
}

func (p *TPengetahuanResensiList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&TPengetahuanResensi{})
}
