package testutil

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/pkg/errors"
)

// TestRequest is used to validate whether the handler returns values based on openapi.yaml
func TestRequest(handler http.HandlerFunc, httpReq *http.Request) error {
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}
	// 長すぎて草
	doc, err := loader.LoadFromFile("../../../../openapi.yaml")
	if err != nil {
		panic(err)
	}
	// Validate document
	_ = doc.Validate(ctx)
	router, err := gorillamux.NewRouter(doc)
	if err != nil {
		panic(err)
	}

	// Find route
	route, pathParams, err := router.FindRoute(httpReq)
	if err != nil {
		panic(err)
	}

	// Validate request
	requestValidationInput := &openapi3filter.RequestValidationInput{
		Request:    httpReq,
		PathParams: pathParams,
		Route:      route,
	}
	err = openapi3filter.ValidateRequest(ctx, requestValidationInput)
	if err != nil {
		return errors.Wrap(err, "failed request validation")
	}

	// Handle that request
	rec := httptest.NewRecorder()
	handler(rec, httpReq)
	response := *rec.Result()
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// Validate response
	responseValidationInput := &openapi3filter.ResponseValidationInput{
		RequestValidationInput: requestValidationInput,
		Status:                 response.StatusCode,
		Header:                 response.Header,
	}
	responseValidationInput.SetBodyBytes(body)

	return errors.Wrap(openapi3filter.ValidateResponse(ctx, responseValidationInput), "failed response validation")
}
