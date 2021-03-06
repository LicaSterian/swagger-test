// Code generated by go-swagger; DO NOT EDIT.

package reviewer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/LicaSterian/swagger-test/models"
)

// AddNewReviewerCreatedCode is the HTTP code returned for type AddNewReviewerCreated
const AddNewReviewerCreatedCode int = 201

/*AddNewReviewerCreated OK

swagger:response addNewReviewerCreated
*/
type AddNewReviewerCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Reviewer `json:"body,omitempty"`
}

// NewAddNewReviewerCreated creates AddNewReviewerCreated with default headers values
func NewAddNewReviewerCreated() *AddNewReviewerCreated {

	return &AddNewReviewerCreated{}
}

// WithPayload adds the payload to the add new reviewer created response
func (o *AddNewReviewerCreated) WithPayload(payload *models.Reviewer) *AddNewReviewerCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add new reviewer created response
func (o *AddNewReviewerCreated) SetPayload(payload *models.Reviewer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddNewReviewerCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddNewReviewerBadRequestCode is the HTTP code returned for type AddNewReviewerBadRequest
const AddNewReviewerBadRequestCode int = 400

/*AddNewReviewerBadRequest invalid input

swagger:response addNewReviewerBadRequest
*/
type AddNewReviewerBadRequest struct {
}

// NewAddNewReviewerBadRequest creates AddNewReviewerBadRequest with default headers values
func NewAddNewReviewerBadRequest() *AddNewReviewerBadRequest {

	return &AddNewReviewerBadRequest{}
}

// WriteResponse to the client
func (o *AddNewReviewerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
