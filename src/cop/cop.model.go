package cop

import "github.com/maulanar/kms/app"

type Cop struct {
	app.Model
	ID                app.NullInt64    `json:"id"               db:"m.id"              gorm:"column:id;primaryKey"`
	Judul             app.NullText     `json:"judul"            db:"m.judul"           gorm:"column:judul"`
	Deskripsi         app.NullText     `json:"deskripsi"        db:"m.deskripsi"       gorm:"column:deskripsi"`
	GambarID          app.NullInt64    `json:"gambar.id"        db:"m.gambar_id"       gorm:"column:gambar_id"`
	GambarFilename    app.NullString   `json:"gambar.filename"  db:"g.filename"        gorm:"-"`
	GambarUrl         app.NullString   `json:"gambar.url"       db:"g.url"             gorm:"-"`
	DokumenID         app.NullInt64    `json:"dokumen.id"       db:"m.dokumen_id"      gorm:"column:dokumen_id"`
	DokumenFilename   app.NullString   `json:"dokumen.filename" db:"d.filename"        gorm:"-"`
	DokumenUrl        app.NullString   `json:"dokumen.url"      db:"d.url"             gorm:"-"`
	CreatedAt         app.NullDateTime `json:"created_at"       db:"m.created_at"      gorm:"column:created_at"`
	CreatedBy         app.NullInt64    `json:"created_by.id"             db:"m.created_by"                                                       gorm:"column:created_by"`
	CreatedByUsername app.NullString   `json:"created_by.username"       db:"cbuser.username"                                                    gorm:"-"`
	UpdatedBy         app.NullInt64    `json:"updated_by.id"             db:"m.updated_by"                                                       gorm:"column:updated_by"`
	UpdatedByUsername app.NullString   `json:"updated_by.username"       db:"ubuser.username"                                                    gorm:"-"`
	DeletedBy         app.NullInt64    `json:"deleted_by.id"             db:"m.deleted_by"                                                       gorm:"column:deleted_by"`
	DeletedByUsername app.NullString   `json:"deleted_by.username"       db:"dbuser.username"                                                    gorm:"-"`
	UpdatedAt         app.NullDateTime `json:"updated_at"       db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt         app.NullDateTime `json:"deleted_at"       db:"m.deleted_at,hide" gorm:"column:deleted_at"`

	//statistik View, like, dislike, komentar
	StatistikView     app.NullInt64 `json:"statistik.view"            db:"(CASE WHEN m.count_view > 0 THEN m.count_view ELSE 0 END)"         gorm:"column:count_view;default:0"`
	StatistikLike     app.NullInt64 `json:"statistik.like"            db:"(SELECT COUNT(*) FROM t_like_cop WHERE t_like_cop.id_cop=m.id)"              gorm:"-"`
	StatistikDislike  app.NullInt64 `json:"statistik.dislike"         db:"(SELECT COUNT(*) FROM t_dislike_cop WHERE t_dislike_cop.id_cop=m.id)"           gorm:"-"`
	StatistikKomentar app.NullInt64 `json:"statistik.komentar"        db:"(SELECT COUNT(*) FROM t_komentar_cop WHERE t_komentar_cop.id_cop=m.id)"          gorm:"-"`
	IsLiked           app.NullBool  `json:"is_liked"                  db:"-"                                                                  gorm:"-"`
	IsDisliked        app.NullBool  `json:"is_disliked"               db:"-"                                                                  gorm:"-"`
}

func (Cop) EndPoint() string {
	return "cop"
}

func (Cop) TableVersion() string {
	return "28.08.231152"
}

func (Cop) TableName() string {
	return "t_cop"
}

func (Cop) TableAliasName() string {
	return "m"
}

func (m *Cop) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_attachments", "g", []map[string]any{{"column1": "g.id", "column2": "m.gambar_id"}})
	m.AddRelation("left", "m_attachments", "d", []map[string]any{{"column1": "d.id", "column2": "m.dokumen_id"}})
	m.AddRelation("left", "m_user", "cbuser", []map[string]any{{"column1": "cbuser.id_user", "column2": "m.created_by"}})
	m.AddRelation("left", "m_user", "ubuser", []map[string]any{{"column1": "ubuser.id_user", "column2": "m.updated_by"}})
	m.AddRelation("left", "m_user", "dbuser", []map[string]any{{"column1": "dbuser.id_user", "column2": "m.deleted_by"}})

	return m.Relations
}

func (m *Cop) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *Cop) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *Cop) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *Cop) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (Cop) OpenAPISchemaName() string {
	return "Cop"
}

func (m *Cop) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type CopList struct {
	app.ListModel
}

func (CopList) OpenAPISchemaName() string {
	return "CopList"
}

func (p *CopList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Cop{})
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
