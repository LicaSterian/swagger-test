// Code generated by go-swagger; DO NOT EDIT.

package reviewer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/LicaSterian/swagger-test/models"
)

// GetReviewerOKCode is the HTTP code returned for type GetReviewerOK
const GetReviewerOKCode int = 200

/*GetReviewerOK Retrieved

swagger:response getReviewerOK
*/
type GetReviewerOK struct {

	/*
	  In: Body
	*/
	Payload *models.Reviewer `json:"body,omitempty"`
}

// NewGetReviewerOK creates GetReviewerOK with default headers values
func NewGetReviewerOK() *GetReviewerOK {

	return &GetReviewerOK{}
}

// WithPayload adds the payload to the get reviewer o k response
func (o *GetReviewerOK) WithPayload(payload *models.Reviewer) *GetReviewerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get reviewer o k response
func (o *GetReviewerOK) SetPayload(payload *models.Reviewer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReviewerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReviewerBadRequestCode is the HTTP code returned for type GetReviewerBadRequest
const GetReviewerBadRequestCode int = 400

/*GetReviewerBadRequest invalid input

swagger:response getReviewerBadRequest
*/
type GetReviewerBadRequest struct {
}

// NewGetReviewerBadRequest creates GetReviewerBadRequest with default headers values
func NewGetReviewerBadRequest() *GetReviewerBadRequest {

	return &GetReviewerBadRequest{}
}

// WriteResponse to the client
func (o *GetReviewerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetReviewerNotFoundCode is the HTTP code returned for type GetReviewerNotFound
const GetReviewerNotFoundCode int = 404

/*GetReviewerNotFound reviewer id not Found

swagger:response getReviewerNotFound
*/
type GetReviewerNotFound struct {
}

// NewGetReviewerNotFound creates GetReviewerNotFound with default headers values
func NewGetReviewerNotFound() *GetReviewerNotFound {

	return &GetReviewerNotFound{}
}

// WriteResponse to the client
func (o *GetReviewerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
