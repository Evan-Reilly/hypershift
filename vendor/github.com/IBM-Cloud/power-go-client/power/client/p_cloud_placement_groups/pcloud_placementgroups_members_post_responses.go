// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_placement_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudPlacementgroupsMembersPostReader is a Reader for the PcloudPlacementgroupsMembersPost structure.
type PcloudPlacementgroupsMembersPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPlacementgroupsMembersPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPlacementgroupsMembersPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPlacementgroupsMembersPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPlacementgroupsMembersPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudPlacementgroupsMembersPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudPlacementgroupsMembersPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPlacementgroupsMembersPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudPlacementgroupsMembersPostOK creates a PcloudPlacementgroupsMembersPostOK with default headers values
func NewPcloudPlacementgroupsMembersPostOK() *PcloudPlacementgroupsMembersPostOK {
	return &PcloudPlacementgroupsMembersPostOK{}
}

/* PcloudPlacementgroupsMembersPostOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPlacementgroupsMembersPostOK struct {
	Payload *models.PlacementGroup
}

func (o *PcloudPlacementgroupsMembersPostOK) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersPostOK  %+v", 200, o.Payload)
}
func (o *PcloudPlacementgroupsMembersPostOK) GetPayload() *models.PlacementGroup {
	return o.Payload
}

func (o *PcloudPlacementgroupsMembersPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlacementGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsMembersPostBadRequest creates a PcloudPlacementgroupsMembersPostBadRequest with default headers values
func NewPcloudPlacementgroupsMembersPostBadRequest() *PcloudPlacementgroupsMembersPostBadRequest {
	return &PcloudPlacementgroupsMembersPostBadRequest{}
}

/* PcloudPlacementgroupsMembersPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPlacementgroupsMembersPostBadRequest struct {
	Payload *models.Error
}

func (o *PcloudPlacementgroupsMembersPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersPostBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudPlacementgroupsMembersPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsMembersPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsMembersPostNotFound creates a PcloudPlacementgroupsMembersPostNotFound with default headers values
func NewPcloudPlacementgroupsMembersPostNotFound() *PcloudPlacementgroupsMembersPostNotFound {
	return &PcloudPlacementgroupsMembersPostNotFound{}
}

/* PcloudPlacementgroupsMembersPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPlacementgroupsMembersPostNotFound struct {
	Payload *models.Error
}

func (o *PcloudPlacementgroupsMembersPostNotFound) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersPostNotFound  %+v", 404, o.Payload)
}
func (o *PcloudPlacementgroupsMembersPostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsMembersPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsMembersPostConflict creates a PcloudPlacementgroupsMembersPostConflict with default headers values
func NewPcloudPlacementgroupsMembersPostConflict() *PcloudPlacementgroupsMembersPostConflict {
	return &PcloudPlacementgroupsMembersPostConflict{}
}

/* PcloudPlacementgroupsMembersPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudPlacementgroupsMembersPostConflict struct {
	Payload *models.Error
}

func (o *PcloudPlacementgroupsMembersPostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersPostConflict  %+v", 409, o.Payload)
}
func (o *PcloudPlacementgroupsMembersPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsMembersPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsMembersPostUnprocessableEntity creates a PcloudPlacementgroupsMembersPostUnprocessableEntity with default headers values
func NewPcloudPlacementgroupsMembersPostUnprocessableEntity() *PcloudPlacementgroupsMembersPostUnprocessableEntity {
	return &PcloudPlacementgroupsMembersPostUnprocessableEntity{}
}

/* PcloudPlacementgroupsMembersPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudPlacementgroupsMembersPostUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudPlacementgroupsMembersPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersPostUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PcloudPlacementgroupsMembersPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsMembersPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsMembersPostInternalServerError creates a PcloudPlacementgroupsMembersPostInternalServerError with default headers values
func NewPcloudPlacementgroupsMembersPostInternalServerError() *PcloudPlacementgroupsMembersPostInternalServerError {
	return &PcloudPlacementgroupsMembersPostInternalServerError{}
}

/* PcloudPlacementgroupsMembersPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPlacementgroupsMembersPostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudPlacementgroupsMembersPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}/members][%d] pcloudPlacementgroupsMembersPostInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudPlacementgroupsMembersPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsMembersPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}