package tag

import "github.com/maulanar/kms/app"

// Tag is the main model of Tag data. It provides a convenient interface for app.ModelInterface
type Tag struct {
	app.Model
	ID        app.NullInt64    `json:"id"         db:"m.id_tag"          gorm:"column:id_tag;primaryKey; not null"`
	Nama      app.NullText     `json:"nama"       db:"m.nama_tag"        gorm:"column:nama_tag"`
	Jenis     app.NullText     `json:"nama"       db:"m.jenis"           gorm:"column:jenis"`
	CreatedAt app.NullDateTime `json:"created_at" db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at" db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at" db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the Tag end point, it used for cache key, etc.
func (Tag) EndPoint() string {
	return "tag"
}

// TableVersion returns the versions of the Tag table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Tag) TableVersion() string {
	return "24.02.211152"
}

// TableName returns the name of the Tag table in the database.
func (Tag) TableName() string {
	return "m_tag"
}

// TableAliasName returns the table alias name of the Tag table, used for querying.
func (Tag) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Tag data in the database, used for querying.
func (m *Tag) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Tag data in the database, used for querying.
func (m *Tag) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Tag data in the database, used for querying.
func (m *Tag) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Tag data in the database, used for querying.
func (m *Tag) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Tag schema, used for querying.
func (m *Tag) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Tag schema in the open api documentation.
func (Tag) OpenAPISchemaName() string {
	return "Tag"
}

// GetOpenAPISchema returns the Open API Schema of the Tag in the open api documentation.
func (m *Tag) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type TagList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the TagList schema in the open api documentation.
func (TagList) OpenAPISchemaName() string {
	return "TagList"
}

// GetOpenAPISchema returns the Open API Schema of the TagList in the open api documentation.
func (p *TagList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Tag{})
}

// ParamCreate is the expected parameters for create a new Tag data.
type ParamCreate struct {
	UseCaseHandler
	Nama app.NullText `json:"nama" db:"m.nama_tag" gorm:"column:nama_tag" validate:"required"`
}

// ParamUpdate is the expected parameters for update the Tag data.
type ParamUpdate struct {
	UseCaseHandler
}

// ParamPartiallyUpdate is the expected parameters for partially update the Tag data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
}

// ParamDelete is the expected parameters for delete the Tag data.
type ParamDelete struct {
	UseCaseHandler
}
