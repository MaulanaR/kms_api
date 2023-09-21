package eventmateri

import "github.com/maulanar/kms/app"

type EventMateri struct {
	app.Model
	ID                  app.NullInt64      `json:"id"                   db:"m.id"                   gorm:"column:id;primaryKey"`
	EventID             app.NullInt64      `json:"event.id"             db:"m.event_id"             gorm:"column:event_id"`
	NamaPengarang       app.NullText       `json:"nama_pengarang"       db:"m.nama_pengarang"       gorm:"column:nama_pengarang"`
	NomorDokumen        app.NullText       `json:"nomor_dokumen"        db:"m.nomor_dokumen"        gorm:"column:nomor_dokumen"`
	TanggalDokumen      app.NullDate       `json:"tanggal_dokumen"      db:"m.tanggal_dokumen"      gorm:"column:tanggal_dokumen"`
	JudulDokumen        app.NullText       `json:"judul_dokumen"        db:"m.judul_dokumen"        gorm:"column:judul_dokumen"`
	Penerbit            app.NullText       `json:"penerbit"             db:"m.penerbit"             gorm:"column:penerbit"`
	KelompokDokumen     app.NullText       `json:"kelompok_dokumen"     db:"m.kelompok_dokumen"     gorm:"column:kelompok_dokumen"`
	KelompokPengetahuan app.NullText       `json:"kelompok_pengetahuan" db:"m.kelompok_pengetahuan" gorm:"column:kelompok_pengetahuan"`
	IsiDokumen          app.NullText       `json:"isi_dokumen"          db:"m.isi_dokumen"          gorm:"column:isi_dokumen"`
	Dokumen             []MateriAttachment `json:"dokumen"              db:"event_materi.id={id}"   gorm:"-"`
	CreatedBy           app.NullInt64      `json:"created_by.id"        db:"m.created_by"           gorm:"column:created_by"`
	CreatedByUsername   app.NullString     `json:"created_by.username"  db:"cbuser.username"        gorm:"-"`
	CreatedAt           app.NullDateTime   `json:"created_at"           db:"m.created_at"           gorm:"column:created_at"`
	UpdatedAt           app.NullDateTime   `json:"updated_at"           db:"m.updated_at"           gorm:"column:updated_at"`
	DeletedAt           app.NullDateTime   `json:"deleted_at"           db:"m.deleted_at,hide"      gorm:"column:deleted_at"`
}

func (EventMateri) EndPoint() string {
	return "event_materi"
}

func (EventMateri) TableVersion() string {
	return "28.08.291152"
}

func (EventMateri) TableName() string {
	return "t_event_materi"
}

func (EventMateri) TableAliasName() string {
	return "m"
}

func (m *EventMateri) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_user", "cbuser", []map[string]any{{"column1": "cbuser.id_user", "column2": "m.created_by"}})
	return m.Relations
}

func (m *EventMateri) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *EventMateri) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *EventMateri) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *EventMateri) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (EventMateri) OpenAPISchemaName() string {
	return "EventMateri"
}

func (m *EventMateri) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type EventMateriList struct {
	app.ListModel
}

func (EventMateriList) OpenAPISchemaName() string {
	return "EventMateriList"
}

func (p *EventMateriList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&EventMateri{})
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

type MateriAttachment struct {
	app.Model
	ID              app.NullInt64    `json:"id"                          db:"mea.id"              gorm:"column:id;primaryKey;autoIncrement"`
	EventMateriID   app.NullInt64    `json:"event_materi.id"             db:"mea.event_materi_id" gorm:"column:event_materi_id"`
	AttachmentID    app.NullInt64    `json:"attachment.id"               db:"mea.attachment_id"   gorm:"column:attachment_id"`
	Filename        app.NullString   `json:"attachment.filename"         db:"a.filename"          gorm:"-"`
	Size            app.NullInt64    `json:"attachment.size"             db:"a.size"              gorm:"-"`
	Extension       app.NullString   `json:"attachment.extension"        db:"a.extension"         gorm:"-"`
	StorageLocation app.NullString   `json:"attachment.storage_location" db:"a.storage_location"  gorm:"-"`
	Url             app.NullString   `json:"attachment.url"              db:"a.url"               gorm:"-"`
	CreatedAt       app.NullDateTime `json:"created_at"                  db:"mea.created_at"      gorm:"column:created_at"`
	UpdatedAt       app.NullDateTime `json:"updated_at"                  db:"mea.updated_at"      gorm:"column:updated_at"`
	DeletedAt       app.NullDateTime `json:"deleted_at"                  db:"mea.deleted_at,hide" gorm:"column:deleted_at"`
}

func (MateriAttachment) TableVersion() string {
	return "28.07.291152"
}

func (MateriAttachment) TableName() string {
	return "m_event_materi_attachments"
}

func (MateriAttachment) TableAliasName() string {
	return "mea"
}

func (m *MateriAttachment) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_attachments", "a", []map[string]any{{"column1": "a.id", "column2": "mea.attachment_id"}})
	return m.Relations
}

func (m *MateriAttachment) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "mea.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *MateriAttachment) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "mea.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *MateriAttachment) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *MateriAttachment) GetSchema() map[string]any {
	return m.SetSchema(m)
}

type Event struct {
	app.Model
	ID                 app.NullInt64    `json:"id"                  db:"m.id"              gorm:"column:id;primaryKey"`
	AttachmentID       app.NullInt64    `json:"attachment.id"       db:"m.attachment_id"   gorm:"column:attachment_id"`
	AttachmentFilename app.NullString   `json:"attachment.filename" db:"a.filename"        gorm:"-"`
	AttachmentUrl      app.NullString   `json:"attachment.url"      db:"a.url"             gorm:"-"`
	Nama               app.NullText     `json:"nama"                db:"m.nama"            gorm:"column:nama"`
	UnitKerja          app.NullText     `json:"unit_kerja"          db:"m.unit_kerja"      gorm:"column:unit_kerja"`
	Uraian             app.NullText     `json:"uraian"              db:"m.uraian"          gorm:"column:uraian"`
	Lokasi             app.NullText     `json:"lokasi"              db:"m.lokasi"          gorm:"column:lokasi"`
	TanggalMulai       app.NullDate     `json:"tanggal_mulai"       db:"m.tanggal_mulai"   gorm:"column:tanggal_mulai"`
	TanggalSelesai     app.NullDate     `json:"tanggal_selesai"     db:"m.tanggal_selesai" gorm:"column:tanggal_selesai"`
	AksesKegiatan      app.NullString   `json:"akses_kegiatan"      db:"m.akses_kegiatan"  gorm:"column:akses_kegiatan"`
	CreatedBy          app.NullInt64    `json:"created_by.id"       db:"m.created_by"      gorm:"column:created_by"`
	CreatedByUsername  app.NullString   `json:"created_by.username" db:"cbuser.username"   gorm:"-"`
	CreatedAt          app.NullDateTime `json:"created_at"          db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt          app.NullDateTime `json:"updated_at"          db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt          app.NullDateTime `json:"deleted_at"          db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

func (Event) EndPoint() string {
	return "events"
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
