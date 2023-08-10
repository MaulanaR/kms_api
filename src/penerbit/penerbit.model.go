package penerbit

import "github.com/maulanar/kms/app"

type Penerbit struct {
	app.Model
	ID        app.NullInt64    `json:"id"         db:"m.id"              gorm:"column:id;primaryKey;AutoIncrement;auto_increment"`
	Nama      app.NullString   `json:"nama"       db:"m.nama"            gorm:"column:nama"`
	CreatedAt app.NullDateTime `json:"created_at" db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at" db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at" db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

func (Penerbit) EndPoint() string {
	return "penerbit"
}

func (Penerbit) TableVersion() string {
	return "28.06.291152"
}

func (Penerbit) TableName() string {
	return "m_penerbit"
}

func (Penerbit) TableAliasName() string {
	return "m"
}

func (m *Penerbit) GetRelations() map[string]map[string]any {

	return m.Relations
}

func (m *Penerbit) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *Penerbit) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *Penerbit) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *Penerbit) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (Penerbit) OpenAPISchemaName() string {
	return "Penerbit"
}

func (m *Penerbit) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type PenerbitList struct {
	app.ListModel
}

func (PenerbitList) OpenAPISchemaName() string {
	return "PenerbitList"
}

func (p *PenerbitList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Penerbit{})
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
