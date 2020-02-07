// Code generated by go-swagger; DO NOT EDIT.

package reviewer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/LicaSterian/swagger-test/models"
)

// GetAllReviewersOKCode is the HTTP code returned for type GetAllReviewersOK
const GetAllReviewersOKCode int = 200

/*GetAllReviewersOK list of reviewers

swagger:response getAllReviewersOK
*/
type GetAllReviewersOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Reviewer `json:"body,omitempty"`
}

// NewGetAllReviewersOK creates GetAllReviewersOK with default headers values
func NewGetAllReviewersOK() *GetAllReviewersOK {

	return &GetAllReviewersOK{}
}

// WithPayload adds the payload to the get all reviewers o k response
func (o *GetAllReviewersOK) WithPayload(payload []*models.Reviewer) *GetAllReviewersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all reviewers o k response
func (o *GetAllReviewersOK) SetPayload(payload []*models.Reviewer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllReviewersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Reviewer, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAllReviewersBadRequestCode is the HTTP code returned for type GetAllReviewersBadRequest
const GetAllReviewersBadRequestCode int = 400

/*GetAllReviewersBadRequest bad input parameter

swagger:response getAllReviewersBadRequest
*/
type GetAllReviewersBadRequest struct {
}

// NewGetAllReviewersBadRequest creates GetAllReviewersBadRequest with default headers values
func NewGetAllReviewersBadRequest() *GetAllReviewersBadRequest {

	return &GetAllReviewersBadRequest{}
}

// WriteResponse to the client
func (o *GetAllReviewersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
