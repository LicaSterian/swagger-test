// Code generated by go-swagger; DO NOT EDIT.

package book

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/LicaSterian/swagger-test/models"
)

// GetBookOKCode is the HTTP code returned for type GetBookOK
const GetBookOKCode int = 200

/*GetBookOK Retrieved

swagger:response getBookOK
*/
type GetBookOK struct {

	/*
	  In: Body
	*/
	Payload *models.Book `json:"body,omitempty"`
}

// NewGetBookOK creates GetBookOK with default headers values
func NewGetBookOK() *GetBookOK {

	return &GetBookOK{}
}

// WithPayload adds the payload to the get book o k response
func (o *GetBookOK) WithPayload(payload *models.Book) *GetBookOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get book o k response
func (o *GetBookOK) SetPayload(payload *models.Book) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBookOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBookBadRequestCode is the HTTP code returned for type GetBookBadRequest
const GetBookBadRequestCode int = 400

/*GetBookBadRequest invalid input

swagger:response getBookBadRequest
*/
type GetBookBadRequest struct {
}

// NewGetBookBadRequest creates GetBookBadRequest with default headers values
func NewGetBookBadRequest() *GetBookBadRequest {

	return &GetBookBadRequest{}
}

// WriteResponse to the client
func (o *GetBookBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetBookNotFoundCode is the HTTP code returned for type GetBookNotFound
const GetBookNotFoundCode int = 404

/*GetBookNotFound bookId not Found

swagger:response getBookNotFound
*/
type GetBookNotFound struct {
}

// NewGetBookNotFound creates GetBookNotFound with default headers values
func NewGetBookNotFound() *GetBookNotFound {

	return &GetBookNotFound{}
}

// WriteResponse to the client
func (o *GetBookNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
