package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/go-swagger/go-swagger/examples/task-tracker/models"
)

/*UpdateTaskOK Task details

swagger:response updateTaskOK
*/
type UpdateTaskOK struct {

	// In: body
	Payload *models.Task `json:"body,omitempty"`
}

// NewUpdateTaskOK creates UpdateTaskOK with default headers values
func NewUpdateTaskOK() *UpdateTaskOK {
	return &UpdateTaskOK{}
}

// WithPayload adds the payload to the update task o k response
func (o *UpdateTaskOK) WithPayload(payload *models.Task) *UpdateTaskOK {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *UpdateTaskOK) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateTaskUnprocessableEntity Validation error

swagger:response updateTaskUnprocessableEntity
*/
type UpdateTaskUnprocessableEntity struct {

	// In: body
	Payload *models.ValidationError `json:"body,omitempty"`
}

// NewUpdateTaskUnprocessableEntity creates UpdateTaskUnprocessableEntity with default headers values
func NewUpdateTaskUnprocessableEntity() *UpdateTaskUnprocessableEntity {
	return &UpdateTaskUnprocessableEntity{}
}

// WithPayload adds the payload to the update task unprocessable entity response
func (o *UpdateTaskUnprocessableEntity) WithPayload(payload *models.ValidationError) *UpdateTaskUnprocessableEntity {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *UpdateTaskUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateTaskDefault update task default

swagger:response updateTaskDefault
*/
type UpdateTaskDefault struct {
	_statusCode int
}

// NewUpdateTaskDefault creates UpdateTaskDefault with default headers values
func NewUpdateTaskDefault(code int) *UpdateTaskDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateTaskDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update task default response
func (o *UpdateTaskDefault) WithStatusCode(code int) *UpdateTaskDefault {
	o._statusCode = code
	return o
}

// WriteResponse to the client
func (o *UpdateTaskDefault) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(o._statusCode)
}
