package event

import (
	"github.com/maulanar/kms/app"
	"github.com/maulanar/kms/src/eventmateri"
)

type Event struct {
	app.Model
	ID                 app.NullInt64             `json:"id"                  db:"m.id"              gorm:"column:id;primaryKey"`
	AttachmentID       app.NullInt64             `json:"attachment.id"       db:"m.attachment_id"   gorm:"column:attachment_id"`
	AttachmentFilename app.NullString            `json:"attachment.filename" db:"a.filename"        gorm:"-"`
	AttachmentUrl      app.NullString            `json:"attachment.url"      db:"a.url"             gorm:"-"`
	Nama               app.NullText              `json:"nama"                db:"m.nama"            gorm:"column:nama"`
	UnitKerja          app.NullText              `json:"unit_kerja"          db:"m.unit_kerja"      gorm:"column:unit_kerja"`
	Uraian             app.NullText              `json:"uraian"              db:"m.uraian"          gorm:"column:uraian"`
	Lokasi             app.NullText              `json:"lokasi"              db:"m.lokasi"          gorm:"column:lokasi"`
	TanggalMulai       app.NullDate              `json:"tanggal_mulai"       db:"m.tanggal_mulai"   gorm:"column:tanggal_mulai"`
	TanggalSelesai     app.NullDate              `json:"tanggal_selesai"     db:"m.tanggal_selesai" gorm:"column:tanggal_selesai"`
	AksesKegiatan      app.NullString            `json:"akses_kegiatan"      db:"m.akses_kegiatan"  gorm:"column:akses_kegiatan"`
	CreatedBy          app.NullInt64             `json:"created_by.id"       db:"m.created_by"      gorm:"column:created_by"`
	CreatedByUsername  app.NullString            `json:"created_by.username" db:"cbuser.username"   gorm:"-"`
	UpdatedBy          app.NullInt64             `json:"updated_by.id"       db:"m.updated_by"       gorm:"column:updated_by"`
	UpdatedByUsername  app.NullString            `json:"updated_by.username" db:"ubuser.username"    gorm:"-"`
	DeletedBy          app.NullInt64             `json:"deleted_by.id"       db:"m.deleted_by"       gorm:"column:deleted_by"`
	DeletedByUsername  app.NullString            `json:"deleted_by.username" db:"dbuser.username"    gorm:"-"`
	CreatedAt          app.NullDateTime          `json:"created_at"          db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt          app.NullDateTime          `json:"updated_at"          db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt          app.NullDateTime          `json:"deleted_at"          db:"m.deleted_at,hide" gorm:"column:deleted_at"`
	Materi             []eventmateri.EventMateri `json:"materi"              db:"event.id={id}"     gorm:"-"`
}

func (Event) EndPoint() string {
	return "events"
}

func (Event) TableVersion() string {
	return "23.10.051152"
}

func (Event) TableName() string {
	return "t_events"
}

func (Event) TableAliasName() string {
	return "m"
}

func (m *Event) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_attachments", "a", []map[string]any{{"column1": "a.id", "column2": "m.attachment_id"}})
	m.AddRelation("left", "m_user", "cbuser", []map[string]any{{"column1": "cbuser.id_user", "column2": "m.created_by"}})
	m.AddRelation("left", "m_user", "ubuser", []map[string]any{{"column1": "ubuser.id_user", "column2": "m.updated_by"}})
	m.AddRelation("left", "m_user", "dbuser", []map[string]any{{"column1": "dbuser.id_user", "column2": "m.deleted_by"}})
	return m.Relations
}

func (m *Event) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *Event) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *Event) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *Event) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (Event) OpenAPISchemaName() string {
	return "Event"
}

func (m *Event) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type EventList struct {
	app.ListModel
}

func (EventList) OpenAPISchemaName() string {
	return "EventList"
}

func (p *EventList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&Event{})
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
