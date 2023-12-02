package totalsummary

import "github.com/maulanar/kms/app"

type TotalSummary struct {
	app.Model
	TotalDocuments    app.NullInt64 `json:"total_documents"    db:"COUNT(*)"                                    gorm:"-"`
	TotalContributors app.NullInt64 `json:"total_contributors" db:"(SELECT COUNT(*) FROM m_user)"               gorm:"-"`
	TotalThreads      app.NullInt64 `json:"total_threads"      db:"(SELECT COUNT(*) FROM t_forum)"              gorm:"-"`
	TotalHits         app.NullInt64 `json:"total_hits"         db:"(SELECT SUM(count_view) FROM t_pengetahuan)" gorm:"-"`
}

func (TotalSummary) EndPoint() string {
	return "total_summaries"
}

func (TotalSummary) TableVersion() string {
	return "28.06.291152"
}

func (TotalSummary) TableName() string {
	return "t_pengetahuan_dokumen"
}

func (TotalSummary) TableAliasName() string {
	return "m"
}

func (m *TotalSummary) GetRelations() map[string]map[string]any {

	return m.Relations
}

func (m *TotalSummary) GetFilters() []map[string]any {
	return m.Filters
}

func (m *TotalSummary) GetSorts() []map[string]any {
	return m.Sorts
}

func (m *TotalSummary) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *TotalSummary) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (TotalSummary) OpenAPISchemaName() string {
	return "TotalSummary"
}

func (m *TotalSummary) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type TotalSummaryList struct {
	app.ListModel
}

func (TotalSummaryList) OpenAPISchemaName() string {
	return "TotalSummaryList"
}

func (p *TotalSummaryList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&TotalSummary{})
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
