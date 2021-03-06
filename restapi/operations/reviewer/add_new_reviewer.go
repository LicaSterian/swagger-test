// Code generated by go-swagger; DO NOT EDIT.

package reviewer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddNewReviewerHandlerFunc turns a function with the right signature into a add new reviewer handler
type AddNewReviewerHandlerFunc func(AddNewReviewerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddNewReviewerHandlerFunc) Handle(params AddNewReviewerParams) middleware.Responder {
	return fn(params)
}

// AddNewReviewerHandler interface for that can handle valid add new reviewer params
type AddNewReviewerHandler interface {
	Handle(AddNewReviewerParams) middleware.Responder
}

// NewAddNewReviewer creates a new http.Handler for the add new reviewer operation
func NewAddNewReviewer(ctx *middleware.Context, handler AddNewReviewerHandler) *AddNewReviewer {
	return &AddNewReviewer{Context: ctx, Handler: handler}
}

/*AddNewReviewer swagger:route POST /reviewer reviewer addNewReviewer

add New Reviewer in Store

*/
type AddNewReviewer struct {
	Context *middleware.Context
	Handler AddNewReviewerHandler
}

func (o *AddNewReviewer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddNewReviewerParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
