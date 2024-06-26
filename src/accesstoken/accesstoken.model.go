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
	AccessToken        app.NullString   `json:"access_token"        db:"m.access_token"                                                                                                                       gorm:"column:access_token;PrimaryKey"`
	ExpiredAt          app.NullDateTime `json:"expired_at"          db:"m.expired_at"                                                                                                                         gorm:"column:expired_at"`
	UserId             app.NullInt64    `json:"user.id"             db:"m.id_user"                                                                                                                            gorm:"column:id_user"`
	OrangId            app.NullInt64    `json:"user.orang.id"       db:"u.id_orang,hide"                                                                                                                      gorm:"-"`
	OrangNama          app.NullString   `json:"user.nama_lengkap"   db:"o.nama"                                                                                                                               gorm:"-"`
	OrangNamaPanggilan app.NullString   `json:"user.nama_panggilan" db:"o.nama_panggilan"                                                                                                                     gorm:"-"`
	OrangJabatan       app.NullString   `json:"user.jabatan"        db:"o.jabatan"                                                                                                                            gorm:"-"`
	OrangEmail         app.NullString   `json:"user.email"          db:"o.email"                                                                                                                              gorm:"-"`
	OrangFoto          app.NullInt64    `json:"user.foto.id"        db:"o.foto"                                                                                                                               gorm:"-"`
	OrangFotoUrl       app.NullString   `json:"user.foto.url"       db:"att.url"                                                                                                                              gorm:"-"`
	OrangFotoNama      app.NullString   `json:"user.foto.nama"      db:"att.filename"                                                                                                                         gorm:"-"`
	OrangUnitKerja     app.NullString   `json:"user.unit_kerja"     db:"o.unit_kerja"                                                                                                                         gorm:"-"`
	OrangStatusLevel   app.NullString   `json:"user.status_level"   db:"o.status_level"                                                                                                                       gorm:"-"`
	OrangNip           app.NullString   `json:"user.nip"            db:"o.nip"                                                                                                                                gorm:"-"`
	Username           app.NullString   `json:"user.username"       db:"u.username"                                                                                                                           gorm:"-"`
	Jenis              app.NullString   `json:"user.jenis"          db:"u.jenis"                                                                                                                              gorm:"-"`
	LeveL              app.NullString   `json:"user.level"          db:"u.level"                                                                                                                              gorm:"-"`
	FollowingTags      []FollowdHastag  `json:"following_tags"      db:"user.id={user.id}"                                                                                                                         gorm:"-"`
	Points             app.NullInt64    `json:"user.total_point"    db:"(SELECT thp.after FROM t_history_points thp WHERE thp.id_user = m.id_user ORDER BY thp.updated_at DESC, thp.created_at DESC LIMIT 1)" gorm:"-"`
	Password           app.NullString   `json:"user.password"       db:"u.password"                                                                                                                           gorm:"-"`
	IpAddress          app.NullString   `json:"ip_address"          db:"m.ip_address"                                                                                                                         gorm:"column:ip_address"`
	CreatedAt          app.NullDateTime `json:"created_at"          db:"m.created_at"                                                                                                                         gorm:"column:created_at"`
	UpdatedAt          app.NullDateTime `json:"updated_at"          db:"m.updated_at"                                                                                                                         gorm:"column:updated_at"`
	DeletedAt          app.NullDateTime `json:"deleted_at"          db:"m.deleted_at,hide"                                                                                                                    gorm:"column:deleted_at"`
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

type LoginStara struct {
	HttpCode     app.NullInt64  `json:"http_code"`
	Message      app.NullString `json:"message"`
	Name         app.NullText   `json:"data.user_info.name"`
	NamaGelar    app.NullText   `json:"data.user_info.nama_gelar"`
	Username     app.NullText   `json:"data.user_info.username"`
	RoleId       app.NullText   `json:"data.user_info.role_id"`
	UserNip      app.NullText   `json:"data.user_info.user_nip"`
	Nomorhp      app.NullText   `json:"data.user_info.nomorhp"`
	Aktif        app.NullText   `json:"data.user_info.aktif"`
	Nipbaru      app.NullText   `json:"data.user_info.nipbaru"`
	Jabatan      app.NullText   `json:"data.user_info.jabatan"`
	IsJab        app.NullText   `json:"data.user_info.is_jab"`
	IsJab2       app.NullText   `json:"data.user_info.is_jab_2"`
	JenisJab     app.NullText   `json:"data.user_info.jenis_jab"`
	IsDpil       app.NullText   `json:"data.user_info.is_dpil"`
	Version      app.NullText   `json:"data.user_info.version"`
	IsUrgent     app.NullText   `json:"data.user_info.is_urgent"`
	IsHut        app.NullText   `json:"data.user_info.is_hut"`
	IsAtasan     app.NullText   `json:"data.user_info.is_atasan"`
	Nik          app.NullText   `json:"data.user_info.nik"`
	JenisKelamin app.NullText   `json:"data.user_info.jenis_kelamin"`
	SNamaAgama   app.NullText   `json:"data.user_info.s_nama_agama"`
	Namaunit     app.NullText   `json:"data.user_info.namaunit"`
	KeySortUnit  app.NullText   `json:"data.user_info.key_sort_unit"`
	KelJab       app.NullText   `json:"data.user_info.kel_jab"`
	SKdJabdetail app.NullText   `json:"data.user_info.s_kd_jabdetail"`
}

// FollowdHastag is the main model of FollowdHastag data. It provides a convenient interface for app.ModelInterface
type FollowdHastag struct {
	app.Model
	ID         app.NullInt64 `json:"id"             db:"m.id"             gorm:"column:id;primaryKey"`
	UserID     app.NullInt64 `json:"user.id"        db:"m.id_user,hide"   gorm:"column:id_user"`
	HastagID   app.NullInt64 `json:"tag.id"         db:"m.id_tag"         gorm:"column:id_tag"`
	HastagNama app.NullText  `json:"tag.nama"       db:"t.nama_tag"       gorm:"-"`
}

// EndPoint returns the FollowdHastag end point, it used for cache key, etc.
func (FollowdHastag) EndPoint() string {
	return "followed_tags"
}

// TableVersion returns the versions of the FollowdHastag table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (FollowdHastag) TableVersion() string {
	return "23.12.0112302"
}

// TableName returns the name of the FollowdHastag table in the database.
func (FollowdHastag) TableName() string {
	return "m_user_followed_tags"
}

// TableAliasName returns the table alias name of the FollowdHastag table, used for querying.
func (FollowdHastag) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the FollowdHastag data in the database, used for querying.
func (m *FollowdHastag) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_tag", "t", []map[string]any{{"column1": "t.id_tag", "column2": "m.id_tag"}})
	return m.Relations
}

// GetFilters returns the filter of the FollowdHastag data in the database, used for querying.
func (m *FollowdHastag) GetFilters() []map[string]any {
	return m.Filters
}

// GetSorts returns the default sort of the FollowdHastag data in the database, used for querying.
func (m *FollowdHastag) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.id", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the FollowdHastag data in the database, used for querying.
func (m *FollowdHastag) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the FollowdHastag schema, used for querying.
func (m *FollowdHastag) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the FollowdHastag schema in the open api documentation.
func (FollowdHastag) OpenAPISchemaName() string {
	return "FollowdHastag"
}

// GetOpenAPISchema returns the Open API Schema of the FollowdHastag in the open api documentation.
func (m *FollowdHastag) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

// FORGOT PASSWORD
type RequestForgotPassword struct {
	app.Model
	Email app.NullString `json:"email" validate:"required"`
}

type ForgotPassword struct {
	app.Model
	Key            app.NullString `json:"key" validate:"required"`
	Password       app.NullString `json:"password" validate:"required"`
	ReTypePassword app.NullString `json:"re_password" validate:"required"`
}
