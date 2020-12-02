// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDivocPortalAPIAPI creates a new DivocPortalAPI instance
func NewDivocPortalAPIAPI(spec *loads.Document) *DivocPortalAPIAPI {
	return &DivocPortalAPIAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,

		JSONProducer: runtime.JSONProducer(),

		PostEnrollmentsHandler: PostEnrollmentsHandlerFunc(func(params PostEnrollmentsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation PostEnrollments has not yet been implemented")
		}),
		PostFacilitiesHandler: PostFacilitiesHandlerFunc(func(params PostFacilitiesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation PostFacilities has not yet been implemented")
		}),
		PostVaccinatorsHandler: PostVaccinatorsHandlerFunc(func(params PostVaccinatorsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation PostVaccinators has not yet been implemented")
		}),
		CreateFacilityUsersHandler: CreateFacilityUsersHandlerFunc(func(params CreateFacilityUsersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation CreateFacilityUsers has not yet been implemented")
		}),
		CreateMedicineHandler: CreateMedicineHandlerFunc(func(params CreateMedicineParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation CreateMedicine has not yet been implemented")
		}),
		CreateProgramHandler: CreateProgramHandlerFunc(func(params CreateProgramParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation CreateProgram has not yet been implemented")
		}),
		GetEnrollmentsHandler: GetEnrollmentsHandlerFunc(func(params GetEnrollmentsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetEnrollments has not yet been implemented")
		}),
		GetFacilitiesHandler: GetFacilitiesHandlerFunc(func(params GetFacilitiesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetFacilities has not yet been implemented")
		}),
		GetFacilityGroupsHandler: GetFacilityGroupsHandlerFunc(func(params GetFacilityGroupsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetFacilityGroups has not yet been implemented")
		}),
		GetFacilityUsersHandler: GetFacilityUsersHandlerFunc(func(params GetFacilityUsersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetFacilityUsers has not yet been implemented")
		}),
		GetMedicinesHandler: GetMedicinesHandlerFunc(func(params GetMedicinesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetMedicines has not yet been implemented")
		}),
		GetProgramsHandler: GetProgramsHandlerFunc(func(params GetProgramsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetPrograms has not yet been implemented")
		}),
		GetVaccinatorsHandler: GetVaccinatorsHandlerFunc(func(params GetVaccinatorsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetVaccinators has not yet been implemented")
		}),

		HasRoleAuth: func(token string, scopes []string) (interface{}, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (hasRole) has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*DivocPortalAPIAPI Digital infra for vaccination certificates */
type DivocPortalAPIAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for the following mime types:
	//   - multipart/form-data
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// HasRoleAuth registers a function that takes an access token and a collection of required scopes and returns a principal
	// it performs authentication based on an oauth2 bearer token provided in the request
	HasRoleAuth func(string, []string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// PostEnrollmentsHandler sets the operation handler for the post enrollments operation
	PostEnrollmentsHandler PostEnrollmentsHandler
	// PostFacilitiesHandler sets the operation handler for the post facilities operation
	PostFacilitiesHandler PostFacilitiesHandler
	// PostVaccinatorsHandler sets the operation handler for the post vaccinators operation
	PostVaccinatorsHandler PostVaccinatorsHandler
	// CreateFacilityUsersHandler sets the operation handler for the create facility users operation
	CreateFacilityUsersHandler CreateFacilityUsersHandler
	// CreateMedicineHandler sets the operation handler for the create medicine operation
	CreateMedicineHandler CreateMedicineHandler
	// CreateProgramHandler sets the operation handler for the create program operation
	CreateProgramHandler CreateProgramHandler
	// GetEnrollmentsHandler sets the operation handler for the get enrollments operation
	GetEnrollmentsHandler GetEnrollmentsHandler
	// GetFacilitiesHandler sets the operation handler for the get facilities operation
	GetFacilitiesHandler GetFacilitiesHandler
	// GetFacilityGroupsHandler sets the operation handler for the get facility groups operation
	GetFacilityGroupsHandler GetFacilityGroupsHandler
	// GetFacilityUsersHandler sets the operation handler for the get facility users operation
	GetFacilityUsersHandler GetFacilityUsersHandler
	// GetMedicinesHandler sets the operation handler for the get medicines operation
	GetMedicinesHandler GetMedicinesHandler
	// GetProgramsHandler sets the operation handler for the get programs operation
	GetProgramsHandler GetProgramsHandler
	// GetVaccinatorsHandler sets the operation handler for the get vaccinators operation
	GetVaccinatorsHandler GetVaccinatorsHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *DivocPortalAPIAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *DivocPortalAPIAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *DivocPortalAPIAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *DivocPortalAPIAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *DivocPortalAPIAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *DivocPortalAPIAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *DivocPortalAPIAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *DivocPortalAPIAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *DivocPortalAPIAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the DivocPortalAPIAPI
func (o *DivocPortalAPIAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}
	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.HasRoleAuth == nil {
		unregistered = append(unregistered, "HasRoleAuth")
	}

	if o.PostEnrollmentsHandler == nil {
		unregistered = append(unregistered, "PostEnrollmentsHandler")
	}
	if o.PostFacilitiesHandler == nil {
		unregistered = append(unregistered, "PostFacilitiesHandler")
	}
	if o.PostVaccinatorsHandler == nil {
		unregistered = append(unregistered, "PostVaccinatorsHandler")
	}
	if o.CreateFacilityUsersHandler == nil {
		unregistered = append(unregistered, "CreateFacilityUsersHandler")
	}
	if o.CreateMedicineHandler == nil {
		unregistered = append(unregistered, "CreateMedicineHandler")
	}
	if o.CreateProgramHandler == nil {
		unregistered = append(unregistered, "CreateProgramHandler")
	}
	if o.GetEnrollmentsHandler == nil {
		unregistered = append(unregistered, "GetEnrollmentsHandler")
	}
	if o.GetFacilitiesHandler == nil {
		unregistered = append(unregistered, "GetFacilitiesHandler")
	}
	if o.GetFacilityGroupsHandler == nil {
		unregistered = append(unregistered, "GetFacilityGroupsHandler")
	}
	if o.GetFacilityUsersHandler == nil {
		unregistered = append(unregistered, "GetFacilityUsersHandler")
	}
	if o.GetMedicinesHandler == nil {
		unregistered = append(unregistered, "GetMedicinesHandler")
	}
	if o.GetProgramsHandler == nil {
		unregistered = append(unregistered, "GetProgramsHandler")
	}
	if o.GetVaccinatorsHandler == nil {
		unregistered = append(unregistered, "GetVaccinatorsHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *DivocPortalAPIAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *DivocPortalAPIAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "hasRole":
			result[name] = o.BearerAuthenticator(name, o.HasRoleAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *DivocPortalAPIAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *DivocPortalAPIAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *DivocPortalAPIAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *DivocPortalAPIAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the divoc portal API API
func (o *DivocPortalAPIAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *DivocPortalAPIAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/enrollments"] = NewPostEnrollments(o.context, o.PostEnrollmentsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/facilities"] = NewPostFacilities(o.context, o.PostFacilitiesHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/vaccinators"] = NewPostVaccinators(o.context, o.PostVaccinatorsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/facility/users"] = NewCreateFacilityUsers(o.context, o.CreateFacilityUsersHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/medicines"] = NewCreateMedicine(o.context, o.CreateMedicineHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/programs"] = NewCreateProgram(o.context, o.CreateProgramHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/enrollments"] = NewGetEnrollments(o.context, o.GetEnrollmentsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/facilities"] = NewGetFacilities(o.context, o.GetFacilitiesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/facility/groups"] = NewGetFacilityGroups(o.context, o.GetFacilityGroupsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/facility/users"] = NewGetFacilityUsers(o.context, o.GetFacilityUsersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/medicines"] = NewGetMedicines(o.context, o.GetMedicinesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/programs"] = NewGetPrograms(o.context, o.GetProgramsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/vaccinators"] = NewGetVaccinators(o.context, o.GetVaccinatorsHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *DivocPortalAPIAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *DivocPortalAPIAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *DivocPortalAPIAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *DivocPortalAPIAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *DivocPortalAPIAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
