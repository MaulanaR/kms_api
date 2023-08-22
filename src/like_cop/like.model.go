package like_cop

import "github.com/maulanar/kms/app"

type LikeCOP struct {
	app.Model
	ID                     app.NullInt64    `json:"id"                  db:"m.id"              gorm:"column:id;AutoIncrement"`
	CopID                  app.NullInt64    `json:"cop.id"              db:"m.id_cop"          gorm:"column:id_cop" validate:"required"`
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

func (LikeCOP) EndPoint() string {
	return "like_cop"
}

func (LikeCOP) TableVersion() string {
	return "28.06.291152"
}

func (LikeCOP) TableName() string {
	return "t_like_cop"
}

func (LikeCOP) TableAliasName() string {
	return "m"
}

func (m *LikeCOP) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_user", "u", []map[string]any{{"column1": "u.id_user", "column2": "m.id_user"}})
	m.AddRelation("left", "m_orang", "o", []map[string]any{{"column1": "o.id_orang", "column2": "u.id_orang"}})
	m.AddRelation("left", "m_attachments", "att", []map[string]any{{"column1": "att.id", "column2": "o.foto"}})
	return m.Relations
}

func (m *LikeCOP) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *LikeCOP) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *LikeCOP) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *LikeCOP) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (LikeCOP) OpenAPISchemaName() string {
	return "Like"
}

func (m *LikeCOP) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type LikeList struct {
	app.ListModel
}

func (LikeList) OpenAPISchemaName() string {
	return "LikeList"
}

func (p *LikeList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&LikeCOP{})
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
