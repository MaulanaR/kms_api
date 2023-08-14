package dislike

import "github.com/maulanar/kms/app"

type Dislike struct {
	app.Model
	ID                     app.NullInt64    `json:"id"                  db:"m.id"              gorm:"column:id;primaryKey"`
	PengetahuanID          app.NullInt64    `json:"pengetahuan.id"      db:"m.id_pengetahuan"  gorm:"column:id_pengetahuan"`
	UserID                 app.NullInt64    `json:"user.id"             db:"m.id_user"         gorm:"column:id_user"`
	UserOrangId            app.NullInt64    `json:"user.orang.id"       db:"u.id_orang,hide"   gorm:"-"`
	UserOrangNama          app.NullString   `json:"user.nama_lengkap"   db:"o.nama"            gorm:"-"`
	UserOrangNamaPanggilan app.NullString   `json:"user.nama_panggilan" db:"o.nama_panggilan"  gorm:"-"`
	UserOrangJabatan       app.NullString   `json:"user.jabatan"        db:"o.jabatan"         gorm:"-"`
	UserOrangEmail         app.NullString   `json:"user.email"          db:"o.email"           gorm:"-"`
	UserOrangFotoID        app.NullInt64    `json:"user.foto.id"        db:"o.foto"            gorm:"-"`
	UserOrangFotoUrl       app.NullString   `json:"user.foto.url"       db:"att.url"           gorm:"-"`
	UserOrangFotoNama      app.NullString   `json:"user.foto.nama"      db:"att.filename"      gorm:"-"`
	UserOrangUnitKerja     app.NullString   `json:"user.unit_kerja"     db:"o.unit_kerja"      gorm:"-"`
	UserOrangUserLevel     app.NullString   `json:"user.user_level"     db:"o.user_level"      gorm:"-"`
	UserOrangStatusLevel   app.NullString   `json:"user.status_level"   db:"o.status_level"    gorm:"-"`
	UserOrangNip           app.NullString   `json:"user.nip"            db:"o.nip"             gorm:"-"`
	UserUsername           app.NullString   `json:"user.username"       db:"u.username"        gorm:"-"`
	CreatedAt              app.NullDateTime `json:"created_at"          db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt              app.NullDateTime `json:"updated_at"          db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt              app.NullDateTime `json:"deleted_at"          db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

func (Dislike) EndPoint() string {
	return "dislike"
}

func (Dislike) TableVersion() string {
	return "28.06.291152"
}

func (Dislike) TableName() string {
	return "t_dislike"
}

func (Dislike) TableAliasName() string {
	return "m"
}

func (m *Dislike) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_user", "u", []map[string]any{{"column1": "u.id_user", "column2": "m.id_user"}})
	m.AddRelation("left", "m_orang", "o", []map[string]any{{"column1": "o.id_orang", "column2": "u.id_orang"}})
	m.AddRelation("left", "m_attachments", "att", []map[string]any{{"column1": "att.id", "column2": "u.foto"}})
	return m.Relations
}

func (m *Dislike) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *Dislike) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *Dislike) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *Dislike) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (Dislike) OpenAPISchemaName() string {
	return "Dislike"
}

func (m *Dislike) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type DislikeList struct {
	app.ListModel
}

func (DislikeList) OpenAPISchemaName() string {
	return "DislikeList"
}

func (p *DislikeList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Dislike{})
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
