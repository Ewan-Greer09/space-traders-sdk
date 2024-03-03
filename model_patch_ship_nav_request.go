/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.   

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PatchShipNavRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchShipNavRequest{}

// PatchShipNavRequest struct for PatchShipNavRequest
type PatchShipNavRequest struct {
	FlightMode *ShipNavFlightMode `json:"flightMode,omitempty"`
}

// NewPatchShipNavRequest instantiates a new PatchShipNavRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchShipNavRequest() *PatchShipNavRequest {
	this := PatchShipNavRequest{}
	var flightMode ShipNavFlightMode = CRUISE
	this.FlightMode = &flightMode
	return &this
}

// NewPatchShipNavRequestWithDefaults instantiates a new PatchShipNavRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchShipNavRequestWithDefaults() *PatchShipNavRequest {
	this := PatchShipNavRequest{}
	var flightMode ShipNavFlightMode = CRUISE
	this.FlightMode = &flightMode
	return &this
}

// GetFlightMode returns the FlightMode field value if set, zero value otherwise.
func (o *PatchShipNavRequest) GetFlightMode() ShipNavFlightMode {
	if o == nil || IsNil(o.FlightMode) {
		var ret ShipNavFlightMode
		return ret
	}
	return *o.FlightMode
}

// GetFlightModeOk returns a tuple with the FlightMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchShipNavRequest) GetFlightModeOk() (*ShipNavFlightMode, bool) {
	if o == nil || IsNil(o.FlightMode) {
		return nil, false
	}
	return o.FlightMode, true
}

// HasFlightMode returns a boolean if a field has been set.
func (o *PatchShipNavRequest) HasFlightMode() bool {
	if o != nil && !IsNil(o.FlightMode) {
		return true
	}

	return false
}

// SetFlightMode gets a reference to the given ShipNavFlightMode and assigns it to the FlightMode field.
func (o *PatchShipNavRequest) SetFlightMode(v ShipNavFlightMode) {
	o.FlightMode = &v
}

func (o PatchShipNavRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchShipNavRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FlightMode) {
		toSerialize["flightMode"] = o.FlightMode
	}
	return toSerialize, nil
}

type NullablePatchShipNavRequest struct {
	value *PatchShipNavRequest
	isSet bool
}

func (v NullablePatchShipNavRequest) Get() *PatchShipNavRequest {
	return v.value
}

func (v *NullablePatchShipNavRequest) Set(val *PatchShipNavRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchShipNavRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchShipNavRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchShipNavRequest(val *PatchShipNavRequest) *NullablePatchShipNavRequest {
	return &NullablePatchShipNavRequest{value: val, isSet: true}
}

func (v NullablePatchShipNavRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchShipNavRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

