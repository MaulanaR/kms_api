package codegentemplate

import "github.com/maulanar/kms/app"

type CodeGenTemplate struct {
	app.Model
	ID app.NullInt64 `json:"id"         db:"m.id"              gorm:"column:id;primaryKey"`
	// AddField : DONT REMOVE THIS COMMENT
	CreatedAt app.NullDateTime `json:"created_at" db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at" db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at" db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

func (CodeGenTemplate) EndPoint() string {
	return "end_point"
}

func (CodeGenTemplate) TableVersion() string {
	return "28.06.291152"
}

func (CodeGenTemplate) TableName() string {
	return "end_point"
}

func (CodeGenTemplate) TableAliasName() string {
	return "m"
}

func (m *CodeGenTemplate) GetRelations() map[string]map[string]any {

	return m.Relations
}

func (m *CodeGenTemplate) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *CodeGenTemplate) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *CodeGenTemplate) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *CodeGenTemplate) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (CodeGenTemplate) OpenAPISchemaName() string {
	return "CodeGenTemplate"
}

func (m *CodeGenTemplate) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type CodeGenTemplateList struct {
	app.ListModel
}

func (CodeGenTemplateList) OpenAPISchemaName() string {
	return "CodeGenTemplateList"
}

func (p *CodeGenTemplateList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&CodeGenTemplate{})
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
