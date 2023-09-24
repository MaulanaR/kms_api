package accesstoken

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

// RESTAPIHandler provides a convenient interface for AccessToken REST API handler.
type RESTAPIHandler struct {
	UseCase UseCaseHandler
}

// injectDeps inject the dependencies of the AccessToken REST API handler.
func (r *RESTAPIHandler) injectDeps(c *fiber.Ctx) error {
	ctx, ok := c.Locals(app.CtxKey).(*app.Ctx)
	if !ok {
		return app.Error().New(http.StatusInternalServerError, "ctx is not found")
	}
	r.UseCase = UseCase(*ctx, app.Query().Parse(c.OriginalURL()))
	return nil
}

// GetByID is the REST API handler for `GET /api/access_token/{id}`.
func (r *RESTAPIHandler) GetByID(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	res, err := r.UseCase.GetByID("")
	if err != nil {
		return app.Error().Handler(c, err)
	}
	if r.UseCase.IsFlat() {
		return c.JSON(res)
	}
	return c.JSON(grest.NewJSON(res).ToStructured().Data)
}

func (r *RESTAPIHandler) Login(c *fiber.Ctx) error {
	err := r.injectDeps(c)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	p := ParamCreate{}
	err = grest.NewJSON(c.Body()).ToFlat().Unmarshal(&p)
	if err != nil {
		return app.Error().Handler(c, app.Error().New(http.StatusBadRequest, err.Error()))
	}
	res, err := r.UseCase.Login(&p)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	if r.UseCase.Query.Get("is_skip_return") == "true" {
		return c.Status(http.StatusOK).JSON(map[string]any{"message": "Success"})
	}
	res, err = r.UseCase.GetByID(res.AccessToken.String)
	if err != nil {
		return app.Error().Handler(c, err)
	}
	if r.UseCase.IsFlat() {
		return c.Status(http.StatusOK).JSON(res)
	}
	return c.Status(http.StatusOK).JSON(grest.NewJSON(res).ToStructured().Data)
}
