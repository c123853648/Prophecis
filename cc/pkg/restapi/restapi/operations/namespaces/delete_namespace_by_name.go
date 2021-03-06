// Code generated by go-swagger; DO NOT EDIT.

package namespaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteNamespaceByNameHandlerFunc turns a function with the right signature into a delete namespace by name handler
type DeleteNamespaceByNameHandlerFunc func(DeleteNamespaceByNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteNamespaceByNameHandlerFunc) Handle(params DeleteNamespaceByNameParams) middleware.Responder {
	return fn(params)
}

// DeleteNamespaceByNameHandler interface for that can handle valid delete namespace by name params
type DeleteNamespaceByNameHandler interface {
	Handle(DeleteNamespaceByNameParams) middleware.Responder
}

// NewDeleteNamespaceByName creates a new http.Handler for the delete namespace by name operation
func NewDeleteNamespaceByName(ctx *middleware.Context, handler DeleteNamespaceByNameHandler) *DeleteNamespaceByName {
	return &DeleteNamespaceByName{Context: ctx, Handler: handler}
}

/*DeleteNamespaceByName swagger:route DELETE /cc/v1/namespaces/namespace/{namespace} Namespaces deleteNamespaceByName

Returns a namespace.

Optional extended description in Markdown.

*/
type DeleteNamespaceByName struct {
	Context *middleware.Context
	Handler DeleteNamespaceByNameHandler
}

func (o *DeleteNamespaceByName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteNamespaceByNameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
