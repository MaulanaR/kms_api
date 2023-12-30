package kompetensi

import "github.com/maulanar/kms/app"

// Kompetensi is the main model of Kompetensi data. It provides a convenient interface for app.ModelInterface
type Kompetensi struct {
	app.Model
	ID app.NullInt64 `json:"id"         db:"m.id_kompetensi"   gorm:"column:id_kompetensi;primaryKey"`
	//id_mapping
	IDMapping app.NullInt64 `json:"id_mapping" db:"id_mapping" gorm:"column:id_mapping"`
	//nama_kompetensi_sdm
	NamaKompetensiSDM app.NullText `json:"nama_kompetensi_sdm"       db:"-" gorm:"-"`
	Nama              app.NullText `json:"nama"       db:"m.nama_kompetensi" gorm:"column:nama_kompetensi"`
	// id_kompetensi_sdm
	IDKompetensiSDM app.NullInt64 `json:"id_kompetensi_sdm" db:"id_kompetensi_sdm" gorm:"column:id_kompetensi_sdm"`
	// id_kategori_kms
	IdKategoriKms app.NullInt64 `json:"id_kategori_kms" db:"id_kategori_kms" gorm:"column:id_kategori_kms"`
	// nama_kategori_kms
	NamaKategoriKms app.NullText     `json:"nama_kategori_kms" db:"nama_kategori_kms" gorm:"column:nama_kategori_kms"`
	CreatedAt       app.NullDateTime `json:"created_at" db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt       app.NullDateTime `json:"updated_at" db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt       app.NullDateTime `json:"deleted_at" db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the Kompetensi end point, it used for cache key, etc.
func (Kompetensi) EndPoint() string {
	return "kompetensi"
}

// TableVersion returns the versions of the Kompetensi table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Kompetensi) TableVersion() string {
	return "28.12.301152"
}

// TableName returns the name of the Kompetensi table in the database.
func (Kompetensi) TableName() string {
	return "m_kompetensi"
}

// TableAliasName returns the table alias name of the Kompetensi table, used for querying.
func (Kompetensi) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Kompetensi data in the database, used for querying.
func (m *Kompetensi) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Kompetensi data in the database, used for querying.
func (m *Kompetensi) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Kompetensi data in the database, used for querying.
func (m *Kompetensi) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Kompetensi data in the database, used for querying.
func (m *Kompetensi) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Kompetensi schema, used for querying.
func (m *Kompetensi) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Kompetensi schema in the open api documentation.
func (Kompetensi) OpenAPISchemaName() string {
	return "Kompetensi"
}

// GetOpenAPISchema returns the Open API Schema of the Kompetensi in the open api documentation.
func (m *Kompetensi) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type KompetensiList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the KompetensiList schema in the open api documentation.
func (KompetensiList) OpenAPISchemaName() string {
	return "KompetensiList"
}

// GetOpenAPISchema returns the Open API Schema of the KompetensiList in the open api documentation.
func (p *KompetensiList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Kompetensi{})
}

// ParamCreate is the expected parameters for create a new Kompetensi data.
type ParamCreate struct {
	UseCaseHandler
	Nama app.NullText `json:"nama" db:"m.nama_kompetensi" gorm:"column:nama_kompetensi" validate:"required"`
}

// ParamUpdate is the expected parameters for update the Kompetensi data.
type ParamUpdate struct {
	UseCaseHandler
}

// ParamPartiallyUpdate is the expected parameters for partially update the Kompetensi data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
}

// ParamDelete is the expected parameters for delete the Kompetensi data.
type ParamDelete struct {
	UseCaseHandler
}

type LoginStara struct {
	HttpCode app.NullInt64  `json:"http_code"`
	Message  app.NullString `json:"message"`
	Token    app.NullText   `json:"data.token"`
}
