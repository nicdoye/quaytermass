// Code generated by go-swagger; DO NOT EDIT.

package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/nicdoye/quaytermass/models"
)

// ListTagImagesReader is a Reader for the ListTagImages structure.
type ListTagImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTagImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListTagImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListTagImagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListTagImagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListTagImagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListTagImagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListTagImagesOK creates a ListTagImagesOK with default headers values
func NewListTagImagesOK() *ListTagImagesOK {
	return &ListTagImagesOK{}
}

/*ListTagImagesOK handles this case with default header values.

Successful invocation
*/
type ListTagImagesOK struct {
}

func (o *ListTagImagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/tag/{tag}/images][%d] listTagImagesOK ", 200)
}

func (o *ListTagImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListTagImagesBadRequest creates a ListTagImagesBadRequest with default headers values
func NewListTagImagesBadRequest() *ListTagImagesBadRequest {
	return &ListTagImagesBadRequest{}
}

/*ListTagImagesBadRequest handles this case with default header values.

Bad Request
*/
type ListTagImagesBadRequest struct {
	Payload *models.APIError
}

func (o *ListTagImagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/tag/{tag}/images][%d] listTagImagesBadRequest  %+v", 400, o.Payload)
}

func (o *ListTagImagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTagImagesUnauthorized creates a ListTagImagesUnauthorized with default headers values
func NewListTagImagesUnauthorized() *ListTagImagesUnauthorized {
	return &ListTagImagesUnauthorized{}
}

/*ListTagImagesUnauthorized handles this case with default header values.

Session required
*/
type ListTagImagesUnauthorized struct {
	Payload *models.APIError
}

func (o *ListTagImagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/tag/{tag}/images][%d] listTagImagesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListTagImagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTagImagesForbidden creates a ListTagImagesForbidden with default headers values
func NewListTagImagesForbidden() *ListTagImagesForbidden {
	return &ListTagImagesForbidden{}
}

/*ListTagImagesForbidden handles this case with default header values.

Unauthorized access
*/
type ListTagImagesForbidden struct {
	Payload *models.APIError
}

func (o *ListTagImagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/tag/{tag}/images][%d] listTagImagesForbidden  %+v", 403, o.Payload)
}

func (o *ListTagImagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTagImagesNotFound creates a ListTagImagesNotFound with default headers values
func NewListTagImagesNotFound() *ListTagImagesNotFound {
	return &ListTagImagesNotFound{}
}

/*ListTagImagesNotFound handles this case with default header values.

Not found
*/
type ListTagImagesNotFound struct {
	Payload *models.APIError
}

func (o *ListTagImagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/tag/{tag}/images][%d] listTagImagesNotFound  %+v", 404, o.Payload)
}

func (o *ListTagImagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
