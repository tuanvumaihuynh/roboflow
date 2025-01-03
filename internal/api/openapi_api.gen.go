// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List QR locations
	// (GET /qr-locations)
	ListQRLocations(w http.ResponseWriter, r *http.Request, params ListQRLocationsParams)
	// Create QR location
	// (POST /qr-locations)
	CreateQRLocation(w http.ResponseWriter, r *http.Request)
	// Delete QR location by id
	// (DELETE /qr-locations/{qrLocationId})
	DeleteQRLocationById(w http.ResponseWriter, r *http.Request, qrLocationId QRLocationId)
	// Get QR location by id
	// (GET /qr-locations/{qrLocationId})
	GetQRLocationById(w http.ResponseWriter, r *http.Request, qrLocationId QRLocationId)
	// Update QR location by id
	// (PUT /qr-locations/{qrLocationId})
	UpdateQRLocationById(w http.ResponseWriter, r *http.Request, qrLocationId QRLocationId)
	// Get raybot command by id
	// (GET /raybot-commands/{raybotCommandId})
	GetRaybotCommandById(w http.ResponseWriter, r *http.Request, raybotCommandId RaybotCommandId)
	// List raybots
	// (GET /raybots)
	ListRaybots(w http.ResponseWriter, r *http.Request, params ListRaybotsParams)
	// Create raybot
	// (POST /raybots)
	CreateRaybot(w http.ResponseWriter, r *http.Request)
	// Delete raybot by id
	// (DELETE /raybots/{raybotId})
	DeleteRaybotById(w http.ResponseWriter, r *http.Request, raybotId RaybotId)
	// Get raybot by id
	// (GET /raybots/{raybotId})
	GetRaybotById(w http.ResponseWriter, r *http.Request, raybotId RaybotId)
	// List raybot commands
	// (GET /raybots/{raybotId}/commands)
	ListRaybotCommands(w http.ResponseWriter, r *http.Request, raybotId RaybotId, params ListRaybotCommandsParams)
	// Create raybot command
	// (POST /raybots/{raybotId}/commands)
	CreateRaybotCommand(w http.ResponseWriter, r *http.Request, raybotId RaybotId)
	// Get step by id
	// (GET /steps/{stepId})
	GetStepById(w http.ResponseWriter, r *http.Request, stepId StepId)
	// Get workflow execution by id
	// (GET /workflow-executions/{workflowExecutionId})
	GetWorkflowExecutionById(w http.ResponseWriter, r *http.Request, workflowExecutionId WorkflowExecutionId)
	// Get workflow execution status by id
	// (GET /workflow-executions/{workflowExecutionId}/status)
	GetWorkflowExecutionStatusById(w http.ResponseWriter, r *http.Request, workflowExecutionId WorkflowExecutionId)
	// List steps by workflow execution id
	// (GET /workflow-executions/{workflowExecutionId}/steps)
	ListStepsByWorkflowExecutionId(w http.ResponseWriter, r *http.Request, workflowExecutionId WorkflowExecutionId, params ListStepsByWorkflowExecutionIdParams)
	// List  workflows
	// (GET /workflows)
	ListWorkflows(w http.ResponseWriter, r *http.Request, params ListWorkflowsParams)
	// Create workflow
	// (POST /workflows)
	CreateWorkflow(w http.ResponseWriter, r *http.Request)
	// Delete workflow by id
	// (DELETE /workflows/{workflowId})
	DeleteWorkflowById(w http.ResponseWriter, r *http.Request, workflowId WorkflowId)
	// Get workflow by id
	// (GET /workflows/{workflowId})
	GetWorkflowById(w http.ResponseWriter, r *http.Request, workflowId WorkflowId)
	// Update workflow by id
	// (PUT /workflows/{workflowId})
	UpdateWorkflowById(w http.ResponseWriter, r *http.Request, workflowId WorkflowId)
	// List workflow executions by workflow id
	// (GET /workflows/{workflowId}/executions)
	ListWorkflowExecutionsByWorkflowId(w http.ResponseWriter, r *http.Request, workflowId WorkflowId, params ListWorkflowExecutionsByWorkflowIdParams)
	// Run workflow by id
	// (POST /workflows/{workflowId}/run)
	RunWorkflowById(w http.ResponseWriter, r *http.Request, workflowId WorkflowId)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// List QR locations
// (GET /qr-locations)
func (_ Unimplemented) ListQRLocations(w http.ResponseWriter, r *http.Request, params ListQRLocationsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create QR location
// (POST /qr-locations)
func (_ Unimplemented) CreateQRLocation(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete QR location by id
// (DELETE /qr-locations/{qrLocationId})
func (_ Unimplemented) DeleteQRLocationById(w http.ResponseWriter, r *http.Request, qrLocationId QRLocationId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get QR location by id
// (GET /qr-locations/{qrLocationId})
func (_ Unimplemented) GetQRLocationById(w http.ResponseWriter, r *http.Request, qrLocationId QRLocationId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update QR location by id
// (PUT /qr-locations/{qrLocationId})
func (_ Unimplemented) UpdateQRLocationById(w http.ResponseWriter, r *http.Request, qrLocationId QRLocationId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get raybot command by id
// (GET /raybot-commands/{raybotCommandId})
func (_ Unimplemented) GetRaybotCommandById(w http.ResponseWriter, r *http.Request, raybotCommandId RaybotCommandId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// List raybots
// (GET /raybots)
func (_ Unimplemented) ListRaybots(w http.ResponseWriter, r *http.Request, params ListRaybotsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create raybot
// (POST /raybots)
func (_ Unimplemented) CreateRaybot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete raybot by id
// (DELETE /raybots/{raybotId})
func (_ Unimplemented) DeleteRaybotById(w http.ResponseWriter, r *http.Request, raybotId RaybotId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get raybot by id
// (GET /raybots/{raybotId})
func (_ Unimplemented) GetRaybotById(w http.ResponseWriter, r *http.Request, raybotId RaybotId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// List raybot commands
// (GET /raybots/{raybotId}/commands)
func (_ Unimplemented) ListRaybotCommands(w http.ResponseWriter, r *http.Request, raybotId RaybotId, params ListRaybotCommandsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create raybot command
// (POST /raybots/{raybotId}/commands)
func (_ Unimplemented) CreateRaybotCommand(w http.ResponseWriter, r *http.Request, raybotId RaybotId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get step by id
// (GET /steps/{stepId})
func (_ Unimplemented) GetStepById(w http.ResponseWriter, r *http.Request, stepId StepId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get workflow execution by id
// (GET /workflow-executions/{workflowExecutionId})
func (_ Unimplemented) GetWorkflowExecutionById(w http.ResponseWriter, r *http.Request, workflowExecutionId WorkflowExecutionId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get workflow execution status by id
// (GET /workflow-executions/{workflowExecutionId}/status)
func (_ Unimplemented) GetWorkflowExecutionStatusById(w http.ResponseWriter, r *http.Request, workflowExecutionId WorkflowExecutionId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// List steps by workflow execution id
// (GET /workflow-executions/{workflowExecutionId}/steps)
func (_ Unimplemented) ListStepsByWorkflowExecutionId(w http.ResponseWriter, r *http.Request, workflowExecutionId WorkflowExecutionId, params ListStepsByWorkflowExecutionIdParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// List  workflows
// (GET /workflows)
func (_ Unimplemented) ListWorkflows(w http.ResponseWriter, r *http.Request, params ListWorkflowsParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create workflow
// (POST /workflows)
func (_ Unimplemented) CreateWorkflow(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete workflow by id
// (DELETE /workflows/{workflowId})
func (_ Unimplemented) DeleteWorkflowById(w http.ResponseWriter, r *http.Request, workflowId WorkflowId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get workflow by id
// (GET /workflows/{workflowId})
func (_ Unimplemented) GetWorkflowById(w http.ResponseWriter, r *http.Request, workflowId WorkflowId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update workflow by id
// (PUT /workflows/{workflowId})
func (_ Unimplemented) UpdateWorkflowById(w http.ResponseWriter, r *http.Request, workflowId WorkflowId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// List workflow executions by workflow id
// (GET /workflows/{workflowId}/executions)
func (_ Unimplemented) ListWorkflowExecutionsByWorkflowId(w http.ResponseWriter, r *http.Request, workflowId WorkflowId, params ListWorkflowExecutionsByWorkflowIdParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Run workflow by id
// (POST /workflows/{workflowId}/run)
func (_ Unimplemented) RunWorkflowById(w http.ResponseWriter, r *http.Request, workflowId WorkflowId) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// ListQRLocations operation middleware
func (siw *ServerInterfaceWrapper) ListQRLocations(w http.ResponseWriter, r *http.Request) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ListQRLocationsParams

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", r.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page", Err: err})
		return
	}

	// ------------- Optional query parameter "pageSize" -------------

	err = runtime.BindQueryParameter("form", true, false, "pageSize", r.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "pageSize", Err: err})
		return
	}

	// ------------- Optional query parameter "sort" -------------

	err = runtime.BindQueryParameter("form", true, false, "sort", r.URL.Query(), &params.Sort)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sort", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListQRLocations(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// CreateQRLocation operation middleware
func (siw *ServerInterfaceWrapper) CreateQRLocation(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateQRLocation(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// DeleteQRLocationById operation middleware
func (siw *ServerInterfaceWrapper) DeleteQRLocationById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "qrLocationId" -------------
	var qrLocationId QRLocationId

	err = runtime.BindStyledParameterWithOptions("simple", "qrLocationId", chi.URLParam(r, "qrLocationId"), &qrLocationId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "qrLocationId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteQRLocationById(w, r, qrLocationId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetQRLocationById operation middleware
func (siw *ServerInterfaceWrapper) GetQRLocationById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "qrLocationId" -------------
	var qrLocationId QRLocationId

	err = runtime.BindStyledParameterWithOptions("simple", "qrLocationId", chi.URLParam(r, "qrLocationId"), &qrLocationId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "qrLocationId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetQRLocationById(w, r, qrLocationId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// UpdateQRLocationById operation middleware
func (siw *ServerInterfaceWrapper) UpdateQRLocationById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "qrLocationId" -------------
	var qrLocationId QRLocationId

	err = runtime.BindStyledParameterWithOptions("simple", "qrLocationId", chi.URLParam(r, "qrLocationId"), &qrLocationId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "qrLocationId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateQRLocationById(w, r, qrLocationId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetRaybotCommandById operation middleware
func (siw *ServerInterfaceWrapper) GetRaybotCommandById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "raybotCommandId" -------------
	var raybotCommandId RaybotCommandId

	err = runtime.BindStyledParameterWithOptions("simple", "raybotCommandId", chi.URLParam(r, "raybotCommandId"), &raybotCommandId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "raybotCommandId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetRaybotCommandById(w, r, raybotCommandId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// ListRaybots operation middleware
func (siw *ServerInterfaceWrapper) ListRaybots(w http.ResponseWriter, r *http.Request) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ListRaybotsParams

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", r.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page", Err: err})
		return
	}

	// ------------- Optional query parameter "pageSize" -------------

	err = runtime.BindQueryParameter("form", true, false, "pageSize", r.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "pageSize", Err: err})
		return
	}

	// ------------- Optional query parameter "sort" -------------

	err = runtime.BindQueryParameter("form", true, false, "sort", r.URL.Query(), &params.Sort)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sort", Err: err})
		return
	}

	// ------------- Optional query parameter "status" -------------

	err = runtime.BindQueryParameter("form", true, false, "status", r.URL.Query(), &params.Status)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "status", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListRaybots(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// CreateRaybot operation middleware
func (siw *ServerInterfaceWrapper) CreateRaybot(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateRaybot(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// DeleteRaybotById operation middleware
func (siw *ServerInterfaceWrapper) DeleteRaybotById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "raybotId" -------------
	var raybotId RaybotId

	err = runtime.BindStyledParameterWithOptions("simple", "raybotId", chi.URLParam(r, "raybotId"), &raybotId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "raybotId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteRaybotById(w, r, raybotId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetRaybotById operation middleware
func (siw *ServerInterfaceWrapper) GetRaybotById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "raybotId" -------------
	var raybotId RaybotId

	err = runtime.BindStyledParameterWithOptions("simple", "raybotId", chi.URLParam(r, "raybotId"), &raybotId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "raybotId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetRaybotById(w, r, raybotId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// ListRaybotCommands operation middleware
func (siw *ServerInterfaceWrapper) ListRaybotCommands(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "raybotId" -------------
	var raybotId RaybotId

	err = runtime.BindStyledParameterWithOptions("simple", "raybotId", chi.URLParam(r, "raybotId"), &raybotId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "raybotId", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params ListRaybotCommandsParams

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", r.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page", Err: err})
		return
	}

	// ------------- Optional query parameter "pageSize" -------------

	err = runtime.BindQueryParameter("form", true, false, "pageSize", r.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "pageSize", Err: err})
		return
	}

	// ------------- Optional query parameter "sort" -------------

	err = runtime.BindQueryParameter("form", true, false, "sort", r.URL.Query(), &params.Sort)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sort", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListRaybotCommands(w, r, raybotId, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// CreateRaybotCommand operation middleware
func (siw *ServerInterfaceWrapper) CreateRaybotCommand(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "raybotId" -------------
	var raybotId RaybotId

	err = runtime.BindStyledParameterWithOptions("simple", "raybotId", chi.URLParam(r, "raybotId"), &raybotId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "raybotId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateRaybotCommand(w, r, raybotId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetStepById operation middleware
func (siw *ServerInterfaceWrapper) GetStepById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "stepId" -------------
	var stepId StepId

	err = runtime.BindStyledParameterWithOptions("simple", "stepId", chi.URLParam(r, "stepId"), &stepId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "stepId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetStepById(w, r, stepId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetWorkflowExecutionById operation middleware
func (siw *ServerInterfaceWrapper) GetWorkflowExecutionById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "workflowExecutionId" -------------
	var workflowExecutionId WorkflowExecutionId

	err = runtime.BindStyledParameterWithOptions("simple", "workflowExecutionId", chi.URLParam(r, "workflowExecutionId"), &workflowExecutionId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "workflowExecutionId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetWorkflowExecutionById(w, r, workflowExecutionId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetWorkflowExecutionStatusById operation middleware
func (siw *ServerInterfaceWrapper) GetWorkflowExecutionStatusById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "workflowExecutionId" -------------
	var workflowExecutionId WorkflowExecutionId

	err = runtime.BindStyledParameterWithOptions("simple", "workflowExecutionId", chi.URLParam(r, "workflowExecutionId"), &workflowExecutionId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "workflowExecutionId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetWorkflowExecutionStatusById(w, r, workflowExecutionId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// ListStepsByWorkflowExecutionId operation middleware
func (siw *ServerInterfaceWrapper) ListStepsByWorkflowExecutionId(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "workflowExecutionId" -------------
	var workflowExecutionId WorkflowExecutionId

	err = runtime.BindStyledParameterWithOptions("simple", "workflowExecutionId", chi.URLParam(r, "workflowExecutionId"), &workflowExecutionId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "workflowExecutionId", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params ListStepsByWorkflowExecutionIdParams

	// ------------- Optional query parameter "sort" -------------

	err = runtime.BindQueryParameter("form", true, false, "sort", r.URL.Query(), &params.Sort)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sort", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListStepsByWorkflowExecutionId(w, r, workflowExecutionId, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// ListWorkflows operation middleware
func (siw *ServerInterfaceWrapper) ListWorkflows(w http.ResponseWriter, r *http.Request) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ListWorkflowsParams

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", r.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page", Err: err})
		return
	}

	// ------------- Optional query parameter "pageSize" -------------

	err = runtime.BindQueryParameter("form", true, false, "pageSize", r.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "pageSize", Err: err})
		return
	}

	// ------------- Optional query parameter "sort" -------------

	err = runtime.BindQueryParameter("form", true, false, "sort", r.URL.Query(), &params.Sort)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sort", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListWorkflows(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// CreateWorkflow operation middleware
func (siw *ServerInterfaceWrapper) CreateWorkflow(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateWorkflow(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// DeleteWorkflowById operation middleware
func (siw *ServerInterfaceWrapper) DeleteWorkflowById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "workflowId" -------------
	var workflowId WorkflowId

	err = runtime.BindStyledParameterWithOptions("simple", "workflowId", chi.URLParam(r, "workflowId"), &workflowId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "workflowId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteWorkflowById(w, r, workflowId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetWorkflowById operation middleware
func (siw *ServerInterfaceWrapper) GetWorkflowById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "workflowId" -------------
	var workflowId WorkflowId

	err = runtime.BindStyledParameterWithOptions("simple", "workflowId", chi.URLParam(r, "workflowId"), &workflowId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "workflowId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetWorkflowById(w, r, workflowId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// UpdateWorkflowById operation middleware
func (siw *ServerInterfaceWrapper) UpdateWorkflowById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "workflowId" -------------
	var workflowId WorkflowId

	err = runtime.BindStyledParameterWithOptions("simple", "workflowId", chi.URLParam(r, "workflowId"), &workflowId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "workflowId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateWorkflowById(w, r, workflowId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// ListWorkflowExecutionsByWorkflowId operation middleware
func (siw *ServerInterfaceWrapper) ListWorkflowExecutionsByWorkflowId(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "workflowId" -------------
	var workflowId WorkflowId

	err = runtime.BindStyledParameterWithOptions("simple", "workflowId", chi.URLParam(r, "workflowId"), &workflowId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "workflowId", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params ListWorkflowExecutionsByWorkflowIdParams

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", r.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "page", Err: err})
		return
	}

	// ------------- Optional query parameter "pageSize" -------------

	err = runtime.BindQueryParameter("form", true, false, "pageSize", r.URL.Query(), &params.PageSize)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "pageSize", Err: err})
		return
	}

	// ------------- Optional query parameter "sort" -------------

	err = runtime.BindQueryParameter("form", true, false, "sort", r.URL.Query(), &params.Sort)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sort", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListWorkflowExecutionsByWorkflowId(w, r, workflowId, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// RunWorkflowById operation middleware
func (siw *ServerInterfaceWrapper) RunWorkflowById(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "workflowId" -------------
	var workflowId WorkflowId

	err = runtime.BindStyledParameterWithOptions("simple", "workflowId", chi.URLParam(r, "workflowId"), &workflowId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "workflowId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RunWorkflowById(w, r, workflowId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/qr-locations", wrapper.ListQRLocations)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/qr-locations", wrapper.CreateQRLocation)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/qr-locations/{qrLocationId}", wrapper.DeleteQRLocationById)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/qr-locations/{qrLocationId}", wrapper.GetQRLocationById)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/qr-locations/{qrLocationId}", wrapper.UpdateQRLocationById)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/raybot-commands/{raybotCommandId}", wrapper.GetRaybotCommandById)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/raybots", wrapper.ListRaybots)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/raybots", wrapper.CreateRaybot)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/raybots/{raybotId}", wrapper.DeleteRaybotById)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/raybots/{raybotId}", wrapper.GetRaybotById)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/raybots/{raybotId}/commands", wrapper.ListRaybotCommands)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/raybots/{raybotId}/commands", wrapper.CreateRaybotCommand)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/steps/{stepId}", wrapper.GetStepById)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/workflow-executions/{workflowExecutionId}", wrapper.GetWorkflowExecutionById)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/workflow-executions/{workflowExecutionId}/status", wrapper.GetWorkflowExecutionStatusById)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/workflow-executions/{workflowExecutionId}/steps", wrapper.ListStepsByWorkflowExecutionId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/workflows", wrapper.ListWorkflows)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/workflows", wrapper.CreateWorkflow)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/workflows/{workflowId}", wrapper.DeleteWorkflowById)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/workflows/{workflowId}", wrapper.GetWorkflowById)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/workflows/{workflowId}", wrapper.UpdateWorkflowById)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/workflows/{workflowId}/executions", wrapper.ListWorkflowExecutionsByWorkflowId)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/workflows/{workflowId}/run", wrapper.RunWorkflowById)
	})

	return r
}
