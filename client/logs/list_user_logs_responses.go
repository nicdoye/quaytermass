// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/nicdoye/quaytermass/models"
)

// ListUserLogsReader is a Reader for the ListUserLogs structure.
type ListUserLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUserLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListUserLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListUserLogsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListUserLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListUserLogsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListUserLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListUserLogsOK creates a ListUserLogsOK with default headers values
func NewListUserLogsOK() *ListUserLogsOK {
	return &ListUserLogsOK{}
}

/*ListUserLogsOK handles this case with default header values.

Successful invocation
*/
type ListUserLogsOK struct {
}

func (o *ListUserLogsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/logs][%d] listUserLogsOK ", 200)
}

func (o *ListUserLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListUserLogsBadRequest creates a ListUserLogsBadRequest with default headers values
func NewListUserLogsBadRequest() *ListUserLogsBadRequest {
	return &ListUserLogsBadRequest{}
}

/*ListUserLogsBadRequest handles this case with default header values.

Bad Request
*/
type ListUserLogsBadRequest struct {
	Payload *models.APIError
}

func (o *ListUserLogsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/logs][%d] listUserLogsBadRequest  %+v", 400, o.Payload)
}

func (o *ListUserLogsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserLogsUnauthorized creates a ListUserLogsUnauthorized with default headers values
func NewListUserLogsUnauthorized() *ListUserLogsUnauthorized {
	return &ListUserLogsUnauthorized{}
}

/*ListUserLogsUnauthorized handles this case with default header values.

Session required
*/
type ListUserLogsUnauthorized struct {
	Payload *models.APIError
}

func (o *ListUserLogsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/logs][%d] listUserLogsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListUserLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserLogsForbidden creates a ListUserLogsForbidden with default headers values
func NewListUserLogsForbidden() *ListUserLogsForbidden {
	return &ListUserLogsForbidden{}
}

/*ListUserLogsForbidden handles this case with default header values.

Unauthorized access
*/
type ListUserLogsForbidden struct {
	Payload *models.APIError
}

func (o *ListUserLogsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/logs][%d] listUserLogsForbidden  %+v", 403, o.Payload)
}

func (o *ListUserLogsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserLogsNotFound creates a ListUserLogsNotFound with default headers values
func NewListUserLogsNotFound() *ListUserLogsNotFound {
	return &ListUserLogsNotFound{}
}

/*ListUserLogsNotFound handles this case with default header values.

Not found
*/
type ListUserLogsNotFound struct {
	Payload *models.APIError
}

func (o *ListUserLogsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/logs][%d] listUserLogsNotFound  %+v", 404, o.Payload)
}

func (o *ListUserLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
