package notifikasi

import "github.com/maulanar/kms/app"

type Notifikasi struct {
	app.Model
	ID app.NullInt64 `json:"id"                  db:"m.id"              gorm:"column:id;primaryKey"`

	Endpoint  app.NullString `json:"endpoint"            db:"m.endpoint"        gorm:"column:endpoint"`
	DataID    app.NullInt64  `json:"data_id"             db:"m.data_id"         gorm:"column:data_id"`
	UserID    app.NullInt64  `json:"user_id"             db:"m.user_id"         gorm:"column:user_id"`
	IsRead    app.NullBool   `json:"is_read"             db:"m.is_read"         gorm:"column:is_read"`
	Data      app.NullJSON   `json:"data"                db:"m.data"            gorm:"column:data"`
	Judul     app.NullText   `json:"judul"               db:"m.judul"           gorm:"column:judul"`
	Deskripsi app.NullText   `json:"deskripsi"           db:"m.deskripsi"       gorm:"column:deskripsi"`

	CreatedBy         app.NullInt64    `json:"created_by.id"       db:"m.created_by"      gorm:"column:created_by"`
	CreatedByUsername app.NullString   `json:"created_by.username" db:"cbuser.username"   gorm:"-"`
	UpdatedBy         app.NullInt64    `json:"updated_by.id"       db:"m.updated_by"      gorm:"column:updated_by"`
	UpdatedByUsername app.NullString   `json:"updated_by.username" db:"ubuser.username"   gorm:"-"`
	DeletedBy         app.NullInt64    `json:"deleted_by.id"       db:"m.deleted_by"      gorm:"column:deleted_by"`
	DeletedByUsername app.NullString   `json:"deleted_by.username" db:"dbuser.username"   gorm:"-"`
	CreatedAt         app.NullDateTime `json:"created_at"          db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt         app.NullDateTime `json:"updated_at"          db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt         app.NullDateTime `json:"deleted_at"          db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

func (Notifikasi) EndPoint() string {
	return "notifikasi"
}

func (Notifikasi) TableVersion() string {
	return "28.06.291152"
}

func (Notifikasi) TableName() string {
	return "notifikasi"
}

func (Notifikasi) TableAliasName() string {
	return "m"
}

func (m *Notifikasi) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_user", "cbuser", []map[string]any{{"column1": "cbuser.id_user", "column2": "m.created_by"}})
	m.AddRelation("left", "m_user", "ubuser", []map[string]any{{"column1": "ubuser.id_user", "column2": "m.updated_by"}})
	m.AddRelation("left", "m_user", "dbuser", []map[string]any{{"column1": "dbuser.id_user", "column2": "m.deleted_by"}})

	return m.Relations
}

func (m *Notifikasi) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *Notifikasi) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *Notifikasi) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *Notifikasi) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (Notifikasi) OpenAPISchemaName() string {
	return "Notifikasi"
}

func (m *Notifikasi) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type NotifikasiList struct {
	app.ListModel
}

func (NotifikasiList) OpenAPISchemaName() string {
	return "NotifikasiList"
}

func (p *NotifikasiList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Notifikasi{})
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
