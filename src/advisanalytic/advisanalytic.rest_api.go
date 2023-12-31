package advisanalytic

import (
	"bytes"
	"encoding/csv"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"grest.dev/grest"

	"github.com/maulanar/kms/app"
)

// REST returns a *RESTAPIHandler.
func REST() *RESTAPIHandler {
	return &RESTAPIHandler{}
}

// RESTAPIHandler provides a convenient interface for AdvisAnalytic REST API handler.
type RESTAPIHandler struct {
	UseCase UseCaseHandler
}

// injectDeps inject the dependencies of the AdvisAnalytic REST API handler.
func (r *RESTAPIHandler) injectDeps(c *fiber.Ctx) error {
	ctx, ok := c.Locals(app.CtxKey).(*app.Ctx)
	if !ok {
		return app.Error().New(http.StatusInternalServerError, "ctx is not found")
	}
	r.UseCase = UseCase(*ctx, app.Query().Parse(c.OriginalURL()))
	return nil
}

// GetByID is the REST API handler for `GET /api/advis_analytics/{id}`.
func (r *RESTAPIHandler) GetByID(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	res, err := r.UseCase.GetByID(c.Params("id"))
	if err != nil {
		return app.Error().Handler(c, err)
	}
	if r.UseCase.IsFlat() {
		return c.JSON(res)
	}
	return c.JSON(grest.NewJSON(res).ToStructured().Data)
}

// Get is the REST API handler for `GET /api/advis_analytics`.
func (r *RESTAPIHandler) Get(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	res, err := r.UseCase.Get()
	if err != nil {
		return app.Error().Handler(c, err)
	}
	res.SetLink(c)
	if r.UseCase.IsFlat() {
		return c.JSON(res)
	}
	return c.JSON(grest.NewJSON(res).ToStructured().Data)
}

// Create is the REST API handler for `POST /api/advis_analytics`.
func (r *RESTAPIHandler) Create(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	p := ParamCreate{}
	p.Ctx = r.UseCase.Ctx
	err = grest.NewJSON(c.Body()).ToFlat().Unmarshal(&p)
	if err != nil {
		return app.Error().Handler(c, app.Error().New(http.StatusBadRequest, err.Error()))
	}
	err = r.UseCase.Create(&p)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	if r.UseCase.Query.Get("is_skip_return") == "true" {
		return c.Status(http.StatusCreated).JSON(map[string]any{"message": "Success"})
	}
	res, err := r.UseCase.GetByID(strconv.Itoa(int(p.ID.Int64)))
	if err != nil {
		return app.Error().Handler(c, err)
	}
	if r.UseCase.IsFlat() {
		return c.Status(http.StatusCreated).JSON(res)
	}
	return c.Status(http.StatusCreated).JSON(grest.NewJSON(res).ToStructured().Data)
}

// UpdateByID is the REST API handler for `PUT /api/advis_analytics/{id}`.
func (r *RESTAPIHandler) UpdateByID(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	p := ParamUpdate{}
	p.Ctx = r.UseCase.Ctx
	err = grest.NewJSON(c.Body()).ToFlat().Unmarshal(&p)
	if err != nil {
		return app.Error().Handler(c, app.Error().New(http.StatusBadRequest, err.Error()))
	}
	err = r.UseCase.UpdateByID(c.Params("id"), &p)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	if r.UseCase.Query.Get("is_skip_return") == "true" {
		return c.JSON(map[string]any{"message": "Success"})
	}
	res, err := r.UseCase.GetByID(c.Params("id"))
	if err != nil {
		return app.Error().Handler(c, err)
	}
	if r.UseCase.IsFlat() {
		return c.JSON(res)
	}
	return c.JSON(grest.NewJSON(res).ToStructured().Data)
}

// PartiallyUpdateByID is the REST API handler for `PATCH /api/advis_analytics/{id}`.
func (r *RESTAPIHandler) PartiallyUpdateByID(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	p := ParamPartiallyUpdate{}
	p.Ctx = r.UseCase.Ctx
	err = grest.NewJSON(c.Body()).ToFlat().Unmarshal(&p)
	if err != nil {
		return app.Error().Handler(c, app.Error().New(http.StatusBadRequest, err.Error()))
	}
	err = r.UseCase.PartiallyUpdateByID(c.Params("id"), &p)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	if r.UseCase.Query.Get("is_skip_return") == "true" {
		return c.JSON(map[string]any{"message": "Success"})
	}
	res, err := r.UseCase.GetByID(c.Params("id"))
	if err != nil {
		return app.Error().Handler(c, err)
	}
	if r.UseCase.IsFlat() {
		return c.JSON(res)
	}
	return c.JSON(grest.NewJSON(res).ToStructured().Data)
}

// DeleteByID is the REST API handler for `DELETE /api/advis_analytics/{id}`.
func (r *RESTAPIHandler) DeleteByID(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	p := ParamDelete{}
	p.Ctx = r.UseCase.Ctx
	err = grest.NewJSON(c.Body()).ToFlat().Unmarshal(&p)
	if err != nil {
		return app.Error().Handler(c, app.Error().New(http.StatusBadRequest, err.Error()))
	}
	err = r.UseCase.DeleteByID(c.Params("id"), &p)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	res := map[string]any{
		"code": http.StatusOK,
		"message": r.UseCase.Ctx.Trans("deleted", map[string]string{
			"advis_analytics": p.EndPoint(),
			"id":              c.Params("id"),
		}),
	}
	return c.JSON(res)
}

func (r *RESTAPIHandler) DownloadTemplateCSV(c *fiber.Ctx) error {
	// Data untuk CSV
	data := [][]string{
		{"Judul 1", "Uraian 1", "http://example.com/img1.jpg", "http://example.com/file1.pdf", "http://example.com/dashboard1"},
		// Tambahkan data lain jika diperlukan
	}

	// Membuat file CSV di memori
	csvData := [][]string{{"Judul", "Uraian", "img_url", "file_url", "dashboard_url"}}
	csvData = append(csvData, data...)

	// Membuat file CSV
	file, err := os.Create("TemplateAdvisAnalytics.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	// Penulisan data ke file CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(csvData)
	if err != nil {
		return err
	}

	// Atur header untuk file CSV
	c.Set("Content-Disposition", "attachment; filename=TemplateAdvisAnalytics.csv")
	c.Set("Content-Type", "text/csv")

	// Menghapus file setelah dikirim
	defer os.Remove("TemplateAdvisAnalytics.csv")

	// Mengirim file CSV sebagai stream
	err = c.SendFile("TemplateAdvisAnalytics.csv")
	if err != nil {
		return err
	}

	return nil
}

func (r *RESTAPIHandler) UploadCSV(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["files"]
	if len(files) == 0 {
		res := map[string]any{
			"code":    http.StatusBadRequest,
			"message": "Tidak ada file yang di-upload",
		}
		return c.JSON(res)
	}

	file := files[0]
	fileContent, err := file.Open()
	if err != nil {
		return err
	}
	defer fileContent.Close()

	// Baca isi file CSV
	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, fileContent)
	if err != nil {
		return err
	}

	// Ubah isi buffer menjadi string
	csvString := buffer.String()

	// Baca sebagai CSV
	reader := csv.NewReader(strings.NewReader(csvString))
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// Lakukan sesuatu dengan data CSV, contohnya tampilkan di console
	for i, row := range records {
		if i != 0 {
			//save data
			p := ParamCreate{}
			p.Judul.Set(row[0])
			p.Uraian.Set(row[1])
			p.ImgUrl.Set(row[2])
			p.FileUrl.Set(row[3])
			p.DashboardUrl.Set(row[4])
			p.Ctx = r.UseCase.Ctx
			err = r.UseCase.Create(&p)
			if err != nil {
				return app.Error().Handler(c, err)
			}
		}
	}

	res := map[string]any{
		"code":    http.StatusOK,
		"message": "Data CSV berhasil di-Import",
	}
	return c.JSON(res)
}
