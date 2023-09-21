package pedoman

import "github.com/maulanar/kms/app"

type Pedoman struct {
	app.Model
	ID                app.NullInt64    `json:"id"                  db:"m.id"              gorm:"column:id;primaryKey;auto_increment;AutoIncrement;"`
	Nama              app.NullText     `json:"nama"                db:"m.nama"            gorm:"column:nama"`
	Data              app.NullText     `json:"data"                db:"m.data"            gorm:"column:data"`
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

func (Pedoman) EndPoint() string {
	return "pedoman"
}

func (Pedoman) TableVersion() string {
	return "28.06.291152"
}

func (Pedoman) TableName() string {
	return "m_pedoman"
}

func (Pedoman) TableAliasName() string {
	return "m"
}

func (m *Pedoman) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_user", "cbuser", []map[string]any{{"column1": "cbuser.id_user", "column2": "m.created_by"}})
	m.AddRelation("left", "m_user", "ubuser", []map[string]any{{"column1": "ubuser.id_user", "column2": "m.updated_by"}})
	m.AddRelation("left", "m_user", "dbuser", []map[string]any{{"column1": "dbuser.id_user", "column2": "m.deleted_by"}})

	return m.Relations
}

func (m *Pedoman) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *Pedoman) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *Pedoman) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *Pedoman) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (Pedoman) OpenAPISchemaName() string {
	return "Pedoman"
}

func (m *Pedoman) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type PedomanList struct {
	app.ListModel
}

func (PedomanList) OpenAPISchemaName() string {
	return "PedomanList"
}

func (p *PedomanList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Pedoman{})
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
