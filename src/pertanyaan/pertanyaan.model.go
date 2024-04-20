package pertanyaan

import "github.com/maulanar/kms/app"

type Pertanyaan struct {
	app.Model
	ID                 app.NullInt64    `json:"id"                   db:"m.id"              gorm:"column:id;primaryKey"`
	Judul              app.NullText     `json:"judul"                db:"m.judul"           gorm:"column:judul"`
	Masalah            app.NullText     `json:"masalah"              db:"m.masalah"         gorm:"column:masalah"`
	Ekspektasi         app.NullText     `json:"ekspektasi"           db:"m.ekspektasi"      gorm:"column:ekspektasi"`
	PakarID            app.NullInt64    `json:"pakar.id"             db:"m.pakar_id"        gorm:"column:pakar_id"`
	OrangNama          app.NullString   `json:"pakar.nama_lengkap"   db:"o.nama"            gorm:"-"`
	OrangNamaPanggilan app.NullString   `json:"pakar.nama_panggilan" db:"o.nama_panggilan"  gorm:"-"`
	OrangJenisKelamin  app.NullString   `json:"pakar.jenis_kelamin"  db:"o.jenis_kelamin"   gorm:"-"`
	OrangEmail         app.NullString   `json:"pakar.email"          db:"o.email"           gorm:"-"`
	OrangFotoID        app.NullInt64    `json:"pakar.foto.id"        db:"o.foto"            gorm:"-"`
	OrangFotoUrl       app.NullString   `json:"pakar.foto.url"       db:"att.url"           gorm:"-"`
	OrangFotoNama      app.NullString   `json:"pakar.foto.nama"      db:"att.filename"      gorm:"-"`
	Kategori           app.NullString   `json:"pakar.kategori"       db:"u.kategori"        gorm:"-"`
	Level              app.NullString   `json:"pakar.level"          db:"u.level"           gorm:"-"`
	Jawaban            []Jawaban        `json:"jawaban"              db:"id={id}"           gorm:"-"`
	CreatedBy          app.NullInt64    `json:"created_by.id"        db:"m.created_by"      gorm:"column:created_by"`
	CreatedByUsername  app.NullString   `json:"created_by.username"  db:"cbuser.username"   gorm:"-"`
	UpdatedBy          app.NullInt64    `json:"updated_by.id"        db:"m.updated_by"      gorm:"column:updated_by"`
	UpdatedByUsername  app.NullString   `json:"updated_by.username"  db:"ubuser.username"   gorm:"-"`
	DeletedBy          app.NullInt64    `json:"deleted_by.id"        db:"m.deleted_by"      gorm:"column:deleted_by"`
	DeletedByUsername  app.NullString   `json:"deleted_by.username"  db:"dbuser.username"   gorm:"-"`
	CreatedAt          app.NullDateTime `json:"created_at"           db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt          app.NullDateTime `json:"updated_at"           db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt          app.NullDateTime `json:"deleted_at"           db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

func (Pertanyaan) EndPoint() string {
	return "pertanyaan"
}

func (Pertanyaan) TableVersion() string {
	return "24.04.201130"
}

func (Pertanyaan) TableName() string {
	return "m_pertanyaan"
}

func (Pertanyaan) TableAliasName() string {
	return "m"
}

func (m *Pertanyaan) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_user", "cbuser", []map[string]any{{"column1": "cbuser.id_user", "column2": "m.created_by"}})
	m.AddRelation("left", "m_user", "ubuser", []map[string]any{{"column1": "ubuser.id_user", "column2": "m.updated_by"}})
	m.AddRelation("left", "m_user", "dbuser", []map[string]any{{"column1": "dbuser.id_user", "column2": "m.deleted_by"}})

	m.AddRelation("left", "m_user", "u", []map[string]any{{"column1": "u.id_user", "column2": "m.pakar_id"}})
	m.AddRelation("left", "m_orang", "o", []map[string]any{{"column1": "o.id_orang", "column2": "u.id_orang"}})
	m.AddRelation("left", "m_attachments", "att", []map[string]any{{"column1": "att.id", "column2": "o.foto"}})

	return m.Relations
}

func (m *Pertanyaan) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *Pertanyaan) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *Pertanyaan) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *Pertanyaan) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (Pertanyaan) OpenAPISchemaName() string {
	return "Pertanyaan"
}

func (m *Pertanyaan) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type PertanyaanList struct {
	app.ListModel
}

func (PertanyaanList) OpenAPISchemaName() string {
	return "PertanyaanList"
}

func (p *PertanyaanList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Pertanyaan{})
}

type ParamCreate struct {
	Judul   app.NullText `json:"judul"   db:"m.judul"   gorm:"column:judul"   validate:"required"`
	Masalah app.NullText `json:"masalah" db:"m.masalah" gorm:"column:masalah" validate:"required"`
	UseCaseHandler
}

type ParamUpdate struct {
	Judul   app.NullText `json:"judul"   db:"m.judul"   gorm:"column:judul"   validate:"required"`
	Masalah app.NullText `json:"masalah" db:"m.masalah" gorm:"column:masalah" validate:"required"`
	UseCaseHandler
}

type ParamPartiallyUpdate struct {
	UseCaseHandler
}

type ParamDelete struct {
	UseCaseHandler
}

type PostJawaban struct {
	Ctx        *app.Ctx     `json:"-"          db:"-"            gorm:"-"`
	Keterangan app.NullText `json:"keterangan" db:"j.keterangan" gorm:"column:keterangan" validate:"required"`
	Jawaban
}

type ParamJawabanDelete struct {
	Ctx *app.Ctx `json:"-" db:"-" gorm:"-"`
	Jawaban
}

// Jawaban
type Jawaban struct {
	app.Model
	ID                app.NullInt64    `json:"id"                  db:"j.id"              gorm:"column:id;primaryKey"`
	IDPertanyaan      app.NullInt64    `json:"id_pertanyaan"       db:"j.id_pertanyaan"   gorm:"column:id_pertanyaan"`
	Keterangan        app.NullText     `json:"keterangan"          db:"j.keterangan"      gorm:"column:keterangan"`
	CreatedBy         app.NullInt64    `json:"created_by.id"       db:"j.created_by"      gorm:"column:created_by"`
	CreatedByUsername app.NullString   `json:"created_by.username" db:"cbuser.username"   gorm:"-"`
	UpdatedBy         app.NullInt64    `json:"updated_by.id"       db:"j.updated_by"      gorm:"column:updated_by"`
	UpdatedByUsername app.NullString   `json:"updated_by.username" db:"ubuser.username"   gorm:"-"`
	DeletedBy         app.NullInt64    `json:"deleted_by.id"       db:"j.deleted_by"      gorm:"column:deleted_by"`
	DeletedByUsername app.NullString   `json:"deleted_by.username" db:"dbuser.username"   gorm:"-"`
	CreatedAt         app.NullDateTime `json:"created_at"          db:"j.created_at"      gorm:"column:created_at"`
	UpdatedAt         app.NullDateTime `json:"updated_at"          db:"j.updated_at"      gorm:"column:updated_at"`
	DeletedAt         app.NullDateTime `json:"deleted_at"          db:"j.deleted_at,hide" gorm:"column:deleted_at"`
}

func (Jawaban) EndPoint() string {
	return "jawaban"
}

func (Jawaban) TableVersion() string {
	return "24.04.131130"
}

func (Jawaban) TableName() string {
	return "m_jawaban"
}

func (Jawaban) TableAliasName() string {
	return "j"
}

func (m *Jawaban) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_user", "cbuser", []map[string]any{{"column1": "cbuser.id_user", "column2": "j.created_by"}})
	m.AddRelation("left", "m_user", "ubuser", []map[string]any{{"column1": "ubuser.id_user", "column2": "j.updated_by"}})
	m.AddRelation("left", "m_user", "dbuser", []map[string]any{{"column1": "dbuser.id_user", "column2": "j.deleted_by"}})

	return m.Relations
}

func (m *Jawaban) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "j.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *Jawaban) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "j.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *Jawaban) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *Jawaban) GetSchema() map[string]any {
	return m.SetSchema(m)
}
