// Code generated by go-swagger; DO NOT EDIT.

package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/nicdoye/quaytermass/models"
)

// GetOrgRobotsReader is a Reader for the GetOrgRobots structure.
type GetOrgRobotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgRobotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrgRobotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetOrgRobotsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetOrgRobotsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetOrgRobotsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetOrgRobotsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrgRobotsOK creates a GetOrgRobotsOK with default headers values
func NewGetOrgRobotsOK() *GetOrgRobotsOK {
	return &GetOrgRobotsOK{}
}

/*GetOrgRobotsOK handles this case with default header values.

Successful invocation
*/
type GetOrgRobotsOK struct {
}

func (o *GetOrgRobotsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/robots][%d] getOrgRobotsOK ", 200)
}

func (o *GetOrgRobotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrgRobotsBadRequest creates a GetOrgRobotsBadRequest with default headers values
func NewGetOrgRobotsBadRequest() *GetOrgRobotsBadRequest {
	return &GetOrgRobotsBadRequest{}
}

/*GetOrgRobotsBadRequest handles this case with default header values.

Bad Request
*/
type GetOrgRobotsBadRequest struct {
	Payload *models.APIError
}

func (o *GetOrgRobotsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/robots][%d] getOrgRobotsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrgRobotsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgRobotsUnauthorized creates a GetOrgRobotsUnauthorized with default headers values
func NewGetOrgRobotsUnauthorized() *GetOrgRobotsUnauthorized {
	return &GetOrgRobotsUnauthorized{}
}

/*GetOrgRobotsUnauthorized handles this case with default header values.

Session required
*/
type GetOrgRobotsUnauthorized struct {
	Payload *models.APIError
}

func (o *GetOrgRobotsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/robots][%d] getOrgRobotsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrgRobotsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgRobotsForbidden creates a GetOrgRobotsForbidden with default headers values
func NewGetOrgRobotsForbidden() *GetOrgRobotsForbidden {
	return &GetOrgRobotsForbidden{}
}

/*GetOrgRobotsForbidden handles this case with default header values.

Unauthorized access
*/
type GetOrgRobotsForbidden struct {
	Payload *models.APIError
}

func (o *GetOrgRobotsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/robots][%d] getOrgRobotsForbidden  %+v", 403, o.Payload)
}

func (o *GetOrgRobotsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgRobotsNotFound creates a GetOrgRobotsNotFound with default headers values
func NewGetOrgRobotsNotFound() *GetOrgRobotsNotFound {
	return &GetOrgRobotsNotFound{}
}

/*GetOrgRobotsNotFound handles this case with default header values.

Not found
*/
type GetOrgRobotsNotFound struct {
	Payload *models.APIError
}

func (o *GetOrgRobotsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/robots][%d] getOrgRobotsNotFound  %+v", 404, o.Payload)
}

func (o *GetOrgRobotsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
