package elibrary

import "github.com/maulanar/kms/app"

type Elibrary struct {
	app.Model
	ID                app.NullInt64    `json:"id"                  db:"m.id"               gorm:"column:id;primaryKey"`
	Pengarang         app.NullText     `json:"pengarang"           db:"m.pengarang"        gorm:"column:pengarang"`
	UnitKerja         app.NullText     `json:"unit_kerja"          db:"m.unit_kerja"       gorm:"column:unit_kerja"`
	LokasiGedung      app.NullText     `json:"lokasi_gedung"       db:"m.lokasi_gedung"    gorm:"column:lokasi_gedung"`
	Ruangan           app.NullText     `json:"ruangan"             db:"m.ruangan"          gorm:"column:ruangan"`
	KategoriBukuId    app.NullInt64    `json:"kategori_buku.id"    db:"m.kategori_buku_id" gorm:"column:kategori_buku_id"`
	KategoriBukuNama  app.NullText     `json:"kategori_buku.nama"  db:"kb.nama"            gorm:"-"`
	JudulBuku         app.NullText     `json:"judul_buku"          db:"m.judul_buku"       gorm:"column:judul_buku"`
	IsiBuku           app.NullText     `json:"isi_buku"            db:"m.isi_buku"         gorm:"column:isi_buku"`
	GambarId          app.NullInt64    `json:"gambar.id"           db:"m.gambar_id"        gorm:"column:gambar_id"`
	GambarUrl         app.NullString   `json:"gambar.url"          db:"gbr.url"            gorm:"-"`
	GambarNama        app.NullString   `json:"gambar.nama"         db:"gbr.filename"       gorm:"-"`
	DokumenId         app.NullInt64    `json:"dokumen.id"          db:"m.dokumen_id"       gorm:"column:dokumen_id"`
	DokumenUrl        app.NullString   `json:"dokumen.url"         db:"doku.url"           gorm:"-"`
	DokumenNama       app.NullString   `json:"dokumen.nama"        db:"doku.filename"      gorm:"-"`
	CreatedBy         app.NullInt64    `json:"created_by.id"       db:"m.created_by"       gorm:"column:created_by"`
	CreatedByUsername app.NullString   `json:"created_by.username" db:"cbuser.username"    gorm:"-"`
	UpdatedBy         app.NullInt64    `json:"updated_by.id"       db:"m.updated_by"       gorm:"column:updated_by"`
	UpdatedByUsername app.NullString   `json:"updated_by.username" db:"ubuser.username"    gorm:"-"`
	DeletedBy         app.NullInt64    `json:"deleted_by.id"       db:"m.deleted_by"       gorm:"column:deleted_by"`
	DeletedByUsername app.NullString   `json:"deleted_by.username" db:"dbuser.username"    gorm:"-"`
	CreatedAt         app.NullDateTime `json:"created_at"          db:"m.created_at"       gorm:"column:created_at"`
	UpdatedAt         app.NullDateTime `json:"updated_at"          db:"m.updated_at"       gorm:"column:updated_at"`
	DeletedAt         app.NullDateTime `json:"deleted_at"          db:"m.deleted_at,hide"  gorm:"column:deleted_at"`

	//statistik View, like, dislike, komentar
	StatistikView     app.NullInt64 `json:"statistik.view"      db:"(CASE WHEN m.count_view > 0 THEN m.count_view ELSE 0 END)"               gorm:"column:count_view;default:0"`
	StatistikKomentar app.NullInt64 `json:"statistik.komentar"  db:"(SELECT COUNT(*) FROM t_komentar WHERE t_komentar.id_elibrary=m.id)" gorm:"-"`
}

func (e *Elibrary) KategoriBukuRefer() string {
	return "m_kategori_buku(id)"
}

func (Elibrary) EndPoint() string {
	return "elibrary"
}

func (Elibrary) TableVersion() string {
	return "23.10.061152"
}

func (Elibrary) TableName() string {
	return "t_elibrary"
}

func (Elibrary) TableAliasName() string {
	return "m"
}

func (m *Elibrary) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_kategori_buku", "kb", []map[string]any{{"column1": "kb.id", "column2": "m.kategori_buku_id"}})
	m.AddRelation("left", "m_attachments", "gbr", []map[string]any{{"column1": "gbr.id", "column2": "m.gambar_id"}})
	m.AddRelation("left", "m_attachments", "doku", []map[string]any{{"column1": "doku.id", "column2": "m.dokumen_id"}})
	m.AddRelation("left", "m_user", "cbuser", []map[string]any{{"column1": "cbuser.id_user", "column2": "m.created_by"}})
	m.AddRelation("left", "m_user", "ubuser", []map[string]any{{"column1": "ubuser.id_user", "column2": "m.updated_by"}})
	m.AddRelation("left", "m_user", "dbuser", []map[string]any{{"column1": "dbuser.id_user", "column2": "m.deleted_by"}})

	return m.Relations
}

func (m *Elibrary) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *Elibrary) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *Elibrary) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *Elibrary) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (Elibrary) OpenAPISchemaName() string {
	return "Elibrary"
}

func (m *Elibrary) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type ElibraryList struct {
	app.ListModel
}

func (ElibraryList) OpenAPISchemaName() string {
	return "ElibraryList"
}

func (p *ElibraryList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Elibrary{})
}

type ParamCreate struct {
	UseCaseHandler
}

type ParamUpdate struct {
	UseCaseHandler
}

type ParamPartiallyUpdate struct {
	UseCaseHandler
}

type ParamDelete struct {
	UseCaseHandler
}
