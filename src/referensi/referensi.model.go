package referensi

import "github.com/maulanar/kms/app"

// Referensi is the main model of Referensi data. It provides a convenient interface for app.ModelInterface
type Referensi struct {
	app.Model
	ID            app.NullInt64    `json:"id"         db:"m.id_referensi"    gorm:"column:id_referensi;primaryKey"`
	NamaReferensi app.NullText     `json:"referensi"  db:"m.nama_referensi"  gorm:"column:nama_referensi"`
	CreatedAt     app.NullDateTime `json:"created_at" db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt     app.NullDateTime `json:"updated_at" db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt     app.NullDateTime `json:"deleted_at" db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the Referensi end point, it used for cache key, etc.
func (Referensi) EndPoint() string {
	return "referensi"
}

// TableVersion returns the versions of the Referensi table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Referensi) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the Referensi table in the database.
func (Referensi) TableName() string {
	return "m_referensi"
}

// TableAliasName returns the table alias name of the Referensi table, used for querying.
func (Referensi) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Referensi data in the database, used for querying.
func (m *Referensi) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Referensi data in the database, used for querying.
func (m *Referensi) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Referensi data in the database, used for querying.
func (m *Referensi) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Referensi data in the database, used for querying.
func (m *Referensi) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Referensi schema, used for querying.
func (m *Referensi) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Referensi schema in the open api documentation.
func (Referensi) OpenAPISchemaName() string {
	return "Referensi"
}

// GetOpenAPISchema returns the Open API Schema of the Referensi in the open api documentation.
func (m *Referensi) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type ReferensiList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the ReferensiList schema in the open api documentation.
func (ReferensiList) OpenAPISchemaName() string {
	return "ReferensiList"
}

// GetOpenAPISchema returns the Open API Schema of the ReferensiList in the open api documentation.
func (p *ReferensiList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Referensi{})
}

// ParamCreate is the expected parameters for create a new Referensi data.
type ParamCreate struct {
	UseCaseHandler
	NamaReferensi app.NullText `json:"referensi" db:"m.nama_referensi" gorm:"column:nama_referensi" validate:"required"`
}

// ParamUpdate is the expected parameters for update the Referensi data.
type ParamUpdate struct {
	UseCaseHandler
}

// ParamPartiallyUpdate is the expected parameters for partially update the Referensi data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
}

// ParamDelete is the expected parameters for delete the Referensi data.
type ParamDelete struct {
	UseCaseHandler
}
