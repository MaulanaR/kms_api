package user

import "github.com/maulanar/kms/app"

// User is the main model of User data. It provides a convenient interface for app.ModelInterface
type User struct {
	app.Model
	ID                 app.NullInt64    `json:"id"             db:"m.id_user"         gorm:"column:id_user;primaryKey"`
	OrangId            app.NullInt64    `json:"orang.id"       db:"m.id_orang,hide"   gorm:"column:id_orang"`
	OrangNama          app.NullString   `json:"nama_lengkap"   db:"o.nama"            gorm:"-"`
	OrangNamaPanggilan app.NullString   `json:"nama_panggilan" db:"o.nama_panggilan"  gorm:"-"`
	OrangJabatan       app.NullString   `json:"jabatan"        db:"o.jabatan"         gorm:"-"`
	OrangEmail         app.NullString   `json:"email"          db:"o.email"           gorm:"-"`
	OrangFoto          app.NullString   `json:"foto"           db:"o.foto"            gorm:"-"`
	OrangUnitKerja     app.NullString   `json:"unit_kerja"     db:"o.unit_kerja"      gorm:"-"`
	OrangUserLevel     app.NullString   `json:"user_level"     db:"o.user_level"      gorm:"-"`
	OrangStatusLevel   app.NullString   `json:"status_level"   db:"o.status_level"    gorm:"-"`
	OrangNip           app.NullString   `json:"nip"            db:"o.nip"             gorm:"-"`
	Username           app.NullString   `json:"username"       db:"m.username"        gorm:"column:username"`
	Jenis              app.NullString   `json:"jenis"          db:"m.jenis"           gorm:"column:jenis"`
	Password           app.NullString   `json:"password"       db:"m.password,hide"   gorm:"column:password"`
	CreatedAt          app.NullDateTime `json:"created_at"     db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt          app.NullDateTime `json:"updated_at"     db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt          app.NullDateTime `json:"deleted_at"     db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the User end point, it used for cache key, etc.
func (User) EndPoint() string {
	return "user"
}

// TableVersion returns the versions of the User table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (User) TableVersion() string {
	return "28.07.011154"
}

// TableName returns the name of the User table in the database.
func (User) TableName() string {
	return "m_user"
}

// TableAliasName returns the table alias name of the User table, used for querying.
func (User) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the User data in the database, used for querying.
func (m *User) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_orang", "o", []map[string]any{{"column1": "o.id_orang", "column2": "m.id_orang"}})
	return m.Relations
}

// GetFilters returns the filter of the User data in the database, used for querying.
func (m *User) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the User data in the database, used for querying.
func (m *User) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the User data in the database, used for querying.
func (m *User) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the User schema, used for querying.
func (m *User) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the User schema in the open api documentation.
func (User) OpenAPISchemaName() string {
	return "User"
}

// GetOpenAPISchema returns the Open API Schema of the User in the open api documentation.
func (m *User) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type UserList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the UserList schema in the open api documentation.
func (UserList) OpenAPISchemaName() string {
	return "UserList"
}

// GetOpenAPISchema returns the Open API Schema of the UserList in the open api documentation.
func (p *UserList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&User{})
}

// ParamCreate is the expected parameters for create a new User data.
type ParamCreate struct {
	UseCaseHandler
	Username app.NullString `json:"username" db:"m.username"      gorm:"column:username" validate:"required"`
	Password app.NullString `json:"password" db:"m.password,hide" gorm:"column:password" validate:"required"`
}

// ParamUpdate is the expected parameters for update the User data.
type ParamUpdate struct {
	UseCaseHandler
}

// ParamPartiallyUpdate is the expected parameters for partially update the User data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
}

// ParamDelete is the expected parameters for delete the User data.
type ParamDelete struct {
	UseCaseHandler
}
