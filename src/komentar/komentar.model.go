package komentar

import "github.com/maulanar/kms/app"

type Komentar struct {
	app.Model
	ID                     app.NullInt64    `json:"id"                       db:"m.id"                 gorm:"column:id;primaryKey"`
	PengetahuanID          app.NullInt64    `json:"pengetahuan.id"           db:"m.id_pengetahuan"     gorm:"column:id_pengetahuan"`
	forumID                app.NullInt64    `json:"forum.id"                 db:"m.id_forum"             gorm:"column:id_forum"`
	LeaderTalkID           app.NullInt64    `json:"leader_talk.id"           db:"m.id_leader_talk"          gorm:"column:id_leader_talk"`
	UserID                 app.NullInt64    `json:"user.id"                  db:"m.id_user"            gorm:"column:id_user"`
	UserOrangId            app.NullInt64    `json:"user.orang.id"            db:"u.id_orang,hide"      gorm:"-"`
	UserOrangNama          app.NullString   `json:"user.nama_lengkap"        db:"o.nama"               gorm:"-"`
	UserOrangNamaPanggilan app.NullString   `json:"user.nama_panggilan"      db:"o.nama_panggilan"     gorm:"-"`
	UserOrangJabatan       app.NullString   `json:"user.jabatan"             db:"o.jabatan"            gorm:"-"`
	UserOrangEmail         app.NullString   `json:"user.email"               db:"o.email"              gorm:"-"`
	UserOrangFotoID        app.NullInt64    `json:"user.foto.id"             db:"o.foto"               gorm:"-"`
	UserOrangFotoUrl       app.NullString   `json:"user.foto.url"            db:"att.url"              gorm:"-"`
	UserOrangFotoNama      app.NullString   `json:"user.foto.nama"           db:"att.filename"         gorm:"-"`
	UserOrangUnitKerja     app.NullString   `json:"user.unit_kerja"          db:"o.unit_kerja"         gorm:"-"`
	UserOrangUserLevel     app.NullString   `json:"user.user_level"          db:"o.user_level"         gorm:"-"`
	UserOrangStatusLevel   app.NullString   `json:"user.status_level"        db:"o.status_level"       gorm:"-"`
	UserOrangNip           app.NullString   `json:"user.nip"                 db:"o.nip"                gorm:"-"`
	UserUsername           app.NullString   `json:"user.username"            db:"u.username"           gorm:"-"`
	ParentKomentarID       app.NullInt64    `json:"parent_komentar.id"       db:"m.id_parent_komentar" gorm:"column:id_parent_komentar"`
	ParentKomentarKomentar app.NullText     `json:"parent_komentar.komentar" db:"pkom.komentar"        gorm:"-"`
	ParentKomentarStatus   app.NullString   `json:"parent_komentar.status"   db:"pkom.status"          gorm:"-"`
	Komentar               app.NullText     `json:"komentar"                 db:"m.komentar"           gorm:"column:komentar" validate:"required"`
	Status                 app.NullString   `json:"status"                   db:"m.status"             gorm:"column:status"`
	CreatedAt              app.NullDateTime `json:"created_at"               db:"m.created_at"         gorm:"column:created_at"`
	UpdatedAt              app.NullDateTime `json:"updated_at"               db:"m.updated_at"         gorm:"column:updated_at"`
	DeletedAt              app.NullDateTime `json:"deleted_at"               db:"m.deleted_at,hide"    gorm:"column:deleted_at"`
}

func (Komentar) EndPoint() string {
	return "komentar"
}

func (Komentar) TableVersion() string {
	return "28.06.301152"
}

func (Komentar) TableName() string {
	return "t_komentar"
}

func (Komentar) TableAliasName() string {
	return "m"
}

func (m *Komentar) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "t_komentar", "pkom", []map[string]any{{"column1": "pkom.id", "column2": "m.id_parent_komentar"}})
	m.AddRelation("left", "m_user", "u", []map[string]any{{"column1": "u.id_user", "column2": "m.id_user"}})
	m.AddRelation("left", "m_orang", "o", []map[string]any{{"column1": "o.id_orang", "column2": "u.id_orang"}})
	m.AddRelation("left", "m_attachments", "att", []map[string]any{{"column1": "att.id", "column2": "o.foto"}})

	return m.Relations
}

func (m *Komentar) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *Komentar) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *Komentar) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *Komentar) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (Komentar) OpenAPISchemaName() string {
	return "Komentar"
}

func (m *Komentar) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type KomentarList struct {
	app.ListModel
}

func (KomentarList) OpenAPISchemaName() string {
	return "KomentarList"
}

func (p *KomentarList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Komentar{})
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
