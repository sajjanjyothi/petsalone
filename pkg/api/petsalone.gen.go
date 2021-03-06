// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/gin-gonic/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns all pets
	// (GET /api/pets)
	GetAllPets(c *gin.Context)
	// Creates a new missing pet
	// (POST /api/pets)
	AddPet(c *gin.Context)
	// Returns a specific pet type
	// (GET /api/pets/{type})
	FindPetByType(c *gin.Context, pType string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(c *gin.Context)

// GetAllPets operation middleware
func (siw *ServerInterfaceWrapper) GetAllPets(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetAllPets(c)
}

// AddPet operation middleware
func (siw *ServerInterfaceWrapper) AddPet(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.AddPet(c)
}

// FindPetByType operation middleware
func (siw *ServerInterfaceWrapper) FindPetByType(c *gin.Context) {

	var err error

	// ------------- Path parameter "type" -------------
	var pType string

	err = runtime.BindStyledParameter("simple", false, "type", c.Param("type"), &pType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter type: %s", err)})
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.FindPetByType(c, pType)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	router.GET(options.BaseURL+"/api/pets", wrapper.GetAllPets)

	router.POST(options.BaseURL+"/api/pets", wrapper.AddPet)

	router.GET(options.BaseURL+"/api/pets/:type", wrapper.FindPetByType)

	return router
}

