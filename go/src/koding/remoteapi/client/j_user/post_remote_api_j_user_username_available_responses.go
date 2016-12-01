package j_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJUserUsernameAvailableReader is a Reader for the PostRemoteAPIJUserUsernameAvailable structure.
type PostRemoteAPIJUserUsernameAvailableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJUserUsernameAvailableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJUserUsernameAvailableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJUserUsernameAvailableUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJUserUsernameAvailableOK creates a PostRemoteAPIJUserUsernameAvailableOK with default headers values
func NewPostRemoteAPIJUserUsernameAvailableOK() *PostRemoteAPIJUserUsernameAvailableOK {
	return &PostRemoteAPIJUserUsernameAvailableOK{}
}

/*PostRemoteAPIJUserUsernameAvailableOK handles this case with default header values.

Request processed succesfully
*/
type PostRemoteAPIJUserUsernameAvailableOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJUserUsernameAvailableOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JUser.usernameAvailable][%d] postRemoteApiJUserUsernameAvailableOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJUserUsernameAvailableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJUserUsernameAvailableUnauthorized creates a PostRemoteAPIJUserUsernameAvailableUnauthorized with default headers values
func NewPostRemoteAPIJUserUsernameAvailableUnauthorized() *PostRemoteAPIJUserUsernameAvailableUnauthorized {
	return &PostRemoteAPIJUserUsernameAvailableUnauthorized{}
}

/*PostRemoteAPIJUserUsernameAvailableUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJUserUsernameAvailableUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJUserUsernameAvailableUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JUser.usernameAvailable][%d] postRemoteApiJUserUsernameAvailableUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJUserUsernameAvailableUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
