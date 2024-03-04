// Code generated by go-swagger; DO NOT EDIT.

package usuarios

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AgregarUsuarioHandlerFunc turns a function with the right signature into a agregar usuario handler
type AgregarUsuarioHandlerFunc func(AgregarUsuarioParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AgregarUsuarioHandlerFunc) Handle(params AgregarUsuarioParams) middleware.Responder {
	return fn(params)
}

// AgregarUsuarioHandler interface for that can handle valid agregar usuario params
type AgregarUsuarioHandler interface {
	Handle(AgregarUsuarioParams) middleware.Responder
}

// NewAgregarUsuario creates a new http.Handler for the agregar usuario operation
func NewAgregarUsuario(ctx *middleware.Context, handler AgregarUsuarioHandler) *AgregarUsuario {
	return &AgregarUsuario{Context: ctx, Handler: handler}
}

/*
AgregarUsuario swagger:route POST /usuarios usuarios agregarUsuario

# Agrega un nuevo usuario

Crea un nuevo usuario con la información proporcionada
*/
type AgregarUsuario struct {
	Context *middleware.Context
	Handler AgregarUsuarioHandler
}

func (o *AgregarUsuario) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAgregarUsuarioParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
