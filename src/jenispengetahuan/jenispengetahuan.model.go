package jenispengetahuan

import "github.com/maulanar/kms/app"

// JenisPengetahuan is the main model of JenisPengetahuan data. It provides a convenient interface for app.ModelInterface
type JenisPengetahuan struct {
	app.Model
	ID        app.NullInt64    `json:"id"         db:"m.id_jenis_pengetahuan"   gorm:"column:id_jenis_pengetahuan;primaryKey"`
	Nama      app.NullText     `json:"nama"       db:"m.nama_jenis_pengetahuan" gorm:"column:nama_jenis_pengetahuan"`
	CreatedAt app.NullDateTime `json:"created_at" db:"m.created_at"             gorm:"column:created_at"`
	UpdatedAt app.NullDateTime `json:"updated_at" db:"m.updated_at"             gorm:"column:updated_at"`
	DeletedAt app.NullDateTime `json:"deleted_at" db:"m.deleted_at,hide"        gorm:"column:deleted_at"`
}

// EndPoint returns the JenisPengetahuan end point, it used for cache key, etc.
func (JenisPengetahuan) EndPoint() string {
	return "jenis_pengetahuan"
}

// TableVersion returns the versions of the JenisPengetahuan table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (JenisPengetahuan) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the JenisPengetahuan table in the database.
func (JenisPengetahuan) TableName() string {
	return "m_jenis_pengetahuan"
}

// TableAliasName returns the table alias name of the JenisPengetahuan table, used for querying.
func (JenisPengetahuan) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the JenisPengetahuan data in the database, used for querying.
func (m *JenisPengetahuan) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the JenisPengetahuan data in the database, used for querying.
func (m *JenisPengetahuan) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the JenisPengetahuan data in the database, used for querying.
func (m *JenisPengetahuan) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the JenisPengetahuan data in the database, used for querying.
func (m *JenisPengetahuan) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the JenisPengetahuan schema, used for querying.
func (m *JenisPengetahuan) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the JenisPengetahuan schema in the open api documentation.
func (JenisPengetahuan) OpenAPISchemaName() string {
	return "JenisPengetahuan"
}

// GetOpenAPISchema returns the Open API Schema of the JenisPengetahuan in the open api documentation.
func (m *JenisPengetahuan) GetOpenAPISchema() map[string]any {
	return m.SetOpenAPISchema(m)
}

func SeederData() error {
	dataSeeder := []JenisPengetahuan{
		{
			ID:   app.NewNullInt64(1),
			Nama: app.NewNullText("Tugas (Panduan Penugasan)"),
		},
		{
			ID:   app.NewNullInt64(2),
			Nama: app.NewNullText("KIAT"),
		},
		{
			ID:   app.NewNullInt64(3),
			Nama: app.NewNullText("Kapitalisasi / Analytic Today"),
		},
		{
			ID:   app.NewNullInt64(4),
			Nama: app.NewNullText("Resensi"),
		},
		{
			ID:   app.NewNullInt64(5),
			Nama: app.NewNullText("Aksi Perubahan"),
		},
		{
			ID:   app.NewNullInt64(6),
			Nama: app.NewNullText("PKS (Pelatihan Kantor Sendiri)"),
		},
		{
			ID:   app.NewNullInt64(7),
			Nama: app.NewNullText("Karya Tulis"),
		},
		{
			ID:   app.NewNullInt64(8),
			Nama: app.NewNullText("Newsletter LC"),
		},
	}
	tx, err := app.DB().Conn("main")
	if err != nil {
		return err
	}
	err = tx.Create(&dataSeeder).Error
	if err != nil {
		return err
	}

	return nil
}

type JenisPengetahuanList struct {
	app.ListModel
}

// OpenAPISchemaName returns the name of the JenisPengetahuanList schema in the open api documentation.
func (JenisPengetahuanList) OpenAPISchemaName() string {
	return "JenisPengetahuanList"
}

// GetOpenAPISchema returns the Open API Schema of the JenisPengetahuanList in the open api documentation.
func (p *JenisPengetahuanList) GetOpenAPISchema() map[string]any {
	return p.SetOpenAPISchema(&JenisPengetahuan{})
}

// ParamCreate is the expected parameters for create a new JenisPengetahuan data.
type ParamCreate struct {
	UseCaseHandler
	Nama app.NullText `json:"nama" db:"m.nama_jenis_pengetahuan" gorm:"column:nama_jenis_pengetahuan" validate:"required"`
}

// ParamUpdate is the expected parameters for update the JenisPengetahuan data.
type ParamUpdate struct {
	UseCaseHandler
}

// ParamPartiallyUpdate is the expected parameters for partially update the JenisPengetahuan data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
}

// ParamDelete is the expected parameters for delete the JenisPengetahuan data.
type ParamDelete struct {
	UseCaseHandler
}
