// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ContainerStateWaiting ContainerStateWaiting is a waiting state of a container.
//
// swagger:model ContainerStateWaiting
type ContainerStateWaiting struct {

	// Message regarding why the container is not yet running.
	Message string `json:"message,omitempty"`

	// (brief) reason the container is not yet running.
	Reason string `json:"reason,omitempty"`
}
