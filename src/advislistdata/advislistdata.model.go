package advislistdata

import "github.com/maulanar/kms/app"

type AdvisListData struct {
	app.Model
	ID                      app.NullInt64    `json:"id"                         db:"m.id"                    gorm:"column:id;primaryKey"`
	SubKategoriID           app.NullInt64    `json:"sub_kategori.id"            db:"m.ref_sub_kategori_id"   gorm:"column:ref_sub_kategori_id"`
	SubKategoriNama         app.NullString   `json:"sub_kategori.nama"          db:"skt.nama_sub_kategori"   gorm:"-"`
	SubKategoriKategoriID   app.NullInt64    `json:"sub_kategori.kategori.id"   db:"skt.ref_kategori_id"     gorm:"-"`
	SubKategoriKategoriNama app.NullString   `json:"sub_kategori.kategori.nama" db:"skt_kt.nama_kategori"    gorm:"-"`
	SumberDataID            app.NullInt64    `json:"sumber_data.id"             db:"m.ref_sumber_data_id"    gorm:"column:ref_sumber_data_id"`
	SumberDataNama          app.NullString   `json:"sumber_data.nama"           db:"smd.nama_sumber_data"    gorm:"-"`
	SumberDataSingkat       app.NullString   `json:"sumber_data.singkat"        db:"smd.singkat_sumber_data" gorm:"-"`
	SumberDataKeterangan    app.NullText     `json:"sumber_data.keterangan"     db:"smd.ket_sumber_data"     gorm:"-"`
	NamaData                app.NullString   `json:"nama_data"                  db:"m.nama_data"             gorm:"column:nama_data"`
	UrlData                 app.NullString   `json:"url_data"                   db:"m.url_data"              gorm:"column:url_data"`
	CreatedBy               app.NullString   `json:"created_by"                 db:"m.created_by"            gorm:"column:created_by"`
	UpdatedBy               app.NullString   `json:"updated_by"                 db:"m.updated_by"            gorm:"column:updated_by"`
	CreatedAt               app.NullDateTime `json:"created_at"                 db:"m.created_at"            gorm:"column:created_at"`
	UpdatedAt               app.NullDateTime `json:"updated_at"                 db:"m.updated_at"            gorm:"column:updated_at"`
	DeletedAt               app.NullDateTime `json:"deleted_at"                 db:"m.deleted_at,hide"       gorm:"column:deleted_at"`
	// CreatedByUsername       app.NullString   `json:"created_by.username"        db:"cbuser.username"         gorm:"-"`
	// UpdatedByUsername       app.NullString   `json:"updated_by.username"        db:"ubuser.username"         gorm:"-"`
	// DeletedBy         app.NullInt64    `json:"deleted_by.id"              db:"m.deleted_by"            gorm:"column:deleted_by"`
	// DeletedByUsername app.NullString   `json:"deleted_by.username"        db:"dbuser.username"         gorm:"-"`
}

func (AdvisListData) EndPoint() string {
	return "advis_list_data"
}

func (AdvisListData) TableVersion() string {
	return "28.06.291152"
}

func (AdvisListData) TableName() string {
	return "t_list_data"
}

func (AdvisListData) TableAliasName() string {
	return "m"
}

func (m *AdvisListData) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "ref_sub_kategoris", "skt", []map[string]any{{"column1": "skt.id", "column2": "m.ref_sub_kategori_id"}})
	m.AddRelation("left", "ref_kategoris", "skt_kt", []map[string]any{{"column1": "skt_kt.id", "column2": "skt.ref_kategori_id"}})
	m.AddRelation("left", "ref_sumber_data", "smd", []map[string]any{{"column1": "smd.id", "column2": "m.ref_sumber_data_id"}})

	// m.AddRelation("left", "m_user", "cbuser", []map[string]any{{"column1": "cbuser.id_user", "column2": "m.created_by"}})
	// m.AddRelation("left", "m_user", "ubuser", []map[string]any{{"column1": "ubuser.id_user", "column2": "m.updated_by"}})
	// m.AddRelation("left", "m_user", "dbuser", []map[string]any{{"column1": "dbuser.id_user", "column2": "m.deleted_by"}})

	return m.Relations
}

func (m *AdvisListData) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *AdvisListData) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *AdvisListData) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *AdvisListData) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (AdvisListData) OpenAPISchemaName() string {
	return "AdvisListData"
}

func (m *AdvisListData) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type AdvisListDataList struct {
	app.ListModel
}

func (AdvisListDataList) OpenAPISchemaName() string {
	return "AdvisListDataList"
}

func (p *AdvisListDataList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&AdvisListData{})
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
