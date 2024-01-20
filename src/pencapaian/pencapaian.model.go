package pencapaian

import "github.com/maulanar/kms/app"

type Pencapaian struct {
	app.Model
	ID app.NullInt64 `json:"id"                     db:"m.id"              gorm:"column:id;primaryKey"`

	Tanggal                app.NullDateTime `json:"tanggal"                db:"m.tanggal"         gorm:"column:tanggal"`
	UserID                 app.NullInt64    `json:"user.id"                db:"m.id_user"         gorm:"column:id_user"       validate:"required"`
	UserOrangId            app.NullInt64    `json:"user.orang.id"          db:"u.id_orang"        gorm:"-"`
	UserOrangNama          app.NullString   `json:"user.nama_lengkap"      db:"o.nama"            gorm:"-"`
	UserOrangNamaPanggilan app.NullString   `json:"user.nama_panggilan"    db:"o.nama_panggilan"  gorm:"-"`
	UserOrangJabatan       app.NullString   `json:"user.jabatan"           db:"o.jabatan"         gorm:"-"`
	UserOrangEmail         app.NullString   `json:"user.email"             db:"o.email"           gorm:"-"`
	UserOrangFotoID        app.NullInt64    `json:"user.foto.id"           db:"o.foto"            gorm:"-"`
	UserOrangFotoUrl       app.NullString   `json:"user.foto.url"          db:"att.url"           gorm:"-"`
	UserOrangFotoNama      app.NullString   `json:"user.foto.nama"         db:"att.filename"      gorm:"-"`
	UserOrangUnitKerja     app.NullString   `json:"user.unit_kerja"        db:"o.unit_kerja"      gorm:"-"`
	UserOrangUserLevel     app.NullString   `json:"user.user_level"        db:"o.user_level"      gorm:"-"`
	UserOrangStatusLevel   app.NullString   `json:"user.status_level"      db:"o.status_level"    gorm:"-"`
	UserOrangNip           app.NullString   `json:"user.nip"               db:"o.nip"             gorm:"-"`
	UserUsername           app.NullString   `json:"user.username"          db:"u.username"        gorm:"-"`
	UserJenis              app.NullString   `json:"user.jenis"             db:"u.jenis"           gorm:"-"`
	UserLevel              app.NullString   `json:"user.level"             db:"u.level"           gorm:"-"`

	HadiahId             app.NullInt64  `json:"hadiah.id"              db:"m.hadiah_id"       gorm:"column:hadiah_id"     validate:"required"`
	HadiahName           app.NullString `json:"hadiah.nama"            db:"h.nama"            gorm:"-"`
	HadiahDesc           app.NullText   `json:"hadiah.deskripsi"       db:"h.deskripsi"       gorm:"-"`
	HadiahPoint          app.NullInt64  `json:"hadiah.point"           db:"h.point"           gorm:"-"`
	HadiahGambarID       app.NullInt64  `json:"hadiah.gambar.id"       db:"h.gambar_id"       gorm:"-"`
	HadiahGambarFilename app.NullString `json:"hadiah.gambar.filename" db:"g.filename"        gorm:"-"`
	HadiahGambarUrl      app.NullString `json:"hadiah.gambar.url"      db:"g.url"             gorm:"-"`
	HadiahIsActive       app.NullBool   `json:"hadiah.is_active"       db:"h.is_active"       gorm:"-"`

	CreatedBy         app.NullInt64    `json:"created_by.id"          db:"m.created_by"      gorm:"column:created_by"`
	CreatedByUsername app.NullString   `json:"created_by.username"    db:"cbuser.username"   gorm:"-"`
	UpdatedBy         app.NullInt64    `json:"updated_by.id"          db:"m.updated_by"      gorm:"column:updated_by"`
	UpdatedByUsername app.NullString   `json:"updated_by.username"    db:"ubuser.username"   gorm:"-"`
	DeletedBy         app.NullInt64    `json:"deleted_by.id"          db:"m.deleted_by"      gorm:"column:deleted_by"`
	DeletedByUsername app.NullString   `json:"deleted_by.username"    db:"dbuser.username"   gorm:"-"`
	CreatedAt         app.NullDateTime `json:"created_at"             db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt         app.NullDateTime `json:"updated_at"             db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt         app.NullDateTime `json:"deleted_at"             db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

func (Pencapaian) EndPoint() string {
	return "pencapaian"
}

func (Pencapaian) TableVersion() string {
	return "28.06.291152"
}

func (Pencapaian) TableName() string {
	return "t_pencapaian"
}

func (Pencapaian) TableAliasName() string {
	return "m"
}

func (m *Pencapaian) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_user", "u", []map[string]any{{"column1": "u.id_user", "column2": "m.id_user"}})
	m.AddRelation("left", "m_orang", "o", []map[string]any{{"column1": "o.id_orang", "column2": "u.id_orang"}})
	m.AddRelation("left", "m_attachments", "att", []map[string]any{{"column1": "att.id", "column2": "o.foto"}})

	m.AddRelation("left", "m_hadiah", "h", []map[string]any{{"column1": "h.id", "column2": "m.hadiah_id"}})
	m.AddRelation("left", "m_attachments", "g", []map[string]any{{"column1": "g.id", "column2": "h.gambar_id"}})

	m.AddRelation("left", "m_user", "cbuser", []map[string]any{{"column1": "cbuser.id_user", "column2": "m.created_by"}})
	m.AddRelation("left", "m_user", "ubuser", []map[string]any{{"column1": "ubuser.id_user", "column2": "m.updated_by"}})
	m.AddRelation("left", "m_user", "dbuser", []map[string]any{{"column1": "dbuser.id_user", "column2": "m.deleted_by"}})

	return m.Relations
}

func (m *Pencapaian) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *Pencapaian) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *Pencapaian) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *Pencapaian) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (Pencapaian) OpenAPISchemaName() string {
	return "Pencapaian"
}

func (m *Pencapaian) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type PencapaianList struct {
	app.ListModel
}

func (PencapaianList) OpenAPISchemaName() string {
	return "PencapaianList"
}

func (p *PencapaianList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Pencapaian{})
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
