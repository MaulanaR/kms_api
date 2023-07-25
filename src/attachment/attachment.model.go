package attachment

import (
	"mime/multipart"

	"github.com/maulanar/kms/app"
)

type Attachment struct {
	app.Model
	ID              app.NullInt64         `json:"id"               db:"m.id"               gorm:"column:id;primaryKey"`
	File            *multipart.FileHeader `json:"file"             db:""                   gorm:"-"`
	Filename        app.NullString        `json:"filename"         db:"m.filename"         gorm:"column:filename"`
	Size            app.NullInt64         `json:"size"             db:"m.size"             gorm:"column:size"`
	Extension       app.NullString        `json:"extension"        db:"m.extension"        gorm:"column:extension"`
	StorageLocation app.NullString        `json:"storage_location" db:"m.storage_location" gorm:"column:storage_location"`
	Url             app.NullString        `json:"url"              db:"m.url"              gorm:"column:url"`
	CreatedAt       app.NullDateTime      `json:"created_at"       db:"m.created_at"       gorm:"column:created_at"`
	UpdatedAt       app.NullDateTime      `json:"updated_at"       db:"m.updated_at"       gorm:"column:updated_at"`
	DeletedAt       app.NullDateTime      `json:"deleted_at"       db:"m.deleted_at,hide"  gorm:"column:deleted_at"`
}

func (Attachment) EndPoint() string {
	return "attachments"
}

func (Attachment) TableVersion() string {
	return "28.06.291152"
}

func (Attachment) TableName() string {
	return "attachments"
}

func (Attachment) TableAliasName() string {
	return "m"
}

func (m *Attachment) GetRelations() map[string]map[string]any {

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
