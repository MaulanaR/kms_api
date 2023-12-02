package dokumenmap

import "github.com/maulanar/kms/app"

type DokumenMap struct {
	app.Model
	ID                app.NullInt64    `json:"id"                  db:"tpdokumen.id_pengetahuan_dokumen"                                                gorm:"column:id_pengetahuan_dokumen;primaryKey"`
	PengetahuanID     app.NullInt64    `json:"-"                   db:"tpdokumen.id_pengetahuan,hide"                                                   gorm:"column:id_pengetahuan"`
	AttachmentID      app.NullInt64    `json:"dokumen.id"          db:"tpdokumen.id_attachment"                                                         gorm:"column:id_attachment"`
	AttachmentNama    app.NullText     `json:"dokumen.nama"        db:"attachment.filename"                                                             gorm:"-"`
	AttachmentUrl     app.NullText     `json:"dokumen.url"         db:"attachment.url"                                                                  gorm:"-"`
	Status            app.NullString   `json:"broadcast.status"    db:"(CASE WHEN CURDATE() BETWEEN start AND end THEN 'aktif' ELSE 'tidak_aktif' END)" gorm:"-"`
	Start             app.NullDate     `json:"broadcast.start"     db:"tpdokumen.start"                                                                 gorm:"column:start"`
	End               app.NullDate     `json:"broadcast.end"       db:"tpdokumen.end"                                                                   gorm:"column:end"`
	CreatedAt         app.NullDateTime `json:"created_at"          db:"tpdokumen.created_at"                                                            gorm:"column:created_at"`
	CreatedBy         app.NullInt64    `json:"created_by.id"       db:"tpdokumen.created_by"                                                            gorm:"column:created_by"`
	CreatedByUsername app.NullString   `json:"created_by.username" db:"cbuser.username"                                                                 gorm:"-"`
	UpdatedBy         app.NullInt64    `json:"updated_by.id"       db:"tpdokumen.updated_by"                                                            gorm:"column:updated_by"`
	UpdatedByUsername app.NullString   `json:"updated_by.username" db:"ubuser.username"                                                                 gorm:"-"`
	DeletedBy         app.NullInt64    `json:"deleted_by.id"       db:"tpdokumen.deleted_by"                                                            gorm:"column:deleted_by"`
	DeletedByUsername app.NullString   `json:"deleted_by.username" db:"dbuser.username"                                                                 gorm:"-"`
	UpdatedAt         app.NullDateTime `json:"updated_at"          db:"tpdokumen.updated_at"                                                            gorm:"column:updated_at"`
	DeletedAt         app.NullDateTime `json:"deleted_at"          db:"tpdokumen.deleted_at,hide"                                                       gorm:"column:deleted_at"`
}

func (DokumenMap) EndPoint() string {
	return "dokumen_map"
}

func (DokumenMap) TableVersion() string {
	return "23.11.020947"
}

func (DokumenMap) TableName() string {
	return "t_pengetahuan_dokumen"
}

func (DokumenMap) TableAliasName() string {
	return "tpdokumen"
}

func (m *DokumenMap) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_attachments", "attachment", []map[string]any{{"column1": "attachment.id", "column2": "tpdokumen.id_attachment"}})
	m.AddRelation("left", "m_user", "cbuser", []map[string]any{{"column1": "cbuser.id_user", "column2": "tpdokumen.created_by"}})
	m.AddRelation("left", "m_user", "ubuser", []map[string]any{{"column1": "ubuser.id_user", "column2": "tpdokumen.updated_by"}})
	m.AddRelation("left", "m_user", "dbuser", []map[string]any{{"column1": "dbuser.id_user", "column2": "tpdokumen.deleted_by"}})

	return m.Relations
}

func (m *DokumenMap) GetFilters() []map[string]any {
	return m.Filters
}

func (m *DokumenMap) GetSorts() []map[string]any {
	return m.Sorts
}

func (m *DokumenMap) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *DokumenMap) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (DokumenMap) OpenAPISchemaName() string {
	return "DokumenMap"
}

func (m *DokumenMap) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type DokumenMapList struct {
	app.ListModel
}

func (DokumenMapList) OpenAPISchemaName() string {
	return "DokumenMapList"
}

func (p *DokumenMapList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&DokumenMap{})
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
