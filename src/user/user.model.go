package user

import "github.com/maulanar/kms/app"

// User is the main model of User data. It provides a convenient interface for app.ModelInterface
type User struct {
	app.Model
	ID                 app.NullInt64    `json:"id"             db:"m.id_user"                                                                                                                            gorm:"column:id_user;primaryKey"`
	OrangId            app.NullInt64    `json:"orang.id"       db:"m.id_orang"                                                                                                                      gorm:"column:id_orang"`
	OrangNama          app.NullString   `json:"nama_lengkap"   db:"o.nama"                                                                                                                               gorm:"-"`
	OrangNamaPanggilan app.NullString   `json:"nama_panggilan" db:"o.nama_panggilan"                                                                                                                     gorm:"-"`
	OrangNip           app.NullString   `json:"nip"            db:"o.nip"                                                                                                                                gorm:"-"`
	OrangNik           app.NullString   `json:"nik"            db:"o.nik"                                                                                                                                gorm:"-"`
	OrangTempatLahir   app.NullString   `json:"tempat_lahir"   db:"o.tempat_lahir"                                                                                                                       gorm:"-"`
	OrangTglLahir      app.NullDate     `json:"tgl_lahir"      db:"o.tgl_lahir"                                                                                                                          gorm:"-"`
	OrangJenisKelamin  app.NullString   `json:"jenis_kelamin"  db:"o.jenis_kelamin"                                                                                                                      gorm:"-"`
	OrangAlamat        app.NullString   `json:"alamat"         db:"o.alamat"                                                                                                                             gorm:"-"`
	OrangEmail         app.NullString   `json:"email"          db:"o.email"                                                                                                                              gorm:"-"`
	OrangTelp          app.NullString   `json:"telp"           db:"o.telp"                                                                                                                               gorm:"-"`
	OrangJabatan       app.NullString   `json:"jabatan"        db:"o.jabatan"                                                                                                                            gorm:"-"`
	OrangFotoID        app.NullInt64    `json:"foto.id"        db:"o.foto"                                                                                                                               gorm:"-"`
	OrangFotoUrl       app.NullString   `json:"foto.url"       db:"att.url"                                                                                                                              gorm:"-"`
	OrangFotoNama      app.NullString   `json:"foto.nama"      db:"att.filename"                                                                                                                         gorm:"-"`
	OrangUnitKerja     app.NullString   `json:"unit_kerja"     db:"o.unit_kerja"                                                                                                                         gorm:"-"`
	OrangUserLevel     app.NullString   `json:"user_level"     db:"o.user_level"                                                                                                                         gorm:"-"`
	OrangStatusLevel   app.NullString   `json:"status_level"   db:"o.status_level"                                                                                                                       gorm:"-"`
	Username           app.NullString   `json:"username"       db:"m.username"                                                                                                                           gorm:"column:username"`
	UsernameStara      app.NullString   `json:"username_stara" db:"m.username_stara"                                                                                                                     gorm:"column:username_stara"`
	Kategori           app.NullString   `json:"kategori"       db:"m.kategori"                                                                                                                           gorm:"column:kategori"           validate:"omitempty,oneof='BPKP' 'UMUM' 'APIP'"`
	Level              app.NullString   `json:"level"          db:"m.level"                                                                                                                              gorm:"column:level"`
	FollowingTags      []FollowdHastag  `json:"following_tags" db:"user.id={id}"                                                                                                                         gorm:"-"`
	Points             app.NullInt64    `json:"total_point"    db:"(SELECT thp.after FROM t_history_points thp WHERE thp.id_user = m.id_user ORDER BY thp.updated_at DESC, thp.created_at DESC LIMIT 1)" gorm:"-"`
	Password           app.NullString   `json:"password"       db:"m.password,hide"                                                                                                                      gorm:"column:password"`
	CreatedAt          app.NullDateTime `json:"created_at"     db:"m.created_at"                                                                                                                         gorm:"column:created_at"`
	UpdatedAt          app.NullDateTime `json:"updated_at"     db:"m.updated_at"                                                                                                                         gorm:"column:updated_at"`
	DeletedAt          app.NullDateTime `json:"deleted_at"     db:"m.deleted_at,hide"                                                                                                                    gorm:"column:deleted_at"`
}

// EndPoint returns the User end point, it used for cache key, etc.
func (User) EndPoint() string {
	return "user"
}

// TableVersion returns the versions of the User table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (User) TableVersion() string {
	return "23.12.0112302"
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
	m.AddRelation("left", "m_attachments", "att", []map[string]any{{"column1": "att.id", "column2": "o.foto"}})
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
	OrangNik   app.NullString `json:"nik"          db:"o.nik"           gorm:"-"               validate:"required"`
	OrangNama  app.NullString `json:"nama_lengkap" db:"o.nama"          gorm:"-"               validate:"required"`
	OrangEmail app.NullString `json:"email"        db:"o.email"         gorm:"-"               validate:"omitempty"`
	Password   app.NullString `json:"password"     db:"m.password,hide" gorm:"column:password" validate:"omitempty"`
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

// FollowdHastag is the main model of FollowdHastag data. It provides a convenient interface for app.ModelInterface
type FollowdHastag struct {
	app.Model
	ID         app.NullInt64 `json:"id"             db:"m.id"             gorm:"column:id;primaryKey"`
	UserID     app.NullInt64 `json:"user.id"        db:"m.id_user"        gorm:"column:id_user"`
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
