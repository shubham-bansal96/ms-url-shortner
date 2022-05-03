// Package classification awesome.
//
// Documentation of our awesome API.
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//     Host: localhost:4232
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - basic
//
//    SecurityDefinitions:
//    basic:
//      type: basic
//
// swagger:meta
package docs

import "github.com/ms-url-shortner/app/model"

// swagger:route POST /ms-url-shortner/getshorturl shortURL getShortURL
// return the short url
// responses:
//   200: shortURLResponse

// This text will appear as description of your response body.
// swagger:parameters getShortURL
type getShortURLWrapper struct {
	// in:body
	Body model.URLDTO
	// in:header
	// required
	Authorization string `json:"Authorization"`
}

// swagger:route GET /ms-url-shortner/ping ping getPing
// return the ping from application
// responses:
//   200: Success
// swagger:parameters getPing
type getPingWrapper struct {
	// in:path
	UserID string `json:"user_id"`
}

// swagger:response shortURLResponse
type getShortURLParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body model.URLDTO
}

// swagger:response Success
type Success struct {
	// in:body
	Success string `json:"success"`
}
