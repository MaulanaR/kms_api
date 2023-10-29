package historypoint

import "github.com/maulanar/kms/app"

type HistoryPoint struct {
	app.Model
	ID                     app.NullInt64  `json:"id"                                    db:"m.id"                           gorm:"column:id;primaryKey"`
	UserID                 app.NullInt64  `json:"user.id"                               db:"m.id_user"                      gorm:"column:id_user"`
	UserOrangId            app.NullInt64  `json:"user.orang.id"                         db:"u.id_orang"                     gorm:"-"`
	UserOrangNama          app.NullString `json:"user.nama_lengkap"                     db:"o.nama"                         gorm:"-"`
	UserOrangNamaPanggilan app.NullString `json:"user.nama_panggilan"                   db:"o.nama_panggilan"               gorm:"-"`
	UserOrangJabatan       app.NullString `json:"user.jabatan"                          db:"o.jabatan"                      gorm:"-"`
	UserOrangEmail         app.NullString `json:"user.email"                            db:"o.email"                        gorm:"-"`
	UserOrangFotoID        app.NullInt64  `json:"user.foto.id"                          db:"o.foto"                         gorm:"-"`
	UserOrangFotoUrl       app.NullString `json:"user.foto.url"                         db:"att.url"                        gorm:"-"`
	UserOrangFotoNama      app.NullString `json:"user.foto.nama"                        db:"att.filename"                   gorm:"-"`
	UserOrangUnitKerja     app.NullString `json:"user.unit_kerja"                       db:"o.unit_kerja"                   gorm:"-"`
	UserOrangUserLevel     app.NullString `json:"user.user_level"                       db:"o.user_level"                   gorm:"-"`
	UserOrangStatusLevel   app.NullString `json:"user.status_level"                     db:"o.status_level"                 gorm:"-"`
	UserOrangNip           app.NullString `json:"user.nip"                              db:"o.nip"                          gorm:"-"`
	UserUsername           app.NullString `json:"user.username"                         db:"u.username"                     gorm:"-"`
	UserJenis              app.NullString `json:"user.jenis"                            db:"u.jenis"                        gorm:"-"`
	UserLevel              app.NullString `json:"user.level"                            db:"u.level"                        gorm:"-"`

	PengetahuanID          app.NullInt64  `json:"pengetahuan.id"                        db:"m.id_pengetahuan"               gorm:"column:id_pengetahuan"`
	JenisPengetahuanID     app.NullInt64  `json:"pengetahuan.jenis_pengetahuan.id"      db:"p.id_jenis_pengetahuan"         gorm:"-"`
	JenisPengtahuanNama    app.NullText   `json:"pengetahuan.jenis_pengetahuan.nama"    db:"jp.nama_jenis_pengetahuan"      gorm:"-"`
	SubJenisPengetahuanID  app.NullInt64  `json:"pengetahuan.subjenis_pengetahuan.id"   db:"p.id_subjenis_pengetahuan"      gorm:"-"`
	SubJenisPengtahuanNama app.NullText   `json:"pengetahuan.subjenis_pengetahuan.nama" db:"sjp.nama_subjenis_pengetahuan"  gorm:"-"`
	LingkupPengetahuanID   app.NullInt64  `json:"pengetahuan.lingkup_pengetahuan.id"    db:"p.id_lingkup_pengetahuan"       gorm:"-"`
	LingkupPengetahuanNama app.NullText   `json:"pengetahuan.lingkup_pengetahuan.nama"  db:"lp.nama_lingkup_pengetahuan"    gorm:"-"`
	StatusPengetahuanID    app.NullInt64  `json:"pengetahuan.status_pengetahuan.id"     db:"p.id_status_pengetahuan"        gorm:"-"`
	StatusPengetahuanNama  app.NullText   `json:"pengetahuan.status_pengetahuan.nama"   db:"status.nama_status_pengetahuan" gorm:"-"`
	Judul                  app.NullText   `json:"pengetahuan.judul"                     db:"p.judul"                        gorm:"-"`
	Ringkasan              app.NullText   `json:"pengetahuan.ringkasan"                 db:"p.ringkasan"                    gorm:"-"`
	ThumbnailID            app.NullInt64  `json:"pengetahuan.thumbnail.id"              db:"p.thumbnail"                    gorm:"-"`
	ThumbnailName          app.NullString `json:"pengetahuan.thumbnail.nama"            db:"attachment.filename"            gorm:"-"`
	ThumbnailUrl           app.NullString `json:"pengetahuan.thumbnail.url"             db:"attachment.url"                 gorm:"-"`

	Before          app.NullInt64 `json:"before"                                db:"m.before"                       gorm:"column:before"`
	AdjustmentPoint app.NullInt64 `json:"adjustment_point"                      db:"m.adjustment_point"             gorm:"column:adjustment_point"`
	After           app.NullInt64 `json:"after"                                 db:"m.after"                        gorm:"column:after"`

	CreatedBy         app.NullInt64    `json:"created_by.id"                         db:"m.created_by"                   gorm:"column:created_by"`
	CreatedByUsername app.NullString   `json:"created_by.username"                   db:"cbuser.username"                gorm:"-"`
	UpdatedBy         app.NullInt64    `json:"updated_by.id"                         db:"m.updated_by"                   gorm:"column:updated_by"`
	UpdatedByUsername app.NullString   `json:"updated_by.username"                   db:"ubuser.username"                gorm:"-"`
	DeletedBy         app.NullInt64    `json:"deleted_by.id"                         db:"m.deleted_by"                   gorm:"column:deleted_by"`
	DeletedByUsername app.NullString   `json:"deleted_by.username"                   db:"dbuser.username"                gorm:"-"`
	CreatedAt         app.NullDateTime `json:"created_at"                            db:"m.created_at"                   gorm:"column:created_at"`
	UpdatedAt         app.NullDateTime `json:"updated_at"                            db:"m.updated_at"                   gorm:"column:updated_at"`
	DeletedAt         app.NullDateTime `json:"deleted_at"                            db:"m.deleted_at,hide"              gorm:"column:deleted_at"`
}

func (HistoryPoint) EndPoint() string {
	return "history_points"
}

func (HistoryPoint) TableVersion() string {
	return "28.06.291152"
}

func (HistoryPoint) TableName() string {
	return "t_history_points"
}

func (HistoryPoint) TableAliasName() string {
	return "m"
}

func (m *HistoryPoint) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_user", "cbuser", []map[string]any{{"column1": "cbuser.id_user", "column2": "m.created_by"}})
	m.AddRelation("left", "m_user", "ubuser", []map[string]any{{"column1": "ubuser.id_user", "column2": "m.updated_by"}})
	m.AddRelation("left", "m_user", "dbuser", []map[string]any{{"column1": "dbuser.id_user", "column2": "m.deleted_by"}})

	m.AddRelation("left", "m_user", "u", []map[string]any{{"column1": "u.id_user", "column2": "m.id_user"}})
	m.AddRelation("left", "m_orang", "o", []map[string]any{{"column1": "o.id_orang", "column2": "u.id_orang"}})
	m.AddRelation("left", "m_attachments", "att", []map[string]any{{"column1": "att.id", "column2": "o.foto"}})

	m.AddRelation("left", "t_pengetahuan", "p", []map[string]any{{"column1": "p.id_pengetahuan", "column2": "m.id_pengetahuan"}})
	m.AddRelation("left", "m_attachments", "attachment", []map[string]any{{"column1": "attachment.id", "column2": "p.thumbnail"}})
	m.AddRelation("left", "m_jenis_pengetahuan", "jp", []map[string]any{{"column1": "jp.id_jenis_pengetahuan", "column2": "p.id_jenis_pengetahuan"}})
	m.AddRelation("left", "m_subjenis_pengetahuan", "sjp", []map[string]any{{"column1": "sjp.id_subjenis_pengetahuan", "column2": "p.id_subjenis_pengetahuan"}})
	m.AddRelation("left", "m_status_pengetahuan", "status", []map[string]any{{"column1": "status.id_status_pengetahuan", "column2": "p.id_status_pengetahuan"}})
	m.AddRelation("left", "m_lingkup_pengetahuan", "lp", []map[string]any{{"column1": "lp.id_lingkup_pengetahuan", "column2": "p.id_lingkup_pengetahuan"}})

	return m.Relations
}

func (m *HistoryPoint) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *HistoryPoint) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	m.AddSort(map[string]any{"column": "m.created_at", "direction": "desc"})
	return m.Sorts
}

func (m *HistoryPoint) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *HistoryPoint) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (HistoryPoint) OpenAPISchemaName() string {
	return "HistoryPoint"
}

func (m *HistoryPoint) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type HistoryPointList struct {
	app.ListModel
}

func (HistoryPointList) OpenAPISchemaName() string {
	return "HistoryPointList"
}

func (p *HistoryPointList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&HistoryPoint{})
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
