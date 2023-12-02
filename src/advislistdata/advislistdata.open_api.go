package advislistdata

import "github.com/maulanar/kms/app"

// OpenAPI is constructor for *openAPI, to autogenerate open api document.
func OpenAPI() *OpenAPIOperation {
	return &OpenAPIOperation{}
}

// OpenAPIOperation embed from app.OpenAPIOperation for simplicity, used for autogenerate open api document.
type OpenAPIOperation struct {
	app.OpenAPIOperation
}

// Base is common detail of advis_list_data open api document component.
func (o *OpenAPIOperation) Base() {
	o.Tags = []string{"AdvisListData"}
	o.HeaderParams = []map[string]any{{"$ref": "#/components/parameters/headerParam.Accept-Language"}}
	o.Responses = map[string]map[string]any{
		"200": {
			"description": "Success",
			"content":     map[string]any{"application/json": &AdvisListData{}}, // will auto create schema $ref: '#/components/schemas/AdvisListData' if not exists
		},
		"400": app.OpenAPIError().BadRequest(),
		"401": app.OpenAPIError().Unauthorized(),
		"403": app.OpenAPIError().Forbidden(),
	}
	o.Securities = []map[string][]string{}
}

// Get is detail of `GET /api/v3/advis_list_data` open api document component.
func (o *OpenAPIOperation) Get() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o // skip for efficiency
	}

	o.Base()
	o.Summary = "Get AdvisListData"
	o.Description = "Use this method to get list of AdvisListData"
	o.QueryParams = []map[string]any{{"$ref": "#/components/parameters/queryParam.Any"}}
	o.Responses = map[string]map[string]any{
		"200": {
			"description": "Success",
			"content":     map[string]any{"application/json": &AdvisListDataList{}}, // will auto create schema $ref: '#/components/schemas/AdvisListData.List' if not exists
		},
		"400": app.OpenAPIError().BadRequest(),
		"401": app.OpenAPIError().Unauthorized(),
		"403": app.OpenAPIError().Forbidden(),
	}
	return o
}

// GetByID is detail of `GET /api/v3/advis_list_data/{id}` open api document component.
func (o *OpenAPIOperation) GetByID() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o // skip for efficiency
	}

	o.Base()
	o.Summary = "Get AdvisListData By ID"
	o.Description = "Use this method to get AdvisListData by id"
	o.PathParams = []map[string]any{{"$ref": "#/components/parameters/pathParam.ID"}}
	return o
}

// Create is detail of `POST /api/v3/advis_list_data` open api document component.
func (o *OpenAPIOperation) Create() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o // skip for efficiency
	}

	o.Base()
	o.Summary = "Create AdvisListData"
	o.Description = "Use this method to create AdvisListData"
	o.Body = map[string]any{"application/json": &ParamCreate{}}
	return o
}

// UpdateByID is detail of `PUT /api/v3/advis_list_data/{id}` open api document component.
func (o *OpenAPIOperation) UpdateByID() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o // skip for efficiency
	}

	o.Base()
	o.Summary = "Update AdvisListData By ID"
	o.Description = "Use this method to update AdvisListData by id"
	o.PathParams = []map[string]any{{"$ref": "#/components/parameters/pathParam.ID"}}
	o.Body = map[string]any{"application/json": &ParamUpdate{}}
	return o
}

// PartiallyUpdateByID is detail of `PATCH /api/v3/advis_list_data/{id}` open api document component.
func (o *OpenAPIOperation) PartiallyUpdateByID() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o // skip for efficiency
	}

	o.Base()
	o.Summary = "Partially Update AdvisListData By ID"
	o.Description = "Use this method to partially update AdvisListData by id"
	o.PathParams = []map[string]any{{"$ref": "#/components/parameters/pathParam.ID"}}
	o.Body = map[string]any{"application/json": &ParamPartiallyUpdate{}}
	return o
}

// DeleteByID is detail of `DELETE /api/v3/advis_list_data/{id}` open api document component.
func (o *OpenAPIOperation) DeleteByID() *OpenAPIOperation {
	if !app.IS_GENERATE_OPEN_API_DOC {
		return o // skip for efficiency
	}

	o.Base()
	o.Summary = "Delete AdvisListData By ID"
	o.Description = "Use this method to delete AdvisListData by id"
	o.PathParams = []map[string]any{{"$ref": "#/components/parameters/pathParam.ID"}}
	o.Body = map[string]any{"application/json": &ParamDelete{}}
	return o
}
