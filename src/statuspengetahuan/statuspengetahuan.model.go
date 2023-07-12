package statuspengetahuan

import "github.com/maulanar/kms/app"

// StatusPengetahuan is the main model of StatusPengetahuan data. It provides a convenient interface for app.ModelInterface
type StatusPengetahuan struct {
	app.Model
	ID        app.NullInt64    `json:"id"         db:"m.id_status_pengetahuan"   gorm:"column:id_status_pengetahuan;primaryKey"`
	Nama      app.NullText     `json:"nama"       db:"m.nama_status_pengetahuan" gorm:"column:nama_status_pengetahuan"`
	CreatedAt app.NullDateTime `json:"created_at" db:"m.created_at"              gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at" db:"m.updated_at"              gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at" db:"m.deleted_at,hide"         gorm:"column:deleted_at"`
}

// EndPoint returns the StatusPengetahuan end point, it used for cache key, etc.
func (StatusPengetahuan) EndPoint() string {
	return "status_pengetahuan"
}

// TableVersion returns the versions of the StatusPengetahuan table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (StatusPengetahuan) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the StatusPengetahuan table in the database.
func (StatusPengetahuan) TableName() string {
	return "m_status_pengetahuan"
}

// TableAliasName returns the table alias name of the StatusPengetahuan table, used for querying.
func (StatusPengetahuan) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the StatusPengetahuan data in the database, used for querying.
func (m *StatusPengetahuan) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the StatusPengetahuan data in the database, used for querying.
func (m *StatusPengetahuan) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the StatusPengetahuan data in the database, used for querying.
func (m *StatusPengetahuan) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the StatusPengetahuan data in the database, used for querying.
func (m *StatusPengetahuan) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the StatusPengetahuan schema, used for querying.
func (m *StatusPengetahuan) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the StatusPengetahuan schema in the open api documentation.
func (StatusPengetahuan) OpenAPISchemaName() string {
	return "StatusPengetahuan"
}

// GetOpenAPISchema returns the Open API Schema of the StatusPengetahuan in the open api documentation.
func (m *StatusPengetahuan) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type StatusPengetahuanList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the StatusPengetahuanList schema in the open api documentation.
func (StatusPengetahuanList) OpenAPISchemaName() string {
	return "StatusPengetahuanList"
}

// GetOpenAPISchema returns the Open API Schema of the StatusPengetahuanList in the open api documentation.
func (p *StatusPengetahuanList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&StatusPengetahuan{})
}

// ParamCreate is the expected parameters for create a new StatusPengetahuan data.
type ParamCreate struct {
	UseCaseHandler
	Nama app.NullText `json:"nama" db:"m.nama_status_pengetahuan" gorm:"column:nama_status_pengetahuan" validate:"required"`
}

// ParamUpdate is the expected parameters for update the StatusPengetahuan data.
type ParamUpdate struct {
	UseCaseHandler
}

// ParamPartiallyUpdate is the expected parameters for partially update the StatusPengetahuan data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
}

// ParamDelete is the expected parameters for delete the StatusPengetahuan data.
type ParamDelete struct {
	UseCaseHandler
}
