/*
 * Estrava
 *
 * Cool Estrava tings
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	//"github.com/gorilla/mux"
)

// UsersApiController binds http requests to an api service and writes the service results to the http response
type UsersApiController struct {
	service UsersApiServicer
	errorHandler ErrorHandler
}

// UsersApiOption for how the controller is set up.
type UsersApiOption func(*UsersApiController)

// WithUsersApiErrorHandler inject ErrorHandler into controller
func WithUsersApiErrorHandler(h ErrorHandler) UsersApiOption {
	return func(c *UsersApiController) {
		c.errorHandler = h
	}
}

// NewUsersApiController creates a default api controller
func NewUsersApiController(s UsersApiServicer, opts ...UsersApiOption) Router {
	controller := &UsersApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the UsersApiController
func (c *UsersApiController) Routes() Routes {
	return Routes{ 
		{
			"UserPost",
			strings.ToUpper("Post"),
			"/user",
			c.UserPost,
		},
	}
}

// UserPost - 
func (c *UsersApiController) UserPost(w http.ResponseWriter, r *http.Request) {
	userPostRequestParam := UserPostRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&userPostRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertUserPostRequestRequired(userPostRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UserPost(r.Context(), userPostRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
