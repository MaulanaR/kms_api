package orang

import "github.com/maulanar/kms/app"

// Orang is the main model of Orang data. It provides a convenient interface for app.ModelInterface
type Orang struct {
	app.Model
	ID            app.NullInt64    `json:"id"             db:"m.id_orang"        gorm:"column:id_orang;primaryKey"`
	Nip           app.NullString   `json:"nip"            db:"m.nip"             gorm:"column:nip"`
	Nama          app.NullString   `json:"nama_lengkap"   db:"m.nama"            gorm:"column:nama"`
	NamaPanggilan app.NullString   `json:"nama_panggilan" db:"m.nama_panggilan"  gorm:"column:nama_panggilan"`
	Jabatan       app.NullString   `json:"jabatan"        db:"m.jabatan"         gorm:"column:jabatan"`
	Email         app.NullString   `json:"email"          db:"m.email"           gorm:"column:email"               validate:"email"`
	FotoID        app.NullInt64    `json:"foto.id"        db:"m.foto"            gorm:"column:foto"`
	FotoUrl       app.NullString   `json:"foto.url"       db:"att.url"           gorm:"-"`
	FotoNama      app.NullString   `json:"foto.nama"      db:"att.filename"      gorm:"-"`
	UnitKerja     app.NullString   `json:"unit_kerja"     db:"m.unit_kerja"      gorm:"column:unit_kerja"`
	UserLevel     app.NullString   `json:"user_level"     db:"m.user_level"      gorm:"column:user_level"`
	StatusLevel   app.NullString   `json:"status_level"   db:"m.status_level"    gorm:"column:status_level"`
	CreatedBy     app.NullInt64    `json:"created_by"     db:"m.created_by"      gorm:"column:created_by"`
	UpdatedBy     app.NullInt64    `json:"updated_by"     db:"m.updated_by"      gorm:"column:updated_by"`
	DeletedBy     app.NullInt64    `json:"deleted_by"     db:"m.deleted_by"      gorm:"column:deleted_by"`
	CreatedAt     app.NullDateTime `json:"created_at"     db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt     app.NullDateTime `json:"updated_at"     db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt     app.NullDateTime `json:"deleted_at"     db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the Orang end point, it used for cache key, etc.
func (Orang) EndPoint() string {
	return "orang"
}

// TableVersion returns the versions of the Orang table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Orang) TableVersion() string {
	return "28.07.011152"
}

// TableName returns the name of the Orang table in the database.
func (Orang) TableName() string {
	return "m_orang"
}

// TableAliasName returns the table alias name of the Orang table, used for querying.
func (Orang) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Orang data in the database, used for querying.
func (m *Orang) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_attachments", "att", []map[string]any{{"column1": "att.id", "column2": "m.foto"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Orang data in the database, used for querying.
func (m *Orang) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Orang data in the database, used for querying.
func (m *Orang) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Orang data in the database, used for querying.
func (m *Orang) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Orang schema, used for querying.
func (m *Orang) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Orang schema in the open api documentation.
func (Orang) OpenAPISchemaName() string {
	return "Orang"
}

// GetOpenAPISchema returns the Open API Schema of the Orang in the open api documentation.
func (m *Orang) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type OrangList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the OrangList schema in the open api documentation.
func (OrangList) OpenAPISchemaName() string {
	return "OrangList"
}

// GetOpenAPISchema returns the Open API Schema of the OrangList in the open api documentation.
func (p *OrangList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Orang{})
}

// ParamCreate is the expected parameters for create a new Orang data.
type ParamCreate struct {
	UseCaseHandler
	Nama    app.NullString `json:"nama_lengkap" db:"m.nama"    gorm:"column:nama"    validate:"required"`
	Jabatan app.NullString `json:"jabatan"      db:"m.jabatan" gorm:"column:jabatan" validate:"required"`
	Nip     app.NullString `json:"nip"          db:"m.nip"     gorm:"column:nip"     validate:"required"`
}

// ParamUpdate is the expected parameters for update the Orang data.
type ParamUpdate struct {
	UseCaseHandler
}

// ParamPartiallyUpdate is the expected parameters for partially update the Orang data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
}

// ParamDelete is the expected parameters for delete the Orang data.
type ParamDelete struct {
	UseCaseHandler
}
