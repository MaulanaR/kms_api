package librarycafe

import "github.com/maulanar/kms/app"

type LibraryCafe struct {
	app.Model
	ID                app.NullInt64 `json:"id"                  db:"m.id"                                                                    gorm:"column:id;primaryKey"`
	JudulKegiatan     app.NullText  `json:"judul_kegiatan"      db:"m.judul_kegiatan"                                                        gorm:"column:judul_kegiatan"`
	NamaPenyelenggara app.NullText  `json:"nama_penyelenggara"  db:"m.nama_penyelenggara"                                                    gorm:"column:nama_penyelenggara"`
	LinkKegiatan      app.NullText  `json:"link_kegiatan"       db:"m.link_kegiatan"                                                         gorm:"column:link_kegiatan"`
	Deskripsi         app.NullText  `json:"deskripsi"           db:"m.deskripsi"                                                             gorm:"column:deskripsi"`
	// UserID                 app.NullInt64    `json:"user.id"             db:"m.id_user"         gorm:"column:id_user"`
	// UserOrangID            app.NullInt64    `json:"user.orang.id"       db:"u.id_orang,hide"   gorm:"-"`
	// UserOrangNama          app.NullString   `json:"user.nama_lengkap"   db:"o.nama"            gorm:"-"`
	// UserOrangNamaPanggilan app.NullString   `json:"user.nama_panggilan" db:"o.nama_panggilan"  gorm:"-"`
	// UserOrangJabatan       app.NullString   `json:"user.jabatan"        db:"o.jabatan"         gorm:"-"`
	// UserOrangEmail         app.NullString   `json:"user.email"          db:"o.email"           gorm:"-"`
	// UserOrangFotoID        app.NullInt64    `json:"user.foto.id"        db:"o.foto"            gorm:"-"`
	// UserOrangFotoUrl       app.NullString   `json:"user.foto.url"       db:"att.url"           gorm:"-"`
	// UserOrangFotoNama      app.NullString   `json:"user.foto.nama"      db:"att.filename"      gorm:"-"`
	// UserOrangUnitKerja     app.NullString   `json:"user.unit_kerja"     db:"o.unit_kerja"      gorm:"-"`
	// UserOrangUserLevel     app.NullString   `json:"user.user_level"     db:"o.user_level"      gorm:"-"`
	// UserOrangStatusLevel   app.NullString   `json:"user.status_level"   db:"o.status_level"    gorm:"-"`
	// UserOrangNip           app.NullString   `json:"user.nip"            db:"o.nip"             gorm:"-"`
	// UserUsername           app.NullString   `json:"user.username"       db:"u.username"        gorm:"-"`
	GambarID          app.NullInt64    `json:"gambar.id"           db:"m.gambar_id"                                                             gorm:"column:gambar_id"`
	GambarUrl         app.NullString   `json:"gambar.url"          db:"gbr.url"                                                                 gorm:"-"`
	GambarNama        app.NullString   `json:"gambar.nama"         db:"gbr.filename"                                                            gorm:"-"`
	DokumenID         app.NullInt64    `json:"dokumen.id"          db:"m.dokumen_id"                                                            gorm:"column:dokumen_id"`
	DokumenUrl        app.NullString   `json:"dokumen.url"         db:"dok.url"                                                                 gorm:"-"`
	DokumenNama       app.NullString   `json:"dokumen.nama"        db:"dok.filename"                                                            gorm:"-"`
	CreatedBy         app.NullInt64    `json:"created_by.id"       db:"m.created_by"                                                            gorm:"column:created_by"`
	CreatedByUsername app.NullString   `json:"created_by.username" db:"cbuser.username"                                                         gorm:"-"`
	UpdatedBy         app.NullInt64    `json:"updated_by.id"       db:"m.updated_by"                                                            gorm:"column:updated_by"`
	UpdatedByUsername app.NullString   `json:"updated_by.username" db:"ubuser.username"                                                         gorm:"-"`
	DeletedBy         app.NullInt64    `json:"deleted_by.id"       db:"m.deleted_by"                                                            gorm:"column:deleted_by"`
	DeletedByUsername app.NullString   `json:"deleted_by.username" db:"dbuser.username"                                                         gorm:"-"`
	CreatedAt         app.NullDateTime `json:"created_at"          db:"m.created_at"                                                            gorm:"column:created_at"`
	UpdatedAt         app.NullDateTime `json:"updated_at"          db:"m.updated_at"                                                            gorm:"column:updated_at"`
	DeletedAt         app.NullDateTime `json:"deleted_at"          db:"m.deleted_at,hide"                                                       gorm:"column:deleted_at"`

	//statistik View, like, dislike, komentar
	StatistikView     app.NullInt64 `json:"statistik.view"      db:"(CASE WHEN m.count_view > 0 THEN m.count_view ELSE 0 END)"               gorm:"column:count_view;default:0"`
	StatistikLike     app.NullInt64 `json:"statistik.like"      db:"(SELECT COUNT(*) FROM t_like WHERE t_like.id_library_cafe=m.id)"         gorm:"-"`
	StatistikDislike  app.NullInt64 `json:"statistik.dislike"   db:"(SELECT COUNT(*) FROM t_dislike WHERE t_dislike.id_library_cafe=m.id)"   gorm:"-"`
	StatistikKomentar app.NullInt64 `json:"statistik.komentar"  db:"(SELECT COUNT(*) FROM t_komentar WHERE t_komentar.id_library_cafe=m.id)" gorm:"-"`
	IsLiked           app.NullBool  `json:"is_liked"            db:"-"                                                                       gorm:"-"`
	IsDisliked        app.NullBool  `json:"is_disliked"         db:"-"                                                                       gorm:"-"`
}

func (LibraryCafe) EndPoint() string {
	return "library_cafe"
}

func (LibraryCafe) TableVersion() string {
	return "23.11.291152"
}

func (LibraryCafe) TableName() string {
	return "t_library_cafe"
}

func (LibraryCafe) TableAliasName() string {
	return "m"
}

func (m *LibraryCafe) GetRelations() map[string]map[string]any {
	//m.AddRelation("left", "m_user", "u", []map[string]any{{"column1": "u.id_user", "column2": "m.id_user"}})
	//m.AddRelation("left", "m_orang", "o", []map[string]any{{"column1": "o.id_orang", "column2": "u.id_orang"}})
	m.AddRelation("left", "m_attachments", "dok", []map[string]any{{"column1": "dok.id", "column2": "m.dokumen_id"}})
	m.AddRelation("left", "m_attachments", "gbr", []map[string]any{{"column1": "gbr.id", "column2": "m.gambar_id"}})
	m.AddRelation("left", "m_user", "cbuser", []map[string]any{{"column1": "cbuser.id_user", "column2": "m.created_by"}})
	m.AddRelation("left", "m_user", "ubuser", []map[string]any{{"column1": "ubuser.id_user", "column2": "m.updated_by"}})
	m.AddRelation("left", "m_user", "dbuser", []map[string]any{{"column1": "dbuser.id_user", "column2": "m.deleted_by"}})

	return m.Relations
}

func (m *LibraryCafe) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *LibraryCafe) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *LibraryCafe) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *LibraryCafe) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (LibraryCafe) OpenAPISchemaName() string {
	return "LibraryCafe"
}

func (m *LibraryCafe) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type LibraryCafeList struct {
	app.ListModel
}

func (LibraryCafeList) OpenAPISchemaName() string {
	return "LibraryCafeList"
}

func (p *LibraryCafeList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&LibraryCafe{})
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
