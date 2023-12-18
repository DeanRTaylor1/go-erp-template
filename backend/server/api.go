// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for UserRole.
const (
	UserRoleAdmin   UserRole = "admin"
	UserRoleManager UserRole = "manager"
	UserRoleUser    UserRole = "user"
)

// Defines values for UserStatus.
const (
	Active   UserStatus = "active"
	Inactive UserStatus = "inactive"
)

// User defines model for User.
type User struct {
	CreatedAt *time.Time          `json:"createdAt,omitempty"`
	Email     openapi_types.Email `json:"email"`
	FirstName *string             `json:"firstName,omitempty"`
	Id        *int64              `json:"id,omitempty"`
	LastName  *string             `json:"lastName,omitempty"`
	Password  string              `json:"password"`
	Role      *UserRole           `json:"role,omitempty"`
	Status    UserStatus          `json:"status"`
	UpdatedAt *time.Time          `json:"updatedAt,omitempty"`
	Username  string              `json:"username"`
}

// UserRole defines model for User.Role.
type UserRole string

// UserStatus defines model for User.Status.
type UserStatus string

// GetUsersParams defines parameters for GetUsers.
type GetUsersParams struct {
	// Offset Page number of the users list
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit Number of users per page
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`
}

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody = User

// PutUsersUserIdJSONRequestBody defines body for PutUsersUserId for application/json ContentType.
type PutUsersUserIdJSONRequestBody = User

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get a list of users
	// (GET /users)
	GetUsers(c *gin.Context, params GetUsersParams)
	// Create a new user
	// (POST /users)
	PostUsers(c *gin.Context)
	// Delete a user
	// (DELETE /users/{userId})
	DeleteUsersUserId(c *gin.Context, userId int)
	// Get a user by ID
	// (GET /users/{userId})
	GetUsersUserId(c *gin.Context, userId int)
	// Update a user
	// (PUT /users/{userId})
	PutUsersUserId(c *gin.Context, userId int)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetUsers operation middleware
func (siw *ServerInterfaceWrapper) GetUsers(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUsersParams

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", c.Request.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter offset: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", c.Request.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter limit: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUsers(c, params)
}

// PostUsers operation middleware
func (siw *ServerInterfaceWrapper) PostUsers(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostUsers(c)
}

// DeleteUsersUserId operation middleware
func (siw *ServerInterfaceWrapper) DeleteUsersUserId(c *gin.Context) {

	var err error

	// ------------- Path parameter "userId" -------------
	var userId int

	err = runtime.BindStyledParameter("simple", false, "userId", c.Param("userId"), &userId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter userId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteUsersUserId(c, userId)
}

// GetUsersUserId operation middleware
func (siw *ServerInterfaceWrapper) GetUsersUserId(c *gin.Context) {

	var err error

	// ------------- Path parameter "userId" -------------
	var userId int

	err = runtime.BindStyledParameter("simple", false, "userId", c.Param("userId"), &userId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter userId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUsersUserId(c, userId)
}

// PutUsersUserId operation middleware
func (siw *ServerInterfaceWrapper) PutUsersUserId(c *gin.Context) {

	var err error

	// ------------- Path parameter "userId" -------------
	var userId int

	err = runtime.BindStyledParameter("simple", false, "userId", c.Param("userId"), &userId)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter userId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutUsersUserId(c, userId)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/users", wrapper.GetUsers)
	router.POST(options.BaseURL+"/users", wrapper.PostUsers)
	router.DELETE(options.BaseURL+"/users/:userId", wrapper.DeleteUsersUserId)
	router.GET(options.BaseURL+"/users/:userId", wrapper.GetUsersUserId)
	router.PUT(options.BaseURL+"/users/:userId", wrapper.PutUsersUserId)
}
