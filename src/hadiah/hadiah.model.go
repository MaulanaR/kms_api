package hadiah

import "github.com/maulanar/kms/app"

type Hadiah struct {
	app.Model
	ID app.NullInt64 `json:"id"                  db:"m.id"              gorm:"column:id;primaryKey"`

	Name           app.NullString `json:"nama"                db:"m.nama"            gorm:"column:nama"`
	Desc           app.NullText   `json:"deskripsi"           db:"m.deskripsi"       gorm:"column:deskripsi"`
	Point          app.NullInt64  `json:"point"               db:"m.point"           gorm:"column:point"`
	GambarID       app.NullInt64  `json:"gambar.id"           db:"m.gambar_id"       gorm:"column:gambar_id"`
	GambarFilename app.NullString `json:"gambar.filename"     db:"g.filename"        gorm:"-"`
	GambarUrl      app.NullString `json:"gambar.url"          db:"g.url"             gorm:"-"`
	IsActive       app.NullBool   `json:"is_active"           db:"m.is_active"       gorm:"column:is_active"`

	CreatedBy         app.NullInt64    `json:"created_by.id"       db:"m.created_by"      gorm:"column:created_by"`
	CreatedByUsername app.NullString   `json:"created_by.username" db:"cbuser.username"   gorm:"-"`
	UpdatedBy         app.NullInt64    `json:"updated_by.id"       db:"m.updated_by"      gorm:"column:updated_by"`
	UpdatedByUsername app.NullString   `json:"updated_by.username" db:"ubuser.username"   gorm:"-"`
	DeletedBy         app.NullInt64    `json:"deleted_by.id"       db:"m.deleted_by"      gorm:"column:deleted_by"`
	DeletedByUsername app.NullString   `json:"deleted_by.username" db:"dbuser.username"   gorm:"-"`
	CreatedAt         app.NullDateTime `json:"created_at"          db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt         app.NullDateTime `json:"updated_at"          db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt         app.NullDateTime `json:"deleted_at"          db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

func (Hadiah) EndPoint() string {
	return "hadiah"
}

func (Hadiah) TableVersion() string {
	return "28.06.291152"
}

func (Hadiah) TableName() string {
	return "m_hadiah"
}

func (Hadiah) TableAliasName() string {
	return "m"
}

func (m *Hadiah) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_attachments", "g", []map[string]any{{"column1": "g.id", "column2": "m.gambar_id"}})

	m.AddRelation("left", "m_user", "cbuser", []map[string]any{{"column1": "cbuser.id_user", "column2": "m.created_by"}})
	m.AddRelation("left", "m_user", "ubuser", []map[string]any{{"column1": "ubuser.id_user", "column2": "m.updated_by"}})
	m.AddRelation("left", "m_user", "dbuser", []map[string]any{{"column1": "dbuser.id_user", "column2": "m.deleted_by"}})

	return m.Relations
}

func (m *Hadiah) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *Hadiah) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *Hadiah) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *Hadiah) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (Hadiah) OpenAPISchemaName() string {
	return "Hadiah"
}

func (m *Hadiah) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type HadiahList struct {
	app.ListModel
}

func (HadiahList) OpenAPISchemaName() string {
	return "HadiahList"
}

func (p *HadiahList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Hadiah{})
}

type ParamCreate struct {
	UseCaseHandler
}

type ParamUpdate struct {
	UseCaseHandler
}

type ParamPartiallyUpdate struct {
	UseCaseHandler
}

type ParamDelete struct {
	UseCaseHandler
}

type RankingPoint struct {
	app.Model
	ID                 app.NullInt64  `json:"id"             db:"m.id_user"                                                                                                                            gorm:"column:id_user;primaryKey"`
	OrangId            app.NullInt64  `json:"orang.id"       db:"m.id_orang"                                                                                                                           gorm:"column:id_orang"`
	OrangNama          app.NullString `json:"nama_lengkap"   db:"o.nama"                                                                                                                               gorm:"-"`
	OrangNamaPanggilan app.NullString `json:"nama_panggilan" db:"o.nama_panggilan"                                                                                                                     gorm:"-"`
	OrangNip           app.NullString `json:"nip"            db:"o.nip,hide"                                                                                                                           gorm:"-"`
	OrangNik           app.NullString `json:"nik"            db:"o.nik,hide"                                                                                                                           gorm:"-"`
	OrangTempatLahir   app.NullString `json:"tempat_lahir"   db:"o.tempat_lahir"                                                                                                                       gorm:"-"`
	OrangTglLahir      app.NullDate   `json:"tgl_lahir"      db:"o.tgl_lahir"                                                                                                                          gorm:"-"`
	OrangJenisKelamin  app.NullString `json:"jenis_kelamin"  db:"o.jenis_kelamin"                                                                                                                      gorm:"-"`
	OrangAlamat        app.NullString `json:"alamat"         db:"o.alamat"                                                                                                                             gorm:"-"`
	OrangEmail         app.NullString `json:"email"          db:"o.email"                                                                                                                              gorm:"-"`
	OrangTelp          app.NullString `json:"telp"           db:"o.telp"                                                                                                                               gorm:"-"`
	OrangJabatan       app.NullString `json:"jabatan"        db:"o.jabatan"                                                                                                                            gorm:"-"`
	OrangFotoID        app.NullInt64  `json:"foto.id"        db:"o.foto"                                                                                                                               gorm:"-"`
	OrangFotoUrl       app.NullString `json:"foto.url"       db:"att.url"                                                                                                                              gorm:"-"`
	OrangFotoNama      app.NullString `json:"foto.nama"      db:"att.filename"                                                                                                                         gorm:"-"`
	OrangUnitKerja     app.NullString `json:"unit_kerja"     db:"o.unit_kerja"                                                                                                                         gorm:"-"`
	OrangUserLevel     app.NullString `json:"user_level"     db:"o.user_level"                                                                                                                         gorm:"-"`
	OrangStatusLevel   app.NullString `json:"status_level"   db:"o.status_level"                                                                                                                       gorm:"-"`
	Username           app.NullString `json:"username"       db:"m.username"                                                                                                                           gorm:"column:username"`
	UsernameStara      app.NullString `json:"username_stara" db:"m.username_stara"                                                                                                                     gorm:"column:username_stara"`
	Kategori           app.NullString `json:"kategori"       db:"m.kategori"                                                                                                                           gorm:"column:kategori"           validate:"omitempty,oneof='BPKP' 'UMUM' 'APIP'"`
	Level              app.NullString `json:"level"          db:"m.level"                                                                                                                              gorm:"column:level"`
	Points             app.NullInt64  `json:"total_point"    db:"(SELECT thp.after FROM t_history_points thp WHERE thp.id_user = m.id_user ORDER BY thp.updated_at DESC, thp.created_at DESC LIMIT 1)" gorm:"-"`
}

func (RankingPoint) EndPoint() string {
	return "hadiah"
}

func (RankingPoint) TableVersion() string {
	return "28.06.291152"
}

func (RankingPoint) TableName() string {
	return "m_user"
}

func (RankingPoint) TableAliasName() string {
	return "m"
}

func (m *RankingPoint) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_orang", "o", []map[string]any{{"column1": "o.id_orang", "column2": "m.id_orang"}})
	m.AddRelation("left", "m_attachments", "att", []map[string]any{{"column1": "att.id", "column2": "o.foto"}})
	return m.Relations
}

func (m *RankingPoint) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *RankingPoint) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "total_point", "direction": "desc"})
	return m.Sorts
}

func (m *RankingPoint) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *RankingPoint) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (RankingPoint) OpenAPISchemaName() string {
	return "RankingPoint"
}

func (m *RankingPoint) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type RankingPointList struct {
	app.ListModel
}

func (RankingPointList) OpenAPISchemaName() string {
	return "RankingPointList"
}

func (p *RankingPointList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&RankingPoint{})
}
