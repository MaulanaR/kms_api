package subjenispengetahuan

import "github.com/maulanar/kms/app"

// SubjenisPengetahuan is the main model of SubjenisPengetahuan data. It provides a convenient interface for app.ModelInterface
type SubjenisPengetahuan struct {
	app.Model
	ID app.NullInt64 `json:"id"                     db:"m.id_subjenis_pengetahuan"   gorm:"column:id_subjenis_pengetahuan;primaryKey"`
	// JenisPengetahuanID   app.NullInt64    `json:"jenis_pengetahuan.id"   db:"m.id_jenis_pengetahuan"      gorm:"column:id_jenis_pengetahuan"`
	// JenisPengetahuanNama app.NullText     `json:"jenis_pengetahuan.nama" db:"jp.nama_jenis_pengetahuan"   gorm:"-"`
	Nama      app.NullText     `json:"nama"                   db:"m.nama_subjenis_pengetahuan" gorm:"column:nama_subjenis_pengetahuan"`
	CreatedAt app.NullDateTime `json:"created_at"             db:"m.created_at"                gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at"             db:"m.updated_at"                gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at"             db:"m.deleted_at,hide"           gorm:"column:deleted_at"`
}

// EndPoint returns the SubjenisPengetahuan end point, it used for cache key, etc.
func (SubjenisPengetahuan) EndPoint() string {
	return "subjenis_pengetahuan"
}

// TableVersion returns the versions of the SubjenisPengetahuan table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (SubjenisPengetahuan) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the SubjenisPengetahuan table in the database.
func (SubjenisPengetahuan) TableName() string {
	return "m_subjenis_pengetahuan"
}

// TableAliasName returns the table alias name of the SubjenisPengetahuan table, used for querying.
func (SubjenisPengetahuan) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the SubjenisPengetahuan data in the database, used for querying.
func (m *SubjenisPengetahuan) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "m_jenis_pengetahuan", "jp", []map[string]any{{"column1": "jp.id_jenis_pengetahuan", "column2": "m.id_jenis_pengetahuan"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the SubjenisPengetahuan data in the database, used for querying.
func (m *SubjenisPengetahuan) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the SubjenisPengetahuan data in the database, used for querying.
func (m *SubjenisPengetahuan) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the SubjenisPengetahuan data in the database, used for querying.
func (m *SubjenisPengetahuan) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the SubjenisPengetahuan schema, used for querying.
func (m *SubjenisPengetahuan) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the SubjenisPengetahuan schema in the open api documentation.
func (SubjenisPengetahuan) OpenAPISchemaName() string {
	return "SubjenisPengetahuan"
}

// GetOpenAPISchema returns the Open API Schema of the SubjenisPengetahuan in the open api documentation.
func (m *SubjenisPengetahuan) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type SubjenisPengetahuanList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the SubjenisPengetahuanList schema in the open api documentation.
func (SubjenisPengetahuanList) OpenAPISchemaName() string {
	return "SubjenisPengetahuanList"
}

// GetOpenAPISchema returns the Open API Schema of the SubjenisPengetahuanList in the open api documentation.
func (p *SubjenisPengetahuanList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&SubjenisPengetahuan{})
}

// ParamCreate is the expected parameters for create a new SubjenisPengetahuan data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the SubjenisPengetahuan data.
type ParamUpdate struct {
	UseCaseHandler
	Nama app.NullText `json:"nama" db:"m.nama_subjenis_pengetahuan" gorm:"column:nama_subjenis_pengetahuan" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the SubjenisPengetahuan data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
}

// ParamDelete is the expected parameters for delete the SubjenisPengetahuan data.
type ParamDelete struct {
	UseCaseHandler
}
