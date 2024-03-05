package bannercarousel

import "github.com/maulanar/kms/app"

type BannerCarousel struct {
	app.Model
	ID                app.NullInt64    `json:"id"                  db:"m.id"              gorm:"column:id;primaryKey"`
	JenisBanner       app.NullString   `json:"jenis_banner"        db:"m.jenis_banner"    gorm:"column:jenis_banner"`
	Order             app.NullInt64    `json:"order"               db:"m.order"           gorm:"column:order"`
	ResponseLink      app.NullText     `json:"response_link"       db:"m.response_link"   gorm:"column:response_link"`
	IsActive          app.NullBool     `json:"is_active"           db:"m.is_active"       gorm:"column:is_active"`
	GambarID          app.NullInt64    `json:"gambar.id"           db:"m.gambar_id"       gorm:"column:gambar_id"`
	GambarFilename    app.NullString   `json:"gambar.filename"     db:"a.filename"        gorm:"-"`
	GambarUrl         app.NullString   `json:"gambar.url"          db:"a.url"             gorm:"-"`
	CreatedBy         app.NullInt64    `json:"created_by.id"       db:"m.created_by"      gorm:"column:created_by"`
	CreatedByUsername app.NullString   `json:"created_by.username" db:"cbuser.username"   gorm:"-"`
	UpdatedBy         app.NullInt64    `json:"updated_by.id"       db:"m.updated_by"      gorm:"column:updated_by"`
	UpdatedByUsername app.NullString   `json:"updated_by.username" db:"ubuser.username"   gorm:"-"`
	DeletedBy         app.NullInt64    `json:"deleted_by.id"       db:"m.deleted_by"      gorm:"column:deleted_by"`
	DeletedByUsername app.NullString   `json:"deleted_by.username" db:"dbuser.username"   gorm:"-"`
	CreatedAt         app.NullDateTime `json:"created_at"          db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt         app.NullDateTime `json:"updated_at"          db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt         app.NullDateTime `json:"deleted_at"          db:"m.deleted_at,hide" gorm:"column:deleted_at"`
}

func (BannerCarousel) EndPoint() string {
	return "banner_carousel"
}

func (BannerCarousel) TableVersion() string {
	return "28.06.291152"
}

func (BannerCarousel) TableName() string {
	return "banner_carousel"
}

func (BannerCarousel) TableAliasName() string {
	return "m"
}

func (m *BannerCarousel) GetRelations() map[string]map[string]any {
	m.AddRelation("left", "m_user", "cbuser", []map[string]any{{"column1": "cbuser.id_user", "column2": "m.created_by"}})
	m.AddRelation("left", "m_user", "ubuser", []map[string]any{{"column1": "ubuser.id_user", "column2": "m.updated_by"}})
	m.AddRelation("left", "m_user", "dbuser", []map[string]any{{"column1": "dbuser.id_user", "column2": "m.deleted_by"}})
	m.AddRelation("left", "m_attachments", "a", []map[string]any{{"column1": "a.id", "column2": "m.gambar_id"}})

	return m.Relations
}

func (m *BannerCarousel) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

func (m *BannerCarousel) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

func (m *BannerCarousel) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *BannerCarousel) GetSchema() map[string]any {
	return m.SetSchema(m)
}

func (BannerCarousel) OpenAPISchemaName() string {
	return "BannerCarousel"
}

func (m *BannerCarousel) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

type BannerCarouselList struct {
	app.ListModel
}

func (BannerCarouselList) OpenAPISchemaName() string {
	return "BannerCarouselList"
}

func (p *BannerCarouselList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&BannerCarousel{})
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
