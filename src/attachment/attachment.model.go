package attachment

import (
	"mime/multipart"

	"github.com/maulanar/kms/app"
)

type Attachment struct {
	app.Model
	ID                app.NullInt64         `json:"id"               db:"m.id"               gorm:"column:id;primaryKey;autoIncrement"`
	File              *multipart.FileHeader `json:"file"             db:"-"                  gorm:"-"`
	Filename          app.NullString        `json:"filename"         db:"m.filename"         gorm:"column:filename"`
	Size              app.NullInt64         `json:"size"             db:"m.size"             gorm:"column:size"`
	Extension         app.NullString        `json:"extension"        db:"m.extension"        gorm:"column:extension"`
	StorageLocation   app.NullString        `json:"storage_location" db:"m.storage_location" gorm:"column:storage_location"`
	Url               app.NullString        `json:"url"              db:"m.url"              gorm:"column:url"`
	CreatedBy         app.NullInt64         `json:"created_by.id"       db:"m.created_by"       gorm:"column:created_by"`
	CreatedByUsername app.NullString        `json:"created_by.username" db:"cbuser.username"    gorm:"-"`
	UpdatedBy         app.NullInt64         `json:"updated_by.id"       db:"m.updated_by"       gorm:"column:updated_by"`
	UpdatedByUsername app.NullString        `json:"updated_by.username" db:"ubuser.username"    gorm:"-"`
	DeletedBy         app.NullInt64         `json:"deleted_by.id"       db:"m.deleted_by"       gorm:"column:deleted_by"`
	DeletedByUsername app.NullString        `json:"deleted_by.username" db:"dbuser.username"    gorm:"-"`
	CreatedAt         app.NullDateTime      `json:"created_at"       db:"m.created_at"       gorm:"column:created_at"`
	UpdatedAt         app.NullDateTime      `json:"updated_at"       db:"m.updated_at"       gorm:"column:updated_at"`
	DeletedAt         app.NullDateTime      `json:"deleted_at"       db:"m.deleted_at,hide"  gorm:"column:deleted_at"`
}

func (Attachment) EndPoint() string {
	return "attachments"
}

func (Attachment) TableVersion() string {
	return "23.10.051152"
}

func (Attachment) TableName() string {
	return "m_attachments"
}

func (Attachment) TableAliasName() string {
	return "m"
}

func (m *Attachment) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_user", "cbuser", []map[string]any{{"column1": "cbuser.id_user", "column2": "m.created_by"}})
	m.AddRelation("left", "m_user", "ubuser", []map[string]any{{"column1": "ubuser.id_user", "column2": "m.updated_by"}})
	m.AddRelation("left", "m_user", "dbuser", []map[string]any{{"column1": "dbuser.id_user", "column2": "m.deleted_by"}})
	return m.Relations
}

func (m *Attachment) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *Attachment) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *Attachment) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *Attachment) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (Attachment) OpenAPISchemaName() string {
	return "Attachment"
}

func (m *Attachment) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type AttachmentList struct {
	app.ListModel
}

func (AttachmentList) OpenAPISchemaName() string {
	return "AttachmentList"
}

func (p *AttachmentList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Attachment{})
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
