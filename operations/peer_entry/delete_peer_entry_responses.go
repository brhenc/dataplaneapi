// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package peer_entry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/models"
)

// DeletePeerEntryAcceptedCode is the HTTP code returned for type DeletePeerEntryAccepted
const DeletePeerEntryAcceptedCode int = 202

/*DeletePeerEntryAccepted Configuration change accepted and reload requested

swagger:response deletePeerEntryAccepted
*/
type DeletePeerEntryAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeletePeerEntryAccepted creates DeletePeerEntryAccepted with default headers values
func NewDeletePeerEntryAccepted() *DeletePeerEntryAccepted {

	return &DeletePeerEntryAccepted{}
}

// WithReloadID adds the reloadId to the delete peer entry accepted response
func (o *DeletePeerEntryAccepted) WithReloadID(reloadID string) *DeletePeerEntryAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete peer entry accepted response
func (o *DeletePeerEntryAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeletePeerEntryAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeletePeerEntryNoContentCode is the HTTP code returned for type DeletePeerEntryNoContent
const DeletePeerEntryNoContentCode int = 204

/*DeletePeerEntryNoContent PeerEntry deleted

swagger:response deletePeerEntryNoContent
*/
type DeletePeerEntryNoContent struct {
}

// NewDeletePeerEntryNoContent creates DeletePeerEntryNoContent with default headers values
func NewDeletePeerEntryNoContent() *DeletePeerEntryNoContent {

	return &DeletePeerEntryNoContent{}
}

// WriteResponse to the client
func (o *DeletePeerEntryNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeletePeerEntryNotFoundCode is the HTTP code returned for type DeletePeerEntryNotFound
const DeletePeerEntryNotFoundCode int = 404

/*DeletePeerEntryNotFound The specified resource was not found

swagger:response deletePeerEntryNotFound
*/
type DeletePeerEntryNotFound struct {
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeletePeerEntryNotFound creates DeletePeerEntryNotFound with default headers values
func NewDeletePeerEntryNotFound() *DeletePeerEntryNotFound {

	return &DeletePeerEntryNotFound{}
}

// WithConfigurationVersion adds the configurationVersion to the delete peer entry not found response
func (o *DeletePeerEntryNotFound) WithConfigurationVersion(configurationVersion int64) *DeletePeerEntryNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete peer entry not found response
func (o *DeletePeerEntryNotFound) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete peer entry not found response
func (o *DeletePeerEntryNotFound) WithPayload(payload *models.Error) *DeletePeerEntryNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete peer entry not found response
func (o *DeletePeerEntryNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePeerEntryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeletePeerEntryDefault General Error

swagger:response deletePeerEntryDefault
*/
type DeletePeerEntryDefault struct {
	_statusCode int
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeletePeerEntryDefault creates DeletePeerEntryDefault with default headers values
func NewDeletePeerEntryDefault(code int) *DeletePeerEntryDefault {
	if code <= 0 {
		code = 500
	}

	return &DeletePeerEntryDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete peer entry default response
func (o *DeletePeerEntryDefault) WithStatusCode(code int) *DeletePeerEntryDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete peer entry default response
func (o *DeletePeerEntryDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete peer entry default response
func (o *DeletePeerEntryDefault) WithConfigurationVersion(configurationVersion int64) *DeletePeerEntryDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete peer entry default response
func (o *DeletePeerEntryDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete peer entry default response
func (o *DeletePeerEntryDefault) WithPayload(payload *models.Error) *DeletePeerEntryDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete peer entry default response
func (o *DeletePeerEntryDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePeerEntryDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
