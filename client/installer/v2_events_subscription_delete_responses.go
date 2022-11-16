// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2EventsSubscriptionDeleteReader is a Reader for the V2EventsSubscriptionDelete structure.
type V2EventsSubscriptionDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2EventsSubscriptionDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV2EventsSubscriptionDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewV2EventsSubscriptionDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2EventsSubscriptionDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV2EventsSubscriptionDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewV2EventsSubscriptionDeleteMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2EventsSubscriptionDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2EventsSubscriptionDeleteNoContent creates a V2EventsSubscriptionDeleteNoContent with default headers values
func NewV2EventsSubscriptionDeleteNoContent() *V2EventsSubscriptionDeleteNoContent {
	return &V2EventsSubscriptionDeleteNoContent{}
}

/* V2EventsSubscriptionDeleteNoContent describes a response with status code 204, with default header values.

Success.
*/
type V2EventsSubscriptionDeleteNoContent struct {
}

func (o *V2EventsSubscriptionDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v2/events/subscription/{subscription_id}][%d] v2EventsSubscriptionDeleteNoContent ", 204)
}

func (o *V2EventsSubscriptionDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewV2EventsSubscriptionDeleteUnauthorized creates a V2EventsSubscriptionDeleteUnauthorized with default headers values
func NewV2EventsSubscriptionDeleteUnauthorized() *V2EventsSubscriptionDeleteUnauthorized {
	return &V2EventsSubscriptionDeleteUnauthorized{}
}

/* V2EventsSubscriptionDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type V2EventsSubscriptionDeleteUnauthorized struct {
	Payload *models.InfraError
}

func (o *V2EventsSubscriptionDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v2/events/subscription/{subscription_id}][%d] v2EventsSubscriptionDeleteUnauthorized  %+v", 401, o.Payload)
}
func (o *V2EventsSubscriptionDeleteUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2EventsSubscriptionDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2EventsSubscriptionDeleteForbidden creates a V2EventsSubscriptionDeleteForbidden with default headers values
func NewV2EventsSubscriptionDeleteForbidden() *V2EventsSubscriptionDeleteForbidden {
	return &V2EventsSubscriptionDeleteForbidden{}
}

/* V2EventsSubscriptionDeleteForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type V2EventsSubscriptionDeleteForbidden struct {
	Payload *models.InfraError
}

func (o *V2EventsSubscriptionDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v2/events/subscription/{subscription_id}][%d] v2EventsSubscriptionDeleteForbidden  %+v", 403, o.Payload)
}
func (o *V2EventsSubscriptionDeleteForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2EventsSubscriptionDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2EventsSubscriptionDeleteNotFound creates a V2EventsSubscriptionDeleteNotFound with default headers values
func NewV2EventsSubscriptionDeleteNotFound() *V2EventsSubscriptionDeleteNotFound {
	return &V2EventsSubscriptionDeleteNotFound{}
}

/* V2EventsSubscriptionDeleteNotFound describes a response with status code 404, with default header values.

Error.
*/
type V2EventsSubscriptionDeleteNotFound struct {
	Payload *models.Error
}

func (o *V2EventsSubscriptionDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v2/events/subscription/{subscription_id}][%d] v2EventsSubscriptionDeleteNotFound  %+v", 404, o.Payload)
}
func (o *V2EventsSubscriptionDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2EventsSubscriptionDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2EventsSubscriptionDeleteMethodNotAllowed creates a V2EventsSubscriptionDeleteMethodNotAllowed with default headers values
func NewV2EventsSubscriptionDeleteMethodNotAllowed() *V2EventsSubscriptionDeleteMethodNotAllowed {
	return &V2EventsSubscriptionDeleteMethodNotAllowed{}
}

/* V2EventsSubscriptionDeleteMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed.
*/
type V2EventsSubscriptionDeleteMethodNotAllowed struct {
	Payload *models.Error
}

func (o *V2EventsSubscriptionDeleteMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /v2/events/subscription/{subscription_id}][%d] v2EventsSubscriptionDeleteMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *V2EventsSubscriptionDeleteMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2EventsSubscriptionDeleteMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2EventsSubscriptionDeleteInternalServerError creates a V2EventsSubscriptionDeleteInternalServerError with default headers values
func NewV2EventsSubscriptionDeleteInternalServerError() *V2EventsSubscriptionDeleteInternalServerError {
	return &V2EventsSubscriptionDeleteInternalServerError{}
}

/* V2EventsSubscriptionDeleteInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type V2EventsSubscriptionDeleteInternalServerError struct {
	Payload *models.Error
}

func (o *V2EventsSubscriptionDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v2/events/subscription/{subscription_id}][%d] v2EventsSubscriptionDeleteInternalServerError  %+v", 500, o.Payload)
}
func (o *V2EventsSubscriptionDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2EventsSubscriptionDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
