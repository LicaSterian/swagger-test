// Code generated by go-swagger; DO NOT EDIT.

package reviewer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteReviewerHandlerFunc turns a function with the right signature into a delete reviewer handler
type DeleteReviewerHandlerFunc func(DeleteReviewerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteReviewerHandlerFunc) Handle(params DeleteReviewerParams) middleware.Responder {
	return fn(params)
}

// DeleteReviewerHandler interface for that can handle valid delete reviewer params
type DeleteReviewerHandler interface {
	Handle(DeleteReviewerParams) middleware.Responder
}

// NewDeleteReviewer creates a new http.Handler for the delete reviewer operation
func NewDeleteReviewer(ctx *middleware.Context, handler DeleteReviewerHandler) *DeleteReviewer {
	return &DeleteReviewer{Context: ctx, Handler: handler}
}

/*DeleteReviewer swagger:route DELETE /reviewer/{id} reviewer deleteReviewer

delete reviewer from store by reviewerId

*/
type DeleteReviewer struct {
	Context *middleware.Context
	Handler DeleteReviewerHandler
}

func (o *DeleteReviewer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteReviewerParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
