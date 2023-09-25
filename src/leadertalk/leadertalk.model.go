package leadertalk

import "github.com/maulanar/kms/app"

type LeaderTalk struct {
	app.Model
	ID                app.NullInt64    `json:"id"                  db:"m.id"                                                                   gorm:"column:id;primaryKey"`
	Nama              app.NullText     `json:"nama"                db:"m.nama"                                                                 gorm:"column:nama"`
	Jabatan           app.NullText     `json:"jabatan"             db:"m.jabatan"                                                              gorm:"column:jabatan"`
	NamaKegiatan      app.NullText     `json:"nama_kegiatan"       db:"m.nama_kegiatan"                                                        gorm:"column:nama_kegiatan"`
	IsiDokumen        app.NullText     `json:"isi_dokumen"         db:"m.isi_dokumen"                                                          gorm:"column:isi_dokumen"`
	DokumenID         app.NullInt64    `json:"dokumen.id"          db:"m.dokumen_id"                                                           gorm:"column:dokumen_id"`
	DokumenFilename   app.NullString   `json:"dokumen.filename"    db:"d.filename"                                                             gorm:"-"`
	DokumenUrl        app.NullString   `json:"dokumen.url"         db:"d.url"                                                                  gorm:"-"`
	CreatedAt         app.NullDateTime `json:"created_at"          db:"m.created_at"                                                           gorm:"column:created_at"`
	UpdatedAt         app.NullDateTime `json:"updated_at"          db:"m.updated_at"                                                           gorm:"column:updated_at"`
	DeletedAt         app.NullDateTime `json:"deleted_at"          db:"m.deleted_at,hide"                                                      gorm:"column:deleted_at"`
	CreatedBy         app.NullInt64    `json:"created_by.id"       db:"m.created_by"                                                           gorm:"column:created_by"`
	CreatedByUsername app.NullString   `json:"created_by.username" db:"cbuser.username"                                                        gorm:"-"`
	UpdatedBy         app.NullInt64    `json:"updated_by.id"       db:"m.updated_by"                                                           gorm:"column:updated_by"`
	UpdatedByUsername app.NullString   `json:"updated_by.username" db:"ubuser.username"                                                        gorm:"-"`
	DeletedBy         app.NullInt64    `json:"deleted_by.id"       db:"m.deleted_by"                                                           gorm:"column:deleted_by"`
	DeletedByUsername app.NullString   `json:"deleted_by.username" db:"dbuser.username"                                                        gorm:"-"`

	//statistik View, like, dislike, komentar
	StatistikView     app.NullInt64 `json:"statistik.view"      db:"(CASE WHEN m.count_view > 0 THEN m.count_view ELSE 0 END)"              gorm:"column:count_view;default:0"`
	StatistikLike     app.NullInt64 `json:"statistik.like"      db:"(SELECT COUNT(*) FROM t_like WHERE t_like.id_leader_talk=m.id)"         gorm:"-"`
	StatistikDislike  app.NullInt64 `json:"statistik.dislike"   db:"(SELECT COUNT(*) FROM t_dislike WHERE t_dislike.id_leader_talk=m.id)"   gorm:"-"`
	StatistikKomentar app.NullInt64 `json:"statistik.komentar"  db:"(SELECT COUNT(*) FROM t_komentar WHERE t_komentar.id_leader_talk=m.id)" gorm:"-"`
	IsLiked           app.NullBool  `json:"is_liked"            db:"-"                                                                      gorm:"-"`
	IsDisliked        app.NullBool  `json:"is_disliked"         db:"-"                                                                      gorm:"-"`
}

func (LeaderTalk) EndPoint() string {
	return "leader_talk"
}

func (LeaderTalk) TableVersion() string {
	return "28.06.291152"
}

func (LeaderTalk) TableName() string {
	return "t_leader_talk"
}

func (LeaderTalk) TableAliasName() string {
	return "m"
}

func (m *LeaderTalk) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_attachments", "d", []map[string]any{{"column1": "d.id", "column2": "m.dokumen_id"}})
	m.AddRelation("left", "m_user", "cbuser", []map[string]any{{"column1": "cbuser.id_user", "column2": "m.created_by"}})
	m.AddRelation("left", "m_user", "ubuser", []map[string]any{{"column1": "ubuser.id_user", "column2": "m.updated_by"}})
	m.AddRelation("left", "m_user", "dbuser", []map[string]any{{"column1": "dbuser.id_user", "column2": "m.deleted_by"}})
	return m.Relations
}

func (m *LeaderTalk) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *LeaderTalk) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.id", "direction": "desc"})
	return m.Sorts
}

func (m *LeaderTalk) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *LeaderTalk) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (LeaderTalk) OpenAPISchemaName() string {
	return "LeaderTalk"
}

func (m *LeaderTalk) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type LeaderTalkList struct {
	app.ListModel
}

func (LeaderTalkList) OpenAPISchemaName() string {
	return "LeaderTalkList"
}

func (p *LeaderTalkList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&LeaderTalk{})
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
