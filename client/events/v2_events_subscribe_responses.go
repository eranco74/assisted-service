// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2EventsSubscribeReader is a Reader for the V2EventsSubscribe structure.
type V2EventsSubscribeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2EventsSubscribeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewV2EventsSubscribeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewV2EventsSubscribeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2EventsSubscribeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV2EventsSubscribeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2EventsSubscribeMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2EventsSubscribeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2EventsSubscribeCreated creates a V2EventsSubscribeCreated with default headers values
func NewV2EventsSubscribeCreated() *V2EventsSubscribeCreated {
	return &V2EventsSubscribeCreated{}
}

/* V2EventsSubscribeCreated describes a response with status code 201, with default header values.

Success.
*/
type V2EventsSubscribeCreated struct {
	Payload *models.EventSubscription
}

func (o *V2EventsSubscribeCreated) Error() string {
	return fmt.Sprintf("[POST /v2/events/subscription][%d] v2EventsSubscribeCreated  %+v", 201, o.Payload)
}
func (o *V2EventsSubscribeCreated) GetPayload() *models.EventSubscription {
	return o.Payload
}

func (o *V2EventsSubscribeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventSubscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2EventsSubscribeUnauthorized creates a V2EventsSubscribeUnauthorized with default headers values
func NewV2EventsSubscribeUnauthorized() *V2EventsSubscribeUnauthorized {
	return &V2EventsSubscribeUnauthorized{}
}

/* V2EventsSubscribeUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type V2EventsSubscribeUnauthorized struct {
	Payload *models.InfraError
}

func (o *V2EventsSubscribeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v2/events/subscription][%d] v2EventsSubscribeUnauthorized  %+v", 401, o.Payload)
}
func (o *V2EventsSubscribeUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2EventsSubscribeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2EventsSubscribeForbidden creates a V2EventsSubscribeForbidden with default headers values
func NewV2EventsSubscribeForbidden() *V2EventsSubscribeForbidden {
	return &V2EventsSubscribeForbidden{}
}

/* V2EventsSubscribeForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type V2EventsSubscribeForbidden struct {
	Payload *models.InfraError
}

func (o *V2EventsSubscribeForbidden) Error() string {
	return fmt.Sprintf("[POST /v2/events/subscription][%d] v2EventsSubscribeForbidden  %+v", 403, o.Payload)
}
func (o *V2EventsSubscribeForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2EventsSubscribeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2EventsSubscribeNotFound creates a V2EventsSubscribeNotFound with default headers values
func NewV2EventsSubscribeNotFound() *V2EventsSubscribeNotFound {
	return &V2EventsSubscribeNotFound{}
}

/* V2EventsSubscribeNotFound describes a response with status code 404, with default header values.

Error.
*/
type V2EventsSubscribeNotFound struct {
	Payload *models.Error
}

func (o *V2EventsSubscribeNotFound) Error() string {
	return fmt.Sprintf("[POST /v2/events/subscription][%d] v2EventsSubscribeNotFound  %+v", 404, o.Payload)
}
func (o *V2EventsSubscribeNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2EventsSubscribeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2EventsSubscribeMethodNotAllowed creates a V2EventsSubscribeMethodNotAllowed with default headers values
func NewV2EventsSubscribeMethodNotAllowed() *V2EventsSubscribeMethodNotAllowed {
	return &V2EventsSubscribeMethodNotAllowed{}
}

/* V2EventsSubscribeMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed.
*/
type V2EventsSubscribeMethodNotAllowed struct {
	Payload *models.Error
}

func (o *V2EventsSubscribeMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v2/events/subscription][%d] v2EventsSubscribeMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *V2EventsSubscribeMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2EventsSubscribeMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2EventsSubscribeInternalServerError creates a V2EventsSubscribeInternalServerError with default headers values
func NewV2EventsSubscribeInternalServerError() *V2EventsSubscribeInternalServerError {
	return &V2EventsSubscribeInternalServerError{}
}

/* V2EventsSubscribeInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type V2EventsSubscribeInternalServerError struct {
	Payload *models.Error
}

func (o *V2EventsSubscribeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v2/events/subscription][%d] v2EventsSubscribeInternalServerError  %+v", 500, o.Payload)
}
func (o *V2EventsSubscribeInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2EventsSubscribeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
