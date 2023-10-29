package accesstoken

import (
	"net/url"

	"github.com/maulanar/kms/app"
)

type AccessTokenHandler struct {
	AccessToken

	// injectable dependencies
	Ctx   *app.Ctx   `json:"-" db:"-" gorm:"-"`
	Query url.Values `json:"-" db:"-" gorm:"-"`
}

// AccessToken is the main model of AccessToken data. It provides a convenient interface for app.ModelInterface
type AccessToken struct {
	app.Model
	AccessToken        app.NullString   `json:"access_token"        db:"m.access_token"    gorm:"column:access_token;PrimaryKey"`
	ExpiredAt          app.NullDateTime `json:"expired_at"          db:"m.expired_at"      gorm:"column:expired_at"`
	UserId             app.NullInt64    `json:"user.id"             db:"m.id_user"         gorm:"column:id_user"`
	OrangId            app.NullInt64    `json:"user.orang.id"       db:"u.id_orang,hide"   gorm:"-"`
	OrangNama          app.NullString   `json:"user.nama_lengkap"   db:"o.nama"            gorm:"-"`
	OrangNamaPanggilan app.NullString   `json:"user.nama_panggilan" db:"o.nama_panggilan"  gorm:"-"`
	OrangJabatan       app.NullString   `json:"user.jabatan"        db:"o.jabatan"         gorm:"-"`
	OrangEmail         app.NullString   `json:"user.email"          db:"o.email"           gorm:"-"`
	OrangFoto          app.NullInt64    `json:"user.foto.id"        db:"o.foto"            gorm:"-"`
	OrangFotoUrl       app.NullString   `json:"user.foto.url"       db:"att.url"           gorm:"-"`
	OrangFotoNama      app.NullString   `json:"user.foto.nama"      db:"att.filename"      gorm:"-"`
	OrangUnitKerja     app.NullString   `json:"user.unit_kerja"     db:"o.unit_kerja"      gorm:"-"`
	OrangStatusLevel   app.NullString   `json:"user.status_level"   db:"o.status_level"    gorm:"-"`
	OrangNip           app.NullString   `json:"user.nip"            db:"o.nip"             gorm:"-"`
	Username           app.NullString   `json:"user.username"       db:"u.username"        gorm:"-"`
	Jenis              app.NullString   `json:"user.jenis"          db:"u.jenis"           gorm:"-"`
	LeveL              app.NullString   `json:"user.level"          db:"u.level"           gorm:"-"`
	Points             app.NullInt64    `json:"user.total_point"    db:"(SELECT thp.after FROM t_history_points thp WHERE thp.id_user = m.id_user ORDER BY thp.updated_at DESC, thp.created_at DESC LIMIT 1)"           gorm:"-"`
	Password           app.NullString   `json:"user.password"       db:"u.password"        gorm:"-"`
	IpAddress          app.NullString   `json:"ip_address"          db:"m.ip_address"      gorm:"column:ip_address"`
	CreatedAt          app.NullDateTime `json:"created_at"          db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt          app.NullDateTime `json:"updated_at"          db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt          app.NullDateTime `json:"deleted_at"          db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

// EndPoint returns the AccessToken end point, it used for cache key, etc.
func (AccessToken) EndPoint() string {
	return "access_token"
}

// TableVersion returns the versions of the AccessToken table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (AccessToken) TableVersion() string {
	return "23.10.291156"
}

// TableName returns the name of the AccessToken table in the database.
func (AccessToken) TableName() string {
	return "access_token"
}

// TableAliasName returns the table alias name of the AccessToken table, used for querying.
func (AccessToken) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the AccessToken data in the database, used for querying.
func (m *AccessToken) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_user", "u", []map[string]any{{"column1": "u.id_user", "column2": "m.id_user"}})
	m.AddRelation("left", "m_orang", "o", []map[string]any{{"column1": "o.id_orang", "column2": "u.id_orang"}})
	m.AddRelation("left", "m_attachments", "att", []map[string]any{{"column1": "att.id", "column2": "o.foto"}})
	return m.Relations
}

// GetFilters returns the filter of the AccessToken data in the database, used for querying.
func (m *AccessToken) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the AccessToken data in the database, used for querying.
func (m *AccessToken) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the AccessToken data in the database, used for querying.
func (m *AccessToken) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the AccessToken schema, used for querying.
func (m *AccessToken) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the AccessToken schema in the open api documentation.
func (AccessToken) OpenAPISchemaName() string {
	return "AccessToken"
}

// GetOpenAPISchema returns the Open API Schema of the AccessToken in the open api documentation.
func (m *AccessToken) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type AccessTokenList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the AccessTokenList schema in the open api documentation.
func (AccessTokenList) OpenAPISchemaName() string {
	return "AccessTokenList"
}

// GetOpenAPISchema returns the Open API Schema of the AccessTokenList in the open api documentation.
func (p *AccessTokenList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&AccessToken{})
}

// ParamCreate is the expected parameters for create a new AccessToken data.
type ParamCreate struct {
	app.Model
	Username app.NullString `json:"username" validate:"required"`
	Password app.NullString `json:"password" validate:"required"`
}

func (m *ParamCreate) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (ParamCreate) OpenAPISchemaName() string {
	return "Login"
}

func (m *ParamCreate) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

// ParamUpdate is the expected parameters for update the AccessToken data.
type ParamUpdate struct {
	UseCaseHandler
}

// ParamPartiallyUpdate is the expected parameters for partially update the AccessToken data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
}

// ParamDelete is the expected parameters for delete the AccessToken data.
type ParamDelete struct {
	UseCaseHandler
}
