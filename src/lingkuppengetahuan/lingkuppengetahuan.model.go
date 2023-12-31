package lingkuppengetahuan

import "github.com/maulanar/kms/app"

// LingkupPengetahuan is the main model of LingkupPengetahuan data. It provides a convenient interface for app.ModelInterface
type LingkupPengetahuan struct {
	app.Model
	ID        app.NullInt64    `json:"id"         db:"m.id_lingkup_pengetahuan"   gorm:"column:id_lingkup_pengetahuan;primaryKey"`
	Nama      app.NullText     `json:"nama"       db:"m.nama_lingkup_pengetahuan" gorm:"column:nama_lingkup_pengetahuan"`
	CreatedAt app.NullDateTime `json:"created_at" db:"m.created_at"               gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at" db:"m.updated_at"               gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at" db:"m.deleted_at,hide"          gorm:"column:deleted_at"`
}

// EndPoint returns the LingkupPengetahuan end point, it used for cache key, etc.
func (LingkupPengetahuan) EndPoint() string {
	return "lingkup_pengetahuan"
}

// TableVersion returns the versions of the LingkupPengetahuan table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (LingkupPengetahuan) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the LingkupPengetahuan table in the database.
func (LingkupPengetahuan) TableName() string {
	return "m_lingkup_pengetahuan"
}

// TableAliasName returns the table alias name of the LingkupPengetahuan table, used for querying.
func (LingkupPengetahuan) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the LingkupPengetahuan data in the database, used for querying.
func (m *LingkupPengetahuan) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the LingkupPengetahuan data in the database, used for querying.
func (m *LingkupPengetahuan) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the LingkupPengetahuan data in the database, used for querying.
func (m *LingkupPengetahuan) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the LingkupPengetahuan data in the database, used for querying.
func (m *LingkupPengetahuan) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the LingkupPengetahuan schema, used for querying.
func (m *LingkupPengetahuan) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the LingkupPengetahuan schema in the open api documentation.
func (LingkupPengetahuan) OpenAPISchemaName() string {
	return "LingkupPengetahuan"
}

// GetOpenAPISchema returns the Open API Schema of the LingkupPengetahuan in the open api documentation.
func (m *LingkupPengetahuan) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type LingkupPengetahuanList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the LingkupPengetahuanList schema in the open api documentation.
func (LingkupPengetahuanList) OpenAPISchemaName() string {
	return "LingkupPengetahuanList"
}

// GetOpenAPISchema returns the Open API Schema of the LingkupPengetahuanList in the open api documentation.
func (p *LingkupPengetahuanList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&LingkupPengetahuan{})
}

// ParamCreate is the expected parameters for create a new LingkupPengetahuan data.
type ParamCreate struct {
	UseCaseHandler
	Nama app.NullText `json:"nama" db:"m.nama_lingkup_pengetahuan" gorm:"column:nama_lingkup_pengetahuan" validate:"required"`
}

// ParamUpdate is the expected parameters for update the LingkupPengetahuan data.
type ParamUpdate struct {
	UseCaseHandler
}

// ParamPartiallyUpdate is the expected parameters for partially update the LingkupPengetahuan data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
}

// ParamDelete is the expected parameters for delete the LingkupPengetahuan data.
type ParamDelete struct {
	UseCaseHandler
}
