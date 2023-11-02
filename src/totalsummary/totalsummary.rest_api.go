package totalsummary

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"grest.dev/grest"

	"github.com/maulanar/kms/app"
)

// REST returns a *RESTAPIHandler.
func REST() *RESTAPIHandler {
	return &RESTAPIHandler{}
}

// RESTAPIHandler provides a convenient interface for TotalSummary REST API handler.
type RESTAPIHandler struct {
	UseCase UseCaseHandler
}

// injectDeps inject the dependencies of the TotalSummary REST API handler.
func (r *RESTAPIHandler) injectDeps(c *fiber.Ctx) error {
	ctx, ok := c.Locals(app.CtxKey).(*app.Ctx)
	if !ok {
		return app.Error().New(http.StatusInternalServerError, "ctx is not found")
	}
	r.UseCase = UseCase(*ctx, app.Query().Parse(c.OriginalURL()))
	return nil
}

// Get is the REST API handler for `GET /api/total_summaries`.
func (r *RESTAPIHandler) Get(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	res, err := r.UseCase.Get()
	if err != nil {
		return app.Error().Handler(c, err)
	}
	if r.UseCase.IsFlat() {
		return c.JSON(res)
	}
	return c.JSON(grest.NewJSON(res).ToStructured().Data)
}
