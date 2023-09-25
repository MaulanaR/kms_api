package dokumen

import "github.com/maulanar/kms/app"

type Dokumen struct {
	app.Model
	ID                    app.NullInt64    `json:"id"                      db:"m.id"                      gorm:"column:id;primaryKey"`
	NamaPengarang         app.NullText     `json:"nama_pengarang"          db:"m.nama_pengarang"          gorm:"column:nama_pengarang"`
	NomorDokumen          app.NullText     `json:"nomor_dokumen"           db:"m.nomor_dokumen"           gorm:"column:nomor_dokumen"`
	TanggalDokumen        app.NullDate     `json:"tanggal_dokumen"         db:"m.tanggal_dokumen"         gorm:"column:tanggal_dokumen"`
	JudulDokumen          app.NullText     `json:"judul_dokumen"           db:"m.judul_dokumen"           gorm:"column:judul_dokumen"`
	Penerbit              app.NullText     `json:"penerbit"                db:"m.penerbit"                gorm:"column:penerbit"`
	KelompokDokumenId     app.NullInt64    `json:"kelompok_dokumen_id"     db:"m.kelompok_dokumen_id"     gorm:"column:kelompok_dokumen_id"`
	KategoriPengetahuanId app.NullInt64    `json:"kategori_pengetahuan_id" db:"m.kategori_pengetahuan_id" gorm:"column:kategori_pengetahuan_id"`
	IsiDokumen            app.NullText     `json:"isi_dokumen"             db:"m.isi_dokumen"             gorm:"column:isi_dokumen"`
	AttachmentId          app.NullInt64    `json:"attachment_id"           db:"m.attachment_id"           gorm:"column:attachment_id"`
	CreatedBy             app.NullInt64    `json:"created_by.id"           db:"m.created_by"              gorm:"column:created_by"`
	CreatedByUsername     app.NullString   `json:"created_by.username"     db:"cbuser.username"           gorm:"-"`
	UpdatedBy             app.NullInt64    `json:"updated_by.id"           db:"m.updated_by"              gorm:"column:updated_by"`
	UpdatedByUsername     app.NullString   `json:"updated_by.username"     db:"ubuser.username"           gorm:"-"`
	DeletedBy             app.NullInt64    `json:"deleted_by.id"           db:"m.deleted_by"              gorm:"column:deleted_by"`
	DeletedByUsername     app.NullString   `json:"deleted_by.username"     db:"dbuser.username"           gorm:"-"`
	CreatedAt             app.NullDateTime `json:"created_at"              db:"m.created_at"              gorm:"column:created_at"`
	UpdatedAt             app.NullDateTime `json:"updated_at"              db:"m.updated_at"              gorm:"column:updated_at"`
	DeletedAt             app.NullDateTime `json:"deleted_at"              db:"m.deleted_at,hide"         gorm:"column:deleted_at"`
}

func (Dokumen) EndPoint() string {
	return "dokumen"
}

func (Dokumen) TableVersion() string {
	return "28.06.291152"
}

func (Dokumen) TableName() string {
	return "dokumen"
}

func (Dokumen) TableAliasName() string {
	return "m"
}

func (m *Dokumen) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_user", "cbuser", []map[string]any{{"column1": "cbuser.id_user", "column2": "m.created_by"}})
	m.AddRelation("left", "m_user", "ubuser", []map[string]any{{"column1": "ubuser.id_user", "column2": "m.updated_by"}})
	m.AddRelation("left", "m_user", "dbuser", []map[string]any{{"column1": "dbuser.id_user", "column2": "m.deleted_by"}})

	return m.Relations
}

func (m *Dokumen) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *Dokumen) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *Dokumen) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *Dokumen) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (Dokumen) OpenAPISchemaName() string {
	return "Dokumen"
}

func (m *Dokumen) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type DokumenList struct {
	app.ListModel
}

func (DokumenList) OpenAPISchemaName() string {
	return "DokumenList"
}

func (p *DokumenList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Dokumen{})
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
