package akademi

import "github.com/maulanar/kms/app"

// Akademi is the main model of Akademi data. It provides a convenient interface for app.ModelInterface
type Akademi struct {
	app.Model
	ID        app.NullInt64    `json:"id"         db:"m.id_akademi"      gorm:"column:id_akademi;primaryKey"`
	Nama      app.NullText     `json:"nama"       db:"m.nama_akademi"    gorm:"column:nama_akademi"`
	CreatedAt app.NullDateTime `json:"created_at" db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at" db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at" db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the Akademi end point, it used for cache key, etc.
func (Akademi) EndPoint() string {
	return "akademi"
}

// TableVersion returns the versions of the Akademi table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Akademi) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the Akademi table in the database.
func (Akademi) TableName() string {
	return "m_akademi"
}

// TableAliasName returns the table alias name of the Akademi table, used for querying.
func (Akademi) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Akademi data in the database, used for querying.
func (m *Akademi) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Akademi data in the database, used for querying.
func (m *Akademi) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Akademi data in the database, used for querying.
func (m *Akademi) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Akademi data in the database, used for querying.
func (m *Akademi) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Akademi schema, used for querying.
func (m *Akademi) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Akademi schema in the open api documentation.
func (Akademi) OpenAPISchemaName() string {
	return "Akademi"
}

// GetOpenAPISchema returns the Open API Schema of the Akademi in the open api documentation.
func (m *Akademi) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type AkademiList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the AkademiList schema in the open api documentation.
func (AkademiList) OpenAPISchemaName() string {
	return "AkademiList"
}

// GetOpenAPISchema returns the Open API Schema of the AkademiList in the open api documentation.
func (p *AkademiList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Akademi{})
}

// ParamCreate is the expected parameters for create a new Akademi data.
type ParamCreate struct {
	UseCaseHandler
	Nama app.NullText `json:"nama" db:"m.nama_akademi" gorm:"column:nama_akademi" validate:"required"`
}

// ParamUpdate is the expected parameters for update the Akademi data.
type ParamUpdate struct {
	UseCaseHandler
}

// ParamPartiallyUpdate is the expected parameters for partially update the Akademi data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
}

// ParamDelete is the expected parameters for delete the Akademi data.
type ParamDelete struct {
	UseCaseHandler
}
