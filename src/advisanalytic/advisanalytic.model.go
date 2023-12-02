package advisanalytic

import "github.com/maulanar/kms/app"

// LIST TABLE ADVIS
// `ref_jenis_visuals`
// `ref_kategoris`
// `ref_sub_kategoris`
// `ref_sumber_data`
// `t_analytics`
// `t_list_data`
// `t_visuals`

type AdvisAnalytic struct {
	app.Model
	ID           app.NullInt64    `json:"id"            db:"m.id"              gorm:"column:id;primaryKey"`
	Judul        app.NullString   `json:"judul"         db:"m.judul"           gorm:"column:judul"`
	Uraian       app.NullText     `json:"uraian"        db:"m.uraian"          gorm:"column:uraian"`
	ImgUrl       app.NullString   `json:"img_url"       db:"m.img_url"         gorm:"column:img_url"`
	FileUrl      app.NullString   `json:"file_url"      db:"m.file_url"        gorm:"column:file_url"`
	DashboardUrl app.NullString   `json:"dashboard_url" db:"m.dashboard_url"   gorm:"column:dashboard_url"`
	CreatedBy    app.NullString   `json:"created_by"    db:"m.created_by"      gorm:"column:created_by"`
	UpdatedBy    app.NullString   `json:"updated_by"    db:"m.updated_by"      gorm:"column:updated_by"`
	CreatedAt    app.NullDateTime `json:"created_at"    db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt    app.NullDateTime `json:"updated_at"    db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt    app.NullDateTime `json:"deleted_at"    db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

func (AdvisAnalytic) EndPoint() string {
	return "advis_analytics"
}

func (AdvisAnalytic) TableVersion() string {
	return "28.06.291152"
}

func (AdvisAnalytic) TableName() string {
	return "t_analytics"
}

func (AdvisAnalytic) TableAliasName() string {
	return "m"
}

func (m *AdvisAnalytic) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_user", "cbuser", []map[string]any{{"column1": "cbuser.id_user", "column2": "m.created_by"}})
	m.AddRelation("left", "m_user", "ubuser", []map[string]any{{"column1": "ubuser.id_user", "column2": "m.updated_by"}})
	m.AddRelation("left", "m_user", "dbuser", []map[string]any{{"column1": "dbuser.id_user", "column2": "m.deleted_by"}})

	return m.Relations
}

func (m *AdvisAnalytic) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *AdvisAnalytic) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *AdvisAnalytic) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *AdvisAnalytic) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (AdvisAnalytic) OpenAPISchemaName() string {
	return "AdvisAnalytic"
}

func (m *AdvisAnalytic) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type AdvisAnalyticList struct {
	app.ListModel
}

func (AdvisAnalyticList) OpenAPISchemaName() string {
	return "AdvisAnalyticList"
}

func (p *AdvisAnalyticList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&AdvisAnalytic{})
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
