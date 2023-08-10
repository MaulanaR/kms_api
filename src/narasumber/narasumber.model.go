package narasumber

import "github.com/maulanar/kms/app"

type Narasumber struct {
	app.Model
	ID        app.NullInt64    `json:"id"         db:"m.id"              gorm:"column:id;primaryKey;AutoIncrement;auto_increment"`
	Nama      app.NullString   `json:"nama"       db:"m.nama"            gorm:"column:nama"`
	CreatedAt app.NullDateTime `json:"created_at" db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at" db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at" db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

func (Narasumber) EndPoint() string {
	return "narasumber"
}

func (Narasumber) TableVersion() string {
	return "28.06.291152"
}

func (Narasumber) TableName() string {
	return "m_narasumber"
}

func (Narasumber) TableAliasName() string {
	return "m"
}

func (m *Narasumber) GetRelations() map[string]map[string]any {

	return m.Relations
}

func (m *Narasumber) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *Narasumber) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *Narasumber) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *Narasumber) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (Narasumber) OpenAPISchemaName() string {
	return "Narasumber"
}

func (m *Narasumber) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type NarasumberList struct {
	app.ListModel
}

func (NarasumberList) OpenAPISchemaName() string {
	return "NarasumberList"
}

func (p *NarasumberList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Narasumber{})
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
