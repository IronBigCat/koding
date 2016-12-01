package j_workspace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJWorkspaceFetchByMachinesReader is a Reader for the PostRemoteAPIJWorkspaceFetchByMachines structure.
type PostRemoteAPIJWorkspaceFetchByMachinesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJWorkspaceFetchByMachinesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJWorkspaceFetchByMachinesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJWorkspaceFetchByMachinesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJWorkspaceFetchByMachinesOK creates a PostRemoteAPIJWorkspaceFetchByMachinesOK with default headers values
func NewPostRemoteAPIJWorkspaceFetchByMachinesOK() *PostRemoteAPIJWorkspaceFetchByMachinesOK {
	return &PostRemoteAPIJWorkspaceFetchByMachinesOK{}
}

/*PostRemoteAPIJWorkspaceFetchByMachinesOK handles this case with default header values.

Request processed succesfully
*/
type PostRemoteAPIJWorkspaceFetchByMachinesOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJWorkspaceFetchByMachinesOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JWorkspace.fetchByMachines][%d] postRemoteApiJWorkspaceFetchByMachinesOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJWorkspaceFetchByMachinesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJWorkspaceFetchByMachinesUnauthorized creates a PostRemoteAPIJWorkspaceFetchByMachinesUnauthorized with default headers values
func NewPostRemoteAPIJWorkspaceFetchByMachinesUnauthorized() *PostRemoteAPIJWorkspaceFetchByMachinesUnauthorized {
	return &PostRemoteAPIJWorkspaceFetchByMachinesUnauthorized{}
}

/*PostRemoteAPIJWorkspaceFetchByMachinesUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJWorkspaceFetchByMachinesUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJWorkspaceFetchByMachinesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JWorkspace.fetchByMachines][%d] postRemoteApiJWorkspaceFetchByMachinesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJWorkspaceFetchByMachinesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
