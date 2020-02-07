// Code generated by go-swagger; DO NOT EDIT.

package book

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UpdateBookOKCode is the HTTP code returned for type UpdateBookOK
const UpdateBookOKCode int = 200

/*UpdateBookOK Updated

swagger:response updateBookOK
*/
type UpdateBookOK struct {
}

// NewUpdateBookOK creates UpdateBookOK with default headers values
func NewUpdateBookOK() *UpdateBookOK {

	return &UpdateBookOK{}
}

// WriteResponse to the client
func (o *UpdateBookOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UpdateBookBadRequestCode is the HTTP code returned for type UpdateBookBadRequest
const UpdateBookBadRequestCode int = 400

/*UpdateBookBadRequest invalid input

swagger:response updateBookBadRequest
*/
type UpdateBookBadRequest struct {
}

// NewUpdateBookBadRequest creates UpdateBookBadRequest with default headers values
func NewUpdateBookBadRequest() *UpdateBookBadRequest {

	return &UpdateBookBadRequest{}
}

// WriteResponse to the client
func (o *UpdateBookBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// UpdateBookNotFoundCode is the HTTP code returned for type UpdateBookNotFound
const UpdateBookNotFoundCode int = 404

/*UpdateBookNotFound bookId not Found

swagger:response updateBookNotFound
*/
type UpdateBookNotFound struct {
}

// NewUpdateBookNotFound creates UpdateBookNotFound with default headers values
func NewUpdateBookNotFound() *UpdateBookNotFound {

	return &UpdateBookNotFound{}
}

// WriteResponse to the client
func (o *UpdateBookNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
