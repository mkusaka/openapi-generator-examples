package api

// This file is auto-generated, don't modify it manually

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

// NewRouter creates a new router for the spec and the given handlers.
// Swagger Petstore - OpenAPI 3.0
//
// This is a sample Pet Store Server based on the OpenAPI 3.0 specification.  You can find out more about
// Swagger at [http://swagger.io](http://swagger.io). In the third iteration of the pet store, we've switched to the design first approach!
// You can now help us improve the API whether it's by making changes to the definition itself or to the code.
// That way, with time, we can improve the API in general, and expose some of the new features in OAS3.
//
// Some useful links:
// - [The Pet Store repository](https://github.com/swagger-api/swagger-petstore)
// - [The source API definition for the Pet Store](https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml)
//
// 1.0.9
//
func NewRouter() http.Handler {

	r := chi.NewRouter()

	return r
}

func optionsHandlerFunc(allowedMethods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Allow", strings.Join(allowedMethods, ", "))
	}
}
